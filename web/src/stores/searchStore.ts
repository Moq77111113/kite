import { createSignal } from "solid-js";

const [isOpen, setIsOpen] = createSignal(false);

export const searchStore = {
  get isOpen() {
    return isOpen();
  },

  open() {
    setIsOpen(true);
  },

  close() {
    setIsOpen(false);
  },

  toggle() {
    setIsOpen((prev) => !prev);
  },
};
