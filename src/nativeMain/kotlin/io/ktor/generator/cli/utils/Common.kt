package io.ktor.generator.cli.utils

import kotlinx.cinterop.addressOf
import kotlinx.cinterop.toKString
import kotlinx.cinterop.usePinned
import platform.posix.PATH_MAX

expect val RESOURCES_PATH: String

fun getResourcePath(path: String): String {
    val filePath = "$RESOURCES_PATH$FS_DELIMETER$path"
    // Remove all '..' and '.'
    val buffer = ByteArray(PATH_MAX)
    val standardized = buffer.usePinned {
        realPath(filePath, it.addressOf(0))
    }
    return standardized ?: filePath
}

fun getResource(path: String): File = File(getResourcePath(path))