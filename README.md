# vue-go-blog
本博客项目基于vue以及gin框架搭建<br>
本意是验证自己前后端能力。<br>
本博客会用于记录我的日常生活，同时也作为我的学习笔记。

# docker(用于快速启mysql和redis)
    cd docker 
    docker-compose up -d
    
# singo(使用golang的gin框架搭建后端接口，使用gorm操控数据库)
    cd backend
    export GOPROXY=https://mirrors.aliyun.com/goproxy/
    export GO111MODULE=on
    go run main.go

# frontend(采用vuejs编写博客前端界面)
    cd frontend
    npm install
    npm run serve
# 界面展示
## 首页
![Alt text](./images/home.png)

## 编辑页
![Alt text](./images/editor.png)

## 管理页
![Alt text](./images/manage.png)

# To list:
    1.登录页面(Done)
    2.图片上传接口(Done)
    3.文章接口(Done)
    4.文章显示(Done)
    5.首页展示(Done)
    6.编辑页展示(Done)
    7.搜索功能(Done)
    8.分页功能(Done)
    9.文章管理(Done)
    10.redis统计访问量
    11.打包docker镜像
    12.搭建gitlab ci/cd实现一键部署
    13.增加视频上传以及管理功能
    14.增加人工智能模块
    
# 参考了
    https://github.com/bydmm/singo
    https://github.com/pppercyWang/twitter-blog-vue

