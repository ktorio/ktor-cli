package io.ktor.generator.bundle

import kotlin.test.assertNotNull

// TODO: figure out how to read properties from resources and deploy with project in Kotlin/Native
object PropertiesBundle {
    private val propToValue: Map<String, String> = mapOf(
        "program.name" to "Ktor CLI generator",
        "ktor.backend.url.description" to "Ktor generator backend url. It is recommended to leave this property as default",
        "generate.command.description" to "Generate new ktor project",
        "run.command.description" to "Run existing ktor project",
        "project.name.description" to "Name of the ktor project",
        "run.arguments.description" to "Arguments that will be passed to your project",
        "unable.to.run.command" to "Unable to run: {0}",
        "error.running.command" to "Error running process: {0}",
        "jdk.11.not.found" to "JDK 11 not found.\nJdk download path: {0}\nDownloading JDK 11 from server, please wait...",
        "jdk.installed.success" to "Successfully installed JDK (Java Development Kit) on your computer.",
        "project.already.exists" to "Project with name {0} already exists",
        "generating.project" to "Generating your ktor project",
        "project.downloaded" to "Project \"{0}\" was downloaded. Running gradle setup...",
        "project.generated" to "Project \"{0}\" was successfully generated.\nYou can execute `ktor run {0}` to run it",
        "project.not.exists" to "Project {0} does not exist"
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
                assert(argId < args.size) { "Insufficient number of arguments provided for $property. Required at least $argId but only ${args.size} found" }

                return@replace args[argId]
            }
        }

        return value
    }

    fun writeMessage(property: String, vararg args: String) = println(message(property, *args))
}