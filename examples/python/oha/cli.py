#!/usr/bin/python3


"""
"""


import click

from oha.commands import (
    config,
    delete,
    insert,
    update,
    reset
)


@click.group(invoke_without_command=True)
def cli() -> None:
    """
    """
    click.echo("The main command")


# Register commands later
cli.add_command(config)
cli.add_command(delete)
cli.add_command(insert)
cli.add_command(update)
cli.add_command(reset)
