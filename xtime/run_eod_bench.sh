#!/bin/bash
cd /Users/luoxin/persons/go/lazygophers/utils/xtime
go test -bench=BenchmarkEndOfDay -benchmem -count=3 -benchtime=5s
