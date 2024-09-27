<template>
  <v-app class="app-background">
    <v-main>
      <PackManager />
    </v-main>
  </v-app>
</template>

<script>
import PackManager from './components/PackManager.vue'

export default {
  name: 'App',

  components: {
    PackManager,
  },
}
</script>

<style scoped>
.app-background {
  background-color: #000;
  color: #ffcc00; /* Optional: Set a default text color */
  min-height: 100vh; /* Ensure it covers the full viewport height */
}
</style>