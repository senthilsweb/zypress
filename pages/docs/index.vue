<template>
<div class="lg:flex">
    <div id="sidebar" class="fixed z-40 inset-0 flex-none h-full bg-black bg-opacity-25 w-full lg:bg-white lg:static lg:h-auto lg:overflow-y-visible lg:pt-0 lg:w-60 xl:w-72 lg:block" :class="mobileNav==false? 'hidden' : 'block'">
        <div id="navwrapper" class="h-full overflow-y-auto scrolling-touch lg:h-auto lg:block lg:relative lg:sticky lg:bg-transparent overflow-hidden lg:top-18 bg-white mr-24 lg:mr-0">
            <div :class="mobileNav==false? 'hidden' : 'block'" class="lg:block h-12 pointer-events-none absolute inset-x-0 z-10 bg-gradient-to-b from-white"></div>
            <nav id="nav" class="px-1 pt-6 overflow-y-auto font-medium text-base sm:px-3 xl:px-5 lg:text-sm pb-10 lg:pt-10 lg:pb-14 sticky?lg:h-(screen-18)">
                <!--<LeftNavColor/>-->
                <LeftNavColor :data="links" />
            </nav>
        </div>
    </div>
    <div id="content-wrapper" class="m-5 pt-5 min-w-0 w-full flex-auto lg:static lg:max-h-full lg:overflow-visible">
        <h1 class="text-5xl leading-none font-extrabold text-gray-900 tracking-tight mb-4">
            <span class="block xl:inline">Getting started with </span>
            <span class="block text-teal-400 xl:inline">Zypress</span>
        </h1>
        <p class="text-2xl tracking-tight text-gray-500 mb-10">Learn how to build elegant blog and documentation websites. </p>
        <div class="w-full flex">
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6 xl:gap-8">
                <section class="flex">
                    <div class="w-full relative text-white overflow-hidden rounded-3xl flex shadow-lg">
                        <div class="w-full flex md:flex-col bg-gradient-to-br from-purple-500 to-indigo-500">
                            <div class="sm:max-w-sm sm:flex-none md:w-auto md:flex-auto flex flex-col items-start relative z-10 p-6 xl:p-8">
                                <h2 class="text-xl font-semibold mb-2 text-shadow">Try it in the IDE</h2>
                                <p class="font-medium text-violet-100 text-shadow mb-4">Learn how to get zypress set up to publish your website.</p>
                                <a class="mt-auto bg-indigo-800 bg-opacity-50 hover:bg-opacity-75 transition-colors duration-200 rounded-xl font-semibold py-2 px-4 inline-flex" href="/docs/getting-started">Start Building</a>
                            </div>
                            <div class="docs_image__1HDuG relative md:pl-6 xl:pl-8 hidden sm:block">
                            </div>
                        </div>
                        <div class="absolute bottom-0 left-0 right-0 h-20 hidden sm:block" style="background: linear-gradient(to top, rgb(135, 94, 245), rgba(135, 94, 245, 0));"></div>
                    </div>
                </section>
                <section class="flex">
                    <div class="w-full relative text-white overflow-hidden rounded-3xl flex shadow-lg">
                        <div class="w-full flex md:flex-col bg-gradient-to-br from-pink-500 to-rose-500">
                            <div class="sm:max-w-sm sm:flex-none md:w-auto md:flex-auto flex flex-col items-start relative z-10 p-6 xl:p-8">
                                <h2 class="text-xl font-semibold mb-2 text-shadow">Content Writing</h2>
                                <p class="font-medium text-rose-100 text-shadow mb-4">Customize the website branding and create content.</p>
                                <a class="mt-auto bg-rose-900 bg-opacity-50 hover:bg-opacity-75 transition-colors duration-200 rounded-xl font-semibold py-2 px-4 inline-flex" href="/docs/website-settings">Start Customization</a>
                            </div>
                            <div class="docs_image__1HDuG relative md:pl-6 xl:pl-8 hidden sm:block">
                            </div>
                        </div>
                        <div class="absolute bottom-0 left-0 right-0 h-20 bg-gradient-to-t from-rose-500 hidden sm:block"></div>
                    </div>
                </section>
                <section class="flex">
                    <div class="w-full relative text-white overflow-hidden rounded-3xl flex shadow-lg">
                        <div class="w-full flex md:flex-col bg-gradient-to-br from-yellow-400 to-orange-500">
                            <div class="sm:max-w-sm sm:flex-none md:w-auto md:flex-auto flex flex-col items-start relative z-10 p-6 xl:p-8">
                                <h2 class="text-xl font-semibold mb-2 text-shadow">Go Live</h2>
                                <p class="font-medium text-amber-100 text-shadow mb-4">Learn how to deploy your website in surge.sh</p>
                                <a class="mt-auto bg-orange-600 bg-opacity-50 hover:bg-opacity-75 transition-colors duration-200 rounded-xl font-semibold py-2 px-4 inline-flex" href="/docs/website-settings">Start Publishing</a>
                            </div>
                            <div class="docs_image__1HDuG relative hidden sm:block">
                                <div class="absolute left-2 bottom-3 xl:bottom-5">
                                </div>
                            </div>
                        </div>
                        <div class="absolute bottom-0 left-0 right-0 h-20 bg-gradient-to-t from-orange-500 hidden sm:block"></div>
                    </div>
                </section>
            </div>
        </div>
    </div>
</div>
</template>

<script>
import docnav from '@/components/docs/docnav.vue'
import {
    mapState
} from 'vuex'
export default {
    layout: 'public',
    components: {
        docnav
    },
    data() {
        return {
            mobileNav: false,
        }
    },
    created() {
        this.$nuxt.$on('evtFloatingButtonClick', (data) => {
            this.mobileNav = !this.mobileNav
        });
    },
    beforeDestroy() {
        this.$nuxt.$off('evtFloatingButtonClick');
    },
    async asyncData({
        $content
    }) {
        const links = await $content('settings').fetch()
        return {
            links
        }
    },
}
</script>
