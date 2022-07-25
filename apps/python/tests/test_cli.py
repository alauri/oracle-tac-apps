#!/usr/bin/python3


"""E2E test suite for the main command"""


from oha import cli

import json


def test_cli_info(runner, static) -> None:
    """Invoke the CLI by asking information about the configuration.

    Returns:
        Nothing
    """
    result = runner.invoke(cli.cli, ["-w", static, "--info"])
    assert result.exit_code == 0
    assert result.output.startswith("{")
    assert result.output.endswith("}\n")


def test_cli_error(runner, static) -> None:
    """Invoke the CLI with a wrong dsn value.

    Returns:
        Nothing
    """
    result = runner.invoke(cli.cli, ["-w", static, "-d", 0])
    assert result.exit_code == 2
    assert "Invalid value for '-d'" in result.output
