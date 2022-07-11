<template>
  <div class="app-container">
    <panel-group />
    <el-row :gutter="30">
      <el-col :span="10">
        <div class="staff-panel">
          <h3>在線客服</h3>
          <AppTable :table-data="tableData" :table-columns="tableColumns" />
        </div>
      </el-col>
      <el-col :span="14">
        <div class="staff-panel">
          <h3>提醒事項</h3>
          <AppTable :table-data="remindData" :table-columns="remindColumns" />
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import PanelGroup from './components/PanelGroup'
import {apiGetAllStaff} from "@/api/staff";
import AppTable from "@/components/AppTable";
import {apiGetActiveReminds} from "@/api/remind";

export default {
  name: 'Dashboard',
  components: {
    PanelGroup,
    AppTable,
  },
  data() {
    return {
      tableData: [],
      tableColumns: [
        { prop: 'id', label: '編號', align: "center", width: 80 },
        { prop: 'name', label: '名稱', align: "center" },
        {
          prop: 'serving_status',
          label: '服務狀態',
          align: "center",
          width: 100,
          render: (h, param) => {
            switch (param.row.serving_status) {
              case 1:
                return <el-tag effect="dark" type="danger">關閉</el-tag>
              case 2:
                return <el-tag effect="dark" type="success">服務中</el-tag>
              case 3:
                return <el-tag effect="dark" type="warning">閒置</el-tag>
            }
          }
        },
      ],
      remindData: [],
      remindColumns: [
        { prop: 'title', label: '標題', width: 200 },
        { prop: 'content', label: '內容' },
      ]
    }
  },
  created() {
    this.fetchData()
    this.fetchRemind()
  },
  methods: {
    async fetchData() {
      try {
        this.loading = true
        const { data } = await apiGetAllStaff()
        this.tableData = data
        this.loading = false
      } catch (err) {
        console.log(err)
        this.loading = false
      }
    },
    async fetchRemind() {
      try {
        this.loading = true
        const { data } = await apiGetActiveReminds()
        this.remindData = data
        this.loading = false
      } catch (err) {
        console.log(err)
        this.loading = false
      }
    },
  }
}
</script>

<style lang="scss" scoped>
.app-container {
  .staff-panel {
    border-radius: 6px;
    padding: 6px 20px 20px 20px;
    background-color: white;
  }
}
</style>
