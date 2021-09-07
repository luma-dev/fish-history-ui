<script lang="ts">
  import { Tooltip as Tooltip0 } from 'carbon-components-svelte';
  import * as copy from 'copy-to-clipboard';
  import * as debounce from 'debounce';
  // eslint-disable-next-line @typescript-eslint/no-explicit-any -- Issue: https://github.com/carbon-design-system/sveld/issues/40
  const Tooltip = Tooltip0 as any;
  export let code: string;
  export let maxLength = Infinity;
  $: open = false;
  const closeDebounced = debounce(() => {
    open = false;
  }, 800);
</script>

<div class="inline-block">
  <Tooltip {open} tooltipBodyId="tooltip-body" hideIcon>
    <p id="tooltip-body">Copied!</p>
    <span
      slot="triggerText"
      class="text-xs cursor-pointer bg-gray-900 hover:bg-gray-700 px-2 py-1 text-white transition-all"
      on:click={() => {
        open = true;
        copy(code);
        closeDebounced();
      }}
    >
      {#if code.length < maxLength}
        {code}
      {:else}
        {code.substr(0, maxLength - 3)}
        <span class="text-pink-200">...</span>
      {/if}
    </span>
  </Tooltip>
</div>
