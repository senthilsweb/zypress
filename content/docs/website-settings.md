---
title: Website Configuration
category: User guide
position: 2
---

# Introduction
This guide walks you thru the initial application wide configuration to fit to your branding and requirements. All of the configurations are static and design time configs meaning it requires re-build and re-dpeloyment to take the new configuration in effect.

# Website Configurations 

Zypress comes with intial defaults however you can override by editing the 

* `.env` file in the root
* and, `/content/settings.json`
* and, `/content/profile/about.json`

## Application Defaults

You can override basic application default is by editing the environment variable `/.env` file at the root. The variables are self-explanatory.

```
APP_NAME='alex-website'
APP_DESCRIPTION=Personal website of Alex Drysdale
API_URL='http://localhost'
BASE_URL='http://localhost'
APP_PORT=5000
APP_BASE_HOST='http://localhost'
API_BASE_HOST='http://localhost'
BIZ_ADDRESS=''
APP_LOGO='https://res.cloudinary.com/nathansweb/image/upload/v1624156571/logos/0001_xwkovt.svg'
API_SECRET=''
```

## Advanced Configuration

cd into `/content` and edit `settings.json` file for application and site wide navigation

## About

The landing page shows the profile and bio suitable for personal website however for company or NGO details is work-in-progress. To edit the content, cd into `/content/profile` and esit  `about.json`.

```
{
    "name": "Alex Drysdale",
    "coverpicture": "https://images.unsplash.com/photo-1520785643438-5bf77931f493?ixlib=rb-=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=8&w=1024&h=1024&q=80",
    "title": "Senior Developer Advocate",
    "bio": "I am an accomplished coder and programmer, and I enjoy using my skills to contribute to the exciting technological advances that happen every day at work. I graduated from the California Institute of Technology in 2016 with a Bachelor's Degree in Software Development. While in school, I earned the 2015 Edmund Gains Award for my exemplary academic performance and leadership skills"
}

```
<alert>
You can use [cloudinary](https://cloudinary.com/), a cloud-based image and video management services
</alert>

