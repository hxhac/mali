<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="category">
          <el-cascader
            ref="elcascader"
            v-model="searchInfo.category"
            :options="categoryOptions"
            :show-all-levels="false"
            :props="{
              value: 'ID',
              label: 'cateName',
              children: 'children',
              expandTrigger: 'hover',
              checkStrictly: true,
            }"
            clearable
            filterable
            @change="handleChange"
          />
        </el-form-item>
        <el-form-item label="brand">
          <el-select v-model="searchInfo.brand" clearable filterable>
            <el-option
              v-for="item in brandOptions"
              :key="item.ID"
              :label="item.brandName"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="goods">
          <el-input v-model="searchInfo.goodsName" placeholder="商品名称" clearable />
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
        <el-form-item label="more">
          <el-input v-model="searchInfo.more" placeholder="商品备注" clearable />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询
          </el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">新增
        </el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="text" @click="deleteVisible = false">取消
            </el-button>
            <el-button size="small" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              size="small"
              style="margin-left: 10px;"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
            >删除
            </el-button>
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

        <el-table-column align="left" label="商品名称" prop="goodsName,brand" min-width="35%">
          <template #default="scope">
            {{ scope.row.goods_brand.brandName }} {{ scope.row.goodsName }}
            <el-tag v-if="scope.row.isStarred" type="warning" size="small" effect="dark">
              <el-icon>
                <StarFilled />
              </el-icon>
            </el-tag>
            <el-tag v-if="scope.row.useTimes" type="info" size="small" effect="dark">
              <el-icon>
                {{ scope.row.useTimes }}
              </el-icon>
            </el-tag>
            <el-tag v-if="scope.row.more" size="small" effect="dark">
              <el-icon>
                <Document />
              </el-icon>
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="商品价格" prop="price" min-width="10%">
          <template #default="scope">¥{{ scope.row.price }}</template>
        </el-table-column>
        <el-table-column align="left" label="评分" prop="score" min-width="10%">
          <template #default="scope">
            <el-rate v-model="scope.row.score" disabled />
          </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组" min-width="10%">
          <template #default="scope">
            <el-button
              type="text"
              icon="edit"
              size="small"
              class="table-button"
              @click="updateGoodsEvaluationFunc(scope.row)"
            >变更
            </el-button>
            <el-button
              type="text"
              icon="delete"
              size="small"
              @click="deleteRow(scope.row)"
            >删除
            </el-button>
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
    <!--	  详情页弹窗-->
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">

        <el-form-item label="商品名称:">
          <el-input v-model="formData.goodsName" clearable placeholder="请输入" />
        </el-form-item>
	      <el-row type="flex" class="row-bg">
		      <el-col :span="8">
			      <el-form-item label="商品价格:">
				      <el-input-number v-model.number="formData.price" placeholder="请输入" />
			      </el-form-item>
		      </el-col>
		      <el-col :span="8">
			      <el-form-item label="品牌:" prop="brand">
				      <el-select v-model="formData.brand" placeholder="请选择" clearable filterable>
					      <el-option
						      v-for="item in brandOptions"
						      :key="item.ID"
						      :label="item.brandName"
						      :value="item.ID"
					      />
				      </el-select>
			      </el-form-item>
		      </el-col>
		      <el-col :span="8">
			      <el-form-item label="分类:" prop="category">
				      <el-cascader
					      ref="elcascader"
					      v-model="formData.category"
					      :options="categoryOptions"
					      :props="{
              value: 'ID',
              label: 'cateName',
              children: 'children',
              expandTrigger: 'hover',
            }"
					      clearable
					      filterable
					      @change="handleChange"
				      />
			      </el-form-item>
		      </el-col>
	      </el-row>
	      <el-row>
		      <el-col :span="8">
			      <el-form-item label="use_times:">
				      <el-input-number v-model.number="formData.useTimes" />
			      </el-form-item>
		      </el-col>
		      <el-col :span="8">
			      <el-form-item label="score:">
				      <el-rate v-model.number="formData.score" />
			      </el-form-item>
		      </el-col>
		      <el-col :span="8">
			      <el-form-item label="star:">
				      <el-switch
					      v-model="formData.isStarred"
					      active-color="#13ce66"
					      inactive-color="#ff4949"
					      active-text="是"
					      inactive-text="否"
					      clearable
				      />
			      </el-form-item>
		      </el-col>
	      </el-row>
        <el-form-item label="more字段:">
          <mavon-editor v-model="formData.more" style="min-height: 400px" />
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
  name: 'GoodsEvaluation',
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
      fileList: []
    }
  },
  computed: {},
  watch: {
    handleChange() {
      this.$refs.elcascader.dropDownVisible = false
    }
  },
  methods: {
  }
}
</script>

