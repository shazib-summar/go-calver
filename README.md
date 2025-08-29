[![Go Reference](https://pkg.go.dev/badge/github.com/shazib-summar/go-calver.svg)](https://pkg.go.dev/github.com/shazib-summar/go-calver)
[![Go Report Card](https://goreportcard.com/badge/github.com/shazib-summar/go-calver)](https://goreportcard.com/report/github.com/shazib-summar/go-calver)
[![Tests Status](https://github.com/shazib-summar/go-calver/actions/workflows/codestyle.yml/badge.svg?branch=main)](https://github.com/shazib-summar/go-calver/actions)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](CODE_OF_CONDUCT.md)

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
- **Version Incrementing**: Increment major, minor, micro, and modifier versions
  while preserving zero-padding
- **Series Management**: Extract version series at different levels (major,
  minor, micro, modifier)
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

    "github.com/shazib-summar/go-calver"
)

func main() {
    format := "Rel-<YYYY>-<0M>-<0D>"
    // Create a new Version object
    ver, err := calver.NewVersion("Rel-2025-07-14", format)
    if err != nil {
        log.Fatal(err)
    }

    // Print the version
    fmt.Println(ver.String()) // Output: Rel-2025-07-14

    // Compare with another version
    other, _ := calver.NewVersion("Rel-2025-07-15", format)
    result, _ := ver.Compare(other)
    fmt.Printf("Comparison result: %d\n", result) // Output: -1 (less than)
}
```

You can pass more than one format strings as well. If multiple formats are
provided, the format that matches the version string the most will be used. For
example, in the following code, the format `"Rel-<YYYY>-<0M>-<0D>"` will be used
at it returns a greater number of matched regex groups than the other formats.

```go
ver, err := NewVersion(
    "Rel-2025-07-14",
    "Rel-<YYYY>",
    "Rel-<YYYY>-<0M>",
    "Rel-<YYYY>-<0M>-<0D>",
)
if err != nil {
    return err
}
fmt.Println(ver.String()) // Rel-2025-07-14
```

## Supported Formats

The library supports all standard CalVer conventions, organized into four levels
that determine the order when comparing versions. Only one convention string may
be used per level in the format string provided to `NewVersion` func.

### Levels and Conventions

| Level        | Description                  | Conventions                               | Example                  |
| ------------ | ---------------------------- | ----------------------------------------- | ------------------------ |
| **Major**    | Primary version identifier   | `<YYYY>`, `<YY>`, `<0Y>`, `<MAJOR>`       | `2025`, `25`, `05`, `12` |
| **Minor**    | Secondary version identifier | `<MM>`, `<0M>`, `<MINOR>`                 | `7`, `07`, `14`          |
| **Micro**    | Tertiary version identifier  | `<WW>`, `<0W>`, `<DD>`, `<0D>`, `<MICRO>` | `1`, `01`, `31`, `42`    |
| **Modifier** | Additional version metadata  | `<MODIFIER>`                              | `alpha`, `beta`, `12:43` |

### Convention Details

| Convention   | Description                                | Regex                |
| ------------ | ------------------------------------------ | -------------------- |
| `<YYYY>`     | 4-digit year                               | `(?P<major>\d{4})`   |
| `<YY>`       | 1-2 digit year                             | `(?P<major>\d{1,2})` |
| `<0Y>`       | 2-digit year (zero-padded)                 | `(?P<major>\d{2})`   |
| `<MAJOR>`    | Major version number                       | `(?P<major>\d+)`     |
| `<MM>`       | 1-2 digit month                            | `(?P<minor>\d{1,2})` |
| `<0M>`       | 2-digit month (zero-padded)                | `(?P<minor>\d{2})`   |
| `<MINOR>`    | Minor version number                       | `(?P<minor>\d+)`     |
| `<WW>`       | 1-2 digit week                             | `(?P<micro>\d{1,2})` |
| `<0W>`       | 2-digit week (zero-padded)                 | `(?P<micro>\d{2})`   |
| `<DD>`       | 1-2 digit day                              | `(?P<micro>\d{1,2})` |
| `<0D>`       | 2-digit day (zero-padded)                  | `(?P<micro>\d{2})`   |
| `<MICRO>`    | Micro version number                       | `(?P<micro>\d+)`     |
| `<MODIFIER>` | Modifier string or additional version part | `(?P<modifier>.*)`   |

## Usage Examples

Complete examples files can be found in the [examples](examples) dir. Example can be run with the following command:

```bash
make examples
```

### Basic Version Creation

```go
// Year-Month-Day format
ver, err := calver.NewVersion("2025-07-14", "<YYYY>-<MM>-<DD>")
if err != nil {
    log.Fatal(err)
}

// Year.Release format
ver, err = calver.NewVersion("2025.R14", "<YYYY>.R<DD>")
if err != nil {
    log.Fatal(err)
}

// Ubuntu-style format
ver, err = calver.NewVersion("22.04.6", "<0Y>.<0M>.<DD>")
if err != nil {
    log.Fatal(err)
}
```

### Version Comparison

```go
verA, _ := calver.NewVersion("2025-07-14", "<YYYY>-<MM>-<DD>")
verB, _ := calver.NewVersion("2025-07-15", "<YYYY>-<MM>-<DD>")

result, err := verA.Compare(verB)
if err != nil {
    log.Fatal(err)
}

switch result {
case -1:
    fmt.Println("verA is older than verB")
case 0:
    fmt.Println("verA equals verB")
case 1:
    fmt.Println("verA is newer than verB")
}
```

### Working with Collections

```go
versions := []string{
    "2025-07-14",
    "2025-07-15",
    "2025-07-13",
}

collection, err := calver.NewCollection(versions, "<YYYY>-<MM>-<DD>")
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

### Version Incrementing

```go
// Create a version
ver, err := calver.NewVersion("2025.07.14", "<YYYY>.<0M>.<0D>")
if err != nil {
    log.Fatal(err)
}

// Increment different parts
err = ver.IncMajor()   // 2025 -> 2026
err = ver.IncMinor()   // 07 -> 08
err = ver.IncMicro()   // 14 -> 15

fmt.Println(ver.String()) // Output: 2026.08.15

// Zero-padding is preserved
ver, _ = calver.NewVersion("2025.01.09", "<YYYY>.<0M>.<0D>")
err = ver.IncMinor()   // 01 -> 02 (preserves zero-padding)
err = ver.IncMicro()   // 09 -> 10 (loses zero-padding)

fmt.Println(ver.String()) // Output: 2025.02.10
```

### Series Management

```go
ver, err := calver.NewVersion("Rel-2025-07-14", "Rel-<YYYY>-<0M>-<0D>")
if err != nil {
    log.Fatal(err)
}

// Get series at different levels
fmt.Println(ver.Series("major"))    // Output: Rel-2025
fmt.Println(ver.Series("minor"))    // Output: Rel-2025-07
fmt.Println(ver.Series("micro"))    // Output: Rel-2025-07-14
fmt.Println(ver.Series("modifier")) // Output: Rel-2025-07-14
fmt.Println(ver.Series(""))         // Output: Rel-2025-07-14 (full version)

// Useful for grouping related versions
majorSeries := ver.Series("major") // "Rel-2025"
minorSeries := ver.Series("minor") // "Rel-2025-07"
```

### Custom Format with Modifiers

```go
// Release format with timestamp modifier
format := "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z"
version := "RELEASE.2025-07-23T15-54-02Z"

ver, err := calver.NewVersion(version, format)
if err != nil {
    log.Fatal(err)
}

fmt.Println(ver.String()) // Output: RELEASE.2025-07-23T15-54-02Z
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
