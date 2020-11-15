# Common Module for Event-driven Architecture

Designed for Command Query Responsibility Segregation (CQRS) and event sourcing.

[Github Repository](https://github.com/ivorscott/devpie-client-events)

## Install

```
go get github.com/ivorscott/devpie-client-events
```

Example:

```go
package main

import (
	"fmt"

	"github.com/ivorscott/devpie-client-events/go/events"
)

func main() {
	command := events.CommandsRegisterUser
	fmt.Println(command)

	event := events.EventsUserRegistered
	fmt.Println(event)
}
```
