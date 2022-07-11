<template>
  <div class="app-container" v-loading="loading">
    <AppPanel title="快捷消息列表">
      <template v-slot:filter>
        <el-input v-model="form.title" placeholder="標題" style="width: 200px;" />
        <el-input v-model="form.content" placeholder="內容" style="width: 200px;" />
        <el-button icon="el-icon-search" type="primary" @click="fetchData(pagination)">搜尋</el-button>
        <el-button icon="el-icon-circle-close" @click="handleClear">清除</el-button>
      </template>
      <template v-slot:action>
        <el-button type="primary" icon="el-icon-plus" @click="openDialogAdd">新增</el-button>
        <el-button type="warning" icon="el-icon-collection-tag" @click="openAddCategoryDialog">分類</el-button>
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
    <AddCategoryDialog ref="AddCategoryDialog" @flush-parent="flushTable" />
    <AddFastReplyDialog ref="AddFastReplyDialog" @flush-parent="flushTable" />
    <EditFastReplyDialog ref="EditFastReplyDialog" @flush-parent="flushTable" />
  </div>
</template>

<script>
import AddFastReplyDialog from '@/views/cs-setting/fast-reply/components/AddFastReplyDialog'
import EditFastReplyDialog from '@/views/cs-setting/fast-reply/components/EditFastReplyDialog'
import {apiDeleteFastReply, apiGetFastReplyList} from "@/api/fast-reply";
import AddCategoryDialog from "@/views/cs-setting/fast-reply/components/AddCategoryDialog";
import AppTable from "@/components/AppTable";
import AppPagination from "@/components/AppPagination";
import AppPanel from "@/components/AppPanel";

export default {
  name: 'FastReplyList',
  components: {
    AppPanel,
    AppTable,
    AppPagination,
    AddFastReplyDialog,
    EditFastReplyDialog,
    AddCategoryDialog,
  },
  data() {
    return {
      loading: false,
      initFilter: {
        title: '',
        content: '',
        status: 0,
      },
      form: {
        title: '',
        content: '',
        status: 0,
      },
      pagination: {
        page: 1,
        page_size: 10,
        total: 0,
      },
      tableData: [],
      tableColumns: [
        { prop: 'id', label: '編號', width: 100 },
        { prop: 'category', label: '分類', width: 130 },
        { prop: 'title', label: '標題', width: 130 },
        { prop: 'content', label: '內容' },
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
      this.$refs.AddFastReplyDialog.show()
    },
    openAddCategoryDialog() {
      this.$refs.AddCategoryDialog.show()
    },
    openDialogEdit(id) {
      this.$refs.EditFastReplyDialog.show(id)
    },
    async fetchData(payload) {
      try {
        this.loading = true
        const params = new URLSearchParams({...this.form, ...(payload)});
        const { data, pagination } = await apiGetFastReplyList(params.toString())
        this.tableData = data
        this.pagination = pagination
        this.loading = false
      } catch (err) {
        console.log(err)
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
        await apiDeleteFastReply(id)
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
