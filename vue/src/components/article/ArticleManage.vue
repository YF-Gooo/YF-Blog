<template>
  <div>
    <el-table
      :data="articleList"
      style="width: 100%">
      <el-table-column
        label="日期"
        width="200">
        <template slot-scope="scope">
          <i class="el-icon-time"></i>
          <span style="margin-left: 10px">{{ scope.row.created_at| getTimestamp }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="作者"
        width="150">
        <template slot-scope="scope">
            <div slot="reference" class="name-wrapper">
              <el-tag size="medium">{{ scope.row.username }}</el-tag>
            </div>
        </template>
      </el-table-column>
      <el-table-column
        label="标题"
        width="150">
        <template slot-scope="scope">
          <span >{{ scope.row.title }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="简介"
        width="200">
        <template slot-scope="scope">
          <span >{{ scope.row.info }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="内容"
        width="300">
        <template slot-scope="scope">
          <span >{{ scope.row.markdown.substr(0, [30]) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="标签"
        width="100">
        <template slot-scope="scope">
          <span >{{ scope.row.tag }}</span>
        </template>
      </el-table-column>
      <el-table-column align= "center" header-align="center" label="操作">
        <template slot-scope="scope">
          <el-button size="mini" icon="el-icon-search" @click="handleCheck(scope.row)"></el-button>
          <el-button size="mini" type="primary" icon="el-icon-edit" @click="handleEdit(scope.row)"></el-button>
          <el-button size="mini" type="danger" icon="el-icon-delete" @click="handleDelete(scope.row)"></el-button>
        </template>
      </el-table-column>
    </el-table>
    <div class="block">
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="currentPage"
        :page-sizes="[14, 18]"
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
      limit: 14,
      currentPage :1
    }
  },
  async mounted(){
    let _this = this;
    if (_this.amsearchkw===""){
      _this.getUserArticleList(_this.start,_this.limit)
    } else{
      _this.getUserSearchArticleList(_this.start,_this.limit,_this.amsearchkw)
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
          if (_this.amsearchkw===""){
            _this.getUserArticleList(_this.start,_this.limit)
          } else{
            _this.getUserSearchArticleList(_this.start,_this.limit,_this.amsearchkw)
          }
        },
        getUserArticleList(start,limit) {
            let _this = this;
            API.getUserArticleList(start,limit)
                .then(
                response => {
                    console.log(response)
                    _this.articleList=response.data.items
                    _this.total = response.data.total
                },
                response => console.log("获取失败" + response)
                );
        },
        getUserSearchArticleList(start,limit,kw) {
            let _this = this;
            API.getUserSearchArticleList(start,limit,kw)
                .then(
                response => {
                    console.log(response)
                    _this.articleList=response.data.items
                    _this.total = response.data.total
                },
                response => console.log("获取失败" + response)
                );
        },
        handleCheck(row){
          window.location.href="/articlepage/"+row.id
        },
        handleEdit(row) {
          window.location.href="/updatearticlepage/"+row.id
        },
        handleDelete(row) {
          let _this = this;
          API.deleteArticle(row.id
          )
          .then(
            response => {
              console.log(response)
              if (response.status!==404) {
                _this.$message({
                  message: "删除成功",
                  type: "success"
                });
                _this.articleList = _this.articleList.filter(function(item) {
                  return item.id != row.id
                });
              } else{
                _this.$message({
                  message: "删除失败",
                  type: "fail"
                })
              }
            },
            response => console.log("删除失败" + response)
          );
        }
    },
  computed: {
    amsearchkw(){      
      return this.$store.state.amkw
    }
  },
  watch:{
    amsearchkw(){
        this.getUserSearchArticleList(0,this.limit,this.$store.state.amkw)
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
<style>
</style>