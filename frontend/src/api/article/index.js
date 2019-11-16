import axios from 'axios';

// 创建视频
const createDoc = form => axios.post('/api/v1/article', form).then(res => res.data);

const getArticle = id => axios.get(`/api/v1/article/${id}`).then(res => res.data);

const getArticleList = () => axios.get(`/api/v1/articles`).then(res => res.data);

const deleteArticle = id => axios.delete(`/api/v1/article/${id}`).then(res => res.data);
export {
  createDoc,
  getArticle,
  getArticleList,
  deleteArticle
};
