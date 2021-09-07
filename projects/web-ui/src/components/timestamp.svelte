<script lang="ts">
  import { _ } from 'svelte-i18n';
  import type { Temporal } from '@js-temporal/polyfill';
  import offset from '../util/timezone-offset-seconds';
  import { TooltipDefinition } from 'carbon-components-svelte';
  import type { TooltipDefinitionProps } from 'carbon-components-svelte/types/TooltipDefinition/TooltipDefinition';
  export let align: TooltipDefinitionProps['align'] = undefined;
  export let direction: TooltipDefinitionProps['direction'] = undefined;
  export let dateTime: Temporal.PlainDateTime;
  $: instant =
    typeof $offset === 'number' && dateTime.toZonedDateTime('UTC').toInstant().subtract({ seconds: $offset });
</script>

<TooltipDefinition
  tooltipText={instant ? `${instant.toString()}\n(${instant.toLocaleString()})` : $_('loading')}
  {align}
  {direction}
>
  {dateTime.toPlainTime().toString()}
</TooltipDefinition>
