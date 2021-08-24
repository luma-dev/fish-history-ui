<script lang="ts">
  import { _ } from 'svelte-i18n';
  import { Temporal } from '@js-temporal/polyfill';
  import { url } from '@roxi/routify';
  import App from '../../../../../components/app.svelte';
  import Cmd from '../../../../../components/cmd.svelte';
  import api from '../../../../../util/api';
  import today from '../../../../../util/today';
  import { params, goto } from '@roxi/routify';
  import type { History } from '../../../../../api/@types';
  import Timestamp from '../../../../../components/timestamp.svelte';

  $: year = Number.parseInt($params.year, 10);
  $: month = Number.parseInt($params.month, 10);
  $: day = Number.parseInt($params.day, 10);

  $: date = new Temporal.PlainDate(year, month, day);
  $: dayOfWeek = date.dayOfWeek;
  $: dayFromTS = date.toZonedDateTime('UTC').epochSeconds;
  $: dayToTS = date.add({ days: 1 }).toZonedDateTime('UTC').epochSeconds - 1;

  $: refetch = async (): Promise<void> => {
    return api.history.chunk
      .$post({
        body: {
          unit: 'day',
          filter: {
            from: dayFromTS,
            to: dayToTS,
          },
        },
      })
      .then(r => {
        histories = r.blocks?.[0]?.histories ?? [];
      });
  };

  $: {
    refetch();
  }
  $: onPrev = () => {
    const prevDate = date.subtract({ days: 1 });
    $goto(`/calendar/${prevDate.year}/${prevDate.month}/${prevDate.day}`);
  };
  $: onNext = () => {
    const nextDate = date.add({ days: 1 });
    $goto(`/calendar/${nextDate.year}/${nextDate.month}/${nextDate.day}`);
  };

  function g(ts: number): Temporal.PlainDateTime {
    return Temporal.Instant.fromEpochSeconds(ts).toZonedDateTimeISO('UTC').toPlainDateTime();
  }

  $: isToday = date.equals($today);

  let histories: History[] | undefined;
</script>

<App on:prev={onPrev} on:next={onNext}>
  <div slot="title">
    <time class={'flex border-b-4 border-solid ' + (isToday ? 'border-blue-200' : 'border-transparent')}>
      <span>
        {$_('date.year', { values: { year } })}
      </span>
      <a href={`/calendar/${year}/${month}`} use:$url>{$_('date.month', { values: { month } })} </a>
      <span>
        {$_('date.day', { values: { day } })}
      </span>
      <span>
        ({$_(`dayOfWeek.short.${dayOfWeek % 7}`)})
      </span>
    </time>
  </div>
  <div class="p-16">
    {#if histories}
      <div class="text-lg">
        {$_('itemCount', { values: { count: histories.length } })}
      </div>
      {#each histories as h}
        <div class="m-2">
          <Timestamp dateTime={g(h.when)} align="start" direction="top" />
          <Cmd code={h.cmd} />
        </div>
      {/each}
    {:else}
      {$_('loading')}
    {/if}
  </div>
</App>
