#!/usr/bin/python3


"""E2E test suite."""


from oha import cli


def test_root(runner) -> None:
    """Invoke the CLI with no commands.

    Returns:
        Nothing
    """
    result = runner.invoke(cli.cli, [])
    assert result.exit_code == 0
    assert result.output == "root called\n"


def test_config(runner) -> None:
    """Invoke the command ``config``.

    Returns:
        Nothing
    """
    result = runner.invoke(cli.cli, ["config"])
    assert result.exit_code == 0
    assert result.output == "root called\nconfig called\n"


def test_delete(runner) -> None:
    """Invoke the command ``delete``.

    Returns:
        Nothing
    """
    result = runner.invoke(cli.cli, ["delete"])
    assert result.exit_code == 0
    assert result.output == "root called\ndelete called\n"


def test_insert(runner) -> None:
    """Invoke the command ``insert``.

    Returns:
        Nothing
    """
    result = runner.invoke(cli.cli, ["insert"])
    assert result.exit_code == 0
    assert result.output == "root called\ninsert called\n"


def test_update(runner) -> None:
    """Invoke the command ``update``.

    Returns:
        Nothing
    """
    result = runner.invoke(cli.cli, ["update"])
    assert result.exit_code == 0
    assert result.output == "root called\nupdate called\n"


def test_reset(runner) -> None:
    """Invoke the command ``reset``.

    Returns:
        Nothing
    """
    result = runner.invoke(cli.cli, ["reset"])
    assert result.exit_code == 0
    assert result.output == "root called\nreset called\n"
