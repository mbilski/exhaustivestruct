# exhaustivestruct

[![Go Report Card](https://goreportcard.com/badge/github.com/mbilski/exhaustivestruct)](https://goreportcard.com/badge/github.com/mbilski/exhaustivestruct)

exhaustivestruct is a go static analysis tool to find structs that have uninitialized fields.

> :warning: This linter is meant to be used only for special cases.
> It is not recommended to use it for all files in a project.

## Installation

```
go get -u github.com/mbilski/exhaustivestruct/cmd/exhaustivestruct
```

## Usage

```
Usage: exhaustivestruct [-flag] [package]

Flags:
  -struct_patterns string
      This is a comma separated list of expressions to match struct packages and names
```

## Example

``` go
type User struct {
  Name string
  Age int
}

var user = User{ // fails with "Age is missing in User"
  Name: "John",
}
```

