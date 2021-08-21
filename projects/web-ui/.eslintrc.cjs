const path = require('path');

/** @type {import("eslint").Linter.Config} */
module.exports = {
  extends: ['luma/svelte'],
  overrides: [{
    files: ['*.ts'],
    parserOptions: { project: path.resolve(__dirname, 'tsconfig.json') },
  }],
};
