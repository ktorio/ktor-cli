plugins {
    kotlin("jvm") version "2.0.20"
}

repositories {
    mavenCentral()
}

val ktorVersion: String by project

dependencies {
    implementation("io.ktor:ktor-server-core-jvm:$ktorVersion")
    implementation("io.ktor:ktor-server-content-negotiation:$ktorVersion")
}
