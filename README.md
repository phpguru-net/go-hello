# Golang

## Resources

- [Go Packages](https://pkg.go.dev)

## Essentials

1. Entry Point

```go
import "fmt"
package main
func main() {
  fmt.Println("Entry point of go application")
}
```

2. Package management with go mod

> go mod init {package-namespace}

**About package namespace:**

- Should be distinct with other package
- It's better to name with this format

> {repository-host}/{repository-name}

```bash
go mod init github.com/phpguru-net/go-solutions
```

**Use the external package**

Eg : https://pkg.go.dev/search?q=phpguru&m=package

```bash
go get github.com/phpguru-net/go-hello
```

## 1. Hello world

### Basic Concept

1. Application Structure
2. Variables, Constants, Datatypes
3. Input, Output
4. Function, Scope
5. Debug
6. Test
7. Conditional Statement
8. Loop
9. Common Packages: Math, String
10. Data Structure
11. File I/O
12. Error

### Excercises
