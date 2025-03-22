plugins {
    id("monitor.common-conventions")
    id("monitor.bundle-conventions")
    alias(libs.plugins.run.task)
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

tasks {
    processResources {
        filesMatching(listOf("paper-plugin.yml")) {
            expand("projectVersion" to project.version)
        }
    }

    runServer {
        minecraftVersion(libs.versions.paper.get().removeSuffix("-R0.1-SNAPSHOT"))
    }
}
