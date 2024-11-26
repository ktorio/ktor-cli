package ktor

import (
	"fmt"
	"github.com/agnivade/levenshtein"
	"slices"
	"strings"
)

var defs = map[string]string{
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

	"auth":                           "Auth",
	"auth-jwt":                       "awt",
	"auth-ldap":                      "ldap",
	"auto-head-response":             "AutoHeadResponse",
	"body-limit":                     "RequestBodyLimit",
	"cachingHeaders":                 "CachingHeaders",
	"call-id":                        "CallId",
	"call-logging":                   "CallLogging",
	"compression":                    "Compression",
	"conditional-headers":            "ConditionalHeaders",
	"content-negotiation":            "ContentNegotiation",
	"cors":                           "CORS",
	"csrf":                           "CSRF",
	"data-conversion":                "DataConversion",
	"double-receive":                 "DoubleReceive",
	"forwarded-header":               "ForwardedHeaders",
	"freemarker":                     "FreeMarker",
	"hsts":                           "HSTS",
	"html-builder":                   "respondHtml",
	"http-redirect":                  "HttpsRedirect",
	"i18n":                           "I18n",
	"jte":                            "Jte",
	"method-override":                "XHttpMethodOverride",
	"metrics":                        "DropwizardMetrics",
	"metrics-micrometer":             "MicrometerMetrics",
	"mustache":                       "Mustache",
	"openapi":                        "openAPI",
	"partial-content":                "PartialContent",
	"pebble":                         "Pebble",
	"rate-limit":                     "RateLimit",
	"request-validation":             "RequestValidation",
	"resources":                      "Resources",
	"sessions":                       "Sessions",
	"status-pages":                   "StatusPages",
	"sse":                            "SSE",
	"swagger":                        "swaggerUI",
	"thymeleaf":                      "Thymeleaf",
	"velocity":                       "Velocity",
	"webjars":                        "Webjars",
	"websockets":                     "Websockets",
	"test-host":                      "testApplication",
	"serialization-kotlinx-cbor":     "Cbor",
	"serialization-kotlinx-json":     "Json",
	"serialization-kotlinx-protobuf": "ProtoBuf",
	"serialization-kotlinx-xml":      "Xml",
	"serialization-gson":             "Gson",
	"serialization-jackson":          "Jackson",
	"websocket-serialization":        "sendSerialized",
}

var shared = map[string]struct{}{
	"serialization-kotlinx-cbor":     {},
	"serialization-kotlinx-json":     {},
	"serialization-kotlinx-protobuf": {},
	"serialization-kotlinx-xml":      {},
	"serialization-gson":             {},
	"serialization-jackson":          {},
	"websocket-serialization":        {},
}

var testing = map[string]struct{}{
	"test-host": {},
}

const ktorGroup = "io.ktor"

var modules map[string]MavenCoords
var ktorServerModules map[string]MavenCoords
var serverModules map[string]MavenCoords
var modulesBySymbol map[string]MavenCoords

func init() {
	modules = make(map[string]MavenCoords)
	for name := range defs {
		_, isTest := testing[name]
		modules[name] = MavenCoords{Artifact: getPrefix(name) + name, Group: ktorGroup, IsTest: isTest}
	}

	serverModules = make(map[string]MavenCoords)
	for name := range defs {
		_, isTest := testing[name]
		serverModules["server-"+name] = MavenCoords{Artifact: "ktor-server-" + name, Group: ktorGroup, IsTest: isTest}
	}

	ktorServerModules = make(map[string]MavenCoords)
	for name := range defs {
		_, isTest := testing[name]
		ktorServerModules["ktor-server-"+name] = MavenCoords{Artifact: "ktor-server-" + name, Group: ktorGroup, IsTest: isTest}
	}

	modulesBySymbol = make(map[string]MavenCoords)
	for name, symbol := range defs {
		_, isTest := testing[name]
		modulesBySymbol[symbol] = MavenCoords{Artifact: getPrefix(name) + name, Group: ktorGroup, IsTest: isTest}
	}
}

func getPrefix(name string) string {
	if _, ok := shared[name]; ok {
		return "ktor-"
	}

	return "ktor-server-"
}

type MavenCoords struct {
	Artifact, Group, Version string
	IsTest                   bool
}

type GradlePlugin struct {
	Id              string
	IsSerialization bool
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

func FindModule(name string) (MavenCoords, int, bool) {
	mc, dist, ok := findIn(modules, strings.ToLower(name), 2)

	if ok {
		return mc, dist, true
	}

	mc, dist, ok = findIn(serverModules, strings.ToLower(name), 2)

	if ok {
		return mc, dist, true
	}

	mc, dist, ok = findIn(ktorServerModules, strings.ToLower(name), 2)

	if ok {
		return mc, dist, true
	}

	mc, dist, ok = findIn(modulesBySymbol, name, 2)

	if ok {
		return mc, dist, true
	}

	return MavenCoords{}, 0, false
}

type mcCandidate struct {
	mc   MavenCoords
	dist int
}

func findIn(mp map[string]MavenCoords, name string, maxDistance int) (MavenCoords, int, bool) {
	if mc, ok := mp[name]; ok {
		return mc, 0, true
	}

	var candidates []mcCandidate

	for m, mc := range mp {
		distance := levenshtein.ComputeDistance(strings.ToLower(name), strings.ToLower(m))

		if distance <= maxDistance {
			candidates = append(candidates, mcCandidate{mc: mc, dist: distance})
		}
	}

	if len(candidates) == 0 {
		return MavenCoords{}, 0, false
	}

	cd := slices.MinFunc(candidates, func(a, b mcCandidate) int {
		return a.dist - b.dist
	})

	return cd.mc, cd.dist, true
}
