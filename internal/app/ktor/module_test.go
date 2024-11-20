package ktor

import (
	"reflect"
	"testing"
)

type testCase struct {
	name          string
	expMc         MavenCoords
	expResult     ModuleResult
	expCandidates []MavenCoords
}

func TestFindModule(t *testing.T) {
	var cases = []testCase{
		{name: "sse", expMc: MavenCoords{Artifact: "ktor-server-sse", Group: ktorGroup}, expResult: ModuleFound},
		{name: "js", expMc: MavenCoords{Artifact: "ktor-client-js", Group: ktorGroup}, expResult: ModuleFound},
		{name: "test-host", expMc: MavenCoords{Artifact: "ktor-server-test-host", Group: ktorGroup, IsTest: true}, expResult: ModuleFound},
		{name: "mock", expMc: MavenCoords{Artifact: "ktor-client-mock", Group: ktorGroup, IsTest: true}, expResult: ModuleFound},
		{name: "client-mock", expMc: MavenCoords{Artifact: "ktor-client-mock", Group: ktorGroup, IsTest: true}, expResult: ModuleFound},
		{name: "ktor-client-mock", expMc: MavenCoords{Artifact: "ktor-client-mock", Group: ktorGroup, IsTest: true}, expResult: ModuleFound},
		{name: "ktor-websocket-serialization", expMc: MavenCoords{Artifact: "ktor-websocket-serialization", Group: ktorGroup}, expResult: ModuleFound},
		{name: "json", expMc: MavenCoords{Artifact: "ktor-serialization-kotlinx-json", Group: ktorGroup}, expResult: ModuleFound},
		{name: "websocket-serialization", expMc: MavenCoords{Artifact: "ktor-websocket-serialization", Group: ktorGroup}, expResult: ModuleFound},
		{name: "content-negotiation", expResult: ModuleAmbiguity, expCandidates: []MavenCoords{
			{Artifact: "ktor-server-content-negotiation", Group: ktorGroup},
			{Artifact: "ktor-client-content-negotiation", Group: ktorGroup},
		}},
		{name: "core", expResult: ModuleAmbiguity, expCandidates: []MavenCoords{
			{Artifact: "ktor-server-core", Group: ktorGroup},
			{Artifact: "ktor-client-core", Group: ktorGroup},
		}},
		{name: "nonexistent", expResult: ModuleNotFound},
		{name: "freemaker", expResult: AlikeModuleFound, expMc: MavenCoords{Artifact: "ktor-server-freemarker", Group: ktorGroup}},
		{name: "ktor-clienp-core", expMc: MavenCoords{Artifact: "ktor-client-core", Group: ktorGroup}, expResult: AlikeModuleFound},
		{name: "websockets-serialization", expMc: MavenCoords{Artifact: "ktor-websocket-serialization", Group: ktorGroup}, expResult: AlikeModuleFound},
		{name: "peble", expMc: MavenCoords{Artifact: "ktor-server-pebble", Group: ktorGroup}, expResult: AlikeModuleFound},
		{name: "server-peble", expMc: MavenCoords{Artifact: "ktor-server-pebble", Group: ktorGroup}, expResult: AlikeModuleFound},
		{name: "ktor-server-peble", expMc: MavenCoords{Artifact: "ktor-server-pebble", Group: ktorGroup}, expResult: AlikeModuleFound},
	}

	for _, c := range cases {
		mc, result, candidates := FindModule(c.name)

		if result != c.expResult {
			t.Errorf("Module '%s': expected result %v, got %v", c.name, c.expResult, result)
		}

		if mc.Artifact != c.expMc.Artifact || mc.Group != c.expMc.Group || mc.IsTest != c.expMc.IsTest {
			t.Errorf("Module '%s': expected Maven coordinates to be %v, got %v", c.name, c.expMc, mc)
		}

		if !reflect.DeepEqual(candidates, c.expCandidates) {
			t.Errorf("Module '%s': expected candidates to be %v, got %v", c.name, c.expCandidates, candidates)
		}
	}
}
