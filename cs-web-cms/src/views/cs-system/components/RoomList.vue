<template>
  <div class="room-list-container" v-loading="loading">
    <el-tabs v-model="activeName" @tab-click="handleClick" lazy stretch>
      <el-tab-pane label="當前" name="Serving" class="tab-pane-item">
        <div v-if="roomList.length > 0">
          <div v-for="(room,idx) in roomList" :key="idx">
            <RoomItem
                :id="room.id"
                :name="room.member_name"
                :op-type="room.op_type"
                :content-type="room.content_type"
                :content="room.content"
                :isRead="room.isRead"
                :unread-count="1"
                :member-id="room.member_id"
                :online-status="room.online_status"
            />
          </div>
        </div>
        <div v-else class="empty-room">
          <svg-icon icon-class="people"></svg-icon>
          <span>- 尚無訪客 -</span>
        </div>
      </el-tab-pane>
      <el-tab-pane label="排隊中" name="Queuing" class="tab-pane-item">
        <div v-if="roomList.length > 0">
          <div v-for="(room,idx) in roomList" :key="idx">
            <RoomItem
                :id="room.id"
                :name="room.member_name"
                message=""
                :isRead="true"
                :unread-count="1"
                :member-id="room.member_id"
                :online-status="room.online_status"
            />
          </div>
        </div>
        <div v-else class="empty-room">
          <svg-icon icon-class="people"></svg-icon>
          <span>- 尚無訪客 -</span>
        </div>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import RoomItem from "@/views/cs-system/components/RoomItem";
import {mapGetters} from "vuex";
import AppPagination from "@/components/AppPagination";

export default {
  name: "RoomList",
  components: {RoomItem, AppPagination},
  computed: {
    ...mapGetters({
      roomList: "cs/roomList",
      activeTab: "cs/activeTab",
    }),
  },
  created() {
    this.fetchStaffRoomList()
  },
  mounted() {
    this.activeName = this.activeTab
  },
  watch: {
    activeTab() {
        this.activeName = this.activeTab
    }
  },
  data() {
    return {
      loading: false,
      activeName: "",
    }
  },
  methods: {
    handleClick() {
      this.$store.commit("cs/RESET")
      this.$store.commit("cs/SET_ACTIVE_TAB", this.activeName)
      if (this.activeName === "Serving") {
        this.fetchStaffRoomList()
      } else {
        this.fetchPendingRoomList()
      }
    },
    async fetchStaffRoomList() {
      await this.$store.dispatch("cs/getStaffRoomList")
    },
    async fetchPendingRoomList() {
      await this.$store.dispatch("cs/getQueuingRoomList")
    },
  }
}
</script>

<style lang="scss" scoped>
.room-list-container {
  border-radius: 6px;
  height: 90vh;
  background-color: white;
  .tab-pane-item {
    padding: 0 14px;
    height: 82vh;
    overflow-y: scroll;
  }
  .empty-text {
    span {
      display: block;
      color: #c3c3c3;
      font-weight: bold;
      text-align: center;
      padding: 10px;
      margin-bottom: 10px;
    }
  }
}
.empty-room {
  display: flex;
  align-items: center;
  flex-direction: column;
  width: 100%;
  color: #e3e3e3;
  font-weight: bold;
  font-size: 22px;
  .svg-icon {
    width: 75px;
    height: 80px;
    margin: 30px 20px;
  }
}
</style>