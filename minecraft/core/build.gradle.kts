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
    implementation(libs.hikaricp) {
        exclude("org.slf4j", "slf4j-api")
    }
    implementation(libs.codec4j.api)
    implementation(libs.codec4j.io.yaml) {
        exclude("org.yaml", "snakeyaml")
    }
    implementation(libs.codec4j.io.gson) {
        exclude("com.google.code.gson", "gson")
    }
    implementation(libs.minio) {
        exclude("com.google.guava", "guava")
        exclude("org.jetbrains", "annotations")
    }

    compileOnly(libs.concurrent.util) {
        exclude("org.slf4j", "slf4j-api")
    }
}
