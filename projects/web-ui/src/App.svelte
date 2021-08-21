<script lang="ts">
  import 'carbon-components-svelte/css/all.css';
  import { Router } from '@roxi/routify';
  import { routes } from '../.routify/routes';
  import Timestamp from './Timestamp.svelte';
  import apiClient from './util/api-client';
  apiClient.history.chunk
    .$post({
      body: {
        unit: 'day',
        limit: 4,
        filter: {
          from: 1628472588,
          to: 1629472588,
        },
      },
    })
    .then(r => {
      console.log(r);
    });
</script>

<Router {routes} />
