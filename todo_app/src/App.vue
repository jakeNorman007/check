<script setup>
import axios from "axios";
import { ref, onMounted } from "vue";
import deleteIcon from "./icons/deleteIcon.vue";
import createIcon from "./icons/createIcon.vue";

const todoItems = ref([]);
const body = ref("");

onMounted(async () => {
  const response = await axios
    .get("http://localhost:42069/")
    .then((res) => res.data)
    .catch((err) => console.error(err));

  const todos = response.map((item) => ({ ID: item.ID, Body: item.Body }));
  todoItems.value = todos;
});

const onSubmit = async () => {
  const response = await axios
    .post("http://localhost:42069/", { Body: body.value })
    .then((res) => res.data)
    .catch((err) => console.error(err));

  todoItems.value.push({
    body: response.Body,
  });

  window.location.reload();
};

function onDelete(itemId) {
  axios
    .delete(`http://localhost:42069/${itemId}`)
    .then((res) => res.data)
    .catch((err) => console.error(err));

  window.location.reload();
}
</script>

<template>
  <div class="bg-neutral-800 text-white h-screen">
    <div class="flex items-center rounded-lg gap-2 ml-[10rem] pt-6 mb-2">
      <label for="new-todo-item" class="text-lg">Create a new todo:</label>
      <input
        v-model="body"
        id="new-todo-item"
        type="text"
        class="border border-black text-black pl-2 rounded-md py-1"
      />
      <button
        @click="onSubmit"
        class="flex gap-2 bg-green-500 rounded-md px-4 py-1 border-t border-green-300 hover:bg-green-700"
      >
        <createIcon />Create
      </button>
    </div>
    <div class="border-b border-green-300 h-1 w-5/6 ml-[10rem]"></div>
    <div
      class="flex flex-col gap-3 mt-3 h-[51rem] mx-[25rem] overflow-y-scroll scrollbar-hide border-b border-neutral-300"
    >
      <div v-for="(item, index) in todoItems" :key="index">
        <div
          class="flex justify-between items-center mx-[1rem] mb-1 bg-green-500 rounded-md border-t border-green-300"
        >
          <p @click="toggleEdit" class="text-2xl pl-[4rem]">{{ item.Body }}</p>
          <button
            @click="onDelete(item.ID)"
            class="p-[3rem] bg-green-500 rounded-r-md hover:bg-red-700"
          >
            <deleteIcon />
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
