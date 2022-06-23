#!/usr/bin/python3


"""E2E test suite for the main command"""


from oha import cli


def test_cli(runner, static) -> None:
    """Invoke the CLI with a different working folder.

    Returns:
        Nothing
    """
    result = runner.invoke(cli.cli, ["-w", static, "-d", 5])
    assert result.exit_code == 2
    assert "Error: Missing command." in result.output


def test_cli_error(runner, static) -> None:
    """Invoke the CLI with a wrong dsn value.

    Returns:
        Nothing
    """
    result = runner.invoke(cli.cli, ["-w", static, "-d", 0])
    assert result.exit_code == 2
    assert "Invalid value for '-d'" in result.output
