# the-game

## Table of Contents
- [the-game](#the-game)
	- [Table of Contents](#table-of-contents)
	- [Overview](#overview)
	- [Development Setup](#development-setup)
		- [Prerequisites](#prerequisites)
		- [Major Dependencies](#major-dependencies)
		- [Installation](#installation)
	- [Development Workflow](#development-workflow)
	- [Suggestions for improvements](#suggestions-for-improvements)
	- [License](#license)

## Overview
The application, `the-game`, utilizes the web framework `gin` for its development. This framework offers built-in support for features such as middleware support, JSON validation, routing, and error management. In addition, the application utilizes `gqlgen` for working with GraphQL. This library allows for the dynamic creation of models and resolvers for queries and mutations, based on the GraphQL schema. The ORM of choice for the application is `gorm`, which facilitates seamless interaction with SQL. Furthermore, `testify` and `go-sqlmock` are employed for writing unit tests, ensuring the robustness and reliability of the application.

## Development Setup
### Prerequisites
- [Docker](https://docs.docker.com/install/)
- [Git](https://git-scm.com/downloads)

### Major Dependencies
- [gin](https://github.com/gin-gonic/gin)
- [gorm](https://github.com/go-gorm/gorm)
- [gqlgen](https://github.com/99designs/gqlgen)
- [testify](https://github.com/stretchr/testify)
- [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)

### Installation
1. Clone the repository
``` shell
git clone git@github.com:samya-ak/the-game.git
```

1. Build required docker containers for source code and database and run the application
```
cd the-game
make docker/start
```

## Development Workflow
When you make any changes to existing codebase, here are the steps that someone must adhere to:

If you make changes to graphql schema located at `./graph/`, use following command to auto genrate models and resolvers.
```
make schema
```

Add required files to staging area. Unstaged files will be stashed by pre-commit.
```
git add ./your-file
```

Install pre-commit (Need to run only once when you're running pre-commit for the first time)
```
make docker/pre-commit-install
```

Run pre-commit before pushing to upstream for auto linting and formatting.
```
make docker/pre-commit
```

Available make commands
``` shell
# build image and start containers while removing unused containers
make docker/start

# stop containers
make docker/stop

# combination of above two commands, stop before starting again
make docker/stop-start

# install pre-commit
make docker/pre-commit-install

# run pre-commit
make docker/pre-commit

# uninstall pre-commit
make docker/pre-commit-uninstall

# run golangci-lint
make docker/golangci-lint:

# generates models, boilerplate resolvers using gqlgen from graphql schema
make schema
```

## Suggestions for improvements
- Integrate tracing to track request lifecycle
- Implement authentication and authorization for enhanced security
- Integrate Metrics for request-response latency, failure rate, status code frequency etc.
- Export logs to remote dashboard
- Implement Migration management and versioning
- Integrate CI/CD to ship app quickly and efficiently
- Implement test coverage dashboards

## License
MIT
