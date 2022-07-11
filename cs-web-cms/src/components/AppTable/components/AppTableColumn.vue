<script>
export default {
  name: 'AppTableColumn',
  props: {
    attrs: {
      type: Object,
      default: () => ({}),
      required: true,
    },
  },
  render(h) {
    const label = typeof this.attrs.label === 'function' ? this.attrs.label() : this.attrs.label;
    const attrs = { ...this.attrs, minWidth: 150, showOverflowTooltip: true, label };
    const scopedSlots = {};
    if (attrs.render) {
      scopedSlots.default = scope => attrs.render(h, scope);
    }
    return h(
        'el-table-column',
        {
          attrs,
          scopedSlots,
        },
        attrs.isGroupHeader ? [attrs.render(h)] : null
    );
  },
};
</script>
