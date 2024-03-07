#!/usr/bin/python3


"""Tests for the module ``ingest.py``"""


from otac.cli import cli

from tests.factory import MockResponse


def test_args(mocker, runner) -> None:
    """Invoke the command ``ingest`` with the options ``iters``.

    Returns:
        Nothing
    """
    MockResponse.fetchone = mocker.Mock(
        side_effect=[
            ("server1", "vm1"),
            ("server1", "vm1"),
        ]
    )
    result = runner.invoke(
        cli,
        [
            "-d",
            "localhost",
            "ingest",
            "--delay",
            0.05,
            "--commit-every",
            5,
        ],
    )

    assert result.exit_code == 0
    assert "(2021,'Abu Dhabi','NaT|1|Car 1|Driver 1')" in result.output
    assert "0 days 00:01:29.103000|2|Car 1|Driver 1')" in result.output
    assert "0 days 00:01:28.827000|3|Car 1|Driver 1')" in result.output
    assert "0 days 00:01:29.026000|4|Car 1|Driver 1')" in result.output
    assert "0 days 00:01:28.718000|5|Car 1|Driver 1')" in result.output
    assert result.output.count("COMMIT") == 2
    assert result.output.count("('server1', 'vm1')") == 2
