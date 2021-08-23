import { locale } from 'svelte-i18n';
import type { Subscriber, Writable } from 'svelte/store';

export const lsLocaleKey = 'locale';

const createLSLocale = (): Writable<string> => {
  let run: Subscriber<string>;
  return {
    subscribe(run0) {
      run = run0;
      run(localStorage[lsLocaleKey]);
      // eslint-disable-next-line @typescript-eslint/no-empty-function
      return () => {};
    },
    set(value) {
      localStorage.setItem(lsLocaleKey, value);
      locale.set(value);
      run(value);
    },
    update(updater) {
      const value = localStorage.getItem(lsLocaleKey);
      const newValue = updater(value ?? '');
      localStorage.setItem(lsLocaleKey, newValue);
      locale.set(newValue);
      run(newValue);
    },
  };
};

export const lsLocale = createLSLocale();
