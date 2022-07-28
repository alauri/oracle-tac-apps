# The Golang application

The Golang version of the application using the driver **cx_Oracle**[^1].


# Installation and first execution

Install Go dependencies and access the help menu:

    $ cd oracle-ha
    $ go install
    $ oracle-ha --help


# Run tests

To run the tests suite type:

    $ cd oracle-ha
    $ go test ./...


[^1]: https://blogs.oracle.com/developers/post/how-to-connect-a-go-program-to-oracle-database-using-godror
