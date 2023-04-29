<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="分类:" :rules="[{required: true}]">
          <el-select v-model="searchInfo.categoryId" placeholder="请选择" clearable filterable required>
            <el-option
              v-for="item in categoryOptions"
              :key="item.ID"
              :label="item.categoryName"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="是否使用中" :rules="[{required: true}]">
          <el-select v-model="searchInfo.isUse" clearable filterable>
            <el-option
              v-for="item in isUseOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="标签">
          <el-select v-model="searchInfo.appLabel" clearable filterable>
            <el-option
              v-for="item in labelOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="工具名称:">
          <el-input v-model="searchInfo.appName" placeholder="工具名称" clearable />
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

        <el-table-column align="left" label="app功能" prop="target" min-width="20%" />

        <el-table-column align="left" label="app名称" prop="appName" min-width="20%">
          <template #default="scope">

            <el-link v-if="scope.row.appUrl" :href="scope.row.appUrl" type="primary">
              {{ scope.row.appName }}
            </el-link> &nbsp;
            <template v-else>
              {{ scope.row.appName }}
            </template> &nbsp;

            <el-tag v-if="scope.row.appLabel >= 0" :type="labelOptions[scope.row.appLabel].color" size="small" effect="dark" :hit="false">
              {{ labelOptions[scope.row.appLabel].label }}
            </el-tag>

          </template>
        </el-table-column>

        <el-table-column align="left" label="评分" prop="score" min-width="15%">
          <template #default="scope">
            <el-rate v-model="scope.row.score" disabled />
          </template>
        </el-table-column>

        <el-table-column align="left" label="评价" prop="appRemark" min-width="25%" />

        <el-table-column align="left" label="按钮组" min-width="20%">
          <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateAppManageFunc(scope.row)">变更</el-button>
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
        <el-form-item label="app功能:">
          <el-input v-model="formData.target" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="app名称:">
          <el-input v-model="formData.appName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="app的url:">
          <el-input v-model="formData.appUrl" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="分类id:">
          <el-select v-model="formData.categoryId" placeholder="请选择" clearable filterable>
            <el-option
              v-for="item in categoryOptions"
              :key="item.ID"
              :label="item.categoryName"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="isUse:">
          <el-select v-model="formData.isUse" placeholder="请选择" clearable filterable>
            <el-option
              v-for="item in isUseOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="app标签:">
          <el-select v-model="formData.appLabel" placeholder="请选择" clearable filterable>
            <el-option
              v-for="item in labelOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="评分:">
          <el-rate v-model.number="formData.score" />
        </el-form-item>

        <el-form-item label="评价:">
          <el-input v-model="formData.appRemark" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="more字段:">
          <Md v-model="formData.appMore" />
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
  name: 'AppManage',
  data() {
    return {
      labelOptions: [{
        value: 0,
        label: 'CORE',
        color: 'success'
      }, {
        value: 1,
        label: 'ENHANCE',
        color: 'warning'
      }, {
        value: 2,
        label: 'DEL',
        color: 'danger'
      }],
      isUseOptions: [{
        value: 1,
        label: '使用中'
      }, {
        value: 0,
        label: '已删除'
      }]
    }
  }
}
</script>

<script setup>
import {
  createAppManage,
  deleteAppManage,
  deleteAppManageByIds,
  updateAppManage,
  findAppManage,
  getAppManageList
} from '@/api/appManage'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { Link } from '@element-plus/icons-vue'
import { getAppCategoryList } from '@/api/appCategory'
import Md from '@/components/md/md.vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  appLabel: 0,
  appMore: '',
  appName: '',
  appRemark: '',
  appUrl: '',
  categoryId: 1,
  isUse: 1,
  target: ''
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(30)
const tableData = ref([])
const searchInfo = ref({
  categoryId: 1,
  // appLabel: 0,
  isUse: 1,
})

// 重置
const onReset = () => {
  searchInfo.value = {
    categoryId: 1,
    // appLabel: 0,
    isUse: 1
  }
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 30
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
  const table = await getAppManageList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const categoryOptions = ref([])
// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
  const category = await getAppCategoryList({ pageSize: 999 })
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
    deleteAppManageFunc(row)
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
  const res = await deleteAppManageByIds({ ids })
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
const updateAppManageFunc = async(row) => {
  const res = await findAppManage({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reappManage
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteAppManageFunc = async(row) => {
  const res = await deleteAppManage({ ID: row.ID })
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
    appLabel: 0,
    appMore: '',
    appName: '',
    appRemark: '',
    appUrl: '',
    categoryId: 1,
    isUse: 1
  }
}
// 弹窗确定
const enterDialog = async() => {
  let res
  switch (type.value) {
    case 'create':
      res = await createAppManage(formData.value)
      break
    case 'update':
      res = await updateAppManage(formData.value)
      break
    default:
      res = await createAppManage(formData.value)
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
