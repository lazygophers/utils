#!/bin/bash
cd /Users/luoxin/persons/go/lazygophers/utils/xtime
go test -bench=BenchmarkEndOfMonth_Global -benchmem -benchtime=5s -count=3 . 2>&1
