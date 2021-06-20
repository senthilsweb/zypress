<template>
<div class="bg-gray-900">
  <div class="max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:py-16 lg:px-8">
    <div class="xl:grid xl:grid-cols-3 xl:gap-8">
      
      <!--Footer Links (Start)-->
      
      <!--Footer Links (end)-->
    
      <!--Newsletter-->
   
       <!--Newsletter-->
    </div>
    <div class="mt-8 border-t border-white pt-8 md:flex md:items-center md:justify-between">
      <div class="flex md:order-2">
        <a href="#" class="ml-6 text-white hover:text-white" v-for="sm in menuitems.socialmedia" :key="sm.title">
          <span class="sr-only"> {{sm.title}} </span>
            <svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
              <path :fill-rule="sm.rule" :d="sm.icon" :clip-rule="sm.rule" />
            </svg>
        </a>
      </div>
      <p class="mt-8 text-base leading-6 text-white md:mt-0 md:order-1">
        copyright
      </p>
    </div>
  </div>
  <BannerFooterCookie v-if="menuitems.bannerHeaderFlashNews.active" />
</div>

</template>
<script>
import menuitems from "@/store/menuitems.json";

import BannerFooterCookie from '@/components/BannerFooterCookie.vue'
export default {
  components: {
   BannerFooterCookie
  },
  data() {
    return {
      menuitems
    };
  },
  methods: {
    classNameByIndex: function (index) {
        //return index % 2 == (0 || 1) ? 'evenodd' : '';
        return 'evenodd';
}
  },
};
</script>
