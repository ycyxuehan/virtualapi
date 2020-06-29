import axios from 'axios';

let prefix="/admin";

export const GetServices=()=>{return axios.get(`${prefix}/services`).then(res=>res.data);};
export const GetService=(name)=>{return axios.get(`${prefix}/service/${name}`).then(res=>res.data);}
export const AddService=()=>{return axios.post(`${prefix}/service`).then(res=>res.data);}
export const AddGroup=(params)=>{return axios.post(`${prefix}/group`, {params:params}).then(res=>res.data);}
export const AddAPI=(params)=>{return axios.post(`${prefix}/api`, {params:params}).then(res=>res.data);}
