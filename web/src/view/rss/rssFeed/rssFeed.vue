<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="rss分类:">
          <el-select v-model="searchInfo.cateId" placeholder="请选择" clearable filterable>
            <el-option
              v-for="item in categoryOptions"
              :key="item.ID"
              :label="item.cateName"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="rss名称:">
          <el-input v-model="searchInfo.rssName" placeholder="商品名称" clearable />
        </el-form-item>
        <el-form-item label="starred">
          <el-select v-model="searchInfo.isStarred" clearable filterable>
            <el-option
              v-for="item in isStarredOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="pause">
          <el-select v-model="searchInfo.isPause" clearable filterable>
            <el-option
              v-for="item in isStarredOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="text" @click="deleteVisible = false">取消</el-button>
            <el-button size="small" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="rss名称" prop="rssName" min-width="30%">
          <template #default="scope">
            <el-link :href="scope.row.url" :underline="false">{{ scope.row.rssName }} &nbsp;</el-link>

<!--	          如果三个月未更新就标记-->
            <el-tag v-if="scope.row.lastUpdated && Date.parse(scope.row.lastUpdated) < new Date(ShowDate(90))" size="small">⚠️</el-tag>

            <el-tag v-if="scope.row.sourceUrl" size="small">
              <el-link :href="scope.row.sourceUrl" _target="blank" :icon="Link" />
            </el-tag>

            <el-tag v-if="scope.row.cateId" size="small" effect="dark">
              {{ scope.row.rss_category.cateName }}
            </el-tag>
            
            <el-tag v-if="scope.row.isStarred" type="success" size="small" effect="dark">
              <el-icon>
                <Check />
              </el-icon>
            </el-tag>
            
            <el-tag v-if="scope.row.isPause" type="danger" size="small" effect="dark">
              <el-icon>
                <Close />
              </el-icon>
            </el-tag>
          </template>
        </el-table-column>
        
        <el-table-column align="left" label="评分" prop="score" min-width="20%">
          <template #default="scope">
            <el-rate v-model="scope.row.score" disabled />
          </template>
        </el-table-column>
        
        <el-table-column align="left" label="备注" prop="remark" min-width="40%" />
        <el-table-column align="left" label="按钮组" min-width="10%">
          <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateRssFeedFunc(scope.row)">变更</el-button>
            <el-button type="text" icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="rss分类:">
          <el-select v-model="formData.cateId" placeholder="请选择" clearable filterable>
            <el-option
              v-for="item in categoryOptions"
              :key="item.ID"
              :label="item.cateName"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="url:">
          <el-input v-model="formData.url" clearable placeholder="请输入" />
        </el-form-item>
	      <el-form-item label="源url:">
		      <el-input v-model="formData.sourceUrl" clearable placeholder="请输入" />
	      </el-form-item>
        <el-form-item label="keywords:">
          <el-input v-model="formData.keywords" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="is_starred:">
          <el-switch v-model="formData.isStarred" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
        <el-form-item label="是否暂停:">
          <el-switch v-model="formData.isPause" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
        <el-form-item label="评分:">
          <el-rate v-model.number="formData.score" />
        </el-form-item>
	      <el-form-item label="最后更新时间:">
		      <el-input v-model="formData.lastUpdated" disabled />
	      </el-form-item>
	      
	      <el-form-item label="备注:">
		      <el-input v-model="formData.remark" type="textarea" :autosize="{ minRows: 4 }" clearable placeholder="请输入" />
	      </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'RssFeed',
  data() {
    return {
      value: '请选择',
      isStarredOptions: [{
        value: true,
        label: '是'
      }, {
        value: false,
        label: '否'
      }],
      fileList: [],
      date: '2021-09-01',
    }
  },
}
</script>

<script setup>
import {
  createRssFeed,
  deleteRssFeed,
  deleteRssFeedByIds,
  updateRssFeed,
  findRssFeed,
  getRssFeedList
} from '@/api/rssFeed'
import { getRssCategoryList } from '@/api/rssCategory'

// 全量引入格式化工具 请按需保留
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { Close, Check, Link, WarningFilled } from '@element-plus/icons-vue'
import { formatTimeToStr, ShowDate } from '@/utils/date.js'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  cateId: null,
  url: '',
  keywords: '',
  isPause: false,
  isStarred: false,
  score: null,
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.isPause === '') {
    searchInfo.value.isPause = null
  }
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getRssFeedList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

const categoryOptions = ref([])
// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
  const category = await getRssCategoryList({ pageSize: 999 })
  if (category.code === 0) {
    categoryOptions.value = category.data.list
  }
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteRssFeedFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
  const res = await deleteRssFeedByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateRssFeedFunc = async(row) => {
  const res = await findRssFeed({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.rerssFeed
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteRssFeedFunc = async(row) => {
  const res = await deleteRssFeed({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    cateId: null,
    url: '',
    keywords: '',
    isPause: false,
    isStarred: false,
    score: null,
  }
}
// 弹窗确定
const enterDialog = async() => {
  let res
  switch (type.value) {
    case 'create':
      res = await createRssFeed(formData.value)
      break
    case 'update':
      res = await updateRssFeed(formData.value)
      break
    default:
      res = await createRssFeed(formData.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功'
    })
    closeDialog()
    getTableData()
  }
}

</script>

<style lang="scss">
.el-table .cell {
  white-space: pre; // 单元格内空格展示为换行
}

.el-tag + .el-tag {
  margin-left: 10px;
}
</style>
