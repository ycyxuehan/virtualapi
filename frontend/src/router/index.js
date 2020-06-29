import Vue from "vue";
import VueRouter from "vue-router";
import Service from "../views/Service.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/service",
    name: "Service",
    component: Service
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
