<script setup>
const props = defineProps(['foo']);
</script>

<template>
  <div>
    this card is a component <span class="text-red"> {{ props.foo }}</span>
  </div>
</template>

<style scoped>
.text-red {
  color: red;
}
</style>
