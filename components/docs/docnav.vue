<template>
<div>
    <ul>
        <li class="mt-8">
            <h5 class="px-3 mb-3 lg:mb-3 uppercase tracking-wide font-semibold text-sm lg:text-xs text-gray-900">{{this.title}}</h5>
            <ul>
                <li v-for="item of data" :key="item.path">
                    <NuxtLink :to="item.path" class="px-3 py-2 transition-colors duration-200 relative block hover:text-gray-900 text-gray-500">
                        <span class=" rounded-md absolute inset-0 bg-cyan-50 opacity-0"></span>
                        <span class="relative"> {{ item.title }}</span>
                    </NuxtLink>
                </li>
            </ul>
        </li>
    </ul>
</div>
</template>

<script>
export default {
    name: "doc-nav",
    data() {
        return {

        }
    },
    props: {
        data: Array,
        title: String
    },
    methods: {

    },
    created() {

    },
}
</script>
