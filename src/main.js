import { createApp } from "vue";
import App from "./App.vue";

import VueGtag from "vue-gtag";

import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.min.js";
import "mousetrap/mousetrap.min.js";

import "./style.css";
import "./xess.css";

createApp(App)
  .use(VueGtag, {
    config: { id: "G-YF11ZH9SYD" },
  })
  .mount("#app");
