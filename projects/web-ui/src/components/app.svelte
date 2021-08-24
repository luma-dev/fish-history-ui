<script lang="ts">
  import { _, locales, format } from 'svelte-i18n';
  import { lsLocale } from '../i18n/ls-locale';
  import { createEventDispatcher } from 'svelte';
  import { Button } from 'carbon-components-svelte';
  import ArrowLeft32 from 'carbon-icons-svelte/lib/ArrowLeft32';
  import ArrowRight32 from 'carbon-icons-svelte/lib/ArrowRight32';

  const dispatch = createEventDispatcher();
</script>

<div id="app">
  <nav class="flex items-center p-1 border-b-1 border-solid">
    <div class="flex-1" />
    <Button kind="tertiary" iconDescription={$_('goPrev')} icon={ArrowLeft32} on:click={() => dispatch('prev')} />
    <div class="text-3xl px-4">
      <slot name="title" />
    </div>
    <Button kind="tertiary" iconDescription={$_('goNext')} icon={ArrowRight32} on:click={() => dispatch('next')} />
    <div class="flex flex-1">
      <div class="flex-1" />
      <div>
        <select bind:value={$lsLocale}>
          {#each $locales as locale}
            <option value={locale}>{$format('$meta.language', { locale })}</option>
          {/each}
        </select>
      </div>
    </div>
  </nav>
  <main>
    <slot />
  </main>
</div>

<style>
  #app {
    display: grid;
    min-height: 100vh;
    grid-template-columns: minmax(0, 1fr);
    grid-template-rows: max-content 1fr auto;
    overflow: hidden;
    --border: #efefef;
  }
  main {
    overflow-y: auto;
  }
  nav {
    border-color: var(--border);
  }
</style>
