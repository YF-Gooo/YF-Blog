<template>
    <el-row :gutter="10">
        <el-menu router :default-active="activeIndex" class="el-menu-demo" mode="horizontal" background-color="#eee" active-text-color="#009a61" @select="handleSelect">
            <el-col :xs="24" :sm="24" :md="4" :lg="4" :xl="4">
                <div style="text-align:center;line-height:50px;">
                    <span style="font-size:30px;font-weight:bold;">Y</span>
                    <span style="font-size:30px;font-weight:bold;">F</span>
                    <span style="color:green;font-size:30px;">B</span>
                    <span style="font-size:25px;">l</span>
                    <span style="font-size:25px;">o</span>
                    <span style="font-size:25px;">g</span>
                </div>
            </el-col>
            <el-col :xs="12" :sm="12" :md="2" :lg="2" :xl="2"><el-menu-item index="/">首页</el-menu-item></el-col>
            <el-col :xs="12" :sm="12" :md="2" :lg="2" :xl="2"><el-menu-item index="/about">关于我</el-menu-item></el-col>
            <el-col v-if="user!==''" :xs="12" :sm="12" :md="2" :lg="2" :xl="2"><el-menu-item index="/createarticlepage">写文章</el-menu-item></el-col>
            <el-col v-if="user!==''" :xs="12" :sm="12" :md="2" :lg="2" :xl="2"><el-menu-item index="/managehome">管理屋</el-menu-item></el-col>
            <el-col :xs="24" :sm="{span: 16, offset: 2}" :md="{span: 4, offset: 0}" :lg="{span: 4, offset: 0}" :xl="{span: 4, offset: 0}" style="padding-top:8px">
              <el-input
                placeholder="请输入内容"
                v-model="text"
                @keyup.enter.native="handlesearch"
                >
                <i slot="prefix" class="el-input__icon el-icon-search"></i>
            </el-input>
          </el-col>
          <div v-if="user!==''">
            <el-col :xs="{span: 5, offset: 8}" :sm="{span: 2, offset: 1}" :md="2" :lg="{span: 1, offset: 3}" :xl="{span: 1, offset: 3}" style="height:56px;text-align:center;"><el-avatar shape="square" :size="56" :src="avtarUrl"></el-avatar></el-col>
            <el-col :xs="4" :sm="2" :md="2" :lg="2" :xl="1" style="padding-top:15px">
              <el-dropdown @command="handleCommand">
                <span class="el-dropdown-link" style="color:grey;font-weight:bolder;font-size:15px;">
                  {{user}}<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                <el-dropdown-menu slot="dropdown" trigger="click">
                  <el-dropdown-item command="/">首页</el-dropdown-item>
                  <el-dropdown-item command="/about">关于我</el-dropdown-item>
                  <el-dropdown-item command="/createarticlepage">写文章</el-dropdown-item>
                  <el-dropdown-item command="/managehome">管理屋</el-dropdown-item>
                  <el-dropdown-item disabled>其他</el-dropdown-item>
                  <el-dropdown-item divided ><el-button size="mini" type="info" @click="logout">注销</el-button></el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>
            </el-col>
          </div>
          <div v-else>
            <el-col :xs="12" :sm="{span: 2, offset: 0}" :md="{span: 2, offset: 3}" :lg="{span: 2, offset: 3}" :xl="{span: 2, offset: 3}"><el-menu-item index="/signin" >登录</el-menu-item></el-col>
            <el-col :xs="12" :sm="2" :md="2" :lg="2" :xl="2" ><el-menu-item index="/signup"><span style="border-radius:4px;background-color:#009a61;padding:15px 25px;color:white;">注册</span></el-menu-item></el-col>
          </div>
        </el-menu>
    </el-row>
</template>
<script>
import * as API from "@/api/user/";
import store from '@/store'
export default {
  data() {
    return {
      avtarUrl: "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png",
      user:"",
      text:"",
    };
  },
  store,
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
    handlesearch(){
      let _this = this;
      _this.$store.commit('changeSearchKW', _this.text)
      _this.$router.push("/");
      _this.text=""
    },
    handleCommand(command) {
      let _this = this;
      _this.$router.push(command);
      }
    }
  }
</script>
<style>
  #search:hover {
    background-color:#eee;
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
