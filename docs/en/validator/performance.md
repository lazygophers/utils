---
title: Performance & Best Practices - Validator
---

# Performance & Best Practices

## Performance

Built-in algorithm variants benchmarked to select optimal implementation:

| Validation | Optimization |
|------------|-------------|
| Message formatting | 5 template implementations (original/builder/compiled/byte-slice/no-fmt) |
| Email | Multiple regex/non-regex approaches compared |

Run benchmarks:

```bash
go test -bench=. ./validator/
```

## Best Practices

**Tag design**: Combine 2-3 tags per field, not one catch-all.

```go
// ✅ Recommended
Email string `validate:"required,email"`
Name  string `validate:"required,min=2,max=50"`

// ❌ Too loose
Email string `validate:"required"`
Name  string `validate:"required"`
```

**Validation timing**: Validate at the first entry point (config loading, request parsing), not deep in business logic.

**Error handling**: Use `ValidationErrors` type assertion to generate user-friendly messages per field.

**Reuse instances**: Use `Default()` or cache a `New()` instance instead of creating per call.
