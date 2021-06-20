<template>
<div>

    <!-- 3 column wrapper -->
    <div class="flex-grow w-full max-w-8xl mx-auto xl:px-8 lg:flex">
        <!-- Left sidebar & main wrapper -->
        <div class="flex-1 min-w-0 bg-white xl:flex">
            <div class="border-b border-gray-200 xl:border-b-0 xl:flex-shrink-0 xl:w-64 xl:border-r xl:border-gray-200 bg-white">
                <div class="h-full pl-4 pr-6 py-6 sm:pl-6 lg:pl-8 xl:pl-0">
                    <!-- Start left column area -->
                    <div class="h-full relative" style="min-height: 12rem;">
                        <div class="absolute inset-0">
                            <nav class="px-2 space-y-1 pt-5" id="doc-left-nav">
                                <NuxtLink v-for="item of menu" :key="item.path" :to="item.path" class="border-transparent text-gray-500 hover:bg-green-400 hover:text-gray-900 group flex items-center px-3 py-2 text-md font-medium border-l-4">
                                    <span class="truncate">
                                        {{ item.title }}
                                    </span>
                                </NuxtLink>
                            </nav>
                        </div>
                    </div>
                    <!-- End left column area -->
                </div>
            </div>

            <div class="bg-white lg:min-w-0 lg:flex-1">
                <div class="h-full py-2 px-4 sm:px-6 lg:px-8">
                    <ul class="space-y-1">
                        <li class="overflow-hidden px-6">
                            <div class="min-w-0 flex-1 md:px-8 lg:px-0 xl:col-span-6">
                                <div class="flex items-center px-6 py-4 md:max-w-3xl md:mx-auto lg:max-w-none lg:mx-0 xl:px-0">
                                    <div class="w-full">
                                        <label for="search" class="sr-only">Search</label>
                                        <div class="relative">
                                            <div class="pointer-events-none absolute inset-y-0 left-0 pl-3 flex items-center">
                                                <!-- Heroicon name: solid/search -->
                                                <svg class="h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                                                    <path fill-rule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clip-rule="evenodd" />
                                                </svg>
                                            </div>
                                            <input id="search" name="search" class="block w-full bg-white border border-gray-300 rounded-md py-4 pl-10 pr-3 text-3xl placeholder-gray-500 focus:outline-none focus:text-gray-900 focus:placeholder-gray-400 focus:ring-1 focus:ring-rose-500 focus:border-rose-500 sm:text-sm" placeholder="Search" type="search">
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </li>
                        <li class="overflow-hidden px-6 py-4">
                            <!-- Start main area-->
                            <div class="border relative" style="min-height: 36rem;">
                                <div class="flex-1 max-h-screen xl:overflow-y-auto">
                                    <div class="inset-0">
                                        <article class="max-w-full p-5 prose md:prose-lg">
                                            <nuxt-content :document="docs" />
                                        </article>
                                    </div>
                                </div>
                            </div>
                            <!-- End main area -->
                        </li>
                        <li class="overflow-hidden py-4">
                            <div class="border-b border-gray-200 px-4 py-4 sm:flex sm:items-center sm:justify-between sm:px-6 lg:px-8">
                                <div class="flex-1 min-w-0">
                                    <NuxtLink v-if="prev" :to="prev.path" class="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium text-gray-600 bg-green-300 hover:bg-green-400 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500">
                                        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
                                        </svg>
                                        {{ prev.title }}
                                    </NuxtLink>
                                </div>
                                <div class="mt-4 flex sm:mt-0 sm:ml-4">
                                    <NuxtLink v-if="next" :to="next.path" class="inline-flex object-none object-right-top items-center px-6 py-3 border border-transparent text-base font-medium text-gray-600 bg-green-300 hover:bg-green-400 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500">
                                        {{ next.title }}
                                        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                                        </svg>
                                    </NuxtLink>
                                </div>
                            </div>
                        </li>
                    </ul>
                </div>
            </div>
        </div>

        <div class="pr-4 sm:pr-6 lg:pr-8 lg:flex-shrink-0 lg:border-l xl:pr-0">
            <div class="h-full pl-6 py-6 lg:w-80">
                <!-- Start right column area -->
                <div class="h-full relative" style="min-height: 16rem;">
                    <div class="absolute inset-0">
                        <div class="text-sm text-gray-600 pb-2 pl-2 uppercase">On this page</div>
                        <div class="flex-1 max-h-screen xl:overflow-y-auto">
                            <nav>
                                <ul>
                                    <li v-for="link of docs.toc" :key="link.id" :class="{ 'pl-8': link.depth === 3, 'pl-5 font-bold': link.depth === 2 }" class="group flex items-center px-3 text-sm leading-8 hover:text-green-900 hover:bg-green-100 focus:outline-none focus:bg-green-200 transition ease-in-out duration-150">
                                        <a :href="`#${link.id}`">{{ link.text }}</a>
                                    </li>
                                </ul>
                            </nav>
                        </div>
                    </div>
                </div>
                <!-- End right column area -->
            </div>
        </div>
    </div>

</div>
</template>

<script>
export default {
    layout: 'docs',
    components: {

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
a.nuxt-link-exact-active {
    @apply bg-green-300;
}
a.nuxt-link-exact-active {
    @apply border-green-900;
}
</style>
