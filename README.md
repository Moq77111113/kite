# 🪁 KITE

**Fork your infrastructure. Don't worship it.**

> [!WARNING]
> Still an MVP. Because the best tools start as scripts that worked once.

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)
[![Go Version](https://img.shields.io/github/go-mod/go-version/Moq77111113/kite)](go.mod)
[![Release](https://img.shields.io/github/v/release/Moq77111113/kite)](https://github.com/Moq77111113/kite/releases)
[![Stars](https://img.shields.io/github/stars/Moq77111113/kite?style=social)](https://github.com/Moq77111113/kite/stargazers)

**Try it now:**

```bash
curl -sfL https://raw.githubusercontent.com/Moq77111113/kite/main/install.sh | sh && kite init
```

---

## What is this?

> Kite is not a package manager. It's a code forking tool.

A CLI tool for copying infrastructure code without the guilt trip.

Docker setups, app configs. Terraform modules. The stuff that lives in your team's repos but nobody can find because it's Slack message 47 on June 12th.

**You don't install templates. You copy them. Then you break them. Like adults.**

## Why does this exist?

Every team has the same ritual:

1. Someone writes a killer Docker setup
2. It gets buried in Slack / a random repo / Steve's brain
3. You find it 6 months later, copy-paste it, ship it
4. Three months pass
5. Nobody remembers where it came from or how it works

**Welcome to infrastructure archaeology.**

Sure, "proper" solutions exist. Package managers. Git submodules. Registry systems. Terraform modules that do 80% of what you need. They all meant well.
But here's what actually happens: you copy Steve's folder, rename a few variables, and ship it before standup.

**Kite doesn’t fix it — it embraces it, and gives you a shovel.**

## How it works

<p align="center">
  <img src="demo.gif" alt="Kite Demo" width="800">
</p>

```bash
# initialize a new registry (or use an existing one)
kite init

# Browse kits in a wholesome web UI
kite serve

# Or just grab what you need
kite add docker-postgres redis-cache

# Files appear. Edit them. They're yours now.
```

> For those who need a little hand-holding despite the rules, here's an example registry: [https://github.com/Moq77111113/kite-registry](https://github.com/Moq77111113/kite-registry).

**That's it. That's the whole tool.**

No daemons. No agents. No "syncing with the mothership."

## What you get

```bash
kite add steves-stuff redis-cache
```

**Result:**

```
├─ redis-cache/
│  └─ docker-compose.yml
└─ steves-stuff/
   └─ fix-everything.sh
```

Real files. Not symlinks. Not "immutable dependencies." Not "please don't touch this."

**Code you can read, edit, and completely destroy at 3 AM if you want.**

## Yes, you could...

- Sure, you could use a package manager.
- You could even `git sparse-checkout` that one folder you actually need.
- You could add a submodule too — if you enjoy existential pain.
- And yes, you could even build your own shadcn registry.

And it's ok!

But sometimes you just want to grab `cleanRestart.sh`, or `steves-business.py` — drop them in your repo, tweak two lines, and move on with your life.

That’s what Kite is for.
No build step. No dependency graph. No ceremony.
Just a bunch of scripts that actually do things.

```
npm install  →  Black box, breaks on update, vendor lock-in
kite add     →  Fork the code, modify it, you're in control
```

## The Rules

1. **No vendor lock-in** — Git repos, not proprietary registries
2. **No magic** — It's `cp -r` with a web UI
3. **No worship** — Edit the code. It's yours. Break it on purpose.
4. **No hand-holding** — If you can't `vim` it, you don't need it
5. **No documentation guilt** — if it works, you’ll document it tomorrow (maybe)

Kite assumes you're an adult who understands that forking code is how things actually get done.

## Setup

**Install:**

```bash
curl -sfL https://raw.githubusercontent.com/Moq77111113/kite/main/install.sh | sh
```

**Use:**

```bash
# Point to a registry (just a Git repo)
kite init --registry git@github.com:your-team/kits.git

# Browse kits in a UI
kite serve --port 8080

# Or skip the browsing
kite list

# Copy what you need
kite add docker-postgres github-ci-deploy
```

## What it does

✅ Browse kits in a wholesome UI (because command lines are for cowards)  
✅ Copy them as real files you can actually edit  
✅ Git-based registries (your repo = your source of truth)  
✅ Self-hosted (no SaaS, no tracking, no "phone home")  
✅ Works with any text files (Docker, Terraform, CI, shell scripts, Steve's notes, whatever)

## What it's NOT

❌ A package manager pretending to solve your problems  
❌ "Enterprise ready" (aka bloated and slow)  
❌ Trying to abstract away complexity you need to understand  
❌ Here to hold your hand while you cargo-cult best practices

It’s `cp -r` with opinions and a search bar.
Everything else is coping.

## Creating a registry

It's just a Git repo with folders:

```
my-kits/
├── docker-postgres/
│   ├── kite.yaml
│   └── docker-compose.yml
├── github-ci-node/
│   ├── kite.yaml
│   └── .github/workflows/ci.yml
└── steves-script-for-magically-fixing-stuff/
    ├── kite.yaml
    └── fix-everything.sh
```

Each kit has a `kite.yaml`:

```yaml
name: docker-postgres
version: 1.0.0
description: PostgreSQL with Docker Compose
tags: [docker, database]
```

Push to Git. Point Kite at it. Ship.

### What people are saying

> "Finally, a tool that admits we all copy-paste code. This is how we actually work."
> — @DevOpsSteve

> "I've been doing this manually for years. Kite just made it 10x faster."
> — @SarahFromSlack

> "Package managers promised to solve this. They lied. Kite delivers."
> — @InfraArchaeologist

## Roadmap

**V0** (now): CLI + web UI + Git registries
**V1**: Template variables, better search, CI integration
**V2**: Analytics, offline mode, maybe a t-shirt

And more, if you ask nicely.

## Contributing

PRs welcome. Issue reports welcome. Philosophical debates about package management welcome.

What we need:

- Tests (of course)
- Docs
- Fixing my terrible Go code
- Error messages that tell the truth
- Coffee funds

## License

MIT. Copy it. Fork it. Sell it to Oracle. Whatever.

## Credits

**shadcn/ui** — proved that copy-paste is a legitimate design pattern  
**Every over-engineered tool ever** — endless motivation  
**YAML** — a constant reminder that pain is the price of simplicity  
**Caffeine and Stack Overflow's copy button** — powering infrastructure since 2008  
**LLMs** — for making me realize how much I miss writing READMEs

**It's what happens when you stop pretending you don't copy code.**

If you've ever searched Slack for 30 minutes looking for "that one docker-compose file Sarah wrote," this tool is for you.

```bash
curl -sfL https://raw.githubusercontent.com/Moq77111113/kite/main/install.sh | sh && kite init
```
