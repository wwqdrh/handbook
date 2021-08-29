#!/usr/bin/env bash
go-fuzz-build wwqdrh/handbook/tools/profile/fuzz
go-fuzz -bin=./fuzz-fuzz.zip -workdir=output
