<script setup>
import { ref } from 'vue';
import users from 'api/users';
import Card from 'components/cards/Card.vue';
import {
  Squares2X2Icon,
  ListBulletIcon,
  UserIcon,
  ArrowsPointingOutIcon,
  ArrowsPointingInIcon,
} from '@heroicons/vue/24/outline';

const isGrid = ref(true);
const isExpandingAll = ref(false);
const items = ref([
  {
    id: '8a8f4d7c-3c7a-4d3d-9d4b-2efab0a0d0c7',
    text: 'Lorem ipsum dolor sit amet',
    subText:
      'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Nulla facilisi. In sed tortor eget lorem posuere lacinia. Quisque fermentum odio eget nunc aliquet, eget placerat elit fringilla.',
    image:
      'https://fastly.picsum.photos/id/27/3264/1836.jpg?hmac=p3BVIgKKQpHhfGRRCbsi2MCAzw8mWBCayBsKxxtWO8g',
    showMore: false,
  },
  {
    id: '1f3b1d74-0d39-4fae-9d1e-3e8d96a7b8d0',
    text: 'Pellentesque habitant',
    subText:
      'Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Fusce vitae semper nunc. Aenean vel dui sit amet nunc commodo ultricies. Morbi eget massa eget risus bibendum sodales. Integer tincidunt tellus vel justo vestibulum, et tincidunt velit aliquet.',
    image:
      'https://fastly.picsum.photos/id/26/4209/2769.jpg?hmac=vcInmowFvPCyKGtV7Vfh7zWcA_Z0kStrPDW3ppP0iGI',
    showMore: false,
  },
  {
    id: '2e3f09b6-6d7d-42c4-8c0c-2d0c7a3f7e5f',
    text: 'Sed egestas',
    subText:
      'Sed egestas, nulla vitae rhoncus feugiat, sapien justo fermentum arcu, at lacinia quam libero eu justo. Nam sed pellentesque justo. Phasellus at erat ac justo consequat bibendum. Vestibulum consectetur augue sit amet felis posuere, a bibendum nibh vulputate. Nunc convallis, leo id vehicula scelerisque, sem velit consectetur justo, non aliquet sapien purus et eros.',
    image:
      'https://fastly.picsum.photos/id/30/1280/901.jpg?hmac=A_hpFyEavMBB7Dsmmp53kPXKmatwM05MUDatlWSgATE',
    showMore: false,
  },
  {
    id: 'b26f843e-0a5e-4a45-8af1-3e3f6cd2f2c8',
    text: 'Vestibulum ante ipsum primis in faucibus',
    subText:
      'Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Donec euismod, leo in bibendum interdum, sem massa venenatis libero, id cursus felis libero at metus. Vivamus felis risus, vehicula vel pretium eget, imperdiet sed metus. Fusce fermentum mi et suscipit rhoncus.',
    image:
      'https://fastly.picsum.photos/id/34/3872/2592.jpg?hmac=4o5QGDd7eVRX8_ISsc5ZzGrHsFYDoanmcsz7kyu8A9A',
    showMore: false,
  },
  {
    id: '9c1c0a88-5e6f-495e-b0a2-9f1d2e6c9c2e',
    text: 'Maecenas',
    subText:
      'Maecenas tincidunt magna sit amet ex sagittis, nec vulputate elit maximus. Nullam auctor, justo at malesuada tristique, mi massa tincidunt sapien, id vulputate sem ligula sit amet ligula. Quisque non condimentum ligula. Aliquam ut metus a lacus dignissim fringilla. Morbi vel sem eget eros hendrerit malesuada.',
    image: '',
    showMore: false,
  },
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

    <div :class="`grid grid-${isGrid ? 'rows' : 'cols'}-3 gap-2`">
      <div class="" v-for="item in items" :key="item">
        <component :is="Card" class="p-1" :image="item.image">
          <template v-slot:card-title> {{ item.id }} </template>
          <template v-slot:card-text>{{ item.text }}</template>
          <template v-slot:card-sub-text>
            <span class="text-ellipsis overflow-hidden">
              {{ ellipse(item) }}
              <span
                class="text-blue-500 hover:text-blue-900 cursor-pointer"
                @click="item.showMore = !item.showMore"
              >
                show {{ item.showMore ? 'less' : 'more' }}...
              </span>
            </span>
          </template>
        </component>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
