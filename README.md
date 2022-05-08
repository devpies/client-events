# Common Module for Event-driven Architecture

Supports commands and events (CQRS).

## Overview

This package serves as a shared library for all message interfaces, commands and events in the system. It ensures consistency and correctness while implementing the event data model across Applications, Microservices, Aggregators and various programming languages.

## How it works

Client's event data model is exported as command and event enums, and message interfaces, to enable easy lookup of the available identifiers in the system.

```typescript
import { Commands, Events } from "@devpie/client-events";

console.log(Commands.EnableAccounting);
// EnableAccounting

console.log(Events.MembershipAdded);
// MembershipAdded
```

Existing interfaces allow us to type check the message body being sent, ensuring correctness of implementation.

```typescript
export interface MembershipAddedEvent {
  id: string;
  type: Events.MembershipAdded;
  metadata: Metadata;
  data: {
    MemberId: string;
    TeamId: string;
    UserId: string;
    Role: string;
    Created: string;
  };
}
```

## Messaging

Messaging systems allow Microservices to exchange messages without coupling them together. Some Microservices emit messages, while others listen to the messages they subscribe to.

A message is a generic term for data that could be either a command or an event. Commands are messages that trigger something to happen (in the future). Events are messages that notify listeners about something that has happened (in the past). Publishers send commands or events without knowing the consumers that may be listening.

### Language Support

This package is written in TypeScript but converted to language targets. Each supported language has its own package.

Supported languages include:

- TypeScript [See package](https://www.npmjs.com/package/@devpie/client-events)
- Golang [See package](https://github.com/ivorscott/client-events/tree/main/go)

### Development

Modify `src/events.ts`, the source of truth, then re-build to update all packages.

```
npm run build
```

### Release

Here's the steps to perform a manual release for Typescript and Go packages (needs to be automated). Publishing Go modules relies on git tags. https://blog.golang.org/publishing-go-modules

```bash
# 1. npm run build
# 2. update ./package.json version
# 3. commit changes to git
# 4. create a new tag for release
git tag v0.0.1
# 5. push new tag
git push origin v0.0.1
# 6. push changes to remote repository
git push origin main
# 7. publish npm module
npm publish
```
