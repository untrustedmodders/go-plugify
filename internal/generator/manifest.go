package generator

import (
	"errors"
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/untrustedmodders/go-plugify/manifest"
)

// typeTables collects the prototypes and enums encountered while converting
// exported functions. Each is written to the manifest once and named from every
// use site, so a callback taken by twenty functions is described once.
//
// Two Go types in different packages can share a short name, and a name is what
// a reference resolves on, so a clash is recorded rather than silently letting
// one definition stand in for the other.
type typeTables struct {
	prototypes map[string]*manifest.Method
	enums      map[string]*manifest.Enum
	clashes    []string
}

func newTypeTables() *typeTables {
	return &typeTables{
		prototypes: map[string]*manifest.Method{},
		enums:      map[string]*manifest.Enum{},
	}
}

// addPrototype files a callback signature and returns the name to refer to it by.
func (t *typeTables) addPrototype(prototype *manifest.Method) string {
	existing, found := t.prototypes[prototype.Name]
	if !found {
		t.prototypes[prototype.Name] = prototype
		return prototype.Name
	}
	if !samePrototype(existing, prototype) {
		t.clash("prototype", prototype.Name)
	}
	return prototype.Name
}

// addEnum files an enumeration and returns the name to refer to it by.
func (t *typeTables) addEnum(enum *manifest.Enum) string {
	existing, found := t.enums[enum.Name]
	if !found {
		t.enums[enum.Name] = enum
		return enum.Name
	}
	if !sameEnum(existing, enum) {
		t.clash("enum", enum.Name)
	}
	return enum.Name
}

func (t *typeTables) clash(kind, name string) {
	clash := fmt.Sprintf("%s %s", kind, name)
	for _, seen := range t.clashes {
		if seen == clash {
			return
		}
	}
	t.clashes = append(t.clashes, clash)
}

// err reports the name clashes found during conversion, if any.
func (t *typeTables) err() error {
	if len(t.clashes) == 0 {
		return nil
	}
	return fmt.Errorf(
		"conflicting definitions share a name, so a reference to one would be ambiguous: %v",
		t.clashes)
}

