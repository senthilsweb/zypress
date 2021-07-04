<template>
<div class="lg:flex">
    <div id="sidebar" class="fixed z-40 inset-0 flex-none h-full bg-black bg-opacity-25 w-full lg:bg-white lg:static lg:h-auto lg:overflow-y-visible lg:pt-0 lg:w-60 xl:w-72 lg:block hidden">
        <div id="navwrapper" class="h-full overflow-y-auto scrolling-touch lg:h-auto lg:block lg:relative lg:sticky lg:bg-transparent overflow-hidden lg:top-18 bg-white mr-24 lg:mr-0">
            <div class="hidden lg:block h-12 pointer-events-none absolute inset-x-0 z-10 bg-gradient-to-b from-white"></div>
            <nav id="nav" class="px-1 pt-6 overflow-y-auto font-medium text-base sm:px-3 xl:px-5 lg:text-sm pb-10 lg:pt-10 lg:pb-14 sticky?lg:h-(screen-18)">
                <LeftNavColor :data="settings"/>
            </nav>
        </div>
    </div>
    <div id="content-wrapper" class="min-w-0 w-full flex-auto lg:static lg:max-h-full lg:overflow-visible">
        <div class="w-full flex">
            <div class="min-w-0 flex-auto px-4 sm:px-6 xl:px-8 pb-24 lg:pb-16">
                <!--Main (Starts)-->
                <main class="mx-auto max-w-7xl">
                    <!--Body (Starts)-->
                    <div class="inset-0">
                        <div v-for="article of articles" :key="article.slug" class="p-2">
                            <NuxtLink :to="'blog/' + article.slug" class="block">
                                <p class="text-xl font-semibold text-gray-900">
                                    {{article.title}}
                                </p>
                                <p class="mt-3 text-base text-gray-500">
                                    {{article.description}}
                                </p>
                            </NuxtLink>
                            <div class="mt-6 flex items-center">
                                <div class="flex-shrink-0">
                                    <a href="#">
                                        <span class="sr-only">{{article.author}}</span>
                                        <img class="h-10 w-10 rounded-full" :src="article.avatar" alt="">
                                    </a>
                                </div>
                                <div class="ml-3">
                                    <p class="text-sm font-medium text-gray-900">
                                        <a href="/">
                                            {{article.author}}
                                        </a>
                                    </p>
                                    <div class="flex space-x-1 text-sm text-gray-500">
                                        <time datetime="article.createdAt">
                                            {{article.createdAt | moment("DD-MMM-YYYY")}}
                                        </time>
                                        <span aria-hidden="true">
                                            &middot;
                                        </span>
                                        <!--<ReadingTime :content="article.body"></ReadingTime>-->
                                        <span aria-hidden="true">
                                            &middot;
                                        </span>
                                        <span>
                                            <span class="inline-flex items-center px-3 py-0.5 rounded-full text-sm font-medium" :class="article.type=='Blog' && 'bg-indigo-100 text-indigo-800' || article.type=='Article' && 'bg-pink-100 text-pink-800' || article.type=='Case Study' && 'bg-green-100 text-green-800'">
                                                {{article.type}}
                                            </span>
                                        </span>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <!--Body (Ends)-->
                    <!--Footer (Starts)-->
                    <div class="mt-12 border-t border-gray-200 pt-6 text-right"><a class="mt-10 text-sm hover:text-gray-900" href="/">Edit this page on GitHub</a></div>
                    <!--Footer (Ends)-->
                </main>
                <!--Main (Ends)-->
            </div>
            <div class="hidden xl:text-sm xl:block flex-none w-64 mr-8">
                <about :aboutus="about" :profile="about" />
            </div>
        </div>
    </div>
</div>
</template>

<script>
import {
    mapState
} from 'vuex'
export default {
    layout: 'public',
    components: {

    },
    async asyncData({
        $content,
        params
    }) {
        const articles = await $content('blog', params.slug)
            .only(['title', 'description', 'coverimage', 'slug', 'tags', 'type', 'author', 'avatar', 'createdAt'])
            .sortBy('createdAt', 'desc')
            .fetch()
        const teams = await $content('profile/teams').fetch()
        const about = await $content('profile/about').fetch()
        return {
            articles,
            teams,
            about
        }
    },
    async fetch({
        store,
        error
    }) {
        try {
            await store.dispatch('settings/fetchSettings')
        } catch (e) {}
    },
    computed: mapState({
        settings: state => state.settings.settings
    }),
    // call fetch only on client-side
    fetchOnServer: false
}
</script>

<style>
.article-card {
    border-radius: 8px;
}

.article-card a {
    background-color: #fff;
    border-radius: 8px;
}

.article-card img div {
    border-radius: 8px 0 0 8px;
}
</style>
