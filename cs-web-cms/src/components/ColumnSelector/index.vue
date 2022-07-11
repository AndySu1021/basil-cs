<template>
  <el-popover>
    <div class="reset-button" @click="reset">{{ $t('reset') }}</div>
    <div v-for="(column, index) in updatedColumns" :key="index">
      <el-checkbox v-model="column.visible" :label="column.label" @change.native="change">
        {{ column.label }}
      </el-checkbox>
      <br />
    </div>
    <el-button slot="reference" type="primary" size="small">{{ $t('columnSelector') }}</el-button>
  </el-popover>
</template>

<script>
import { deepCopy } from '@/utils';

export default {
  name: 'ColumnSelector',
  props: {
    columns: {
      type: Array,
      default: () => [],
    },
    columnsKey: {
      type: String,
      default: '',
    },
  },
  data() {
    return {
      collectionKey: this.columnsKey ? `${this.$route.path}/${this.columnsKey}` : this.$route.path,
      defaultColumns: [],
      updatedColumns: [],
      unselectableColumns: [],
      columnSelectorCollection: {},
    };
  },
  created() {
    this.unselectableColumns = this.columns.filter(column => !column.prop);
    let selectableColumns = [];
    this.columnSelectorCollection = JSON.parse(localStorage.getItem('columnSelectorCollection')) || {};
    if (this.columnSelectorCollection[this.collectionKey]) {
      selectableColumns = this.columnSelectorCollection[this.collectionKey];
    } else {
      selectableColumns = this.columns.reduce((acc, cur) => {
        if (cur.prop) {
          acc.push({ ...cur });
        }
        return acc;
      }, []);
    }
    this.defaultColumns = deepCopy(selectableColumns);
    this.updatedColumns = deepCopy(selectableColumns);
    this.change();
  },
  methods: {
    reset() {
      this.updatedColumns = this.columns.reduce((acc, cur) => {
        if (cur.prop) {
          acc.push({ ...cur, visible: true });
        }
        return acc;
      }, []);
      this.change();
    },
    change() {
      const columns = [...this.unselectableColumns, ...this.updatedColumns];
      this.$emit('column-change', columns);
      this.columnSelectorCollection[this.collectionKey] = this.updatedColumns;
      localStorage.setItem('columnSelectorCollection', JSON.stringify(this.columnSelectorCollection));
    },
  },
};
</script>

<style lang="scss" scoped>
.reset-button {
  text-align: right;
  cursor: pointer;
}
.el-button {
  margin: 0 10px;
}
</style>
