<script lang="ts">
  import format from 'date-fns/format';

  export let time: Date;
</script>

<time class="underline underline-dotted">{format(time, 'Pp')}</time>
