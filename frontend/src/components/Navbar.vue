
<template>
    <el-menu
    class="el-menu-demo"
    mode="horizontal"
    background-color="#545c64"
    text-color="#fff"
    active-text-color="#ffd04b"
    :default-active="$route.path"
    router>
    <el-row style="height:60px;">
    <el-col :span="2"><el-menu-item index="/" >首页</el-menu-item></el-col>
    <el-col :span="2"><el-menu-item index="/me">关于我</el-menu-item></el-col>
    <el-col :span="2"><el-menu-item index="/signin">登录</el-menu-item></el-col>
    <el-col :span="16" class="main-title">Welcome to YF-Blog~~</el-col>
    </el-row>
    </el-menu>
</template>
<script>
  export default {
    data() {
      return {
        activeIndex: '1',
        activeIndex2: '1',
        squareUrl: "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png",
        user:"USER"
      };
    },
    methods: {
    }
  }
</script>
<style>
  el-menu {
    margin: 0%;
    padding: 0%;
    height:60px;
  }
  el-row{
    height:60px;
    margin:0%;
    padding: 0%;
  }
  el-col{
    height:60px;
    margin:0%;
    padding: 0%;
  }
  el-avatar{
    height:60px;
    margin:0%;
    padding: 0%;
  }
  .main-title{
    height:60px;
    margin:0%;
    padding-left: 18%;
    text-align: left;
    line-height: 60px;
    font-size: 30px;
    color:aliceblue
  }
  .main-title:hover{
    height:60px;
    margin:0%;
    padding-left: 18%;
    text-align: left;
    line-height: 60px;
    font-size: 29px;
    color:white;
    opacity: 10%;
  }
  .username{
    text-align: left;
    line-height: 60px;
    color:wheat;
    opacity: 10%;
  }
  .username:hover{
    text-align: left;
    line-height: 60px;
    color:wheat;
  }
  .avatar{
    height:60px;
  }
  .el-dropdown-link {
    cursor: pointer;
    color: #409EFF;
    text-align: left;
    line-height: 60px;
    color: white;
  }
  .el-icon-arrow-down {
    font-size: 12px;
  }
</style>
