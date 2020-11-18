<template>
  <b-container id="content" fluid>
    <b-table sticky-header="500px" striped hover :items="items"></b-table>
  </b-container>
</template>

<script>
import axios from "axios";
export default {
  name: "Parte2",
  data() {
    return {
      items: [],
    };
  },
  methods: {
    load: function () {
      axios
        .get("http://localhost:3000/buyers")
        .then((response) => {
          this.items = response.data.items;
        })
        .catch((error) => {
          console.error(error);
        });
    },
  },
  created() {
    this.load();
  },
};
</script>
<style>
#content {
  margin-top: 40px;
}
</style>
