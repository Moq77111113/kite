# Contributing to Kite

Thanks for considering contributing to Kite. We need all the help we can get.

## Philosophy

Kite exists to make infrastructure code **forkable, editable, and yours**. If your contribution moves us away from that goal, we'll probably reject it. If it moves us toward that goal, you're a hero.

**What we care about:**

- Simplicity over features
- Clarity over cleverness
- Working code over perfect code
- Real-world use cases over hypothetical edge cases

**What we don't care about:**

- Enterprise buzzwords
- Your 47-step authentication flow
- Frameworks that solve problems nobody has

---

## Getting Started

### Prerequisites

- Go 1.24+ (because we live on the edge)
- pnpm (for the web UI)
- Git (obviously)
- Coffee (optional but recommended)

### Local Setup

```bash
# Clone the repo
git clone https://github.com/Moq77111113/kite.git
cd kite

# Build everything (Go binary + web UI)
make build

# Or just the Go binary
go build -o kite ./cmd/kite.go

# Run it
./kite --help
```

### Running the Web UI in Dev Mode

```bash
# Terminal 1: Build and run the Go server
make kite -- serve --port 8080

# Terminal 2: Run the web UI in dev mode
cd web
pnpm install
pnpm dev
```

The dev server will proxy API requests to the Go server.

---

## Making Changes

### Architecture (Yes, We Have One)

Kite follows a *slightly pretentious* hexagonal/clean architecture. Before you roll your eyes—it actually makes sense here. We have multiple interfaces (CLI + web UI) hitting the same business logic, so the separation helps. Probably.

```
kite/
├── cmd/                    # Main entry point (shocking)
├── internal/
│   ├── application/        # Use cases / business logic
│   │   ├── add/            # Adding kits to projects
│   │   ├── describe/       # Describing registry items
│   │   ├── init/           # Initializing projects
│   │   ├── list/           # Listing installed kits
│   │   ├── remove/         # Removing kits
│   │   └── update/         # Updating kits
│   ├── domain/             # Core domain models (the pure stuff)
│   │   ├── install/        # Installation domain logic
│   │   ├── models/         # Domain entities
│   │   ├── parse/          # Parsing logic
│   │   ├── port/           # Port interfaces (hexagonal buzzword)
│   │   ├── repo/           # Repository patterns
│   │   ├── scan/           # Scanning logic
│   │   └── types/          # Domain types
│   ├── infra/              # Infrastructure (the dirty bits)
│   │   ├── api/            # HTTP API handlers (for web UI)
│   │   ├── cli/            # CLI commands (cobra stuff)
│   │   ├── persistence/    # Config file handling (JSON hell)
│   │   └── storage/        # Git operations & registry clients
│   └── version/            # Version info (GoReleaser magic)
├── pkg/console/            # Pretty CLI output (colors!)
├── web/                    # SolidJS web UI (currently EMBEDDED, because MVP)
│   ├── dist/               # Built assets that get embedded in the Go binary
│   ├── src/                # Actual SolidJS code
│   └── embed.go            # The magic that shoves the web UI into the binary
└── docs/                   # Documentation (you're reading it)
```

**The Embedded Web Situation:**

Yeah, the web UI is currently embedded directly into the Go binary (`web/embed.go`). This is an *MVP decision*. Eventually™ we might separate it into a proper client-server architecture, but for now, having a single binary that contains both CLI and web UI is actually convenient. Don't @ me.

**Naming conventions:**

- CLI commands live in `internal/infra/cli/<command>/`
- Application use cases live in `internal/application/<command>/`
- Each command/use case has its own package
- Domain stuff stays pure (no infrastructure dependencies)
- Keep files focused—one responsibility per file (we're not monsters)

### What We Need

**High priority:**

- Tests (we have zero, it's embarrassing)
- Better error messages (current ones are... optimistic)
- Documentation improvements
- Refactoring messy code (looking at you, registry clients)

**Medium priority:**

- Kit variables/interpolation
- Search improvements in the web UI
- CI/CD integration (docker image, GitHub Actions)

**Low priority:**

- Analytics (who used what)
- Offline mode
- Features nobody asked for

---

## Pull Request Process

1. **Fork the repo** (you know how this works)

2. **Create a branch** with a clear name:

   ```bash
   git switch -c feat/your-feature-name
   ```

3. **Make your changes**

   - Write clear commit messages
   - Keep commits focused (one thing per commit)
   - Test locally (yes, even though we have no tests)

4. **Push and open a PR**

   - Describe what you changed and why
   - Reference any related issues
   - Include screenshots/GIFs if it's a UI change

5. **Wait for review**
   - We'll try to respond quickly
   - Don't take feedback personally—we're all learning here

### PR Guidelines

**Good PRs:**

- Solve a real problem
- Include context (why did you make this change?)
- Are small and focused
- Don't introduce 15 new dependencies

**Bad PRs:**

- "Refactor everything to use framework X"
- Add features without explaining why
- Break existing functionality
- Include commented-out code or debug prints

---

## Code Style

**Go:**

- Follow standard Go conventions (`gofmt` is your friend)
- Keep functions short and readable
- Error handling: wrap errors with context
- Comments: Try to make code self-explanatory; use comments for "why," not "what"

**TypeScript/SolidJS:**

- Use TypeScript (`any` guys should be ashamed)
- Follow the existing component structure
- Keep components small and composable

**General:**

- No commented-out code in commits
- No `console.log` debugging left in production code
- If you add a TODO, create an issue for it

---

## Testing

We don't have tests yet. This is not a flex—it's a cry for help.

If you want to be a hero, help us add:

- Unit tests for core logic
- Integration tests for registry clients
- E2E tests for CLI commands

Use Go's standard `testing` package. No fancy frameworks needed.

---

## Building for Release

Releases are automated via GoReleaser when a tag is pushed:

```bash
git tag v0.2.0
git push origin v0.2.0
```

GitHub Actions will:

1. Build the web UI
2. Run tests (when we have them)
3. Build binaries for Linux and macOS
4. Create a GitHub release with artifacts

---

## Documentation

If you're adding a new command or feature:

- Update the README if needed
- Add a doc in `docs/` if it's complex
- Include usage examples

Keep docs concise. Nobody reads walls of text.

---

## Getting Help

**Questions?**

- Open an issue with the "question" label
- Tag it clearly
- Be specific

**Found a bug?**

- Open an issue with steps to reproduce
- Include your OS, Go version, and Kite version (`kite version`)
- Bonus points for a minimal test case

**Want to propose a feature?**

- Open an issue first (before writing code)
- Explain the problem you're solving
- Describe your proposed solution
- Wait for feedback

---

## Code of Conduct

**The rules:**

1. Be respectful
2. Be constructive
3. Don't be an asshole

That's it. We're all here to build something useful.

---

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

## Final Notes

Kite is a side project built by people who got tired of infrastructure tooling complexity. We're not trying to build the next unicorn startup or create a foundation.

We just want a simple tool that does one thing well: **copy files from Git repos into your project.**

If that resonates with you, welcome aboard. Let's build something useful.

**Now go write some code.**
