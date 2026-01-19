import { For, Show, createSignal, onCleanup } from "solid-js";

interface TagsCarouselProps {
  tags: string[];
  maxVisible?: number;
}

export default function TagsCarousel(props: TagsCarouselProps) {
  const maxVisible = () => props.maxVisible ?? 3;

  let tagsContainerRef: HTMLDivElement | undefined;
  let scrollInterval: number | undefined;
  let leaveTimeout: number | undefined;
  const [isHovering, setIsHovering] = createSignal(false);

  const handleMouseMove = (e: MouseEvent) => {
    if (!tagsContainerRef || !isHovering()) return;

    const rect = tagsContainerRef.getBoundingClientRect();
    const x = e.clientX - rect.left;
    const width = rect.width;
    const edgeThreshold = 50;

    if (scrollInterval) {
      clearInterval(scrollInterval);
      scrollInterval = undefined;
    }

    if (x < edgeThreshold) {
      const speed = Math.max(1, (edgeThreshold - x) / 5);
      scrollInterval = setInterval(() => {
        if (tagsContainerRef) {
          tagsContainerRef.scrollLeft -= speed;
        }
      }, 16);
    } else if (x > width - edgeThreshold) {
      const speed = Math.max(1, (x - (width - edgeThreshold)) / 5);
      scrollInterval = setInterval(() => {
        if (tagsContainerRef) {
          tagsContainerRef.scrollLeft += speed;
        }
      }, 16);
    }
  };

  const handleMouseLeave = () => {
    if (scrollInterval) {
      clearInterval(scrollInterval);
      scrollInterval = undefined;
    }

    if (tagsContainerRef) {
      tagsContainerRef.scrollTo({
        left: 0,
        behavior: "smooth",
      });

      leaveTimeout = setTimeout(() => {
        setIsHovering(false);
      }, 300);
    } else {
      setIsHovering(false);
    }
  };

  const handleMouseEnter = () => {
    if (leaveTimeout) {
      clearTimeout(leaveTimeout);
      leaveTimeout = undefined;
    }
    setIsHovering(true);
  };

  onCleanup(() => {
    if (scrollInterval) {
      clearInterval(scrollInterval);
    }
    if (leaveTimeout) {
      clearTimeout(leaveTimeout);
    }
  });

  return (
    <div class="flex-1 min-w-0 overflow-hidden relative">
      <div
        ref={tagsContainerRef}
        onMouseEnter={handleMouseEnter}
        onMouseMove={handleMouseMove}
        onMouseLeave={handleMouseLeave}
        class="flex items-center gap-1.5 overflow-x-auto scrollbar-hide"
        style={{ "scrollbar-width": "none" }}
      >
        <For each={props.tags}>
          {(tag, index) => (
            <span
              class="px-2.5 py-1 text-xs rounded-md bg-muted/50 text-muted-foreground capitalize font-medium whitespace-nowrap flex-shrink-0"
              classList={{
                hidden: index() >= maxVisible() && !isHovering(),
              }}
            >
              {tag}
            </span>
          )}
        </For>
        <Show when={props.tags.length > maxVisible()}>
          <span
            class="px-2 py-1 text-xs text-muted-foreground font-medium shrink-0"
            classList={{
              hidden: isHovering(),
            }}
          >
            +{props.tags.length - maxVisible()}
          </span>
        </Show>
      </div>
    </div>
  );
}
