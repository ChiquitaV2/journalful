<template>
  <div>
    <h1>Welcome to Journalful!</h1>
    <p>This is your personal research article management system.</p>
    <el-button type="primary" @click="goToArticles">View Articles</el-button>
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const goToArticles = () => {
  router.push('/articles')
}
</script>
