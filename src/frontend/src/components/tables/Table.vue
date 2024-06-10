<script setup>
import {
  PencilSquareIcon,
  ArchiveBoxArrowDownIcon,
} from '@heroicons/vue/24/outline';
import Pagination from 'components/paginations/Pagination.vue';

const props = defineProps({
  'data-column': { type: Array, default: [] },
  'data-source': { type: Array, default: [] },
});

function onChangePage(event) {
  console.log('PAGE CHANGED: ', event);
}

function onChangeSize(event) {
  console.log('SIZE CHANGED: ', event);
}
</script>

<template>
  <div>
    <div class="rounded-xl bg-white p-8 justify-center shadow">
      <table class="w-full table-auto mb-5">
        <caption class="caption-top font-bold text-2xl">
          Products
        </caption>

        <thead class="text-left" v-if="dataColumn.length > 0">
          <th v-for="(column, index) in dataColumn" :key="index">
            {{ column.caption }}
          </th>
          <th class="text-center">Action</th>
        </thead>
        <tbody>
          <tr
            v-for="(data, index) in dataSource"
            :key="index"
            class="border-y-1 border-x-0 border hover:bg-sky-50"
          >
            <td v-for="(column, i) in dataColumn" :key="i" class="py-2 w-px">
              {{ data[column.field] }}
            </td>

            <td class="w-1/12 text-center py-2">
              <button class="px-1">
                <PencilSquareIcon
                  class="size-6 text-blue-500 hover:text-blue-600"
                />
              </button>
              <button class="px-1">
                <ArchiveBoxArrowDownIcon
                  class="size-6 text-red-500 hover:text-red-600"
                />
              </button>
            </td>
          </tr>
        </tbody>
      </table>

      <component
        class="justify-end pt-5"
        :is="Pagination"
        :total="10"
        @on-change-page="(v) => onChangePage(v)"
        @on-change-size="(v) => onChangeSize(v)"
      />
    </div>
  </div>
</template>

<style scoped></style>
