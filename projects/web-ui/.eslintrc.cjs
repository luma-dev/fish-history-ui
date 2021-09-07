const configureBase = require('@luma-dev/eslint-config-base/configure');
const configureSvelte = require('@luma-dev/eslint-config-svelte/configure');

/** @type {import('eslint').Linter.Config} */
module.exports = {
  extends: ['@luma-dev/svelte'],
  overrides: [...configureBase(__dirname), ...configureSvelte(__dirname)],
};
