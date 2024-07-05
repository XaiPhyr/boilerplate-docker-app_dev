<script setup>
import { ref } from 'vue';
import users from 'api/users';
import Card from 'components/cards/Card.vue';
import Table from 'components/tables/Table.vue';
import Draggables from 'components/draggables/Draggables.vue';
import {
  Squares2X2Icon,
  ListBulletIcon,
  UserIcon,
  ArrowsPointingOutIcon,
  ArrowsPointingInIcon,
} from '@heroicons/vue/24/outline';

const isGrid = ref(true);
const isExpandingAll = ref(false);
const products = ref([
  {
    name: 'Product 1',
    description:
      'Lorem ipsum dolor sit amet consectetur, adipisicing elit. Deserunt vel\
    tempore tempora odio modi possimus quia, excepturi nesciunt provident\
    repudiandae architecto necessitatibus earum quis nam perspiciatis eaque\
    iusto ad mollitia.',
    price: 10.99,
  },
  {
    name: 'Product 2',
    description:
      'Lorem ipsum dolor sit amet consectetur adipisicing elit. Ipsa inventore ut\
    sit dignissimos corporis repellat veritatis autem, adipisci aut quia\
    deleniti obcaecati ullam, id quasi voluptas provident earum et laudantium?\
    Consequuntur praesentium, nobis dolore distinctio quidem eveniet ipsam fugit\
    esse nam consectetur dolorum temporibus aliquam nostrum libero cupiditate\
    tenetur, ab nemo voluptas aspernatur magnam facilis velit rerum commodi\
    nulla. Neque!',
    price: 19.99,
  },
  {
    name: 'Product 3',
    description: 'This is the description for Product 3.',
    price: 25.5,
  },
  {
    name: 'Product 4',
    description: 'This is the description for Product 4.',
    price: 15.75,
  },
  {
    name: 'Product 5',
    description: 'This is the description for Product 5.',
    price: 30.25,
  },
  {
    name: 'Product 6',
    description: 'This is the description for Product 6.',
    price: 12.99,
  },
  {
    name: 'Product 7',
    description: 'This is the description for Product 7.',
    price: 22.5,
  },
  {
    name: 'Product 8',
    description: 'This is the description for Product 8.',
    price: 17.49,
  },
  {
    name: 'Product 9',
    description: 'This is the description for Product 9.',
    price: 28.99,
  },
  {
    name: 'Product 10',
    description: 'This is the description for Product 10.',
    price: 35.0,
  },
]);

const productColumns = ref([
  { caption: 'Name', field: 'name' },
  { caption: 'Description', field: 'description' },
  { caption: 'Price', field: 'price' },
]);

async function getUser() {
  users
    .get()
    .then((data) => {
      console.log(data);
    })
    .catch(() => {
      //
    })
    .finally(() => {
      //
    });
}

function viewAsGrid() {
  isGrid.value = !isGrid.value;
}

function ellipse(item, index) {
  if (item.showMore) {
    return item.subText;
  }

  return `${item.subText.substring(0, 5)}...`;
}

function expandAll() {
  isExpandingAll.value = !isExpandingAll.value;

  for (const item of items.value) {
    item.showMore = isExpandingAll.value;
  }
}
</script>

<template>
  <div>
    <button
      class="border bg-white rounded px-4 py-1"
      @click="viewAsGrid()"
      :title="`${isGrid ? 'Grid' : 'List'}`"
    >
      <ListBulletIcon v-if="isGrid" class="size-6 text-blue-500" />
      <Squares2X2Icon v-else class="size-6 text-blue-500" />
    </button>

    <!-- <button class="border bg-white rounded p-1" @click="getUser()">
      <UserIcon class="size-6 text-blue-500" />
    </button> -->

    <button class="border bg-white rounded px-4 py-1" @click="expandAll()">
      <ArrowsPointingInIcon
        v-if="isExpandingAll"
        class="size-6 text-blue-500"
      />
      <ArrowsPointingOutIcon v-else class="size-6 text-blue-500" />
    </button>

    <component :is="Draggables" />
  </div>
</template>

<style scoped></style>
