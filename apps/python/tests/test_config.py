#!/usr/bin/python3


"""Tests for the module ``config.py``"""


import shutil
import os
import toml

from oha.cli import cli

from tests.factory import MockResponse


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


def test_no_args(mocker, runner, static) -> None:
    """Invoke the command ``config`` with no options.

    Returns:
        Nothing
    """
    MockResponse.fetchone = mocker.Mock(side_effect=[
        ('server1', 'vm1'), (1, ), (0, ),
        ('server1', 'vm1'), (1, ), (0, ),
    ])
    result = runner.invoke(cli, ["-w", static, "config"])

    assert result.exit_code == 0


def test_info(mocker, runner, static) -> None:
    """Invoke the command ``config`` with the option ``info``.

    Returns:
        Nothing
    """
    MockResponse.fetchone = mocker.Mock(side_effect=[
        ('server1', 'vm1'), (1, ), (0, ),
        ('server1', 'vm1'), (1, ), (0, ),
    ])
    result = runner.invoke(cli, ["-w", static, "config", "--info"])

    assert result.exit_code == 0
    assert result.output.startswith("(")
    assert "Usage:" not in result.output


def test_username(mocker, runner, static) -> None:
    """Invoke the sub-command ``driver`` with the option ``username``.

    Returns:
        Nothing
    """
    MockResponse.fetchone = mocker.Mock(side_effect=[
        ('server1', 'vm1'), (1, ), (0, ),
        ('server1', 'vm1'), (1, ), (0, ),
    ])
    result = runner.invoke(cli, ["-w", static,
                                 "config", "database",
                                 "--username", "fake"])

    assert result.exit_code == 0
    assert "Configuration updated" in result.output

    _toml = toml.load(os.path.join(static, "config.toml"))
    assert _toml["database"]["username"] == "fake"


def test_password(mocker, runner, static) -> None:
    """Invoke the sub-command ``driver`` with the option ``password``.

    Returns:
        Nothing
    """
    MockResponse.fetchone = mocker.Mock(side_effect=[
        ('server1', 'vm1'), (1, ), (0, ),
        ('server1', 'vm1'), (1, ), (0, ),
    ])
    result = runner.invoke(cli, ["-w", static,
                                 "config", "database",
                                 "--password", "fake"])

    assert result.exit_code == 0
    assert "Configuration updated" in result.output

    _toml = toml.load(os.path.join(static, "config.toml"))
    assert _toml["database"]["password"] == "fake"


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
