package io.ktor.generator.cli.utils

import io.ktor.client.*
import io.ktor.client.features.json.*
import io.ktor.client.features.json.serializer.*
import io.ktor.generator.cli.installer.*
import kotlinx.cinterop.*
import platform.posix.getcwd
import platform.posix.getpwuid
import platform.posix.getuid
import platform.posix.realpath
import kotlin.text.*
import kotlinx.cinterop.allocArray

actual val FS_DELIMETER: String = "/"

actual fun unzip(zipFile: File, outputDir: Directory) {
    runProcess("unzip ${zipFile.path} -d ${outputDir.path}")
}

actual fun homePath(): String =
    getpwuid(getuid())?.pointed?.pw_dir?.toKString() ?: throw Exception("Failed to locate home dir")

actual fun addExecutablePermissions(file: File) {
    runProcess("chmod +x ${file.path}")
}

actual fun realPath(path: String, buffer: CPointer<ByteVar>): String? {
    return realpath(path, buffer)?.toKString()
}

actual fun getCwd(buffer: CPointer<ByteVar>, size: Int) = getcwd(buffer, size.toULong())