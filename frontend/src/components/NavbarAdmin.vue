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
    <el-col :span="2"><el-menu-item index="/me">文章列表</el-menu-item></el-col>
    <el-col :span="2"><el-menu-item index="/editor">随笔</el-menu-item></el-col>
    <el-col :span="2"><el-menu-item index="/ml">机器学习</el-menu-item></el-col>
    <el-col :span="2"><el-menu-item index="/cs">计算机科学</el-menu-item></el-col>
    <el-col :span="2"><el-menu-item index="/signup">注册新用户</el-menu-item></el-col>
    <el-col :span="8" class="main-title"></el-col>
    <el-col :span="1" class="avatar"><el-avatar shape="square" :size="60" :src="squareUrl"></el-avatar></el-col>
    <el-col :span="1">
    <el-dropdown>
      <span class="el-dropdown-link">
        {{user}}<i class="el-icon-arrow-down el-icon--right"></i>
      </span>
      <el-dropdown-menu slot="dropdown">
        <el-dropdown-item>首页</el-dropdown-item>
        <el-dropdown-item>个人信息修改</el-dropdown-item>
        <el-dropdown-item divided>注销</el-dropdown-item>
      </el-dropdown-menu>
    </el-dropdown>
    </el-col>
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
        user:"YIFAN"
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
</style>