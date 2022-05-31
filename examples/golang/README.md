# Oracle HA in Golang

This is the Golang version of the application using the driver **cx_Oracle**[^1]
to get advantage of the Oracle High Availability.


# Installation and first execution

Move under the folder **oha** and run:

    $ go build -o oracle-ha

This will install all the dependencies and create the script *oracle-ha* you
can use to invoke the CLI:

    $ ./oracle-ha --help


# Run tests

To run the tests suite type:

    $ go test ./...


[^1]: https://blogs.oracle.com/developers/post/how-to-connect-a-go-program-to-oracle-database-using-godror
