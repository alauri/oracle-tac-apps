#!/usr/bin/python3


"""Alter tables with command ``reset`` and restore original configuration."""


import click

import cx_Oracle


@click.command()
@click.pass_context
def reset(ctx) -> None:
    """Reset database data to factory"""

    tableraw = ctx.obj.conf['database']['tableraw']
    tablejson = ctx.obj.conf['database']['tablejson']

    try:
        # Truncate table raw
        query = f"TRUNCATE TABLE {tableraw}"
        ctx.obj.cur.execute(query)
        click.echo(f"[+] - {query}")

        # Alter identity column for table raw
        query = f"ALTER TABLE {tableraw} " \
                f"MODIFY(ID GENERATED AS IDENTITY (START WITH 1))"
        ctx.obj.cur.execute(query)
        click.echo(f"[+] - {query}")

        # Truncate table json
        query = f"TRUNCATE TABLE {tablejson}"
        ctx.obj.cur.execute(query)
        click.echo(f"[+] - {query}")

        # Alter identity column for table json
        query = f"ALTER TABLE {tablejson} " \
                f"MODIFY(ID GENERATED AS IDENTITY (START WITH 1))"
        ctx.obj.cur.execute(query)
        click.echo(f"[+] - {query}")
    except cx_Oracle.DatabaseError as err:
        click.echo(err)
        ctx.exit(1)

    click.echo("[+] - All tables have been altered.")
