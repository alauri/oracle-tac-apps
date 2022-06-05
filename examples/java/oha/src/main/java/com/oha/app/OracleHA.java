package com.oha.app;

import com.oha.app.cmd.*;

import picocli.CommandLine;
import picocli.CommandLine.Command;
import picocli.CommandLine.Model.CommandSpec;
import picocli.CommandLine.Spec;

import java.io.File;
import java.util.concurrent.Callable;


@Command(name = "oracle-ha",
         mixinStandardHelpOptions = true,
         version = "oracle-ha 0.1.0",
         description = "A Java application for the High Availability in Oracle.",
         subcommands = {
             Config.class,
             Delete.class,
             Insert.class,
             Reset.class,
             Update.class
         })
class OracleHA implements Callable<Integer> {

    @Spec CommandSpec spec;

    @Override
    public Integer call() throws Exception {
        spec.commandLine().getOut().println("root called.");
        return 0;
    }

    public static void main(String... args) {
        int exitCode = new CommandLine(new OracleHA()).execute(args);
        System.exit(exitCode);
    }
}
