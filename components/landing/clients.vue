<template>
    <div class="bg-green-500">
        <div class="max-w-7xl mx-auto py-16 px-4 sm:py-20 sm:px-6 lg:px-8">
            <h2 class="text-4xl font-extrabold text-gray-900">
                Leading engineering organizations are building with Zynomi
            </h2>
            <div class="flow-root mt-8 lg:mt-10">   
                <div class="-mt-4 -ml-8 flex flex-wrap justify-between lg:-ml-4">
                    <div v-for="client in menuitems.companies" :key="client.name" class="mt-4 ml-8 flex flex-grow flex-shrink-0 lg:flex-grow-0 lg:ml-4">
                        <img class="h-12" v-bind:src="client.logo" v-bind:alt="client.name"> 
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import menuitems from "@/store/menuitems.json";
export default {
    components: {

    },
    data() {
        return {
           menuitems
        };
    },
    methods: {

    },
};
</script>