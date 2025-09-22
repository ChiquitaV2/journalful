<script setup lang="ts">

</script>

<template>
  <header>
    <AuthState>
      <template #default="{ loggedIn, clear }">
        <button v-if="loggedIn" @click="clear">Logout</button>
        <NuxtLink v-else to="/api/auth/zitadel">Login</NuxtLink>
      </template>
      <template #placeholder>
        <button disabled>Loading...</button>
      </template>
    </AuthState>
  </header>
</template>
<style scoped>

</style>