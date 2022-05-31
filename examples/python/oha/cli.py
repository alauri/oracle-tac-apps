#!/usr/bin/python3


"""
"""


import click

from oha.cmd import (
    config,
    delete,
    insert,
    update,
    reset
)


@click.group(invoke_without_command=True)
def cli() -> None:
    """Oracle High Availability CLI in Python
    """
    click.echo("root called")


# Register commands
cli.add_command(config)
cli.add_command(delete)
cli.add_command(insert)
cli.add_command(update)
cli.add_command(reset)
