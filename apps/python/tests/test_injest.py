#!/usr/bin/python3


"""Tests for the module ``injest.py``"""


from oha.cli import cli


def test_no_args(runner, static) -> None:
    """Invoke the command ``injest`` with no options.

    Returns:
        Nothing
    """
    result = runner.invoke(cli, ["-w", static,
                                 "injest"])

    assert result.exit_code == 0

    output = [l for l in result.output.split("\n") if l]
    assert output == [
        "[1/1] - INSERT INTO raw_table(timestamp,sensorid,data) " \
            "VALUES(1658351188,3,Right=False|Left=True)",
        "[1/1] - COMMIT"]


def test_args(runner, static) -> None:
    """Invoke the command ``injest`` with the options ``iters``.

    Returns:
        Nothing
    """
    result = runner.invoke(cli, ["-w", static,
                                 "injest",
                                 "--iters", 5,
                                 "--delay", 0.05,
                                 "--commit-every", 5])

    assert result.exit_code == 0

    output = [l for l in result.output.split("\n") if l]
    assert output == [
        "[1/5] - INSERT INTO raw_table(timestamp,sensorid,data) " \
            "VALUES(1658351188,3,Right=False|Left=True)",
        "[2/5] - INSERT INTO raw_table(timestamp,sensorid,data) " \
            "VALUES(1658351188,2,Right=False|Left=True)",
        "[3/5] - INSERT INTO raw_table(timestamp,sensorid,data) " \
            "VALUES(1658351188,2,Right=False|Left=True)",
        "[4/5] - INSERT INTO raw_table(timestamp,sensorid,data) " \
            "VALUES(1658351188,1,Right=False|Left=True)",
        "[5/5] - INSERT INTO raw_table(timestamp,sensorid,data) " \
            "VALUES(1658351188,1,Right=False|Left=True)",
        "[5/5] - COMMIT"]
