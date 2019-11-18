<template>
    <div style = "height:100%">
        <el-container>
        <el-aside width="250px"><SideBar></SideBar></el-aside>
        <el-main><ArticleManage></ArticleManage></el-main>
        </el-container>
    </div>
</template>
<script>
import SideBar from '@/layers/SideBar'
import ArticleManage from '@/components/article/ArticleManage'
export default {
    name: 'ManageHome',
    components: {
        SideBar,
        ArticleManage,
    }
}
</script>
<style>
</style>