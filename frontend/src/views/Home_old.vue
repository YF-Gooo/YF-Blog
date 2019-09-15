<template>
  <div class="home">
    <el-row>
      <el-col :span="2"> </el-col>
      <el-col :span="18" v-for="(o, index) in 6" :key="o" :offset="0">
        <el-card :body-style="{ padding: '0px'}" shadow="hover">
          <div class="article" style="padding: 0px;">
            <span>好吃的汉堡</span>
            <div>好吃的汉堡</div>
            <div class="bottom clearfix">
              <time class="time">{{ currentDate }}</time>
            </div>
          </div>
        </el-card>
      
      </el-col>
    </el-row>
  </div>
</template>

<script>
// @ is an alias to /src
import store from "@/store";
export default {
  name: 'home',
  components: {
    // Admin
  },
  data() {
    return {
      currentDate: new Date()
    };
  },
  store
}
</script>


<style>
  .time {
    font-size: 13px;
    color: #999;
  }
  
  .bottom {
    margin-top: 13px;
    line-height: 12px;
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
  .article{
    height: 100px;
    display: inline-block;
    text-align: left
  }
</style>