# Why Use Git Hooks

This project demonstrates how to use **Git Hooks** in a Go project to automate tasks like code formatting, running tests, linting, and enforcing commit message conventions.

Git hooks help **ensure code quality before it even reaches the repository**.

---

## Why use Git Hooks?

* [x] Automate repetitive tasks
* [x] Prevent broken commits (tests and lint checks run before committing)
* [x] Enforce commit message standards
* [x] Prevent pushing code that fails tests
* [x] Improve collaboration and maintain high code quality

---

## Project Structure

```
why-use-git-hoocks/
├── .git/
│   └── hooks/
│       ├── commit-msg.sample
│       ├── pre-commit.sample
│       ├── pre-push.sample
├── main.go
├── transaction/
│   ├── transaction.go
│   ├── transaction_test.go
└── README.md
```

---

## Available Hooks

### 1. `pre-commit`

Runs before a commit is created.
In this project, it:

* Formats the code (`go fmt`)
* Runs the linter (`golangci-lint`)
* Executes all tests (`go test`)

If any step fails, the commit is blocked.

---

### 2. `commit-msg`

Runs after you type your commit message but before the commit is finalized.
Validates that the commit message follows the [Conventional Commits](https://www.conventionalcommits.org/) standard.

**Expected format:**

```
feat: add parallel transaction processing
fix: fix concurrency bug in ProcessTransactions
```

If the message does not follow the standard, the commit will be rejected.

---

### 3. `pre-push`

Runs before pushing to a remote repository (`git push`).
Executes the tests to prevent pushing broken code.

## How to Set Up the Hooks

1. Copy the `.sample` files into the `.git/hooks` directory and remove the `.sample` extension:

   ```bash
   cp commit-msg.sample .git/hooks/commit-msg
   cp pre-commit.sample .git/hooks/pre-commit
   cp pre-push.sample .git/hooks/pre-push
   ```

2. Make them executable:

   ```bash
   chmod +x .git/hooks/commit-msg
   chmod +x .git/hooks/pre-commit
   chmod +x .git/hooks/pre-push
   ```

3. Test the hooks:

   ```bash
   git commit -m "feat: example commit"
   git push
   ```

---

## Recommended Dependencies

* [Go](https://go.dev/) (>= 1.20)
* [GolangCI-Lint](https://golangci-lint.run/):

  ```bash
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  ```

---

## Example Workflow

1. You modify your code and run `git add .`
2. When you run `git commit`, the `pre-commit` hook:

   * Formats the code
   * Lints the code
   * Runs tests
3. The `commit-msg` hook validates the commit message.
4. When you run `git push`, the `pre-push` hook runs the tests again.

If everything passes, the code is safely pushed.

---