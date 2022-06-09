#!/usr/bin/python3


"""E2E test suite for the main command."""


from oha import cli


def test_cli(runner, static) -> None:
    """Invoke the CLI with a different working folder.

    Returns:
        Nothing
    """
    result = runner.invoke(cli.cli, ["-w", static])
    assert result.exit_code == 2
    assert "Error: Missing command." in result.output
