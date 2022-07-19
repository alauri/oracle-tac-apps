package com.oha.app;

import java.io.*;

import picocli.CommandLine;

import static org.junit.Assert.*;

import org.junit.Test;


/**
 * Unit test for OracleHA app.
 */
public class OracleHATest 
{
    OracleHA app = new OracleHA();
    CommandLine cmd = new CommandLine(app);
    StringWriter sw = new StringWriter();

    /**
     * Invoke the CLI with no commands.
     */
    @Test
    public void TestRootCommand() {
        cmd.setOut(new PrintWriter(sw));
        int exitCode = cmd.execute();

        assertEquals(0, exitCode);
        assertEquals("root called.\n", sw.toString());
    }

    /**
     * Invoke the command ``config``.
     */
    @Test
    public void TestConfigCommand() {
        cmd.setOut(new PrintWriter(sw));
        int exitCode = cmd.execute("config");

        assertEquals(0, exitCode);
        assertEquals("config called.\n", sw.toString());
    }

    /**
     * Invoke the command ``delete``.
     */
    @Test
    public void TestDeleteCommand() {
        cmd.setOut(new PrintWriter(sw));
        int exitCode = cmd.execute("delete");

        assertEquals(0, exitCode);
        assertEquals("delete called.\n", sw.toString());
    }

    /**
     * Invoke the command ``insert``.
     */
    @Test
    public void TestInsertCommand() {
        cmd.setOut(new PrintWriter(sw));
        int exitCode = cmd.execute("insert");

        assertEquals(0, exitCode);
        assertEquals("insert called.\n", sw.toString());
    }

    /**
     * Invoke the command ``reset``.
     */
    @Test
    public void TestResetCommand() {
        cmd.setOut(new PrintWriter(sw));
        int exitCode = cmd.execute("reset");

        assertEquals(0, exitCode);
        assertEquals("reset called.\n", sw.toString());
    }

    /**
     * Invoke the command ``update``.
     */
    @Test
    public void TestUpdateCommand() {
        cmd.setOut(new PrintWriter(sw));
        int exitCode = cmd.execute("update");

        assertEquals(0, exitCode);
        assertEquals("update called.\n", sw.toString());
    }
}
