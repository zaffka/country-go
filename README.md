# Country Package

[![Go Reference](https://pkg.go.dev/badge/github.com/zaffka/country-go.svg)](https://pkg.go.dev/github.com/zaffka/country-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/zaffka/country-go)](https://goreportcard.com/report/github.com/zaffka/country-go)
[![Tests](https://github.com/zaffka/country-go/actions/workflows/tests.yaml/badge.svg)](https://github.com/zaffka/country-go/actions/workflows/tests.yaml)

A lightweight Go package providing ISO 3166-1 country code lookups and validations.

## Features

- Lookup countries by:
  - Name (case-insensitive)
  - Alpha-2 code (e.g., "US")
  - Alpha-3 code (e.g., "USA") 
  - ISO numeric code (e.g., 840)
- Simplified country existence checks
- Zero dependencies (except embedded JSON data)


## Installation

```bash
go get github.com/zaffka/country-go@latest
```

## Usage
```go
import "github.com/zaffka/country-go"

// Lookup by name
c, err := country.ByName("canada") // Case-insensitive

// Lookup by Alpha-2 code
c, err := country.ByAlpha2Code("ca")

// Check if code exists
if country.ExistsAlpha2("FR") {
    // France exists
}
```
