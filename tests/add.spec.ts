import { test, expect } from '@playwright/test';
import type { Page } from 'playwright';

import { launchApp } from './utils';

test('add works', async () => {
    const page = await launchApp();

    const a = page.locator('#a');
    const b = page.locator('#b');
    const sum = page.locator('#sum');

    await expect(a).toBeVisible();
    await expect(b).toBeVisible();
    await expect(sum).toBeVisible();

    await a.fill('100');
    await b.fill('100');

    await expect(sum).toHaveValue('200');

    await a.fill('-100');
    await b.fill('-100');

    await expect(sum).toHaveValue('-200');
});

test('invalid input adds red border and 0 sum', async () => {
    const page = await launchApp();

    const a = page.locator('#a');
    const b = page.locator('#b');
    const sum = page.locator('#sum');

    await expect(a).toBeVisible();
    await expect(b).toBeVisible();
    await expect(sum).toBeVisible();

    await a.fill('a');
    await b.fill('b');

    await expect(a).toHaveClass(/border-red-500 focus:ring-red-200/);
    await expect(b).toHaveClass(/border-red-500 focus:ring-red-200/);
    await expect(sum).toHaveValue('0');
});
