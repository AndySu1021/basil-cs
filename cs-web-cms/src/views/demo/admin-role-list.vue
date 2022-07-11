<template>
  <div class="app-container" v-loading="loading">
    <AppPanel title="管理組列表">
      <template v-slot:filter>
        <el-input v-model="form.filter.roomID" placeholder="房間編號" style="width: 100px;" />
        <el-select style="width: 150px;" v-model="form.filter.staffID" placeholder="職員">
          <el-option label="全部" :value="0"></el-option>
          <el-option v-for="(staff,idx) in staffs" :key="idx" :label="staff.name" :value="staff.id"></el-option>
        </el-select>
        <el-input v-model="form.filter.content" placeholder="訊息內容" style="width: 250px;" />
        <el-button type="primary" @click="fetchData">搜尋</el-button>
      </template>
      <template v-slot:action>
        <el-button type="primary" icon="el-icon-plus" @click="openDialogAdd">新增</el-button>
      </template>
      <template v-slot:body>
        <AppTable
            :table-data="tableData"
            :table-columns="tableColumns"
            :show-operation="true"
            :handle-edit="handleEdition"
            :handle-delete="handleDeletion"
        />
        <AppPagination
            :data="form.pagination"
            @change="handleChange"
        />
      </template>
    </AppPanel>
    <AddDialog ref="AddDialog" />
  </div>
</template>

<script>
import AppTable from "@/components/AppTable";
import AppPanel from "@/components/AppPanel";
import AppPagination from "@/components/AppPagination";
import AddDialog from "@/views/admin-manage/components/AddDialog";

export default {
  name: 'AdminRoleList',
  components: {
    AppTable,
    AppPanel,
    AppPagination,
    AddDialog,
  },
  data() {
    return {
      form: {
        filter: {
          roomID: '',
          staffID: 0,
          content: '',
        },
        pagination: {
          page: 1,
          pageSize: 10,
          total: 300
        }
      },
      loading: false,
      tableData: [{
        date: '2016-05-02',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄',
        address1: '上海市普陀区金沙江路 1516 弄',
        address2: '上海市普陀区金沙江路 1516 弄',
        address3: '上海市普陀区金沙江路 1516 弄',
      }, {
        date: '2016-05-04',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1517 弄',
        address1: '上海市普陀区金沙江路 1516 弄',
        address2: '上海市普陀区金沙江路 1516 弄',
        address3: '上海市普陀区金沙江路 1516 弄',
      }, {
        date: '2016-05-01',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1519 弄',
        address1: '上海市普陀区金沙江路 1516 弄',
        address2: '上海市普陀区金沙江路 1516 弄',
        address3: '上海市普陀区金沙江路 1516 弄',
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1516 弄',
        address1: '上海市普陀区金沙江路 1516 弄',
        address2: '上海市普陀区金沙江路 1516 弄',
        address3: '上海市普陀区金沙江路 1516 弄',
      }],
      tableColumns: [
        { prop: 'date', label: '日期', width: 120 },
        { prop: 'name', label: '姓名', width: 120 },
        { prop: 'address', label: '地址' },
        { prop: 'address1', label: '地址1', width: 100 },
        { prop: 'address2', label: '地址2' },
        { prop: 'address3', label: '地址3' },
      ]
    }
  },
  methods: {
    async fetchData() {
      try {
        this.loading = true
        const { data } = await apiAdminRoleList()
        this.tableData = data.admin_roles
      } catch (err) {
        console.log(err)
      } finally {
        this.loading = false
      }
    },
    flushTable() {
      this.fetchData()
    },
    openDialogAdd() {
      this.$refs.AddDialog.show()
    },
    openDialogEdit(admin_role_id) {
      const payload = { dialogType: 'EDIT', admin_role_id }
      this.$refs.DialogRole.show(payload)
    },
    async handleDelete(admin_role_id) {
      try {
        await this.$confirmDelete()
        this.loading = true
        const { metadata } = await apiAdminRoleDelete({ admin_role_id })
        this.$showSuccessMessage(metadata.desc)
        await this.fetchData()
      } catch (err) {
        console.log(err)
      } finally {
        this.loading = false
      }
    },
    handleChange(payload) {
      console.log(payload)
    },
    handleEdition() {
      this.$refs.DialogAddStaff.show()
    },
    handleDeletion() {
      console.log("delete")
    },
  },
  created() {
    this.fetchData()
  }
}
</script>

<style lang="scss" scoped>
.app-panel {
  height: 500px;
  .panel-title {
    font-size: 28px;
    font-weight: bold;
  }
  .panel-filter{
    margin-top: 20px;
    background-color: white;
    padding: 20px;
    & > * {
      margin-left: 10px;
    }
    & > *:first-child {
      margin-left: 0;
    }
  }
}
</style>
