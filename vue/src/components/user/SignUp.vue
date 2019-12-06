<template>
  <div id="signup_wrap">
    <h1>用户注册</h1>
    <el-input v-model="nickname" placeholder="NickName"></el-input>
    <el-input v-model="user_name" placeholder="请输入用户名"></el-input>
    <el-input v-model="password" placeholder="请输入密码" type="password"></el-input>
    <el-input v-model="password_confirm" placeholder="请再次输入密码" type="password"></el-input>

    <el-upload
      class="avatar-uploader"
      action="https://jsonplaceholder.typicode.com/posts/"
      :show-file-list="false"
      :on-success="handleAvatarSuccess"
      :before-upload="beforeAvatarUpload" :tip="上传头像">
      <img v-if="imageUrl" :src="imageUrl" class="avatar">
      <i v-else style="line-height:120px;" class="el-icon-plus avatar-uploader-icon"></i>
      <div class="el-upload__text">头像上传</em></div>
      <div class="el-upload__tip" slot="tip">只能上传jpg/png文件，且不超过500kb</div>
    </el-upload>
    <el-button @click="signup">注册</el-button>
  </div>
</template>

<script>
import * as API from "@/api/user/";
export default {
  name: "signup",
  data() {
    return {
      nickname:"",
      user_name: "",
      password: "",
      password_confirm:"",
      imgurl:""
    };
  },
  methods: {
    signup: function() {
      let _this = this;
      if (this.nickname.length < 3 ||this.nickname.length >15) {
        this.$message.error("用户名不能为空或小于三个字符");
        return;
      }

      if (this.user_name.length < 3||this.user_name.length >15) {
        this.$message.error("用户名不能为空或小于六个字符");
        return;
      }

      if (this.password.length < 8 || this.password.length >40) {
        this.$message.error("密码不能为空或小于八个字符");
        return;
      }

      if (this.password !=this.password_confirm) {
        this.$message.error("两次输入密码不一样");
        return;
      }
      
      let obj = {
        nickname:_this.nickname,
        user_name: _this.user_name,
        password: _this.password,
        password_confirm:_this.password_confirm,
      };
      API.signUp(obj
        )
        .then(
          response => {
            console.log(response)
            _this.$message({
              message: "注册成功",
              type: "success"
            });
            _this.$router.push("/signin");
          },
          response => console.log("注册失败" + response)
        );
    },
    handleAvatarSuccess(res, file) {
        this.imageUrl = URL.createObjectURL(file.raw);
      },
    beforeAvatarUpload(file) {
      const isJPG = file.type === 'image/jpeg';
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isJPG) {
        this.$message.error('上传头像图片只能是 JPG 格式!');
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!');
      }
      return isJPG && isLt2M;
    }
  }
}
</script>

<style >
#signup_wrap {
  width: 300px;
  margin: 100px auto;
}

#signup_wrap h1 {
  color: #383a42;
  padding: 10px;
}

#signup_wrap div {
  padding-bottom: 20px;
}
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409EFF;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 150px;
  height: 100px;
  line-height: 100px;
  text-align: center;
}
</style>