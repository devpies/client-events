# Common Module for Event-driven Architecture

Common module for event-driven architecture. Designed for Command Query Responsibility Segregation (CQRS).

The TypeScript code, under `src`, contains the source of truth. [Quicktype](https://quicktype.io/) is used to convert Typescript source code to specific language targets. Each supported language has its own package.

Supported languages include:

- TypeScript
- Golang
- Python [See package](https://pypi.org/project/devpie-client-events/)

## Development

If you modify `src/index.ts`, you need to re-build to update the packages.

```
npm run build
```

![cqrs architecture](cqrs.png)
