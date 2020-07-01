import axios from "axios";

let prefix = "/admin";

export const GetServices = () => {
  return axios.get(`${prefix}/services`).then((res) => res.data);
};
export const GetService = (name) => {
  return axios.get(`${prefix}/service/${name}`).then((res) => res.data);
};
export const AddService = (data) => {
  return axios.post(`${prefix}/service`, data).then((res) => res.data);
};
export const AddGroup = (params) => {
  return axios
    .post(`${prefix}/group`, null, { params: params })
    .then((res) => res.data);
};
export const AddAPI = (params) => {
  return axios
    .post(`${prefix}/api`, null, { params: params })
    .then((res) => res.data);
};
export const AddMethod = (params, data) => {
    return axios
      .post(`${prefix}/api/method`,  data, { params: params })
      .then((res) => res.data);
  };

export const ChineseColumnName = (column) => {
  switch (column) {
    case "method":
      return "请求方法";
    case "request":
      return "请求数据";
    case "response":
      return "响应数据";
    case "headers":
      return "HTTP请求头";
    case "queries":
      return "URL参数";
  }
  return "";
};

export const HTTPMethods = [
    "GET",
    "POST",
    "PUT",
    "DELETE"
]