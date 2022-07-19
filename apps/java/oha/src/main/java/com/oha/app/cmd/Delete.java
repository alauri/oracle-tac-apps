package com.oha.app.cmd;

import picocli.CommandLine.Command;
import picocli.CommandLine.Model.CommandSpec;
import picocli.CommandLine.Spec;

import java.io.File;

@Command(name = "delete",
         mixinStandardHelpOptions = true,
         description = "Delete records from the table.")
public class Delete implements Runnable {

    @Spec CommandSpec spec;

    @Override
    public void run() {
        spec.commandLine().getOut().println("delete called.");
    }
}

