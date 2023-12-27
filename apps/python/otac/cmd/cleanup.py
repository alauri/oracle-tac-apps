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
@click.option(
    "--delay", type=float, default=0.25, help="time to wait before the next iteration"
)
@click.option(
    "--commit-every",
    type=int,
    default=1,
    help="after how many operations perform a commit",
)
@click.pass_context
def cleanup(ctx, delay: float, commit_every: int) -> None:
    """Update records within the database"""

    step = 1
    curr = 1
    try:
        while True:
            query = f"SELECT * FROM telemetry WHERE id={curr}"
            res = ctx.obj.cur.execute(query).fetchone()

            # Check empty query result
            if res is None:
                click.echo(f"[{curr}/inf] - No row to clean up. Exit.")
                break
            else:
                click.echo(f"[{curr}/inf] - {query}")

            _, year, track, data = res
            lt, ln, team, driver = data.split("|")
            lt = lt.replace("0 days ", "") if lt != "NaT" else lt

            query = (
                f"INSERT INTO data"
                f"(year,track,laptime,lapnumber,team,driver) "
                f"VALUES({year},'{track}','{lt}',{ln},'{team}','{driver}')"
            )

            # Execute query
            ctx.obj.cur.execute(query)
            click.echo(f"[{curr}/inf] - {query}")

            # Commit changes
            if step % commit_every == 0:
                ctx.obj.conn.commit()
                click.echo(f"[{curr}/inf] - COMMIT")

            step += 1
            curr += 1
            time.sleep(delay)
    except cx_Oracle.DatabaseError as err:
        click.echo(err)
        ctx.exit(1)
    except KeyboardInterrupt as _:
        click.echo("Error - Interrupted by the user")
        ctx.exit(1)
    finally:
        # Check for the last commit
        ctx.obj.conn.commit()
        click.echo(f"[{curr}/inf] - COMMIT")
