/// <reference types="electron" />

declare interface Window {
    api: {
        add: (number, number) => Promise<number>;
    };
}
