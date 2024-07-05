<script setup>
import {
  ChevronDoubleLeftIcon,
  ChevronDoubleRightIcon,
  ChevronLeftIcon,
  ChevronRightIcon,
} from '@heroicons/vue/24/outline';

const props = defineProps({
  total: { type: Number, default: 0 },
});

const emit = defineEmits({
  'on-change-page': (payload) => {
    return true;
  },
  'on-change-size': (payload) => {
    return true;
  },
});
</script>

<template>
  <div v-if="total > 0" class="md:flex flex-row gap-2">
    <div class="py-1 basis-3/4">
      <label for="" class="">Show</label>
      <select
        name=""
        id=""
        class="border rounded mx-2 px-2"
        @change="emit('on-change-size', $event.target.value)"
      >
        <option value="0">All</option>
        <option value="5">5</option>
        <option value="10">10</option>
        <option value="25">25</option>
        <option value="100">100</option>
      </select>
    </div>

    <button
      class="hover:bg-blue-50 bg-white px-4 py-1"
      @click="emit('on-change-page', 0.1)"
    >
      <ChevronDoubleLeftIcon class="size-4 text-blue-500" />
    </button>
    <button
      class="hover:bg-blue-50 bg-white px-4 py-1"
      @click="emit('on-change-page', 0.2)"
    >
      <ChevronLeftIcon class="size-4 text-blue-500" />
    </button>

    <button
      v-for="i in total"
      :key="i"
      class="hover:bg-blue-50 bg-white px-4 py-1"
      @click="emit('on-change-page', i)"
    >
      {{ i }}
    </button>

    <button
      class="hover:bg-blue-50 bg-white px-4 py-1"
      @click="emit('on-change-page', 0.3)"
    >
      <ChevronRightIcon class="size-4 text-blue-500" />
    </button>
    <button
      class="hover:bg-blue-50 bg-white px-4 py-1"
      @click="emit('on-change-page', 0.4)"
    >
      <ChevronDoubleRightIcon class="size-4 text-blue-500" />
    </button>
  </div>
</template>

<style scoped></style>
