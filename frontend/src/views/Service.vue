<template>
  <div class="home">
    <Header :style="{background: '#fff'}">Service {{service.name}}</Header>
    <service-container :service="service"/>
  </div>
</template>

<script>
// @ is an alias to /src
import ServiceContainer from "@/components/Service.vue";
import {GetService} from "@/api/api.js"
export default {
  name: "Service",
  components: {
    ServiceContainer
  },
  data(){
    return {
      service:{}
    }
  },
  methods: {
    getService:function(){
      let name = this.$route.query.service;
      console.info("actived menu:", name)
      GetService(name).then(res=>{
        console.info("service:", res)
        if(res.code == 0) {
          this.service = res.data;
        }
      })
    }
  },
  mounted(){
    this.getService();
  }
};
</script>
