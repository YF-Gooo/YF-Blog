import axios from 'axios';

// 创建视频
const signUp = form => axios.post('/api/v1/admin/signup', form).then(res => res.data);

// 读视频详情
const getUser = name => axios.get(`/api/v1/admin/getuser/${name}`).then(res => res.data);

// 读取视频列表
const signIn = form => axios.get('/api/v1/admin/signin').then(res => res.data);

export {
  signUp,
  getUser,
  signIn,
};

createDoc() {
  let _this = this;
  let obj = {
      title:_this.title,
      info: _this.description,
      content: _this.content,
  };
  API.createDoc(obj
      )
      .then(
      response => {
          console.log(response)
          _this.$message({
          message: "上传成功",
          type: "success"
          });
      },
      response => console.log("上传失败" + response)
      );
},