package io.ktor.generator.cli.utils

import io.ktor.generator.bundle.*
import kotlinx.cinterop.memScoped
import platform.posix.pclose
import platform.posix.popen

internal fun runProcess(command: String) {
    val filePtr = popen(command, "r")
    if (filePtr == null) {
        PropertiesBundle.writeMessage("unable.to.run.command", command)
        return
    }

    memScoped {
        handleOutput(filePtr, ::print)
    }

    val status = pclose(filePtr)
    if (status == -1) {
        PropertiesBundle.writeMessage("error.running.command", command)
    }
}