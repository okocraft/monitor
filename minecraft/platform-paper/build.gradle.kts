plugins {
    alias(libs.plugins.bundler)
    alias(libs.plugins.run.paper)
}

jcommon {
    setupPaperRepository()
    commonDependencies {
        implementation(projects.monitorCore)
        compileOnly(libs.platform.paper)

        testImplementation(libs.platform.paper)
    }
}

bundler {
    copyToRootBuildDirectory("Monitor-Paper-${project.version}")
    replacePluginVersionForPaper(project.version)
}

tasks {
    runServer {
        minecraftVersion(libs.versions.paper.get().removeSuffix("-R0.1-SNAPSHOT"))
    }
    shadowJar {
        minimize {
            exclude("net.okocraft.monitor.platform.paper.MonitorPaper")
        }
    }
}
