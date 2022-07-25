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
        (1, ), (0, ),
        (None, 1662674400, 3, "Right=False|Left=True")
    ])
    result = runner.invoke(cli, ["-w", static, 'cleanup'])
    assert result.exit_code == 0

    output = [l for l in result.output.split("\n") if l]
    assert len(output) == 3
    assert output[0].startswith("[1/1] - SELECT * FROM raw_table WHERE id=1")
    assert output[1].startswith("[1/1] - INSERT INTO json_table(timestamp")
    assert output[-1] == "[1/1] - COMMIT"


def test_args(mocker, runner, static) -> None:
    """Invoke the command ``cleanup`` with the options ``iters``.

    Returns:
        Nothing
    """
    MockResponse.fetchone = mocker.Mock(side_effect=[
        (1, ), (0, ),
        (None, 1662674400, 3, "Right=False|Left=True"),
        (None, 1658351188, 2, "Right=False|Left=True"),
        (None, 1658351188, 2, "Right=False|Left=True"),
        (None, 1658351188, 1, "Right=False|Left=True"),
        (None, 1659996000, 1, "Right=False|Left=True")
    ])
    result = runner.invoke(cli, ["-w", static, 'cleanup',
                                 "--iters", 5,
                                 "--delay", 0.05,
                                 "--commit-every", 2
                                 ])

    assert result.exit_code == 0

    output = [l for l in result.output.split("\n") if l]
    assert len(output) == 13
    assert output[0].startswith("[1/5] - SELECT * FROM raw_table WHERE id=1")
    assert output[1].startswith("[1/5] - INSERT INTO json_table(timestamp")

    assert output[2].startswith("[2/5] - SELECT * FROM raw_table WHERE id=2")
    assert output[3].startswith("[2/5] - INSERT INTO json_table(timestamp")
    assert output[4] == "[2/5] - COMMIT"

    assert output[5].startswith("[3/5] - SELECT * FROM raw_table WHERE id=3")
    assert output[6].startswith("[3/5] - INSERT INTO json_table(timestamp")

    assert output[7].startswith("[4/5] - SELECT * FROM raw_table WHERE id=4")
    assert output[8].startswith("[4/5] - INSERT INTO json_table(timestamp")
    assert output[9] == "[4/5] - COMMIT"

    assert output[10].startswith("[5/5] - SELECT * FROM raw_table WHERE id=5")
    assert output[11].startswith("[5/5] - INSERT INTO json_table(timestamp")
    assert output[12] == "[5/5] - COMMIT"
