<template>
  <div class="article-list" >
    <div class="post" v-for="(item,index) of articleList" :key="index">
      <router-link :to="'/articlepage/'+item.id" class="title">
          {{item.title}}
      </router-link>
      <p class="info">
          <span>摘要：</span>{{item.info}}
      </p>
      <p class="content">
          <span>内容：</span>{{item.markdown.substr(0, [100])}}
      </p>
      <ul class="list-inline">
        <li>
          <span>作者：{{item.username}}</span>
        </li>
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
        <li>
          <span>时间：{{item.created_at| getTimestamp}}</span>
        </li>
      </ul>
    </div>
    <div class="block">
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="currentPage"
        :page-sizes="[6, 12]"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total">
      </el-pagination>
    </div>
  </div>
</template>

<script>
import * as API from "@/api/article/";
export default {
  data() {
    return {
      articleList:[],
      total: 0,
      start: 0,
      limit: 6,
      currentPage :1
    }
  },
  async mounted(){
    let _this = this;
    if (_this.searchkw===""){
      _this.getArticleList(_this.start,_this.limit)
    } else{
      _this.getSearchArticleList(_this.start,_this.limit,_this.searchkw)
    }
  },
  methods: {
        handleSizeChange(val) {
          let _this = this;
          _this.limit=val
        },
        handleCurrentChange(val) {
          let _this = this;
          _this.currentPage=val
          _this.start=_this.limit*(_this.currentPage-1)
          if (_this.searchkw===""){
            _this.getArticleList(_this.start,_this.limit)
          } else{
            _this.getSearchArticleList(_this.start,_this.limit,_this.searchkw)
          }
        },
        getArticleList(start,limit) {
            let _this = this;
            API.getArticleList(start,limit)
                .then(
                response => {
                    console.log(response)
                    _this.articleList=response.data.items
                    _this.total = response.data.total
                },
                response => console.log("获取失败" + response)
                );
        },
        getSearchArticleList(start,limit,kw) {
            let _this = this;
            API.getSearchArticleList(start,limit,kw)
                .then(
                response => {
                    console.log(response)
                    _this.articleList=response.data.items
                    _this.total = response.data.total
                },
                response => console.log("获取失败" + response)
                );
        },
    },
  computed: {
    searchkw(){      
      return this.$store.state.kw
    }
  },
  watch:{
    searchkw(){
        this.getSearchArticleList(0,this.limit,this.$store.state.kw)
    }
  },
  filters: {
      getTimestamp: function (timestamp) {
        console.log(timestamp)
        if (timestamp) {
          var date = new Date(timestamp * 1000); //时间戳为10位需*1000(为秒)，时间戳为13位的话不需乘1000(毫秒)
          var Y = date.getFullYear();
          var M =
            date.getMonth() + 1 < 10 ?
            "0" + (date.getMonth() + 1) :
            date.getMonth() + 1;
          var D = date.getDate() < 10 ? "0" + date.getDate() : date.getDate();
          var h = date.getHours() < 10 ? "0" + date.getHours() : date.getHours();
          var m =
            date.getMinutes() < 10 ? "0" + date.getMinutes() : date.getMinutes();
          var s = date.getSeconds();
          console.log(Y + "-" + M + "-" + D + " " + h + ":" + m + ":" + s);
          return Y + "-" + M + "-" + D + " " + h + ":" + m + ":" + s;
        }
      }
    }
}
</script>

<style lang="scss" scoped>
.article-list {
  margin: 0 20%;
  margin-top:30px;
  .title{
    font-size : 20px;
    font-weight : bolder;

  }
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
    padding-left: 0px;
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
