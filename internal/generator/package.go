package generator

import (
	"fmt"
	"go/token"
	"os"
	"path/filepath"
	"slices"

	"golang.org/x/tools/go/packages"
)

type PackageData struct {
	Pkg  *packages.Package
	Fset *token.FileSet
}

type LoadPackageFilter struct {
	// Package full path, e.g. `github.com/untrustedmodders/go-plugify` or `github.com/untrustedmodders/go-plugify/manifest`
	Packages []string

	// File path patterns, e.g. `**/*.go`
	PathPatterns []string
}

func ParsePackages(patterns string, verbose bool, filter *LoadPackageFilter) ([]*Package, *Package, error) {
	fset := token.NewFileSet()

	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax |
			packages.NeedTypes | packages.NeedTypesInfo | packages.NeedImports | packages.NeedDeps | packages.LoadAllSyntax |
			packages.LoadImports | packages.NeedExportFile,

		Tests:      false,
		Env:        append(os.Environ(), "CGO_ENABLED=1"),
		BuildFlags: []string{"-tags", "cgo"},

		Fset: fset,
	}

	pkgs, err := packages.Load(cfg, patterns)

	if err != nil {
		return nil, nil, fmt.Errorf("error loading package: %w", err)
	}

	if len(pkgs) == 0 {
		return nil, nil, fmt.Errorf("no packages found in patterns: %s", patterns)
	}

	var mainGen *Package
	var gens []*Package
	for _, pkg := range pkgs {

		if filter != nil {
			if len(filter.Packages) > 0 && !slices.Contains(filter.Packages, pkg.PkgPath) {
				if verbose {
					fmt.Printf("Skipping package: %s (not in filter's package list)\n", pkg.Name)
				}

				continue
			}

			var matched bool
			for _, pattern := range filter.PathPatterns {
				match, err := filepath.Match(pattern, pkg.PkgPath)
				if err != nil {
					return nil, nil, fmt.Errorf("failed to match pattern '%s': %w", pattern, err)
				}

				if match {
					matched = true
					break
				}
			}
			if !matched {
				if verbose {
					fmt.Printf("Skipping package: %s (not in filter's pattern list)\n", pkg.Name)
				}

				continue
			}
		}

		gen := &Package{Pkg: pkg, Fset: fset, verbose: verbose}
		if pkg.Name == "main" {
			if mainGen != nil {
				return nil, nil, fmt.Errorf("multiple main files unsupported: main files: [%s %s]", mainGen.Pkg.Dir, pkg.Dir)
			}
			mainGen = gen
			continue
		}

		gens = append(gens, gen)
	}

	if mainGen == nil {
		return nil, nil, fmt.Errorf("no main package found in patterns: %s", patterns)
	}

	return gens, mainGen, nil
}
