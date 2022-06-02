package io.ktor.generator.cli.installer

import io.ktor.generator.cli.installer.Architecture.*
import io.ktor.generator.cli.utils.*
import kotlinx.cinterop.alloc
import kotlinx.cinterop.memScoped
import kotlinx.cinterop.ptr
import kotlinx.cinterop.toKString
import okio.FileSystem
import okio.Path.Companion.toPath
import platform.posix.uname
import platform.posix.utsname

actual val rootKtorDirName: String = ".ktor"
actual val jdkDownloadUrl: String
    get() = when (getArchitecture()) {
        X86_64 -> "https://download.java.net/java/ga/jdk11/openjdk-11_osx-x64_bin.tar.gz"
        ARM_64 -> "https://cdn.azul.com/zulu/bin/zulu11.54.25-ca-jdk11.0.14.1-macosx_aarch64.tar.gz"
    }
actual val jdkArchiveName: String = "openjdk-11.tar.gz"

actual fun unpackJdk(archive: File, outputDir: Directory) {
    val tempPath = Directory.current().path
    runProcess("tar -xvf ${archive.path} -C $tempPath")
    FileSystem.SYSTEM.atomicMove("$tempPath/jdk-11.jdk".toPath(), outputDir.path.toPath())
}

actual fun isGradleWrapper(file: File): Boolean = file.name.contains("gradlew")

actual fun setEnv(varName: String, value: String) {
    platform.posix.setenv(varName, value, 1)
}

actual fun getJdkContentsHome(directory: Directory?): Directory? =
    directory
        ?.subdir(KtorInstaller.JAVA_CONTENTS)
        ?.subdir(KtorInstaller.JAVA_CONTENTS_HOME)

private enum class Architecture {
    X86_64, ARM_64
}

private fun getArchitecture(): Architecture {
    val architectureName = memScoped {
        val systemInfo = alloc<utsname>()
        uname(systemInfo.ptr)
        systemInfo.machine.toKString()
    }
    return when (architectureName) {
        "x86_64" -> X86_64
        else -> ARM_64
    }
}