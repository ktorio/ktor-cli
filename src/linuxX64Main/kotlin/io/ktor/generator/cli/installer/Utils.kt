package io.ktor.generator.cli.installer

import io.ktor.generator.cli.utils.*
import okio.FileSystem
import okio.Path.Companion.toPath

actual val rootKtorDirName: String = ".ktor"
actual val jdkDownloadUrl: String = "https://download.java.net/java/ga/jdk11/openjdk-11_linux-x64_bin.tar.gz"
actual val jdkArchiveName: String = "openjdk-11.tar.gz"

actual fun unpackJdk(archive: File, outputDir: Directory) {
    val tempPath = Directory.current().path
    runProcess("tar -xvf ${archive.path} -C $tempPath")
    FileSystem.SYSTEM.atomicMove("$tempPath/jdk-11".toPath(), outputDir.path.toPath())
}

actual fun isGradleWrapper(file: File): Boolean = file.name.contains("gradlew")

actual fun setEnv(varName: String, value: String) {
    platform.posix.setenv(varName, value, 1)
}

actual fun getJdkContentsHome(directory: Directory?): Directory? = directory