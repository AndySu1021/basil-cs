<template>
  <div class="app-container" v-loading="loading">
    <AppPanel title="訊息列表">
      <template v-slot:filter>
        <el-input v-model="form.room_id" placeholder="房間編號" style="width: 100px;" />
        <el-select style="width: 150px;" v-model="form.staff_id" placeholder="職員">
          <el-option label="全部" :value="0"></el-option>
          <el-option v-for="(staff,idx) in staffs" :key="idx" :label="staff.name" :value="staff.id"></el-option>
        </el-select>
        <el-input v-model="form.content" placeholder="訊息內容" style="margin-right: 10px;width: 250px;" />
        <el-button icon="el-icon-search" type="primary" @click="fetchData(pagination)">搜尋</el-button>
        <el-button icon="el-icon-circle-close" @click="handleClear">清除</el-button>
      </template>
      <template v-slot:body>
        <AppTable
            :table-data="tableData"
            :table-columns="tableColumns"
        />
        <AppPagination
            :data="pagination"
            @change="handleChange"
        />
      </template>
    </AppPanel>
  </div>
</template>

<script>
import {apiGetStaffList} from "@/api/staff";
import {deepCopy} from "@/utils";
import {apiGetMessageList} from "@/api/message";
import moment from "moment";
import AppTable from "@/components/AppTable";
import AppPagination from "@/components/AppPagination";
import AppPanel from "@/components/AppPanel";

export default {
  name: 'MessageList',
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
        content: '',
      },
      form: {
        room_id: '',
        staff_id: 0,
        content: '',
      },
      pagination: {
        page: 1,
        page_size: 10,
        total: 0,
      },
      tableData: [],
      tableColumns: [
        { prop: 'room_id', label: '房間編號', align: "center", width: 100 },
        {
          prop: 'sender_type',
          label: '發送方式',
          align: "center",
          width: 140,
          render: (h, param) => {
            switch (param.row.sender_type) {
              case 0:
                return <span>系統發送</span>
              case 1:
                return <span>會員 -> 客服</span>
              case 2:
                return <span>客服 -> 會員</span>
            }
          }
        },
        { prop: 'sender_name', label: '發送人', align: "center", width: 140 },
        {
          prop: 'content_type',
          label: '消息類型',
          align: "center",
          width: 120,
          render: (h, scope) => {
            switch (scope.row.content_type) {
              case 1:
                return <span>文字</span>
              case 2:
                return <span>圖片</span>
              default:
                return <span>通知</span>
            }
          }
        },
        { prop: 'content', label: '消息內容' },
        {
          prop: 'ts',
          label: '發送時間',
          align: "center",
          width: 190,
          render: (h, scope) => {
            return <span>{this.getTime(scope.row.ts)}</span>
          }
        },
      ]
    }
  },
  methods: {
    getTime(timestamp) {
      return moment(timestamp*1000).format('YYYY-MM-DD HH:mm:ss');
    },
    async fetchData(payload) {
      try {
        this.loading = true
        let params = deepCopy(this.form)
        if(params.room_id === '') {
          params.room_id = 0
        }
        const query = new URLSearchParams({...params, ...(payload)});
        const { data, pagination } = await apiGetMessageList(query.toString())
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
