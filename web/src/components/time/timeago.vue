<template>
  <timeago ref="timeago" :datetime="date" />
</template>

<script>
export default {
  name: 'Timeago',
  props: ['date'],
  data() {
    return {
      name: this.date // 把传过来的值赋值给新的变量
    }
  },
}
</script>

<style scoped>

</style>