<script setup>
import Card from 'components/cards/Card.vue';
async function getUser() {
  const { VITE_API_ENDPOINT: api } = process.env;
  const { fetch } = window;

  const response = await fetch(`${api}/users`, { method: 'GET' });
  console.log(response);
}
</script>

<template>
  <div>
    <button @click="getUser()">Submit</button>
    <h1>Dashboards</h1>
    <component :is="Card" foo="hello"></component>
  </div>
</template>

<style scoped></style>
