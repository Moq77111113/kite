import { CopyButton } from "./CopyButton";

interface CodeBlockProps {
  content: string;
  filename?: string;
}

export function CodeBlock(props: CodeBlockProps) {
  return (
    <div class="rounded-lg border border-border bg-card overflow-hidden">
      <div class="px-4 py-2 border-b border-border bg-secondary flex items-center justify-between">
        <span class="text-sm font-medium text-secondary-foreground">
          {props.filename || "Code"}
        </span>
        <CopyButton text={props.content} />
      </div>
      <pre class="p-4 text-sm text-card-foreground overflow-x-auto">
        {props.content}
      </pre>
    </div>
  );
}
