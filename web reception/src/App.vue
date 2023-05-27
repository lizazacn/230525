<template>
  <el-config-provider>
    <router-view />
  </el-config-provider>
</template>

<script setup lang="ts">
import { ElConfigProvider } from 'element-plus';
import { reactive, ref, inject } from 'vue'

const socket = inject('socket') as any;

socket.on("connect", (...e: any) => {
  console.log("链接成功", e);


  (window as any).id = e[0]

});


</script>
<style></style>
