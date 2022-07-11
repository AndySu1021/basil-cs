<template>
  <div class="app-container" v-loading="loading">
    <AppPanel title="站點列表" :has-filter="false">
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
      </template>
    </AppPanel>
    <AddSiteDialog ref="AddSiteDialog" @flush-parent="flushTable" />
    <EditSiteDialog ref="EditSiteDialog" @flush-parent="flushTable" />
  </div>
</template>

<script>
import AppTable from "@/components/AppTable";
import AppPanel from "@/components/AppPanel";
import AppPagination from "@/components/AppPagination";
import AddSiteDialog from "@/views/cms/site/components/AddSiteDialog";
import EditSiteDialog from "@/views/cms/site/components/EditSiteDialog";
import EditAvatarDialog from "@/layout/components/EditAvatarDialog";
import {apiDeleteSite, apiGetSiteList} from "@/api/site";

export default {
  name: 'SiteList',
  components: {
    AddSiteDialog,
    EditSiteDialog,
    AppTable,
    AppPanel,
    AppPagination,
    EditAvatarDialog,
  },
  data() {
    return {
      loading: false,
      tableData: [],
      tableColumns: [
        { prop: 'id', label: '編號', align: "center", width: 100 },
        { prop: 'name', label: '名稱', align: "center" },
        { prop: 'code', label: '代號', align: "center" },
        {
          prop: 'status',
          label: '狀態',
          align: "center",
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
    this.fetchData()
  },
  methods: {
    flushTable() {
      this.fetchData()
    },
    openDialogAdd() {
      this.$refs.AddSiteDialog.show()
    },
    openDialogEdit(id) {
      this.$refs.EditSiteDialog.show(id)
    },
    async fetchData() {
      try {
        this.loading = true
        const { data } = await apiGetSiteList()
        this.tableData = data
        this.loading = false
      } catch (err) {
        console.log(err)
        this.loading = false
      }
    },
    async handleDelete(id) {
      try {
        await this.$confirmDelete()
        this.loading = true
        await apiDeleteSite(id)
        this.$showSuccessMessage("刪除成功")
        await this.fetchData(this.pagination)
        this.loading = false
      } catch (err) {
        console.log(err)
        this.loading = false
      }
    },
  },
}
</script>
