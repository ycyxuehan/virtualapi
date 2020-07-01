<template>
  <div class="layout" id="app">
    <Sider :style="{position: 'fixed', height: '100vh', left: 0, overflow: 'auto'}">
      <menu-container :services="serviceNames" />
    </Sider>
    <Layout :style="{marginLeft: '200px'}">
      <router-view />
    </Layout>
  </div>
</template>
<script>
import { GetServices } from "./api/api";
import MenuContainer from "./components/Menu.vue";
export default {
  name: "App",
  components: {
    MenuContainer
  },
  data() {
    return {
      serviceNames: []
    };
  },
  methods: {
    getServices: function() {
      GetServices().then(res => {
        if (res.code == 0) {
          this.serviceNames = res.data;
        }
        console.info("services:", this.serviceNames);
      });
    }
  },
  mounted() {
    this.getServices();
  }
};
</script>
<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.layout {
  border: 1px solid #d7dde4;
  background: #f5f7f9;
  position: relative;
  border-radius: 4px;
  overflow: hidden;
}
.layout-header-bar {
  background: #fff;
  box-shadow: 0 1px 1px rgba(0, 0, 0, 0.1);
}
</style>
