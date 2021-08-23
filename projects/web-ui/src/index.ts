import 'windi.css';
import './styles.css';
import './i18n';
import Root from './root.svelte';

const root = new Root({
  target: document.body,
});

export default root;
