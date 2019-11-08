<template>
    <div style="margin:2% 20%;">
        <el-row>
        <el-button style="margin-top:100px;" round  @click="getArticle(1)">发布</el-button>
        </el-row>
        <mavon-editor style="min-height: 600px;" ref=md v-model="markdown" :language="language" :toolbarsFlag="toolbarsFlag" :subfield="subfield" :defaultOpen="preview" :editable="editable"  ></mavon-editor>
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
            title: '',
            description: '',
            markdown : "",
            preview : "preview",
            editable : false,
            subfield : false,
            toolbarsFlag : false,
            language : "zh-CN",

        }
    },
    methods: {
        getArticle(id) {
            let _this = this;
            API.getArticle(id
                )
                .then(
                response => {
                    console.log(response)
                    _this.$message({
                    message: "上传成功",
                    type: "success"
                    });
                    _this.markdown=response.data.markdown
                },
                response => console.log("上传失败" + response)
                );
        },
    }
}
</script>
<style>

</style>