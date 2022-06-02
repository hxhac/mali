<template>
  <div class="wrap">
    <el-input v-model="text" placeholder="输入你要复制的文本" style="width:30%" />
    <el-button
      type="primary"
      class="copybut"
      :data-clipboard-text="text"
      size="medium"
      @click="copytext"
    >复制文本</el-button>
  </div>
</template>

<script>
import Clipboard from 'clipboard'

export default {
  name: 'Clipboard',
  data() {
    return {
      text: '' // 在  data-clipboard-text 写入你要复制的文本
    }
  },
  mounted() {
    // 在dom 挂载完成的时候执行一次，防止点击两次才能复制的问题
    /* this.copy = new Clipboard(".copybut"); */
    // 传入的必须是字符串
  },
  methods: {
    copytext() {
      const that = this
      /* this.copy = new Clipboard(".copybut"); */
      const clipboard = new Clipboard('.copybut')
      clipboard.on('success', function(e) {
        that.$message({
          message: '复制成功',
          type: 'success'
        })
        that.text = ''
        clipboard.destroy() // 成功之后删除当前 的 clipboard案例，防止每次叠加并多次执行
      })
      clipboard.on('error', function(e) {
        that.$message({
          message: '复制失败',
          type: 'success'
        })
      })
    }
  }
}
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.wrap {
	width: 1200px;
	margin: 0 auto;
}
</style>
