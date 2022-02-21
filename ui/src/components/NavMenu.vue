<template>
  <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
    <el-menu-item index="1"><router-link to="/">Home</router-link></el-menu-item>
    <el-menu-item index="2"><router-link to="/employees">员工</router-link></el-menu-item>
  </el-menu>
</template>

<script>
  export default {
    data(){
      return {
        activeIndex: '1'
      }
    }
  }

</script>