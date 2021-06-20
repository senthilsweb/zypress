<template>
<div id="app" class="bg-white">
    <div class="relative overflow-hidden">
        <!--<div class="fixed top-0 z-40 lg:z-50 w-full max-w-8xl mx-auto">-->
       <Header/>
       <!--</div>-->
        <main class="flex-1 relative z-0 overflow-y-auto focus:outline-none" tabindex="0" x-data x-init="$el.focus()">
            <div class="max-w-full mx-auto">
            <!-- Replace with your content -->
            <nuxt />
            <!-- /End replace -->
            </div>
        </main>
        <Footer />
    </div>
</div>
</template>

<script>
import Header from '@/components/landing/header.vue'
import Footer from '@/components/landing/Footer.vue'

export default {
    components: {
        Header,
        Footer,
    },
    props: {
       
    },
    data() {
        return {
           
        };
    },
    methods: {
        
    },
};
</script>
