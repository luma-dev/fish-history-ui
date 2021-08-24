import { Temporal } from '@js-temporal/polyfill';
import { Readable, writable } from 'svelte/store';

function createToday(): Readable<Temporal.PlainDate> {
  const { subscribe, set } = writable<Temporal.PlainDate>(Temporal.Now.plainDateISO('UTC'));

  void (async () => {
    set(Temporal.Now.plainDateISO('UTC'));
    await new Promise(resolve => setTimeout(resolve, 1000));
  })();

  return {
    subscribe,
  };
}

const today = createToday();
export default today;
