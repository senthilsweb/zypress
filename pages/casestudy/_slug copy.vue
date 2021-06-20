<template>
<div class="m-4">
    <div>
        <section class="flex lg:h-screen w-screen lg:overflow-hidden xs:flex-col lg:flex-row">
            <!--Left cover  (starts)-->
            <!--<div class="relative lg:w-1/2 xs:w-full xs:h-84 lg:h-full">
                <img v-if="casestudy.coverimage" class="object-cover w-full" :src="casestudy.coverimage" />
            </div>-->
            <!--Left cover (Ends)-->
            <!--Middle Column (starts)-->
            <div class="flex-1 relative xs:py-8 xs:px-8 lg:w-1/2 xs:w-full">

                <!-- This example requires Tailwind CSS v2.0+ -->
                <div class="bg-gray-300">
                    <div class="max-w-7xl py-3 px-3 sm:px-6 lg:px-4">
                        <div class="flex items-center justify-between flex-wrap">
                            <div class="w-0 flex-1 flex items-center">
                                <p class="font-medium text-green-800 truncate">
                                    <span class="md:hidden">
                                        {{casestudy.title}}
                                    </span>
                                    <span class="hidden md:inline font-bold text-2xl text-green-800">
                                        {{casestudy.title}}
                                    </span>
                                </p>
                            </div>
                            <div class="order-3 mt-2 flex-shrink-0 w-full sm:order-2 sm:mt-0 sm:w-auto">
                                <button type="submit" class="inline-flex justify-center py-2 px-4 border border-transparent text-sm leading-5 font-medium rounded-md text-white bg-green-600 hover:bg-green-500 focus:outline-none focus:border-green-700 focus:shadow-outline-green active:bg-green-700 transition duration-150 ease-in-out">
                                    Download
                                </button>
                            </div>

                        </div>
                    </div>
                </div>

                <!--Body Flex Box # 1 (Starts)-->

                <div class="bg-white border shadow h-full overflow-y-scroll">
                    <div class="flex mx-5 py-4">

                        <article class="prose md:prose-lg">
                            <nuxt-content :document="casestudy" />
                        </article>
                    </div>

                </div>

                <!--Body Flex Box # 1 (Ends)-->
            </div>
            <!--Middle Column (Ends)-->
            <!--Right Column (starts)-->
            <div class="w-1/5 pr-4">
                <div class="shadow bg-white text-base mb-8 ml-4 mt-0 mr-4">
                    <nav>
                        <ul>
                            <li v-for="link of casestudy.toc" :key="link.id" :class="{ 'pl-8': link.depth === 3, 'pl-5 font-bold': link.depth === 2 }" class="group flex items-center px-3 py-2 text-lg leading-8 hover:text-gray-900 hover:bg-gray-100 focus:outline-none focus:bg-gray-200 transition ease-in-out duration-150">
                                <a :href="`#${link.id}`">{{ link.text }}</a>
                            </li>
                        </ul>
                    </nav>
                </div>
            </div>
            <!--Right Column (Ends)-->
        </section>
    </div>

</div>
</template>

<script>
export default {
    components: {

    },
    async asyncData({
        $content,
        params
    }) {
        const casestudy = await $content('casestudies', params.slug).fetch()
        return {
            casestudy
        }
    }
}
</script>
