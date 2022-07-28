# The application

The application is a simple command line interface written in the following
programming languages:

- :end: Python   
- :on: Golang   
- :soon: Java     

Its goal is to provide the same behaviour across all the examples and so all the
CLIs will have same commands, flags and outcomes.

All the applications use the driver *cx_Oracle* provided by Oracle to interact
with the database and all of them will initialize a single connection as long as
the applicaton lives.

## The CLI commands

The full list of commands can be retrieved by accessing the help menu of the
CLI:

    oracle-ha --help

which will return the following:

- **cleanup**: Update records within the database;
- **config**: Update sections of the TOML configuration file;
- **injest**: Insert new records within the table;
- **remove**: Delete records from the table;
- **reset**: Reset database data to factory.
