package io.ktor.generator.cli.utils

import io.ktor.client.*
import io.ktor.client.features.json.*
import io.ktor.client.features.json.serializer.*
import io.ktor.generator.cli.installer.*
import kotlinx.cinterop.*
import kotlin.text.*
import kotlinx.cinterop.allocArray
import platform.posix.*

actual val FS_DELIMETER: String = "/"

actual fun unzip(zipFile: File, outputDir: Directory) {
    runProcess("unzip ${zipFile.path} -d ${outputDir.path}")
}

actual fun homePath(): String =
    getEnv(HOME_VAR) ?: getpwuid(getuid())?.pointed?.pw_dir?.toKString() ?: throw Exception("Failed to locate home dir")

actual fun addExecutablePermissions(file: File) {
    runProcess("chmod +x ${file.path}")
}

actual fun realPath(path: String, buffer: CPointer<ByteVar>): String? {
    return realpath(path, buffer)?.toKString()
}

actual fun getCwd(buffer: CPointer<ByteVar>, size: Int) = getcwd(buffer, size.toULong())

actual fun makeDir(path: String) {
    mkdir(path, (S_IWOTH or S_IROTH or S_IRUSR or S_IWUSR or S_IRGRP or S_IWGRP or S_IEXEC).toUInt())
}

actual fun createFile(path: String) {
    creat(path, (S_IWOTH or S_IROTH or S_IRUSR or S_IWUSR or S_IRGRP or S_IWGRP or S_IEXEC).toUInt())
}