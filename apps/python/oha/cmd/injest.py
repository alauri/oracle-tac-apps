#!/usr/bin/python3


"""Command ``injest`` is used store new records within the db.

It can repeat the same operation in loop or a defined numbers of times. It can
be possible to define a delay between one operation and the next one and also
after how many operations commit the changes.
"""


import time
import click

import cx_Oracle


@click.command()
@click.option('--iters',
              type=int,
              default=1,
              help='repeat the same operation a given number of times')
@click.option('--delay',
              type=float,
              default=0.25,
              help='time to wait before the next iteration')
@click.option('--commit-every',
              type=int,
              default=1,
              help='after how many operations perform a commit')
@click.pass_context
def injest(ctx,
           iters: int,
           delay: float,
           commit_every: int) -> None:
    """Insert new records within the table"""

    # Set connection attriutes
    ctx.obj.conn.module = "oha.cmd.injest"

    # Define query parameters
    data = open(ctx.obj.conf["injest"]["dumpfile"]).readlines()
    try:
        for step, line in enumerate(data[:iters]):
            step += 1

            # Prepare the query
            ctx.obj.conn.action = "INS raw"
            query = f"INSERT INTO {ctx.obj.conf['database']['tableraw']}" \
                    f"(year,track,data) " \
                    f"VALUES({line.strip()})"

            # Execute query
            ctx.obj.cur.execute(query)
            click.echo(f"[{step}/{iters}] - {query}")

            # Commit changes
            if step % commit_every == 0:
                ctx.obj.conn.action = "COM raw"
                ctx.obj.conn.commit()
                click.echo(f"[{step}/{iters}] - COMMIT")

            # Wait before the next iteration
            time.sleep(delay)

        # Check the last commit
        if iters % commit_every != 0:
            ctx.obj.conn.action = "COM raw"
            ctx.obj.conn.commit()
            click.echo(f"[{iters}/{iters}] - COMMIT")
    except cx_Oracle.DatabaseError as err:
        click.echo(err)
        ctx.exit(1)
    except KeyboardInterrupt as _:
        click.echo("Error - Interrupted by the user")
        ctx.exit(1)
    finally:
        # Check for the last commit
        if iters % commit_every != 0:
            ctx.obj.conn.action = "COM raw"
            ctx.obj.conn.commit()
            click.echo(f"[{iters}/{iters}] - COMMIT")
