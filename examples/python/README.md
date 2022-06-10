# Oracle HA in Python

This is the Python version of the application using the driver **cx_Oracle**[^1]
to get advantage of the Oracle High Availability.


# Installation and first execution

Use **Poetry**[^2] to install all the dependencies of the project:

    $ poetry install

This will create a new virtualenv and install the dependencies and the command
*oracle-ha* you can use to invoke the CLI:

    $ poetry shell
    $ oracle-ha --help


# Run tests

To run the tests suite use PyTest[^3]:

    $ poetry shell
    $ pytest


[^1]: https://oracle.github.io/python-cx_Oracle/
[^2]: https://python-poetry.org/
[^3]: https://docs.pytest.org/
