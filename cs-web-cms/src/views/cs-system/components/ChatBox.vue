<template>
  <div class="chat-box-container">
    <div v-if="activeRoomId === '' || activeTab === 'Queuing'" class="chat-box-mask"></div>
    <el-button v-if="activeTab === 'Queuing'" class="accept-btn" type="primary" @click="handleAccept" :disabled="activeRoomId === ''">開始諮詢</el-button>
    <ChatBoxHeader />
    <div id="chat-box-body" class="chat-box-body">
      <div v-for="(msg,idx) in messageList" :key="idx">
        <div style="margin-bottom: 16px;">
          <SystemMessage v-if="msg.sender_type === 0" :op-type="msg.op_type" :content="msg.content" />
          <MemberMessage
              v-else-if="msg.sender_type === 1"
              :name="msg.sender_name"
              :content-type="msg.content_type"
              :content="msg.content"
              :timestamp="msg.ts"
          />
          <StaffMessage
              v-else
              :name="msg.sender_name"
              :content-type="msg.content_type"
              :content="msg.content"
              :timestamp="msg.ts"
          />
        </div>
      </div>
    </div>
    <ChatBoxFooter />
  </div>
</template>

<script>
import ChatBoxHeader from "./ChatBoxHeader";
import ChatBoxFooter from "./ChatBoxFooter";
import StaffMessage from "./StaffMessage";
import MemberMessage from "./MemberMessage";
import SystemMessage from "./SystemMessage";
import {mapGetters} from "vuex";
import {apiGetStaffList} from "@/api/staff";
import {apiAcceptRoom, apiUpdateRoom} from "@/api/room";

export default {
  name: "ChatBox",
  components: {SystemMessage, MemberMessage, ChatBoxHeader, ChatBoxFooter, StaffMessage},
  computed: {
    ...mapGetters({
      messageList: 'cs/messageList',
      activeTab: 'cs/activeTab',
      activeRoomId: 'cs/activeRoomId',
    }),
  },
  mounted() {
    this.scrollToBottom()
  },
  updated:function(){
    this.scrollToBottom();
  },
  methods:{
    isNull(value) {
      return value == null || typeof (value) == 'undefined';
    },
    scrollToBottom () {
      this.$nextTick(() => {
        let elem = document.getElementById('chat-box-body');
        if(!this.isNull(elem)){
          elem.scrollTop = elem.scrollHeight;
        }
      })
    },
    async handleAccept() {
      try {
        this.loading = true
        if(this.activeRoomId === "") {
          return
        }
        await apiAcceptRoom(this.activeRoomId)
        await this.$store.commit('cs/SET_ACTIVE_TAB', "Serving")
        await this.fetchRoomList()
        const params = new URLSearchParams({
          room_id: this.activeRoomId,
          page: 1,
          page_size: 10,
        });
        await this.$store.dispatch("cs/getRoomMessageList", params.toString())
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
  },
  data() {
    return {}
  }
}
</script>

<style lang="scss" scoped>
.chat-box-container {
  position: relative;
  width: 100%;
  background-color: #f4f5f6;
  border-radius: 6px;
  height: 90vh;
  .chat-box-mask {
    width: 100%;
    height: 100%;
    border-radius: 6px;
    position: absolute;
    top: 0;
    left: 0;
    background-color: #ffffff;
    z-index: 7;
    border-radius: 6px;
    opacity: 0.7;
  }
  ::v-deep .accept-btn {
    position: absolute;
    z-index: 8;
    bottom: 20px;
    margin: 0;
    left: 50%;
    -ms-transform: translateX(-50%);
    transform: translateX(-50%);
  }
  .chat-box-body {
    height: 60vh;
    margin: 12px 0 20px 0;
    padding: 0 15px;
    overflow-y: scroll;
    .message-content {
      display: block;
      margin: 20px;
      color: Black;
      font-weight: bold;
    }
  }
}
</style>