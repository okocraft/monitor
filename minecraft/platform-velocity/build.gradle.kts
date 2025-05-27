plugins {
    alias(libs.plugins.bundler)
    alias(libs.plugins.run.velocity)
}

repositories {
    maven {
        name = "paper"
        url = uri("https://repo.papermc.io/repository/maven-public/")
        content {
            includeGroup("ca.spottedleaf")
        }
    }
}

jcommon {
    setupPaperRepository()
    commonDependencies {
        implementation(projects.monitorCore)
        implementation(libs.mysql) // Velocity does not bundle MySQL driver
        implementation(libs.concurrent.util) {
            exclude("org.slf4j", "slf4j-api")
        }
        compileOnly(libs.platform.velocity)
    }
}

bundler {
    copyToRootBuildDirectory("Monitor-Velocity-${project.version}")
    replacePluginVersionForVelocity(project.version)
}

tasks {
    runVelocity {
        velocityVersion(libs.versions.velocity.get())
    }
    shadowJar {
        minimize {
            exclude("net.okocraft.monitor.platform.velocity.MonitorVelocity",)
            exclude(dependency(libs.mysql.get()))
        }
    }
}
