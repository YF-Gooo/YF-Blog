<template>
  <div class="common-footer">
    <p class="copyright">
      <a href="http://www.miibeian.gov.cn/" target="_blank">{{ siteInfo.icp }}</a>
      <span>|</span>
      {{ siteInfo.copyright }}
      <span>|</span>
      <span v-html="siteInfo.copyright_desc"></span>
    </p>
  </div>
</template>

<script>
export default {
  data () {
    return {
      siteInfo: {
        icp: '粤ICP备17117756号',
        copyright: '版权所有 © 2018 - 2019',
        copyright_desc: '以商业目的使用本网站内容需获许可，非商业目的使用授权遵循CC BY-NC 4.0'
      }
    }
  }
}
</script>
<style>
  .common-footer{
    bottom:0;
    margin:15px 0;
    margin-top:100px;
    width:100%;
    font-weight:300;
    line-height:25px;
    text-align:center;

  }
</style>
