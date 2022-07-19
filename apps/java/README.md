# Oracle HA in Python

This is the Java version of the application using the driver **cx_Oracle**[^1]
to get advantage of the Oracle High Availability.


# Installation and first execution

Use **Maven**[^2] to install all the dependencies of the project:

    $ cd oha
    $ mvn package

This will compile the project with a pre-defined structure by Maven itself. To
execute the CLI, type the following command:

    $ java -cp "<abs-picocli-jar-path>:<target-jar-path>" com.oha.app.OracleHA --help


# Run tests

To run the tests suite use:

    $ mvn test


[^1]: https://oracle.github.io/python-cx_Oracle/
[^2]: https://maven.apache.org
