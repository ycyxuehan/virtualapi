<template>
  <div class="hello" id="service-container">
    <Row>独立API</Row>
    <interface-container :apis="service.apis"></interface-container>
    <Row>分隔符</Row>
    <Row v-for="(group, index) in service.groups" :key="index">
      <Row>API Group {{group.name}}</Row>
      <Row>
        <interface-container :apis="group.apis"></interface-container>
      </Row>
    </Row>
    
  </div>
</template>

<script>
import InterfaceContainer from "./Interface";
export default {
  name: "ServiceContainer",
  components: { InterfaceContainer },
  props: ["service"],
  data() {
    return {
      column: [
        { title: "项目", key: "name" },
        { title: "值", key: "value" }
      ]
    };
  },
  methods: {
    getTable: function(method) {
      let res = [];
      for (let key in method) {
        res.push({ name: key, value: method[key] });
      }
      return res;
    }
  },
  mounted() {}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
