#!/usr/bin/python3


"""A Python application for the High Availability in Oracle"""


from typing import Dict

import os

import click

import toml

import cx_Oracle

from oha.cmd import (
    config,
    delete,
    insert,
    update,
    reset
)


class OracleHA:

    filename = ""

    def __init__(self, wd: str, d: int):
        OracleHA.filename = os.path.join(os.path.abspath(wd), "config.toml")
        self.conf = OracleHA.read_toml()
        self.driver(d)

    @staticmethod
    def read_toml() -> Dict:
        """Read TOML configuration file

        Returns:
            Nothing
        """
        return toml.load(OracleHA.filename)

    def driver(self, connstr: int) -> None:
        """Initialize the Oracle driver.

        Args:
            connstr: the index of the connection string to use.

        Returns:
            Nothing
        """
        self.conn = cx_Oracle.connect(user=self.conf["driver"]["username"],
                                      password=self.conf["driver"]["password"],
                                      dsn=self.conf["driver"][f"dsn{connstr}"])
        self.cur = self.conn.cursor()


@click.group()
@click.option("-w", "--workdir",
              type=str,
              default=os.path.join(os.path.dirname(__file__), "../.."),
              help="the absolute path of the configuration folder")
@click.option("-d", "--dsn",
              type=click.IntRange(min=1, max=5),
              default=1,
              help="the connection string to use")
@click.pass_context
def cli(ctx, workdir: str, dsn: int) -> None:
    """Oracle High Availability CLI in Python"""

    # Initialize Click context with TOML configuration file
    ctx.obj = OracleHA(workdir, dsn)


# Register commands
cli.add_command(config)
cli.add_command(delete)
cli.add_command(insert)
cli.add_command(update)
cli.add_command(reset)
