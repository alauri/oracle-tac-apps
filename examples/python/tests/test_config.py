#!/usr/bin/python3


"""Tests for the module ``config.py``"""


import shutil
import os

from oha.cli import cli, OracleHA


def setup_module(module) -> None:
    """Setup for the current module.

    Temporarly duplicate the TOML configuration file.

    Returns:
        Nothing
    """
    from .conftest import get_static
    src = os.path.join(get_static(), "config.toml")
    dst = os.path.join(get_static(), "config.bak")
    shutil.copyfile(src, dst)


def test_no_args(runner, static) -> None:
    """Invoke the command ``config`` with no options.

    Returns:
        Nothing
    """
    result = runner.invoke(cli, ["-w", static, "config"])

    assert result.exit_code == 0


def test_info(runner, static) -> None:
    """Invoke the command ``config`` with the option ``info``.

    Returns:
        Nothing
    """
    result = runner.invoke(cli, ["-w", static, "config", "--info"])

    assert result.exit_code == 0
    assert "Current configuration" in result.output


def test_username(runner, static) -> None:
    """Invoke the sub-command ``driver`` with the option ``username``.

    Returns:
        Nothing
    """
    result = runner.invoke(cli, ["-w", static,
                                 "config", "driver",
                                 "--username", "fake"])

    assert result.exit_code == 0
    assert "Configuration updated" in result.output

    _toml = OracleHA.read_toml()
    assert _toml["driver"]["username"] == "fake"


def test_password(runner, static) -> None:
    """Invoke the sub-command ``driver`` with the option ``password``.

    Returns:
        Nothing
    """
    result = runner.invoke(cli, ["-w", static,
                                 "config", "driver",
                                 "--password", "fake"])

    assert result.exit_code == 0
    assert "Configuration updated" in result.output

    _toml = OracleHA.read_toml()
    assert _toml["driver"]["password"] == "fake"


def test_table(runner, static) -> None:
    """Invoke the sub-command ``database`` with the option ``table``.

    Returns:
        Nothing
    """
    result = runner.invoke(cli, ["-w", static,
                                 "config", "database",
                                 "--table", "fake"])

    assert result.exit_code == 0
    assert "Configuration updated" in result.output

    _toml = OracleHA.read_toml()
    assert _toml["database"]["table"] == "fake"


def teardown_module(module) -> None:
    """Teardown for the current module.

    Delete temporary copy of the TOML configuration file.

    Returns:
        Nothing
    """
    from .conftest import get_static
    src = os.path.join(get_static(), "config.bak")
    dst = os.path.join(get_static(), "config.toml")
    shutil.copyfile(src, dst)
    os.remove(src)
