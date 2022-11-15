<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useToast, POSITION } from 'vue-toastification'

import { ping } from './api/ping'

const toast = useToast()

onMounted(() => {
  ping()
    .then((res) => {
      console.log(`${res.data},后端服务状态正常。`)
    })
    .catch((err) => {
      toast.error(`后端服务状态异常,${err.message}`, {
        position: POSITION.TOP_RIGHT
      })
    })
})
</script>

<template>
  <RouterView></RouterView>
</template>

<style lang="scss">
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

:root {
  --core-bg: #222429;
  --dark-bg: #0f0f13;
  --primary: #1b1b1f;
  --secondary: #17171b;
  --white: #ffffff;
  --settings-icon-color: #c1c7cd;
  --settings-icon-hover: #9fa7ac;
  --body-color: #273346;
  --input-bg: #f8f8fa;
  --input-chat-color: #a2a2a2;
}

body {
  display: flex;
  min-height: 100vh;
  justify-content: center;
  align-items: center;
  background: linear-gradient(200deg, #e3c5eb, #a9c1ed);
  overflow: hidden;
  font-family: 'lucida grande', 'lucida sans unicode', lucida, helvetica, 'Hiragino Sans GB',
    'Microsoft YaHei', 'WenQuanYi Micro Hei', sans-serif;
}
</style>
