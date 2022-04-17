<template>
	<div>
		<div class="gva-search-box">
			<el-form :inline="true" :model="searchInfo" class="demo-form-inline">
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
						<el-button icon="delete" size="small" style="margin-left: 10px;"
						           :disabled="!multipleSelection.length"
						           @click="deleteVisible = true">删除
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
				<el-table-column type="selection" width="55"/>
				<el-table-column align="left" label="商品名称" prop="name" min-width="10%"/>
				<el-table-column align="left" label="商品分类" prop="goods_category.cate_name"
				                 min-width="5%"/>
				<el-table-column align="left" label="价格" prop="price" min-width="5%"/>
				<el-table-column align="left" label="综合评分" prop="rank" min-width="5%"/>
				<el-table-column align="left" label="remark字段" prop="remark" min-width="30%"/>
				<el-table-column align="left" label="按钮组" min-width="10%">
					<template #default="scope">
						<el-button type="text" icon="edit" size="small" class="table-button"
						           @click="updateGoodsEvaluationFunc(scope.row)">变更
						</el-button>
						<el-button type="text" icon="delete" size="small"
						           @click="deleteRow(scope.row)">删除
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
		<el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
			<el-form :model="formData" label-position="right" label-width="80px">
				<el-form-item label="name字段:">
					<el-input v-model="formData.name" clearable placeholder="请输入"/>
				</el-form-item>
				<el-form-item label="price字段:">
					<el-input v-model.number="formData.price" clearable placeholder="请输入"/>
				</el-form-item>
				<el-form-item label="cateId字段:">
					<!--          <el-input v-model.number="formData.goods_category.cate_name" clearable placeholder="请输入"/>-->
					<!--          <el-input v-model="formData.goods_category.cate_name" clearable placeholder="请输入"/>-->
					<el-select v-model="formData.goods_category.cate_name" placeholder="请选择下拉选择"
					           clearable
					           :style="{width: '100%'}">
						<el-option v-for="(item, index) in dataOptions" :key="index"
						           :label="item.label"
						           :value="item.value" :disabled="item.disabled"></el-option>
					</el-select>
				</el-form-item>
				<!--				<el-form-item label="rank字段:">-->
				<!--					<el-input-number v-model="formData.rank" style="width:100%" :precision="2"-->
				<!--					                 clearable/>-->
				<!--				</el-form-item>-->
				
				<el-form-item label="remark字段:">
					<el-input v-model="formData.remark" clearable placeholder="请输入"/>
				</el-form-item>
				<!--        <el-form-item label="userId字段:">-->
				<!--          <el-input v-model.number="formData.userId" clearable placeholder="请输入"/>-->
				<!--        </el-form-item>-->
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
			dataOptions: [],//下拉菜单
		}
	},
	methods: {
		categoryInfo() {
			this.dataOptions.push({value: '请选择', label: ''})
			// this.categoryInfo({}).then(response => {
			// 	response.data.forEach(item => {
			//
			// 	})
			// })
			f
		}
	},
	mounted() {
		this.categoryInfo()
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
	getGoodsEvaluationList
} from '@/api/goodsEvaluation'

// 全量引入格式化工具 请按需保留
import {getDictFunc, formatDate, formatBoolean, filterDict} from '@/utils/format'
import {ElMessage, ElMessageBox} from 'element-plus'
import {ref} from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
	name: '',
	price: 0,
	cateId: 0,
	rank: 0,
	detail: '',
	remark: '',
	userId: 0,
	goods_category: {cate_name: '', cate_id: 0},
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
const getTableData = async () => {
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
const setOptions = async () => {
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
		deleteGoodsEvaluationFunc(row)
	})
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async () => {
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
	const res = await deleteGoodsEvaluationByIds({ids})
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
const updateGoodsEvaluationFunc = async (row) => {
	const res = await findGoodsEvaluation({ID: row.ID})
	type.value = 'update'
	if (res.code === 0) {
		formData.value = res.data.regoodsEvaluation
		dialogFormVisible.value = true
	}
}

// 删除行
const deleteGoodsEvaluationFunc = async (row) => {
	const res = await deleteGoodsEvaluation({ID: row.ID})
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
		name: '',
		price: 0,
		cateId: 0,
		rank: 0,
		detail: '',
		remark: '',
		userId: 0,
		goods_category: []
	}
}
// 弹窗确定
const enterDialog = async () => {
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
</script>

<style>
</style>
