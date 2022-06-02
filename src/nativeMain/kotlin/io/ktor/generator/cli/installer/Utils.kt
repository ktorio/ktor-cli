package io.ktor.generator.cli.installer

import io.ktor.generator.cli.utils.*
import kotlinx.cinterop.toKString
import platform.posix.getenv

internal fun getEnv(name: String): String? = getenv(name)?.toKString()

expect val rootKtorDirName: String
expect val jdkDownloadUrl: String

expect val jdkArchiveName: String
expect fun unpackJdk(archive: File, outputDir: Directory)
expect fun isGradleWrapper(file: File): Boolean
expect fun setEnv(varName: String, value: String)
expect fun getJdkContentsHome(directory: Directory?): Directory?