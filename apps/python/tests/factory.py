#!/usr/bin/env python3


"""Mock classes to simulate cx_Oracle's objects and their behaviours."""


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
