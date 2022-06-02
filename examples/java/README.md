# Oracle HA in Python

This is the Java version of the application using the driver **cx_Oracle**[^1]
to get advantage of the Oracle High Availability.


# Installation and first execution

Use **Maven**[^2] to install all the dependencies of the project:

    $ cd oha
    $ mvn compile

This will compile the project with a pre-defined structure by Maven itself.


# Run tests

To run the tests suite use:

    $ mvn test


[^1]: https://oracle.github.io/python-cx_Oracle/
[^2]: https://maven.apache.org
