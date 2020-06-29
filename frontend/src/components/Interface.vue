<template>
  <div class="hello" id="interface-container">
        <Row v-for="(api, index) in apis" :key="index">
      API {{api.name}}
      <Tag v-for="(method, index) in api.methods" :key="index" color="blue">{{method.method}}</Tag>
      <Card style v-for="(method, index) in api.methods" :key="index">
        <div style="text-align:center">
          <Table :columns="column" :data="getTable(method)"></Table>
        </div>
      </Card>
    </Row>
  </div>
</template>

<script>
export default {
  name: "InterfaceContainer",
  props: ["apis"],
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
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
