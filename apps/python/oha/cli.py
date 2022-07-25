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
@click.option("-w", "--workdir",
              type=str,
              default=os.path.join(os.path.dirname(__file__), "../.."),
              help="The absolute path of the working folder")
@click.option("-d", "--dsn",
              type=click.IntRange(min=1, max=5),
              default=1,
              help="The connection string to use")
@click.pass_context
def cli(ctx, workdir: str, dsn: int) -> None:
    """Oracle High Availability CLI in Python"""

    # Read the configuration file
    tomlfile = toml.load(os.path.join(os.path.abspath(workdir), "config.toml"))

    if ctx.invoked_subcommand is None:
        click.echo(cli.get_help(ctx))
        ctx.exit(0)

    # Initialize Click context with TOML configuration file
    try:
        ctx.obj = OracleHA(tomlfile, workdir, dsn)

        # Retrieve the ID of the first row from the raw table
        query = f"SELECT id " \
                f"FROM {ctx.obj.conf['database']['tableraw']} " \
                f"WHERE rownum=1"
        headraw = ctx.obj.cur.execute(query).fetchone()
        headraw = 0 if headraw is None else int(headraw[0])

        # Retrieve the ID of the first row from the json table
        query = f"SELECT COUNT(*) " \
                f"FROM {ctx.obj.conf['database']['tablejson']}"
        tail = ctx.obj.cur.execute(query).fetchone()[0]
        tail += headraw

        ctx.obj.conf["injest"]["head"] = headraw
        ctx.obj.conf["cleanup"]["tail"] = tail
    except cx_Oracle.DatabaseError as err:
        click.echo(err)
        ctx.exit(1)


# Register commands
cli.add_command(config)
cli.add_command(remove)
cli.add_command(injest)
cli.add_command(cleanup)
cli.add_command(reset)
