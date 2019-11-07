import axios from 'axios';

// 创建视频
const createDoc = form => axios.post('/api/v1/article', form).then(res => res.data);


export {
  createDoc
};
