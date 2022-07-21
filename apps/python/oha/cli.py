#!/usr/bin/python3


"""A Python application for the High Availability in Oracle"""


from typing import Dict

import os

import click

import toml

import cx_Oracle

from oha.cmd import (
    config,
    remove,
    injest,
    cleanup,
    reset
)


class OracleHA:

    filename = ""

    def __init__(self, wd: str, d: int):
        OracleHA.filename = os.path.join(os.path.abspath(wd), "config.toml")
        self.conf = OracleHA.read_toml()
        self.database(d)

    @staticmethod
    def read_toml() -> Dict:
        """Read TOML configuration file

        Returns:
            Nothing
        """
        return toml.load(OracleHA.filename)

    def database(self, connstr: int) -> None:
        """Initialize the Oracle database.

        Args:
            connstr: the index of the connection string to use.

        Returns:
            Nothing
        """
        self.conn = cx_Oracle.connect(
            user=self.conf["database"]["username"],
            password=self.conf["database"]["password"],
            dsn=self.conf["database"][f"dsn{connstr}"])
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
    try:
        ctx.obj = OracleHA(workdir, dsn)
    except cx_Oracle.DatabaseError as err:
        click.echo(err)
        ctx.exit(1)


# Register commands
cli.add_command(config)
cli.add_command(remove)
cli.add_command(injest)
cli.add_command(cleanup)
cli.add_command(reset)
