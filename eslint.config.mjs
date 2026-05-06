// ESLint flat config for api-snippets JS files.
// Enforces prettier formatting and prefer-const; most other rules are relaxed.
import js from '@eslint/js';
import prettier from 'eslint-plugin-prettier/recommended';
import globals from 'globals';

export default [
  js.configs.recommended,
  prettier,
  {
    languageOptions: {
      globals: { ...globals.browser, ...globals.commonjs, ...globals.node },
      ecmaVersion: 2022,
    },
    rules: {
      'prettier/prettier': [
        'error',
        { singleQuote: true, trailingComma: 'es5' },
      ],
      'no-console': 'off',
      'no-unused-vars': 'off',
      'no-undef': 'off',
      'no-useless-escape': 'off',
      'prefer-const': 'error',
    },
  },
  { ignores: ['tools/', 'node_modules/'] },
];
