<template>
  <div class="pagination-wrap">
    <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :page-sizes="[10, 20, 30, 40]"
        :layout="layout"
        :current-page="paginationData.page"
        :page-size="paginationData.page_size"
        :total="paginationData.total"
        :hide-on-single-page="false"
    >
    </el-pagination>
  </div>
</template>

<script>
export default {
  name: "AppPagination",
  data() {
    return {
      paginationData: {
        page: 1,
        page_size: 10,
        total: 0,
      },
    };
  },
  props: {
    data: {
      type: Object,
      required: true,
    },
    layout: {
      type: String,
      default: "total, sizes, prev, pager, next, jumper",
    }
  },
  watch: {
    data: {
      handler(newValue) {
        this.paginationData = { ...newValue };
      },
      immediate: true,
      deep: true,
    },
  },
  methods: {
    handleCurrentChange(page) {
      const payload = { ...this.paginationData, page };
      this.$emit('change', payload);
    },
    handleSizeChange(pageSize) {
      const payload = { ...this.paginationData, pageSize, page: 1 };
      this.$emit('change', payload);
    },
  },
}
</script>

<style lang="scss" scoped>
.pagination-wrap {
  margin-top: 20px;
  text-align: center;
}
::v-deep .el-pager li.active {
  color: #409EFF;
  cursor: default;
}
</style>