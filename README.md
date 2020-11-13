# Common Module For Events

This module supplies the event types and event definitions for Devpie's CQRS System Architecture.

The same events are specified in multiple languages. Supported languages include:

- TypeScript
- Golang
- and Python

The base definition is defined in Typescript under `events.ts` and converted into other languages using [Quicktype](https://quicktype.io/).

## Usage

To build code for both Python and Go run: `npm run build`

![cqrs architecture](cqrs.png)
