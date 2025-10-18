import { marked } from 'marked';
import { createMemo } from 'solid-js';

export function MarkdownRenderer({ markdown }: { markdown: string }) {
  const rendered = createMemo(() => {
    if (!markdown) return '';
    return marked.parse(markdown, { async: false });
  });
  return (
    <div
      class="prose prose-slate dark:prose-invert max-w-none"
      innerHTML={rendered()}
    />
  );
}
