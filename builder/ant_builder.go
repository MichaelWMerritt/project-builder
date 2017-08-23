package builder

import (
	"github.com/michaelwmerritt/project-builder/model"
	"os/exec"
	"strings"
)

type AntBuilder struct{}

func NewAntBuilder() AntBuilder {
	return AntBuilder{}
}

func (antBuilder AntBuilder) Build(release model.Release, modules []model.Module, path string) (err error) {
	buildModuleMap := make(map[string]model.Module)
	moduleDependencyMap := make(map[string]bool)

	for _, module := range modules {
		if module.BuildModule {
			buildModuleMap[module.Id] = module
		} else {
			if !module.BuildInfrastructure {
				moduleDependencyMap[module.Id] = true
				for _, dependency := range module.Dependencies {
					moduleDependencyMap[dependency] = true
				}
			}
		}
	}

	for _, module := range buildModuleMap {
		customerDependencies := ""
		uiDependencies := ""
		if module.BuilderDependencies {
			keys := make([]string, 0, len(moduleDependencyMap))
			newKeys := make([]string, 0, len(moduleDependencyMap) - 1)
			customerDependency := ""
			for k := range moduleDependencyMap {
				keys = append(keys, k)
			}
			for i := 0; i < len(keys); i++ {
				if i == 0 {
					customerDependency = keys[0]
				} else {
					newKeys = append(newKeys, keys[i])
				}
			}

			customerDependencies = "-Dcustomer.dependencies=" + customerDependency
			uiDependencies = antBuilder.createDependencyString("-Dui.dependencies=", newKeys)
		}
		cmd := exec.Command("ant", "-f", module.CreateBuildFilePath(path), customerDependencies, uiDependencies)
		err = executeCommand(cmd)
	}
	return
}

func (antBuilder AntBuilder) createDependencyString(tag string, keys []string) string {
	 return tag + strings.Join(keys, ",")
}
