<template>
  <b-container fluid id="content">
    <b-calendar block name="date" v-model="date" locale="en"></b-calendar>
    <b-button
      id="cargar"
      block
      variant="primary"
      v-on:click="load()"
      >Cargar</b-button
    >
  </b-container>
</template>
<script>
import axios from "axios";
export default {
  name: "Parte1",
  data: function () {
    return {
      date: new Date().toISOString().substr(0, 10),
      res: "",
    };
  },
  methods: {
    load: function () {
      axios
        .post("http://localhost:3000/load", { date: this.date })
        .then((response) =>
          this.$bvModal.msgBoxOk(response.data, {
            title: "Respuesta",
            size: "sm",
            buttonSize: "sm",
            okVariant: "success",
            headerClass: "p-2 border-bottom-0",
            footerClass: "p-2 border-top-0",
            centered: true,
          })
        )
        .catch((error) => {
          console.error(error);
        });
    },
  },
};
</script>
<style>
#content {
  margin-top: 40px;
}
#cargar {
  margin-top: 40px;
}
</style>
