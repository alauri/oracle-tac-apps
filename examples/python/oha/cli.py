#!/usr/bin/python3


"""A Python application for the High Availability in Oracle"""


from typing import Dict

import os

import click

import toml

from oha.cmd import (
    config,
    delete,
    insert,
    update,
    reset
)


class OracleHA:

    filename = ""

    def __init__(self, wd: str):
        OracleHA.filename = os.path.join(os.path.abspath(wd), "config.toml")
        self.conf = OracleHA.read_toml()

    @staticmethod
    def read_toml() -> Dict:
        """Read TOML configuration file

        Returns:
            Nothing
        """
        return toml.load(OracleHA.filename)


@click.group()
@click.option("-w", "--workdir",
              type=str,
              default=os.path.join(os.path.dirname(__file__), "../.."),
              help="the absolute path of the configuration folder")
@click.pass_context
def cli(ctx, workdir: str) -> None:
    """Oracle High Availability CLI in Python"""

    # Initialize Click context with TOML configuration file
    ctx.obj = OracleHA(workdir)


# Register commands
cli.add_command(config)
cli.add_command(delete)
cli.add_command(insert)
cli.add_command(update)
cli.add_command(reset)
