<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '@roxi/routify';
  import getYear from 'date-fns/getYear';
  import getMonth from 'date-fns/getMonth';

  onMount(() => {
    const now = Date.now();
    $goto(`/calendar/${getYear(now)}/${getMonth(now)}`);
  });
</script>

<h1>Redirecting...</h1>
