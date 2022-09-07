#!/usr/bin/python3


"""The Python application"""


from typing import Dict

import pprint
import os

import click

import toml

import cx_Oracle

from oha.cmd import (
    remove,
    ingest,
    cleanup,
    reset
)


class OracleTAC:

    def __init__(self, tomlfile: Dict, workdir: str, dsn: str):
        self.workdir = workdir
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
            dsn=self.conf["dsn"][connstr])

        # Instrumentation: Set up the client identifier (ACCHK_REPORT)
        self.conn.client_identifier = "oracle.tac.py"

        # Get the cursor
        self.cur = self.conn.cursor()


@click.group(invoke_without_command=True)
@click.option("-w", "--workdir",
              type=str,
              default=os.path.join(os.path.dirname(__file__), "../.."),
              help="The absolute path of the working folder")
@click.option("--config/--no-config",
              type=bool,
              default=False,
              help="Show the current configuration")
@click.option("--ping/--no-ping",
              type=bool,
              default=False,
              help="Check database connection")
@click.option("-d", "--dsn",
              type=str,
              default="localhost",
              help="The connection string to use")
@click.pass_context
def cli(ctx, workdir: str, config: bool, ping: bool, dsn: int) -> None:
    """Oracle High Availability CLI in Python"""

    # Define teardown callbacks
    ctx.call_on_close(_on_close)

    # Read the configuration file
    tomlfile = toml.load(os.path.join(os.path.abspath(workdir), "config.toml"))

    # Check the given DSN
    if dsn not in tomlfile["dsn"]:
        raise click.UsageError("Invalid value for '-d/--dsn'")

    # Check the config flag. If it's True, show the current configuration and
    # exit
    if config:
        pprint.pprint(tomlfile, indent=4)
        ctx.exit(0)

    if ctx.invoked_subcommand is None and not ping:
        click.echo(cli.get_help(ctx))
        ctx.exit(0)

    # Initialize database connection and retrieve context info
    ctx.obj = OracleTAC(tomlfile, workdir, dsn)
    click.echo(f"[+] - {_get_db_info()}")

    # Instrumentation: Set up current module (ACCHK_REPORT)
    ctx.obj.conn.module = "oha.cli"

    # Check the database is reachable
    if ping:
        _ = ctx.obj.conn.ping()
        click.echo("[+] - Database reachable")
        ctx.exit(0)

    # Initialize Click context with TOML configuration file
    try:
        # Retrieve the ID of the first row from the raw table
        query = f"SELECT id " \
                f"FROM {ctx.obj.conf['database']['tableraw']} " \
                f"WHERE rownum=1"
        headraw = ctx.obj.cur.execute(query).fetchone()
        headraw = 0 if headraw is None else int(headraw[0])

        # Retrieve how many rows are in the clean table
        query = f"SELECT COUNT(*) " \
                f"FROM {ctx.obj.conf['database']['tablejson']}"
        tail = ctx.obj.cur.execute(query).fetchone()[0]
        tail += headraw

        ctx.obj.conf["cleanup"]["tail"] = tail
    except cx_Oracle.DatabaseError as err:
        click.echo(err)
        ctx.exit(1)


@click.pass_context
def _get_db_info(ctx) -> str:
    """Query the database to retrieve info about the context.

    Args:
        ctx: the Click context.

    Returns:
        DB context info.
    """
    # Instrumentation: Set up module action (ACCHK_REPORT)
    ctx.obj.conn.action = "SELECT.system.context"

    query = "SELECT " \
            "SYS_CONTEXT('USERENV', 'DB_UNIQUE_NAME') AS DB_UNIQUE_NAME, " \
            "SYS_CONTEXT('USERENV', 'SERVER_HOST') AS HOST " \
            "FROM DUAL"
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
