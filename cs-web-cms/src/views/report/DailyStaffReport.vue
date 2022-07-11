<template>
  <div class="app-container" v-loading="loading">
    <AppPanel title="每日客服報表">
      <template v-slot:filter>
        <el-date-picker
            style="margin-right: 10px;"
            v-model="dateRange"
            type="daterange"
            range-separator="至"
            start-placeholder="開始日期"
            end-placeholder="結束日期"
            value-format="yyyy-MM-dd"
        >
        </el-date-picker>
        <el-button type="primary" icon="el-icon-search" @click="fetchData">搜尋</el-button>
      </template>
      <template v-slot:body>
        <AppTable
            :table-data="tableData"
            :table-columns="tableColumns"
        />
      </template>
    </AppPanel>
  </div>
</template>

<script>
import AppTable from '@/components/AppTable'
import moment from "moment";
import {apiGetDailyStaffReport} from "@/api/report";
import AppPanel from "@/components/AppPanel";

export default {
  name: 'DailyStaffReport',
  components: {
    AppTable,
    AppPanel,
  },
  created() {
    this.fetchData()
  },
  data() {
    let now = moment().format("yyyy-MM-DD")
    let before = moment().subtract(9, 'days').format("yyyy-MM-DD")
    return {
      loading: false,
      dateRange: [before, now],
      tableData: [],
      tableColumns: [
        {
          prop: 'date',
          label: '日期',
          render: (h, scope) => {
            return moment(scope.row.date).format("yyyy-MM-DD")
          }
        },
        { prop: 'staff_name', label: '客服名稱' },
        { prop: 'total_member', label: '總接待人數' },
        { prop: 'total_scored_member', label: '總評分人數' },
        { prop: 'average_score', label: '平均分數' },
      ]
    }
  },
  methods: {
    async fetchData() {
      try {
        this.loading = true
        const params = new URLSearchParams({
          start_date: this.dateRange[0],
          end_date: this.dateRange[1],
        });
        const { data } = await apiGetDailyStaffReport(params.toString())
        this.tableData = data
        this.loading = false
      } catch (err) {
        console.log(err)
        this.loading = false
      }
    },
  },
}
</script>
