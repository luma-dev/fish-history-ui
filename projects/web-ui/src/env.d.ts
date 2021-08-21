/// <reference types="svelte" />
/// <reference types="vite/client" />

declare interface ImportMeta {
  env: {
    MODE: string;
    BASE_URL: string;
    PROD: boolean;
    DEV: boolean;
    // custom
    VITE_API_PORT: string;
  };
}
