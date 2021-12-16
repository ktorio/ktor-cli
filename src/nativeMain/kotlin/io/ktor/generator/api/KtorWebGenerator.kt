package io.ktor.generator.api

import io.ktor.client.*
import io.ktor.client.request.*
import io.ktor.generator.cli.installer.*
import io.ktor.generator.cli.utils.*
import io.ktor.generator.configuration.json.*
import io.ktor.http.*

interface KtorGeneratorWeb {
    /**
     * Downloads jdk archive and returns its byte content
     * */
    suspend fun downloadJdkArchive(): ByteArray

    /**
     * Generates Ktor project and returns its byte content
     * */
    suspend fun generateKtorProject(configuration: SelectedProjectConfiguration): ByteArray

    /**
     * Generates Ktor project and returns its byte content
     * */
    suspend fun genProjectSettings(): ProjectSettingsTemplate
}

class KtorGeneratorWebImpl(val client: HttpClient, private val ktorBackendHost: String) : KtorGeneratorWeb {
    override suspend fun downloadJdkArchive(): ByteArray = client.fetchZipContent(jdkDownloadUrl)

    override suspend fun generateKtorProject(configuration: SelectedProjectConfiguration): ByteArray =
        client.fetchZipContent(
            "$ktorBackendHost/project/generate",
            httpMethod = HttpMethod.Post
        ) {
            contentType(ContentType.Application.Json)

            body = configuration
        }

    override suspend fun genProjectSettings(): ProjectSettingsTemplate =
        client.get("$ktorBackendHost/project/settings")
}