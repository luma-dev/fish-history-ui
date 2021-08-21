<script lang="ts">
  import { Temporal } from '@js-temporal/polyfill';
  import Cmd from '../../../components/cmd.svelte';
  import { Button } from 'carbon-components-svelte';
  import TrashCan16 from 'carbon-icons-svelte/lib/TrashCan16';
  import { params, goto } from '@roxi/routify';
  import type { HistoryBlock } from '../../../api/@types';
  import apiClient from '../../../util/api-client';
  import { onMount, onDestroy } from 'svelte';
  import { title, onPrev, onNext } from '../../../util/view';
  import { lastDayOfMonth, firstDayOfWeek, lastDayOfWeek } from '../../../util/temporal';

  $: year = Number.parseInt($params.year, 10);
  $: month = Number.parseInt($params.month, 10);

  $: monthFrom = new Temporal.PlainDate(year, month, 1);
  $: monthTo = lastDayOfMonth(monthFrom);
  $: spanFrom = firstDayOfWeek(monthFrom);
  $: spanTo = lastDayOfWeek(monthTo);

  $: refetch = () => {
    console.log({ year, month, monthFrom, monthTo, spanFrom, spanTo });
    apiClient.history.chunk
      .$post({
        body: {
          unit: 'day',
          limit,
          filter: {
            from: spanFrom.toZonedDateTime('UTC').epochSeconds,
            to: spanTo.toZonedDateTime('UTC').epochSeconds,
          },
        },
      })
      .then(r => {
        if (r.blocks) {
          table = pad(r.blocks);
          console.log(r.blocks);
        }
      });
  };

  $: {
    $title = `${year}年${month}月`;
    $onPrev = () => {
      $goto(`/${month === 1 ? year - 1 : year}/${month === 1 ? 12 : month - 1}`);
    };
    $onNext = () => {
      $goto(`/${month === 12 ? year + 1 : year}/${month === 12 ? 1 : month + 1}`);
    };
    refetch();
  }

  function pad(blocks: HistoryBlock[]): HistoryBlock[][] {
    const table: HistoryBlock[][] = [];
    let inWeek: HistoryBlock[] = [];
    let i = 0;
    for (let now = spanFrom; Temporal.PlainDateTime.compare(now, spanTo) < 0; now = now.add({ days: 1 })) {
      console.log('now', { now });
      if (i < blocks.length && blocks[i].from <= now.toZonedDateTime('UTC').epochSeconds) {
        console.log('block', { block: blocks[i] });
        inWeek.push(blocks[i]);
        i++;
      } else {
        inWeek.push({
          count: 0,
          from: now.toZonedDateTime('UTC').epochSeconds,
          to: now.add({ days: 1 }).toZonedDateTime('UTC').epochSeconds,
        });
      }
      if (inWeek.length === 7) {
        table.push(inWeek);
        inWeek = [];
      }
    }
    return table;
  }

  const header = ['日', '月', '火', '水', '木', '金', '土'];
  const limit = 3;

  let table: HistoryBlock[][] | undefined;
</script>

{#if table}
  <div class="month">
    <div class="row">
      {#each header as h}
        <div class="column header">{h}</div>
      {/each}
    </div>
    {#each table as row}
      <div class="row">
        {#each row as column}
          <div class="column">
            <time>
              {#if Temporal.Instant.fromEpochSeconds(column.from).toZonedDateTimeISO('UTC').month === 1}
                {Temporal.Instant.fromEpochSeconds(column.from).toZonedDateTimeISO('UTC').month}月
              {/if}
              {Temporal.Instant.fromEpochSeconds(column.from).toZonedDateTimeISO('UTC').day}日
            </time>
            {#each column.histories ?? [] as h}
              <div>
                <Cmd>
                  {h.cmd}
                </Cmd>
              </div>
            {/each}
            <span>
              {#if column.count - limit > 0}
                ...他 {column.count - limit} 件
              {/if}
            </span>
          </div>
        {/each}
      </div>
    {/each}
  </div>
{:else}
  Loading...
{/if}

<style>
  .month {
    height: 100%;
    width: 100%;
    overflow: hidden;
    display: grid;
    grid-template-rows: auto;
    grid-auto-rows: 1fr;
    border: 0.5px solid var(--border);
  }
  .row {
    display: grid;
    grid-auto-columns: 1fr;
    grid-auto-flow: column;
    overflow: hidden;
  }
  .column {
    border: 0.5px solid var(--border);
    overflow: hidden;
  }
  .header {
    font-size: 1.3rem;
    padding: 2px;
    text-align: center;
  }
  time {
    display: block;
    text-align: center;
  }
</style>
