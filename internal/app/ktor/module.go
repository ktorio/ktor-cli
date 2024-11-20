package ktor

import (
	"fmt"
	"github.com/agnivade/levenshtein"
	"math"
	"strings"
)

var serverDefs = map[string]string{
	"cio":             "CIO",
	"config-yaml":     "YamlConfig",
	"core":            "Application",
	"jetty":           "Jetty",
	"jetty-jakarta":   "Jetty",
	"netty":           "Netty",
	"servlet":         "ServletApplicationEngine",
	"servlet-jakarta": "ServletApplicationEngineJakarta",
	"tomcat":          "TomcatApplicationEngine",
	"tomcat-jakarta":  "TomcatApplicationEngineJakarta",

	// plugins

	"auth":                "Authentication",
	"auth-jwt":            "awt",
	"auth-ldap":           "ldap",
	"auto-head-response":  "AutoHeadResponse",
	"body-limit":          "RequestBodyLimit",
	"cachingHeaders":      "CachingHeaders",
	"call-id":             "CallId",
	"call-logging":        "CallLogging",
	"compression":         "Compression",
	"conditional-headers": "ConditionalHeaders",
	"content-negotiation": "ContentNegotiation",
	"cors":                "CORS",
	"csrf":                "CSRF",
	"data-conversion":     "DataConversion",
	"double-receive":      "DoubleReceive",
	"forwarded-header":    "ForwardedHeaders",
	"freemarker":          "FreeMarker",
	"hsts":                "HSTS",
	"html-builder":        "respondHtml",
	"http-redirect":       "HttpsRedirect",
	"i18n":                "I18n",
	"jte":                 "Jte",
	"method-override":     "XHttpMethodOverride",
	"metrics":             "DropwizardMetrics",
	"metrics-micrometer":  "MicrometerMetrics",
	"mustache":            "Mustache",
	"openapi":             "openAPI",
	"partial-content":     "PartialContent",
	"pebble":              "Pebble",
	"rate-limit":          "RateLimit",
	"request-validation":  "RequestValidation",
	"resources":           "Resources",
	"sessions":            "Sessions",
	"status-pages":        "StatusPages",
	"sse":                 "SSE",
	"swagger":             "swaggerUI",
	"thymeleaf":           "Thymeleaf",
	"velocity":            "Velocity",
	"webjars":             "Webjars",
	"websockets":          "Websockets",
	"test-host":           "testApplication",
}

var clientDefs = map[string]string{
	"android":       "Android",
	"apache":        "Apache",
	"apache5":       "Apache5",
	"cio":           "CIO",
	"core":          "HttpClient",
	"curl":          "Curl",
	"darwin":        "Darwin",
	"darwin-legacy": "DarwinLegacy",
	"ios":           "Ios",
	"java":          "Java",
	"jetty":         "Jetty",
	"js":            "Js",
	"mock":          "MockEngine",
	"okhttp":        "OkHttp",
	"winhttp":       "WinHttp",

	// plugins
	"auth":                "Auth",
	"bom-remover":         "BOMRemover",
	"call-id":             "CallId",
	"content-negotiation": "ContentNegotiation",
	"encoding":            "ContentEncoding",
	"logging":             "Logging",
	"resources":           "Resources",
	"websockets":          "Websockets",
}

var sharedDefs = map[string]string{
	"serialization-kotlinx-cbor":     "Cbor",
	"serialization-kotlinx-json":     "Json",
	"serialization-kotlinx-protobuf": "ProtoBuf",
	"serialization-kotlinx-xml":      "Xml",
	"serialization-gson":             "Gson",
	"serialization-jackson":          "Jackson",
	"websocket-serialization":        "WebsocketsSerialization",
}

var testDeps = map[string]struct{}{
	"test-host": {},
	"mock":      {},
}

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

func AllModules() []string {
	var modules []string

	for d, _ := range serverDefs {
		modules = append(modules, "ktor-server-"+d)
	}

	for d, _ := range clientDefs {
		modules = append(modules, "ktor-client-"+d)
	}

	for d, _ := range sharedDefs {
		modules = append(modules, "ktor-"+d)
	}

	return modules
}

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

type ModuleResult int

const (
	ModuleNotFound ModuleResult = iota
	ModuleAmbiguity
	AlikeModuleFound
	ModuleFound
)

