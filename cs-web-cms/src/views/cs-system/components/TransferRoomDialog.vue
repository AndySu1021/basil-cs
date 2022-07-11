<template>
  <AppDialog ref="dialog" title="確定要轉接嗎?" width="500px">
    <div slot="body">
      <el-form
          ref="dialogForm"
          :model="formData"
          :rules="rules"
          label-width="100px"
      >
        <el-form-item label="轉接客服:" prop="staffId">
          <el-select v-model="formData.staffId" placeholder="請選擇">
            <el-option
                v-for="(staff,idx) in staffs"
                :key="idx"
                :label="staff.name"
                :value="staff.id">
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
import {apiTransferRoom} from "@/api/room";
import {mapGetters} from "vuex";
import {apiGetAvailableStaff} from "@/api/staff";

export default {
  name: 'TransferRoomDialog',
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
      staffs: [],
      initialFormData: {
        staffId: null
      },
      formData: {
        staffId: null
      },
      rules: {
        staffId: [{ required: true, message: '必填', trigger: 'blur' }],
      },
    }
  },
  methods: {
    show() {
      this.fetchStaffs()
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
            await apiTransferRoom(this.activeRoomId, {
              staff_id: this.formData.staffId,
            })
            this.$store.commit("cs/RESET")
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
    async fetchStaffs() {
      try {
        this.loading = true
        const { data } = await apiGetAvailableStaff()
        this.staffs = data
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
