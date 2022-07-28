package io.ktor.generator.bundle

import com.github.ajalt.mordant.rendering.TextColors.green
import com.github.ajalt.mordant.rendering.TextColors.red
import com.github.ajalt.mordant.terminal.Terminal
import kotlin.test.assertNotNull

private const val JDK_LICENSE_LINK = "https://www.oracle.com/a/tech/docs/jdk11-lium.pdf"

const val ANSWER_YES = "Y"

// TODO: figure out how to read properties from resources and deploy with project in Kotlin/Native
object PropertiesBundle {
    private val propToValue: Map<String, String> = mapOf(
        "program.name" to "Ktor CLI generator",
        "generate.command.description" to "Generate new ktor project",
        "run.command.description" to "Run existing ktor project",
        "project.name.description" to "Name of the ktor project",
        "unable.to.run.command" to "Unable to run: {0}",
        "error.running.command" to "Error running process: {0}",
        "jdk.11.not.found" to "JDK 11 not found.\nJdk download path: {0}\nDownloading JDK 11 from server, please wait...",
        "jdk.installed.success" to "Successfully installed JDK (Java Development Kit) on your computer.",
        "project.already.exists" to "Project with name {0} already exists",
        "generating.project" to "Generating your ktor project",
        "project.downloaded" to "Project \"{0}\" was downloaded. Running gradle setup...",
        "project.generated" to "Project \"{0}\" was successfully generated.\nYou can execute `ktor start {0}` to start it",
        "project.not.exists" to "Project {0} does not exist",
        "project.not.have.gradlew" to "Invalid project. Project \"{0}\" does not have gradlew file",
        "download.jdk.legal.message" to "JDK not found\n" +
                "JDK is a software licensed by Oracle, Inc. under the terms available at $JDK_LICENSE_LINK.\n" +
                "JDK download path: {0}.\n" +
                "By typing \"$ANSWER_YES\" you agree with these terms and completion of installation [Y/n]: ",
        "jdk.legal.rejected" to "JDK is required to proceed. JDK not found. Quitting...",
        "jdk.setup.failed" to "Failed to setup JDK",
        "error.happened" to "Error happened: {0}"
    )

    private val argumentRegex = "\\{\\d+}".toRegex()

    fun message(property: String, vararg args: String): String {
        assert(property in propToValue) { "Unknown property $property in resource bundle" }
        var value = propToValue[property]!!

        while (value.contains(argumentRegex)) {
            value = value.replace(argumentRegex) {
                val argId = it.value.drop(1).dropLast(1).toIntOrNull()

                assertNotNull(argId, "Bad argument format in $property")
                assert(argId >= 0) { "Negative argument index in $property: $argId. Must be non-negative" }
                assert(argId < args.size) { "Insufficient number of arguments provided for $property. Required at least ${argId + 1} but only ${args.size} found" }

                return@replace args[argId]
            }
        }

        return value
    }

    fun writeMessage(property: String, vararg args: String) = println(message(property, *args))

    fun askQuestion(property: String, vararg args: String): Boolean {
        print(message(property, *args))
        return readLine() == ANSWER_YES
    }
    fun writeErrorMessage(property: String, vararg args: String) {
        println()
        Terminal().println(red(message(property, *args)))
    }
    fun writeSuccessMessage(property: String, vararg args: String) {
        println()
        Terminal().println(green(message(property, *args)))
    }
}