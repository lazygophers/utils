#!/bin/bash
cd xtime
go test -bench=BenchmarkEndOfDay -benchmem -count=5 -benchtime=3s
