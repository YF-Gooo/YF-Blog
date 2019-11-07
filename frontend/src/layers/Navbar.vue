
<template>
      <div class="navbar">
        <el-row :gutter="10">
          <el-col :span="3">
            <div style="height:50px;">
            <span style="font-size:30px;font-weight:bold;">Y</span>
            <span style="font-size:30px;font-weight:bold;">F</span>
            <span style="color:green;font-size:30px;">B</span>
            <span style="font-size:25px;">l</span>
            <span style="font-size:25px;">o</span>
            <span style="font-size:25px;">g</span>
            </div>
          </el-col>
          <el-col class="nav-menu" :span="1"><router-link to="/">首页</router-link></el-col>
          <el-col class="nav-menu" :span="1"><router-link to="/about">关于我</router-link></el-col>
          <el-col v-if="user!==''" class="nav-menu" :span="1"><router-link to="/editor">写文章</router-link></el-col>
          <el-col :span="6" :offset="4">
            <el-input
              placeholder="请输入内容"
              v-model="text">
              <i slot="prefix" class="el-input__icon el-icon-search"></i>
            </el-input>
          </el-col>
          <div v-if="user!==''">
            <el-col :span="1" :offset="3" style="height:50px;"><el-avatar shape="square" :size="50" :src="avtarUrl"></el-avatar></el-col>
            <el-col :span="0.5" >
              <el-dropdown @command="handleCommand">
                <span class="el-dropdown-link" style="color:grey;font-weight:bolder;font-size:15px">
                  {{user}}<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                <el-dropdown-menu slot="dropdown" trigger="click">
                  <el-dropdown-item command="/">首页</el-dropdown-item>
                  <el-dropdown-item command="/about">关于我</el-dropdown-item>
                  <el-dropdown-item command="/about">管理</el-dropdown-item>
                  <el-dropdown-item disabled>其他</el-dropdown-item>
                  <el-dropdown-item divided ><el-button size="mini" type="info" @click="logout">注销</el-button></el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>
            </el-col>
          </div>
          <div v-else>
            <el-col class="nav-menu" :span="0.5" :offset="3"><router-link to="/signin">登录</router-link></el-col>
            <el-col class="nav-register" :span="0.5" ><router-link to="/signup" style="color:white">免费注册</router-link></el-col>
          </div>
        </el-row>
      </div>
</template>
<script>
import * as API from "@/api/user/";
export default {
  data() {
    return {
      avtarUrl: "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png",
      user:"",
      text:""
    };
  },
  async mounted(){
    const {status,data:{data:{nickname}}} = await this.$axios.get("/api/v1/user/me")
    if (status===200){
      this.user=nickname
    }
  },
  methods:{
    logout: function() {
      let _this = this;
      API.logOut()
        .then(
          response => {
            console.log(response)
            _this.$message({
              message: "登出成功",
              type: "success"
            });
            window.location.href="/"
          },
          response => console.log("登出失败" + response)
        );
      },
    handleCommand(command) {
      let _this = this;
      _this.$router.push(command);
      }
    }
  }
</script>
<style>
  .navbar{
    background-color:#eee;
    padding-left:10%;
    line-height:50px;
    border-bottom: 1px solid #dfe6ec;
    text-align:center;
  }
  .navbar .el-col{
    text-align:center;
  }

  .navbar .nav-register{
    background-color:#009a61;
    border-radius: 4px;
    padding:0 13px;
    margin-left:10px;
  }

  .navbar .nav-menu:hover{
    border-radius: 4px;
    background-color:#dcdfe6;
    opacity:0.5;
  }
  .el-dropdown-link {
    cursor: pointer;
    color: black;
  }
  .el-icon-arrow-down {
    font-size: 12px;
  }
  a{
    color: #606266;
    text-decoration:none;
  }
</style>

