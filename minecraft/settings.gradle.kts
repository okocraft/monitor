pluginManagement {
    includeBuild("build-logic")

    repositories {
        mavenCentral()
        gradlePluginPortal()
    }
}

rootProject.name = "monitor"
val prefix = rootProject.name.lowercase()

enableFeaturePreview("TYPESAFE_PROJECT_ACCESSORS")

sequenceOf(
    "core",

    "platform-paper",
    "platform-velocity",
).forEach {
    include("$prefix-$it")
    project(":$prefix-$it").projectDir = file(it)
}
