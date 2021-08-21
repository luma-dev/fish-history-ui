import * as store from 'svelte/store';

export const title = store.writable('');
export const onPrev = store.writable<(() => void) | null>(null);
export const onNext = store.writable<(() => void) | null>(null);
