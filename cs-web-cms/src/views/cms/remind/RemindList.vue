<template>
  <div class="app-container" v-loading="loading">
    <AppPanel title="提醒列表">
      <template v-slot:filter>
        <el-input v-model="form.content" placeholder="內容" style="width: 200px;" />
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
    <AddRemindDialog ref="AddRemindDialog" @flush-parent="flushTable" />
    <EditRemindDialog ref="EditRemindDialog" @flush-parent="flushTable" />
  </div>
</template>

<script>
import AppTable from "@/components/AppTable";
import AppPanel from "@/components/AppPanel";
import AppPagination from "@/components/AppPagination";
import AddRemindDialog from "@/views/cms/remind/components/AddRemindDialog";
import EditRemindDialog from "@/views/cms/remind/components/EditRemindDialog";
import EditAvatarDialog from "@/layout/components/EditAvatarDialog";
import {apiDeleteRemind, apiGetRemindList} from "@/api/remind";

export default {
  name: 'RemindList',
  components: {
    AddRemindDialog,
    EditRemindDialog,
    AppTable,
    AppPanel,
    AppPagination,
    EditAvatarDialog,
  },
  data() {
    return {
      loading: false,
      initFilter: {
        content: '',
        status: 0,
      },
      pagination: {
        page: 1,
        page_size: 10,
        total: 0,
      },
      form: {
        content: '',
        status: 0,
      },
      tableData: [],
      tableColumns: [
        { prop: 'id', label: '編號', width: 100 },
        { prop: 'title', label: '標題', width: 200 },
        { prop: 'content', label: '內容' },
        {
          prop: 'status',
          label: '狀態',
          align: "center",
          width: 100,
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
      this.$refs.AddRemindDialog.show()
    },
    openDialogEdit(id) {
      this.$refs.EditRemindDialog.show(id)
    },
    async fetchData(payload) {
      try {
        this.loading = true
        delete this.pagination.total
        const params = new URLSearchParams({...this.form, ...(payload)});
        const { data, pagination } = await apiGetRemindList(params.toString())
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
        await apiDeleteRemind(id)
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
