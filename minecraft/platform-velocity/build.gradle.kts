plugins {
    alias(libs.plugins.bundler)
}

jcommon {
    setupPaperRepository()
    commonDependencies {
        implementation(projects.monitorCore)
        compileOnly(libs.platform.velocity)
    }
}

bundler {
    copyToRootBuildDirectory("Monitor-Velocity-${project.version}")
    replacePluginVersionForVelocity(project.version)
}

