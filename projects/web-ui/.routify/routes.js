
/**
 * @roxi/routify 2.18.3
 * File generated Sun Aug 22 2021 02:08:00 GMT+0900 (Japan Standard Time)
 */

export const __version = "2.18.3"
export const __timestamp = "2021-08-21T17:08:00.858Z"

//buildRoutes
import { buildClientTree } from "@roxi/routify/runtime/buildRoutes"

//imports


//options
export const options = {}

//tree
export const _tree = {
  "name": "_folder",
  "filepath": "/_folder.svelte",
  "root": true,
  "ownMeta": {},
  "absolutePath": "/home/luma/workspace/0-hisview/main/projects/web-ui/src/pages/_folder.svelte",
  "children": [
    {
      "isFile": true,
      "isDir": false,
      "file": "_fallback.svelte",
      "filepath": "/_fallback.svelte",
      "name": "_fallback",
      "ext": "svelte",
      "badExt": false,
      "absolutePath": "/home/luma/workspace/0-hisview/main/projects/web-ui/src/pages/_fallback.svelte",
      "importPath": "../src/pages/_fallback.svelte",
      "isLayout": false,
      "isReset": false,
      "isIndex": false,
      "isFallback": true,
      "isPage": false,
      "ownMeta": {},
      "meta": {
        "recursive": true,
        "preload": false,
        "prerender": true
      },
      "path": "/_fallback",
      "id": "__fallback",
      "component": () => import('../src/pages/_fallback.svelte').then(m => m.default)
    },
    {
      "isFile": false,
      "isDir": true,
      "file": "[year]",
      "filepath": "/[year]",
      "name": "[year]",
      "ext": "",
      "badExt": false,
      "absolutePath": "/home/luma/workspace/0-hisview/main/projects/web-ui/src/pages/[year]",
      "children": [
        {
          "isFile": false,
          "isDir": true,
          "file": "[month]",
          "filepath": "/[year]/[month]",
          "name": "[month]",
          "ext": "",
          "badExt": false,
          "absolutePath": "/home/luma/workspace/0-hisview/main/projects/web-ui/src/pages/[year]/[month]",
          "children": [
            {
              "isFile": false,
              "isDir": true,
              "file": "[day]",
              "filepath": "/[year]/[month]/[day]",
              "name": "[day]",
              "ext": "",
              "badExt": false,
              "absolutePath": "/home/luma/workspace/0-hisview/main/projects/web-ui/src/pages/[year]/[month]/[day]",
              "children": [
                {
                  "isFile": true,
                  "isDir": false,
                  "file": "index.svelte",
                  "filepath": "/[year]/[month]/[day]/index.svelte",
                  "name": "index",
                  "ext": "svelte",
                  "badExt": false,
                  "absolutePath": "/home/luma/workspace/0-hisview/main/projects/web-ui/src/pages/[year]/[month]/[day]/index.svelte",
                  "importPath": "../src/pages/[year]/[month]/[day]/index.svelte",
                  "isLayout": false,
                  "isReset": false,
                  "isIndex": true,
                  "isFallback": false,
                  "isPage": true,
                  "ownMeta": {},
                  "meta": {
                    "recursive": true,
                    "preload": false,
                    "prerender": true
                  },
                  "path": "/:year/:month/:day/index",
                  "id": "__year__month__day_index",
                  "component": () => import('../src/pages/[year]/[month]/[day]/index.svelte').then(m => m.default)
                }
              ],
              "isLayout": false,
              "isReset": false,
              "isIndex": false,
              "isFallback": false,
              "isPage": false,
              "ownMeta": {},
              "meta": {
                "recursive": true,
                "preload": false,
                "prerender": true
              },
              "path": "/:year/:month/:day"
            },
            {
              "isFile": true,
              "isDir": false,
              "file": "index.svelte",
              "filepath": "/[year]/[month]/index.svelte",
              "name": "index",
              "ext": "svelte",
              "badExt": false,
              "absolutePath": "/home/luma/workspace/0-hisview/main/projects/web-ui/src/pages/[year]/[month]/index.svelte",
              "importPath": "../src/pages/[year]/[month]/index.svelte",
              "isLayout": false,
              "isReset": false,
              "isIndex": true,
              "isFallback": false,
              "isPage": true,
              "ownMeta": {},
              "meta": {
                "recursive": true,
                "preload": false,
                "prerender": true
              },
              "path": "/:year/:month/index",
              "id": "__year__month_index",
              "component": () => import('../src/pages/[year]/[month]/index.svelte').then(m => m.default)
            }
          ],
          "isLayout": false,
          "isReset": false,
          "isIndex": false,
          "isFallback": false,
          "isPage": false,
          "ownMeta": {},
          "meta": {
            "recursive": true,
            "preload": false,
            "prerender": true
          },
          "path": "/:year/:month"
        },
        {
          "isFile": true,
          "isDir": false,
          "file": "index.svelte",
          "filepath": "/[year]/index.svelte",
          "name": "index",
          "ext": "svelte",
          "badExt": false,
          "absolutePath": "/home/luma/workspace/0-hisview/main/projects/web-ui/src/pages/[year]/index.svelte",
          "importPath": "../src/pages/[year]/index.svelte",
          "isLayout": false,
          "isReset": false,
          "isIndex": true,
          "isFallback": false,
          "isPage": true,
          "ownMeta": {},
          "meta": {
            "recursive": true,
            "preload": false,
            "prerender": true
          },
          "path": "/:year/index",
          "id": "__year_index",
          "component": () => import('../src/pages/[year]/index.svelte').then(m => m.default)
        }
      ],
      "isLayout": false,
      "isReset": false,
      "isIndex": false,
      "isFallback": false,
      "isPage": false,
      "ownMeta": {},
      "meta": {
        "recursive": true,
        "preload": false,
        "prerender": true
      },
      "path": "/:year"
    },
    {
      "isFile": true,
      "isDir": false,
      "file": "index.svelte",
      "filepath": "/index.svelte",
      "name": "index",
      "ext": "svelte",
      "badExt": false,
      "absolutePath": "/home/luma/workspace/0-hisview/main/projects/web-ui/src/pages/index.svelte",
      "importPath": "../src/pages/index.svelte",
      "isLayout": false,
      "isReset": false,
      "isIndex": true,
      "isFallback": false,
      "isPage": true,
      "ownMeta": {},
      "meta": {
        "recursive": true,
        "preload": false,
        "prerender": true
      },
      "path": "/index",
      "id": "_index",
      "component": () => import('../src/pages/index.svelte').then(m => m.default)
    }
  ],
  "isLayout": true,
  "isReset": false,
  "isIndex": false,
  "isFallback": false,
  "isPage": false,
  "isFile": true,
  "file": "_folder.svelte",
  "ext": "svelte",
  "badExt": false,
  "importPath": "../src/pages/_folder.svelte",
  "meta": {
    "recursive": true,
    "preload": false,
    "prerender": true
  },
  "path": "/",
  "id": "__folder",
  "component": () => import('../src/pages/_folder.svelte').then(m => m.default)
}


export const {tree, routes} = buildClientTree(_tree)

