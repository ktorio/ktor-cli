package io.ktor.generator.cli.installer

import io.ktor.generator.cli.utils.*
import okio.FileSystem
import okio.Path.Companion.toPath

actual val rootKtorDirName: String = ".ktor."
actual val jdkDownloadUrl: String = "https://download.java.net/java/ga/jdk11/openjdk-11_windows-x64_bin.zip"
actual val jdkArchiveName: String = "jdk-11"

actual fun unpackJdk(archive: File, outputDir: Directory) {
    unzip(zipFile = archive, outputDir = Directory.current())

    FileSystem.SYSTEM.atomicMove("${Directory.current().path}\\jdk-11".toPath(), outputDir.path.toPath())
}

actual fun isGradleWrapper(file: File): Boolean = file.name.contains("gradlew")