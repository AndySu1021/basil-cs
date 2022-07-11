<template>
  <div class="app-container" v-loading="loading">
    <AppPanel title="每日標籤報表">
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
import {apiGetDailyTagReport} from "@/api/report";
import moment from "moment";
import AppPanel from "@/components/AppPanel";

export default {
  name: 'DailyTagReport',
  components: {
    AppPanel,
    AppTable,
  },
  created() {
    this.fetchData()
  },
  data() {
    let now = moment().format("yyyy-MM-DD")
    let before = moment().subtract(9, 'days').format("yyyy-MM-DD")
    return {
      dateRange: [before, now],
      loading: false,
      tableData: [],
      tableColumns: [],
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
        const { data } = await apiGetDailyTagReport(params.toString())
        let columns = data.columns
        let items = data.items
        this.tableColumns = [{ prop: 'date', label: '日期' }]
        for(let i = 0; i < columns.length; i++){
          this.tableColumns.push({ prop: columns[i].key, label: columns[i].label })
        }
        this.tableData = []
        let keys = Object.keys(items)
        for(let i = 0; i < keys.length; i++){
          this.tableData.push({date: keys[i], ...items[keys[i]]})
        }
        this.loading = false
      } catch (err) {
        console.log(err)
        this.loading = false
      }
    },
  },
}
</script>
