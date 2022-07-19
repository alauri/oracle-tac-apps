package com.oha.app.cmd;

import picocli.CommandLine.Command;
import picocli.CommandLine.Model.CommandSpec;
import picocli.CommandLine.Spec;

import java.io.File;

@Command(name = "insert",
         mixinStandardHelpOptions = true,
         description = "Insert new records into the table.")
public class Insert implements Runnable {

    @Spec CommandSpec spec;

    @Override
    public void run() {
        spec.commandLine().getOut().println("insert called.");
    }
}

