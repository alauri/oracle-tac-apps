#!/usr/bin/python3


"""Tests for the module ``injest.py``"""


from oha.cli import cli

from tests.factory import MockResponse


def test_no_args(mocker, runner, static) -> None:
    """Invoke the command ``injest`` with no options.

    Returns:
        Nothing
    """
    MockResponse.fetchone = mocker.Mock(side_effect=[
        ('server1', 'vm1'), (1, ), (0, ),
        ('server1', 'vm1'), (1, ), (0, ),
    ])
    result = runner.invoke(cli, ["-w", static,
                                 "injest"])

    assert result.exit_code == 0
    assert "(2021,'Abu Dhabi','NaT|1|Car 1|Driver 1')" in result.output
    assert result.output.count("COMMIT") == 1
    assert result.output.count("('server1', 'vm1')") == 2


def test_args(mocker, runner, static) -> None:
    """Invoke the command ``injest`` with the options ``iters``.

    Returns:
        Nothing
    """
    MockResponse.fetchone = mocker.Mock(side_effect=[
        ('server1', 'vm1'), (1, ), (0, ),
        ('server1', 'vm1'), (1, ), (0, ),
    ])
    result = runner.invoke(cli, ["-w", static,
                                 "injest",
                                 "--iters", 5,
                                 "--delay", 0.05,
                                 "--commit-every", 5])

    assert result.exit_code == 0
    assert "(2021,'Abu Dhabi','NaT|1|Car 1|Driver 1')" in result.output
    assert "0 days 00:01:29.103000|2|Car 1|Driver 1')" in result.output
    assert "0 days 00:01:28.827000|3|Car 1|Driver 1')" in result.output
    assert "0 days 00:01:29.026000|4|Car 1|Driver 1')" in result.output
    assert "0 days 00:01:28.718000|5|Car 1|Driver 1')" in result.output
    assert result.output.count("COMMIT") == 1
    assert result.output.count("('server1', 'vm1')") == 2
