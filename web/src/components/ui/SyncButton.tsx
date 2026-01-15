import { createSignal } from "solid-js";
import { RefreshIcon } from "./icons";
import { syncRegistry } from "@/api/kits";

interface SyncButtonProps {
  lastSync?: string;
  onSyncComplete?: (lastSync?: string) => void;
}

function formatTimeAgo(dateString?: string): string {
  if (!dateString) return "Never";

  const date = new Date(dateString);
  const now = new Date();
  const diffMs = now.getTime() - date.getTime();
  const diffMins = Math.floor(diffMs / 60000);

  if (diffMins < 1) return "Just now";
  if (diffMins < 60) return `${diffMins}m ago`;

  const diffHours = Math.floor(diffMins / 60);
  if (diffHours < 24) return `${diffHours}h ago`;

  const diffDays = Math.floor(diffHours / 24);
  return `${diffDays}d ago`;
}

export default function SyncButton(props: SyncButtonProps) {
  const [syncing, setSyncing] = createSignal(false);

  const handleSync = async () => {
    if (syncing()) return;

    setSyncing(true);
    try {
      const result = await syncRegistry();
      props.onSyncComplete?.(result.lastSync);
    } catch (error) {
      console.error("Sync failed:", error);
    } finally {
      setSyncing(false);
    }
  };

  return (
    <button
      onClick={handleSync}
      disabled={syncing()}
      class="p-2 cursor-pointer rounded-md hover:bg-sidebar-accent text-sidebar-foreground transition-colors disabled:opacity-50"
      aria-label="Sync registry"
      title={`Last synced: ${formatTimeAgo(props.lastSync)}`}
    >
      <RefreshIcon class={syncing() ? "animate-spin" : ""} />
    </button>
  );
}
