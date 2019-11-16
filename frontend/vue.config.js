module.exports = {
  devServer: {
    proxy: {
      '^/api': {
        target: 'https://127.0.0.1:3000',
        ws: true,
        changeOrigin: true,
      },
      "^/static":{
        target: 'https://127.0.0.1:3000',
        ws: true,
        changeOrigin: true,
      }
    },
  },
};
