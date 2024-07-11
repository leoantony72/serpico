<template>
  <div class="app">
    <Header />
    <Sidebar />
    <router-view></router-view>
  </div>
</template>

<script setup>
import Header from "./components/Navigation.vue";
import Sidebar from "./components/Sidebar.vue";
// import Upload from "./components/Organize.vue";
import { RouterView } from 'vue-router';


</script>

<style>
</style>
