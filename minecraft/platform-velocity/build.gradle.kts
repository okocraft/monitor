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
            relocate("com.fasterxml.jackson", "net.okocraft.monitor.lib.jackson")
            relocate("com.google", "net.okocraft.monitor.lib.google")
            relocate("com.mysql", "net.okocraft.monitor.lib.mysql")
            relocate("com.zaxxer.hikari", "net.okocraft.monitor.lib.hikari")
            relocate("dev.siroshun.jfun", "net.okocraft.monitor.lib.jfun")
            relocate("dev.siroshun.codec4j", "net.okocraft.monitor.lib.codec4j")
            relocate("google.protobuf", "net.okocraft.monitor.lib.protobuf")
            relocate("it.unimi.dsi.fastutil", "net.okocraft.monitor.lib.fastutil")
            relocate("io.minio", "net.okocraft.monitor.lib.minio")
            relocate("javax", "net.okocraft.monitor.lib.javax")
            relocate("kotlin", "net.okocraft.monitor.lib.kotlin")
            relocate("okhttp3", "net.okocraft.monitor.lib.okhttp3")
            relocate("okio", "net.okocraft.monitor.lib.okio")
            relocate("org.apache", "net.okocraft.monitor.lib.apache")
            relocate("org.bouncycastle", "net.okocraft.monitor.lib.bouncycastle")
            relocate("org.checkerframework", "net.okocraft.monitor.lib.checkerframework")
            relocate("org.simpleframework", "net.okocraft.monitor.lib.simpleframework")
            relocate("org.xerial.snappy", "net.okocraft.monitor.lib.snappy")
            exclude("net.okocraft.monitor.platform.velocity.MonitorVelocity")
            exclude(dependency(libs.mysql.get()))
        }
    }
}
