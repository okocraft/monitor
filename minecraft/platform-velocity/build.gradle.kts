plugins {
    id("monitor.common-conventions")
}

project.extra["monitor.plugin-name"] = "Monitor-Velocity"

repositories {
    maven {
        url = uri("https://repo.papermc.io/repository/maven-public/")
    }
}

dependencies {
    implementation(projects.monitorCore)
    compileOnly(libs.platform.velocity)
}

