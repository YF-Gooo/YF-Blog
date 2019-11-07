<template>
  <div id="sign_wrap">
    <h1>后台管理</h1>
    <el-input v-model="user_name" placeholder="请输入用户名"></el-input>
    <el-input v-model="password" placeholder="请输入密码" type="password"></el-input>
    <el-button type="primary" @click="signin">登录</el-button>
  </div>
</template>

<script>
import * as API from "@/api/user/";
export default {
  name: "signin",
  data() {
    return {
      user_name: "",
      password: "",
    };
  },
  mounted: function() {},
  methods: {
    signin: function() {
      let _this = this;
      if (this.user_name.length < 6) {
        this.$message.error("用户名不能为空或小于六个字符");
        return;
      }

      if (this.password.length < 8) {
        this.$message.error("密码不能为空或小于八个字符");
        return;
      }
      let obj = {
        user_name: _this.user_name,
        password: _this.password
      };
      API.signIn(obj
        )
        .then(
          response => {
            console.log(response)
            _this.$message({
              message: "登录成功",
              type: "success"
            });
            delete _this.password;
            window.location.href="/"
          },
          response => console.log("登录失败" + response)
        );
      }
    }
  }
</script>

<style>
#sign_wrap {
  width: 300px;
  margin: 100px auto;
}

#sign_wrap h1 {
  color: #383a42;
  padding: 10px;
}

#sign_wrap div {
  padding-bottom: 20px;
}
</style>