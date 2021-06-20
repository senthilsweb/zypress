<template>
<div>
    <div class="bg-white pt-20 pb-20 px-4 sm:px-6 lg:pt-20 lg:pb-28 lg:px-8">
        <div class="relative max-w-lg mx-auto lg:max-w-7xl">
            <h2 class="text-3xl tracking-tight font-extrabold text-gray-600 sm:text-4xl">
                Recent Posts
            </h2>

            <div class="mt-6 pt-10 grid gap-4 lg:grid-cols-4 lg:gap-x-4 lg:gap-y-8">
                <div v-for="article of articles" :key="article.slug" class="flex flex-col rounded-lg shadow-lg overflow-hidden">
                    <NuxtLink :to="{ name: 'casestudy-slug', params: { slug: article.slug } }">
                        <div class="flex-shrink-0">
                            <img v-if="article.coverimage" class="h-48 w-full object-cover" :src="article.coverimage" />
                        </div>
                    </NuxtLink>
                    <div class="flex-1 bg-white p-6 flex flex-col justify-between">
                        <div class="flex-1">
                            <a href="#" class="inline-block">
                                <span class="inline-flex items-center px-3 py-0.5 rounded-full text-sm font-medium" :class="article.type=='Blog' && 'bg-indigo-100 text-indigo-800' || article.type=='Article' && 'bg-pink-100 text-pink-800' || article.type=='Case Study' && 'bg-green-100 text-green-800'">
                                    {{article.type}}
                                </span>
                            </a>
                            <NuxtLink :to="{ name: 'casestudy-slug', params: { slug: article.slug } }" class="block">
                                <h3 class="mt-2 text-xl leading-7 font-semibold text-gray-900">
                                    {{article.title | prune(25)}}
                                </h3>
                                <p class="mt-3 text-base leading-6 text-gray-500">
                                    {{article.summary | prune(100)}}
                                </p>
                            </NuxtLink>
                        </div>
                        <div class="mt-6 flex items-center">
                            <div class="ml-3">
                                <p class="text-sm leading-5 font-medium text-gray-900">
                                    <a href="#" class="hover:underline">
                                        {{article.author}}
                                    </a>
                                </p>
                                <div class="flex text-sm leading-5 text-gray-500">

                                    {{article.date | moment("DD-MMM-YYYY")}}

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
        </div>
    </div>
</div>
</template>

<script>
export default {
    components: {

    },
    props: {
        articles: Array
    },
    filters: {
        prune(val, len) {
            let n = (len === undefined) ? 100 : len;
            return s(val).prune(n).value()
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
