package plugify

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/untrustedmodders/go-plugify/internal/generator"
	"github.com/untrustedmodders/go-plugify/manifest"
)

type GenerateParams struct {
	Version     string
	Description string
	Author      string
	Website     string
	License     string

	Platforms    []string
	Dependencies []string
	Conflicts    []string
}

type LoadPackageFilter struct {
	// Package full path, e.g. `github.com/untrustedmodders/go-plugify` or `github.com/untrustedmodders/go-plugify/manifest`
	Packages []string

	// File path patterns, e.g. `**/*.go`
	PathPatterns []string
}

func Generate(
	manifestName string,
	pluginName string,
	entry string,

	verbose bool,
	params GenerateParams,
	packageFilter *LoadPackageFilter,
) error {
	gens, mainGen, err := generator.ParsePackages("./...", verbose, (*generator.LoadPackageFilter)(packageFilter))
	if err != nil {
		return err
	}

	var exports []generator.ExportedFunction
	for _, gen := range gens {
		pkgExports, err := gen.Generate(mainGen)
		if err != nil {
			return fmt.Errorf("failed to generate package %s: %w", gen.Pkg.Name, err)
		}

		exports = append(exports, pkgExports...)
	}

	pkgExports, err := mainGen.GenerateMain()
	if err != nil {
		return fmt.Errorf("failed to generate main package: %w", err)
	}

	exports = append(exports, pkgExports...)

	if pluginName == "" {
		pluginName = mainGen.Pkg.Name
	}

	pluginEntry := entry
	if pluginEntry == "" {
		pluginEntry = pluginName
	}

	err = writeManifest(
		exports,
		manifestName,
		pluginName,
		params.Version,
		params.Description,
		params.Author,
		params.Website,
		params.License,
		params.Platforms,
		pluginEntry,
		params.Dependencies,
		params.Conflicts,
	)
	if err != nil {
		return fmt.Errorf("failed to write manifest: %w", err)
	}

	return nil
}

func writeManifest(
	exports []generator.ExportedFunction,
	output string,
	pluginName string,
	version string,
	desc string,
	author string,
	website string,
	license string,
	platforms []string,
	pluginEntry string,
	dependencies []string,
	conflicts []string,
) error {
	var pluginDependencies []manifest.Dependency
	for _, dependency := range dependencies {
		pluginDependencies = append(pluginDependencies, manifest.Dependency{Name: dependency})
	}

	var pluginConflicts []manifest.Conflict
	for _, conflict := range conflicts {
		pluginConflicts = append(pluginConflicts, manifest.Conflict{Name: conflict})
	}

	man := manifest.Manifest{
		Schema:       "https://raw.githubusercontent.com/untrustedmodders/plugify/refs/heads/main/schemas/plugin.schema.json",
		Name:         pluginName,
		Version:      version,
		Description:  desc,
		Author:       author,
		Website:      website,
		License:      license,
		Platforms:    platforms,
		Entry:        pluginEntry,
		Dependencies: pluginDependencies,
		Conflicts:    pluginConflicts,
		Language:     "golang",
		Methods:      generator.ConvertToManifestMethods(exports),
	}

	// Write JSON
	data, err := json.MarshalIndent(man, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	outputFile := output
	if outputFile == "" {
		outputFile = pluginName + ".pplugin"
	}

	err = os.WriteFile(outputFile, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing manifest file %s: %w", outputFile, err)
	}

	fmt.Printf("Generated manifest: %s (%d methods)\n", outputFile, len(man.Methods))
	return nil
}
