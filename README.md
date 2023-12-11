# Oracle Transparent Application Continuity (oracle-tac)

End users get frustrated with unreachable or inconsistent data and this
can also be very disappointing for developers and DBAs who must ensure
relaible services always ready to return information; the **Oracle High
Availability (HA)** helps to garantiee a service is always up-and-running.

As a general description:

> Availability is the degree to which an application and database service is
> available[^1]

Oracle provides a wide range of features to cover all the possible use cases in
order to achive the HA; one of the most important is the **Transparent
Application Continuity (TAC)**. TAC can recover database session after
recoverable outages, with no need for DBA to have any knowledge of the
application and with no need for developers to make application code changes.

The aim of this project is to demonstrate how TAC can be accomplished with
different programming languages.

All the code examples you'll find in this repo use the **cx-Oracle** driver to
connect to Oracle database. All the guidelines and how-to are provided for all
of them.


## The project structure

The project is split into two parts:

 - **apps**: folder with example applications in different programming
   languages;
 - **orchestrate**: the IT automantion tool to deploy all or some of the
   applications on the remotes you want;
 - **producer**: the data generator.


[^1]: https://docs.oracle.com/en/database/oracle/oracle-database/19/haovw/overview-of-ha.html
