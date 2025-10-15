import "@/styles/fire.css";
import { For, Show } from "solid-js";

interface FireOverlayProps {
  isActive: boolean;
}

export default function FireOverlay(props: FireOverlayProps) {
  const debrisParticles = Array.from({ length: 40 }, (_, i) => ({
    id: i,
    left: `${10 + ((i * 2) % 80)}%`,
    bottom: `${15 + ((i * 3) % 60)}%`,
    delay: `${(i * 0.05) % 1}s`,
    animation: `debrisFly${(i % 4) + 1} ${
      0.8 + (i % 5) * 0.1
    }s ease-out infinite`,
  }));

  const shockwaves = Array.from({ length: 8 }, (_, i) => ({
    id: i,
    left: `${15 + ((i * 12) % 70)}%`,
    bottom: `${20 + ((i * 8) % 40)}%`,
    delay: `${(i * 0.12) % 1}s`,
  }));

  return (
    <Show when={props.isActive}>
      <div class="fixed inset-0 pointer-events-none z-[200] overflow-hidden">
        <div class="explosion-container">
          <div class="flash-overlay" />

          <div class="shell shell-1" />
          <div class="shell shell-2" />
          <div class="shell shell-3" />
          <div class="shell shell-4" />
          <div class="shell shell-5" />

          <For each={shockwaves}>
            {(wave) => (
              <div
                class="shockwave"
                style={{
                  left: wave.left,
                  bottom: wave.bottom,
                  animation: `shockwave 0.6s ease-out infinite`,
                  "animation-delay": wave.delay,
                }}
              />
            )}
          </For>

          <div class="explosion explosion-1" />
          <div class="explosion explosion-2" />
          <div class="explosion explosion-3" />
          <div class="explosion explosion-4" />
          <div class="explosion explosion-5" />
          <div class="explosion explosion-6" />
          <div class="explosion explosion-7" />
          <div class="explosion explosion-8" />

          <For each={debrisParticles}>
            {(particle) => (
              <div
                class="debris"
                style={{
                  left: particle.left,
                  bottom: particle.bottom,
                  animation: particle.animation,
                  "animation-delay": particle.delay,
                }}
              />
            )}
          </For>

          <div class="smoke smoke-1" />
          <div class="smoke smoke-2" />
          <div class="smoke smoke-3" />
        </div>
      </div>
    </Show>
  );
}
