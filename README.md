# vue-go-blog
本博客项目基于vue以及gin框架搭建<br>
本意是验证自己前后端能力。<br>
本博客会用于记录我的日常生活，同时也作为我的学习笔记。


# singo
    cd backend
    export GOPROXY=https://mirrors.aliyun.com/goproxy/
    export GO111MODULE=on
    go run main.go

# frontend
    cd singo
    npm install
    npm run serve

# to list:
    1.登录页面(Done)
    2.图片上传接口(Done)
    3.文章接口(Done)
    4.文章显示(Done)
    5.首页展示(Done)
    6.编辑页展示(Done)
    7.搜索功能(Pending)
    8.文章管理
    
# 参考了
    https://github.com/bydmm/singo
    https://github.com/pppercyWang/twitter-blog-vue


sudo add-apt-repository \
  　　　　"deb [arch=amd64] https://download.docker.com/linux/ubuntu \
  　　　　$(lsb_release -cs) \
  　　　 stable"