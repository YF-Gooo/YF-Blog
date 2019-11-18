<template>
    <div style="margin:2% 10%;">
        <div style="margin: 20px 0;padding-left:20px;width:20%;">
        <el-divider content-position="left"><span style = "font-size:20px;font-weight:bolder;color:grey;">标题</span></el-divider>
        <el-input type="text" placeholder="请输入标题" v-model="title" maxlength="20" show-word-limit >
        </el-input>
        </div>
        <div style="margin: 20px 0;padding-left:20px; width:30%;">
        <el-divider content-position="left" style = "font-size:15px;font-weight:bolder;color:grey;">描述</el-divider>
        <el-input type="textarea" placeholder="请输入描述" v-model="info" maxlength="30" show-word-limit>
        </el-input>
        </div>
        <mavon-editor style="min-height: 600px" ref=md v-model="markdown" @imgAdd="imgAdd" @imgDel="imgDel" @save="saveDoc"></mavon-editor>
        <div style="margin: 20px 0;padding-left:20px;width:20%;">
            <el-divider content-position="left" style = "font-size:15px;font-weight:bolder;color:grey;">标签</el-divider>
            <el-input type="text" placeholder="请输入标签';'作为分割符" v-model="tag" maxlength="30" show-word-limit >
            </el-input>
        </div>
        <el-row>
            <el-button style="margin-top:30px;" round  @click="updateDoc">更新</el-button>
            <el-button type="success" style="margin-top:30px;" round>草稿</el-button>
        </el-row>
    </div>
</template>
<script>
// Local Registration
import { mavonEditor } from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
import * as API from "@/api/article/";
export default {
    name: 'updatearticlepage',
    components: {
        mavonEditor
    },
    data(){
        return {
            id :0,
            title: '',
            info: '',
            markdown : "",
            tag:""
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
                response => console.log("获取失败" + response)
                );
        },
        // 绑定@imgAdd event
        imgAdd(pos, $file){
            // 第一步.将图片上传到服务器.
            var formdata = new FormData();
            formdata.append('image', $file);
            this.img_file[pos] = $file;
            this.$axios({
                url: '/api/v1/uploadimage',
                method: 'post',
                data: formdata,
                headers: { 'Content-Type': 'multipart/form-data'},
            }).then((url) => {
                // 第二步.将返回的url替换到文本原位置![...](0) -> ![...](url)
                /**
                * $vm 指为mavonEditor实例，可以通过如下两种方式获取
                * 1. 通过引入对象获取: `import {mavonEditor} from ...` 等方式引入后，`$vm`为`mavonEditor`
                * 2. 通过$refs获取: html声明ref : `<mavon-editor ref=md ></mavon-editor>，`$vm`为 `this.$refs.md`
                */
                this.$refs.md.$img2Url(pos, url.data.msg);
            })
        },

        imgDel(pos) {
            delete this.img_file[pos];
        },

        // updateDoc(markdown, text) {
        //     console.log("markdown内容: " + markdown);
        //     console.log("html内容:" + text);
        // },

        saveDoc(markdown, text) {
            let _this = this;
            let obj = {
                id:_this.id,
                title:_this.title,
                info: _this.info,
                markdown: markdown,
            };
            console.log(text);
            API.updateArticle(_this.id,obj
                )
                .then(
                response => {
                console.log(response)
                if (response.status!==404) {
                    _this.$message({
                    message: "更新成功",
                    type: "success"
                    });
                } else{
                    _this.$message({
                    message: "更新失败",
                    type: "fail"
                    })
                }
                },
                response => console.log("更新失败" + response)
                );
        },
        
        updateDoc() {
            let _this = this;
            let obj = {
                id:_this.id,
                title:_this.title,
                info: _this.info,
                markdown :_this.markdown,
            };
            console.log(obj);
            API.updateArticle(_this.id,obj
                )
                .then(
                response => {
                console.log(response)
                if (response.status!==404) {
                    _this.$message({
                    message: "更新成功",
                    type: "success"
                    });
                } else{
                    _this.$message({
                    message: "更新失败",
                    type: "fail"
                    })
                }
                },
                response => console.log("更新失败" + response)
                );
        },
    }
}
</script>
<style>
#editor {
    margin: auto;
    width: 50%;
    height: 580px;
}
.mveditor{
    margin: auto;
    width: 50%;
    
}
</style>