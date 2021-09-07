<script lang="ts">
  import { _ } from 'svelte-i18n';
  import { Temporal } from '@js-temporal/polyfill';
  import Cmd from '../../../../components/cmd.svelte';
  import { params, goto } from '@roxi/routify';
  import type { HistoryBlock } from '../../../../api/@types';
  import api from '../../../../util/api';
  import today from '../../../../util/today';
  import { lastDayOfMonth, firstDayOfWeek, lastDayOfWeek } from '../../../../util/temporal';
  import App from '../../../../components/app.svelte';

  $: year = Number.parseInt($params.year, 10);
  $: month = Number.parseInt($params.month, 10);
  $: monthYear = Temporal.PlainYearMonth.from({ year, month });
  $: isThisMonth = monthYear.equals($today.toPlainYearMonth());

  $: monthFrom = new Temporal.PlainDate(year, month, 1);
  $: monthTo = lastDayOfMonth(monthFrom);
  $: spanFrom = firstDayOfWeek(monthFrom);
  $: spanTo = lastDayOfWeek(monthTo);
  $: spanFromTS = spanFrom.toZonedDateTime('UTC').epochSeconds;
  $: spanToTS = spanTo.add({ days: 1 }).toZonedDateTime('UTC').epochSeconds - 1;

  $: refetch = async (): Promise<void> => {
    return api.history.chunk
      .$post({
        body: {
          unit: 'day',
          limit,
          filter: {
            from: spanFromTS,
            to: spanToTS,
          },
        },
      })
      .then((r) => {
        if (r.blocks) {
          table = pad(r.blocks);
        }
      });
  };

  $: {
    void refetch();
  }
  $: onPrev = () => {
    $goto(`/calendar/${month === 1 ? year - 1 : year}/${month === 1 ? 12 : month - 1}`);
  };
  $: onNext = () => {
    $goto(`/calendar/${month === 12 ? year + 1 : year}/${month === 12 ? 1 : month + 1}`);
  };

  function pad(blocks: HistoryBlock[]): HistoryBlock[][] {
    const table: HistoryBlock[][] = [];
    let inWeek: HistoryBlock[] = [];
    let i = 0;
    for (let now = spanFrom; Temporal.PlainDateTime.compare(now, spanTo) <= 0; now = now.add({ days: 1 })) {
      if (i < blocks.length && blocks[i].from <= now.toZonedDateTime('UTC').epochSeconds) {
        inWeek.push(blocks[i]);
        i += 1;
      } else {
        inWeek.push({
          count: 0,
          from: now.toZonedDateTime('UTC').epochSeconds,
          to: now.add({ days: 1 }).toZonedDateTime('UTC').epochSeconds - 1,
        });
      }
      if (inWeek.length === 7) {
        table.push(inWeek);
        inWeek = [];
      }
    }
    return table;
  }

  function g(u: number): Temporal.PlainDate {
    return Temporal.Instant.fromEpochSeconds(u).toZonedDateTimeISO('UTC').toPlainDate();
  }

  function inMonth(d: Temporal.PlainDate): boolean {
    return Temporal.PlainDate.compare(monthFrom, d) <= 0 && Temporal.PlainDate.compare(d, monthTo) <= 0;
  }

  $: header = [
    $_('dayOfWeek.short.0'),
    $_('dayOfWeek.short.1'),
    $_('dayOfWeek.short.2'),
    $_('dayOfWeek.short.3'),
    $_('dayOfWeek.short.4'),
    $_('dayOfWeek.short.5'),
    $_('dayOfWeek.short.6'),
  ];
  const limit = 3;
  function isToday(d: Temporal.PlainDate): boolean {
    return d.equals($today);
  }

  let table: HistoryBlock[][] | undefined;
</script>

<App on:prev={onPrev} on:next={onNext}>
  <div slot="title">
    <time class={`flex border-b-4 border-solid ${isThisMonth ? 'border-blue-200' : 'border-transparent'}`}>
      {$_('yearMonth.full', { values: { year, month } })}
    </time>
  </div>
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
            <div
              class={inMonth(g(column.from)) ? 'column' : 'column out'}
              on:click={() => $goto(`/calendar/${g(column.from).year}/${g(column.from).month}/${g(column.from).day}`)}
            >
              <time class={`flex justify-center ${isToday(g(column.from)) ? 'bg-blue-200' : ''}`}>
                {#if g(column.from).day === 1}
                  <span>
                    {$_('monthDay.month', { values: { month: g(column.from).month } })}
                  </span>
                {/if}
                <span>
                  {$_('monthDay.day', {
                    values: { day: g(column.from).day },
                  })}
                </span>
              </time>
              {#each column.histories ?? [] as h}
                <div class="m-1">
                  <Cmd code={h.cmd} maxLength={28} />
                </div>
              {/each}
              <span class="text-xs">
                {#if column.count - limit > 0}
                  {$_('othersCount', { values: { count: column.count - limit } })}
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
</App>

<style>
  .month {
    height: 100%;
    width: 100%;
    display: grid;
    grid-template-rows: auto;
    grid-auto-rows: 1fr;
    border: 0.5px solid var(--border);
    min-width: 0;
    overflow-wrap: anywhere;
  }
  .row {
    display: grid;
    grid-auto-columns: 1fr;
    grid-auto-flow: column;
    min-width: 0;
  }
  .column {
    border: 0.5px solid var(--border);
    min-width: 0;
  }
  .column:not(.header) {
    cursor: pointer;
  }
  .column:not(.header):hover {
    background: hsl(0deg 0% 98%);
  }
  .column.out {
    background: hsl(0deg 0% 94%);
  }
  .header {
    font-size: 1.3rem;
    padding: 2px;
    text-align: center;
  }
</style>
