#!/bin/bash
cd /Users/luoxin/persons/go/lazygophers/utils/validator
go test -run=^$ -bench="^BenchmarkValidateField" -benchmem -benchtime=500ms . 2>&1 | tee VALIDATEFIELD_BENCH_RESULTS.txt
