<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '@roxi/routify';
  import today from '../util/today';

  onMount(() => {
    $goto(`/calendar/${$today.year}/${$today.month}`);
  });
</script>

<h1>Redirecting...</h1>
