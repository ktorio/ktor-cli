package io.ktor.generator.cli.utils

import io.ktor.client.*
import io.ktor.client.features.json.*
import io.ktor.client.features.json.serializer.*
import io.ktor.generator.bundle.*
import io.ktor.generator.cli.installer.*
import kotlinx.cinterop.*
import kotlin.text.*
import kotlinx.cinterop.allocArray
import platform.posix.*

actual val FS_DELIMETER: String = "/"

actual fun unzip(zipFile: File, outputDir: Directory) {
    assertCommandExists("unzip")
    runProcess("unzip ${zipFile.path} -d ${outputDir.path}")
}

actual fun homePath(): String =
    getEnv(HOME_VAR) ?: getpwuid(getuid())?.pointed?.pw_dir?.toKString() ?: throw Exception("Failed to locate home dir")

private val RW_PERMISSIONS = (S_IWOTH or S_IROTH or S_IRUSR or S_IWUSR or S_IRGRP or S_IWGRP).toUInt()
private val RWE_PERMISSIONS = RW_PERMISSIONS or S_IEXEC.toUInt()

actual fun addExecutablePermissions(file: File): Boolean {
    return chmod(file.path, RWE_PERMISSIONS) == 0
}

actual fun realPath(path: String, buffer: CPointer<ByteVar>): String? {
    return realpath(path, buffer)?.toKString()
}

actual fun getCwd(buffer: CPointer<ByteVar>, size: Int) = getcwd(buffer, size.toULong())

actual fun makeDir(path: String): Boolean {
    return mkdir(path, RWE_PERMISSIONS) == 0
}

actual fun createFile(path: String): Boolean {
    return creat(path, RWE_PERMISSIONS) == 0
}