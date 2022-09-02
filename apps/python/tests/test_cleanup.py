#!/usr/bin/python3


"""Tests for the module ``cleanup.py``"""


from oha.cli import cli

from tests.factory import MockResponse


def test_no_args(mocker, runner, static) -> None:
    """Invoke the command ``cleanup`` with no options.

    Returns:
        Nothing
    """
    MockResponse.fetchone = mocker.Mock(side_effect=[
        ('server1', 'vm1'), (1, ), (0, ),
        (1,2021,'Abu Dhabi','NaT|1|Car 1|Driver 1'),
        ('server1', 'vm1'), (1, ), (0, ),
    ])
    result = runner.invoke(cli, ["-w", static, '-d', 1, 'cleanup'])

    assert result.exit_code == 0
    assert ",'NaT',1,'Car 1','Driver 1')" in result.output

def test_args(mocker, runner, static) -> None:
    """Invoke the command ``cleanup`` with the options ``iters``.

    Returns:
        Nothing
    """
    MockResponse.fetchone = mocker.Mock(side_effect=[
        ('server1', 'vm1'), (1, ), (0, ),
        (1,2021,'Abu Dhabi','NaT|1|Car 1|Driver 1'),
        (2,2021,'Abu Dhabi','0 days 00:01:29.103000|2|Car 1|Driver 1'),
        (3,2021,'Abu Dhabi','0 days 00:01:28.827000|3|Car 1|Driver 1'),
        (4,2021,'Abu Dhabi','0 days 00:01:29.026000|4|Car 1|Driver 1'),
        (5,2021,'Abu Dhabi','0 days 00:01:28.718000|5|Car 1|Driver 1'),
        ('server1', 'vm1'), (1, ), (0, ),
    ])
    result = runner.invoke(cli, ["-w", static, '-d', 1, 'cleanup',
                                 "--iters", 5,
                                 "--delay", 0.05,
                                 "--commit-every", 2
                                 ])

    assert result.exit_code == 0
    assert ",'NaT',1,'Car 1','Driver 1')" in result.output
    assert ",'00:01:29.103000',2,'Car 1','Driver 1')" in result.output
    assert ",'00:01:28.827000',3,'Car 1','Driver 1')" in result.output
    assert ",'00:01:29.026000',4,'Car 1','Driver 1')" in result.output
    assert ",'00:01:28.718000',5,'Car 1','Driver 1')" in result.output
