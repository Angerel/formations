<template>
    <li>{{item.label}}</li>
</template>

<script setup lang="ts">
    import {ListItem} from "@/models/ListItem";

    defineProps<{
        item:ListItem
    }>();
</script>

<script lang="ts">
export default {
    name: "ListItemView"
}
</script>

<style scoped>

</style>