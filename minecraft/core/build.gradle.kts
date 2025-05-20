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
    implementation(libs.codec4j.api)
    implementation(libs.codec4j.io.yaml)
    compileOnly(libs.concurrent.util)
}
