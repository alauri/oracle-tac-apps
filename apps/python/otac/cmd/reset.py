#!/usr/bin/python3


"""Alter tables with command ``reset`` and restore original configuration."""


import click

import cx_Oracle


@click.command()
@click.pass_context
def reset(ctx) -> None:
    """Reset database data to factory"""

    try:
        query = "TRUNCATE TABLE telemetry"
        ctx.obj.cur.execute(query)
        click.echo(f"[+] - {query}")

        query = "ALTER TABLE telemetry MODIFY(ID GENERATED AS IDENTITY (START WITH 1))"
        ctx.obj.cur.execute(query)
        click.echo(f"[+] - {query}")

        query = "TRUNCATE TABLE data"
        ctx.obj.cur.execute(query)
        click.echo(f"[+] - {query}")

        query = "ALTER TABLE data MODIFY(ID GENERATED AS IDENTITY (START WITH 1))"
        ctx.obj.cur.execute(query)
        click.echo(f"[+] - {query}")
    except cx_Oracle.DatabaseError as err:
        click.echo(err)
        ctx.exit(1)

    click.echo("[+] - All tables have been altered.")
