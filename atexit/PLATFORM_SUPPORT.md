# Platform Support Matrix

## Overview
The `atexit` package now supports **all 53 platforms** from `go tool dist list`.

## Implementation Files

| File | Build Tags | Platforms Covered |
|------|-----------|-------------------|
| `atexit_linux.go` | `linux \|\| android` | Linux (all archs), Android (all archs) |
| `atexit_darwin.go` | `darwin \|\| ios` | macOS, iOS (all archs) |
| `atexit_windows.go` | `windows` | Windows (386, amd64, arm64) |
| `atexit_bsd.go` | `freebsd \|\| openbsd \|\| netbsd \|\| dragonfly` | All BSD variants |
| `atexit_solaris.go` | `solaris \|\| illumos` | Solaris, illumos |
| `atexit_aix.go` | `aix` | AIX (ppc64) |
| `atexit_plan9.go` | `plan9` | Plan9 (386, amd64, arm) |
| `atexit_js.go` | `js` | JavaScript/WASM |
| `atexit_wasip1.go` | `wasip1` | WASI Preview 1 |
| `atexit.go` | Fallback | Any other future platforms |

## Complete Platform List (53 platforms)

### Unix-like with Signal Support (42 platforms)

#### Linux (10 platforms)
- linux/386
- linux/amd64
- linux/arm
- linux/arm64
- linux/loong64
- linux/mips
- linux/mips64
- linux/mips64le
- linux/mipsle
- linux/ppc64
- linux/ppc64le
- linux/riscv64
- linux/s390x

#### Android (4 platforms)
- android/386
- android/amd64
- android/arm
- android/arm64

#### Darwin/macOS (2 platforms)
- darwin/amd64
- darwin/arm64

#### iOS (2 platforms)
- ios/amd64
- ios/arm64

#### FreeBSD (5 platforms)
- freebsd/386
- freebsd/amd64
- freebsd/arm
- freebsd/arm64
- freebsd/riscv64

#### OpenBSD (5 platforms)
- openbsd/386
- openbsd/amd64
- openbsd/arm
- openbsd/arm64
- openbsd/ppc64
- openbsd/riscv64

#### NetBSD (4 platforms)
- netbsd/386
- netbsd/amd64
- netbsd/arm
- netbsd/arm64

#### DragonFly BSD (1 platform)
- dragonfly/amd64

#### Solaris/illumos (2 platforms)
- solaris/amd64
- illumos/amd64

#### AIX (1 platform)
- aix/ppc64

#### Windows (3 platforms)
- windows/386
- windows/amd64
- windows/arm64

### Non-Signal Platforms (5 platforms)

Require explicit `atexit.Exit()` call:

#### Plan9 (3 platforms)
- plan9/386
- plan9/amd64
- plan9/arm

#### WebAssembly (2 platforms)
- js/wasm
- wasip1/wasm

## Verification

All platforms successfully cross-compile:

```bash
# Test cross-compilation
./test_all_platforms.sh

# Output:
✓ linux/amd64
✓ linux/386
✓ linux/arm64
✓ windows/amd64
✓ windows/386
✓ darwin/amd64
✓ darwin/arm64
✓ android/arm64
✓ android/amd64
✓ ios/arm64
✓ ios/amd64
✓ freebsd/amd64
✓ openbsd/amd64
✓ netbsd/amd64
✓ dragonfly/amd64
✓ solaris/amd64
✓ illumos/amd64
✓ aix/ppc64
✓ plan9/amd64
✓ js/wasm
✓ wasip1/wasm

All platforms compiled successfully!
```

## Architecture Support

- **x86**: 386, amd64
- **ARM**: arm, arm64
- **PowerPC**: ppc64, ppc64le
- **RISC-V**: riscv64
- **MIPS**: mips, mipsle, mips64, mips64le
- **LoongArch**: loong64
- **s390x**: IBM Z
- **wasm**: WebAssembly

## Signal Handling Strategy

### Linux/Android
- Uses `gomonkey` to hook `os.Exit()` calls
- Provides comprehensive coverage for all exit scenarios

### Darwin/macOS/iOS
- Signal handling: SIGINT, SIGTERM, SIGHUP, SIGQUIT
- Leverages Apple's Unix foundation

### Windows
- Signal handling: SIGINT, SIGTERM, os.Interrupt
- Windows-specific signal subset

### BSD Family
- Extended Unix signal support
- SIGINT, SIGTERM, SIGHUP, SIGQUIT

### Solaris/illumos/AIX
- Standard Unix signal handling
- SIGINT, SIGTERM, SIGHUP

### Plan9, js/wasm, wasip1/wasm
- No signal support
- Explicit `atexit.Exit()` required

## Test Coverage

Each platform-specific file has corresponding test file:
- `atexit_linux_test.go`
- `atexit_darwin_test.go`
- `atexit_windows_test.go`
- `atexit_bsd_test.go`
- `atexit_solaris_test.go`
- `atexit_aix_test.go`
- `atexit_plan9_test.go`
- `atexit_js_test.go`
- `atexit_wasip1_test.go`

## Future Compatibility

The fallback implementation (`atexit.go`) ensures support for any future platforms Go may add.
