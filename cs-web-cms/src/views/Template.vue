<template>
  <div></div>
</template>

<script>
export default {
  name: '',
  props: {
    propName: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      key: ''
    }
  },
  watch: {
    data(newValue, oldValue) {}
  },
  computed: {
    name() {
      return this.data
    }
  },
  methods: {
    name() {}
  },
  created() {},
  mounted() {},
  destroyed() {}
}
</script>

<style lang="scss" scoped></style>
