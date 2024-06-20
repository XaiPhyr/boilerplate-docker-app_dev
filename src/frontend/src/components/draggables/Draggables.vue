<script setup>
import { ref } from 'vue';
const items = ref([
  'admin',
  'moderator',
  'member',
  'guest',
  'developer',
  'designer',
  'analyst',
  'manager',
  'support',
  'contributor',
]);

const selectedRoles = ref([]);

function onDrag(event) {
  const { dataTransfer, target } = event;
  const { dataset } = document.querySelector(`#${target.id}`);

  dataTransfer.setData('text/plain', dataset.value);
  dataTransfer.setData('source', target.id);
}

function onDragOver(event) {
  event.preventDefault();
}

function onDropInside(event) {
  event.preventDefault();

  const { dataTransfer, target } = event;
  if (target.id === 'target') {
    const source = dataTransfer.getData('source');
    if (source) {
      const value = dataTransfer.getData('text/plain');
      const data = document.querySelector(`#${source}`);
      data.classList.add('text-red-500');

      selectedRoles.value.push(value);

      data.addEventListener('dragstart', (event) => {
        const { dataTransfer } = event;
        dataTransfer.setData('source-from-target', source);
      });

      target.appendChild(data);
    }
  }
}

function onDragOverOutside(event) {
  event.preventDefault();
}

function onDropOutside(event) {
  event.preventDefault();

  const { dataTransfer, target } = event;
  if (target.id === 'outside') {
    const source = dataTransfer.getData('source-from-target');
    if (source) {
      const t = document.querySelector('#target');
      const s = t.querySelector(`#${source}`);
      s.classList.remove('text-red-500');

      const sources = document.querySelector('#sources');
      console.log(sources);
      sources.appendChild(s);
    }
  }
}
</script>

<template>
  <div class="">
    <div
      id="outside"
      class=""
      @drop="onDropOutside($event)"
      @dragover="onDragOverOutside($event)"
    >
      <div
        id="target"
        class="border flex flex-row flex-wrap gap-4 p-1 w-1/4 h-auto text-center mb-5"
        @drop="onDropInside($event)"
        @dragover="onDragOver($event)"
      ></div>

      <div class="flex flex-row gap-2" id="sources">
        <div
          :id="`source-${item}`"
          v-for="(item, index) in items"
          :key="index"
          class="border p-1 w-1/4 text-center mb-5"
          draggable="true"
          @dragstart="onDrag($event)"
          :data-value="item"
        >
          {{ item }}
        </div>
      </div>
    </div>

    <input
      id="input_text"
      type="text"
      class="border p-1 w-1/4 text-center mb-5"
    />
  </div>
</template>

<style scoped></style>
