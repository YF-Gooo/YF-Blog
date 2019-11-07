<template>
    <div style="margin:2% 10%;">
        <div style="margin: 20px 0;padding-left:20px;width:20%;">
        <el-divider content-position="left">标题</el-divider>
        <el-input type="text" placeholder="请输入标题" v-model="title" maxlength="20" show-word-limit >
        </el-input>
        </div>
        <div style="margin: 20px 0;padding-left:20px; width:30%;">
        <el-divider content-position="left">描述</el-divider>
        <el-input type="textarea" placeholder="请输入描述" v-model="description" maxlength="30" show-word-limit>
        </el-input>
        </div>
        <mavon-editor style="min-height: 600px;" ref=md v-model="text" @imgAdd="imgAdd" @imgDel="imgDel" @save="saveDoc" @change="updateDoc"></mavon-editor>
        <el-row>
            <el-button style="margin-top:30px;" round  @click="saveDoc">发布</el-button>
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
    name: 'editor',
    components: {
        mavonEditor
    },
    data(){
        return {
            title: '',
            description: '',
            img_file : {},
            text : ""
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
        updateDoc(markdown, text) {
            console.log("markdown内容: " + markdown);
            console.log("html内容:" + text);
        },
        saveDoc(markdown, text) {
            let _this = this;
            let obj = {
                title:_this.title,
                info: _this.description,
                content: text,
            };
            API.createDoc(obj
                )
                .then(
                response => {
                    console.log(response)
                    _this.$message({
                    message: "上传成功",
                    type: "success"
                    });
                },
                response => console.log("上传失败" + response)
                );
        },
    }
}
</script>
<style>
#editor {
    margin: auto;
    width: 80%;
    height: 580px;
}
.mveditor{
    margin: auto;
    width: 100%;
    
}
</style>