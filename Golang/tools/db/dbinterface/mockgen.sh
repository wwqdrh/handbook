#!/usr/bin/env bash
mockgen -destination mocks_test.go -package dbinterface wwqdrh/handbook/tools/db/dbinterface DB,Transaction
