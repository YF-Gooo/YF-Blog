<template>
    <div>
        <p class="title">
          <span>{{this.title}}</span>
        </p>
        <p class="info">
          <span>作者：{{this.username}}</span>
        </p>
        <p class="info">
          <span>简介：{{this.info}}</span>{{this.info}}
        </p>
        <p class="content">
        <mavon-editor style="min-height: 600px;" ref=md v-model="markdown" :language="language" :toolbarsFlag="toolbarsFlag" :subfield="subfield" :defaultOpen="preview" :editable="editable" ></mavon-editor>
        </p>
    </div>
</template>
<script>
import * as API from "@/api/article/";
import { mavonEditor } from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
export default {
    name: 'articlepage',
    components: {
        mavonEditor
    },
    data(){
        return {
            id :0,
            title: '',
            info: '',
            markdown : "",
            username : "",
            preview : "preview",
            editable : false,
            subfield : false,
            toolbarsFlag : false,
            language : "zh-CN",
            ishljs:false,
            boxShadow:false,

        }
    },
    mounted(){
        this.id = this.$route.params.id
        this.getArticle(this.id)
    },
    methods: {
        getArticle(id) {
            let _this = this;
            API.getArticle(id
                )
                .then(
                response => {
                    console.log(response)
                    _this.title = response.data.title
                    _this.info = response.data.info
                    _this.markdown=response.data.markdown
                    _this.username=response.data.username
                },
                response => console.log("上传失败" + response)
                );
        },
    }
}
</script>
<style>
    p{
        padding-bottom:10px;
    }
    .title{
        font-size : 30px;
        font-weight : bolder;
        text-align:left;
        color: #666;
    }
    .info{
        font-size : 15px;
        font-weight : bolder;
        text-align:left;
        color: #666;
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
</style>