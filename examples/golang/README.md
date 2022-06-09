# Oracle HA in Golang

This is the Golang version of the application using the driver **cx_Oracle**[^1]
to get advantage of the Oracle High Availability.


# Installation and first execution

To install the CLI move under the folder **oracle-ha** and type:

    $ go install

This will compile and install the CLI in your Golang's bin folder and make it
available at system level:

    $ oracle-ha --help


# Run tests

To run the tests suite type:

    $ go test ./...


[^1]: https://blogs.oracle.com/developers/post/how-to-connect-a-go-program-to-oracle-database-using-godror
