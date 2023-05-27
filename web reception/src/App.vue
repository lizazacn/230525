<template>
  <el-config-provider>
    <router-view />
  </el-config-provider>
</template>

<script setup lang="ts">
import { ElConfigProvider } from 'element-plus';
import { Socket } from 'socket.io-client';
import { reactive, ref, inject } from 'vue'

import { provide } from 'vue'

const socket = inject('socket') as Socket;

socket.on("connect", (...e: any) => {
  console.log("链接成功", e);

  provide('id', e[0] || '0')
});


</script>
<style></style>
