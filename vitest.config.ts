import { defineConfig } from 'vitest/config';

export default defineConfig({
    test: {
        restoreMocks: true,
        unstubEnvs: true,
        unstubGlobals: true,
        typecheck: { enabled: true },
    },
});
