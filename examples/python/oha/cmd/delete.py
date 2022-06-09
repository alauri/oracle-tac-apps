#!/usr/bin/python3


"""Command ``delete`` is used to delete one or more records from the db.

It can repeat the same operation in loop or a defined numbers of times. It can
be possible to define a delay between one operation and the next one and also
after how many operations commit the changes.
"""


import time
import click


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
@click.pass_obj
def delete(obj,
           loop: bool,
           iters: int,
           delay: float,
           commit_every: int) -> None:
    """Delete records from the table."""

    # Define query parameters
    table = obj.conf["database"]["table"]
    conds = {"id": 0}

    iters = 0 if loop else iters
    step = 1
    try:
        while loop or step <= iters:
            # Prepare query with updated conditions
            pairs = [f"{cond}={val + step}" for cond, val in conds.items()]
            query = f"DELETE FROM {table} WHERE {', '.join(pairs)}"

            # TODO: replace with call to cx_Oracle driver
            click.echo(f"[{step}/{iters}] - {query}")

            # Commit changes
            if step % commit_every == 0:
                # TODO: insert a database commit operation here
                click.echo(f"[{step}/{iters}] - COMMIT")

            step += 1
            time.sleep(delay)
    except KeyboardInterrupt as _:
        click.echo("Error - Interrupted by the user")

    # Check the last commit
    if iters % commit_every != 0:
        # TODO: insert a database commit operation here
        click.echo(f"[{iters}/{iters}] - COMMIT")
