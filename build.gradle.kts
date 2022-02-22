val ktor_version: String by project
val kotlinx_cli_version: String by project
val mordant_version: String by project

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
    linuxX64 {
        binaries {
            executable {
                entryPoint = "main"
            }
        }
        compilations["main"].enableEndorsedLibs = true
    }
    mingwX64 {
        compilations["main"].enableEndorsedLibs = true

        binaries {
            executable {
                entryPoint = "main"
            }
        }
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
        val nativeMain by creating {
            dependencies {
                implementation("io.ktor:ktor-client-core:$ktor_version")
                implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:1.5.2-native-mt")
                implementation("io.ktor:ktor-client-serialization:$ktor_version")
                implementation("com.squareup.okio:okio:3.0.0")
                implementation("io.ktor:ktor-client-curl:$ktor_version")
                implementation("org.jetbrains.kotlinx:kotlinx-cli:$kotlinx_cli_version")
                implementation("com.github.ajalt.mordant:mordant:$mordant_version")
            }
        }
        val nativeTest by creating {
            dependencies {
                implementation(kotlin("test"))
            }
        }

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

        listOf(
            linuxX64Test,
            mingwX64Test,
            macosX64Test
        ).forEach { module ->
            module.dependsOn(nativeTest)
        }
    }
}