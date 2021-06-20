<template>
<div class="lg:flex">
    <div id="sidebar" class="fixed z-40 inset-0 flex-none h-full bg-black bg-opacity-25 w-full lg:bg-white lg:static lg:h-auto lg:overflow-y-visible lg:pt-0 lg:w-60 xl:w-72 lg:block hidden">
        <div id="navwrapper" class="h-full overflow-y-auto scrolling-touch lg:h-auto lg:block lg:relative lg:sticky lg:bg-transparent overflow-hidden lg:top-18 bg-white mr-24 lg:mr-0">
            <div class="hidden lg:block h-12 pointer-events-none absolute inset-x-0 z-10 bg-gradient-to-b from-white"></div>
            <nav id="nav" class="px-1 pt-6 overflow-y-auto font-medium text-base sm:px-3 xl:px-5 lg:text-sm pb-10 lg:pt-10 lg:pb-14 sticky?lg:h-(screen-18)">
                <LeftNavColor />
                <docnav :data="menu" :title="docs.title"/>
            </nav>
        </div>
    </div>
    <div id="content-wrapper" class="min-w-0 w-full flex-auto lg:static lg:max-h-full lg:overflow-visible">
        <div class="w-full flex">
            <div class="min-w-0 flex-auto px-4 sm:px-6 xl:px-8 pb-24 lg:pb-16">
                <!--Main (Starts)-->
                <main class="mt-16 mx-auto max-w-7xl px-4 sm:mt-10">
                    <!--Body (Starts)-->
                    <div class="inset-0">
                        <article class="max-w-full p-5 prose md:prose-lg">
                            <nuxt-content :document="docs" />
                        </article>
                    </div>
                    <!--Body (Ends)-->
                    <!--Footer (Starts)-->
                        <div class="flex leading-6 font-medium">
                            <NuxtLink v-if="prev" :to="prev.path" class="flex mr-8 transition-colors duration-200 hover:text-gray-900">
                                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
                                </svg>
                                {{ prev.title }}
                            </NuxtLink>
                            <NuxtLink v-if="next" :to="next.path" class="flex text-right ml-auto transition-colors duration-200 hover:text-gray-900">
                                {{ next.title }}
                                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                                </svg>
                            </NuxtLink>
                        </div>
                        <div class="mt-12 border-t border-gray-200 pt-6 text-right"><a class="mt-10 text-sm hover:text-gray-900" href="/">Edit this page on GitHub</a></div>
                    
                    <!--Footer (Ends)-->
                </main>
                <!--Main (Ends)-->
            </div>
            <div class="hidden xl:text-sm xl:block flex-none w-64 pl-8 mr-8">
                <div class="flex flex-col justify-between overflow-y-auto sticky max-h-(screen-18) pt-10 pb-6 top-18">
                    <div class="mb-8">
                        <toc :data="docs.toc" />
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
</template>

<script>
import toc from '@/components/docs/toc.vue'
import docnav from '@/components/docs/docnav.vue'
export default {
    layout: 'public',
    components: {
        toc,
        docnav
    },
    async asyncData({
        $content,
        params
    }) {
        const menu = await $content('docs').only(['title', 'slug', 'path'])
            .sortBy('position', 'asc')
            .fetch()
        const docs = await $content('docs', params.slug).fetch()
        const [prev, next] = await $content('docs', params.path)
            .only(['title', 'slug', 'path'])
            .sortBy('createdAt', 'asc')
            .surround(params.slug)
            .fetch()
        //this.$store.commit('CHANGE_TOC',docs)
        //console.log(JSON.stringify(docs.toc))
        //$nuxt.$emit('evtDocs', docs.toc)
        return {
            menu,
            docs,
            prev,
            next
        }
    }
}
</script>

<style>
/*a.nuxt-link-exact-active {
    @apply bg-green-300;
}

a.nuxt-link-exact-active {
    @apply border-green-900;
}*/
</style>
