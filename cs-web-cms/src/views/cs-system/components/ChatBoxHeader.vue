<template>
  <div>
    <div class="header-container">
      <div class="header-text">諮詢房 #{{activeRoomId}}</div>
      <div>
        <el-tooltip content="轉接" placement="top">
          <el-button type="success" icon="el-icon-position" circle @click="openTransferRoomDialog"></el-button>
        </el-tooltip>
        <el-tooltip content="評分" placement="top">
          <el-button type="warning" icon="el-icon-star-off" circle @click="sendScore"></el-button>
        </el-tooltip>
        <el-tooltip content="結束" placement="top">
          <el-button type="danger" icon="el-icon-switch-button" circle @click="openCloseRoomDialog"></el-button>
        </el-tooltip>
      </div>
    </div>
    <CloseRoomDialog ref="CloseRoomDialog" />
    <TransferRoomDialog ref="TransferRoomDialog" />
  </div>
</template>

<script>
import CloseRoomDialog from "@/views/cs-system/components/CloseRoomDialog";
import {mapGetters} from "vuex";
import {sendSocketMessage} from "@/utils/ws";
import TransferRoomDialog from "@/views/cs-system/components/TransferRoomDialog";

export default {
  name: "ChatBoxHeader",
  components: {CloseRoomDialog, TransferRoomDialog},
  computed: {
    ...mapGetters({
      activeRoomId: 'cs/activeRoomId',
    }),
  },
  data() {
    return {
      loading: false,
      roomId: "",
    }
  },
  methods: {
    openCloseRoomDialog() {
      this.$refs.CloseRoomDialog.show()
    },
    openTransferRoomDialog() {
      this.$refs.TransferRoomDialog.show()
    },
    async sendScore() {
      await this.$confirmSendScore()
      sendSocketMessage({
        op_type: 3,
        room_id: this.activeRoomId,
        content_type: 1,
        content: ""
      })
    },
  }
}
</script>

<style lang="scss" scoped>
.header-container {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  background-color: #fff;
  border-radius: 6px 6px 0 0;
  font-weight: bold;
  color: white;
  padding: 10px;
  .header-text {
    color: #777;
  }
}
</style>