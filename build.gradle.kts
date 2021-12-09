import org.jetbrains.kotlin.gradle.plugin.KotlinSourceSet

val ktor_version: String by project
val kotlinx_cli_version: String by project

plugins {
    kotlin("multiplatform") version "1.5.31"
    kotlin("plugin.serialization") version "1.5.31"
}

group = "me.user"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
    maven("https://maven.pkg.jetbrains.space/public/p/ktor/eap")
}

kotlin {
    val hostOs = System.getProperty("os.name")
    val isMingwX64 = hostOs.startsWith("Windows")
    val nativeTarget = when {
        hostOs == "Mac OS X" -> macosX64("native")
        hostOs == "Linux" -> linuxX64("native")
        isMingwX64 -> mingwX64("native")
        else -> throw GradleException("Host OS is not supported in Kotlin/Native.")
    }

    nativeTarget.apply {
        binaries {
            executable {
                entryPoint = "main"
            }
        }

    }
    linuxX64 {
        binaries {
            executable {
                entryPoint = "main"
            }
        }
        compilations["main"].enableEndorsedLibs = true
    }
    mingwX64 {
        binaries {
            executable {
                entryPoint = "main"
            }
        }
        compilations["main"].enableEndorsedLibs = true
    }
    macosX64("macosX64") {
        binaries {
            executable {
                entryPoint = "main"
            }
        }
        compilations["main"].enableEndorsedLibs = true
    }
    sourceSets {
        val nativeMain by getting {
            dependencies {
                implementation("io.ktor:ktor-client-core:$ktor_version")
                implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:1.5.2-native-mt")
                implementation("io.ktor:ktor-client-serialization:$ktor_version")
                implementation("com.squareup.okio:okio:3.0.0")
                implementation("io.ktor:ktor-client-curl:$ktor_version")
                implementation("org.jetbrains.kotlinx:kotlinx-cli:$kotlinx_cli_version")
            }
        }
        val nativeTest by getting

        val linuxX64Main by getting
        val linuxX64Test by getting

        val mingwX64Main by getting
        val mingwX64Test by getting

        val macosX64Main by getting
        val macosX64Test by getting

        listOf(
            linuxX64Main,
            mingwX64Main,
            macosX64Main
        ).forEach { module ->
            module.dependsOn(nativeMain)
        }
    }
}