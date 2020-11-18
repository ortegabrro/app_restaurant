<template>
  <b-container fluid>
    <b-table striped hover :items="items"></b-table>
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
          this.items = response.data;
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
