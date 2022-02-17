package io.ktor.generator.cli.utils

import platform.posix._pclose
import platform.posix._popen
import kotlinx.cinterop.CPointer
import platform.posix.FILE

actual fun openPipe(command: String, access: String): CPointer<FILE>? = _popen(command, access)
actual fun closePipe(filePtr: CPointer<FILE>): Int = _pclose(filePtr)