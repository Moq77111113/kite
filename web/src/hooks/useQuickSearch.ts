import { searchStore } from '@/stores/searchStore';
import type { KitSummary } from '@/types/kit';
import { useNavigate } from '@tanstack/solid-router';
import { createEffect, createSignal, onCleanup } from 'solid-js';

export const useQuickSearch = (kits: KitSummary[]) => {
  const [query, setQuery] = createSignal('');
  const [selectedIndex, setSelectedIndex] = createSignal(0);
  const navigate = useNavigate();
  let inputRef: HTMLInputElement | undefined;

  const filteredKits = () => {
    const q = query().toLowerCase();
    if (!q) return kits.slice(0, 8);

    return kits
      .filter(
        (t) =>
          t.name.toLowerCase().includes(q) ||
          t.description?.toLowerCase().includes(q) ||
          t.tags.some((tag) => tag.toLowerCase().includes(q))
      )
      .slice(0, 8);
  };

  const selectKit = (kit: KitSummary) => {
    navigate({ to: '/kits/$name', params: { name: kit.name } });
    searchStore.close();
  };

  createEffect(() => {
    if (!searchStore.isOpen) return;

    const handleKeyDown = (e: KeyboardEvent) => {
      const handlers: Record<string, () => void> = {
        Escape: () => {
          e.preventDefault();
          searchStore.close();
        },
        ArrowDown: () => {
          e.preventDefault();
          setSelectedIndex((i) => Math.min(i + 1, filteredKits().length - 1));
        },
        ArrowUp: () => {
          e.preventDefault();
          setSelectedIndex((i) => Math.max(i - 1, 0));
        },
        Enter: () => {
          e.preventDefault();
          const selected = filteredKits()[selectedIndex()];
          if (selected) selectKit(selected);
        },
      };

      handlers[e.key]?.();
    };

    document.addEventListener('keydown', handleKeyDown);
    onCleanup(() => document.removeEventListener('keydown', handleKeyDown));
  });

  createEffect(() => {
    if (query()) setSelectedIndex(0);
  });

  createEffect(() => {
    if (searchStore.isOpen) {
      setQuery('');
      setSelectedIndex(0);
      setTimeout(() => inputRef?.focus(), 0);
    }
  });

  return {
    query,
    setQuery,
    selectedIndex,
    filteredKits,
    selectKit,
    inputRef: (ref: HTMLInputElement) => (inputRef = ref),
  };
};
