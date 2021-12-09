package io.ktor.generator.bundle

import kotlin.test.Test
import kotlin.test.assertEquals

class BundleParsingTest {
    @Test
    fun testSubstitutions() {
        assertEquals(PropertiesBundle.message("unable.to.run.command", "cmd"), "Unable to run: cmd")
        assertEquals(PropertiesBundle.message("error.running.command", "cmd -rf"), "Error running process: cmd -rf")
        assertEquals(
            PropertiesBundle.message("jdk.11.not.found", "http://jdk.down/load"),
            "JDK 11 not found.\nJdk download path: http://jdk.down/load\nDownloading JDK 11 from server, please wait..."
        )
        assertEquals(
            PropertiesBundle.message("project.already.exists", "ExistingName"),
            "Project with name ExistingName already exists"
        )
        assertEquals(
            PropertiesBundle.message("project.downloaded", "name"),
            "Project \"name\" was downloaded. Running gradle setup..."
        )
        assertEquals(
            PropertiesBundle.message("project.generated", "project-name"),
            "Project \"project-name\" was successfully generated.\nYou can execute `ktor run project-name` to run it"
        )
        assertEquals(PropertiesBundle.message("project.not.exists", "p0"), "Project p0 does not exist")
    }

    private fun assertError(expectedError: Boolean, message: String, block: () -> Any) {
        var errorHappened = false
        try {
            block()
        } catch (_: Error) {
            errorHappened = true
        }

        assert(errorHappened == expectedError) { message }
    }

    private fun assertErrorInBundle(property: String) {
        assertError(true, "Expected error for \"$property\" with 0 args") { PropertiesBundle.message(property) }
    }

    private fun assertNoErrorInBundle(property: String) {
        assertError(false, "Error was not expected for \"$property\"") { PropertiesBundle.message(property) }
    }

    @Test
    fun testWrongNumberOfArguments() {
        assertErrorInBundle("unable.to.run.command")
        assertErrorInBundle("error.running.command")
        assertErrorInBundle("jdk.11.not.found")
        assertErrorInBundle("project.already.exists")
        assertErrorInBundle("project.downloaded")
        assertErrorInBundle("project.generated")
        assertErrorInBundle("project.not.exists")
    }

    @Test
    fun testSimpleMessagesNoError() {
        assertNoErrorInBundle("program.name")
        assertNoErrorInBundle("ktor.backend.url.description")
        assertNoErrorInBundle("generate.command.description")
        assertNoErrorInBundle("run.command.description")
        assertNoErrorInBundle("project.name.description")
        assertNoErrorInBundle("run.arguments.description")
        assertNoErrorInBundle("jdk.installed.success")
        assertNoErrorInBundle("generating.project")
    }
}