import Check from "lucide-solid/icons/check";
import Copy from "lucide-solid/icons/copy";
import { createSignal, Show } from "solid-js";

interface CopyButtonProps {
  text: string;
  class?: string;
}

export function CopyButton(props: CopyButtonProps) {
  const [copied, setCopied] = createSignal(false);

  const handleCopy = async () => {
    try {
      await navigator.clipboard.writeText(props.text);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    } catch (err) {
      console.error("Failed to copy:", err);
    }
  };

  return (
    <button
      onClick={handleCopy}
      class={`p-2 rounded-md border border-border bg-secondary hover:bg-accent transition-colors ${
        props.class || ""
      }`}
      title="Copy to clipboard"
    >
      <Show when={copied()} fallback={<Copy size={16} />}>
        <Check size={16} class="text-green-500" />
      </Show>
    </button>
  );
}
