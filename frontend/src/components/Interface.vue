<template>
  <div class="hello" id="interface-container">
    <Row v-for="(api, index) in apis" :key="index">
      <Row class="api-title">
        <Col span="12"> <h3> API {{api.name}}</h3> </Col>
        <Col span="10">
          <Tag v-for="(method, index) in api.methods" :key="index" color="blue">{{method.method}}</Tag>
        </Col>
        <Col span="2">
          <Button type="primary" @click="showAddMethodDlg(api)">添加方法</Button>
        </Col>
      </Row>
      <Card style v-for="(method, index) in api.methods" :key="index" dis-hover class="api-card">
        <div style="text-align:center; width:50%;">
          <Table :columns="column" :data="getTable(method)" :show-header="false"></Table>
        </div>
      </Card>
    </Row>
    <Modal
      v-model="addMethodDlgVis"
      :title="title"
      @on-ok="addMethod"
      >
      <Row class="edit-row">
        <Col span="6">
          <span>请求方法：</span>
        </Col>
        <Col span="18">
          <Select v-model="addMethodData.method">
            <Option v-for="(method, index) in availableMethods" :value="method" :key="index">{{method}}</Option>
          </Select>
        </Col>
        
      </Row>
      <Row class="edit-row">
        <Col span="6">
          <span>请求头：</span>
        </Col>
        <Col span="18">
          <Input v-model="headers" placeholder=""></Input>
        </Col>
        
      </Row>
      <Row class="edit-row">
        <Col span="6">
          <span>URL参数：</span>
        </Col>
        <Col span="18">
         <Input v-model="queries" placeholder=""></Input>
        </Col>
        
      </Row>
      <Row class="edit-row">
        <Col span="6">
          <span>请求数据：</span>
        </Col>
        <Col span="18">
          <Input v-model="request" placeholder=""></Input>
        </Col>
        
      </Row>
      <Row class="edit-row">
        <Col span="6">
          <span>响应数据：</span>
        </Col>
        <Col span="18">
          <Input v-model="response" placeholder=""></Input>
        </Col>
        
      </Row>

    </Modal>
  </div>
</template>

<script>
import { ChineseColumnName, HTTPMethods, AddMethod } from "@/api/api.js";
export default {
  name: "InterfaceContainer",
  props: ["apis", "group", "service"],
  data() {
    return {
      column: [
        { title: "项目", key: "name" },
        { title: "值", key: "value" }
      ],
      addMethodDlgVis: false,
      title: '添加一个请求方法',
      addMethodData: {
        method:"",
        request:{},
        response:{},
        headers:{},
        queries:{},
      },
              request:"{}",
        response:"{}",
        headers:"{}",
        queries:"{}",
      availableMethods:[],
      api: null,
    };
  },
  methods: {
    getTable: function(method) {
      let res = [];
      for (let key in method) {
        res.push({ name: ChineseColumnName(key), value: method[key] });
      }
      return res;
    },
    chineseColumnName: function(column) {
      return ChineseColumnName(column);
    },
    addMethod : function() {
      console.info("addmethod:", this.service, this.group, this.api.name)
      this.addMethodData.request = JSON.parse(this.request);
      this.addMethodData.response = JSON.parse(this.response);
      this.addMethodData.headers = JSON.parse(this.headers);
      this.addMethodData.queries = JSON.parse(this.queries);
      AddMethod({"api":this.api.name, "group":this.group, "service":this.service}, this.addMethodData).then(res=>{
        console.info("add result:", res)
        if(res.code == 0){
          console.info("added")
          this.api.methods[this.addMethodData.method] = this.addMethodData;
        }
      })
    },
    showAddMethodDlg: function(api){
      let methods = [];
      for(let m in HTTPMethods){
        let exists = false;
        for(let method in api.methods){
          if(HTTPMethods[m] == method){
            exists = true;
            break
          }
        }
        if(!exists){
          methods.push(HTTPMethods[m]);
        }
      }
      this.availableMethods = methods;
      this.addMethodDlgVis = true;
      this.api = api;
    }
  },
  mounted(){
    console.info("interfaces:", this.apis)
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
