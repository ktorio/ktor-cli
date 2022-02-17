package io.ktor.generator.cli.utils

import kotlinx.cinterop.ByteVar
import kotlinx.cinterop.CPointer
import kotlinx.cinterop.pointed
import kotlinx.cinterop.toKString
import platform.posix.getpwuid
import platform.posix.getuid
import platform.posix.realpath
import kotlinx.cinterop.allocArray
import platform.posix.getcwd

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