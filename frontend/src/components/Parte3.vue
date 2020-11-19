<template>
  <b-container fluid id="content">
    <b-form-input
      name="id"
      v-model="id"
      placeholder="Ingresar ID Comprador"
    ></b-form-input>

    <b-button id="enviar" variant="primary" v-on:click="onSubmit()"
      >Enviar</b-button
    >

    <h6>Historial Compras</h6>
    <b-table sticky-header="300px" :items="compras" head-variant="light">
    </b-table>

    <h6>Compradores igual IP</h6>
    <b-table sticky-header="300px" :items="ips" head-variant="light"> </b-table>
  </b-container>
</template>
<script>
import axios from "axios";
export default {
  name: "Parte3",
  data: function () {
    return {
      id: "",
      compras: [],
      ips: [],
    };
  },
  methods: {
    onSubmit: function () {
      console.log(this.id);
      axios
        .post("http://localhost:3000/buyer", { id: this.id })
        .then((response) => {
          this.compras= response.data.shops;
          this.ips = response.data.equalip;
        })
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