import { addMessages, register, init, getLocaleFromNavigator } from 'svelte-i18n';
import { lsLocaleKey } from './ls-locale';

addMessages('en', { $meta: { language: 'English' } });
addMessages('ja', { $meta: { language: '日本語' } });

register('en', () => import('./en.json'));
register('ja', () => import('./ja.json'));

const fallbackLocale = 'en';
if (localStorage[lsLocaleKey] === getLocaleFromNavigator()) localStorage.removeItem(lsLocaleKey);

init({
  fallbackLocale,
  initialLocale: localStorage[lsLocaleKey] || getLocaleFromNavigator(),
});
