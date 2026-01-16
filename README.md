# cp-mgmt-api-go-sdk

Check Point API Go Development Kit simplifies the use of the Check Point Management APIs. The kit contains the API library files, and sample files demonstrating the capabilities of the library. The kit is compatible with Golang version 1.12.

## Content

`APIFiles` - the API library project.

`add_access_rule` - demonstrates a basic flow of using the APIs: performs a login command, adds an access rule to the top of the access policy layer, and publishes the changes.

`add_host` - demonstrates creating a new host.

`show_hosts` - demonstrates showing all of the hosts of the domain.

`discard_sessions` - demonstrates how to discard the changes to the database for un-published sessions.

`find_duplicate_ip` - demonstrates searching for all the hosts that share the same IP address.

`auto_publish` - demonstrates using auto publish feature by running multiple threads that create new network groups.

## Instructions

The SDK can be imported to another Go project:

```bash
go get github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles
```

```go
package main

import "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
```

### Running the examples

Clone the repository:

```sh
git clone https://github.com/CheckPointSW/cp-mgmt-api-go-sdk.git
cd cp-mgmt-api-go-sdk
```

Build and run the project:

```sh
go build main.go
./main <example argument> ["discard"/"rule"/"add_host"/"show_hosts"/"dup_ip"/"auto_publish"]
```
