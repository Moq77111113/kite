# 🪁 KITE

> Fork your infrastructure. Don't worship it.

> [!WARNING]  
> It’s still an MVP.
> Because the best tools start as scripts that worked once — and never stopped.

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.24-00ADD8.svg)](https://go.dev/)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)

### What is Kite?

Kite is a CLI tool for copying, remixing, and actually owning your infrastructure code. Terraform modules, CI configs, scripts, Kubernetes manifests—if it's code, Kite lets you fork it, break it, and make it yours. No package manager nonsense, no magical servers, no dependency drama.

Why? Because infrastructure shouldn't be downloaded like a plugin—it should be **owned**, **read**, **edited**, and **broken** by the people who run it. If you want a black box, go install something else. If you want to actually understand your stack, Kite is your new best friend.

## Philosophy

Every dev has done it.

You find a perfect setup on GitHub, copy-paste it, tell yourself you'll clean it later.

Six months later it's a shrine of YAML files nobody understands.

The person who wrote it is gone.

The pipeline still works — somehow — so no one touches it.

Welcome to DevOps archaeology.

Kite exists because that's **insane**. If you can `git clone`, you don't need another black box. You need a way to fork, own, and evolve code without pretending it's sacred.

## The Tool

Kite copies real files into your repo. Not links. Not "dependencies". Real. Editable. Breakable. Yours.
**Example:**

```bash
kite add docker-postgres
# Instantly, your infra grows a brain.
```

```
your-project/
infrastructure/
├─ docker-postgres/
│  ├─ docker-compose.yml
│  └─ .env.example
└─ github-actions/
	└─ .github/workflows/deploy.yml
```

Edit it, mess it up, fix it, commit it. That's ownership.

And yes — v1 will make Kite **CI/CD-ready**.  
Your pipelines can now `kite add --ci` templates automatically, without any human in the loop.  
GitHub Actions, GitLab CI, any runner: your infrastructure is cloned, editable, and ready **as part of your build**.

## The Rules

- No vendor lock-in
- No magical registries
- No "please don't touch this file" bullshit
- If you can't open it in Vim, it doesn't belong in your stack

> Kite doesn't try to protect you from your own mistakes. It assumes you're smart enough to make them on purpose.

## Setup

**Install:**

```bash
curl -ssL https://raw.githubusercontent.com/moq77111113/kite/main/install.sh | bash
# One day we'll have a website.
```

or clone, build, install. Whatever makes you feel alive.

**Init your registry:**

```bash
kite init --registry git@github.com:your-org/templates.git
```

**Add some templates:**

```bash
kite add docker-postgres github-actions
```

**Serve it yourself:**

```bash
kite serve --registry git@github.com:your-org/templates.git --port 8080
```

Includes a Web UI. You're welcome.

## Why It Exists

- Because every "modern" DevOps tool is trying to turn engineers into consumers
- Because people forgot that "open source" means _own the damn code_
- Because you don't need YAML-as-a-service to spin up a database

Kite is a middle finger to dependency culture. You don't "install" your infrastructure — you **fork** it.

## What It's Not

- Not a package manager
- Doesn't hold your hand
- Doesn't automatically fix your mistakes
- Not SaaS
- Not "enterprise ready"

It's a glorified `cp -r` with opinions. And that's exactly why it works.

## Roadmap

- **V0:** Git registries, HTTP registries via kite serve, Web UI, basic commands
- **V1:** Template variables, search, sane errors, CI/CD integration (dedicated commands, docker image)
- **V2:** Analytics, offline mode, maybe a t-shirt

## Contributing

PRs welcome. We promise not to create a foundation.

We need:

- Tests
- Templates
- Error messages that don't gaslight users
- More caffeine

## License

MIT. Copy it. Fork it. Rename it. We literally don't care.

> If you've ever copy-pasted a docker-compose.yml at 2 AM, you already understand why Kite exists. Welcome home.

## Credits

- **shadcn/ui** — for proving that copy-paste can be a design principle.
- **Every over-engineered DevOps tool** — thank you for the motivation.
- **The inventor of YAML** — a daily reminder that pain is a teacher.
- **Caffeine, sarcasm, and the Stack Overflow copy button** — the holy trinity.
- **LLMs** — for helping humanity automate its own redundancy.

> You don’t need another platform. You just need to read your damn code.
