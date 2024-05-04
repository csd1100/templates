import { contextBridge, ipcRenderer } from 'electron';

contextBridge.exposeInMainWorld('versions', {
    node: () => process.versions.node,
    chrome: () => process.versions.chrome,
    electron: () => process.versions.electron,
});

contextBridge.exposeInMainWorld('api', {
    add: async (a: number, b: number): Promise<number> => {
        return await ipcRenderer.invoke('add', a, b);
    },
});
