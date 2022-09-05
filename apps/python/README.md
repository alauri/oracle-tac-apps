# The Python application

The Python version of the application uses the driver **cx_Oracle**[^1] to
connect to the Oracle database.


# Installation and first execution

Use **Poetry**[^2] to install dependencies and enter the shell to access the
help menu:

    $ poetry install
    $ poetry shell
    $ oracle-tac-py --help


# Run tests

To run unit tests with PyTest[^3]:

    $ poetry shell
    $ pytest


# New driver release

In May 2022 a new version of the driver has been released. As explained in the
release note[^4], this is a library renamed, major version successor of the old
driver.

Support to the new version of the driver will be added soon.


[^1]: https://oracle.github.io/python-cx_Oracle/
[^2]: https://python-poetry.org/
[^3]: https://docs.pytest.org/
[^4]: https://python-oracledb.readthedocs.io/en/latest/user_guide/appendix_c.html#upgradecomparison
