<template>
<div>
    <main class="mt-16 mx-auto max-w-7xl px-4 sm:mt-10">
        <dochero class="pb-10"/>
        <actionpanel />
        <!--Footer (Starts)-->
        <div class="mt-12 border-t border-gray-200 pt-6 text-right">
            <a class="mt-10 text-sm hover:text-gray-900" href="/">
                Edit this page on GitHub
            </a>
        </div>
        <!--Footer (Ends)-->
    </main>
</div>
</template>

<script>
import dochero from '@/components/docs/dochero.vue'
import actionpanel from '@/components/docs/actionpanel.vue'
export default {
    layout: 'doc2',
    components: {
        dochero,
        actionpanel
    }
}
</script>

<style>
a.nuxt-link-exact-active {
    @apply bg-green-300;
}

a.nuxt-link-exact-active {
    @apply border-green-900;
}
</style>
