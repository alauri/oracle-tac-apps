#!/usr/bin/python3


import click


@click.command()
def update() -> None:
    """Update records into the table.
    """
    click.echo("update called")
