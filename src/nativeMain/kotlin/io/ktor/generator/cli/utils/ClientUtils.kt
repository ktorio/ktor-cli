package io.ktor.generator.cli.utils

import io.ktor.client.*
import io.ktor.client.call.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.*
import io.ktor.util.*
import io.ktor.utils.io.*

suspend fun HttpClient.downloadZip(
    urlString: String,
    zipFile: File,
    httpMethod: HttpMethod = HttpMethod.Get,
    block: HttpRequestBuilder.() -> Unit = {}
) {
    val zipBytes =
        request<HttpStatement> {
            url.takeFrom(urlString)
            method = httpMethod
            apply(block)
        }
            .receive<ByteArray>()
    zipFile.writeContent(zipBytes)
}