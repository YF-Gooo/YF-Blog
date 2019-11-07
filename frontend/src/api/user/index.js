import axios from 'axios';

// 注册
const signUp = form => axios.post('/api/v1/user/register', form).then(res => res.data);

// 获取用户
const getUser = name => axios.get(`/api/v1/admin/getuser/${name}`).then(res => res.data);

// 登录
const signIn = form => axios.post('/api/v1/user/login',form).then(res => res.data);

// 登出
const logOut = () => axios.delete('/api/v1/user/logout').then(res => res.data);

export {
  signUp,
  getUser,
  signIn,
  logOut,
};
