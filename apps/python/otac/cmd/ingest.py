#!/usr/bin/python3


"""Command ``ingest`` is used store new records within the db.

It can repeat the same operation in loop or a defined numbers of times. It can
be possible to define a delay between one operation and the next one and also
after how many operations commit the changes.
"""

# from socket import AF_INET, SOCK_DGRAM, socket
import socket

import click
import time

import cx_Oracle


@click.command()
@click.option(
    "--port",
    type=int,
    default=1053,
    help="the UDP server port to connect to",
)
@click.option(
    "--delay", type=float, default=0.25, help="time to wait before the next iteration"
)
@click.option(
    "--commit-every",
    type=int,
    default=5,
    help="after how many operations perform a commit",
)
@click.pass_context
def ingest(ctx, delay: float, port: int, commit_every: int) -> None:
    """Insert new records within the table"""

    with open("/tmp/test.txt", "w") as outfile:
        outfile.write("Start\n")

        # Connect to UDP server
        soc = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)
        outfile.write(f"{type(soc)}\n")
        soc.sendto("Python".encode(), ("127.0.0.1", port))
        outfile.write("Sent\n")

        # Define query parameters
        counter = 1
        try:
            # Think about a EOF from the UDP server
            while True:
                # Get message from server
                buffer = soc.recvfrom(1024)
                msg = buffer[0].decode()
                outfile.write(f"{msg}\n")

                # Execute query
                query = f"INSERT INTO raw_tel(year,track,data) VALUES({msg.strip()})"
                ctx.obj.cur.execute(query)
                click.echo(f"[{counter}/inf] - {query}")

                # Check the last commit
                if counter % commit_every == 0:
                    ctx.obj.conn.commit()
                    click.echo(f"[{counter}/inf] - COMMIT")

                # No more data, stop processing
                if msg == "EOF":
                    click.echo(
                        f"[{counter}/inf] - No more data will be sent, exiting..."
                    )
                    break

                # Increase counter
                counter += 1
                time.sleep(delay)
        except cx_Oracle.DatabaseError as err:
            click.echo(err)
            ctx.exit(1)
        except KeyboardInterrupt as _:
            click.echo("Error - Interrupted by the user")
            ctx.exit(1)
        finally:
            # Perfom the last commit in case something is left
            # Check for the last commit
            ctx.obj.conn.commit()
            click.echo(f"[{counter}/inf] - COMMIT")
