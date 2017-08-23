package builder

import (
	"github.com/michaelwmerritt/project-builder/model"
	"os/exec"
)

type AntBuilder struct{}

func NewAntBuilder() AntBuilder {
	return AntBuilder{}
}

func (antBuilder AntBuilder) Build(release model.Release, modules []model.Module, path string) (err error) {
	buildModuleMap := make(map[*model.Module]bool)
	moduleDependencyMap := make(map[string]bool)

	for _, module := range modules {
		if module.BuildModule {
			buildModuleMap[&module] = module.BuilderDependencies
		} else {
			if !module.BuildInfrastructure {
				moduleDependencyMap[module.Id] = true
				for _, dependency := range module.Dependencies {
					moduleDependencyMap[dependency] = true
				}
			}
		}
	}

	for key, requiresBuildDependencies := range buildModuleMap {
		dependencies := "-Dui.dependencies="
		if requiresBuildDependencies {
			dependencies += antBuilder.createDependencyString(moduleDependencyMap)
		}
		cmd := exec.Command("ant", "-f", key.CreateBuildFilePath(path), dependencies)
		err = executeCommand(cmd)
	}
	return
}

func (antBuilder AntBuilder) createDependencyString(dependencyMap map[string]bool) (dependencies string) {
	for key := range dependencyMap {
		dependencies = dependencies + "," + key
	}
	return
}
