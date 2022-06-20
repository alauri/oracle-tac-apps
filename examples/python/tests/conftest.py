#!/usr/bin/python3


"""Test suite configuration"""


import os

import pytest

from click.testing import CliRunner


__all__ = ["get_static"]


def get_static() -> str:
    """Define the absolute path of the static folder.

    Returns:
        The absolute path of the static folder.
    """
    _staticpath = os.path.abspath(os.path.dirname(__file__))
    _staticpath = os.path.join(_staticpath, "static")
    return _staticpath


@pytest.fixture(scope="session")
def static() -> str:
    """Forward the absolute path of the static folder.

    Returns:
        The absolute path of the static folder.
    """
    return get_static()


@pytest.fixture(scope="session")
def runner() -> CliRunner:
    """Instantiate a Click CLI Runner.

    Returns:
       A Click CLI Runner instance.
    """
    runner = CliRunner()
    return runner


@pytest.fixture(autouse=True)
def wrap_every_test(mocker):
    """Wrap every single test with action that must be occur before and after.

    Returns:
        Nothing
    """
    # Setup: fill with any logic you want
    mocker.patch("cx_Oracle.connect")

    yield

    # Teardown : fill with any logic you want
