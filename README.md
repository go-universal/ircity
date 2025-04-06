# ircity

![GitHub Tag](https://img.shields.io/github/v/tag/go-universal/ircity?sort=semver&label=version)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-universal/ircity.svg)](https://pkg.go.dev/github.com/go-universal/ircity)
[![License](https://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/go-universal/ircity/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-universal/ircity)](https://goreportcard.com/report/github.com/go-universal/ircity)
![Contributors](https://img.shields.io/github/contributors/go-universal/ircity)
![Issues](https://img.shields.io/github/issues/go-universal/ircity)

`ircity` is a Go package that provides an in-memory database of provinces and cities in Iran. It allows you to retrieve lists of provinces and cities, as well as find specific provinces or cities by their codes.

## Installation

To install the package, run:

```sh
go get github.com/go-universal/ircity
```

## Usage

Here is an example of how to use the `ircity` package:

```go
package main

import (
    "fmt"
    "github.com/go-universal/ircity"
)

func main() {
    // Get list of provinces
    provinces := ircity.Provinces()
    fmt.Println("Provinces:", provinces)

    // Find a province by code
    province := ircity.FindProvince(1)
    if province != nil {
        fmt.Println("Found Province:", *province)
    } else {
        fmt.Println("Province not found")
    }

    // Get list of cities in a province
    cities := ircity.Cities(1)
    fmt.Println("Cities in Province 1:", cities)

    // Find a city by code
    city := ircity.FindCity(101)
    if city != nil {
        fmt.Println("Found City:", *city)
    } else {
        fmt.Println("City not found")
    }
}
```

## API

### `Provinces() []Province`

Returns a list of all provinces.

```go
provinces := ircity.Provinces()
fmt.Println(provinces)
```

### `FindProvince(code uint) *Province`

Finds and returns a province by its code. Returns `nil` if the province is not found.

- `code` (uint): The code of the province to find.

```go
province := ircity.FindProvince(1)
if province != nil {
    fmt.Println("Found Province:", *province)
} else {
    fmt.Println("Province not found")
}
```

### `Cities(province uint) []City`

Returns a list of cities in the specified province.

- `province` (uint): The code of the province whose cities you want to retrieve.

```go
cities := ircity.Cities(1)
fmt.Println("Cities in Province 1:", cities)
```

### `FindCity(code uint) *City`

Finds and returns a city by its code. Returns `nil` if the city is not found.

- `code` (uint): The code of the city to find.

```go
city := ircity.FindCity(101)
if city != nil {
    fmt.Println("Found City:", *city)
} else {
    fmt.Println("City not found")
}
```

## Data Structure

### `Province`

Represents a province in Iran.

- `Code` (uint): The unique code of the province.
- `Name` (string): The name of the province.

### `City`

Represents a city in Iran.

- `Code` (uint): The unique code of the city.
- `Name` (string): The name of the city.
- `ProvinceCode` (uint): The code of the province to which the city belongs.
- `Province` (\*Province): The province to which the city belongs.

## License

This project is licensed under the ISC License. See the [LICENSE](g:\go-universal\ircity\LICENSE) file for details.
