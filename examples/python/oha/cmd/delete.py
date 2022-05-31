#!/usr/bin/python3


import click


@click.command()
def delete() -> None:
    """Delete records from the table.
    """
    click.echo("delete called")
