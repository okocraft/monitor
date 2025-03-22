plugins {
    id("monitor.common-conventions")
}

project.extra["monitor.plugin-name"] = "Monitor-Paper"

repositories {
    maven {
        url = uri("https://repo.papermc.io/repository/maven-public/")
    }
}

dependencies {
    implementation(projects.monitorCore)
    compileOnly(libs.platform.paper)
}
