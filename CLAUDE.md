# Cherrypick Test Project

## Code Conventions
- Use camelCase for variables and functions
- Use PascalCase for classes and interfaces
- All functions must have JSDoc comments
- Prefer functional array methods (`Array.map()`, `Array.filter()`) over manual `for` loops with `push()`

## TypeScript
- Never use `any` type — use `unknown` and validate, or define a proper interface (e.g., `UserData`)
- `any` disables type checking entirely and hides real bugs

## Error Handling
- Never silently swallow errors with `console.log` + `return null`
- Either let errors propagate naturally or throw a typed error; silent failures make debugging impossible in production

## Security
- Never hardcode secrets in source code — use environment variables (`process.env.SECRET_KEY`) or a secrets manager

## Testing
- Use vitest for unit tests
- Tests go in `tests/` directory
