<script setup lang="ts">
  import {ListItem} from "@/models/ListItem";
  import ListItemView from "@/components/ListItemView.vue";

  const list:ListItem[] = [
      new ListItem("Test 1"),
      new ListItem("2"),
      new ListItem("Test 3")
  ];
</script>

<template>
    <ul>
        <ListItemView v-for="(item, index) in list" :item="item" :key="index" />
    </ul>
</template>
