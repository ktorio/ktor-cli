package io.ktor.generator.cli.utils

import io.ktor.generator.bundle.*
import kotlinx.cinterop.CPointer
import kotlinx.cinterop.memScoped
import platform.posix.FILE
import platform.posix.exit

expect fun openPipe(command: String, access: String): CPointer<FILE>?
expect fun closePipe(filePtr: CPointer<FILE>): Int
fun libcurlExists(): Boolean = commandExists("curl")

internal fun runProcess(command: String): Int {
    val filePtr = openPipe(command, "r")
    if (filePtr == null) {
        PropertiesBundle.writeErrorMessage("unable.to.run.command", command)
        return 1
    }

    memScoped {
        handleOutput(filePtr, ::print)
    }

    val status = closePipe(filePtr)
    println(" ___ DBG __ STATUS: $status")
    if (status == -1) {
        PropertiesBundle.writeErrorMessage("error.running.command", command)
    }

    return status
}

internal fun commandExists(commandName: String): Boolean =
    runProcess("command -v $commandName") == 0

fun assertCommandExists(commandName: String) {
    if (!commandExists(commandName)) {
        PropertiesBundle.writeErrorMessage("command.does.not.exist", commandName)
        exit(1)
    }
}

fun assertLibcurlExists() {
    if (!libcurlExists()) {
        PropertiesBundle.writeErrorMessage("command.does.not.exist", "libcurl")
        exit(1)
    }
}