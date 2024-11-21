package ktor

import (
	"fmt"
	"github.com/agnivade/levenshtein"
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

	"auth":                "Auth",
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

const ktorGroup = "io.ktor"

var modules map[string]MavenCoords
var ktorServerModules map[string]MavenCoords
var serverModules map[string]MavenCoords
var modulesBySymbol map[string]MavenCoords

func init() {
	modules = make(map[string]MavenCoords)
	for name := range defs {
		modules[name] = MavenCoords{Artifact: "ktor-server-" + name, Group: ktorGroup}
	}

	serverModules = make(map[string]MavenCoords)
	for name := range defs {
		serverModules["server-"+name] = MavenCoords{Artifact: "ktor-server-" + name, Group: ktorGroup}
	}

	ktorServerModules = make(map[string]MavenCoords)
	for name := range defs {
		ktorServerModules["ktor-server-"+name] = MavenCoords{Artifact: "ktor-server-" + name, Group: ktorGroup}
	}

	modulesBySymbol = make(map[string]MavenCoords)
	for name, symbol := range defs {
		modulesBySymbol[symbol] = MavenCoords{Artifact: "ktor-server-" + name, Group: ktorGroup}
	}
}

type MavenCoords struct {
	Artifact, Group string
}

func (mc *MavenCoords) String() string {
	return fmt.Sprintf("%s:%s", mc.Group, mc.Artifact)
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

func findIn(mp map[string]MavenCoords, name string, maxDistance int) (MavenCoords, int, bool) {
	if mc, ok := mp[name]; ok {
		return mc, 0, true
	}

	for m, mc := range mp {
		distance := levenshtein.ComputeDistance(strings.ToLower(name), m)

		if distance <= maxDistance {
			return mc, distance, true
		}
	}

	return MavenCoords{}, 0, false
}
