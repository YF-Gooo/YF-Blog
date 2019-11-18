import axios from 'axios';

// 创建视频
const createDoc = form => axios.post('/api/v1/article', form).then(res => res.data);

const getArticle = id => axios.get(`/api/v1/article/${id}`).then(res => res.data);

const updateArticle = (id,form) => axios.put(`/api/v1/article/${id}`,form).then(res => res.data);

const getArticleList = (start,limit) => axios.get('/api/v1/articles',{params:{start,limit}}).then(res => res.data);

const getUserArticleList = (start,limit) => axios.get('/api/v1/userarticles',{params:{start,limit}}).then(res => res.data);

const getSearchArticleList = (start,limit,kw) => axios.get('/api/v1/searcharticles',{params:{start,limit,kw}}).then(res => res.data);

const getUserSearchArticleList = (start,limit,kw) => axios.get('/api/v1/usersearcharticles',{params:{start,limit,kw}}).then(res => res.data);

const deleteArticle = id => axios.delete(`/api/v1/article/${id}`).then(res => res.data);

export {
  createDoc,
  getArticle,
  getArticleList,
  deleteArticle,
  updateArticle,
  getUserArticleList,
  getSearchArticleList,
  getUserSearchArticleList
};
