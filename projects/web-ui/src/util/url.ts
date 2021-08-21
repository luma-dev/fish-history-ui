/* eslint-disable import/prefer-default-export */
import store from 'svelte/store';
import { page } from '@roxi/routify';

function createUnit(): store.Readable<string> {
  const unit = '';
  return {
    subscribe(run) {
      run(unit);
      return page.subscribe(p => {
        run(p.path.match(/^\/([^/]*)/)?.[1] ?? '');
      });
    },
  };
}

export const unit = createUnit();
/* eslint-enable import/prefer-default-export */
