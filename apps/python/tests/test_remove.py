#!/usr/bin/python3


"""Tests for the module ``remove.py``"""


from otac.cli import cli

from tests.factory import MockResponse


def test_no_args(mocker, runner, static) -> None:
    """Invoke the command ``remove`` with no options.

    Returns:
        Nothing
    """
    MockResponse.fetchone = mocker.Mock(side_effect=[
        ('server1', 'vm1'), (1, ), (0, ),
        ('server1', 'vm1'), (1, ), (0, ),
    ])
    result = runner.invoke(cli, ["-w", static, '-d', 'localhost',
                                 "remove",
                                 "--delay", 0.05])

    assert result.exit_code == 0
    assert result.output.count("FROM json_tel WHERE LapTime='NaT'") == 1
    assert result.output.count("COMMIT") == 1
    assert result.output.count("('server1', 'vm1')") == 2


def test_iters(mocker, runner, static) -> None:
    """Invoke the command `remove` with the options ``iters``.

    Returns:
        Nothing
    """
    MockResponse.fetchone = mocker.Mock(side_effect=[
        ('server1', 'vm1'), (1, ), (0, ),
        ('server1', 'vm1'), (1, ), (0, ),
    ])
    result = runner.invoke(cli, ["-w", static, '-d', 'localhost', 'remove',
                                 "--iters", 5,
                                 "--delay", 0.05,
                                 "--commit-every", 2
                                 ])

    assert result.exit_code == 0
    assert result.output.count("FROM json_tel WHERE LapTime='NaT'") == 5
    assert result.output.count("COMMIT") == 3
    assert result.output.count("('server1', 'vm1')") == 2
