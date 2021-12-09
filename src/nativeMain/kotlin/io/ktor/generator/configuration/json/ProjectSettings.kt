package io.ktor.generator.configuration.json

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class ProjectSettings(
    @SerialName(SETTINGS_PROJECT_NAME)
    val name: String,
    @SerialName(SETTINGS_COMPANY_WEBSITE)
    val companyWebsite: String,
    @SerialName(SETTINGS_ENGINE)
    val ktorEngine: String,
    @SerialName(SETTINGS_BUILD_SYSTEM)
    val buildSystemType: String,
    @SerialName(SETTINGS_KTOR_VERSION)
    val ktorVersion: String,
    @SerialName(SETTINGS_KOTLIN_VERSION)
    val kotlinVersion: String
)

enum class Engine(val engineName: String, val engineClass: String) {
    NETTY("netty", "Netty"),
    JETTY("jetty", "Jetty"),
    CIO("cio", "CIO"),
    TOMCAT("tomcat", "Tomcat")
}

enum class BuildSystemType {
    GRADLE, GRADLE_KTS, MAVEN;
}

val SUPPORTED_ENGINES = Engine.values().map { it.name }

val SUPPORTED_BUILD_SYSTEMS = BuildSystemType.values().map { it.name }