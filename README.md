![Go Report Card](https://goreportcard.com/badge/github.com/renatosaksanni/date-range-overlap)

# Date Range Overlap

This project provides a Go wrapper to check for overlapping date ranges. It utilizes a simple and efficient algorithm to determine if two date ranges overlap.

# Summary of "The Overlapping Date Range Test"
The algorithm for determining if two date ranges overlap is as follows: Two periods overlap if the start date of one is before the end date of the other, and vice versa. This can be implemented in Go using the following condition:
```go
if (period.End.After(selection.Start) && period.Start.Before(selection.End)) {
    // period overlaps selection
}
```
For more information, visit the original article: [The Overlapping Date Range Test](https://logic-and-trick.com/page/the-overlapping-date-range-test.html).

# Setup
### Prerequisites
- Go 1.20 or higher
### Install Dependencies
Clone the repository and run:
```sh
go mod tidy
```
# Usage
### Running the Main Program
To run the main program:
```sh
go run cmd/main.go
```
### Running Tests
To run the unit tests:
```sh
go test ./... -coverprofile=coverage.out
```
