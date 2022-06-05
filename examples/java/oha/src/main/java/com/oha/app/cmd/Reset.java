package com.oha.app.cmd;

import picocli.CommandLine.Command;
import picocli.CommandLine.Model.CommandSpec;
import picocli.CommandLine.Spec;

import java.io.File;

@Command(name = "reset",
         mixinStandardHelpOptions = true,
         description = "Reset the data to factory.")
public class Reset implements Runnable {

    @Spec CommandSpec spec;

    @Override
    public void run() {
        spec.commandLine().getOut().println("reset called.");
    }
}

