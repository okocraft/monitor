plugins {
    id("monitor.common-conventions")
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

dependencies {
    implementation(libs.hikaricp)
    implementation(libs.configapi.format.yaml)
    compileOnly(libs.concurrent.util)
}
