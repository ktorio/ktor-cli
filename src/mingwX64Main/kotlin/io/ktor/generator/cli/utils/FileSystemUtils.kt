package io.ktor.generator.cli.utils

import io.ktor.generator.cli.installer.*

actual val FS_DELIMETER: String = "\\"

actual fun unzip(zipFile: File, outputDir: Directory) {
    val zipjsBat = getResourcePath("zipjs.bat")
    runProcess("call $zipjsBat unzip -source ${zipFile.path} -destination $outputDir -keep yes")
}

actual fun homePath(): String = getEnv("USERPROFILE") ?: throw Exception("Couldn't locate user home path")

actual fun addExecutablePermissions(file: File) {}