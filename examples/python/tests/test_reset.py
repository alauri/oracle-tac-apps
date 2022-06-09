#!/usr/bin/python3


"""Tests for the module ``reset.py``"""


from oha.cli import cli


def test_no_args(runner, static) -> None:
    """Invoke the command ``reset`` with no options.

    Returns:
        Nothing
    """
    result = runner.invoke(cli, ["-w", static, 'update'])

    assert result.exit_code == 0

    assert result.output == "[+] - Database has been reset"
