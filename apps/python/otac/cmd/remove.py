#!/usr/bin/python3


"""Command ``remove`` is used to delete one or more records from the db.

It can repeat the same operation in loop or a defined numbers of times. It can
be possible to define a delay between one operation and the next one and also
after how many operations commit the changes.
"""


import time

import click

import cx_Oracle


@click.command()
@click.option(
    "--delay", type=float, default=5, help="time to wait before the next iteration"
)
@click.option(
    "--commit-every",
    type=int,
    default=1,
    help="after how many operations perform a commit",
)
@click.pass_context
def remove(ctx, delay: float, commit_every: int) -> None:
    """Delete records from the table"""

    step = 1
    try:
        while True:
            query = "DELETE FROM data WHERE LapTime='NaT'"

            # Execute query
            ctx.obj.cur.execute(query)
            click.echo(f"[{step}/inf] - {query}")

            # Commit changes
            if step % commit_every == 0:
                ctx.obj.conn.commit()
                click.echo(f"[{step}/inf] - COMMIT")

            step += 1
            time.sleep(delay)
    except cx_Oracle.DatabaseError as err:
        click.echo(err)
        ctx.exit(1)
    except KeyboardInterrupt as _:
        click.echo("Error - Interrupted by the user")
        ctx.exit(1)
    finally:
        ctx.obj.conn.commit()
        click.echo(f"[{step}/inf] - COMMIT")
