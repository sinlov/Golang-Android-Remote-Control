#!/usr/bin/env bash

rm -rf java/
rm -rf go/

mkdir -p java
mkdir -p go

cd java
flatc -j ../SimulationEvent.fbs
cd ..
cd go
flatc -g ../SimulationEvent.fbs