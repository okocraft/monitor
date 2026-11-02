# Monitor for NeoForge

This server-side module targets Minecraft 26.2, NeoForge 26.2.0.88, and Java 25. It uses the existing Monitor core and database schema. Build it with `gradlew :monitor-platform-neoforge:build` from `minecraft/`.

Place the resulting jar in the server's `mods/` directory and put `config.yml` in the NeoForge server's `config/monitor/` directory. The configuration format matches the Paper and Velocity modules. Set `server.name` to a distinct name for each server. The module requires the same MySQL schema as Monitor's existing installations; `upload` and `discord-webhook` retain the core configuration format.

The module records player connections, chat, commands, item renames, sign edits, and world identity. Ore notifications use the configured block keys and formats. The `/monitor` command is available when `command.enabled` is true; console operators and server operators can use it, and LuckPerms nodes are honored when LuckPerms is installed. Server status polling remains a Velocity responsibility. The module does not add a HUD or scoreboard.
