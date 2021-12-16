package io.ktor.generator

import DEFAULT_KTOR_URL
import createHttpClient
import io.ktor.generator.api.*
import io.ktor.generator.cli.installer.*
import io.ktor.generator.cli.utils.*
import io.ktor.generator.configuration.json.*
import kotlinx.coroutines.TimeoutCancellationException
import kotlinx.coroutines.runBlocking
import kotlinx.coroutines.withTimeout
import kotlin.test.Test
import kotlin.test.assertEquals
import kotlin.test.assertFailsWith
import kotlin.test.assertTrue

private fun nativeResource(name: String): File = File("src/nativeTest/resources/$name")

private fun createService(): KtorGeneratorWeb {
    val client = createHttpClient()
    return KtorGeneratorWebImpl(client, ktorBackendHost = DEFAULT_KTOR_URL)
}

class GeneratorIntegrationTests {

    @Test
    fun testUnzipWorks() {
        unzip(nativeResource("file.zip"), Directory.current())
        val unzippedFile = Directory.current().file("file.txt")
        assertEquals(unzippedFile.readText(), nativeResource("file.txt").readText())
        unzippedFile.delete()
    }

    @Test
    fun testDownloadJdkWorks() {
        val service = createService()
        val jdkContent = runBlocking {
            service.downloadJdkArchive()
        }
        assertTrue(jdkContent.isNotEmpty(), "Jdk archive should not be empty")

        val jdkFile = Directory.home().createFileIfNeeded(jdkArchiveName)
        jdkFile.writeContent(jdkContent)
        val outputDir = Directory.current().createDirIfNeeded("jdkUnpacked")
        assertTrue(outputDir.exists(), "JDK dir not found")

        unpackJdk(jdkFile, outputDir)
        jdkFile.delete()
        assertTrue(outputDir.content().isNotEmpty(), "Directory with JDK should be not empty after unpack")

        val contentsHome = outputDir.subdir(KtorInstaller.JAVA_CONTENTS).subdir(KtorInstaller.JAVA_CONTENTS_HOME)
        assertTrue(contentsHome.exists(), "JDK contains /Contents/Home dir")
        assertTrue(contentsHome.content().isNotEmpty(), "JDK Contents/Home must be not empty")
        outputDir.delete()
    }

    @Test
    fun testGeneration() {
        val ktorInstaller = KtorInstaller(createService())
        ktorInstaller.downloadKtorProject("ktor-sample-gen")

        val projectFolder = Directory.current().subdir("ktor-sample-gen")
        assertTrue(projectFolder.exists(), "Project folder was not created: ${projectFolder.path}, exists: ${projectFolder.exists()}")
        assertTrue(projectFolder.content().isNotEmpty(), "Project folder should not ve empty")
        projectFolder.delete()
    }

    object MockService : KtorGeneratorWeb {
        private val realService = createService()

        override suspend fun downloadJdkArchive(): ByteArray = realService.downloadJdkArchive()

        override suspend fun genProjectSettings(): ProjectSettingsTemplate = realService.genProjectSettings()

        override suspend fun generateKtorProject(configuration: SelectedProjectConfiguration): ByteArray =
            nativeResource("test-project.zip").readContent()
    }

    @Test
    fun testProjectRun() {
        val ktorInstaller = KtorInstaller(MockService)
        ktorInstaller.downloadKtorProject("ktor-sample-run")
        val projectDir = Directory.current().subdir("ktor-sample-run")

        runBlocking {
            ktorInstaller.runKtorProject("ktor-sample-run", args = emptyList())
        }

        val expectedOutput = projectDir.file("TestOutput.txt")

        assertTrue(expectedOutput.exists(), "Test project must populate output")
        assertEquals(
            "Hello, world!",
            expectedOutput.readText(),
            "Test project must populate output file with \"Hello, world!\" content"
        )
        expectedOutput.delete()
        projectDir.delete()
    }
}