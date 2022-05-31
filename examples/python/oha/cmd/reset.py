#!/usr/bin/python3


import click


@click.command()
def reset() -> None:
    """Reset database data to factory.
    """
    click.echo("reset called")
