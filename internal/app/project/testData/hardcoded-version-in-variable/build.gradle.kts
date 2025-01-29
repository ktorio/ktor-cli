plugins {
    kotlin("jvm") version "2.0.20"
}

repositories {
    mavenCentral()
}

val ktorVersion = "3.0.1"

dependencies {
    implementation("group:artifact:1.2.3")
    implementation("io.ktor:ktor-server-core:$ktorVersion")
}
