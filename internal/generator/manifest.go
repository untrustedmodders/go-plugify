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
			RetType:     convertReturnType(f.ReturnType),
		}
	}

	return methods
}

func convertParams(params []ParamInfo) []manifest.Property {
	result := make([]manifest.Property, len(params))
	for i, p := range params {
		result[i] = convertParamType(p.Type, false)
	}
	return result
}

func convertParamType(t TypeInfo, ignoreRef bool) manifest.Property {
	if t.IsFunc {
		// Function parameter with prototype
		return manifest.Property{
			Type: "function",
			//Name:        param.Name,
			//Description: param.Description,
			Ref: t.IsRef && !ignoreRef,
			Prototype: &manifest.Method{
				Name:        t.FuncSig.Name,
				FuncName:    "_",
				Description: t.FuncSig.Description,
				ParamTypes:  convertParams(t.FuncSig.Params),
				RetType:     convertReturnType(t.FuncSig.Return),
			},
		}
	}

	if t.IsArray && t.ElemType != nil && t.ElemType.IsEnum {
		// Array of enum parameter

		prop := convertParamType(*t.ElemType, ignoreRef)
		return manifest.Property{
			Type:       t.TypeString,
			Ref:        t.IsRef && !ignoreRef,
			Enumerator: prop.Enumerator,
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
			Type: t.TypeString,
			Ref:  t.IsRef,
			Enumerator: &manifest.EnumObject{
				Name:        t.EnumTypeName,
				Description: t.Description,
				Values:      enumValues,
			},
		}
	}

	return manifest.Property{
		Type: t.TypeString,
		Ref:  t.IsRef,
	}
}

func convertReturnType(t TypeInfo) manifest.Property {
	return convertParamType(t, true)
}
