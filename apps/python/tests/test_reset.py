#!/usr/bin/python3


"""Tests for the module ``reset.py``"""


from oha.cli import cli

from tests.factory import MockResponse


def test_no_args(mocker, runner, static) -> None:
    """Invoke the command ``reset`` with no options.

    Returns:
        Nothing
    """
    MockResponse.fetchone = mocker.Mock(side_effect=[(1, ), (1, )])
    result = runner.invoke(cli, ["-w", static, 'reset'])

    assert result.exit_code == 0
    assert result.output == "[+] - Database has been reset\n"
