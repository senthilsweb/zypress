<template>
<section v-if="isContactFormVisible" class="z-10 fixed inset-0 overflow-hidden" aria-labelledby="slide-over-title" role="dialog" aria-modal="true">
    <div class="absolute inset-0 overflow-hidden">
        <div class="absolute inset-0" aria-hidden="true"></div>
        <div class="absolute inset-y-0 right-0 pl-10 max-w-full flex">
            <div class="w-screen max-w-md">
                <div class="h-full divide-y divide-gray-200 flex flex-col bg-white shadow-xl">
                    <div class="flex-1 h-0 overflow-y-auto">
                        <header class="space-y-1 py-6 px-4 bg-green-700 sm:px-6">
                            <div class="flex items-center justify-between space-x-3">
                                <h2 class="text-lg leading-7 font-medium text-white">
                                    Contact Sales
                                </h2>
                                <div class="h-7 flex items-center">
                                    <button v-on:click="isContactFormVisible=false" aria-label="Close panel" class="text-green-200 hover:text-white transition ease-in-out duration-150">
                                        <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                                        </svg>
                                    </button>
                                </div>
                            </div>
                            <div>
                                <p class="text-sm leading-5 text-green-300">
                                    Get a live demonstration from a Marketing & Platform specialist and see first-hand how the zyrpa Platform can exponentially increase your company's business value.
                                </p>
                            </div>
                        </header>
                        <div class="flex-1 flex flex-col justify-between">
                            <div class="px-4 divide-y divide-gray-200 sm:px-6">
                                <div class="space-y-6 pt-6 pb-5">
                                    <form name="frmWebsettings" id="frmWebsettings" class="space-y-8 divide-y divide-gray-200" @submit.prevent="SaveAndUpdate">
                                        <div class="sm:overflow-hidden">
                                            <div class="bg-white space-y-6">
                                                <div class="grid grid-cols-1 gap-y-6 gap-x-4 sm:grid-cols-6">
                                                    <div class="sm:col-span-3">
                                                        <label for="first_name" class="block text-sm font-medium text-gray-700">First name</label>
                                                        <div class="mt-1">
                                                            <input type="text" name="first_name" id="first_name" autocomplete="given-name" class="py-3 px-4 block w-full shadow-sm focus:ring-green-500 focus:border-green-500 border-gray-300">
                                                        </div>
                                                    </div>
                                                    <div class="sm:col-span-3">
                                                        <label for="last_name" class="block text-sm font-medium text-gray-700">Last name</label>
                                                        <div class="mt-1">
                                                            <input type="text" name="last_name" id="last_name" autocomplete="family-name" class="py-3 px-4 block w-full shadow-sm focus:ring-green-500 focus:border-green-500 border-gray-300">
                                                        </div>
                                                    </div>
                                                    <div class="sm:col-span-6">
                                                        <label for="name" class="block text-sm font-medium text-gray-700">
                                                            Company
                                                        </label>
                                                        <div class="mt-1 rounded-md shadow-sm">
                                                            <input type="text" id="company" name="company" autocomplete="name" class="py-3 px-4 flex-1 focus:ring-green-500 focus:border-green-500 block w-full min-w-0 sm:text-sm border-gray-300">
                                                        </div>
                                                    </div>
                                                    <div class="sm:col-span-6">
                                                        <label for="Email" class="block text-sm font-medium text-gray-700">
                                                            Email
                                                        </label>
                                                        <div class="mt-1 rounded-md shadow-sm">
                                                            <input type="text" id="Email" name="Email" autocomplete="name" class="py-3 px-4 flex-1 focus:ring-green-500 focus:border-green-500 block w-full min-w-0 sm:text-sm border-gray-300">
                                                        </div>
                                                    </div>
                                                    <div class="sm:col-span-6">
                                                        <label for="phone_number" class="block text-sm font-medium text-gray-700">Phone Number</label>
                                                        <div class="mt-1 relative rounded-md shadow-sm">
                                                            <div class="absolute inset-y-0 left-0 flex items-center">
                                                                <label for="country" class="sr-only">Country</label>
                                                                <select id="country" name="country" class="h-full py-0 pl-4 pr-8 border-transparent bg-transparent text-gray-500 focus:ring-green-500 focus:border-green-500">
                                                                    <option>US</option>
                                                                    <option>CA</option>
                                                                    <option>EU</option>
                                                                </select>
                                                            </div>
                                                            <input type="text" name="phone_number" id="phone_number" autocomplete="tel" class="py-3 px-4 block w-full pl-20 focus:ring-green-500 focus:border-green-500 border-gray-300" placeholder="+1 (555) 987-6543">
                                                        </div>
                                                    </div>
                                                    <div class="sm:col-span-6">
                                                        <label for="about" class="block text-sm font-medium text-gray-700">
                                                            Description
                                                        </label>
                                                        <div class="mt-1">
                                                            <textarea id="description" name="description" rows="3" class="shadow-sm focus:ring-green-500 focus:border-green-500 block w-full sm:text-sm border-gray-300"></textarea>
                                                        </div>
                                                        <p class="mt-2 text-sm text-gray-500">Write a few sentences about the requirements</p>
                                                    </div>
                                                    <div class="sm:col-span-6">
                                                        <Dropdownlist name="Product" label="Interested in" :data="{'data' : [{'code' : 'zyrpa-cloud', 'name':'Zyrpa Cloud'},{'code' : 'zyrpa-api', 'name':'Zyrpa api'},{'code' : 'zyrpa-bot', 'name':'Zyrpa Bot'},{'code' : 'Other', 'name':'Other'}]}" />
                                                    </div>
                                                    <div class="sm:col-span-6">
                                                        <div class="flex items-start">
                                                            <div class="flex-shrink-0">
                                                                <!-- Enabled: "bg-green-600", Not Enabled: "bg-gray-200" -->
                                                                <button type="button" class="bg-gray-200 relative inline-flex flex-shrink-0 h-6 w-11 border-2 border-transparent rounded-full cursor-pointer transition-colors ease-in-out duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500" role="switch" aria-checked="false">
                                                                    <span class="sr-only">Agree to policies</span>
                                                                    <!-- Enabled: "translate-x-5", Not Enabled: "translate-x-0" -->
                                                                    <span aria-hidden="true" class="translate-x-0 inline-block h-5 w-5 rounded-full bg-white shadow transform ring-0 transition ease-in-out duration-200"></span>
                                                                </button>
                                                            </div>
                                                            <div class="ml-3">
                                                                <p class="text-base text-gray-500">
                                                                    By selecting this, you agree to the
                                                                    <a href="#" class="font-medium text-gray-700 underline">Privacy Policy</a>
                                                                    and
                                                                    <a href="#" class="font-medium text-gray-700 underline">Cookie Policy</a>.
                                                                </p>
                                                            </div>
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>

                                        </div>
                                    </form>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="w-full flex-shrink-0 px-4 py-4 space-x-4 flex justify-end">
                        <span class="inline-flex rounded-md shadow-sm">
                            <button type="button" class="py-2 px-4 border border-gray-300 text-sm leading-5 font-medium text-gray-700 hover:text-gray-500 focus:outline-none focus:border-blue-300 focus:shadow-outline-blue active:bg-gray-50 active:text-gray-800 transition duration-150 ease-in-out">
                                Reset
                            </button>
                        </span>
                        <span class="inline-flex rounded-md shadow-sm">
                            <button v-on:click="sendContactUs" type="submit" class="inline-flex justify-center py-2 px-4 border border-transparent text-sm leading-5 font-medium  text-white bg-green-600 hover:bg-green-500 focus:outline-none focus:border-green-700 focus:shadow-outline-green active:bg-green-700 transition duration-150 ease-in-out">
                                Save
                            </button>
                        </span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</section>
</template>

<script>
import Dropdownlist from "@/components/Dropdownlist.vue";
export default {
    components: {
        Dropdownlist
    },
    data() {
        return {
            isContactFormVisible: false
        }
    },
    methods: {

    },
    computed: {

    },
    mounted() {

    },
    created() {
        this.$nuxt.$on('evtShowContactSales', (data) => {
            this.isContactFormVisible = true;
        });
    },
    beforeDestroy() {
        this.$nuxt.$off('evtShowContactSales');
    }
};
</script>
