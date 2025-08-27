# go-calver

A Go library for parsing, validating, and manipulating Calendar Versioning
(CalVer) strings according to the [CalVer specification](https://calver.org/).

## What is CalVer?

Calendar Versioning (CalVer) is a versioning scheme that uses calendar dates for
version numbers. Unlike Semantic Versioning (SemVer) which focuses on API
compatibility, CalVer emphasizes when something was released, making it ideal
for projects that release frequently or on a schedule.

## Features

- **Flexible Format Support**: Supports all standard CalVer conventions
  including `<YYYY>`, `<YY>`, `<0Y>`, `<MM>`, `<0M>`, `<WW>`, `<0W>`, `<DD>`,
  `<0D>`, `<MINOR>`, `<MICRO>`, and `<MODIFIER>`
- **Format Validation**: Ensures version strings match their specified format
- **Comparison Operations**: Compare CalVer versions with proper precedence
  handling
- **Collections**: Sort and manage collections of CalVer objects
- **Zero Dependencies**: Pure Go implementation with no external dependencies
- **Comprehensive Testing**: Extensive test coverage for all functionality
- **Unlimited Format Support**: Supports any format string since users control
  the format - the only requirement is to use the CalVer conventions correctly

## Installation

```bash
go get github.com/shazib-summar/go-calver
```

## Quick Start

```go
package main

import (
	"fmt"
	"log"

	"github.com/shazib-summar/go-calver/calver"
)

func main() {
	format := "Rel-<YYYY>-<0M>-<0D>"
	// Create a new CalVer object
	ver, err := calver.NewCalVer(format, "Rel-2025-07-14")
	if err != nil {
		log.Fatal(err)
	}

	// Print the version
	fmt.Println(ver.String()) // Output: Rel-2025-07-14

	// Compare with another version
	other, _ := calver.NewCalVer(format, "Rel-2025-07-15")
	result, _ := ver.Compare(other)
	fmt.Printf("Comparison result: %d\n", result) // Output: -1 (less than)
}

```

## Supported Formats

The library supports all standard CalVer conventions:

| Convention   | Description                                | Example                  |
| ------------ | ------------------------------------------ | ------------------------ |
| `<YYYY>`     | 4-digit year                               | `2025`                   |
| `<YY>`       | 1-2 digit year                             | `25`                     |
| `<0Y>`       | 2-digit year (zero-padded)                 | `05`                     |
| `<MAJOR>`    | The "major" part                           | `12, 02, 123`            |
| `<MM>`       | 1-2 digit month                            | `7` or `12`              |
| `<0M>`       | 2-digit month (zero-padded)                | `07` or `12`             |
| `<MINOR>`    | Minor version number                       | `14`                     |
| `<WW>`       | 1-2 digit week                             | `1` or `52`              |
| `<0W>`       | 2-digit week (zero-padded)                 | `01` or `52`             |
| `<DD>`       | 1-2 digit day                              | `1` or `31`              |
| `<0D>`       | 2-digit day (zero-padded)                  | `01` or `31`             |
| `<MICRO>`    | Micro version number                       | `42`                     |
| `<MODIFIER>` | Modifier string or additional version part | `alpha`, `beta`, `12:43` |

## Usage Examples

Complete examples files can be found in the [examples](examples) dir

### Basic Version Creation

```go
// Year-Month-Day format
calver, err := calver.NewCalVer("<YYYY>-<MM>-<DD>", "2025-07-14")
if err != nil {
    log.Fatal(err)
}

// Year.Release format
calver, err = calver.NewCalVer("<YYYY>.R<DD>", "2025.R14")
if err != nil {
    log.Fatal(err)
}

// Ubuntu-style format
calver, err = calver.NewCalVer("<0Y>.<0M>.<DD>", "22.04.6")
if err != nil {
    log.Fatal(err)
}
```

### Version Comparison

```go
calver1, _ := calver.NewCalVer("<YYYY>-<MM>-<DD>", "2025-07-14")
calver2, _ := calver.NewCalVer("<YYYY>-<MM>-<DD>", "2025-07-15")

result, err := calver1.Compare(calver2)
if err != nil {
    log.Fatal(err)
}

switch result {
case -1:
    fmt.Println("calver1 is older than calver2")
case 0:
    fmt.Println("calver1 equals calver2")
case 1:
    fmt.Println("calver1 is newer than calver2")
}
```

### Working with Collections

```go
versions := []string{
    "2025-07-14",
    "2025-07-15",
    "2025-07-13",
}

collection, err := calver.NewCollection("<YYYY>-<MM>-<DD>", versions...)
if err != nil {
    log.Fatal(err)
}

// Sort the collection
sort.Sort(collection)

// Print sorted versions
for _, v := range collection {
    fmt.Println(v.String())
}
```

### Custom Format with Modifiers

```go
// Release format with timestamp modifier
format := "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z"
version := "RELEASE.2025-07-23T15-54-02Z"

calver, err := calver.NewCalVer(format, version)
if err != nil {
    log.Fatal(err)
}

fmt.Println(calver.String()) // Output: RELEASE.2025-07-23T15-54-02Z
```

## Testing

Run the test suite:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major
changes, please open an issue first to discuss what you would like to change.

### Development Setup

1. Fork the repository
2. Clone your fork: `git clone https://github.com/yourusername/go-calver.git`
3. Create a feature branch: `git checkout -b feature/amazing-feature`
4. Make your changes and add tests
5. Run tests: `go test ./...`
6. Commit your changes: `git commit -m 'Add amazing feature'`
7. Push to the branch: `git push origin feature/amazing-feature`
8. Open a Pull Request

## License

This project is licensed under the Apache License 2.0 - see the
[LICENSE](LICENSE) file for details.

## Acknowledgments

- [CalVer.org](https://calver.org/) for the Calendar Versioning specification
- The Go community for best practices and testing patterns
- My playful niece Abigail without whom this would've been done much sooner.
