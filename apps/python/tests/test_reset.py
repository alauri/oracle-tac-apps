#!/usr/bin/python3


"""Tests for the module ``reset.py``"""


from oha.cli import cli

from tests.factory import MockResponse


def test_no_args(mocker, runner, static) -> None:
    """Invoke the command ``reset`` with no options.

    Returns:
        Nothing
    """
    MockResponse.fetchone = mocker.Mock(side_effect=[
        ('server1', 'vm1'), (1, ), (0, ),
        ('server1', 'vm1'), (1, ), (0, ),
    ])
    result = runner.invoke(cli, ["-w", static, 'reset'])

    assert result.exit_code == 0
    assert result.output.count("('server1', 'vm1')") == 2
    assert result.output.count("TRUNCATE TABLE") == 2
    assert result.output.count("ALTER TABLE") == 2
    assert "All tables have been altered." in result.output
