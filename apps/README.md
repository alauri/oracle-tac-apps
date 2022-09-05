# The application

The application is a simple Command Line Interface written in the following
programming languages:

- :end: Python
- :on: Golang
- :soon: Java

We provide different implementations for this application to demonstrate that
the *Transparent Application Continuity* in Oracle databases can be achieved in
several programming languages.


## The CLI commands

The full list of commands can be retrieved by accessing the help menu of the
CLI:

    oracle-tac-<py|go|java> --help

which will return the following:

- **cleanup**: Read and store cleaned data into a different table;
- **ingest**: Insert new records within the table;
- **remove**: Delete records from the table;
- **reset**: Reset database data to factory.
