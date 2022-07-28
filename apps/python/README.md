# The Python application

The Python version of the application using the driver **cx_Oracle**[^1].


# Installation and first execution

Use **Poetry**[^2] to install dependencies and enter the shell to access the
help menu:

    $ poetry install
    $ poetry shell
    $ oracle-ha --help


# Run tests

To run unit tests with PyTest[^3]:

    $ poetry shell
    $ pytest


# Use the new driver version

In May 2022 a new version of the driver has been released. As explained in the
release note[^4], this is a library renamed, major version successor of the old
driver.


[^1]: https://oracle.github.io/python-cx_Oracle/
[^2]: https://python-poetry.org/
[^3]: https://docs.pytest.org/
[^4]: https://python-oracledb.readthedocs.io/en/latest/user_guide/appendix_c.html#upgradecomparison
