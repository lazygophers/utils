# Go Test File Reorganization Plan

## Overview
This document outlines the comprehensive plan to reorganize test files in the Go project to follow the standard naming convention where each source file `pkg.go` should have a corresponding test file `pkg_test.go`, with unit tests and benchmarks consolidated into the same file.

## Current Issues Identified

### 1. Files with Multiple Test Files That Need Consolidation

#### `/fake/` directory
**Source files:** `fake.go`, `address.go`, `names.go`
**Current test files:**
- `address_test.go` ✅ (correctly named)
- `basic_test.go` → should be merged into `fake_test.go`
- `benchmark_test.go` → should be merged into `fake_test.go`
- `example_test.go` → should be merged into `fake_test.go`
- `extreme_bench_test.go` → should be merged into `fake_test.go`
- `extreme_example_test.go` → should be merged into `fake_test.go`
- `fake_test.go` ✅ (correctly named, consolidation target)
- `names_test.go` ✅ (correctly named)
- `nano_bench_test.go` → should be merged into `fake_test.go`
- `optimization_bench_test.go` → should be merged into `fake_test.go`
- `performance_first_bench_test.go` → should be merged into `fake_test.go`
- `verification_test.go` → should be merged into `fake_test.go`

#### `/runtime/` directory
**Source files:** `runtime.go`, `exit.go`, `system_*.go`, `exit_signal*.go`
**Current test files need major consolidation:**
- `runtime_test.go` ✅ (correctly named)
- `exit_test.go` ✅ (correctly named)
- `system_test.go` ✅ (correctly named)
- Multiple coverage/edge test files that should be merged:
  - `comprehensive_test.go` → merge into `runtime_test.go`
  - `coverage_test.go` → merge into `runtime_test.go`
  - `error_path_test.go` → merge into `runtime_test.go`
  - `error_paths_coverage_test.go` → merge into `runtime_test.go`
  - `exit_complete_test.go` → merge into `exit_test.go`
  - `exit_function_coverage_test.go` → merge into `exit_test.go`
  - `final_coverage_test.go` → merge into `runtime_test.go`
  - `missing_coverage_test.go` → merge into `runtime_test.go`
  - `missing_lines_test.go` → merge into `runtime_test.go`
  - `mock_error_test.go` → merge into `runtime_test.go`
  - `panic_edge_cases_test.go` → merge into `runtime_test.go`
  - `path_error_coverage_test.go` → merge into `runtime_test.go`
  - `stack_edge_test.go` → merge into `runtime_test.go`

#### `/cryptox/` directory
**Source files:** `aes.go`, `des.go`, `ecdh.go`, `ecdsa.go`, `hash_*.go`, `rsa.go`, `uuid.go`
**Test files needing consolidation:**
- `aes_test.go` ✅ and `aes_100_coverage_test.go` → merge both into `aes_test.go`
- `des_test.go` ✅ and `des_100_coverage_test.go` → merge both into `des_test.go`
- `ecdh_test.go` ✅ and `ecdsa_test.go` ✅ (correctly named)
- `ecc_100_coverage_test.go` → merge into `ecdh_test.go` and `ecdsa_test.go`
- `hash_basic_test.go`, `hash_crc_test.go`, `hash_fnv_test.go`, `hash_hmac_test.go` ✅ (correctly named)
- `rsa_test.go` ✅ and `rsa_100_coverage_test.go` → merge both into `rsa_test.go`
- `uuid_test.go` ✅ (correctly named)
- `missing_coverage_test.go` → distribute content to appropriate test files

#### `/stringx/` directory
**Source files:** `rand.go`, `string.go`, `strings.go`, `unicode.go`, `utf16.go`
**Test files needing consolidation:**
- `rand_test.go`, `string_test.go`, `strings_test.go`, `unicode_test.go`, `utf16_test.go` ✅ (correctly named)
- `benchmark_test.go` → merge into appropriate individual test files
- `complete_coverage_test.go` → merge into appropriate individual test files
- `missing_coverage_test.go` → merge into appropriate individual test files

#### `/human/` directory
**Source files:** `human.go`, `locale*.go`, `options.go`
**Test files needing consolidation:**
- `human_test.go` ✅ (correctly named)
- `locale_test.go` ✅ (correctly named for general locale functionality)
- Multiple test files that should be merged:
  - `api_test.go` → merge into `human_test.go`
  - `complete_test.go` → merge into `human_test.go`
  - `coverage_test.go` → merge into `human_test.go`
  - `edge_cases_test.go` → merge into `human_test.go`

### 2. Missing Test Files for Source Files

#### `/validator/` directory (Major Issues)
**Missing proper test files for:**
- `validator.go` → needs `validator_test.go` (currently has `validator_simple_test.go`)
- `custom_validators.go` → needs `custom_validators_test.go`
- `engine.go` → needs `engine_test.go`
- `options.go` → needs `options_test.go`
- `types.go` → needs `types_test.go`
- All locale files (`locale_*.go`) → currently share `locale_test.go`

**Current test files:**
- `locale_test.go` (tests multiple locale files)
- `simple_coverage_test.go` → should be merged into proper test files
- `validator_simple_test.go` → should be renamed to `validator_test.go`

#### `/app/` directory
**Source files without proper test files:**
- `env.go`, `release.go`, `test.go`, `beta.go`, `alpha.go`, `debug.go`, `info.go`
- All currently use shared `app_test.go`

