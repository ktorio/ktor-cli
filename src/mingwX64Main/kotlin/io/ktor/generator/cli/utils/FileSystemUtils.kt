package io.ktor.generator.cli.utils

import io.ktor.generator.cli.installer.*
import io.ktor.utils.io.core.*
import kotlinx.cinterop.ByteVar
import kotlinx.cinterop.CPointer
import kotlinx.cinterop.allocArray
import kotlinx.cinterop.toKString
import platform.posix.*

actual val FS_DELIMETER: String = "\\"

actual fun unzip(zipFile: File, outputDir: Directory) {
    val zipjsBat = getResourcePath("zipjs.bat")
    runProcess("call $zipjsBat unzip -source ${zipFile.path} -destination $outputDir -keep yes")
}

actual fun homePath(): String = getEnv("USERPROFILE") ?: throw Exception("Couldn't locate user home path")

actual fun addExecutablePermissions(file: File): Boolean = true

actual fun realPath(path: String, buffer: CPointer<ByteVar>): String? {
    return _fullpath(buffer, path, PATH_MAX)?.toKString()
}

actual fun getCwd(buffer: CPointer<ByteVar>, size: Int) = _getcwd(buffer, size)
actual fun makeDir(path: String): Boolean {
    return mkdir(path) == 0
}

actual fun createFile(path: String): Boolean {
    return _creat(path, (S_IWOTH or S_IROTH or S_IRUSR or S_IWUSR or S_IRGRP or S_IWGRP or S_IEXEC)) == 0
}