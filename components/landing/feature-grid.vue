<template>
<!-- Feature section with grid -->
    <div class="relative bg-white py-16 sm:py-24 lg:py-32">
        <div class="mx-auto max-w-md px-4 text-center sm:max-w-3xl sm:px-6 lg:px-8 lg:max-w-7xl">
            <h2 class="text-base font-semibold tracking-wider text-cyan-600 uppercase">{{menuitems.featuregrid[0].smaltitle}}</h2>
            <p class="mt-2 text-3xl font-extrabold text-gray-900 tracking-tight sm:text-4xl">
                {{menuitems.featuregrid[0].largetitle}}
            </p>
            <p class="mt-5 max-w-prose mx-auto text-xl text-gray-500">
                {{menuitems.featuregrid[0].description}}
            </p>
            <div v-for="(featuregrid,index) in menuitems.featuregrid" :key="featuregrid.smaltitle">
                    <!-- <div class="relative" v-if="featuregrid.children" > -->
                        <div class="mt-12">
                            <div class="grid grid-cols-1 gap-8 sm:grid-cols-2 lg:grid-cols-3">
                                    <div class="flex mx-auto" v-for="childmenu in featuregrid.children" :key="childmenu.smalltitle" >
                                        <div class="pt-6">
                                    <div class="flow-root bg-gray-50 rounded-lg px-6 pb-8">
                                        <div class="-mt-6">
                                            <div>
                                                <span class="inline-flex items-center justify-center p-3 bg-gradient-to-r from-teal-500 to-cyan-600 rounded-md shadow-lg">
                                                    <!-- Heroicon name: outline/cloud-upload -->
                                                    <svg class="h-6 w-6 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" v-bind:d="childmenu.svg" />
                                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" v-bind:d="childmenu.svg2" />
                                                    </svg>
                                                </span>
                                            </div>
                                            <h3 class="mt-8 text-lg font-medium text-gray-900 tracking-tight">{{childmenu.smalltitle}}</h3>
                                            <p class="mt-5 text-base text-gray-500">
                                                {{childmenu.description}}
                                            </p>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    <!-- </div> -->
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