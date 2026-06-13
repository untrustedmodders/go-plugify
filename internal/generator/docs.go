package generator

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/packages"
)

// DocComment represents parsed documentation from comments
type DocComment struct {
	Description  string
	ParamDescs   map[string]string // param name -> description
	ReturnDesc   string
	EnumValueMap map[string]string // enum value name -> description
}

// parseDocComment parses doxygen-style comments and extracts @param, @return, @brief, etc.
func parseDocComment(commentGroup *ast.CommentGroup) DocComment {
	doc := DocComment{
		ParamDescs:   make(map[string]string),
		EnumValueMap: make(map[string]string),
	}

	if commentGroup == nil {
		return doc
	}

	var descriptionLines []string
	var briefDesc string
	inDescription := true

	for _, comment := range commentGroup.List {
		text := strings.TrimSpace(strings.TrimPrefix(comment.Text, "//"))
		text = strings.TrimSpace(strings.TrimPrefix(text, "/*"))
		text = strings.TrimSpace(strings.TrimSuffix(text, "*/"))
		text = strings.TrimSpace(strings.TrimPrefix(text, "*"))

		// Skip plugify:export directives
		if strings.HasPrefix(text, "plugify:export") {
			continue
		}

		// Parse @brief tag
		if strings.HasPrefix(text, "@brief") {
			inDescription = false
			parts := strings.SplitN(text, "@brief", 2)
			if len(parts) == 2 {
				briefDesc = strings.TrimSpace(parts[1])
			}
			continue
		}

		// Parse @param tag
		if strings.HasPrefix(text, "@param") {
			inDescription = false
			parts := strings.Fields(text)
			if len(parts) >= 3 {
				paramName := parts[1]
				paramDesc := strings.Join(parts[2:], " ")
				doc.ParamDescs[paramName] = paramDesc
			}
			continue
		}

		// Parse @return tag
		if strings.HasPrefix(text, "@return") {
			inDescription = false
			parts := strings.SplitN(text, "@return", 2)
			if len(parts) == 2 {
				doc.ReturnDesc = strings.TrimSpace(parts[1])
			}
			continue
		}

		// Collect description lines
		if inDescription && text != "" {
			descriptionLines = append(descriptionLines, text)
		}
	}

	// Use @brief if provided, otherwise use the collected description lines
	if briefDesc != "" {
		doc.Description = briefDesc
	} else {
		doc.Description = strings.Join(descriptionLines, " ")
	}

	return doc
}

// findTypeDelegateDoc finds and parses doxygen-style documentation for a delegate type
func findTypeDelegateDoc(pkg *packages.Package, typeName string) DocComment {
	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				if typeSpec.Name.Name == typeName {
					// Check for doc comment above the type
					if genDecl.Doc != nil {
						return parseDocComment(genDecl.Doc)
					}
				}
			}
		}
	}
	return DocComment{
		ParamDescs:   make(map[string]string),
		EnumValueMap: make(map[string]string),
	}
}

// findTypeComment finds the comment for a type declaration in AST
func findTypeComment(pkg *packages.Package, typeName string) string {
	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				if typeSpec.Name.Name == typeName {
					// First check for doc comment above the type
					if genDecl.Doc != nil {
						docComment := parseDocComment(genDecl.Doc)
						if docComment.Description != "" {
							return docComment.Description
						}
					}

					// Then check for inline comment on the same line
					if typeSpec.Comment != nil {
						comment := typeSpec.Comment.Text()
						comment = strings.TrimSpace(strings.TrimPrefix(comment, "//"))
						comment = strings.TrimSpace(strings.TrimPrefix(comment, "/*"))
						comment = strings.TrimSpace(strings.TrimSuffix(comment, "*/"))
						return comment
					}
				}
			}
		}
	}
	return ""
}

// findConstComment finds the comment for a constant declaration in AST
func findConstComment(pkg *packages.Package, constName string) string {
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

				for _, name := range valueSpec.Names {
					if name.Name == constName {
						// Check for comment on the same line or doc comment
						if valueSpec.Doc != nil {
							docComment := parseDocComment(valueSpec.Doc)
							return docComment.Description
						}
						if valueSpec.Comment != nil {
							comment := valueSpec.Comment.Text()
							comment = strings.TrimSpace(strings.TrimPrefix(comment, "//"))
							comment = strings.TrimSpace(strings.TrimPrefix(comment, "/*"))
							comment = strings.TrimSpace(strings.TrimSuffix(comment, "*/"))
							return comment
						}
					}
				}
			}
		}
	}
	return ""
}
