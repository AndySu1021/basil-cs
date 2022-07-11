<template>
  <div class="app-container" v-loading="loading">
    <AppPanel title="常見問題列表">
      <template v-slot:filter>
        <el-input v-model="form.question" placeholder="問題" style="width: 200px;" />
        <el-select style="width: 100px;" v-model="form.status" placeholder="狀態">
          <el-option label="全部" :value="0"></el-option>
          <el-option label="啟用" :value="1"></el-option>
          <el-option label="禁用" :value="2"></el-option>
        </el-select>
        <el-button icon="el-icon-search" type="primary" @click="fetchData(pagination)">搜尋</el-button>
        <el-button icon="el-icon-circle-close" @click="handleClear">清除</el-button>
      </template>
      <template v-slot:action>
        <el-button type="primary" icon="el-icon-plus" @click="openDialogAdd">新增</el-button>
      </template>
      <template v-slot:body>
        <AppTable
            :table-data="tableData"
            :table-columns="tableColumns"
            :show-operation="true"
            :handle-edit="openDialogEdit"
            :handle-delete="handleDelete"
        />
        <AppPagination
            :data="pagination"
            @change="handleChange"
        />
      </template>
    </AppPanel>
    <AddFAQDialog ref="AddFAQDialog" @flush-parent="flushTable" />
    <EditFAQDialog ref="EditFAQDialog" @flush-parent="flushTable" />
  </div>
</template>

<script>
import AddFAQDialog from '@/views/cs-setting/faq/components/AddFAQDialog'
import EditFAQDialog from '@/views/cs-setting/faq/components/EditFAQDialog'
import AppTable from "@/components/AppTable";
import AppPagination from "@/components/AppPagination";
import AppPanel from "@/components/AppPanel";
import {apiDeleteFAQ, apiGetFAQList} from "@/api/faq";

export default {
  name: 'FAQList',
  components: {
    AppPanel,
    AppTable,
    AppPagination,
    AddFAQDialog,
    EditFAQDialog,
  },
  data() {
    return {
      loading: false,
      initFilter: {
        question: '',
        status: 0,
      },
      form: {
        question: '',
        status: 0,
      },
      pagination: {
        page: 1,
        page_size: 10,
        total: 0,
      },
      tableData: [],
      tableColumns: [
        { prop: 'id', label: '編號' },
        { prop: 'question', label: '問題' },
        { prop: 'answer', label: '答案' },
        {
          prop: 'status',
          label: '狀態',
          render: (h, scope) => {
            switch (scope.row.status) {
              case 1:
                return <el-tag effect="dark" type="success">啟用</el-tag>
              case 2:
                return <el-tag effect="dark" type="danger">禁用</el-tag>
            }
          }
        },
      ]
    }
  },
  created() {
    this.fetchData(this.pagination)
  },
  methods: {
    flushTable() {
      this.fetchData(this.pagination)
    },
    openDialogAdd() {
      this.$refs.AddFAQDialog.show()
    },
    openDialogEdit(id) {
      this.$refs.EditFAQDialog.show(id)
    },
    async fetchData(payload) {
      try {
        this.loading = true
        const params = new URLSearchParams({...this.form, ...(payload)});
        const { data, pagination } = await apiGetFAQList(params.toString())
        this.tableData = data
        this.pagination = pagination
      } catch (err) {
        console.log(err)
      } finally {
        this.loading = false
      }
    },
    handleChange(payload) {
      this.fetchData(payload)
    },
    async handleDelete(id) {
      try {
        await this.$confirmDelete()
        this.loading = true
        await apiDeleteFAQ(id)
        this.$showSuccessMessage("刪除成功")
        await this.fetchData(this.pagination)
        this.loading = false
      } catch (err) {
        console.log(err)
        this.loading = false
      }
    },
    handleClear() {
      this.form = { ...this.initFilter }
    },
  },
}
</script>
