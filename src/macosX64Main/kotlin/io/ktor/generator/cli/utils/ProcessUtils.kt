package io.ktor.generator.cli.utils

import platform.posix.pclose
import platform.posix.popen
import kotlinx.cinterop.CPointer
import platform.posix.FILE

actual fun openPipe(command: String, access: String): CPointer<FILE>? = popen(command, access)
actual fun closePipe(filePtr: CPointer<FILE>): Int = pclose(filePtr)