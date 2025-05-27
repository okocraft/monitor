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
            relocate("ca.spottedleaf.concurrentutil", "net.okocraft.monitor.lib.concurrentutil")
            relocate("it.unimi.dsi.fastutil", "net.okocraft.monitor.lib.fastutil")
            relocate("com.zaxxer.hikari", "net.okocraft.monitor.lib.hikari")
            relocate("com.mysql", "net.okocraft.monitor.lib.mysql")
            relocate("google.protobuf", "net.okocraft.monitor.lib.protobuf")
            relocate("com.google.protobuf", "net.okocraft.monitor.lib.protobuf")
            relocate("dev.siroshun.jfun", "net.okocraft.monitor.lib.jfun")
            relocate("dev.siroshun.codec4j", "net.okocraft.monitor.lib.codec4j")
            exclude("net.okocraft.monitor.platform.velocity.MonitorVelocity")
            exclude(dependency(libs.mysql.get()))
        }
    }
}
