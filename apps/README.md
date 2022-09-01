# The application

The application is a simple command line interface written in the following
programming languages:

- :end: Python   
- :on: Golang   
- :soon: Java     

The goal is to demonstrate that by using the driver provided by Oracle is
possible to ensure the same behaviour in different programming languages.


## The CLI commands

The full list of commands can be retrieved by accessing the help menu of the
CLI:

    oracle-tac-<py|go|java> --help

which will return the following:

- **cleanup**: Read and store cleaned data into a different table;
- **ingest**: Insert new records within the table;
- **remove**: Delete records from the table;
- **reset**: Reset database data to factory.
