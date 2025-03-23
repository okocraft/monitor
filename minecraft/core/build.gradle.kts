plugins {
    id("monitor.common-conventions")
}

dependencies {
    implementation(libs.hikaricp)
    implementation(libs.configapi.format.yaml)
}
