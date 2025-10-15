import { createSignal, onMount, onCleanup, For } from "solid-js";

interface Section {
  id: string;
  label: string;
}

interface SectionNavProps {
  sections: Section[];
}

export function SectionNav(props: SectionNavProps) {
  const [activeSection, setActiveSection] = createSignal("");

  onMount(() => {
    const observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            setActiveSection(entry.target.id);
          }
        });
      },
      {
        rootMargin: "-100px 0px -66%",
        threshold: 0,
      }
    );

    props.sections.forEach((section) => {
      const element = document.getElementById(section.id);
      if (element) {
        observer.observe(element);
      }
    });

    onCleanup(() => observer.disconnect());
  });

  const scrollToSection = (id: string) => {
    const element = document.getElementById(id);
    if (element) {
      element.scrollIntoView({ behavior: "smooth", block: "start" });
    }
  };

  return (
    <nav class="sticky top-24 space-y-2">
      <div class="text-xs font-semibold text-muted-foreground mb-3 uppercase tracking-wider">
        On this page
      </div>
      <div class="relative">
        <div class="absolute left-0 top-0 bottom-0 w-px bg-border" />
        <ul class="space-y-2">
          <For each={props.sections}>
            {(section) => (
              <li class="relative">
                <button
                  onClick={() => scrollToSection(section.id)}
                  class={`block pl-3 text-sm transition-colors text-left w-full ${
                    activeSection() === section.id
                      ? "text-foreground font-medium"
                      : "text-muted-foreground hover:text-foreground"
                  }`}
                >
                  <div
                    class={`absolute left-0 top-1/2 -translate-y-1/2 w-px transition-all ${
                      activeSection() === section.id
                        ? "h-full bg-foreground"
                        : "h-0 bg-border"
                    }`}
                  />
                  {section.label}
                </button>
              </li>
            )}
          </For>
        </ul>
      </div>
    </nav>
  );
}
