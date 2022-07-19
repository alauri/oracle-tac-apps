#!/usr/bin/python3


"""Command ``update`` to change already existing records within the db.

It can repeat the same operation in loop or a defined numbers of times. It can
be possible to define a delay between one operation and the next one and also
after how many operations commit the changes.
"""


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
              default=0.25,
              help='time to wait before the next iteration')
@click.option('--commit-every',
              type=int,
              default=1,
              help='after how many operations perform a commit')
@click.pass_context
def update(ctx,
           loop: bool,
           iters: int,
           delay: float,
           commit_every: int) -> None:
    """Update records within the database"""

    # Define query parameters
    table = ctx.obj.conf["database"]["table"]
    args = {"DEPARTMENT_NAME": 0}
    conds = {"DEPARTMENT_ID": 1}

    iters = 0 if loop else iters
    step = 1
    try:
        while loop or step <= iters:
            # Prepare query with updated conditions
            sets = [f"{arg}='pippo{val + step}'" for arg, val in args.items()]
            wheres = [f"{arg}={val + step}" for arg, val in conds.items()]
            query = (f"UPDATE {table} "
                     f"SET {', '.join(sets)} "
                     f"WHERE {', '.join(wheres)}")

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