func FindModule(name string) (MavenCoords, ModuleResult, []MavenCoords) {
	serverCandidates, clientCandidates, sharedCandidates := findModuleCandidates(name, 2)
	mc := MavenCoords{}

	var exactServerCandidates []MavenCoords
	for _, c := range serverCandidates {
		if c.Dist == 0 {
			exactServerCandidates = append(exactServerCandidates, c.Mc)
		}
	}

	var exactClientCandidates []MavenCoords
	for _, c := range clientCandidates {
		if c.Dist == 0 {
			exactClientCandidates = append(exactClientCandidates, c.Mc)
		}
	}

	var exactSharedCandidates []MavenCoords
	for _, c := range sharedCandidates {
		if c.Dist == 0 {
			exactSharedCandidates = append(exactSharedCandidates, c.Mc)
		}
	}

	if len(exactServerCandidates) > 0 && len(exactClientCandidates) == 0 && len(exactSharedCandidates) == 0 {
		return exactServerCandidates[0], ModuleFound, nil
	}

	if len(exactServerCandidates) == 0 && len(exactClientCandidates) > 0 && len(exactSharedCandidates) == 0 {
		return exactClientCandidates[0], ModuleFound, nil
	}

	if len(exactServerCandidates) == 0 && len(exactClientCandidates) == 0 && len(exactSharedCandidates) > 0 {
		return exactSharedCandidates[0], ModuleFound, nil
	}

	if len(serverCandidates) == 0 && len(clientCandidates) == 0 && len(sharedCandidates) == 0 {
		return mc, ModuleNotFound, nil
	}

	if (len(exactServerCandidates) > 0 && len(exactClientCandidates) > 0) || (len(exactServerCandidates) > 0 && len(exactSharedCandidates) > 0) {
		var candidates []MavenCoords
		for _, m := range exactServerCandidates {
			candidates = append(candidates, m)
		}
		for _, m := range exactClientCandidates {
			candidates = append(candidates, m)
		}
		for _, m := range exactSharedCandidates {
			candidates = append(candidates, m)
		}

		return mc, ModuleAmbiguity, candidates
	}

	if len(serverCandidates) > 0 {
		minDistMcc := mcCandidate{Dist: math.MaxInt}
		for _, mcc := range serverCandidates {
			if mcc.Dist == 0 {
				return mc, ModuleFound, nil
			}
			if mcc.Dist < minDistMcc.Dist {
				minDistMcc = mcc
			}
		}

		return minDistMcc.Mc, AlikeModuleFound, nil
	}

	if len(clientCandidates) > 0 {
		minDistMcc := mcCandidate{Dist: math.MaxInt}
		for _, mcc := range clientCandidates {
			if mcc.Dist == 0 {
				return mc, ModuleFound, nil
			}
			if mcc.Dist < minDistMcc.Dist {
				minDistMcc = mcc
			}
		}

		return minDistMcc.Mc, AlikeModuleFound, nil
	}

	if len(sharedCandidates) > 0 {
		minDistMcc := mcCandidate{Dist: math.MaxInt}
		for _, mcc := range sharedCandidates {
			if mcc.Dist == 0 {
				return mc, ModuleFound, nil
			}
			if mcc.Dist < minDistMcc.Dist {
				minDistMcc = mcc
			}
		}

		return minDistMcc.Mc, AlikeModuleFound, nil
	}

	return mc, ModuleNotFound, nil
}

func findModuleCandidates(name string, maxDistance int) (serverCandidates []mcCandidate, clientCandidates []mcCandidate, sharedCandidates []mcCandidate) {
	name = strings.ToLower(name)

	for k, alias := range serverDefs {
		k = strings.ToLower(k)

		if k == name || "server-"+k == name || "ktor-server-"+k == name || strings.ToLower(alias) == strings.ToLower(name) {
			serverCandidates = append(serverCandidates, mcCandidate{Mc: MavenCoords{Artifact: "ktor-server-" + k, Group: ktorGroup, IsTest: isTest(k)}})
		}
	}

	for k, alias := range serverDefs {
		k = strings.ToLower(k)

		for _, fullKey := range []string{k, "server-" + k, "ktor-server-" + k, strings.ToLower(alias)} {
			if dist := levenshtein.ComputeDistance(name, fullKey); dist <= maxDistance && !hasArtifact(serverCandidates, "ktor-server-"+k) {
				serverCandidates = append(serverCandidates, mcCandidate{Mc: MavenCoords{Artifact: "ktor-server-" + k, Group: ktorGroup, IsTest: isTest(k)}, Dist: dist})
			}
		}
	}

	for k, alias := range clientDefs {
		k = strings.ToLower(k)

		if k == name || "client-"+k == name || "ktor-client-"+k == name || strings.ToLower(alias) == strings.ToLower(name) {
			clientCandidates = append(clientCandidates, mcCandidate{Mc: MavenCoords{Artifact: "ktor-client-" + k, Group: ktorGroup, IsTest: isTest(k)}})
		}
	}

	for k, alias := range clientDefs {
		k = strings.ToLower(k)

		for _, fullKey := range []string{k, "client-" + k, "ktor-client-" + k, strings.ToLower(alias)} {
			if dist := levenshtein.ComputeDistance(name, fullKey); dist <= maxDistance && !hasArtifact(clientCandidates, "ktor-client-"+k) {
				clientCandidates = append(clientCandidates, mcCandidate{Mc: MavenCoords{Artifact: "ktor-client-" + k, Group: ktorGroup, IsTest: isTest(k)}, Dist: dist})
			}
		}
	}

	for k, alias := range sharedDefs {
		k = strings.ToLower(k)

		if k == name || "ktor-"+k == name || strings.ToLower(alias) == strings.ToLower(name) {
			sharedCandidates = append(sharedCandidates, mcCandidate{Mc: MavenCoords{Artifact: "ktor-" + k, Group: ktorGroup, IsTest: isTest(k)}})
		}
	}

	for k, alias := range sharedDefs {
		k = strings.ToLower(k)

		for _, fullKey := range []string{k, "ktor-" + k, strings.ToLower(alias)} {
			if dist := levenshtein.ComputeDistance(name, fullKey); dist <= maxDistance && !hasArtifact(clientCandidates, "ktor-"+k) {
				sharedCandidates = append(sharedCandidates, mcCandidate{Mc: MavenCoords{Artifact: "ktor-" + k, Group: ktorGroup, IsTest: isTest(k)}, Dist: dist})
			}
		}
	}

	return serverCandidates, clientCandidates, sharedCandidates
}

func isTest(artifact string) bool {
	_, ok := testDeps[artifact]
	return ok
}

func hasArtifact(mcs []mcCandidate, artifact string) bool {
	for _, m := range mcs {
		if m.Mc.Artifact == artifact {
			return true
		}
	}

	return false
}

type mcCandidate struct {
	Mc   MavenCoords
	Dist int
}
