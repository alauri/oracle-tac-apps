package com.oha.app.cmd;

import picocli.CommandLine.Command;
import picocli.CommandLine.Model.CommandSpec;
import picocli.CommandLine.Spec;

import java.io.File;

@Command(name = "update",
         mixinStandardHelpOptions = true,
         description = "Update records into the table.")
public class Update implements Runnable {

    @Spec CommandSpec spec;

    @Override
    public void run() {
        spec.commandLine().getOut().println("update called.");
    }
}

