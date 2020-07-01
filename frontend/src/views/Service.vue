<template>
  <div class="home">
    <Header class="group-title"><Row>
      <Col span="22">
        <h3>Service {{service.name}} <Tag>端口：{{service.port}}</Tag></h3>
      </Col>
      <Col span="2">
        <Button type="primary" @click="showAddServiceDlg">添加服务</Button>
      </Col>
    </Row></Header>
    <service-container :service="service"/>
    <Modal
      v-model="addServiceDlgVis"
      title="添加一个服务"
      @on-ok="addService"
      >
      <Row>
        <Col span="8">
          <span>服务名称</span>
        </Col>
        <Col span="16">
          <Input v-model="addServiceData.name" placeholder=""></Input>
        </Col>
      </Row>
      <Col span="8">
          <span>服务端口</span>
        </Col>
        <Col span="16">
          <Input v-model="addServiceData.port" placeholder=""></Input>
        </Col>
      </Row>
             <Col span="8">
          <span>API前缀</span>
        </Col>
        <Col span="16">
          <Input v-model="addServiceData.prefix" placeholder=""></Input>
        </Col>
      </Row>
             <Col span="8">
          <span>说明</span>
        </Col>
        <Col span="16">
          <Input v-model="addServiceData.description" placeholder=""></Input>
        </Col>
      </Row>
    </Modal>
  </div>
</template>

<script>
// @ is an alias to /src
import ServiceContainer from "@/components/Service.vue";
import {GetService, AddService} from "@/api/api.js"
export default {
  name: "Service",
  components: {
    ServiceContainer
  },
  data(){
    return {
      service:{},
      addServiceData:{
        name: "",
      },
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
    },
    addService: function(){
      AddService(this.addServiceData).then(res=>{
        console.info("add service:", res)
        if(res.code == 0){

        }
      })
    },
    showAddServiceDlg: function(){
      this.addServiceDlgVis = true
    }
  },
  mounted(){
    this.getService();
  }
};
</script>
