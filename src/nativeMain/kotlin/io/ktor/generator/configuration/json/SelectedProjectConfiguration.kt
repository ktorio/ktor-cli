package io.ktor.generator.configuration.json

import kotlinx.serialization.Serializable

@Serializable
data class SelectedProjectConfiguration(
    val settings: ProjectSettings,
    val features: List<String>,
    val addDefaultRoutes: Boolean = true,
    val configurationOption: ConfigurationOptions = ConfigurationOptions.CODE,
    val addWrapper: Boolean = false
)

@Serializable
enum class ConfigurationOptions {
    CODE, HOCON
}