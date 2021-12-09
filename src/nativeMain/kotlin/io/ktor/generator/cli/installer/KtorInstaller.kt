package io.ktor.generator.cli.installer

import io.ktor.client.*
import io.ktor.client.request.*
import io.ktor.generator.bundle.*
import io.ktor.generator.cli.utils.*
import io.ktor.generator.configuration.json.*
import io.ktor.http.*
import kotlinx.cinterop.toKString
import kotlinx.coroutines.runBlocking
import platform.posix.*

class KtorInstaller(private val client: HttpClient, private val ktorBackendHost: String) {
    private val ktorRootDir: Directory by lazy { Directory.home().createDirIfNeeded(rootKtorDirName) }
    private val ktorRcFile: File by lazy { ktorRootDir.createFileIfNeeded(KTOR_RC_FILENAME) }

    private fun runGradle(gradleFile: File, task: String, javaHome: String, args: List<String> = emptyList()) {
        setenv(JAVA_HOME, javaHome, 1)
        addExecutablePermissions(gradleFile)
        runProcess("${gradleFile.path} $task ${args.joinToString(" ")}")
    }

    private fun getRcProperties(): Map<String, String> = ktorRcFile.readLines().associate {
        val (name, value) = it.split(" ")
        name to value
    }

    private fun addRcProperty(name: String, value: String) {
        ktorRcFile.writeText("$name $value\n")
    }

    private fun getRcProperty(name: String): String? = getRcProperties()[name]

    private fun createAdHockJdkDir(): Directory = ktorRootDir.createDirIfNeeded(JDK_INSTALLED_DIR_PATH)

    private fun initKtorRootIfAbsent() {
        assert(ktorRootDir.exists())
        assert(ktorRcFile.exists())
    }

    private fun findCustomJdk(): Directory? =
        ktorRootDir
            .content()
            .filterIsInstance<Directory>()
            .find { it.name == JDK_INSTALLED_DIR_PATH }
            ?.subdir(JAVA_CONTENTS)
            ?.subdir(JAVA_CONTENTS_HOME)

    private fun customJdkIsInstalled(): Boolean = findCustomJdk() != null

    private fun getJavaHome(): String? = getEnv(JAVA_HOME)

    private fun hasJavaHome11(): Boolean {
        val javaHome = getJavaHome() ?: return false
        return javaHome.contains("11")
    }

    private fun jdkIsInstalled(): Boolean =
        getRcProperty(JAVA_HOME) != null || customJdkIsInstalled() || hasJavaHome11()

    private fun installJdkIfAbsent() {
        initKtorRootIfAbsent()
        if (jdkIsInstalled()) {
            if (getRcProperty(JAVA_HOME) == null) {
                addRcProperty(JAVA_HOME, findCustomJdk()?.path ?: getJavaHome()!!)
            }
            return
        }

        val jdkArchiveFile = Directory.home().createFileIfNeeded(jdkArchiveName)

        PropertiesBundle.writeMessage("jdk.11.not.found", jdkDownloadUrl)
        runBlocking {
            client.downloadZip(jdkDownloadUrl, jdkArchiveFile)
        }
        PropertiesBundle.writeMessage("jdk.installed.success")

        val jdkDir = createAdHockJdkDir()

        unpackJdk(archive = jdkArchiveFile, outputDir = jdkDir)
        jdkArchiveFile.delete()

        val newJdkPath = findCustomJdk()?.path ?: throw Exception("Failed to setup JDK")
        addRcProperty(JAVA_HOME, newJdkPath)
    }

    fun downloadKtorProject(projectName: String) {
        installJdkIfAbsent()

        val currentDir = Directory.current()
        if (currentDir.subdir(projectName).exists()) {
            PropertiesBundle.writeMessage("project.already.exists", projectName)
            return
        }

        val projectZip = currentDir.createFileIfNeeded("$projectName.zip")

        PropertiesBundle.writeMessage("generating.project")
        runBlocking {
            val defaultKtorVersion =
                client.get<ProjectSettingsTemplate>("$ktorBackendHost/project/settings").ktorVersion.default

            client.downloadZip(
                "$ktorBackendHost/project/generate",
                projectZip,
                httpMethod = HttpMethod.Post
            ) {
                contentType(ContentType.Application.Json)

                body = SelectedProjectConfiguration(
                    settings = ProjectSettings(
                        name = projectName,
                        companyWebsite = "example.com",
                        ktorEngine = "NETTY",
                        buildSystemType = "GRADLE_KTS",
                        ktorVersion = defaultKtorVersion,
                        kotlinVersion = "LAST_KOTLIN_VERSION",
                    ), features = emptyList(), addWrapper = true
                )
            }
        }

        val projectDir = currentDir.createDirIfNeeded(projectName)

        unzip(zipFile = projectZip, projectDir)
        projectZip.delete()

        PropertiesBundle.writeMessage("project.downloaded", projectName)

        val ktorJavaHome = getRcProperty(JAVA_HOME)!!
        val gradleFile = projectDir.gradleWrapper() ?: return

        chdir(projectDir.path)
        runGradle(gradleFile, GRADLE_BUILD, ktorJavaHome)

        PropertiesBundle.writeMessage("project.generated", projectName)
    }

    fun runKtorProject(path: String, args: List<String>) {
        installJdkIfAbsent()
        val ktorJavaHome = getRcProperty(JAVA_HOME)!!

        val projectDir = Directory.current().subdir(path)
        if (!projectDir.exists()) {
            PropertiesBundle.writeMessage("project.not.exists")
            return
        }

        val gradleFile = projectDir.gradleWrapper() ?: return
        chdir(projectDir.path)
        runGradle(gradleFile, GRADLE_RUN, ktorJavaHome, args)
    }

    companion object {
        const val KTOR_RC_FILENAME: String = "ktor.rc"

        const val JAVA_HOME: String = "JAVA_HOME"

        const val JAVA_CONTENTS: String = "Contents"

        const val JAVA_CONTENTS_HOME: String = "Home"

        const val JDK_INSTALLED_DIR_PATH = "jdk11"

        const val GRADLE_BUILD = "build"

        const val GRADLE_RUN = "run"
    }

    private fun Directory.gradleWrapper(): File? = content().filterIsInstance<File>().find(::isGradleWrapper)
}