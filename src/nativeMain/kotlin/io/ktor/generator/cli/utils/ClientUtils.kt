package io.ktor.generator.cli.utils

import io.ktor.client.*
import io.ktor.client.call.*
import io.ktor.client.features.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.*
import io.ktor.util.*
import io.ktor.utils.io.*

suspend fun HttpClient.fetchZipContent(
    urlString: String,
    httpMethod: HttpMethod = HttpMethod.Get,
    block: HttpRequestBuilder.() -> Unit = {}
): ByteArray =
    request<HttpStatement> {
        url.takeFrom(urlString)
        method = httpMethod
        onDownload { bytesSentTotal, contentLength ->
            val progress = 100 * bytesSentTotal / contentLength
            if (progress != (100 * (bytesSentTotal - 1024) / contentLength)) print("\rProgress: $progress %")
            if (progress == 100L) {
                println()
            }
        }
        apply(block)
    }
        .receive()