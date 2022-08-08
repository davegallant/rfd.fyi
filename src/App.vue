<template>
  <body>
    <table class="table table-dark table-hover">
      <thead>
        <tr>
          <th scope="col">Deal</th>
          <th scope="col">Views</th>
          <th scope="col">Last Post</th>
        </tr>
      </thead>
      <tbody>
        <loading
          v-model:active="isLoading"
          color="#ccc"
          opacity="0"
          loader="bars"
          :is-full-page="false"
        >
        </loading>
        <tr scope="row" v-for="topic in topics" :key="topic.topic_id">
          <td scope="col">
            <a
              :href="'https://forums.redflagdeals.com' + topic.web_path"
              target="_blank"
              >{{ topic.title }}</a
            >
          </td>
          <td scope="col">{{ topic.total_views }}</td>
          <td scope="col">{{ formatDate(topic.last_post_time) }}</td>
        </tr>
      </tbody>
    </table>
  </body>
</template>

<script>
import axios from "axios";
import moment from "moment";
import Loading from "vue-loading-overlay";
import "vue-loading-overlay/dist/vue-loading.css";

export default {
  data() {
    return {
      topics: [],
      isLoading: false,
    };
  },
  mounted() {
    this.isLoading = true;
    axios
      .get("http://localhost:8081/api/v1/topics")
      .then((response) => {
        this.topics = response.data;
        this.isLoading = false;
      })
      .catch((err) => {
        console.log(err.response);
      });
  },
  props: ["date"],
  computed: {
    formatDate() {
      return (v) => {
        return moment(String(v)).format("hh:mm A z (MM/DD)");
      };
    },
  },
  components: {
    Loading,
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
