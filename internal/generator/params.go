package generator

import (
	"fmt"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/packages"
)

type ExportedFunction struct {
	ExportName  string
	FuncName    string
	Params      []ParamInfo
	ReturnType  TypeInfo
	Description string

	originalFuncName string
	packageImport    packageImport
}

type ParamInfo struct {
	Name        string
	Type        TypeInfo
	Description string
}

type EnumValue struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Value       int64  `json:"value"`
}

type TypeInfo struct {
	TypeString string
	GoBaseType string
	IsRef      bool
	IsFunc     bool
	IsEnum     bool
	IsAlias    bool
	IsArray    bool

	AliasTypeName string
	EnumTypeName  string
	EnumValues    []EnumValue
	ElemType      *TypeInfo
	FuncSig       *FuncSignature
	Description   string

	packageImport packageImport
}

type FuncSignature struct {
	Name        string
	Params      []ParamInfo
	Return      TypeInfo
	Description string
}

type mappedType int

const (
	mappedTypeInvalid mappedType = iota
	mappedTypeBasic
	mappedTypePointer
	mappedTypeSlice
	mappedTypeInterface
	mappedTypeFunc
)

type paramError struct {
	fileName       string
	fileLine       int
	fileLineOffset int

	funcName string
	paramNum int
	varName  string
	varType  string

	err error
}

func (p *paramError) Error() string {
	return fmt.Sprintf("%s:%d:%d func %s: argument %d (%s): %v", p.fileName, p.fileLine, p.fileLineOffset, p.fileName, p.paramNum, p.varName, p.err)
}

func (g *Package) newParamError(p token.Pos, funcName string, paramNum int, varName string, varType string, err error) *paramError {
	pos := g.Fset.Position(p)
	return &paramError{
		fileName:       pos.Filename,
		fileLine:       pos.Line,
		fileLineOffset: pos.Column,

		funcName: funcName,
		paramNum: paramNum,
		varName:  varName,
		varType:  varType,

		err: err,
	}
}

func (p *paramError) Add(err string) *paramError {
	p.err = fmt.Errorf("%s: %w", err, p.err)
	return p
}
func (g *Package) checkExportedType(v *types.Var, tn *types.TypeName) *paramError {
	pkg := tn.Pkg()
	if pkg == nil || pkg.Name() == "main" {
		return nil
	}
	if tn.Exported() {
		return nil
	}
	return g.newParamError(
		v.Pos(), "", 0, v.Name(), tn.Name(),
		fmt.Errorf("cannot use unexported type %s from package %s in an exported function signature", tn.Name(), pkg.Path()),
	)
}

func (g *Package) extractParams(sig *types.Signature, info *types.Info, pkg *packages.Package) ([]ParamInfo, *paramError) {
	params := sig.Params()

	if params == nil {
		return nil, nil
	}

	result := make([]ParamInfo, params.Len())
	for i := 0; i < params.Len(); i++ {
		param := params.At(i)

		t, _, err := g.mapTypeInfo(param, param.Type(), "", info, pkg, false)
		if err != nil {
			return nil, err
		}

		result[i] = ParamInfo{
			Name: param.Name(),
			Type: t,
		}
	}

	return result, nil
}

