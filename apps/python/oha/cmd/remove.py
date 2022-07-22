#!/usr/bin/python3


"""Command ``remove`` is used to delete one or more records from the db.

It can repeat the same operation in loop or a defined numbers of times. It can
be possible to define a delay between one operation and the next one and also
after how many operations commit the changes.
"""


import datetime
import time

import click

import cx_Oracle


@click.command()
@click.option('--loop/--no-loop',
              type=bool,
              default=False,
              help="repeat the same operation forever")
@click.option('--iters',
              type=int,
              default=1,
              help='repeat the same operation a given number of times')
@click.option('--delay',
              type=float,
              default=5,
              help='time to wait before the next iteration')
@click.option('--commit-every',
              type=int,
              default=1,
              help='after how many operations perform a commit')
@click.option('--older',
              type=int,
              default=10,
              help='days threshold to start removing data')
@click.pass_context
def remove(ctx,
           loop: bool,
           iters: int,
           delay: float,
           commit_every: int,
           older: int) -> None:
    """Delete records from the table"""
 
    # Define the query day, the day where to start from below query
    tstm = datetime.datetime(2022, 9, 9, 0, 0)

    # Define query parameters
    table = ctx.obj.conf["database"]["tablejson"]

    iters = 0 if loop else iters
    step = 1
    try:
        while loop or step <= iters:
            # Prepare query with updated conditions
            date = (tstm - datetime.timedelta(days=older))
            date = date.strftime("%Y-%m-%d %H:%M:%S")
            cond = f"timestamp <= to_date('{date}','yyyy-mm-dd hh24:mi:ss')"
            query = f"DELETE FROM {table} WHERE {cond}"

            # Execute query
            ctx.obj.cur.execute(query)
            click.echo(f"[{step}/{iters}] - {query}")

            # Commit changes
            if step % commit_every == 0:
                ctx.obj.conn.commit()
                click.echo(f"[{step}/{iters}] - COMMIT")

            step += 1
            time.sleep(delay)

        # Check the last commit
        if iters % commit_every != 0:
            ctx.obj.conn.commit()
            click.echo(f"[{iters}/{iters}] - COMMIT")
    except cx_Oracle.DatabaseError as err:
        click.echo(err)
        ctx.exit(1)
    except KeyboardInterrupt as _:
        click.echo("Error - Interrupted by the user")
        ctx.exit(1)
