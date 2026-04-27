# Cherrypick Test Project

## Code Conventions
- Use camelCase for variables, functions, and object property names
- Use PascalCase for classes and interfaces
- All functions must have JSDoc comments
- Avoid `any` type — use `unknown` and validate, or define a proper interface
- Prefer `Array.map()` over manual `for` loops with `push()`
- Do not silently swallow errors with `console.log` + `return null`; let errors propagate or throw a typed error

## Security
- Never hardcode secrets in source code; use environment variables (`process.env.SECRET_KEY`) or a secrets manager

## Testing
- Use vitest for unit tests
- Tests go in `tests/` directory
