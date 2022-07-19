#!/usr/bin/python3


"""Tests for the module ``remove.py``"""


from oha.cli import cli


def test_no_args(runner, static) -> None:
    """Invoke the command ``remove`` with no options.

    Returns:
        Nothing
    """
    result = runner.invoke(cli, ["-w", static, 'remove'])

    assert result.exit_code == 0

    output = [l for l in result.output.split("\n") if l]
    assert output == ['[1/1] - DELETE FROM test WHERE DEPARTMENT_ID=1',
                      '[1/1] - COMMIT']


def test_iters(runner, static) -> None:
    """Invoke the command `remove` with the options ``iters``.

    Returns:
        Nothing
    """
    result = runner.invoke(cli, ["-w", static, 'remove',
                                 "--iters", 5,
                                 "--delay", 0.05,
                                 "--commit-every", 2
                                 ])

    assert result.exit_code == 0

    output = [l for l in result.output.split("\n") if l]
    assert output == ['[1/5] - DELETE FROM test WHERE DEPARTMENT_ID=1',
                      '[2/5] - DELETE FROM test WHERE DEPARTMENT_ID=2',
                      '[2/5] - COMMIT',
                      '[3/5] - DELETE FROM test WHERE DEPARTMENT_ID=3',
                      '[4/5] - DELETE FROM test WHERE DEPARTMENT_ID=4',
                      '[4/5] - COMMIT',
                      '[5/5] - DELETE FROM test WHERE DEPARTMENT_ID=5',
                      '[5/5] - COMMIT']