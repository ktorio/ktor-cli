package io.ktor.generator.cli.utils

import io.ktor.generator.bundle.*
import platform.posix.pclose
import platform.posix.popen
import kotlinx.cinterop.CPointer
import platform.posix.FILE
import platform.posix.exit

actual fun openPipe(command: String, access: String): CPointer<FILE>? = popen(command, access)
actual fun closePipe(filePtr: CPointer<FILE>): Int = pclose(filePtr)