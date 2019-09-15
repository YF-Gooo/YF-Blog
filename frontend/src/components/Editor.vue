<template>
    <mavon-editor class="mveditor" ref=md @imgAdd="$imgAdd" @imgDel="$imgDel"></mavon-editor>
</template>
<script>
// Local Registration
import { mavonEditor } from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
export default {
    name: 'editor',
    components: {
        mavonEditor
        // or 'mavon-editor': mavonEditor
    },
    methods: {
    // 绑定@imgAdd event
    $imgAdd(pos, $file){
        // 第一步.将图片上传到服务器.
        var formdata = new FormData();
        formdata.append('image', $file);
        axios({
            url: 'server url',
            method: 'post',
            data: formdata,
            headers: { 'Content-Type': 'multipart/form-data' },
        }).then((url) => {
            // 第二步.将返回的url替换到文本原位置![...](./0) -> ![...](url)
            // $vm.$img2Url 详情见本页末尾
            this.$refs.md.$img2Url(pos, url);
        })
    }
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
    height: 800px;
}
</style>