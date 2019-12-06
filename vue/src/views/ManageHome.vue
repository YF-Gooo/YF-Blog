<template>
    <div style = "height:100%">
        <el-container>
            <el-header>
                <el-col :xs="{span:10, offset: 0}" :sm="{span: 8, offset: 0}" :md="{span: 6, offset: 0}" :lg="{span: 4, offset: 0}" :xl="{span: 4, offset: 0}">
                    <el-input
                        placeholder="筛选"
                        v-model="text"
                        @keyup.enter.native="handlesearch">
                        <i slot="prefix" class="el-input__icon el-icon-search"></i>
                    </el-input>
                </el-col>
            </el-header>
            <el-raw>
                <el-col :xs="{span:4, offset: 0}" :sm="{span: 3, offset: 0}" :md="{span: 3, offset: 0}" :lg="{span: 3, offset: 0}" :xl="{span: 4, offset: 0}">
                    <SideBar></SideBar>
                </el-col>
                <el-col :xs="{span: 15, offset:3}" :sm="{span: 17, offset: 4}" :md="{span: 19, offset: 2}" :lg="{span: 19, offset: 1}" :xl="{span: 19, offset: 1}">
                    <ArticleManage></ArticleManage>
                </el-col>
            </el-raw>
        </el-container>
    </div>
</template>
<script>
import store from '@/store'
import SideBar from '@/layers/SideBar'
import ArticleManage from '@/components/article/ArticleManage'
export default {
    name: 'ManageHome',
    components: {
        SideBar,
        ArticleManage,
    },
    data() {
        return {
        text:"",
        };
    },
    store,
    methods: {
        handlesearch(){
        let _this = this;
        _this.$store.commit('changeAMSearchKW', _this.text)
        _this.text=""
        },
    }
}
</script>
<style>
</style>