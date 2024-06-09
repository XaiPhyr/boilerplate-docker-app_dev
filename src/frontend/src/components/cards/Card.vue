<script setup>
const props = defineProps({
  data: { type: String },
  size: { type: String, default: 'md' },
  image: { type: [String, Object], default: '' },
});
</script>

<template>
  <div class="">
    <div :class="`max-w-${size} bg-white rounded-xl shadow-md overflow-hidden`">
      <div class="md:flex">
        <div v-if="image" class="md:shrink-0">
          <img
            class="h-full w-full object-cover md:h-25 md:w-20"
            :src="image"
            alt="lorem picsum"
          />
        </div>
        <div class="p-8">
          <div class="font-extrabold text-sky-500 pb-1">
            <slot name="card-title" />
          </div>
          <div class="">
            <slot name="card-text" />
          </div>
          <div class="text-xs">
            <slot name="card-sub-text" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
