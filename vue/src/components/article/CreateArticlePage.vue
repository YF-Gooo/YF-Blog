<template>
    <el-container>
        <el-main>
            <el-col :xs="{span: 14, offset: 5}" :sm="{span: 12, offset: 2}" :md="{span: 6, offset: 2}" :lg="{span: 6, offset: 2}" :xl="{span: 6, offset: 2}">
                <el-divider content-position="left"><el-button type="text" size="medium" style="color:black">标题</el-button></el-divider>
                <el-input type="text" placeholder="请输入标题" v-model="title" maxlength="20" show-word-limit >
                </el-input>
            </el-col>
            <el-col :xs="{span: 14, offset: 5}" :sm="{span: 12, offset: 2}" :md="{span: 10, offset: 4}" :lg="{span: 10, offset: 4}" :xl="{span: 6, offset: 2}">
                <el-divider content-position="left"><el-button type="text" size="medium" style="color:black">描述</el-button></el-divider>
                <el-input type="textarea" placeholder="请输入描述" v-model="info" maxlength="50" show-word-limit>
                </el-input>
            </el-col>
            <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
                <mavon-editor class="editor" style="min-height: 600px" ref=md v-model="markdown" :language="language"  @imgAdd="imgAdd" @imgDel="imgDel" @save="saveDoc"></mavon-editor>
            </el-col>
            <el-col :xs="{span: 14, offset: 5}" :sm="{span: 10, offset: 2}" :md="{span: 5, offset: 2}" :lg="{span: 5, offset: 2}" :xl="{span: 5, offset: 4}">
                <el-divider content-position="left"><el-button type="text" size="medium" style="color:black">标签</el-button></el-divider>
                <el-input type="text" placeholder="请输入标签';'作为分割符" v-model="tag" maxlength="30" show-word-limit >
                </el-input>
            </el-col>
            <el-col :xs="24" :sm="{span: 10, offset: 2}" :md="{span: 10, offset: 5}" :lg="{span: 10, offset: 5}" :xl="{span: 10, offset: 5}">
                <el-button style="margin-top:30px;" round  @click="uploadDoc" :loading="uploading">发布</el-button>
                <el-button type="success" style="margin-top:30px;" round>草稿</el-button>
            </el-col>
        </el-main>
    </el-container>
</template>
<script>
// Local Registration
import { mavonEditor } from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
import * as API from "@/api/article/";
export default {
    name: 'createarticlepage',
    components: {
        mavonEditor
    },
    data(){
        return {
            title: '',
            info: '',
            img_file : {},
            markdown : "",
            tag:"",
            language : "zh-CN",
            uploading:false,
        }
    },
    methods: {
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
                title:_this.title,
                info: _this.info,
                markdown: markdown,
            };
            console.log(obj);
            console.log(text);
            API.createDoc(obj
                )
                .then(
                response => {
                    console.log(response)
                    _this.$message({
                    message: "上传成功",
                    type: "success"
                    });
                _this.$router.push("/articlepage/"+response.data.id);
                },
                response => console.log("上传失败" + response)
                );
        },
        
        uploadDoc() {
            let _this = this;
            _this.uploading=true
            let obj = {
                title:_this.title,
                info: _this.info,
                markdown :_this.markdown,
            };
            console.log(obj);
            API.createDoc(obj
                )
                .then(
                response => {
                    console.log(response)
                    _this.$message({
                    message: "上传成功",
                    type: "success"
                    });
                _this.uploading=false
                _this.$router.push("/articlepage/"+response.data.id);
                },
                response => console.log("上传失败" + response)
                );
        },
    }
}
</script>
<style>
   .editor{
       margin:2%;
       margin-right:0;
   }
  .el-row {
    margin-bottom: 20px;
    &:last-child {
      margin-bottom: 0;
    }
  }
</style>