func (g *Package) extractReturnType(results *types.Tuple, info *types.Info, pkg *packages.Package) (TypeInfo, mappedType, *paramError) {
	if results == nil || results.Len() == 0 {
		return TypeInfo{TypeString: "void"}, mappedTypeInvalid, nil
	}

	res := results.At(0)
	return g.mapTypeInfo(res, res.Type(), "", info, pkg, false)
}
func (g *Package) mapTypeInfo(v *types.Var, bt types.Type, typeName string, info *types.Info, pkg *packages.Package, isRef bool) (TypeInfo, mappedType, *paramError) {
	switch t := bt.(type) {

	// Handle pointer types - set isRef and unwrap
	case *types.Pointer:
		typeInfo, _, err := g.mapTypeInfo(v, t.Elem(), "", info, pkg, true)
		if err != nil {
			return TypeInfo{}, mappedTypeInvalid, err.Add("failed to map pointer type:")
		}

		typeInfo.GoBaseType = "*" + typeInfo.GoBaseType
		return typeInfo, mappedTypePointer, nil

	// Handle interface{} (any) - check this before slices since []any needs to detect any
	case *types.Interface:
		// Check if it's the empty interface (any)
		if t.NumMethods() == 0 {
			return TypeInfo{
				TypeString: "any",
				GoBaseType: "any",
				IsRef:      isRef,
			}, mappedTypeInterface, nil
		}
		// Non-empty interfaces are treated as pointers
		return TypeInfo{
			TypeString: getPtrType(),
			GoBaseType: "any",
			IsRef:      isRef,
		}, mappedTypePointer, nil

	// Handle slices/arrays
	case *types.Slice:
		// check is it ref?
		elemType, _, err := g.mapTypeInfo(v, t.Elem(), "", info, pkg, false)
		if err != nil {
			return TypeInfo{}, mappedTypeInvalid, err.Add("failed to map slice type:")
		}

		return TypeInfo{
			TypeString: elemType.TypeString + "[]",
			GoBaseType: "[]" + elemType.GoBaseType,
			IsRef:      isRef,
			IsArray:    true,
			ElemType:   &elemType,
		}, mappedTypeSlice, nil

	// Handle function types
	case *types.Signature:
		params, err := g.extractParams(t, info, pkg)
		if err != nil {
			return TypeInfo{}, mappedTypeInvalid, err.Add("failed to map func param")
		}
		retType, _, err := g.extractReturnType(t.Results(), info, pkg)
		if err != nil {
			return TypeInfo{}, mappedTypeInvalid, err.Add("failed to map func return")
		}

		// Get type-level documentation for delegate with doxygen parsing
		var delegateDoc DocComment
		if pkg != nil {
			if typeName == "" {
				typeName = t.String()
			}
			delegateDoc = findTypeDelegateDoc(pkg, typeName)
		}

		// Apply parameter descriptions from delegate documentation
		for i := range params {
			if desc, ok := delegateDoc.ParamDescs[params[i].Name]; ok {
				params[i].Description = desc
			}
		}

		// Apply return description from delegate documentation
		if delegateDoc.ReturnDesc != "" {
			retType.Description = delegateDoc.ReturnDesc
		}

		funcName := "callback"
		// Try to get function type name from context if available
		return TypeInfo{
			TypeString: "function",
			GoBaseType: "any",
			IsRef:      isRef,
			IsFunc:     true,
			FuncSig: &FuncSignature{
				Name:        funcName,
				Params:      params,
				Return:      retType,
				Description: delegateDoc.Description,
			},
		}, mappedTypeFunc, err

	case *types.Alias:
		obj := t.Obj()

		if err := g.checkExportedType(v, t.Obj()); err != nil {
			return TypeInfo{}, mappedTypeInvalid, err
		}

		aliasName := obj.Name()
		baseType := t.Rhs()

		typeInfo, mappedType, err := g.mapTypeInfo(v, baseType, aliasName, info, pkg, isRef)
		if err != nil {
			return TypeInfo{}, mappedTypeInvalid, err.Add("failed to map alias type")
		}

		switch mappedType {
		case mappedTypeFunc:
			typeInfo.FuncSig.Name = aliasName
		case mappedTypeBasic, mappedTypeSlice:
			typeInfo.EnumTypeName = aliasName

			if pkg != nil {
				typeInfo.Description = findTypeComment(pkg, typeInfo.EnumTypeName)
			}

			typeInfo.IsAlias = true
		}

		return typeInfo, mappedType, nil
	case *types.Named:
		obj := t.Obj()

		if err := g.checkExportedType(v, obj); err != nil {
			return TypeInfo{}, mappedTypeInvalid, err
		}

		typeName := obj.Name()
		baseType := t.Underlying()

		typeInfo, mappedType, err := g.mapTypeInfo(v, baseType, typeName, info, pkg, isRef)
		if err != nil {
			return TypeInfo{}, mappedTypeInvalid, err.Add("failed to map named type")
		}

		switch mappedType {
		case mappedTypeFunc:
			typeInfo.FuncSig.Name = typeName
		case mappedTypeBasic, mappedTypeSlice, mappedTypeInterface:
			typeInfo.EnumTypeName = typeName

			if pkg != nil {
				typeInfo.Description = findTypeComment(pkg, typeInfo.EnumTypeName)
			}

			typeInfo.IsEnum = true
			typeInfo.EnumValues = findEnumValues(pkg, t.Obj())
		}
		typeInfo.packageImport = g.prepareImport(obj)

		return typeInfo, mappedType, nil

	case *types.Struct:
		switch t.NumFields() {
		case 1:
			field0 := t.Field(0)
			if field0.Name() == "M" && field0.Type().Underlying().String() == "[4][4]float32" {
				return TypeInfo{
					TypeString: "mat4x4",
					IsRef:      isRef,
				}, mappedTypeBasic, nil
			}
		case 2:
			field0 := t.Field(0)
			field1 := t.Field(1)

			if field0.Name() == "X" && field1.Name() == "Y" &&
				field0.Type().Underlying().String() == "float32" && field1.Type().Underlying().String() == "float32" {
				return TypeInfo{
					TypeString: "vec2",
					IsRef:      isRef,
				}, mappedTypeBasic, nil
			}
		case 3:
			field0 := t.Field(0)
			field1 := t.Field(1)
			field2 := t.Field(2)

			if field0.Name() == "X" && field1.Name() == "Y" && field2.Name() == "Z" &&
				field0.Type().Underlying().String() == "float32" && field1.Type().Underlying().String() == "float32" && field2.Type().Underlying().String() == "float32" {
				return TypeInfo{
					TypeString: "vec3",
					IsRef:      isRef,
				}, mappedTypeBasic, nil
			}
		case 4:
			field0 := t.Field(0)
			field1 := t.Field(1)
			field2 := t.Field(2)
			field3 := t.Field(3)

			if field0.Name() == "X" && field1.Name() == "Y" && field2.Name() == "Z" && field3.Name() == "W" &&
				field0.Type().Underlying().String() == "float32" && field1.Type().Underlying().String() == "float32" && field2.Type().Underlying().String() == "float32" && field3.Type().Underlying().String() == "float32" {
				return TypeInfo{
					TypeString: "vec4",
					IsRef:      isRef,
				}, mappedTypeBasic, nil
			}
		}

		return TypeInfo{
			TypeString: getPtrType(),
			GoBaseType: t.Underlying().String(),
			IsRef:      isRef,
		}, mappedTypePointer, nil
	// Handle basic type
	case *types.Basic:
		return TypeInfo{
			TypeString: mapBasicType(t),
			GoBaseType: t.Name(),
			IsRef:      isRef,
		}, mappedTypeBasic, nil
	// Default - unknown type
	default:
		return TypeInfo{
			TypeString: getPtrType(),
			IsRef:      isRef,
		}, mappedTypePointer, nil
	}
}
