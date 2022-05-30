#!/usr/bin/python3

"""Test suite configuration."""


import pytest


from click.testing import CliRunner


@pytest.fixture
def runner() -> CliRunner:
    """Instantiate a Click CLI Runner.

    Returns:
       A Click CLI Runner instance.
    """
    runner = CliRunner()
    return runner


@pytest.fixture(autouse=True)
def wrap_every_test():
    """Wrap every single test with action that must be occur before and after.

    Returns:
        Nothing
    """
    # Setup: fill with any logic you want

    yield

    # Teardown : fill with any logic you want
