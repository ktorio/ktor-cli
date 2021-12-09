package io.ktor.generator.configuration.json

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class ProjectSettingsTemplate(
    @SerialName(PROJECT_NAME)
    val name: TextSetting,
    @SerialName(COMPANY_WEBSITE)
    val companyWebsite: TextSetting,
    @SerialName(ENGINE)
    val engine: OptionalSetting,
    @SerialName(KTOR_VERSION)
    var ktorVersion: OptionalSetting,
    @SerialName(KOTLIN_VERSION)
    var kotlinVersion: OptionalSetting,
    @SerialName(BUILD_SYSTEM)
    val buildSystem: OptionalSetting,
    @SerialName(CONFIGURATION)
    val configuration: OptionalSetting
)

@Serializable
data class TextSetting(
    val default: String
)

@Serializable
data class OptionalSetting(
    val options: List<Option>,
    @SerialName(OPTION_DEFAULT)
    val default: String
) {
    @Serializable
    data class Option(
        val id: String,
        val name: String
    )
}