<script setup>
import {
  createGoodsEvaluation,
  deleteGoodsEvaluation,
  deleteGoodsEvaluationByIds,
  updateGoodsEvaluation,
  findGoodsEvaluation,
  getGoodsEvaluationList,
} from '@/api/goodsEvaluation'
import {
  getGoodsBrandList
} from '@/api/goodsBrand'
import {
  getGoodsCategoryList,
} from '@/api/goodsCategory'

// 全量引入格式化工具 请按需保留
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { Document, StarFilled } from '@element-plus/icons-vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  goodsName: '',
  price: 1,
  brand: null,
  category: null,
  score: null,
  isStarred: false,
  useTimes: 1,
  more: '',
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
  if (searchInfo.value.isStarred === '') {
    searchInfo.value.isStarred = null
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
  const table = await getGoodsEvaluationList({
    page: page.value,
    pageSize: pageSize.value, ...searchInfo.value
  })
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
const brandOptions = ref([])
const categoryOptions = ref([])
const setOptions = async() => {
  // 产品品牌
  const brand = await getGoodsBrandList({ pageSize: 999 })
  if (brand.code === 0) {
    brandOptions.value = brand.data.list
  }

  // 产品分类数据
  const category = await getGoodsCategoryList({ pageSize: 999 })
  if (category.code === 0) {
    categoryOptions.value = category.data.list
  }
}

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
    deleteGoodsEvaluationFunc(row)
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
  const res = await deleteGoodsEvaluationByIds({ ids })
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
const updateGoodsEvaluationFunc = async(row) => {
  const res = await findGoodsEvaluation({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.regoodsEvaluation
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteGoodsEvaluationFunc = async(row) => {
  const res = await deleteGoodsEvaluation({ ID: row.ID })
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
    goodsName: '',
    price: 1,
    brand: null,
    category: null,
    score: null,
    isStarred: false,
    useTimes: 1,
    more: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
  let res
  switch (type.value) {
    case 'create':
      res = await createGoodsEvaluation(formData.value)
      break
    case 'update':
      res = await updateGoodsEvaluation(formData.value)
      break
    default:
      res = await createGoodsEvaluation(formData.value)
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

// 弹窗控制标记
// const dialogDocVisible = ref(false)
//
// // 打开弹窗
// const openDialogDoc = () => {
// 	type.value = 'create'
// 	dialogDocVisible.value = true
// }
//
// // 关闭弹窗
// const closeDialogDoc = () => {
// 	dialogDocVisible.value = false
// }

</script>

<style lang="scss">
.el-table .cell {
  white-space: pre-line; // 单元格内空格展示为换行
}

// [elment ui 使用标签时两个标签会连在一起 - SegmentFault 思否](https://segmentfault.com/q/1010000011729618)
.el-tag + .el-tag {
  margin-left: 10px;
}

//.el-table__fixed-body-wrapper {
//	z-index: auto !important;
//}

//.table-expand {
//	padding-left: 60px;
//	font-size: 0;
//
//	label {
//		width: 90px;
//		color: #99a9bf;
//
//		.el-form-item {
//			margin-right: 0;
//			margin-bottom: 0;
//			width: 50%;
//		}
//
//	}
//}
.popover-box {
	width: auto;
	height: auto;
	max-width: 70%;
	max-height: 70%;
	//overflow: auto;
}

//.popover-box::-webkit-scrollbar {
//	display: none; /* Chrome Safari */
//}
</style>
