#!/usr/bin/python3


"""Command ``config`` helps managing the TOML configuration file by listing or
updating the current configuration.

It supports sub-commands to manage all the TOML sections with a Click's option
for each value of the section.
"""


import pprint
import click
import toml
import os


@click.group(invoke_without_command=True)
@click.option("--info/--no-info",
              type=bool,
              default=False,
              help="print current stored configurations")
@click.pass_context
def config(ctx, info: bool) -> None:
    """Update sections of the TOML configuration file"""

    # Print current configuration info and exit
    if info:
        pprint.pprint(ctx.obj.conf, indent=4)
        ctx.exit(0)

    if ctx.invoked_subcommand is None:
        click.echo(config.get_help(ctx))


@click.command()
@click.option("--username",
              type=str,
              default=None,
              help="update the driver's username")
@click.option("--password",
              type=str,
              default=None,
              help="update the driver's password")
@click.pass_obj
def database(obj, **kwargs) -> None:
    """update section 'database'"""

    # Collect only those values different from None and update the TOML file
    updates = {(arg, val) for arg, val in kwargs.items() if val is not None}
    obj.conf["database"].update(updates)
    toml.dump(obj.conf, open(os.path.join(obj.workdir, "config.toml"), "w"))

    click.echo("[+] - Configuration updated")


config.add_command(database)
