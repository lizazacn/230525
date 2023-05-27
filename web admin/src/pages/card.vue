<template>
  <el-row>
    <el-col :span="8">  <el-card class="box-card">
      <!--    <div v-for="o in 4" :key="o" class="text item">{{ 'List item ' + o }}</div>-->
      <div class="title">
        <b>ip: 127.0.0.1&nbsp; &nbsp;</b>
        <b style="color: green"> 13秒前在线</b>
      </div>
      <p>
        {
        name:张三,
        classNumber:三年级二班,
        }
      </p>
      <el-button>进入年龄页面</el-button>
      <el-button>进入性别页面</el-button>
    </el-card></el-col>
<!--    <el-col :span="8">  <el-card class="box-card">-->
<!--      &lt;!&ndash;    <div v-for="o in 4" :key="o" class="text item">{{ 'List item ' + o }}</div>&ndash;&gt;-->
<!--      <div class="title">-->
<!--        <b>ip: 127.0.0.1 &nbsp; &nbsp;</b>-->
<!--        <b style="color: green">9秒前在线</b>-->
<!--      </div>-->
<!--      <p>-->
<!--        {-->
<!--        name:张三,-->
<!--        classNumber:三年级二班,-->
<!--        }-->
<!--      </p>-->
<!--      <el-button>进入年龄页面</el-button>-->
<!--      <el-button>进入性别页面</el-button>-->
<!--    </el-card></el-col>-->
<!--    <el-col :span="8">  <el-card class="box-card">-->
<!--      &lt;!&ndash;    <div v-for="o in 4" :key="o" class="text item">{{ 'List item ' + o }}</div>&ndash;&gt;-->
<!--      <div class="title">-->
<!--        <b>ip: 127.0.0.1&nbsp; &nbsp;</b>-->
<!--        <b style="color: red">离线</b>-->
<!--      </div>-->
<!--      <p>-->
<!--        {-->
<!--        name:张三,-->
<!--        classNumber:三年级二班,-->
<!--        }-->
<!--      </p>-->
<!--      <el-button>进入年龄页面</el-button>-->
<!--      <el-button>进入性别页面</el-button>-->
<!--    </el-card></el-col>-->
<!--    <el-col :span="8">  <el-card class="box-card">-->
<!--      &lt;!&ndash;    <div v-for="o in 4" :key="o" class="text item">{{ 'List item ' + o }}</div>&ndash;&gt;-->
<!--      <div class="title">-->
<!--        <b>ip: 127.0.0.1&nbsp; &nbsp;</b>-->
<!--        <b style="color: green">1秒前在线</b>-->
<!--      </div>-->
<!--      <p>-->
<!--        {-->
<!--        name:张三,-->
<!--        classNumber:等待获取,-->
<!--        }-->
<!--      </p>-->
<!--      <el-button>进入年龄页面</el-button>-->
<!--      <el-button>进入性别页面</el-button>-->
<!--    </el-card></el-col>-->

  </el-row>
  <p>
    以上P标签中应该是json格式，是可以收缩的，我不会修改，请帮我修改。
    <br>
    卡片是有新的客户进来后再添加新的卡片出现

  </p>
</template>
<style scoped>
.text {
  font-size: 14px;
}

.item {
  padding: 18px 0;
}

.box-card {
  width: 480px;
}
</style>
<script setup lang="ts">
</script>