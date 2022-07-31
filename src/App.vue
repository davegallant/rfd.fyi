<template>

<body>
<table class="table table-striped table-dark">
  <thead>
    <tr>
      <td scope="col">Deal</td>
      <td scope="col">Views</td>
      <td scope="col">Last Post</td>
    </tr>
  </thead>
  <tbody>
  <tr scope="row" v-for="topic in topics" :key="topic.topic_id">
    <td scope="col">{{ topic.title }}</td>
    <td scope="col">{{ topic.total_views }}</td>
    <td scope="col">{{ topic.last_post_time }}</td>
  </tr>
  </tbody>
</table>

</body>

</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      topics: []
    }
  },
  mounted() {
    axios
      .get('http://localhost:8081/api/v1/topics')
      .then((response) => {
        this.topics = response.data
      }).catch(err => {
        console.log(err.response);
      })
  }
}

</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
