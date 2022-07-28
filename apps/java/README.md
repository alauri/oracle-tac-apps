# The Java application

The Python version of the application using the driver **cx_Oracle**.


# Installation and first execution

Use **Maven**[^1] to install dependencies and use the command *java* to access
the help menu:

    $ cd oha
    $ mvn package
    $ java -cp "<abs-picocli-jar-path>:<target-jar-path>" com.oha.app.OracleHA --help


# Run tests

To run the tests suite use:

    $ mvn test


[^1]: https://maven.apache.org
