<template>
<div class="hidden px-4 md:flex-1 md:flex md:items-middle md:justify-between md:space-x-12">
    <div v-if="showBranding=='true'" class="flex1 items-center object-none object-center">
        <a href="#">
            <span class="sr-only">Zynomi</span>
            <nuxt-link tag="img" to="/" alt="Logo" width="200" class="h-8 w-auto sm:h-10" :src="require('@/assets/wlogo-full-black.svg')">
            </nuxt-link>
        </a>
    </div>

    <div class="-mr-2 -my-2 md:hidden">
        <button type="button" class="bg-white rounded-md p-2 inline-flex items-center justify-center text-gray-400 hover:text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-green-500" aria-expanded="false">
            <span class="sr-only">Open menu</span>
            <!-- Heroicon name: outline/menu -->
            <svg class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
        </button>
    </div>

    <div class="bg-white ">
        <nav class="flex space-x-10 m-4">
            <div v-for="(menu,index) in menuitems.menus" :key="menu.link">
                <div class="relative" v-if="menu.children">
                    <button @click="toggleMenu(index)" type="button" class="uppercase group
                  inline-flex items-center space-x-2 text-base leading-6 font-medium focus:outline-none
                  transition ease-in-out duration-150" :class="textColor!='green' ? 'text-${textColor}' : 'text-green-500  hover:text-green-600  focus:text-green-600'">
                        <span>{{menu.title}}</span>
                        <!-- Item active: "text-gray-600", Item inactive: "text-gray-400" -->
                        <svg class="transition ease-in-out duration-150" viewBox="0 0 20 20" fill="currentColor" :class="textColor!='green' ? 'text-${textColor} h-5 w-5 group-hover:text-white group-focus:text-white' : 'text-green-400 h-5 w-5 group-hover:text-green-500 group-focus:text-green-500'">
                            <path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd" />
                        </svg>
                    </button>
                    <!--Flyout Menu (Start)-->
                    <div v-if="isVisible[index]" :id="'cmenu_'+index" class="absolute left-1/2 transform -translate-x-1/2 mt-3 px-2 w-screen max-w-md sm:px-0 z-50">
                        <div class="rounded-lg shadow-lg">
                            <div class="rounded-lg shadow-xs overflow-hidden">
                                <div class="z-20 relative grid gap-6 bg-white px-5 py-6 sm:gap-8 sm:p-8">
                                    <!-- suresh -->
                                    <router-link v-bind:to="childmenu.link" v-for="childmenu in menu.children" :key="childmenu.link" class="-m-3 p-3 flex items-start space-x-4 rounded-lg hover:bg-gray-50 transition ease-in-out duration-150">
                                        <svg class="flex-shrink-0 h-6 w-6 text-green-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="childmenu.icon" />
                                        </svg>
                                        <div class="space-y-1">
                                            <p class="text-base leading-6 font-medium text-gray-900">
                                                {{childmenu.title}}
                                            </p>
                                        </div>
                                    </router-link>
                                    <!-- <a :href="childmenu.link" v-for="childmenu in menu.children" :key="childmenu.link" class="-m-3 p-3 flex items-start space-x-4 rounded-lg hover:bg-gray-50 transition ease-in-out duration-150">
                                        <svg class="flex-shrink-0 h-6 w-6 text-green-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="childmenu.icon" />
                                        </svg>
                                        <div class="space-y-1">
                                            <p class="text-base leading-6 font-medium text-gray-900">
                                                {{childmenu.title}}
                                            </p>
                                        </div>
                                    </a> -->
                                </div>
                            </div>
                        </div>
                    </div>
                    <!--Flyout Menu (End)-->
                </div>
                <div v-else>
                    <router-link v-bind:to="menu.link" class="uppercase text-base leading-6 font-medium  focus:outline-none transition ease-in-out duration-150 nav-link"  :class="textColor!='green' ? 'text-${textColor}' : 'text-green-500  hover:text-green-600  focus:text-green-600' ">{{menu.title}}</router-link>
                    <!-- <a :href="menu.link" 
                    @click="toggle(index)"
                    class="uppercase text-base leading-6 font-medium  focus:outline-none transition ease-in-out duration-150"  :class="textColor!='green' ? 'text-${textColor}' : 'text-green-500  hover:text-green-600  focus:text-green-600' ">
                        {{menu.title}}
                    </a> -->
                </div>
            </div>
        </nav>
    </div>
    <div class="flex items-center space-x-8 -mt-2">
        <!--Action button in the header menu (Start)-->
        <span v-if="showActionButtons=='true' && menu.hidden==false" v-for="menu in menuitems.actions" :key="menu.link" class="inline-flex rounded-md shadow-sm">
            <a href="#" class="uppercase inline-flex items-center justify-center px-4 py-2 border border-transparent text-base leading-6 font-medium rounded-md text-white bg-green-600 hover:bg-green-500 focus:outline-none focus:border-green-700 focus:shadow-outline-green active:bg-green-700 transition ease-in-out duration-150">
                {{menu.title}}
            </a>
        </span>
        <!--Action button in the header menu (End)-->
        <!--Profile (Start)-->
        <button v-if="isAuthenticated=='true'" class="p-1 border-2 border-transparent text-gray-400 rounded-full hover:text-gray-500 focus:outline-none focus:text-gray-500 focus:bg-gray-100 transition duration-150 ease-in-out" aria-label="Notifications">
            <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
            </svg>
        </button>

        <!-- Profile dropdown -->
        <div v-if="isAuthenticated=='true'" @click="toggleProfile" class="ml-3 relative pr-10">
            <div>
                <button class="flex text-sm border-2 border-transparent rounded-full focus:outline-none focus:border-gray-300 transition duration-150 ease-in-out" id="user-menu" aria-label="User menu" aria-haspopup="true">
                    <img class="h-8 w-8 rounded-full" src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80" alt="">
                </button>
            </div>
            <!--
            Profile dropdown panel, show/hide based on dropdown state.

            Entering: "transition ease-out duration-200"
              From: "transform opacity-0 scale-95"
              To: "transform opacity-100 scale-100"
            Leaving: "transition ease-in duration-75"
              From: "transform opacity-100 scale-100"
              To: "transform opacity-0 scale-95"
          -->
            <div v-if="isVisibleProfile" class="origin-top-right absolute right-0 mt-2 w-48 shadow-lg">
                <div class="py-1 rounded-md bg-white shadow-xs" role="menu" aria-orientation="vertical" aria-labelledby="user-menu">
                    <a href="/settings" class="block px-4 py-2 text-sm leading-5 text-gray-700 hover:bg-gray-100 focus:outline-none focus:bg-gray-100 transition duration-150 ease-in-out" role="menuitem">Your Profile</a>
                    <a href="/settings" class="block px-4 py-2 text-sm leading-5 text-gray-700 hover:bg-gray-100 focus:outline-none focus:bg-gray-100 transition duration-150 ease-in-out" role="menuitem">Settings</a>
                    <a href="/signin" class="block px-4 py-2 text-sm leading-5 text-gray-700 hover:bg-gray-100 focus:outline-none focus:bg-gray-100 transition duration-150 ease-in-out" role="menuitem" @click="logout">Sign out</a>
                </div>
            </div>
        </div>

        <!--Profile (End)-->
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
    <div class="absolute top-0 inset-x-0 p-2 transition transform origin-top md:hidden">
        <div class="rounded-lg shadow-md bg-white ring-1 ring-black ring-opacity-5 overflow-hidden">
            <div class="px-5 pt-4 flex items-center justify-between">
                <div>
                    <nuxt-link tag="img" to="/" alt="Logo" width="200" class="pt-2" :src="require('@/assets/zlogo.svg')">
                    </nuxt-link>
                </div>
                <div class="-mr-2">
                    <button type="button" class="bg-white rounded-md p-2 inline-flex items-center justify-center text-gray-400 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-cyan-600">
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
                    <a v-for="(menu,index) in menuitems.menus" :key="menu.link" href="#" class="block px-3 py-2 rounded-md text-base font-medium text-green-900 hover:bg-green-50">
                        {{menu.title}}
                    </a>
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
    props: {
        showBranding: String,
        showActionButtons: String,
        isAuthenticated: String,
        textColor: String,
        logo: String
    },
    data() {
        let isVisible = new Array(20).fill(false);
        return {
            menuitems,
            isVisible,
            isVisibleProfile: false,
            activeIndex: null
        };
    },
    methods: {
        toggleMenu(index) {
            this.$set(this.isVisible, index, !this.isVisible[index])
            const toggleMenuIndex = index
            this.isVisible.forEach((value, index) => {
                if (index == toggleMenuIndex && this.isVisible[index] == true)
                    this.$set(this.isVisible, index, true)
                else
                    this.$set(this.isVisible, index, false)
            });
        },
        toggleProfile() {
            this.isVisibleProfile = !this.isVisibleProfile;
        },
        logout() {
            store.dispatch('logout')
        },
        toggle(index){
            alert(index)
            this.textColor ='green'
            //this.activeIndex = index
        }
    },
};
</script>
