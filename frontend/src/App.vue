<template>
  <div>
    <b-navbar type="light" variant="light" id="app">
      <b-navbar-nav>
        <b-button-group>
          <b-button to='/parte1'>Parte 1</b-button>
          <b-button to='/parte2'>Parte 2</b-button>
          <b-button to='/parte3'>Parte 3</b-button>
        </b-button-group>
      </b-navbar-nav>
    </b-navbar>
    <b-container fluid>
      <router-view/>
    </b-container>
  </div>
</template>

<script>
export default {
  name: "App",
};
</script>
