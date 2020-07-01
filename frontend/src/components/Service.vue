<template>
  <div class="hello" id="service-container">
    <Card class="group-card" dis-hover>
      <Row slot="title">
        <Col span="22">
          <h3>API</h3>
        </Col>
        <Col span="2">
          <Button type="primary" style="width:80%;" @click="showAddApiDlg('')">添加API</Button>
        </Col>
      </Row>
      <!-- <p>独立API</p> -->
      <interface-container :apis="service.apis" :service="service.name" :group="''"></interface-container>
    </Card>
    <!-- <Row class="separator"></Row> -->
    <Row class="group-title">
      <Col span="22">
        <h3>已分组API</h3>
      </Col>
      <Col span="2">
        <Button type="primary" style="width:80%;" @click="showAddGoupDlg()">添加组</Button>
      </Col>
    </Row>
    <Row v-for="(group, index) in service.groups" :key="index">
      <Card class="group-card" dis-hover>
        <Row slot="title">
          <Col span="22">
            <h3>API组{{group.name}}</h3>
          </Col>
          <Col span="2">
            <Button type="primary" style="width:80%;" @click="showAddApiDlg(group.name)">添加API</Button>
          </Col>
        </Row>
        <interface-container :apis="group.apis" :service="service.name" :group="group.name"></interface-container>
      </Card>
    </Row>
    <Modal v-model="addAPIDlgVis" title="添加一个API">
      <Row>
        <Col span="8">
          <span>{{label}}</span>
        </Col>
        <Col span="16">
          <Input v-model="addApiData.name" placeholder></Input>
        </Col>
      </Row>
    </Modal>
  </div>
</template>

<script>
import InterfaceContainer from "./Interface";
import { AddAPI, AddGroup } from "@/api/api.js";
export default {
  name: "ServiceContainer",
  components: { InterfaceContainer },
  props: ["service"],
  data() {
    return {
      column: [
        { title: "项目", key: "name" },
        { title: "值", key: "value" }
      ],
      addApiData: {
        group: "",
        name: "",
        groupMode: false
      },
      addAPIDlgVis: false,
      label: ""
    };
  },
  methods: {
    getTable: function(method) {
      let res = [];
      for (let key in method) {
        res.push({ name: key, value: method[key] });
      }
      return res;
    },
    showAddApiDlg: function(group) {
      this.addApiData.group = group;
      this.addApiData.groupMode = false;

      this.label = "API名称";
      this.addAPIDlgVis = true;
    },
    addApi: function() {
      if (this.addApiData.groupMode) {
        AddGroup({
          service: this.service.name,
          group: this.addApiData.name
        }).then(res => {
          if (res.code == 0) {
            //
          }
        });
      } else {
        AddAPI({
          service: this.service.name,
          group: this.addApiData.group,
          api: this.addApiData.name
        }).then(res => {
          if (res.code == 0) {
            //
          }
        });
      }
    },
    showAddGoupDlg: function() {
      this.label = "组名称";
      this.addApiData.groupMode = true;
      this.addAPIDlgVis = true;
    }
  },
  mounted() {
    console.info("comservice",this.service)
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
