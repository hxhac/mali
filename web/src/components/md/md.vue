<template>
  <mavon-editor
    ref="md"
    style="min-height: 300px"
    @imgAdd="imgAdd"
    @imgDel="imgDel"
  />
</template>

<script>

import axios from 'axios'
import { useUserStore } from '@/pinia/modules/user'
import { ref } from 'vue'

const base = ref(import.meta.env.VITE_BASE_API).value
const userStore = useUserStore()

export default {
  name: 'Md',
  methods: {
    imgAdd(pos, $file) {
      // 第一步.将图片上传到服务器.
      var formdata = new FormData()
	    // beforeImageUpload($file)
      formdata.append('file', $file)
      this.uploadFileRequest(`/fileUploadAndDownload/upload`, formdata).then(resp => {
        var json = resp.data.data // 取出上传成功后的url
        var msg = resp.msg

        if (resp.status === 200) {
          this.$refs.md.$imglst2Url([[pos, json.file.url]])
        } else {
          this.$message({ type: resp.status, message: msg })
        }
      })
    },
    uploadFileRequest(url, params) {
      console.log(base)
      return axios({
        method: 'post',
        url: `${base}${url}`,
        data: params,
        headers: {
          'Content-Type': 'multipart/form-data',
          'x-token': userStore.token
        }
      })
    },
    // imgDel(pos) {
    //   console.log(pos)
    //   deleteFile(pos).then(resp => {
    // 	  var json = resp.data.data // 取出上传成功后的url
    // 	  var msg = resp.msg
    // 	  if (resp.status === 200) {
    //       this.$message({ type: resp.status, message: msg })
    // 	  }
    //   })
    // },
  }
}
</script>

<style scoped>

</style>