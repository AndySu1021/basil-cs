<template>
  <div class="app-container" v-loading="loading">
    <AppPanel title="房間列表">
      <template v-slot:filter>
        <el-input v-model="form.room_id" placeholder="房間編號" style="width: 100px;" />
        <el-select style="width: 150px;" v-model="form.staff_id" placeholder="職員">
          <el-option label="全部" :value="0"></el-option>
          <el-option v-for="(staff,idx) in staffs" :key="idx" :label="staff.name" :value="staff.id"></el-option>
        </el-select>
        <el-select style="width: 100px;" v-model="form.status" placeholder="狀態">
          <el-option label="全部" :value="0"></el-option>
          <el-option label="等待中" :value="1"></el-option>
          <el-option label="排隊中" :value="2"></el-option>
          <el-option label="服務中" :value="3"></el-option>
          <el-option label="已關閉" :value="4"></el-option>
        </el-select>
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
import {apiGetRoomList} from "@/api/room";
import {deepCopy} from "@/utils";
import AppPanel from "@/components/AppPanel";
import moment from "moment";

export default {
  name: 'RoomList',
  components: {
    AppPanel,
    AppTable,
    AppPagination,
  },
  created() {
    this.fetchStaffList()
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
        { prop: 'id', label: '編號' },
        { prop: 'member_name', label: '會員名稱' },
        { prop: 'staff_name', label: '職員名稱' },
        { prop: 'tag_name', label: '標籤' },
        {
          prop: 'status',
          label: '狀態',
          render: (h, scope) => {
            switch (scope.row.status) {
              case 1:
                return <el-tag effect="dark" type="info">等待中</el-tag>
              case 2:
                return <el-tag effect="dark" type="warning">排隊中</el-tag>
              case 3:
                return <el-tag effect="dark" type="success">服務中</el-tag>
              case 4:
                return <el-tag effect="dark" type="danger">已關閉</el-tag>
            }
          }
        },
        {
          prop: 'created_at',
          label: '開始時間',
          width: 180,
          render: (h, scope) => {
            return moment(scope.row.created_at).format("YYYY-MM-DD HH:mm:ss")
          }
        },
        {
          prop: 'closed_at',
          label: '結束時間',
          width: 180,
          render: (h, scope) => {
            if (scope.row.closed_at.Valid) {
              return moment(scope.row.closed_at.Time).format("YYYY-MM-DD HH:mm:ss")
            } else {
              return ""
            }
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
        const { data, pagination } = await apiGetRoomList(query.toString())
        this.tableData = data
        this.pagination = pagination
        this.loading = false
      } catch (err) {
        console.log(err)
        this.loading = false
      }
    },
    async fetchStaffList() {
      try {
        this.loading = true
        const params = new URLSearchParams({
          name: "",
          status: 0,
          serving_status: 0,
          page: 1,
          page_size: 100,
        });
        const { data } = await apiGetStaffList(params.toString())
        this.staffs = data
      } catch (err) {
        console.log(err)
      } finally {
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
