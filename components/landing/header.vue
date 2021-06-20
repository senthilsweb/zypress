<template>
<header class="">
    <BannerHeaderFlashNews v-if="menuitems.bannerHeaderFlashNews.active" />
    <div class=" bg-gray-900 pt-4 pb-2">
        <nav class="max-w-7xl mx-auto flex items-center justify-between px-4 sm:px-6" aria-label="Global">
            <div class="flex items-center flex-1">
                <div class="flex items-center justify-between w-full md:w-auto">
                    <a href="#">
                        <span class="sr-only">Docs</span>
                        <img tag="img" to="/" alt="Logo" width="200" class="h-8 w-auto sm:h-10" :src="menuitems.branding.logo.source"/>
                    </a>
                    <div @click="toggleMobileMenu" class="-mr-2 flex items-center md:hidden">
                        <button type="button" class="bg-gray-900 rounded-md p-2 inline-flex items-center justify-center text-gray-400 hover:bg-gray-800 focus:outline-none focus:ring-2 focus-ring-inset focus:ring-white" aria-expanded="false">
                            <span class="sr-only">Open main menu</span>
                            <!-- Heroicon name: outline/menu -->
                            <svg class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
                            </svg>
                        </button>
                    </div>
                </div>
                <div class="hidden  space-x-8 md:flex md:ml-10">
                    <div v-for="(menu,index) in menuitems.menus" @mouseover="showMenu(index)" :key="menu.link">
                        <div class="relative" v-if="menu.children">
                            <button type="button" class="uppercase group inline-flex items-center space-x-2 text-base leading-6 font-medium focus:outline-none transition ease-in-out duration-150" :class="textColor=='white' ? 'text-white' : 'text-green-500  hover:text-green-600  focus:text-green-600'">
                                <span>{{menu.title}}</span>
                                <!-- Item active: "text-gray-600", Item inactive: "text-gray-400" -->
                                <svg class="transition ease-in-out duration-150" viewBox="0 0 20 20" fill="currentColor" :class="textColor=='white' ? 'text-white h-5 w-5 group-hover:text-white group-focus:text-white' : 'text-green-400 h-5 w-5 group-hover:text-green-500 group-focus:text-green-500'">
                                    <path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd" />
                                </svg>
                            </button>
                            <!--Main Menu (Start)-->
                            <!--Flyout Menu (Start)-->
                            <div class="absolute z-10 inset-x-0 transform shadow-lg">
                                <div class="absolute inset-0 flex" aria-hidden="true">
                                    <div class="bg-white w-1/2"></div>
                                    <div class="bg-gray-50 w-1/2"></div>
                                </div>
                                <div @focus="showMenu(index)" @mouseleave="hideMenu(index)" v-if="isVisible[index]" :id="'cmenu_'+index" class="absolute left-1/2 transform -translate-x-1/2 mt-3 px-2 w-screen max-w-md sm:px-0 z-50">
                                    <div class="rounded-lg shadow-lg">
                                        <div class="rounded-lg shadow-xs overflow-hidden">
                                            <div class="z-20 relative grid gap-6 bg-white px-5 py-6 sm:gap-8 sm:p-8">
                                                <a :href="childmenu.link" v-for="childmenu in menu.children" :key="childmenu.link" class="-m-3 p-3 flex items-start space-x-4 rounded-lg hover:bg-gray-50 transition ease-in-out duration-150">
                                                    <svg class="flex-shrink-0 h-6 w-6 text-green-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="childmenu.icon" />
                                                    </svg>
                                                    <div class="space-y-1">
                                                        <p class="text-base leading-6 font-medium text-gray-900">
                                                            {{childmenu.title}}
                                                        </p>
                                                    </div>
                                                </a>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <!--Flyout Menu (End)-->
                        </div>
                        <div v-else>
                            <a :href="menu.link" class="uppercase text-base leading-6 font-medium  focus:outline-none transition ease-in-out duration-150" :class="textColor=='white' ? 'text-white' : 'text-green-500  hover:text-green-600  focus:text-green-600'">
                                {{menu.title}}
                            </a>
                        </div>
                        <!--Main Menu (End)-->
                    </div>
                </div>
            </div>
            <div class="hidden md:flex md:items-center md:space-x-6">
                <div class="flex items-center space-x-8 -mt-2">
                    <!--Action button in the header menu (Start)-->
                    <a v-if="menu.hidden==false" 
                        v-for="menu in menuitems.actions" 
                        :key="menu.link" :href="menu.link"
                        @click="$nuxt.$emit('evtShowContactSales')"
                        class="inline-flex items-center px-4 py-2 border border-transparent text-base font-medium  text-white transition ease-in-out duration-150" :class="menu.class">
                        {{menu.title}}
                    </a>

                    <!--Action button in the header menu (End)-->
                </div>

            </div>
        </nav>
    </div>

    <!--
        Mobile menu, show/hide based on menu open state.

        Entering: "duration-150 ease-out"
          From: "opacity-0 scale-95"
          To: "opacity-100 scale-100"
        Leaving: "duration-100 ease-in"
          From: "opacity-100 scale-100"
          To: "opacity-0 scale-95"
      -->
    <div v-if="mobileMenuVisiblity" class="absolute top-0 inset-x-0 p-2 transition transform origin-top md:hidden">
        <div class="rounded-lg shadow-md bg-white ring-1 ring-black ring-opacity-5 overflow-hidden">
            <div class="px-5 pt-4 flex items-center justify-between">
                <div>
                    <img class="h-8 w-auto" :src="require('@/assets/wlogo-alone.svg')" alt="">
                </div>
                <div class="-mr-2">
                    <button type="button" @click="toggleMobileMenu" class="bg-white rounded-md p-2 inline-flex items-center justify-center text-gray-400 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-cyan-600">
                        <span class="sr-only">Close menu</span>
                        <!-- Heroicon name: outline/x -->
                        <svg class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                        </svg>
                    </button>
                </div>
            </div>
            <div class="pt-5 pb-6">
                <div class="px-2 space-y-1">
                    <a v-for="(menu,index) in menuitems.menus" :key="menu.link" :id="menu.title+'_'+index" href="#" class="block px-3 py-2 rounded-md text-base font-medium text-gray-900 hover:bg-gray-50">
                        {{menu.title}}
                    </a>
                </div>
                <!--<div class="mt-6 px-5">
                            <a href="#" class="block text-center w-full py-3 px-4 rounded-md shadow bg-gradient-to-r from-teal-500 to-cyan-600 text-white font-medium hover:from-teal-600 hover:to-cyan-700">Start free trial</a>
                        </div>
                        <div class="mt-6 px-5">
                            <p class="text-center text-base font-medium text-gray-500">Existing customer? <a href="#" class="text-gray-900 hover:underline">Login</a></p>
                        </div>-->
            </div>
        </div>
    </div>
    <contactSales/>
</header>

</template>

<script>
import BannerHeaderFlashNews from '@/components/BannerHeaderFlashNews.vue'
import contactSales from "@/components/landing/contactSales.vue";
import menuitems from "@/store/menuitems.json";

export default {
    name: "header",
    components: {
        BannerHeaderFlashNews,
        contactSales
    },
    data() {
        let isVisible = new Array(20).fill(false);
        return {
            menuitems,
            isVisible,
            isVisibleProfile: false,
            mobileMenuVisiblity: false
        };
    },
    props: {
        showBranding: Boolean,
        showActionButtons: Boolean,
        isAuthenticated: Boolean,
        textColor: String,
        logo: String
    },

    methods: {
        showMenu(index) {
            this.isVisible.forEach((value, index) => {
                this.$set(this.isVisible, index, false)
            })
            this.$set(this.isVisible, index, true)
        },
        hideMenu(index) {
            this.$set(this.isVisible, index, false)
        },
        toggleMobileMenu() {
            this.mobileMenuVisiblity = !this.mobileMenuVisiblity;
        }
    },
    async fetch() {

    }
}
</script>
