import Prism from 'prismjs';
import 'prismjs/components/prism-bash';
import 'prismjs/components/prism-json';
import 'prismjs/components/prism-jsx';
import 'prismjs/components/prism-markup';
import 'prismjs/components/prism-tsx';
import 'prismjs/components/prism-typescript';
import 'prismjs/components/prism-yaml';
import 'prismjs/themes/prism-tomorrow.css';
import { createEffect } from 'solid-js';
import { CopyButton } from './CopyButton';
import { MarkdownRenderer } from './Markdown';
interface CodeBlockProps {
  content: string;
  filename?: string;
  language?: string;
}

export function CodeBlock(props: CodeBlockProps) {
  let codeRef: HTMLElement | undefined;
  const lang = () => detectLanguage(props.filename, props.language);
  const isMarkdown = lang() === 'markdown';

  createEffect(() => {
    if (!codeRef || isMarkdown) return;
    codeRef.className = `language-${lang()}`;
    codeRef.textContent = props.content;
    try {
      Prism.highlightElement(codeRef);
    } catch (err) {
      console.warn('Highlighting failed:', err);
    }
  });

  return (
    <div class="rounded-xl border border-border overflow-hidden shadow-sm bg-[hsl(var(--code-bg))]">
      <div class="px-5 py-3 border-b border-border bg-muted/30 flex items-center justify-between">
        <span class="text-sm font-semibold text-foreground font-mono">
          {props.filename || 'Code'}
        </span>
        <CopyButton text={props.content} />
      </div>
      {isMarkdown ? (
        <div class="p-5">
          <MarkdownRenderer markdown={props.content} />
        </div>
      ) : (
        <pre class="p-5 text-sm overflow-x-auto font-mono leading-relaxed">
          <code ref={codeRef} class={`language-${lang()}`}>
            {props.content}
          </code>
        </pre>
      )}
    </div>
  );
}
// ...existing code...

function detectLanguage(filename?: string, explicit?: string): string {
  if (explicit) return explicit;
  if (!filename) return 'plaintext';

  const ext = filename.split('.').pop()?.toLowerCase();

  switch (ext) {
    case 'ts':
      return 'typescript';
    case 'tsx':
      return 'tsx';
    case 'js':
      return 'javascript';
    case 'jsx':
      return 'jsx';
    case 'html':
    case 'xml':
      return 'markup';
    case 'json':
      return 'json';
    case 'css':
      return 'css';
    case 'yml':
    case 'yaml':
      return 'yaml';
    case 'sh':
      return 'bash';
    case 'md':
    case 'mdx':
      return 'markdown';
    default:
      return ext || 'plaintext';
  }
}
