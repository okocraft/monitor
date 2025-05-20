plugins {
    alias(libs.plugins.jcommon)
}

jcommon {
    javaVersion = JavaVersion.VERSION_21

    commonDependencies {
        compileOnlyApi(libs.annotations)
        compileOnlyApi(libs.adventure)
        compileOnlyApi(libs.adventure.text.serializer.gson)
        compileOnlyApi(libs.adventure.text.serializer.plain)
        compileOnlyApi(libs.slf4j)

        testImplementation(platform(libs.junit.bom))
        testImplementation(libs.junit.jupiter)
        testImplementation(libs.adventure)
        testImplementation(libs.adventure.text.serializer.gson)
        testImplementation(libs.adventure.text.serializer.plain)
        testRuntimeOnly("org.junit.platform:junit-platform-launcher")
        testRuntimeOnly(libs.slf4j.simple)
    }
}
