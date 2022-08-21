<template>
  <body>
    <input
      class="form-control bg-dark text-light"
      type="text"
      placeholder="Search"
      v-model="filter"
      ref="search"
    />
    <table class="table table-dark table-hover">
      <thead class="thead thead-light text-muted">
        <tr>
          <th scope="col">Deal</th>
          <th scope="col">Views</th>
          <!-- TODO:score  -->
          <!-- <th scope="col">Score</th> -->
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
        <tr
          scope="row"
          v-for="(topic, index) in filteredTopics"
          :key="`topic.topic_id-${index}`"
        >
          <td scope="col">
            <a
              :href="`https://forums.redflagdeals.com${topic.web_path}`"
              target="_blank"
              v-html="highlightMatches(topic.title)"
            ></a>
          </td>
          <td scope="col">{{ topic.total_views }}</td>
          <td scope="col">{{ topic.score }}</td>
          <!-- <td scope="col">{{ topic.votes.up - topic.votes.down }}</td> -->
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
import Mousetrap from "mousetrap";

import "vue-loading-overlay/dist/vue-loading.css";

export default {
  data() {
    return {
      topics: [],
      isLoading: false,
      filter: "",
    };
  },
  mounted() {
    this.isLoading = true;
    this.fetchDeals();
    Mousetrap.bind("/", this.focusSearch, "keyup");
    Mousetrap.bind("r", this.loadDeals);
    this.interval = setInterval(() => this.fetchDeals(), 10000);
  },
  methods: {
    focusSearch() {
      this.$refs.search.focus();
    },
    fetchDeals() {
      axios
        .get("api/v1/topics")
        .then((response) => {
          this.topics = response.data;
          this.isLoading = false;
        })
        .catch((err) => {
          console.log(err.response);
        });
    },
  },
  props: ["date"],
  computed: {
    formatDate() {
      return (v) => {
        return moment(String(v)).format("hh:mm A z (MM/DD)");
      };
    },
    filteredTopics() {
      return this.topics.filter((row) => {
        const titles = row.title.toString().toLowerCase();
        const searchTerm = this.filter.toLowerCase();

        return titles.includes(searchTerm);
      });
    },
    highlightMatches() {
      return (v) => {
        if (this.filter == "") return v;
        const matchExists = v.toLowerCase().includes(this.filter.toLowerCase());
        if (!matchExists) return v;

        const re = new RegExp(this.filter, "ig");
        return v.replace(re, (matchedText) => `<mark>${matchedText}</mark>`);
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
