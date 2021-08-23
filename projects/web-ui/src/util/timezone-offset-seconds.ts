import { writable, Readable } from 'svelte/store';
import api from './api';

function createTimezoneOffsetSeconds(): Readable<number | null> {
  const { subscribe, set } = writable<number | null>(null);
  void api.timezone.$get().then(offset => {
    set(offset.offset);
  });

  return {
    subscribe,
  };
}

const timezoneOffsetSeconds = createTimezoneOffsetSeconds();

export default timezoneOffsetSeconds;
