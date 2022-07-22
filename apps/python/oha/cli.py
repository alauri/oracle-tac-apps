#!/usr/bin/python3


"""A Python application for the High Availability in Oracle"""


from typing import Dict

import pprint
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

    def __init__(self, tomlfile: Dict, workdir: str, dsn: int):
        self.workdir = workdir
        self.conf = tomlfile
        self.database(dsn)

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


@click.group(invoke_without_command=True)
@click.option('--info/--no-info',
              type=bool,
              default=False,
              help="Information about the current configuration")
@click.option("-w", "--workdir",
              type=str,
              default=os.path.join(os.path.dirname(__file__), "../.."),
              help="The absolute path of the working folder")
@click.option("-d", "--dsn",
              type=click.IntRange(min=1, max=5),
              default=1,
              help="The connection string to use")
@click.pass_context
def cli(ctx, info: bool, workdir: str, dsn: int) -> None:
    """Oracle High Availability CLI in Python"""

    # Read the configuration file
    tomlfile = toml.load(os.path.join(os.path.abspath(workdir), "config.toml"))

    if info:
        pprint.pprint(tomlfile, indent=4)
        ctx.exit(0)

    if ctx.invoked_subcommand is None:
        click.echo(cli.get_help(ctx))
        ctx.exit(0)

    # Initialize Click context with TOML configuration file
    try:
        ctx.obj = OracleHA(tomlfile, workdir, dsn)
    except cx_Oracle.DatabaseError as err:
        click.echo(err)
        ctx.exit(1)


# Register commands
cli.add_command(config)
cli.add_command(remove)
cli.add_command(injest)
cli.add_command(cleanup)
cli.add_command(reset)
