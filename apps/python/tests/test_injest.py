#!/usr/bin/python3


"""Tests for the module ``injest.py``"""


from oha.cli import cli


def test_no_args(runner, static) -> None:
    """Invoke the command ``injest`` with no options.

    Returns:
        Nothing
    """
    result = runner.invoke(cli, ["-w", static, 'injest'])

    assert result.exit_code == 0

    output = [l for l in result.output.split("\n") if l]
    assert output == ['[1/1] - INSERT INTO test(DEPARTMENT_ID, DEPARTMENT_NAME) VALUES(1, \'pippo\')',
                      '[1/1] - COMMIT']


def test_args(runner, static) -> None:
    """Invoke the command ``delete`` with the options ``iters``.

    Returns:
        Nothing
    """
    result = runner.invoke(cli, ["-w", static, 'injest',
                                 "--iters", 5,
                                 "--delay", 0.05,
                                 "--commit-every", 2
                                 ])

    assert result.exit_code == 0

    output = [l for l in result.output.split("\n") if l]
    assert output == ['[1/5] - INSERT INTO test(DEPARTMENT_ID, DEPARTMENT_NAME) VALUES(1, \'pippo\')',
                      '[2/5] - INSERT INTO test(DEPARTMENT_ID, DEPARTMENT_NAME) VALUES(2, \'pippo\')',
                      '[2/5] - COMMIT',
                      '[3/5] - INSERT INTO test(DEPARTMENT_ID, DEPARTMENT_NAME) VALUES(3, \'pippo\')',
                      '[4/5] - INSERT INTO test(DEPARTMENT_ID, DEPARTMENT_NAME) VALUES(4, \'pippo\')',
                      '[4/5] - COMMIT',
                      '[5/5] - INSERT INTO test(DEPARTMENT_ID, DEPARTMENT_NAME) VALUES(5, \'pippo\')',
                      '[5/5] - COMMIT']
