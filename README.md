## go-ucsm-sdk: Go library for Cisco UCS

`go-ucsm-sdk` is a Go library for interfacing with UCS Manager XML API

### History

This library is a fork from https://github.com/dnaeon/go-ucs

## Documentation

The library implements the following subset of Cisco UCS Manager XML API:

-   AaaLogin
-   AaaRefresh
-   AaaKeepAlive
-   AaaLogout
-   ConfigResolveDn
-   ConfigResolveDns
-   ConfigResolveClass
-   ConfigResolveClasses
-   ConfigResolveChildren
-   orgResolveElements
-   ConfigConfMo
-   ConfigConfMos
-   ConfigEstimateImpact
-   LsInstantiateNTemplate
-   LsInstantiateNNamedTemplate

See Cisco UCS Manager XML API Programmer's Guide at https://www.cisco.com/c/en/us/td/docs/unified_computing/ucs/sw/api/b_ucs_api_book/b_ucs_api_book_chapter_00.html

## Installation

In order to install `go-ucsm-sdk` execute the following command:

```bash
   go get -v github.com/cloud104/tks-go-ucsm-sdk
```

## Tests

```
TBD
```

## Examples

Check the included docs/examples from this repository.

[server-get](docs/exemples/server-get_example.md) is a good starting point to see how to use the library.


Please note that most of the examples are using high-level utility calls to abstart a caller from XML API.
