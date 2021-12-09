package io.ktor.generator.cli.utils

import io.ktor.client.*
import io.ktor.client.features.json.*
import io.ktor.client.features.json.serializer.*
import io.ktor.generator.cli.installer.*
import kotlinx.cinterop.pointed
import kotlinx.cinterop.toKString
import platform.posix.getpwuid
import platform.posix.getuid
import kotlin.text.*

actual val FS_DELIMETER: String = "/"

actual fun unzip(zipFile: File, outputDir: Directory) {
    runProcess("unzip ${zipFile.path} -d ${outputDir.path}")
}

actual fun homePath(): String =
    getpwuid(getuid())?.pointed?.pw_dir?.toKString() ?: throw Exception("Failed to locate home dir")

actual fun addExecutablePermissions(file: File) {
    runProcess("chmod +x ${file.path}")
}