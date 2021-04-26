# Contract

## System messages are either commands or events

- Commands are things to be done. (in the future)
- Events are things that have happened. (in the past)

- All messages have an id, a subject, metadata and data

## Event sourcing stream naming conventions

Category streams

    {Category}

Command Category streams

    {Category}.command

Entity streams

    {Category}.123

Command streams

    {Category}.command.123
