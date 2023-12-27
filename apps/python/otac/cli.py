#!/usr/bin/python3


"""The Python application"""


from typing import Dict
from os.path import join, dirname

import pprint

import click

from toml import load

import cx_Oracle

from otac.cmd import remove, ingest, cleanup, reset


# Read the configuration file
TOMLFILE = load(join(dirname(__file__), "../../config.toml"))


class OracleTAC:
    def __init__(self, tomlfile: Dict, dsn: str):
        self.conf = tomlfile
        self.database(dsn)

    def database(self, connstr: str) -> None:
        """Initialize the Oracle database.

        Args:
            connstr: the index of the connection string to use.

        Returns:
            Nothing
        """
        # Initialize connection and set the tracing metadata
        self.conn = cx_Oracle.connect(
            user=self.conf["database"]["username"],
            password=self.conf["database"]["password"],
            dsn=self.conf["dsn"][connstr],
        )

        # Get the cursor
        self.cur = self.conn.cursor()


@click.group(invoke_without_command=True)
@click.option(
    "--config/--no-config",
    type=bool,
    default=False,
    help="show database configuration",
)
@click.option(
    "--ping/--no-ping", type=bool, default=False, help="test database connection"
)
@click.option(
    "-d",
    "--dsn",
    type=click.Choice(TOMLFILE["dsn"]),
    default="localhost",
    help="the database connection string to use",
)
@click.pass_context
def cli(ctx, config: bool, ping: bool, dsn: str) -> None:
    """Oracle High Availability CLI in Python"""

    # Define teardown callbacks
    ctx.call_on_close(_on_close)

    # Check the config flag. If it's True, show the current configuration and
    # exit
    if config:
        pprint.pprint(TOMLFILE, indent=4)
        ctx.exit(0)

    if ctx.invoked_subcommand is None and not ping:
        click.echo(cli.get_help(ctx))
        ctx.exit(0)

    # Initialize database connection and retrieve context info
    ctx.obj = OracleTAC(TOMLFILE, dsn)
    click.echo(f"[+] - {_get_db_info()}")

    # Check the database is reachable
    if ping:
        _ = ctx.obj.conn.ping()
        click.echo("[+] - Database reachable")
        ctx.exit(0)


@click.pass_context
def _get_db_info(ctx) -> str:
    """Query the database to retrieve info about the context.

    Args:
        ctx: the Click context.

    Returns:
        DB context info.
    """
    query = (
        "SELECT "
        "SYS_CONTEXT('USERENV', 'DB_UNIQUE_NAME') AS DB_UNIQUE_NAME, "
        "SYS_CONTEXT('USERENV', 'SERVER_HOST') AS HOST "
        "FROM DUAL"
    )
    db_ctx = ctx.obj.cur.execute(query).fetchone()
    return db_ctx


@click.pass_context
def _on_close(ctx) -> None:
    """Teardown connections and configuration.

    Returns:
        Nothing
    """
    # Check db cursor
    if ctx.obj is None:
        return

    try:
        # Print DB context info
        click.echo(f"[+] - {_get_db_info()}")

        # Close the cursor
        ctx.obj.cur.close()
    except cx_Oracle.DatabaseError as err:
        click.echo(err)
        ctx.exit(1)


# Register commands
cli.add_command(remove)
cli.add_command(ingest)
cli.add_command(cleanup)
cli.add_command(reset)
