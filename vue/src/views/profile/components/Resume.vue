<template>
<div style="text-indent:0px;text-align:center">
    <!-- personal --><p></p>
    <h2 style="text-align: center;">个人简介</h2>
    <p style="margin-top:10px;text-align: center;font-size:18px; font-weight: bold;">YF</p>
    <p id="brief" style="text-align: center;">来自江苏无锡，资深宅男，芳龄25</p>
    <p id="brief" style="text-align: center;">菜鸟数据科学家，人工智Z工程师</p>
    <p id="brief" style="text-align: center;">野生的程序猿，喜欢唱，跳，rap～music～～</p>
    <p id="brief" style="text-align: center;">擅长写ugly前端的界面</p>
    <p id="brief" style="text-align: center;">热衷于研究把自己绕晕的后端架构</p>
    <p id="brief" style="text-align: center;">喜欢用躲在docker里，做一只安静的土拨鼠</p>
    <p id="brief" style="text-align: center;"><img :src="imgUrl"></p>
    <p><!--end personal --></p>
    <p><!-- tech tree --></p>
</div>
</template>

<script>
export default {
    data(){
        return {
            imgUrl:""
        }
    },
    mounted(){
        let urlTemp = "assets/mice.jpeg";
        this.imgUrl = require("@/"+urlTemp)
    }
}
</script>
