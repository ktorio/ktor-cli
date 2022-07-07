package io.ktor.generator.cli.utils

import io.ktor.generator.cli.installer.*
import io.ktor.generator.cli.installer.getEnv
import kotlinx.cinterop.ByteVar
import kotlinx.cinterop.CPointer
import kotlinx.cinterop.pointed
import kotlinx.cinterop.toKString
import kotlinx.cinterop.allocArray
import platform.posix.*

private const val HOME_VAR: String = "HOME"

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
    mkdir(path, (S_IWOTH or S_IROTH or S_IROTH or S_IRUSR or S_IWUSR or S_IRGRP or S_IWGRP or S_IEXEC).toUShort())
}

actual fun createFile(path: String) {
    creat(path, (S_IWOTH or S_IROTH or S_IRUSR or S_IWUSR or S_IRGRP or S_IWGRP or S_IEXEC).toUShort())
}