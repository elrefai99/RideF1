# Repository Guidelines

## Project Structure & Module Organization

The repository is at an early stage. `app/` is the intended home for application source, and `docs/` is for project documentation; both are currently empty. Root-level files contain licensing, authorship, code ownership, and repository configuration. Keep new code and assets in the relevant directory rather than adding unrelated files at the root.

## Build, Test, and Development Commands

No build system, package manifest, or test runner is present yet, so there are no project commands to run. When adding a language or framework, include its setup and standard build, test, and local development commands in the project documentation and keep dependencies declared in its manifest.

## Coding Style & Naming Conventions

No language-specific formatter or lint configuration exists yet. Follow the conventions of the chosen language and framework, use consistent formatting, and add formatter/linter configuration alongside the first application code. Prefer descriptive, lowercase directory and file names; follow the ecosystem's standard naming for modules and tests.

## Testing Guidelines

There is no test suite or coverage requirement yet. Add tests with new behavior, keep them near the structure prescribed by the chosen framework, and document the command to run them. Use names that describe the behavior under test (for example, `user_settings_test` where that pattern fits the language).

## Commit & Pull Request Guidelines

The Git history currently contains only the baseline commit, so no established commit-message convention can be inferred. Use short, imperative commit subjects (for example, `Add settings screen`). Pull requests should explain the change and its motivation, mention relevant issues, and include screenshots for visible UI changes. Note any test or build commands run; if none are available, say so.

## Security & Configuration

Do not commit credentials or local environment files. `.env` files are ignored; provide safe example configuration with placeholder values if the application needs environment variables. Keep generated output, logs, and coverage artifacts out of version control as configured in `.gitignore`.


## Permissions

Global rule:

- Ask the user first before making any code change.
- Show the intended change for review when possible.
- Wait for the user to accept or reject the change before editing project code.
- Always keep the changes in the code to be as simple as possible, straight to the point, with no comments added

### Allow

- Read tracked project files needed for the task.
- Read source code under `src/`, configuration under `config/`, and docs such as `README.md`, `CLAUDE.md`, and this file.
- Create new source or documentation files when they are required for the requested change.
- Edit application code, route files, models, middleware, utilities, tests, and markdown documentation.
- Update `package.json` when the task explicitly requires script or dependency changes.
- Run safe repo-local commands such as `rg`, `ls`, `sed`, `git status`, `pnpm lint`, and `pnpm test`.

### Deny

- Do not read, print, or copy secrets from `.env`, `.env.dev`, or any credential file.
- Do not modify `.env`, `.env.dev`, or other secret-bearing files unless the user explicitly asks.
- Do not modify `node_modules/`, generated caches, or log files.
- Do not change `pnpm-lock.yaml` unless dependency work is part of the task.
- Do not delete files, rename major directories, or rewrite large parts of the codebase without explicit approval.
- Do not run destructive git or shell commands such as `git reset --hard`, `git checkout --`, or broad `rm` operations.
- Do not alter deployment/infrastructure files (`Dockerfile`, `docker-compose.yml`, `ecosystem.config.cjs`) unless the task explicitly requires it.
