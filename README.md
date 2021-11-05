# Conjur-Go

https://github.com/pkcwong/conjur-go

![GitWorkflow](https://github.com/conjur-dev/conjur-go/workflows/master/badge.svg)

A full stack boilerplate with Golang + ReactJS.

## Prerequisites and Installation

Install the following frameworks and packages.

- [Go](https://golang.org/)
- [NodeJS](https://nodejs.org/en/)
- [Yarn](https://yarnpkg.com/)

Verify the installation.

```bash
go version
node --version
yarn --version
```

Clone this repository.

```bash
git clone https://github.com/conjur-dev/conjur-go.git <PROJECT_NAME>
cd <PROJECT_NAME>
git submodule update --init --recursive
```

Building backend.

```bash
go build
```

Building frontend.

```bash
cd client
yarn build
```

## Development

Start the backend server, default listening on port :8080.

```bash
go run .
```

Start a development server for frontend, default listening on port :3000.

```bash
cd client
yarn start
```

## Production

Create an optimized frontend production build.

```bash
cd client
yarn build
```

Start the backend server.

```bash
go run .
```
