package generator

import (
	"go/ast"
	"go/constant"
	"go/types"
	"sort"
	"strings"

	"golang.org/x/tools/go/packages"
)

func findEnumValues(pkg *packages.Package, typeObj types.Object) []EnumValue {
	// Get the package where this type is defined
	typePkg := typeObj.Pkg()
	if typePkg == nil {
		return nil
	}

	// Get the type we're looking for
	enumType := typeObj.Type()
	enumTypeName := typeObj.Name()

	// If we have AST access, use it to find constants with explicit type annotations
	if pkg != nil {
		enumValues := findEnumValuesFromAST(pkg, enumTypeName)
		if len(enumValues) > 0 {
			return enumValues
		}
	}

	return findEnumValuesInScope(pkg, typePkg, enumType)
}

// findEnumValuesFromAST extracts enum values by examining the AST to find constants
// with explicit type annotations matching the enum type name
func findEnumValuesFromAST(pkg *packages.Package, enumTypeName string) []EnumValue {
	var enumValues []EnumValue

	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			for _, spec := range genDecl.Specs {
				valueSpec, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}

				// Check if this constant has an explicit type annotation
				if valueSpec.Type == nil {
					continue
				}

				// Get the type name from the AST
				var typeName string
				switch t := valueSpec.Type.(type) {
				case *ast.Ident:
					typeName = t.Name
				default:
					continue
				}

				// Only include constants explicitly declared with this enum type
				if typeName != enumTypeName {
					continue
				}

				// Extract all constants in this spec
				for _, name := range valueSpec.Names {
					// Get the constant object from type info
					obj := pkg.TypesInfo.ObjectOf(name)
					if obj == nil {
						continue
					}

					constObj, ok := obj.(*types.Const)
					if !ok {
						continue
					}

					// Extract the constant value
					val := constObj.Val()
					if val == nil {
						continue
					}

					// Convert to int64
					var intValue int64
					switch val.Kind() {
					case constant.Int:
						if i, ok := constant.Int64Val(val); ok {
							intValue = i
						} else {
							continue
						}
					default:
						continue
					}

					// Get description
					var description string
					if valueSpec.Doc != nil {
						docComment := parseDocComment(valueSpec.Doc)
						description = docComment.Description
					}
					if description == "" && valueSpec.Comment != nil {
						comment := valueSpec.Comment.Text()
						comment = strings.TrimSpace(strings.TrimPrefix(comment, "//"))
						comment = strings.TrimSpace(strings.TrimPrefix(comment, "/*"))
						comment = strings.TrimSpace(strings.TrimSuffix(comment, "*/"))
						description = comment
					}

					enumValues = append(enumValues, EnumValue{
						Name:        name.Name,
						Value:       intValue,
						Description: description,
					})
				}
			}
		}
	}

	// Sort by value to ensure consistent ordering
	sort.Slice(enumValues, func(i, j int) bool {
		return enumValues[i].Value < enumValues[j].Value
	})

	return enumValues
}

func findEnumValuesInScope(pkg *packages.Package, typePkg *types.Package, enumType types.Type) []EnumValue {
	var enumValues []EnumValue

	// Fallback: Iterate through all objects in the package scope
	scope := typePkg.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)

		// Check if it's a constant
		constObj, ok := obj.(*types.Const)
		if !ok {
			continue
		}

		// Check if the constant's type matches our enum type
		if !types.Identical(constObj.Type(), enumType) {
			continue
		}

		// Extract the constant value
		val := constObj.Val()
		if val == nil {
			continue
		}

		// Convert to int64
		var intValue int64
		switch val.Kind() {
		case constant.Int:
			// Get the int64 value
			if i, ok := constant.Int64Val(val); ok {
				intValue = i
			} else {
				// Value too large for int64, skip
				continue
			}
		default:
			// Not an integer constant, skip
			continue
		}

		// Try to get comment for this constant
		var description string
		if pkg != nil {
			description = findConstComment(pkg, constObj.Name())
		}

		enumValues = append(enumValues, EnumValue{
			Name:        constObj.Name(),
			Value:       intValue,
			Description: description,
		})
	}

	// Sort by value to ensure consistent ordering
	sort.Slice(enumValues, func(i, j int) bool {
		return enumValues[i].Value < enumValues[j].Value
	})

	return enumValues
}
