[![Build Status](https://travis-ci.org/MarcGrol/golangAnnotations.svg?branch=master)](https://travis-ci.org/MarcGrol/golangAnnotations)
[![Coverage Status](https://coveralls.io/repos/github/MarcGrol/golangAnnotations/badge.svg)](https://coveralls.io/github/MarcGrol/golangAnnotations)
# Golang annotations

[Detailed explanation](https://github.com/MarcGrol/golangAnnotations/wiki)

## Summary

The golangAnnotations-tool parses your golang source-code into an intermediate representation.
Using this intermediate representation, the tool uses your annotations to generate source code that would be cumbersome and error-prone to write manually. Bottom line, a lot less code needs to be written.

Example:
    
    // @RestOperation( method = "GET", path = "/person/{uid}" )
    func (s Service) getPerson(uid string) (*Person,error) {
        ...
    } 

## Getting the software

    $ go get -u -t github.com/MarcGrol/golangAnnotations

## Testing and installing

    $ make gen
    $ make test
    $ make install
    
    or
    
    $ make

## Currently supported annotations

This first implementation provides the following kind of annotations:
- web-services (jax-rs like):
    - Generate server-side http-handling for a "service"
    - Generate client-side http-handling for a "service"
    - Generate helpers to ease integration testing of your services

- event-sourcing:
    - Describe which events belong to which aggregate
    - Type-strong boiler-plate code to build an aggregate from individual events
    - Type-strong boiler-plate code to wrap and unwrap events into an envelope so that it can be easily stored and emitted

## How to use http-server related annotations ("jax-rs"-like)?

A regular golang struct definition with our own "RestService" and "RestOperation"-annotations. Observe that [./examples/rest/tourService.go](./examples/rest/tourService.go) is used as input.

    // @RestService( path = "/api" )
    type Service struct {
       ...
    }
    
    // @RestOperation( method = "GET", path = "/person/{uid}" )
    func (s Service) getPerson(uid string) (*Person, error) {
        ...
    }        

Observe that ./examples/rest/$tourService.go have been generated.

[Example](https://github.com/MarcGrol/golangAnnotations/wiki/example-of-generated-code) of the generated http handler.

## How to use event-sourcing related annotations?

A regular golang struct definition with our own "Event"-annotation.
    
    // @Event( aggregate = Tour" )
    type TourEtappeCreated struct {
        ...
    }        

Observe that ./examples/event/$wrappers.go and ./examples/event/$aggregates.go have been created in ./examples/structExample.

### Command to trigger code-generation:

We use the "go:generate" mechanism to trigger our goAnnotations-executable.
In order to trigger this mechanisme we use a '//go:genarate' comment with the command to be executed.

example:

    //go:generate golangAnnotations -input-dir .

So can can use the regular toolchain to trigger code-genaration

    $ cd ${GOPATH/src/github.com/MarcGrol/golangAnnotations
    $ go generate ./...
    // go imports will fix all the imports
    $ for i in `find . -name "*.go"`; do goimports -w -local github.com/ ${i}; done
    // fixes formatting for generated code
    $ for i in `find . -name "*.go"`; do gofmt -s -w ${i}; done
    
