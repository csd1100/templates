import prettier from 'eslint-config-prettier';
import js from '@eslint/js';
import globals from 'globals';

export default [
    {
        files: ['**/*.js'],
        plugins: {
            prettier: prettier,
        },
        languageOptions: {
            ecmaVersion: 'latest',
            globals: {
                ...globals.browser,
                ...globals.node,
                ...globals.jest,
            },
        },
        rules: {
            ...js.configs.recommended.rules,
            ...prettier.rules,
        },
    },
];
