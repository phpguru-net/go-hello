# Overview

## How to run test?

### Run al test

```bash
go test -timeout 30s go-hello.phpguru.net/test/gross_store
```

### Specific function

```bash
# -run ^TestAddItem$
go test -timeout 30s -run ^TestAddItem$ go-hello.phpguru.net/test/gross_store
```

### Streaming output

```bash
go test -v -timeout 30s go-hello.phpguru.net/test/gross_store
```
