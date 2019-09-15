<template>
  <div >
      <!-- 图片上传组件辅助-->
      <el-upload
        class="avatar-uploader"
        :action="serverUrl"
        name="file"
        :headers="header"
        :show-file-list="false"
        :on-success="uploadSuccess"
        :on-error="uploadError"
        :before-upload="beforeUpload">
      </el-upload>
      <quill-editor class="blog-editor"
        v-model="content"
        ref="myQuillEditor"
        :options="editorOption"
        @change="onEditorChange($event)"
      >
      </quill-editor>
    </div>
</template>
<script>
  const toolbarOptions = [
    ['bold', 'italic', 'underline', 'strike'],        // toggled buttons
    [{'header': 1}, {'header': 2}],               // custom button values
    [{'list': 'ordered'}, {'list': 'bullet'}],
    [{'indent': '-1'}, {'indent': '+1'}],          // outdent/indent
    [{'direction': 'rtl'}],                         // text direction
    [{'size': ['small', false, 'large', 'huge']}],  // custom dropdown
    [{'header': [1, 2, 3, 4, 5, 6, false]}],
    [{'color': []}, {'background': []}],          // dropdown with defaults from theme
    [{'font': []}],
    [{'align': []}],
    ['link', 'image','video'],
    ['clean']
 
  ]
  export default {
    data() {
      return {
        quillUpdateImg: false, // 根据图片上传状态来确定是否显示loading动画，刚开始是false,不显示
        content: null,
        imgQuality: 0.5, //压缩图片的质量
        max_width :500,
        max_height :300,
        editorOption: {
          placeholder: '',
          theme: 'snow',  // or 'bubble'
          modules: {
            toolbar: {
              container: toolbarOptions,
              handlers: {
                'image': function (value) {
                  if (value) {
                    // 触发input框选择图片文件
                    document.querySelector('.avatar-uploader input').click()
                  } else {
                    this.quill.format('image', false);
                  }
                }
              }
            }
          }
        },
        serverUrl: '/api/v1/uploadimage',  // 这里写你要上传的图片服务器地址
        header: {
          // token: sessionStorage.token
        } // 有的图片服务器要求请求头需要有token
      }
    },
    methods: {
      onEditorChange({editor, html, text}) {//内容改变事件
        console.log("---内容改变事件---")
        this.content = html
        console.log(html)
      },
      // upload.vue

 
      dataURItoBlob(dataURI, type) {
        var binary = atob(dataURI.split(',')[1]);
        var array = [];
        for(var i = 0; i < binary.length; i++) {
          array.push(binary.charCodeAt(i));
        }
        return new Blob([new Uint8Array(array)], {type: type});
      },
      
      beforeUpload(param) {
        //对图片进行压缩
        const imgSize = param.size / 1024 / 1024

          const _this = this
          return new Promise(resolve => {
            const reader = new FileReader()
            const image = new Image()
            image.onload = (imageEvent) => {
              const canvas = document.createElement('canvas');
              const context = canvas.getContext('2d');
              console.log(image.width)
              console.log(image.height)
              if(image.width > _this.max_width||image.height>_this.max_height) {
              const imgQuality=image.width/(_this.max_width)
              const width = image.width * (1/imgQuality)
              const height = image.height * (1/imgQuality)
              image.width = width
              image.height = height
              }
              console.log(image.width)
              console.log(image.height)
              canvas.width = image.width;
              canvas.height = image.height;
              context.clearRect(0, 0, image.width, image.height);
              context.drawImage(image, 0, 0, image.width, image.height);
              const dataUrl = canvas.toDataURL(param.type);
              const blobData = _this.dataURItoBlob(dataUrl, param.type);
              resolve(blobData)
            }
            reader.onload = (e => { image.src = e.target.result; });
            reader.readAsDataURL(param);
          })
      },
 
      uploadSuccess(res, file) {
        // res为图片服务器返回的数据
        // 获取富文本组件实例
        console.log(res);
        let quill = this.$refs.myQuillEditor.quill
        // 如果上传成功
        if (res.status == 200 ) {
          // 获取光标所在位置
          let length = quill.getSelection().index;
          // 插入图片  res.msg为服务器返回的图片地址
          quill.insertEmbed(length, "image", res.msg)
          // 调整光标到最后
          quill.setSelection(length + 1)
        } else {
          this.$message.error('图片插入失败')
        }
        // loading动画消失
        this.quillUpdateImg = false
      },
      // 富文本图片上传失败
      uploadError() {
        // loading动画消失
        this.quillUpdateImg = false
        this.$message.error('图片插入失败')
      }
    }
  }
</script>
<style>
.blog-editor {
  width: 90%;
  margin:0 auto;
  text-align: center;
}
.ql-container{
  min-height:200px; 
}
</style>


