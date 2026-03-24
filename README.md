# Go-Tdarr-Module

A Go client library for the [Tdarr](https://tdarr.io) API v2.

## Installation
```bash
go get github.com/dennix85/Go-Tdarr-Module/v2
```

## Usage
```go
import tdarr "github.com/dennix85/Go-Tdarr-Module/v2"

client := tdarr.NewClient("http://localhost:8265",
    tdarr.WithToken("your-api-token"),
)
```
