import { createApp } from "vue";
import App from "./App.vue";
import VueGtag from "vue-gtag";
import { createRouter, createWebHashHistory } from "vue-router";

import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.min.js";

import "./theme.css";

const routes = [];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

const app = createApp(App).use(VueGtag, {
  config: { id: "G-YF11ZH9SYD" },
});

app.use(router);
console.log(router);
app.mount("#app");
