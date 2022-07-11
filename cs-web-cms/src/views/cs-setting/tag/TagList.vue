<template>
  <div class="app-container" v-loading="loading">
    <AppPanel title="標籤列表">
      <template v-slot:filter>
        <el-input v-model="form.name" placeholder="名稱" style="width: 200px;" />
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
    <AddTagDialog ref="AddTagDialog" @flush-parent="flushTable" />
    <EditTagDialog ref="EditTagDialog" @flush-parent="flushTable" />
  </div>
</template>

<script>
import AddTagDialog from '@/views/cs-setting/tag/components/AddTagDialog'
import EditTagDialog from '@/views/cs-setting/tag/components/EditTagDialog'
import {apiDeleteTag, apiGetTagList} from "@/api/tag";
import AppTable from "@/components/AppTable";
import AppPagination from "@/components/AppPagination";
import AppPanel from "@/components/AppPanel";

export default {
  name: 'TagList',
  components: {
    AppPanel,
    AppTable,
    AppPagination,
    AddTagDialog,
    EditTagDialog,
  },
  data() {
    return {
      loading: false,
      initFilter: {
        name: '',
        status: 0,
      },
      form: {
        name: '',
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
        { prop: 'name', label: '名稱' },
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
      this.$refs.AddTagDialog.show()
    },
    openDialogEdit(id) {
      this.$refs.EditTagDialog.show(id)
    },
    async fetchData(payload) {
      try {
        this.loading = true
        const params = new URLSearchParams({...this.form, ...(payload)});
        const { data, pagination } = await apiGetTagList(params.toString())
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
        await apiDeleteTag(id)
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
