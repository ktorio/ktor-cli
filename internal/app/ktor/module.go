package ktor

import (
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/network"
	"slices"
	"strings"
)

const SerPluginId = "org.jetbrains.kotlin.plugin.serialization"
const SerPluginKotlinId = "plugin.serialization"
const KotlinJvmPluginId = "org.jetbrains.kotlin.jvm"
const KmpPluginId = "org.jetbrains.kotlin.multiplatform"

const ktorGroup = "io.ktor"

type MavenCoords struct {
	Artifact, Group, Version string
	IsTest                   bool
}

type GradlePlugin struct {
	Id              string
	IsSerialization bool
}

type ModuleResult int

const (
	ModuleNotFound ModuleResult = iota
	ModuleAmbiguity
	SimilarModulesFound
	ModuleFound
)

func ParseMavenCoords(s string) (MavenCoords, bool) {
	parts := strings.Split(s, ":")

	if len(parts) == 1 {
		return MavenCoords{Artifact: parts[0]}, true
	}

	if len(parts) == 2 {
		return MavenCoords{Group: parts[0], Artifact: parts[1]}, true
	}

	if len(parts) == 3 {
		return MavenCoords{Group: parts[0], Artifact: parts[1], Version: parts[2]}, true
	}

	return MavenCoords{}, false
}

func (mc *MavenCoords) String() string {
	return fmt.Sprintf("%s:%s", mc.Group, mc.Artifact)
}

func (mc *MavenCoords) RoughlySame(other MavenCoords) bool {
	if mc.Group != other.Group {
		return false
	}

	if strings.HasPrefix(mc.Artifact, other.Artifact) || strings.HasPrefix(other.Artifact, mc.Artifact) {
		return true
	}

	return false
}

func DependentPlugins(mc MavenCoords) []GradlePlugin {
	var plugs []GradlePlugin

	if strings.HasPrefix(mc.Artifact, "ktor-serialization-kotlinx") {
		plugs = append(plugs, GradlePlugin{Id: "org.jetbrains.kotlin.plugin.serialization", IsSerialization: true})
	}

	return plugs
}

func FindModule(artifacts []network.Artifact) (coords MavenCoords, result ModuleResult, candidates []MavenCoords) {
	if len(artifacts) == 0 {
		result = ModuleNotFound
	} else {
		var exactArtifacts []network.Artifact
		var similarArtifacts []network.Artifact
		for _, a := range artifacts {
			if a.Distance == 0 {
				exactArtifacts = append(exactArtifacts, a)
			} else {
				similarArtifacts = append(similarArtifacts, a)
			}
		}

		if len(exactArtifacts) == 1 {
			result = ModuleFound
			coords = MavenCoords{
				Artifact: exactArtifacts[0].Name,
				Group:    exactArtifacts[0].Group,
				IsTest:   exactArtifacts[0].IsTest,
			}
		} else if len(exactArtifacts) > 0 {
			result = ModuleAmbiguity
			sortArtifacts(exactArtifacts)

			for _, a := range exactArtifacts {
				candidates = append(candidates, MavenCoords{
					Artifact: a.Name,
					Group:    a.Group,
					IsTest:   a.IsTest,
				})
			}
		} else {
			sortArtifacts(similarArtifacts)
			result = SimilarModulesFound

			for _, a := range similarArtifacts {
				candidates = append(candidates, MavenCoords{
					Artifact: a.Name,
					Group:    a.Group,
					IsTest:   a.IsTest,
				})
			}
		}
	}

	return
}

var prefixes = []string{"ktor-server-", "ktor-client-", "ktor-"}

func sortArtifacts(artifacts []network.Artifact) {
	slices.SortFunc(artifacts, func(x, y network.Artifact) int {
		if x.Distance == y.Distance {
			var xIndex, yIndex int
			for i, pr := range prefixes {
				if strings.HasPrefix(x.Name, pr) {
					xIndex = i
					break
				}
			}

			for i, pr := range prefixes {
				if strings.HasPrefix(y.Name, pr) {
					yIndex = i
					break
				}
			}

			return xIndex - yIndex
		} else {
			return x.Distance - y.Distance
		}
	})
}
