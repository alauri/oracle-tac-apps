#!/usr/bin/python3


"""Command ``cleanup`` to change already existing records within the db.

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
def cleanup(ctx,
            loop: bool,
            iters: int,
            delay: float,
            commit_every: int) -> None:
    """Update records within the database"""

    # Get the first id to read from the table raw
    tail = ctx.obj.conf['cleanup']['tail']

    iters = 0 if loop else iters
    step = 1
    try:
        while loop or step <= iters:
            query = f"SELECT * " \
                    f"FROM {ctx.obj.conf['database']['tableraw']} " \
                    f"WHERE id={tail}"
            res = ctx.obj.cur.execute(query).fetchone()
            click.echo(f"[{step}/{iters}] - {query}")

            # Get and clean information
            if res is None:
                click.echo(f"[{step}/{iters}] - No row to clean up. Exit.")
                break

            _, year, track, data = res
            lt, ln, team, driver = data.split("|")
            lt = lt.replace("0 days ", "") if lt != "NaT" else lt

            # Prepare the query
            query = f"INSERT INTO {ctx.obj.conf['database']['tablejson']}" \
                    f"(year,track,laptime,lapnumber,team,driver) " \
                    f"VALUES({year},'{track}','{lt}',{ln},'{team}','{driver}')"

            # Execute query
            ctx.obj.cur.execute(query)
            click.echo(f"[{step}/{iters}] - {query}")

            # Commit changes
            if step % commit_every == 0:
                ctx.obj.conn.commit()
                click.echo(f"[{step}/{iters}] - COMMIT")

            step += 1
            tail += 1
            time.sleep(delay)
    except cx_Oracle.DatabaseError as err:
        click.echo(err)
        ctx.exit(1)
    except KeyboardInterrupt as _:
        click.echo("Error - Interrupted by the user")
        ctx.exit(1)
    finally:
        # Check for the last commit
        if iters % commit_every != 0:
            ctx.obj.conn.commit()
            click.echo(f"[{iters}/{iters}] - COMMIT")
