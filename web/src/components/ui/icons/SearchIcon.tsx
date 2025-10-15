import type { IconProps } from "./types";

export const SearchIcon = (props: IconProps) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={props.size || 18}
    height={props.size || 18}
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    class={props.class}
  >
    <circle cx="11" cy="11" r="8" />
    <path d="m21 21-4.35-4.35" />
  </svg>
);
