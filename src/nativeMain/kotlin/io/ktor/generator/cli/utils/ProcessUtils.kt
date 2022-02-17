package io.ktor.generator.cli.utils

import io.ktor.generator.bundle.*
import kotlinx.cinterop.CPointer
import kotlinx.cinterop.memScoped
import platform.posix.FILE

expect fun openPipe(command: String, access: String): CPointer<FILE>?
expect fun closePipe(filePtr: CPointer<FILE>): Int

internal fun runProcess(command: String) {
    val filePtr = openPipe(command, "r")
    if (filePtr == null) {
        PropertiesBundle.writeMessage("unable.to.run.command", command)
        return
    }

    memScoped {
        handleOutput(filePtr, ::print)
    }

    val status = closePipe(filePtr)
    if (status == -1) {
        PropertiesBundle.writeMessage("error.running.command", command)
    }
}