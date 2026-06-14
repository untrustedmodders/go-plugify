package generator

import (
	"github.com/untrustedmodders/go-plugify/manifest"
)

func ConvertToManifestMethods(funcs []ExportedFunction) []manifest.Method {
	methods := make([]manifest.Method, len(funcs))
	for i, f := range funcs {
		methods[i] = manifest.Method{
			Name:        f.ExportName,
			FuncName:    f.FuncName,
			Description: f.Description,
			ParamTypes:  convertParams(f.Params),
			RetType:     convertReturnType(ParamInfo{"", f.ReturnType, f.Description}),
		}
	}

	return methods
}

func convertParams(params []ParamInfo) []manifest.Property {
	result := make([]manifest.Property, len(params))
	for i, p := range params {
		result[i] = convertParamType(p, false)
	}
	return result
}

func convertParamType(p ParamInfo, ignoreRef bool) manifest.Property {
	t := p.Type

	if t.IsFunc {
		// Function parameter with prototype
		return manifest.Property{
			Type: "function",
			//Name:        param.Name,
			//Description: param.Description,
			Ref: t.IsRef && !ignoreRef,
			Prototype: &manifest.Method{
				Name:        t.FuncSig.Name,
				Description: t.FuncSig.Description,
				FuncName:    "_",
				ParamTypes:  convertParams(t.FuncSig.Params),
				RetType:     convertReturnType(ParamInfo{"", t.FuncSig.Return, t.FuncSig.Description}),
			},
		}
	}

	if t.IsArray && t.ElemType != nil && t.ElemType.IsEnum {
		// Array of enum parameter

		prop := convertParamType(ParamInfo{p.Name, *p.Type.ElemType, p.Description}, ignoreRef)

		return manifest.Property{
			Name:        p.Name,
			Description: p.Description,
			Type:        t.TypeString,
			Ref:         t.IsRef && !ignoreRef,
			Enumerator:  prop.Enumerator,
		}
	}

	if t.IsEnum {
		var enumValues = make([]manifest.EnumValue, len(t.EnumValues))
		for i, v := range t.EnumValues {
			enumValues[i] = manifest.EnumValue{
				Name:        v.Name,
				Description: v.Description,
				Value:       v.Value,
			}
		}

		// Enum parameter
		return manifest.Property{
			Name:        p.Name,
			Description: p.Description,
			Type:        t.TypeString,
			Ref:         t.IsRef,
			Enumerator: &manifest.EnumObject{
				Name:        t.EnumTypeName,
				Description: t.Description,
				Values:      enumValues,
			},
		}
	} else if t.IsAlias {
		return manifest.Property{
			Name:        p.Name,
			Description: p.Description,
			Type:        t.TypeString,
			Ref:         t.IsRef,
			Alias: &manifest.Alias{
				Name:        t.EnumTypeName,
				Description: t.Description,
			},
		}
	}

	return manifest.Property{
		Name:        p.Name,
		Description: p.Description,
		Type:        t.TypeString,
		Ref:         t.IsRef,
	}
}

func convertReturnType(p ParamInfo) manifest.Property {
	return convertParamType(p, true)
}
