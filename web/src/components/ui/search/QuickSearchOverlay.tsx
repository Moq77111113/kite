import { type ParentProps } from "solid-js";

interface QuickSearchOverlayProps extends ParentProps {
  onClose: () => void;
}

export default function QuickSearchOverlay(props: QuickSearchOverlayProps) {
  return (
    <div
      class="fixed inset-0 bg-black/50 backdrop-blur-sm z-[100] flex items-start justify-center pt-[15vh] animate-in fade-in duration-200"
      onClick={props.onClose}
    >
      <div
        class="bg-card border border-border rounded-xl shadow-2xl w-full max-w-2xl mx-4 overflow-hidden animate-in slide-in-from-top-4 duration-200"
        onClick={(e) => e.stopPropagation()}
      >
        {props.children}
      </div>
    </div>
  );
}
