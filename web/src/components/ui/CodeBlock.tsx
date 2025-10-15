import { CopyButton } from "./CopyButton";

interface CodeBlockProps {
  content: string;
  filename?: string;
}

export function CodeBlock(props: CodeBlockProps) {
  return (
    <div
      class="rounded-xl border border-border overflow-hidden shadow-sm"
      style={{ background: "var(--code-bg)" }}
    >
      <div class="px-5 py-3 border-b border-border bg-muted/30 flex items-center justify-between">
        <span class="text-sm font-semibold text-foreground font-mono">
          {props.filename || "Code"}
        </span>
        <CopyButton text={props.content} />
      </div>
      <pre class="p-5 text-sm text-card-foreground overflow-x-auto font-mono leading-relaxed">
        {props.content}
      </pre>
    </div>
  );
}
