# JAM Stack in vue.js/Nuxt.js and Tailwind.css

## Pre-requisites

* Node.JS
* NPM
* VS.code editor


## Build Setup

```bash
# install dependencies
$ npm install

# serve with hot reload at localhost:8000 (or change the port in nust.config.js)
$ npm run dev

# build for production and launch server
$ npm run build


# generate static project
$ npm run generate

#launch server
$ npm run start

#build  docker image

docker build -t web.frontend.nuxtjs:dev .

#run the docker
docker run -it -p 8000:80 web.frontend.nuxtjs:dev

#delete node_module folder and subfolders
rm -r node_modules

#Remove any stopped containers and all unused images (not just dangling images)
docker system prune -a

#Remove dangling images
docker images -f dangling=true

#Remove
docker image prune

```

For detailed explanation on how things work, check out [Nuxt.js docs](https://nuxtjs.org).
