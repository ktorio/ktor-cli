plugins {
    kotlin("jvm") version "2.0.20"
}

repositories {
    mavenCentral()
}

dependencies {
    implementation("io.ktor:ktor-server-webjars")
    implementation(platform("io.ktor:ktor-bom:2.3.13"))
}
