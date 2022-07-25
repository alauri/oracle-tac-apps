#!/usr/bin/python3


"""Tests for the module ``reset.py``"""


from oha.cli import cli

from tests.factory import MockResponse


def test_no_args(mocker, runner, static) -> None:
    """Invoke the command ``reset`` with no options.

    Returns:
        Nothing
    """
    MockResponse.fetchone = mocker.Mock(side_effect=[(1, ), (0, )])
    result = runner.invoke(cli, ["-w", static, 'reset'])

    assert result.exit_code == 0
    output = [l for l in result.output.split("\n") if l]
    assert output == [
        "[ALTERING] - TRUNCATE TABLE raw_table",
        "[ALTERING] - ALTER TABLE raw_table MODIFY(ID GENERATED AS IDENTITY " \
            "(START WITH 1))",
        "[ALTERING] - TRUNCATE TABLE json_table",
        "[ALTERING] - ALTER TABLE json_table MODIFY(ID GENERATED AS IDENTITY " \
            "(START WITH 1))",
        "[ALTERING] - All tables have been altered."]
