<template>
  <div>
    <el-upload
      ref="upload"
      action=""
      :limit="limit"
      :http-request="httpRequest"
      :file-list="fileList"
      list-type="picture"
      :before-upload="beforeUpload"
      accept="image/jpeg, image/png"
    >
      <el-button v-if="previewImage === ''" size="small">
        點擊上傳
      </el-button>
    </el-upload>

    <div v-if="previewImage" class="preview-wrap">
      <el-image class="preview-wrap__image" fit="cover" :src="previewImage" />
      <i
        v-if="previewImage"
        class="preview-wrap__icon el-icon-delete"
        @click="handleRemove"
      />
    </div>
  </div>
</template>

<script>
export default {
  name: 'ImageUpload',
  props: {
    onRemove: {
      type: Function,
      required: true
    },
    limit: {
      type: Number,
      default: 1
    },
    httpRequest: {
      type: Function,
      required: true
    },
    beforeUpload: {
      type: Function
    },
    previewImage: {
      type: String
    }
  },
  data() {
    return {
      fileList: []
    }
  },
  methods: {
    handleRemove() {
      this.$refs.upload.clearFiles()
      this.onRemove()
    }
  }
}
</script>

<style lang="scss" scoped>
::v-deep .el-upload-list {
  display: none;
}
.preview-wrap {
  position: relative;
  padding-top: 5px;
  width: 300px;
  margin-top: 10px;

  &:hover {
    .preview-wrap__icon {
      opacity: 1;
    }

    .preview-wrap__image:after {
      opacity: 1;
    }
  }

  &__image {
    width: 100%;
    transition: 0.3s all ease-in-out;

    &:after {
      content: '';
      display: block;
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      background: rgba(0, 0, 0, 0.6);
      transition: 0.3s all ease-in-out;
      opacity: 0;
    }
  }

  &__icon {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    cursor: pointer;
    font-size: 24px;
    color: #fff;
    opacity: 0;
    transition: 0.3s all ease-in-out;
  }
}
</style>
