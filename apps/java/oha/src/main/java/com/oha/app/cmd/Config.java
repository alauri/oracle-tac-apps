package com.oha.app.cmd;

import picocli.CommandLine.Command;
import picocli.CommandLine.Model.CommandSpec;
import picocli.CommandLine.Spec;

import java.io.File;

@Command(name = "config",
         mixinStandardHelpOptions = true,
         description = "Configure the application.")
public class Config implements Runnable {

    @Spec CommandSpec spec;

    @Override
    public void run() {
        spec.commandLine().getOut().println("config called.");
    }
}

