import { createSignal } from "solid-js";

const [isOnFire, setIsOnFire] = createSignal(false);

export const fireStore = {
  get burning() {
    return isOnFire();
  },

  boom() {
    setIsOnFire(true);
  },

  calm() {
    setIsOnFire(false);
  },
};
