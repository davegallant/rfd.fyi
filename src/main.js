import { createApp } from "vue";
import { createGtm } from "vue-gtm";
import App from "./App.vue";

import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.min.js";
import "mousetrap/mousetrap.min.js";

import "./style.css";
import "./xess.css";

const app = createApp(App).mount("#app");

app.use(
  createGtm({
    id: "G-YF11ZH9SYD",
  })
);
