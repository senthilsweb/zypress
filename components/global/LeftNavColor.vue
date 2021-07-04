<template>
<div>
    <ul>
        <li v-if="!item.hidden" v-for="item in data.settings.navigation" :key="item.title">
            <a v-bind:href="item.to" class="flex items-center px-3 hover:text-gray-900 transition-colors duration-200 mb-4">
                <div class="mr-3 rounded-md bg-gradient-to-br" :class="item.class">
                    <svg class="h-6 w-6" viewBox="0 0 24 24">
                        <path v-for="data in item.svg" :key="data.d" :d="data.d" :fill="data.fill"></path>
                    </svg>
                </div>
                {{item.title}}
            </a>
        </li>
    </ul>
</div>
</template>
<script>
export default {
    name: "left-nav",
    data() {
       return {
            
        }
    },
    props: {
        data:Object
    },
}
</script>
