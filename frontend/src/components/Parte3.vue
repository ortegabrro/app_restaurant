<template>
  <b-container fluid id="content">
    <b-form-input
      v-model="text"
      placeholder="Ingresar ID Comprador"
    ></b-form-input>
    <div class="mt-2">Value: {{ text }}</div>
    <b-button id="enviar" variant="primary" v-on:click="load()"
      >Enviar</b-button
    >

    <h6>Historial Compras</h6>
    <b-table
      sticky-header="300px"
      :items="compras"
      caption-top
      head-variant="light"
    >
    </b-table>

    <h6>Compradores igual IP</h6>
    <b-table
      sticky-header="300px"
      :items="ips"
      caption-top
      head-variant="light"
    >
    </b-table>
  </b-container>
</template>
<script>
import axios from "axios";
export default {
  name: "Parte3",
  data() {
    return {
      text: "",
      id: "",
      compras: [],
      ips: [],
    };
  },
  methods: {
    onSubmit: function () {
      axios
        .post("http://localhost:3000/buyer", { id: this.id })
        .then((response) => response)
        .catch((error) => {
          console.error(error);
        });
    },
  },
};
</script>
<style>
#content {
  margin-top: 20px;
}
#enviar {
  margin-top: 10px;
  margin-block-end: 30px;
}
</style>