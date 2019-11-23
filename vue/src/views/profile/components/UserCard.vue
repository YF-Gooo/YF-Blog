<template>
  <el-card style="margin-bottom:20px;">
    <div slot="header" class="clearfix">
      <i class="el-icon-s-custom"></i>
      <span>About me</span>
    </div>

    <div class="user-profile">
      <div class="box-center">
          <div><span>Name:{{name}}</span></div>
          <div><span>Email:{{email}}</span> </div>
          <div><span>Role:{{role}}</span> </div>
      </div>
    </div>

    <div class="user-bio">
      <div class="user-education user-bio-section">
        <div class="user-bio-section-header"><i class="el-icon-notebook-2"></i><span>Education</span></div>
        <div class="user-bio-section-body">
          <div class="text-muted">
            Master of Data Science, University College London
          </div>
        </div>
      </div>

      <div class="user-skills user-bio-section">
        <div class="user-bio-section-header"><i class="el-icon-monitor"></i><span>Skills</span></div>
        <div class="user-bio-section-body">
          <div class="progress-item">
            <span>Vue</span>
            <el-progress :percentage="40" />
          </div>
          <div class="progress-item">
            <span>JavaScript</span>
            <el-progress :percentage="41" />
          </div>
          <div class="progress-item">
            <span>Css</span>
            <el-progress :percentage="42" />
          </div>
          <div class="progress-item">
            <span>Python</span>
            <el-progress :percentage="80" />
          </div>
          <div class="progress-item">
            <span>Golang</span>
            <el-progress :percentage="50" />
          </div>
          <div class="progress-item">
            <span>Machine Learning</span>
            <el-progress :percentage="60" />
          </div>
          <div class="progress-item">
            <span>Big Data</span>
            <el-progress :percentage="45" />
          </div>
          <div class="progress-item">
            <span>Devops</span>
            <el-progress :percentage="39" />
          </div>
        </div>
      </div>
    </div>
  </el-card>
</template>
<script>

export default {
  data() {
    return {
      name: 'YF',
      email: 'ucakyj@gmail.com',
      avatar: '',
      role: 'DATA SCIENTIST'
    }
  }
}
</script>

<style lang="scss" scoped>
 .box-center {
   margin: 0 auto;
   display: table;
 }

 .text-muted {
   color: #777;
 }

 .user-profile {
   .user-name {
     font-weight: bold;
   }

   .box-center {
     padding-top: 10px;
     text-align:left;
     span {
       color: #606266;
       font-size: 14px;
     }
   }

   .user-role {
     padding-top: 10px;
     font-weight: 400;
     font-size: 14px;
   }

   .box-social {
     padding-top: 30px;

     .el-table {
       border-top: 1px solid #dfe6ec;
     }
   }

   .user-follow {
     padding-top: 20px;
   }
 }

 .user-bio {
   margin-top: 20px;
   color: #606266;

   span {
     padding-left: 4px;
   }

   .user-bio-section {
     font-size: 14px;
     padding: 15px 0;

     .user-bio-section-header {
       border-bottom: 1px solid #dfe6ec;
       padding-bottom: 10px;
       margin-bottom: 10px;
       font-weight: bold;
     }
   }
 }
</style>