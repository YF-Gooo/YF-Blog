<template>
    <div style = "width:80%;height:100%">
      <el-radio-group v-model="isCollapse" style="margin-bottom: 20px;">
      <el-radio-button :label="false">展开</el-radio-button>
      <el-radio-button :label="true">收起</el-radio-button>
      </el-radio-group>
      <el-menu default-active="1-4-1" class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose" :collapse="isCollapse">
      <el-submenu index="1">
          <template slot="title">
          <i class="el-icon-menu"></i>
          <span slot="title">管理中心</span>
          </template>
          <el-menu-item-group>
          <span slot="title">文章管理</span>
            <el-menu-item index="1-1">选项1</el-menu-item>
            <el-menu-item index="1-2">选项2</el-menu-item>
            </el-menu-item-group>
          <el-menu-item-group>
          <span slot="title">视频管理</span>
            <el-menu-item index="1-1">选项1</el-menu-item>
            <el-menu-item index="1-2">选项2</el-menu-item>
            </el-menu-item-group>
          <el-menu-item-group>
          <span slot="title">用户管理</span>
            <el-menu-item index="1-1">选项1</el-menu-item>
            <el-menu-item index="1-2">选项2</el-menu-item>
            </el-menu-item-group>
          </el-submenu>
      <el-menu-item index="2">
          <i class="el-icon-setting"></i>
          <span slot="title">设置中心</span>
      </el-menu-item>
      </el-menu>
    </div>
</template>

<script>
import store from '@/store'
export default {
  data() {
    return {
      text:"",
      isCollapse: true
    };
  },
  store,
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath);
    },
    handlesearch(){
    let _this = this;
    _this.$store.commit('changeAMSearchKW', _this.text)
    _this.text=""
    },
  }
}
</script>

<style>

  .el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 200px;
    min-height: 400px;
  }
</style>