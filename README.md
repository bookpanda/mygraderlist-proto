# MyGraderList Proto

MyGraderList is a web app that lets students assess the difficulties and worthiness of each DSA grader problem in their respective courses.

MyGraderList Proto is the central library that contains all the protobuf messages and services used by the MyGraderList project.

## Technologies

-   protobuf

## Getting Started

### Prerequisites

-   golang 1.21 or [later](https://go.dev)
-   makefile

### Installation

1. Clone this repo
2. Run `go mod download` to download all the dependencies.

### Usage
1. Edit `*.proto` files
2. Run `make proto` to generate the golang implementation of the protobuf files.
