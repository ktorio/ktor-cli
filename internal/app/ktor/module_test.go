package ktor

import (
	"github.com/ktorio/ktor-cli/internal/app/network"
	"reflect"
	"testing"
)

type testCase struct {
	artifacts     []network.Artifact
	expMc         MavenCoords
	expResult     ModuleResult
	expCandidates []MavenCoords
}

func TestFindModule(t *testing.T) {
	var cases = []testCase{
		{
			artifacts: []network.Artifact{artifactOf("server-sse", 0)},
			expMc:     MavenCoords{Artifact: "ktor-server-sse", Group: MavenGroup}, expResult: ModuleFound,
		},
		{
			artifacts: []network.Artifact{artifactOf("client-js", 0)},
			expMc:     MavenCoords{Artifact: "ktor-client-js", Group: MavenGroup}, expResult: ModuleFound,
		},
		{
			artifacts: []network.Artifact{artifactOf("server-test-host", 0)},
			expMc:     MavenCoords{Artifact: "ktor-server-test-host", Group: MavenGroup}, expResult: ModuleFound,
		},
		{
			artifacts: []network.Artifact{testArtifactOf("client-mock", 0)},
			expMc:     MavenCoords{Artifact: "ktor-client-mock", Group: MavenGroup, IsTest: true}, expResult: ModuleFound,
		},
		{
			artifacts: []network.Artifact{
				artifactOf("client-content-negotiation", 0),
				artifactOf("server-content-negotiation", 0),
			},
			expMc: MavenCoords{}, expResult: ModuleAmbiguity, expCandidates: []MavenCoords{
				{Artifact: "ktor-server-content-negotiation", Group: MavenGroup},
				{Artifact: "ktor-client-content-negotiation", Group: MavenGroup},
			},
		},
		{
			artifacts: []network.Artifact{
				artifactOf("client-core", 0),
				artifactOf("server-core", 0),
			},
			expMc: MavenCoords{}, expResult: ModuleAmbiguity, expCandidates: []MavenCoords{
				{Artifact: "ktor-server-core", Group: MavenGroup},
				{Artifact: "ktor-client-core", Group: MavenGroup},
			},
		},
		{
			artifacts: []network.Artifact{
				artifactOf("client-core", 0),
				artifactOf("server-pore", 1),
			},
			expMc: MavenCoords{Artifact: "ktor-client-core", Group: MavenGroup}, expResult: ModuleFound,
		},
		{
			artifacts: []network.Artifact{},
			expResult: ModuleNotFound,
		},
		{
			artifacts: []network.Artifact{artifactOf("server-freemarker", 1)},
			expResult: SimilarModulesFound, expCandidates: []MavenCoords{
				{Artifact: "ktor-server-freemarker", Group: MavenGroup},
			},
		},
	}

	for _, c := range cases {
		mc, result, candidates := FindModule(c.artifacts)

		if result != c.expResult {
			t.Errorf("expected result to be %v, got %v", c.expResult, result)
		}

		if mc.Artifact != c.expMc.Artifact {
			t.Errorf("expected Maven artifact name to be %s, got %s", c.expMc.Artifact, mc.Artifact)
		}

		if mc.Group != c.expMc.Group {
			t.Errorf("expected Maven group to be %s, got %s", c.expMc.Group, mc.Group)
		}

		if mc.IsTest != c.expMc.IsTest {
			t.Errorf("expected artifact's test=%v, got %v", c.expMc.IsTest, mc.IsTest)
		}

		if !reflect.DeepEqual(candidates, c.expCandidates) {
			t.Errorf("expected candidates to be %v, got %v", c.expCandidates, candidates)
		}
	}
}

func artifactOf(name string, distance int) network.Artifact {
	return network.Artifact{
		Name:     "ktor-" + name,
		Group:    MavenGroup,
		IsTest:   false,
		Distance: distance,
	}
}

func testArtifactOf(name string, distance int) network.Artifact {
	return network.Artifact{
		Name:     "ktor-" + name,
		Group:    MavenGroup,
		IsTest:   true,
		Distance: distance,
	}
}
