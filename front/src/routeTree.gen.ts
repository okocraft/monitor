/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file was automatically generated by TanStack Router.
// You should NOT make any changes in this file as it will be overwritten.
// Additionally, you should also exclude this file from your linter and/or formatter to prevent it from being checked or modified.

// Import Routes

import { Route as rootRoute } from './routes/__root'
import { Route as IndexImport } from './routes/index'
import { Route as UsersIndexImport } from './routes/users/index'
import { Route as ExampleIndexImport } from './routes/example/index'
import { Route as GoogleResultIndexImport } from './routes/google/result/index'
import { Route as GoogleLoginIndexImport } from './routes/google/login/index'

// Create/Update Routes

const IndexRoute = IndexImport.update({
  id: '/',
  path: '/',
  getParentRoute: () => rootRoute,
} as any)

const UsersIndexRoute = UsersIndexImport.update({
  id: '/users/',
  path: '/users/',
  getParentRoute: () => rootRoute,
} as any)

const ExampleIndexRoute = ExampleIndexImport.update({
  id: '/example/',
  path: '/example/',
  getParentRoute: () => rootRoute,
} as any)

const GoogleResultIndexRoute = GoogleResultIndexImport.update({
  id: '/google/result/',
  path: '/google/result/',
  getParentRoute: () => rootRoute,
} as any)

const GoogleLoginIndexRoute = GoogleLoginIndexImport.update({
  id: '/google/login/',
  path: '/google/login/',
  getParentRoute: () => rootRoute,
} as any)

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/': {
      id: '/'
      path: '/'
      fullPath: '/'
      preLoaderRoute: typeof IndexImport
      parentRoute: typeof rootRoute
    }
    '/example/': {
      id: '/example/'
      path: '/example'
      fullPath: '/example'
      preLoaderRoute: typeof ExampleIndexImport
      parentRoute: typeof rootRoute
    }
    '/users/': {
      id: '/users/'
      path: '/users'
      fullPath: '/users'
      preLoaderRoute: typeof UsersIndexImport
      parentRoute: typeof rootRoute
    }
    '/google/login/': {
      id: '/google/login/'
      path: '/google/login'
      fullPath: '/google/login'
      preLoaderRoute: typeof GoogleLoginIndexImport
      parentRoute: typeof rootRoute
    }
    '/google/result/': {
      id: '/google/result/'
      path: '/google/result'
      fullPath: '/google/result'
      preLoaderRoute: typeof GoogleResultIndexImport
      parentRoute: typeof rootRoute
    }
  }
}

// Create and export the route tree

export interface FileRoutesByFullPath {
  '/': typeof IndexRoute
  '/example': typeof ExampleIndexRoute
  '/users': typeof UsersIndexRoute
  '/google/login': typeof GoogleLoginIndexRoute
  '/google/result': typeof GoogleResultIndexRoute
}

export interface FileRoutesByTo {
  '/': typeof IndexRoute
  '/example': typeof ExampleIndexRoute
  '/users': typeof UsersIndexRoute
  '/google/login': typeof GoogleLoginIndexRoute
  '/google/result': typeof GoogleResultIndexRoute
}

export interface FileRoutesById {
  __root__: typeof rootRoute
  '/': typeof IndexRoute
  '/example/': typeof ExampleIndexRoute
  '/users/': typeof UsersIndexRoute
  '/google/login/': typeof GoogleLoginIndexRoute
  '/google/result/': typeof GoogleResultIndexRoute
}

export interface FileRouteTypes {
  fileRoutesByFullPath: FileRoutesByFullPath
  fullPaths: '/' | '/example' | '/users' | '/google/login' | '/google/result'
  fileRoutesByTo: FileRoutesByTo
  to: '/' | '/example' | '/users' | '/google/login' | '/google/result'
  id:
    | '__root__'
    | '/'
    | '/example/'
    | '/users/'
    | '/google/login/'
    | '/google/result/'
  fileRoutesById: FileRoutesById
}

export interface RootRouteChildren {
  IndexRoute: typeof IndexRoute
  ExampleIndexRoute: typeof ExampleIndexRoute
  UsersIndexRoute: typeof UsersIndexRoute
  GoogleLoginIndexRoute: typeof GoogleLoginIndexRoute
  GoogleResultIndexRoute: typeof GoogleResultIndexRoute
}

const rootRouteChildren: RootRouteChildren = {
  IndexRoute: IndexRoute,
  ExampleIndexRoute: ExampleIndexRoute,
  UsersIndexRoute: UsersIndexRoute,
  GoogleLoginIndexRoute: GoogleLoginIndexRoute,
  GoogleResultIndexRoute: GoogleResultIndexRoute,
}

export const routeTree = rootRoute
  ._addFileChildren(rootRouteChildren)
  ._addFileTypes<FileRouteTypes>()

/* ROUTE_MANIFEST_START
{
  "routes": {
    "__root__": {
      "filePath": "__root.tsx",
      "children": [
        "/",
        "/example/",
        "/users/",
        "/google/login/",
        "/google/result/"
      ]
    },
    "/": {
      "filePath": "index.tsx"
    },
    "/example/": {
      "filePath": "example/index.tsx"
    },
    "/users/": {
      "filePath": "users/index.tsx"
    },
    "/google/login/": {
      "filePath": "google/login/index.tsx"
    },
    "/google/result/": {
      "filePath": "google/result/index.tsx"
    }
  }
}
ROUTE_MANIFEST_END */
