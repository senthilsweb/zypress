---
title: Installation
description: 'Install @nuxt/content in only two steps in your Nuxt project.'
category: Getting started
position: 1
items:
  - Item1
  - Item2
  - Item3
link: https://codesandbox.io/embed/nuxt-content-l164h?hidenavigation=1&theme=dark
---
# 02 Installation
Add `@nuxt/content` dependency to your project:

# Welcome to Zypress

This a demo application to showcase the development of documentaion website

```
function getList(obj)
{
  alert("Hello")
}
```

<alert>

Hello Sureshbabu

</alert>



<code-group>
  <code-block label="Yarn" active>

  ```bash
  yarn add @nuxt/content
  ```

  </code-block>
  <code-block label="NPM">

  ```bash
  npm install @nuxt/content
  ```

  </code-block>
</code-group>

<list :items="items"></list>


<alert>

Check out an info alert with a `codeblock` and a [link](/themes/docs)!

</alert>

<badge>v1.2+</badge>




Then, add `@nuxt/content` to the `modules` section of `nuxt.config.js`:

```js[nuxt.config.js]
{
  modules: [
    '@nuxt/content'
  ],
  content: {
    // Options
  }
}
```

## TypeScript

Add the types to your "types" array in tsconfig.json after the `@nuxt/types` (Nuxt 2.9.0+) or `@nuxt/vue-app` entry.

**tsconfig.json**

```json
{
  "compilerOptions": {
    "types": [
      "@nuxt/types",
      "@nuxt/content"
    ]
  }
}
```

> **Why?**
>
> Because of the way nuxt works the `$content` property on the context has to be merged into the nuxt `Context` interface via [declaration merging](https://www.typescriptlang.org/docs/handbook/declaration-merging.html). Adding `@nuxt/content` to your types will import the types from the package and make typescript aware of the additions to the `Context` interface.