<template>
  <div id="sign_wrap">
    <h1>后台管理</h1>
    <el-input v-model="avatar" placeholder="头像"></el-input>
    <el-input v-model="name" placeholder="请输入用户名"></el-input>
    <el-input v-model="password" placeholder="请输入密码" type="password"></el-input>
    <el-input v-model="val_password" placeholder="请再次输入密码" type="val_password"></el-input>
    <el-button @click="signup">注册</el-button>
  </div>
</template>

<script>
import * as API from "@/api/user/";
export default {
  name: "signin",
  data() {
    return {
      avatar:"",
      name: "",
      password: "",
      val_password:"",
      hasName: false // 用户名被占
    };
  },
  mounted: function() {},
  methods: {
    signup: function() {
      let _this = this;
      if (this.name.length < 6) {
        this.$message.error("用户名不能为空或小于六个字符");
        return;
      }

      if (this.password.length < 6) {
        this.$message.error("密码不能为空或小于六个字符");
        return;
      }

      if (this.password !=this.val_password) {
        this.$message.error("两次输入密码不一样");
        return;
      }
      API.getUser(this.name).then(
        response => {
          if (response.body.name === _this.name) {
            _this.$message.error("该用户已存在");
            _this.name = "";
            // 由于异步，name的改变比正常流执行得慢，所以不能通过判断name去执行是否post
            // 所以我把post移入else中，而不是在外面通过判断name执行
          } else {
            let obj = {
              avatar:_this.avatar,
              name: _this.name,
              password: _this.password
            };

            API.signUp({
                userInfo: obj
              })
              .then(
                response => {
                  _this.$message({
                    message: "注册成功",
                    type: "success"
                  });
                },
                response => console.log("注册失败" + response)
              );
          }
        },
        response => console.log(response)
      );
    },
  }
};
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