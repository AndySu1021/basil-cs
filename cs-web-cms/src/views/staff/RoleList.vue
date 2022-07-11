<template>
  <div class="app-container" v-loading="loading">
    <AppPanel title="角色列表">
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
    <AddRoleDialog ref="AddRoleDialog" @flush-parent="flushTable" />
    <EditRoleDialog ref="EditRoleDialog" @flush-parent="flushTable" />
  </div>
</template>

<script>
import AppTable from "@/components/AppTable";
import AppPanel from "@/components/AppPanel";
import AppPagination from "@/components/AppPagination";
import AddRoleDialog from "@/views/staff/components/AddRoleDialog";
import EditRoleDialog from "@/views/staff/components/EditRoleDialog";
import {apiDeleteRole, apiGetRoleList} from "@/api/role";

export default {
  name: 'RoleList',
  components: {
    AddRoleDialog,
    EditRoleDialog,
    AppTable,
    AppPanel,
    AppPagination,
  },
  data() {
    return {
      loading: false,
      initFilter: {
        name: '',
      },
      form: {
        name: '',
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
      this.$refs.AddRoleDialog.show()
    },
    openDialogEdit(id) {
      this.$refs.EditRoleDialog.show(id)
    },
    async fetchData(payload) {
      try {
        this.loading = true
        delete this.pagination.total
        const params = new URLSearchParams({...this.form, ...(payload)});
        const { data, pagination } = await apiGetRoleList(params.toString())
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
        await apiDeleteRole(id)
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
    }
  },
}
</script>
