#!/usr/bin/python3


"""Command ``update`` to change already existing records within the db.

It can repeat the same operation in loop or a defined numbers of times. It can
be possible to define a delay between one operation and the next one and also
after how many operations commit the changes.
"""


from datetime import datetime
import time
import click
import json

import cx_Oracle

from oha.cmd import consts


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

    # Get database information
    raw_table = ctx.obj.conf["database"]["tableraw"]
    json_table = ctx.obj.conf["database"]["tablejson"]
    checkpoint = ctx.obj.conf["cleanup"]["checkpoint"]

    iters = 0 if loop else iters
    step = 1
    try:
        while loop or step <= iters:
            # TODO: handle the error when a checkpoint does not exist
            query = f"SELECT * " \
                    f"FROM {raw_table} " \
                    f"WHERE id={checkpoint}"
            res = ctx.obj.cur.execute(query)
            click.echo(f"[{step}/{iters}] - {query}")

            # Get and clean information
            _, timestamp, sensorid, data = res.fetchone()
            timestamp = f"to_date('{datetime.fromtimestamp(timestamp)}'," \
                        "'yyyy-mm-dd hh24:mi:ss')"
            sensorid = consts.SENSORS[sensorid]
            data = json.dumps(dict([tuple(p.split("=")) for p in data.split("|")]))

            # Prepare the query
            query = f"INSERT INTO {json_table}(timestamp,sensorid,data) " \
                    f"VALUES({timestamp},'{sensorid}','{data}')"

            # Execute query
            try:
                ctx.obj.cur.execute(query)
                click.echo(f"[{step}/{iters}] - {query}")
            except cx_Oracle.IntegrityError as err:
                click.echo(err)
                click.exit(1)

            # Commit changes
            if step % commit_every == 0:
                ctx.obj.conn.commit()
                click.echo(f"[{step}/{iters}] - COMMIT")

            step += 1
            checkpoint += 1
            time.sleep(delay)

        # TODO: save checkpoint

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
