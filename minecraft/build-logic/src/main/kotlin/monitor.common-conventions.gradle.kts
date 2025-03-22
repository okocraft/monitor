plugins {
    `java-library`
}

repositories {
    mavenCentral()
}

val libs = extensions.getByType(org.gradle.accessors.dm.LibrariesForLibs::class)

dependencies {
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

val javaVersion = JavaVersion.VERSION_21
val charset = Charsets.UTF_8

java {
    sourceCompatibility = javaVersion
    targetCompatibility = javaVersion
}

tasks {
    compileJava {
        options.encoding = charset.name()
        options.release.set(javaVersion.ordinal + 1)
    }

    processResources {
        filteringCharset = charset.name()
    }

    test {
        useJUnitPlatform()

        testLogging {
            events("passed", "skipped", "failed")
        }
    }
}
