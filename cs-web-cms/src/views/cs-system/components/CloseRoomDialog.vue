<template>
  <AppDialog ref="dialog" title="確定要結束嗎?" width="500px">
    <div slot="body">
      <el-form
          ref="dialogForm"
          :model="formData"
          :rules="rules"
          label-width="80px"
      >
        <el-form-item label="標籤:" prop="tagId">
          <el-select v-model="formData.tagId" placeholder="請選擇">
            <el-option
                v-for="(tag,idx) in tags"
                :key="idx"
                :label="tag.name"
                :value="tag.id">
            </el-option>
          </el-select>
        </el-form-item>
      </el-form>
    </div>
    <div slot="footer">
      <el-button @click="handleCancel">取消</el-button>
      <el-button type="primary" @click="submitForm">確認</el-button>
    </div>
  </AppDialog>
</template>

<script>
import AppDialog from '@/components/AppDialog'
import { deepCopy } from '@/utils'
import {apiCloseRoom} from "@/api/room";
import {apiGetAvailableTags} from "@/api/tag";
import {mapGetters} from "vuex";

export default {
  name: 'CloseRoomDialog',
  components: {
    AppDialog
  },
  computed: {
    ...mapGetters({
      activeRoomId: 'cs/activeRoomId',
      activeTab: 'cs/activeTab',
    }),
  },
  data() {
    return {
      tagID: 0,
      tags: [],
      initialFormData: {
        tagId: null
      },
      formData: {
        tagId: null
      },
      rules: {
        tagId: [{ required: true, message: '必填', trigger: 'blur' }],
      },
    }
  },
  methods: {
    show() {
      this.fetchTags()
      this.formData = deepCopy(this.initialFormData)
      this.$refs.dialog.show()
    },
    handleCancel() {
      this.$refs.dialog.hide()
    },
    submitForm() {
      this.$refs.dialogForm.validate(async (valid) => {
        if (valid) {
          this.$refs.dialog.toggleLoadingFullScreen()
          try {
            await apiCloseRoom(this.activeRoomId, {
              tag_id: this.formData.tagId,
            })
            this.$store.commit("cs/RESET")
            this.$store.commit("cs/RESET_ROOM_INFO")
            await this.fetchRoomList()
            this.$showSuccessMessage("操作成功", this.afterSubmit)
          } catch (error) {
            this.$refs.dialog.toggleLoadingFullScreen()
          }
        }
      })
    },
    afterSubmit() {
      this.$refs.dialog.afterSubmit()
    },
    async fetchTags() {
      try {
        this.loading = true
        const { data } = await apiGetAvailableTags()
        this.tags = data
      } catch (err) {
        console.log(err)
      } finally {
        this.loading = false
      }
    },
    async fetchRoomList() {
      const params = new URLSearchParams({
        status: this.activeTab === "Serving" ? 2 : 1,
        page: 1,
        page_size: 10
      });
      await this.$store.dispatch("cs/getStaffRoomList", params.toString())
    }
  }
}
</script>
