<template>
  <div class="app-container" v-loading="loading">
    <AppPanel title="會員列表">
      <template v-slot:filter>
        <el-input v-model="form.mobile" placeholder="手機號" style="width: 160px;" />
        <el-input v-model="form.email" placeholder="電子信箱" style="width: 250px;" />
        <el-button icon="el-icon-search" type="primary" @click="fetchData(pagination)">搜尋</el-button>
        <el-button icon="el-icon-circle-close" @click="handleClear">清除</el-button>
      </template>
      <template v-slot:body>
        <AppTable
            :table-data="tableData"
            :table-columns="tableColumns"
        />
        <AppPagination :data="pagination" @change="handleChange" />
      </template>
    </AppPanel>
  </div>
</template>

<script>
import AppTable from '@/components/AppTable'
import AppPagination from '@/components/AppPagination'
import {apiGetStaffList} from "@/api/staff";
import {deepCopy} from "@/utils";
import AppPanel from "@/components/AppPanel";
import {apiGetMemberList} from "@/api/member";
import moment from "moment";

export default {
  name: 'RoomList',
  components: {
    AppPanel,
    AppTable,
    AppPagination,
  },
  created() {
    this.fetchData(this.pagination)
  },
  data() {
    return {
      loading: false,
      staffs: [],
      initFilter: {
        room_id: '',
        staff_id: 0,
        status: 0,
      },
      form: {
        room_id: '',
        staff_id: 0,
        status: 0,
      },
      pagination: {
        page: 1,
        page_size: 10,
        total: 0,
      },
      tableData: [],
      tableColumns: [
        { prop: 'id', label: '編號', align: "center", width: 100 },
        {
          prop: 'type',
          label: '類型',
          align: "center",
          render: (h, scope) => {
            switch (scope.row.type) {
              case 1:
                return <span>一般會員</span>
              case 2:
                return <span>訪客</span>
            }
          }
        },
        { prop: 'site_name', label: '所屬站點', align: "center" },
        { prop: 'name', label: '會員名稱', align: "center" },
        { prop: 'real_name', label: '會員姓名', align: "center" },
        { prop: 'mobile', label: '手機號', align: "center" },
        { prop: 'email', label: '信箱', align: "center", width: 200 },
        {
          prop: 'online_status',
          label: '狀態',
          width: 120,
          align: "center",
          render: (h, scope) => {
            switch (scope.row.online_status) {
              case 1:
                return <el-tag effect="dark" type="success">在線</el-tag>
              case 2:
                return <el-tag effect="dark" type="error">離線</el-tag>
            }
          }
        },
        {
          prop: 'created_at',
          label: '首次訪問',
          align: "center",
          width: 200,
          render: (h, scope) => {
            return moment(scope.row.created_at).format("YYYY-MM-DD HH:mm:ss")
          }
        },
      ]
    }
  },
  methods: {
    async fetchData(payload) {
      try {
        this.loading = true
        let params = deepCopy(this.form)
        if(params.room_id === '') {
          params.room_id = 0
        }
        const query = new URLSearchParams({...params, ...(payload)});
        const { data, pagination } = await apiGetMemberList(query.toString())
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
    handleClear() {
      this.form = { ...this.initFilter }
    },
  },
}
</script>
