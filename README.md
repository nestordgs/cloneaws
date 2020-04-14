# Clone repos from _AWS CodeCommit_

Console `Cli` to clone repositories from _`AWS CodeCommit`_ without `urlencode` username, password manualy.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

* Go

### Installing

First go to root project and run

```bash
go run install
```

And use as

```bash
cloneaws clone --profile="your-profile" --url="https://url" --projectName="test"
```

If you want to run without install in the root project use

```bash
go run main.go clone --profile="your-profile" --url="https://url" --projectName="test"
```

## Authors

* **Nestor Gutiérrez**