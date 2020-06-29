<template>
  <div id="menu-container">
    <Menu active-name="1-0" theme="dark" width="auto" :open-names="['1']" @on-select="selectService">
      <Submenu name="1">
        <template slot="title">
          服务
        </template>
        <MenuItem v-for="(svc, index) in services" :key="index" :name='index' >{{svc}}</MenuItem>
      </Submenu>
    </Menu>
  </div>
</template>

<script>
export default {
  name: "MenuContainer",
  props: ['services'],
  data() {
    return {
      activeMenu:-1,
    };
  },
  methods:{
    selectService: function(name) {
      console.info("name", name)
      if(this.activeMenu == name){
        return
      }
      this.activeMenu = name;
      if(name >= 0) {
        //this.activeMenu = this.services[name];
        this.$router.push("/service?service=" + this.services[name])
      }
    }
  },
  mounted(){
      console.info("menus:", this.$route.params)
  }
};
</script>

<style>
</style>