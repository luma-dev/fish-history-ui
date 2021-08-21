/** @type {import("eslint").parserOptions} */
module.exports = {
  overrides: [
    {
      files: ['*.ts'],
      rules: {
        'import/no-unresolved': ['error', {
          ignore: ['^[^.].*\\.(?:css)$'],
        }],
      },
    },
    {
      files: ['*.svelte'],

      parserOptions: {
        ecmaVersion: 2019,
        sourceType: 'module',
      },
      env: {
        es6: true,
        browser: true,
      },
      plugins: [
        'svelte3',
        '@typescript-eslint',
        'prettier',
      ],
      parser: '@typescript-eslint/parser',
      processor: 'svelte3/svelte3',
      settings: {
        'svelte3/typescript': true,
      },
    },
  ],
};
