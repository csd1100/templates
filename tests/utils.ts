import { ElectronApplication, Page, _electron } from 'playwright';

export async function launchApp(): Promise<Page> {
    const electronApp: ElectronApplication = await _electron.launch({
        args: ['./.vite/build/main.js'],
    });

    return electronApp.firstWindow();
}
