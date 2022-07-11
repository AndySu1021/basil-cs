<template>
  <div class="app-table-container">
    <el-table
        :data="tableData"
        style="width: 100%"
        header-cell-class-name="app-table-header"
        empty-text="暫無數據"
    >
      <AppTableColumn v-for="(column,idx) in tableColumns" :key="idx" :attrs="column"/>
      <el-table-column
          v-if="showOperation"
          fixed="right"
          label="操作"
          width="160"
          header-align="center"
          align="center"
      >
        <template slot-scope="scope">
          <el-button
              size="small"
              @click.native.prevent="handleEdit(scope.row.id)"
          >
            編輯
          </el-button>
          <el-button
              type="danger"
              size="small"
              @click="handleDelete(scope.row.id)"
          >
            刪除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import AppTableColumn from "@/components/AppTable/components/AppTableColumn";

export default {
  name: "Table",
  components: {
    AppTableColumn,
  },
  props: {
    tableData: {
      type: [Array,Object],
      default: [],
    },
    tableColumns: {
      type: [Array,Object],
      required: true,
    },
    showOperation: {
      type: Boolean,
      default: false,
    },
    handleEdit: {
      type: Function,
      default: (id) => console.log("edit id: "+id),
    },
    handleDelete: {
      type: Function,
      default: (id) => console.log("delete id: "+id),
    }
  },
}
</script>
<style lang="scss">
.app-table-header {
  background: #F5F5F5 !important;
  color: #666;
  font-weight: bold;
  font-size: 15px;
}
</style>
<style lang="scss" scoped>
.app-table-container {
  width: 100%;
}
</style>