func (t *typeTables) sortedPrototypes() []*manifest.Method {
	if len(t.prototypes) == 0 {
		return nil
	}
	out := make([]*manifest.Method, 0, len(t.prototypes))
	for _, p := range t.prototypes {
		out = append(out, p)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

func (t *typeTables) sortedEnums() []*manifest.Enum {
	if len(t.enums) == 0 {
		return nil
	}
	out := make([]*manifest.Enum, 0, len(t.enums))
	for _, e := range t.enums {
		out = append(out, e)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

// PopulateManifest describes the exported functions in the manifest, setting
// its methods along with the prototype and enum tables they refer to.
func PopulateManifest(m *manifest.Manifest, funcs []ExportedFunction) error {
	tables := newTypeTables()

	methods := make([]manifest.Method, len(funcs))
	for i, f := range funcs {
		methods[i] = manifest.Method{
			Name:        f.ExportName,
			FuncName:    f.FuncName,
			Description: f.Description,
			ParamTypes:  tables.convertParams(f.Params),
			RetType:     tables.convertReturnType(ParamInfo{"", f.ReturnType, f.ReturnType.Description}),
		}
	}

	if err := errors.Join(tables.err(), duplicateExports(methods)); err != nil {
		return err
	}

	sort.Slice(methods, func(i, j int) bool { return methods[i].Name < methods[j].Name })

	m.Methods = methods
	m.Prototypes = tables.sortedPrototypes()
	m.Enums = tables.sortedEnums()
	return nil
}

// duplicateExports reports an export name or generated symbol used more than
// once. Plugify refuses such a manifest at load, and cgo refuses a repeated
// symbol at link time, but neither says which Go declaration caused it; naming
// them here puts the problem next to the source that wrote it.
func duplicateExports(methods []manifest.Method) error {
	var problems []string
	seenName := make(map[string]bool, len(methods))
	seenSymbol := make(map[string]bool, len(methods))

	for _, m := range methods {
		if seenName[m.Name] {
			problems = append(problems, fmt.Sprintf("method %q is exported more than once", m.Name))
		}
		seenName[m.Name] = true

		if seenSymbol[m.FuncName] {
			problems = append(problems, fmt.Sprintf("symbol %q is generated more than once", m.FuncName))
		}
		seenSymbol[m.FuncName] = true
	}

	if len(problems) == 0 {
		return nil
	}

	sort.Strings(problems)
	return errors.New(strings.Join(slices.Compact(problems), "; "))
}

func (t *typeTables) convertParams(params []ParamInfo) []manifest.Property {
	result := make([]manifest.Property, len(params))
	for i, p := range params {
		result[i] = t.convertParamType(p, false)
	}
	return result
}

func (t *typeTables) convertParamType(p ParamInfo, ignoreRef bool) manifest.Property {
	ty := p.Type

	if ty.IsFunc {
		// Function parameter. The signature is filed under its own name, and its
		// parameters are converted first so nested types are filed too.
		prototype := &manifest.Method{
			Name:        ty.FuncSig.Name,
			Description: ty.FuncSig.Description,
			FuncName:    "_",
			ParamTypes:  t.convertParams(ty.FuncSig.Params),
			RetType:     t.convertReturnType(ParamInfo{"", ty.FuncSig.Return, ty.FuncSig.Return.Description}),
		}

		return manifest.Property{
			Type:        "function",
			Name:        p.Name,
			Description: p.Description,
			Ref:         ty.IsRef && !ignoreRef,
			Prototype:   t.addPrototype(prototype),
		}
	}

	if ty.IsArray && ty.ElemType != nil && ty.ElemType.IsEnum {
		// Array of enum: the element carries the enum, the array carries the type.
		elem := t.convertParamType(ParamInfo{p.Name, *ty.ElemType, p.Description}, ignoreRef)

		return manifest.Property{
			Name:        p.Name,
			Description: p.Description,
			Type:        ty.TypeString,
			Ref:         ty.IsRef && !ignoreRef,
			Enum:        elem.Enum,
		}
	}

	if ty.IsEnum {
		values := make([]manifest.Value, len(ty.EnumValues))
		for i, v := range ty.EnumValues {
			values[i] = manifest.Value{
				Name:        v.Name,
				Description: v.Description,
				Value:       v.Value,
			}
		}

		return manifest.Property{
			Name:        p.Name,
			Description: p.Description,
			Type:        ty.TypeString,
			Ref:         ty.IsRef,
			Enum: t.addEnum(&manifest.Enum{
				Name:        ty.EnumTypeName,
				Description: ty.Description,
				Values:      values,
			}),
		}
	}

	if ty.IsAlias {
		return manifest.Property{
			Name:        p.Name,
			Description: p.Description,
			Type:        ty.TypeString,
			Ref:         ty.IsRef,
			Alias: &manifest.Alias{
				Name:        ty.EnumTypeName,
				Description: ty.Description,
			},
		}
	}

	return manifest.Property{
		Name:        p.Name,
		Description: p.Description,
		Type:        ty.TypeString,
		Ref:         ty.IsRef,
	}
}

func (t *typeTables) convertReturnType(p ParamInfo) manifest.Property {
	return t.convertParamType(p, true)
}

// Definitions are compared on the parts that make up the type. Descriptions are
// left out: they mean nothing to a caller, and two spellings of the same type
// that differ only in wording are not a clash.

func sameProperty(a, b *manifest.Property) bool {
	return a.Type == b.Type && a.Ref == b.Ref &&
		a.Prototype == b.Prototype && a.Enum == b.Enum
}

func sameEnum(a, b *manifest.Enum) bool {
	if len(a.Values) != len(b.Values) {
		return false
	}
	for i := range a.Values {
		if a.Values[i].Name != b.Values[i].Name || a.Values[i].Value != b.Values[i].Value {
			return false
		}
	}
	return true
}

func samePrototype(a, b *manifest.Method) bool {
	if len(a.ParamTypes) != len(b.ParamTypes) {
		return false
	}
	if !sameProperty(&a.RetType, &b.RetType) {
		return false
	}
	for i := range a.ParamTypes {
		if !sameProperty(&a.ParamTypes[i], &b.ParamTypes[i]) {
			return false
		}
	}
	return true
}
