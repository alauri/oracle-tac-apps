#!/usr/bin/python3


"""E2E test suite for the main command"""


from otac import cli
from tests.factory import MockResponse


def test_usage(runner) -> None:
    """Invoke the CLI with no commands, expecting an 'Usage ...' message.

    Returns:
        Nothing
    """
    result = runner.invoke(cli.cli, [])
    assert result.exit_code == 0
    assert result.output.startswith("Usage: ")


def test_config(runner) -> None:
    """Invoke the CLI by asking information about the current configuration.

    Returns:
        Nothing
    """
    result = runner.invoke(cli.cli, ["--config"])

    assert result.exit_code == 0
    assert result.output.startswith("{")
    assert "Usage:" not in result.output


def test_ping(mocker, runner) -> None:
    """Invoke the CLI by asking to ping the database.

    Returns:
        Nothing
    """
    MockResponse.fetchone = mocker.Mock(
        side_effect=[
            ("server1", "vm1"),
            ("server1", "vm1"),
        ]
    )
    result = runner.invoke(cli.cli, ["-d", "localhost", "--ping"])

    assert result.exit_code == 0
    assert "[+] - Database reachable" in result.output


def test_error(runner) -> None:
    """Invoke the CLI with a wrong dsn value.

    Returns:
        Nothing
    """
    result = runner.invoke(cli.cli, ["-d", "wrong", "ingest"])
    assert result.exit_code == 2
    assert "Invalid value for '-d' / '--dsn'" in result.output
