#!/usr/bin/python3


"""Tests for the module ``update.py``"""


from oha.cli import cli


def test_no_args(runner, static) -> None:
    """Invoke the command ``update`` with no options.

    Returns:
        Nothing
    """
    result = runner.invoke(cli, ["-w", static, 'update'])

    assert result.exit_code == 0

    output = [l for l in result.output.split("\n") if l]
    assert output == ['[1/1] - UPDATE test SET DEPARTMENT_NAME=\'pippo1\' WHERE DEPARTMENT_ID=2',
                      '[1/1] - COMMIT']


def test_args(runner, static) -> None:
    """Invoke the command ``update`` with the options ``iters``.

    Returns:
        Nothing
    """
    result = runner.invoke(cli, ["-w", static, 'update',
                                 "--iters", 5,
                                 "--delay", 0.05,
                                 "--commit-every", 2
                                 ])

    assert result.exit_code == 0

    output = [l for l in result.output.split("\n") if l]
    assert output == [
        "[1/5] - UPDATE test SET DEPARTMENT_NAME=\'pippo1\' WHERE DEPARTMENT_ID=2",
        "[2/5] - UPDATE test SET DEPARTMENT_NAME=\'pippo2\' WHERE DEPARTMENT_ID=3",
        '[2/5] - COMMIT',
        "[3/5] - UPDATE test SET DEPARTMENT_NAME=\'pippo3\' WHERE DEPARTMENT_ID=4",
        "[4/5] - UPDATE test SET DEPARTMENT_NAME=\'pippo4\' WHERE DEPARTMENT_ID=5",
        '[4/5] - COMMIT',
        "[5/5] - UPDATE test SET DEPARTMENT_NAME=\'pippo5\' WHERE DEPARTMENT_ID=6",
        '[5/5] - COMMIT'
    ]
