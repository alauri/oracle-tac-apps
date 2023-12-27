#!/usr/bin/env python3


"""Mock classes to simulate cx_Oracle's objects and their behaviours."""


from typing import Tuple


class MockResponse:
    def fetchone(self, *args, **kwargs):
        pass


class MockCursor:
    def execute(self, *args, **kwargs):
        return MockResponse()

    def close(self, *args, **kwargs):
        return


class MockOracle:
    def __init__(self, *args, **kwargs):
        pass

    def cursor(self, *args, **kwargs):
        return MockCursor()

    def commit(self, *args, **kwargs):
        return MockCursor()

    def ping(self, *args, **kwargs):
        return None


class MockUDPServer:
    def __init__(self, *args, **kwargs) -> None:
        self.curr = 0
        self.data = [
            "(2021,'Abu Dhabi','NaT|1|Car 1|Driver 1')",
            "0 days 00:01:29.103000|2|Car 1|Driver 1')",
            "0 days 00:01:28.827000|3|Car 1|Driver 1')",
            "0 days 00:01:29.026000|4|Car 1|Driver 1')",
            "0 days 00:01:28.718000|5|Car 1|Driver 1')",
            "EOF",
        ]

    def sendto(self, *args, **kwargs) -> None:
        pass

    def recvfrom(self, *args, **kwargs) -> Tuple[bytes]:
        res = self.data[self.curr]
        self.curr += 1
        return (res.encode(),)
