<script lang="ts">
  import { goto } from '@roxi/routify';
  import { Button } from 'carbon-components-svelte';
  import ArrowLeft32 from 'carbon-icons-svelte/lib/ArrowLeft32';
  import ArrowRight32 from 'carbon-icons-svelte/lib/ArrowRight32';
  import { OverflowMenu, OverflowMenuItem } from 'carbon-components-svelte';
  import { unit } from '../util/url';
  import { title, onPrev, onNext } from '../util/view';
  import type { View } from '../types';

  const views: View[] = [
    { text: '日', unit: 'day' },
    { text: '週', unit: 'week' },
    { text: '月', unit: 'month' },
    { text: '年', unit: 'year' },
  ];
  const rev = (unit: string) => {
    for (const view of views) {
      if (view.unit === unit) return view;
    }
    return views[2];
  };
  $: name = rev($unit || 'month').text;
</script>

<div id="app">
  <nav>
    <Button kind="tertiary" iconDescription="戻る" icon={ArrowLeft32} on:click={$onPrev} />
    <Button kind="tertiary" iconDescription="進む" icon={ArrowRight32} on:click={$onNext} />
    <div class="spacing" />
    <time>
      {$title}
    </time>
    <div class="spacing" />
    <OverflowMenu style="width: auto;" flipped>
      <div slot="menu" style="padding: 1rem; color: red;">
        {name}
      </div>
      {#each views as view}
        <OverflowMenuItem text={view.text} on:click={() => $goto(`/${view.unit}`)} />
      {/each}
    </OverflowMenu>
  </nav>
  <main>
    <slot />
  </main>
</div>

<style>
  #app {
    display: grid;
    min-height: 100vh;
    height: 100vh;
    max-height: 100vh;
    grid-template-columns: minmax(0, 1fr);
    grid-template-rows: max-content 1fr auto;
    overflow: hidden;
    --border: hsl(16deg 100% 90%);
  }
  nav {
    display: flex;
    align-items: center;
    padding: 0.2rem;
    border-bottom: 0.5px solid var(--border);
  }
  .spacing {
    flex-grow: 1;
  }
  time {
    font-size: 2rem;
  }
</style>