#### Other directories with missing test files:
- `/urlx/query.go` → needs `query_test.go` (currently has `urlx_test.go`)
- `/routine/cache.go` and `/routine/group.go` → need individual test files
- Various platform-specific files in `/runtime/` and `/atexit/`

## Reorganization Strategy

### Phase 1: Consolidate Existing Test Files
1. **Merge benchmark and coverage test files** into their corresponding main test files
2. **Preserve all test functions** - no tests should be lost
3. **Combine benchmarks** with unit tests in the same file
4. **Maintain test package structure** and imports

### Phase 2: Rename Test Files to Follow Convention
1. **Rename incorrectly named test files** to match their source files
2. **Split shared test files** where multiple source files currently share one test file
3. **Create missing test files** for source files without tests

### Phase 3: Clean Up and Organize
1. **Remove duplicate test cases** that might exist across multiple files
2. **Organize test functions** logically within each file
3. **Ensure consistent test naming** and structure

## Detailed Reorganization Actions

### High Priority Consolidations

#### 1. `/fake/` directory
```bash
# Merge into fake_test.go:
# - basic_test.go
# - benchmark_test.go
# - example_test.go
# - extreme_bench_test.go
# - extreme_example_test.go
# - nano_bench_test.go
# - optimization_bench_test.go
# - performance_first_bench_test.go
# - verification_test.go

# Keep separate (correctly named):
# - address_test.go (tests address.go)
# - names_test.go (tests names.go)
```

#### 2. `/runtime/` directory
```bash
# Merge into runtime_test.go:
# - comprehensive_test.go
# - coverage_test.go
# - error_path_test.go
# - final_coverage_test.go
# - missing_coverage_test.go
# - missing_lines_test.go
# - mock_error_test.go
# - panic_edge_cases_test.go
# - path_error_coverage_test.go
# - stack_edge_test.go

# Merge into exit_test.go:
# - exit_complete_test.go
# - exit_function_coverage_test.go
# - error_paths_coverage_test.go

# Keep separate:
# - system_test.go (tests system_*.go files)
```

#### 3. `/cryptox/` directory
```bash
# Merge coverage files into main test files:
# aes_100_coverage_test.go → aes_test.go
# des_100_coverage_test.go → des_test.go
# rsa_100_coverage_test.go → rsa_test.go
# ecc_100_coverage_test.go → distribute to ecdh_test.go and ecdsa_test.go
# missing_coverage_test.go → distribute to appropriate test files
```

### Medium Priority Reorganizations

#### 1. `/stringx/` directory
```bash
# Distribute benchmark_test.go content to:
# - rand_test.go
# - string_test.go
# - strings_test.go
# - unicode_test.go
# - utf16_test.go

# Merge coverage files into appropriate individual test files
```

#### 2. `/human/` directory
```bash
# Merge into human_test.go:
# - api_test.go
# - complete_test.go
# - coverage_test.go
# - edge_cases_test.go

# Keep locale_test.go for shared locale functionality
```

### Low Priority (Create Missing Test Files)

#### 1. `/validator/` directory
```bash
# Rename:
# validator_simple_test.go → validator_test.go

# Create new test files:
# - custom_validators_test.go
# - engine_test.go
# - options_test.go
# - types_test.go

# Consider splitting locale_test.go if needed for individual locale files
```

#### 2. Other missing test files
```bash
# /urlx/: Create query_test.go (may need to split from urlx_test.go)
# /app/: Consider creating individual test files or keeping shared app_test.go
# /routine/: Create cache_test.go and group_test.go
```

## Implementation Guidelines

### File Merging Process
1. **Preserve all test functions** - copy, don't move
2. **Combine imports** and remove duplicates
3. **Maintain benchmark functions** alongside unit tests
4. **Keep test helper functions** and test data
5. **Preserve build tags** and platform-specific tests

### Testing Strategy
1. **Run tests before and after** each reorganization to ensure no regressions
2. **Verify test coverage** remains the same or improves
3. **Check that benchmarks still run** correctly
4. **Validate build tags** work on appropriate platforms

### File Naming Convention
- Source file: `package.go`
- Test file: `package_test.go`
- All unit tests, benchmarks, and examples in the same test file
- Use build tags for platform-specific tests if needed

## Expected Benefits

1. **Simplified test discovery** - easy to find tests for any source file
2. **Improved maintainability** - tests and benchmarks in one place
3. **Consistent project structure** - follows Go conventions
4. **Easier CI/CD integration** - standard test file patterns
5. **Better developer experience** - predictable test file locations

## Risk Assessment

### Low Risk
- **candy/**, **cache/** directories - already well-organized
- **xtime/** directories - mostly following conventions

### Medium Risk
- **cryptox/**, **stringx/** - multiple coverage files to merge
- **human/** - several small test files to consolidate

### High Risk
- **fake/** - many benchmark files with potential conflicts
- **runtime/** - complex platform-specific code and many coverage files
- **validator/** - may require significant restructuring

## Success Criteria

1. Every `.go` source file has a corresponding `_test.go` file with the same base name
2. No separate benchmark or coverage-only test files remain
3. All existing tests pass without modification to test logic
4. Test coverage percentage remains the same or improves
5. Build and test times are not negatively impacted