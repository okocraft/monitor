package net.okocraft.monitor.core.command;

public class ArgumentList {

    private final String[] args;
    private int currentIndex;

    public ArgumentList(String[] args) {
        this.args = args;
    }

    public String next() {
        if (!this.hasNext()) {
            throw new IllegalStateException("No more args");
        }
        return this.args[this.currentIndex++];
    }

    public boolean hasNext() {
        return this.currentIndex < this.args.length;
    }
}
