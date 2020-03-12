package clean

import (
	"fmt"
	"strings"

	"github.com/go-openapi/analysis"
	"github.com/go-openapi/spec"
)

// RemoveUnusedDefinitions removes defenitions if they aren't found in refs
func RemoveUnusedDefinitions(swspec *spec.Swagger) {
	aSpec := analysis.New(swspec)
	foundDefRefsNames := aSpec.AllDefinitionReferences()
	foundDefNames := make([]string, len(foundDefRefsNames), len(foundDefRefsNames))
	for i, name := range foundDefRefsNames {
		fmt.Println(name)
		foundDefNames[i] = strings.Replace(name, "#/definitions/", "", 1)
	}
	fmt.Println("----------------------------------------------")
	for name := range swspec.Definitions {
		fmt.Println(name)
		if inStringList(foundDefNames, name) {
			continue
		}
		delete(swspec.Definitions, name)
	}
}

func inStringList(list []string, value string) bool {
	for _, s := range list {
		if s == value {
			return true
		}
	}
	return false
}
