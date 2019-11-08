<template>
  <div class="article-list" >
    <div class="post" v-for="(item,index) of articleList" :key="index">
      <h2 class="title">
          {{item.title}}
      </h2>
      <p class="info">
          <span>摘要：</span>{{item.info}}
      </p>
      <p class="content">
          <span>内容：</span>{{item.markdown}}
      </p>
      <ul class="list-inline">
        <li>
          <span class="link-black text-sm">
            <i class="el-icon-share" />
            Share
          </span>
        </li>
        <li>
          <span class="link-black text-sm">
            <i class="el-icon-apple">Like </i>
          </span>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import * as API from "@/api/article/";
export default {
  data() {
    return {
      articleList:[]
    }
  },
  async mounted(){
        this.getArticleList()
  },
  methods: {
        getArticleList() {
            let _this = this;
            API.getArticleList()
                .then(
                response => {
                    console.log(response)
                    _this.articleList=response.data
                },
                response => console.log("获取失败" + response)
                );
        },
    }
}
</script>

<style lang="scss" scoped>
.article-list {
  margin: 0 20%;
  margin-top:30px;
  .article-block {
    .description {
      display: block;
      padding: 2px 0;
    }

    .articlename{
      font-size: 16px;
      color: #000;
    }


    .img-circle {
      border-radius: 50%;
      width: 40px;
      height: 40px;
      float: left;
    }
    .info{
      font-weight: bolder;
    }
    span {
      font-weight: 500;
      font-size: 12px;
    }
  }

  .post {
    font-size: 14px;
    border-bottom: 1px solid #d2d6de;
    margin-bottom: 15px;
    padding-bottom: 15px;
    color: #666;
    text-align:left;
    padding-left: 10px;
    .image {
      width: 100%;
      height: 100%;

    }

    .article-images {
      padding-top: 20px;
    }
  }

  .list-inline {
    padding-left: 10px;
    margin-left: -5px;
    list-style: none;

    li {
      display: inline-block;
      padding-right: 5px;
      padding-left: 5px;
      font-size: 13px;
    }

    .link-black {

      &:hover,
      &:focus {
        color: #999;
      }
    }
  }

}

.box-center {
  margin: 0 auto;
  display: table;
}

.text-muted {
  color: #777;
}
</style>
