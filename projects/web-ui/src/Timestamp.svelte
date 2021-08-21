<script lang="ts">
  import Time from './Time.svelte';
  const size = 30;

  const time = new Date();
</script>

<div class="h-16 flex content-center items-center p-0">
  <svg
    class="h-16 w-16"
    viewBox="-50 -50 100 100"
    xmlns="http://www.w3.org/2000/svg"
    xmlns:xlink="http://www.w3.org/1999/xlink"
  >
    <rect x={-size / 2} y="-50" width="6" height="100" class="fill-gray-600" />
    <circle cx="-10" cy="0" r={size} class="fill-teal-500" />
  </svg>
  <Time {time} />
</div>

<style>
</style>
