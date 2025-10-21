Kite — web UI

Yes, there's a web UI. It's tiny, opinionated, and embedded into the Go binary so you can ship one executable and still pretend this is "modern." 

Quick notes:

- Built with Solid + Vite. Static assets are compiled and go magik's `embed` package is used to bake them into the Go binary.
- Run `kite serve` to start the Go server that serves the embedded UI (no Node process required — you're welcome).
- Tweak `web/src` and run your usual Vite build to regenerate assets if you feel like changing the thing you weren't supposed to touch.
- Vite dev server will proxy API requests to the Go backend for local development.

Philosophy (TL;DR): small JS app + Go glue = no drama. Copy it. Break it. Tell no one.

You're done here. (you know how to run a vite app, right?)
