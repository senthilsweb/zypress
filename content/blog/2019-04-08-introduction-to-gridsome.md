---
title: Introduction to Gridsome
path: /casestudy/introduction-to-gridsome
date: 2019-04-08
summary: Gridsome is a Vue.js-powered, modern site generator for building the fastest possible websites for any Headless CMS, APIs or Markdown-files. Gridsome makes it easy and fun for developers to create fast, beautiful websites without needing to become a performance expert.
background: background goes here
introduction: introduction goes here
industries: [Healthcare, Telecom]
coverimage: https://source.unsplash.com/random
author: Senthilnathan Karuppaiah
avatar: https://res.cloudinary.com/nathansweb/image/upload/v1623612257/profile/sk_profile_sq.png
type: Blog
tags:
  - Mobile Apps
  - Field data capture
  - Internet Of Things
platforms: [Web]
os: 
  -
    name: IOS
    version: 2.3.0
    icon: javascript.png
    store: http://www.store.link.goes.here.com/
  -
    name: Android
    version: 2.7.3
    icon: android.png
    store: http://www.store.link.goes.here.com/
screenshots:
  -
    title: Splash screen
    path: /image/path/goes/here
  -
    title: Dashboard
    path: /image/path/goes/here
  
technologies:
  -
    name: Vue.js
    version: 3.x.x
    icon: vue.svg
  -
    name: JavaScript
    icon: javascript.svg  
  -
    name: Gridsome
    version: 2.7.3
    icon: gridsome.svg
---

![background](./images/blog_bg_1.jpg)

> Gridsome is a Vue.js-powered, modern site generator for building the fastest possible websites for any Headless CMS, APIs or Markdown-files. Gridsome makes it easy and fun for developers to create fast, beautiful websites without needing to become a performance expert.

### Why Gridsome?

- **Local development with hot-reloading** - See code changes in real-time.
- **Data source plugins** - Use it for any popular Headless CMSs, APIs or Markdown-files.
- **File-based page routing** - Quickly create and manage routes with files.
- **Centralized data managment** - Pull data into a local, unified GraphQL data layer.
- **Vue.js for frontend** - A lightweight and approachable front-end framework.
- **Auto-optimized code** - Get code-splitting and asset optimization out-of-the-box.
- **Static files generation** - Deploy securely to any CDN or static web host.

[Learn more about how Gridsome works](/docs/how-it-works)

```js
<template>
  <Layout>
    <div class="container-inner mx-auto my-16">
      <h1 class="text-4xl font-bold leading-tight">{{ $page.case.title }}</h1>
      <div class="text-xl text-gray-600 mb-8">{{ $page.case.date }}</div>
      <div class="markdown-body" v-html="$page.post.content" />
    </div>
  </Layout>
</template>
```


### Prerequisites
You should have basic knowledge about HTML, CSS, [Vue.js](https://vuejs.org) and how to use the [Terminal](https://www.linode.com/docs/tools-reference/tools/using-the-terminal/). Knowing how [Vue Single File components](https://vuejs.org/v2/guide/single-file-components.html) & [GraphQL](https://www.graphql.com/) works is a plus, but not required. Gridsome is a great way to learn both.

Gridsome requires **Node.js** and recommends **Yarn**. [How to setup](/docs/prerequisites)

![background](./images/background.jpg)

### 1. Install Gridsome CLI tool

Using yarn:
`yarn global add @gridsome/cli`

Using npm:
`npm install --global @gridsome/cli`

### 2. Create a Gridsome project

1. `gridsome create my-gridsome-site` to create a new project </li>
2. `cd my-gridsome-site` to open folder
3. `gridsome develop` to start local dev server at `http://localhost:8080`
4. Happy coding 🎉🙌

### 3. Next steps

1. Create `.vue` components in the `/pages` directory to create page routes.
2. Use `gridsome build` to generate static files in a `/dist` folder


- [How it works](/docs/how-it-works)
- [How Pages work](/docs/pages)
- [How to deploy](/docs/deployment)