
<template>
  <div :key="buyer.id">
    <p>{{ buyer.name }}</p>
    <p>{{ buyer.age }}</p>
  </div>
</template>

<script>
export default {
  name: "Buyer",
  props: ["buyer"]
};
</script>