
<template>
  <div class="home">
    <el-container>
        <el-container>
        <el-main>
            <el-row>
                <div v-for="o in 8" :key="o" class="text item">
                    <el-card class="box-card" :body-style="{ padding: '0px',height:'150px'}" shadow="hover">
                    <el-col :span="20"> 
                    <div class="article-title">
                    {{'文章标题' + o }}
                    </div>
                    <div class="article-content">
                    {{'文章描述:'+'博客园是一个面向开发者的知识分享社区。自创建以来,博客园一直致力并专注于为开发者打造一个纯净的技术交流社区,推动并帮助开发者通过互联网分享知识,从而让更多'}}
                    </div>
                    
                    </el-col>
                    <el-col :span="4">
                    <div class="bottom clearfix">
                        <span class="article-comment">
                            <el-badge :value="1" class="item" type="primary">
                                <el-button size="small">评论</el-button>
                            </el-badge>
                            <el-badge :value="3" class="item" type="primary">
                                <el-button size="small">转发</el-button>
                            </el-badge>
                        </span>
                        <div class="time">
                        <time >{{ currentDate }}</time>
                        </div>
                    </div>
                    </el-col>
                    </el-card>
                </div>
            </el-row>
            <div class="block">
                <el-pagination
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
                :current-page="currentPage4"
                :page-sizes="[100, 200, 300, 400]"
                :page-size="100"
                layout="total, sizes, prev, pager, next, jumper"
                :total="400">
                </el-pagination>
            </div>
        </el-main>
        <el-aside width="400px">
            <el-calendar v-model="currentDate">
            </el-calendar>
            <el-carousel indicator-position="outside" style="margin-top: 100px;">
                <el-carousel-item v-for="item in 4" :key="item">
                <h3>{{ item }}</h3>
                </el-carousel-item>
            </el-carousel>
        </el-aside>
    </el-container>
    <el-footer>此博客归属于YF</el-footer>
    </el-container>
  </div>
</template>

<script>
// @ is an alias to /src
import store from "@/store";
import Navbar from '@/components/Navbar'
export default {
  methods: {
      handleSizeChange(val) {
        console.log(`每页 ${val} 条`);
      },
      handleCurrentChange(val) {
        console.log(`当前页: ${val}`);
      }
    },
  name: 'home',
  components: {
      Navbar,
    // Admin
  },
  data() {
    return {
      currentDate: new Date(),
      currentPage4: 4,
    };
  },
  store,
  
}
</script>


<style>
  .item {
    margin-right: 20px;
    }
  .time {
    font-size: 10px;
    color: #999;
    text-align: right;
  }
  
  .bottom {
    margin-top: 100px;
    margin-left: 20px;
    line-height: 12px;
    position: relative
  }

  .button {
    padding: 0;
    float: right;
  }

  .clearfix:before,
  .clearfix:after {
      display: table;
      content: "";
  }
  
  .clearfix:after {
      clear: both
  }
  .article-title{
    padding: 0%;
    text-align: left;
    margin-left: 20px;
  }
  .article-content{
    padding: 0%;
    padding-right: 20%;
    margin-top: 10px;
    font-size: 13px;
    margin-left: 20px;
    text-align: left;
  }
  .el-card{
    margin-left: 4%;
  }
  .el-card__body{ 
    text-align: left;
  }
  .article-comment{
      padding-left: 5%;
  }
  .el-carousel__item h3 {
    color: #475669;
    font-size: 18px;
    opacity: 0.75;
    line-height: 300px;
    margin: 0;
  }
  
  .el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
  }
  
  .el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
  }

</style>