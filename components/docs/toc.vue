<template>
<div>
    <h5 class="text-gray-900 uppercase tracking-wide font-semibold mb-3 text-sm lg:text-xs">On this page</h5>
    <ul>
        <li v-for="link of data" :key="link.id" :class="{ 'pl-8': link.depth === 3, 'pl-5 font-bold': link.depth === 2 }" class="group flex items-center px-3 text-sm leading-8 hover:text-gray-900 text-gray-500">
            <a :href="`#${link.id}`">{{ link.text }}</a>
        </li>
    </ul>
</div>
</template>

<script>
export default {
    name: "toc",
    data() {
        return {
            //data:this.$store.state.toc
        }
    },
    props: {
        data: Array
    },
    methods: {

    }, 
    created() {
        
    },
}
</script>
