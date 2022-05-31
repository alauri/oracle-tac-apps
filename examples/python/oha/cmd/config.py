#!/usr/bin/python3


import click


@click.command()
def config() -> None:
    """Configure the application.
    """
    click.echo("config called")
