# recotem-cli

## Usage

### Install

```
$ go install recotem.org/cli/recotem@latest
```

### Use recotem command

```
$ recotem <command> <sub-command> ...
```

## Development

### OpenAPI

#### Setup

```
$ go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
```

#### Update Generated Code

Update `pkg/openapi/recotem.yaml ` and then

```
$ oapi-codegen -generate types,client pkg/openapi/recotem.yaml > pkg/openapi/recotem.gen.go
```
 
