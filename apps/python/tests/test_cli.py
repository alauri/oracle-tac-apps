#!/usr/bin/python3


"""E2E test suite for the main command"""


from oha import cli

import json


def test_usage(runner, static) -> None:
    """Invoke the CLI with no commands and see how it works.

    Returns:
        Nothing
    """
    result = runner.invoke(cli.cli, ["-w", static])
    assert result.exit_code == 0
    assert result.output.startswith("Usage: ")


def test_config(runner, static) -> None:
    """Invoke the CLI by asking information about the current configuration.

    Returns:
        Nothing
    """
    result = runner.invoke(cli.cli, ["-w", static, "--config"])

    assert result.exit_code == 0
    assert result.output.startswith("{")
    assert "Usage:" not in result.output


def test_ping(runner, static) -> None:
    """Invoke the CLI and check the connection to the database.

    Returns:
        Nothing
    """
    result = runner.invoke(cli.cli, ["-w", static, "--ping"])

    assert result.exit_code == 0
    assert "Database reachable" in result.output


def test_error(runner, static) -> None:
    """Invoke the CLI with a wrong dsn value.

    Returns:
        Nothing
    """
    result = runner.invoke(cli.cli, ["-w", static, "-d", 0])
    assert result.exit_code == 2
    assert "Invalid value for '-d'" in result.output
