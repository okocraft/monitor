package net.okocraft.monitor.core.command;

import java.util.List;
import java.util.concurrent.CompletableFuture;

public interface Command {

    void execute(CommandSender sender, String[] args);

    default CompletableFuture<List<String>> tabComplete(CommandSender sender, String[] args) {
        return CompletableFuture.completedFuture(List.of());
    }
}
