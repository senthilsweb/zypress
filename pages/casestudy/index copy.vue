<template>
<div class="m-4">
    <div>
        <section class="flex lg:h-screen w-screen lg:overflow-hidden xs:flex-col lg:flex-row">
             <!--Left Column (starts)-->
            <div class="w-1/5 pr-4">
               
            </div>
            <!--Left Column (Ends)-->

            <!--Main Content Column (starts)-->
            <div class="flex-1 relative h-full overflow-y-scroll mr-4">
                <!--Body Flex Box # 1 (Starts)-->
                <div class="bg-white border  shadow">
                    <div class="flex mx-5 py-2">
                        <div class="mt-3 grid">
                            <!--Body goes here (Start)-->
                            <div class="mt-12 grid gap-5 max-w-lg mx-auto lg:grid-cols-3 lg:max-w-none">
                                <div v-for="article of articles" :key="article.slug" class="flex flex-col rounded-lg shadow-lg overflow-hidden">
                                    <NuxtLink :to="{ name: 'casestudy-slug', params: { slug: article.slug } }">
                                        <div class="flex-shrink-0">
                                            <img v-if="article.coverimage" class="h-48 w-full object-cover" :src="article.coverimage" />
                                        </div>
                                    </NuxtLink>
                                    <div class="flex-1 bg-white p-6 flex flex-col justify-between">
                                        <div class="flex-1">
                                            <p class="text-sm leading-5 font-medium text-indigo-600">
                                                <a href="#" class="hover:underline">
                                                    tag#1
                                                </a>
                                            </p>
                                            <NuxtLink :to="{ name: 'casestudy-slug', params: { slug: article.slug } }" class="block">
                                                <h3 class="mt-2 text-xl leading-7 font-semibold text-gray-900">
                                                    {{article.title}}
                                                </h3>
                                                <p class="mt-3 text-base leading-6 text-gray-500">
                                                    {{article.summary}}
                                                </p>
                                            </NuxtLink>
                                        </div>
                                        <div class="mt-6 flex items-center">
                                            <div class="ml-3">
                                                <p class="text-sm leading-5 font-medium text-gray-900">
                                                    <a href="#" class="hover:underline">
                                                        zynomi
                                                    </a>
                                                </p>
                                                <div class="flex text-sm leading-5 text-gray-500">

                                                    03-Apr-2021

                                                    <span class="mx-1">
                                                        &middot;
                                                    </span>
                                                    <span>
                                                        5 min read
                                                    </span>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <!--Body goes here (ends)-->
                        </div>
                    </div>
                    <footer class="border-t flex justify-end px-5 py-4">
                        <p>Footer</p>
                    </footer>
                </div>
                <!--Body Flex Box # 1 (Ends)-->
            </div>
            <!--Main Content ColumnColumn (Ends)-->
           
        </section>
    </div>

</div>
</template>

<script>
import SideBarGridVerticalNavigation from "@/components/SideBarGridVerticalNavigation.vue";
export default {
    components: {
        SideBarGridVerticalNavigation
    },
    async asyncData({
        $content,
        params
    }) {
        const articles = await $content('casestudies', params.slug)
            .only(['title', 'summary', 'coverimage', 'slug', 'tags'])
            .sortBy('createdAt', 'desc')
            .fetch()
        return {
            articles
        }
    }
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
