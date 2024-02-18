import { test, expect } from '@playwright/test';

test('has title', async ({ page }) => {
    await page.goto('http://localhost:8080/');

    // Expect a title "to contain" a substring.
    await expect(page).toHaveTitle(/Saludos/);
});

// test('test', async ({ page }) => {
//     await page.goto('http://localhost:8080/');
//     await page.getByRole('link', { name: 'Favicon Saludos' }).click();
//     await page.getByRole('navigation').click();
//     await page.getByRole('navigation').click();
// });