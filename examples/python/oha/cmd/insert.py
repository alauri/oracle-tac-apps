#!/usr/bin/python3


import click


@click.command()
def insert() -> None:
    """Insert new records into the table.
    """
    click.echo("insert called")
