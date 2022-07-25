#!/usr/bin/python3


"""Tests for the module ``remove.py``"""


from oha.cli import cli

from tests.factory import MockResponse


def test_no_args(mocker, runner, static) -> None:
    """Invoke the command ``remove`` with no options.

    Returns:
        Nothing
    """
    MockResponse.fetchone = mocker.Mock(side_effect=[(1, ), (1, )])
    result = runner.invoke(cli, ["-w", static,
                                 "remove",
                                 "--delay", 0.05,
                                 "--older", 10])

    assert result.exit_code == 0

    output = [l for l in result.output.split("\n") if l]
    query = "DELETE FROM json_table " \
            "WHERE timestamp <= " \
            "to_date('2022-08-30 00:00:00','yyyy-mm-dd hh24:mi:ss')"

    assert output == [f"[1/1] - {query}",
                       "[1/1] - COMMIT"]


def test_iters(mocker, runner, static) -> None:
    """Invoke the command `remove` with the options ``iters``.

    Returns:
        Nothing
    """
    MockResponse.fetchone = mocker.Mock(side_effect=[(1, ), (1, )])
    result = runner.invoke(cli, ["-w", static, 'remove',
                                 "--iters", 5,
                                 "--delay", 0.05,
                                 "--commit-every", 2
                                 ])

    assert result.exit_code == 0

    output = [l for l in result.output.split("\n") if l]
    query = "DELETE FROM json_table " \
            "WHERE timestamp <= " \
            "to_date('2022-08-30 00:00:00','yyyy-mm-dd hh24:mi:ss')"

    assert output == [f"[1/5] - {query}",
                      f"[2/5] - {query}",
                       "[2/5] - COMMIT",
                      f"[3/5] - {query}",
                      f"[4/5] - {query}",
                       "[4/5] - COMMIT",
                      f"[5/5] - {query}",
                       "[5/5] - COMMIT"]
