<template>
<div class="room-item-container" :class="{selected: this.id === activeRoomId}" @click="handleClick">
  <div v-if="!isRead" class="unread-count">{{unreadCount}}</div>
  <div class="top-container">
    <div class="avatar">
      <img :src="require('@/assets/user.jpg')"/>
      <span class="dot1"></span>
      <span class="dot2" :class="status === 2 ? 'grey' : 'green'"></span>
    </div>
    <div class="item-text">
      <div style="color: #333;font-weight: bold;"><span style="color: #666;font-weight: bold;margin-right: 8px;">#{{id}}</span>{{name}}</div>
      <div v-if="opType === 2" class="message-text" :class="isRead ? '' : 'unread'">
        <span v-if="contentType === 1">{{content | decrypt }}</span>
        <span v-else>會員上傳一張圖片</span>
      </div>
      <div v-else-if="opType === 3" class="message-text" :class="isRead ? '' : 'unread'">
        <span>已發送評分請求</span>
      </div>
      <div v-else-if="opType === 4" class="message-text" :class="isRead ? '' : 'unread'">
        <span>已完成評分</span>
      </div>
    </div>
  </div>
</div>
</template>

<script>
import {apiGetRoomInfo} from "@/api/room";
import {mapGetters} from "vuex";

export default {
  name: "RoomItem",
  computed: {
    ...mapGetters({
      activeRoomId: 'cs/activeRoomId',
      roomInfo: 'cs/roomInfo',
    }),
  },
  props: {
    id: Number,
    name: String,
    opType: Number,
    contentType: Number,
    content: String,
    isRead: Boolean,
    unreadCount: Number,
    memberId: Number,
    onlineStatus: Number,
  },
  data() {
    return {
      status: 2,
    }
  },
  created() {
    this.status = this.onlineStatus
  },
  methods: {
    handleClick() {
      this.$store.commit("cs/SET_MESSAGE_READ", this.id)
      this.$store.commit("cs/SET_ACTIVE_ROOM", this.id)
      const params = new URLSearchParams({
        room_id: this.id,
        page: 1,
        page_size: 10,
      });
      this.fetchRoomInfo()
      this.$store.dispatch("cs/getRoomMessageList", params.toString())
    },
    async fetchRoomInfo() {
      try {
        this.loading = true
        const { data } = await apiGetRoomInfo(this.id)
        this.$store.commit("cs/SET_ROOM_INFO", data)
        this.status = data.online_status
      } catch (err) {
        console.log(err)
      } finally {
        this.loading = false
      }
    },
  }
}
</script>

<style lang="scss" scoped>
.room-item-container {
  width: 100%;
  padding: 16px 16px 13px 16px;
  margin: auto auto 10px;
  border-radius: 6px;
  position: relative;
  .unread-count {
    position: absolute;
    right: 16px;
    top: 50%;
    transform: translateY(-50%);
    background-color: #F56C6C;
    padding: 4px 7px;
    color: white;
    font-weight: bold;
    border-radius: 20px;
    font-size: 12px;
  }
  .top-container {
    display: flex;
    align-items: center;
    position: relative;
    .avatar {
      position: relative;
      img {
        width: 50px;
        height: 50px;
        border-radius: 50%;
      }
      .dot1 {
        display: inline-block;
        border-radius: 50%;
        position: absolute;
        top: 33px;
        right: -2px;
        border: 7px solid white;
      }
      .dot2 {
        display: inline-block;
        border-radius: 50%;
        position: absolute;
        top: 35px;
        right: 0;
        &.grey {
          border: 5px solid #b4b4b4;
        }
        &.green {
          border: 5px solid #31A24C;
        }
      }
    }
    .item-text {
      margin-left: 10px;
      .message-text {
        margin-top: 6px;
        width: 100%;
        span {
          font-size: 15px;
          color: #777;
          overflow: hidden;
          white-space:nowrap;
          text-overflow:ellipsis;
          display:inline-block;
          width: 100%;
        }
        &.unread {
          span {
            font-weight: bold;
            color: #333;
          }
        }
      }
    }
  }
  &:hover {
    background-color: rgba(0, 0, 0, 0.03);
    cursor: pointer;
  }
  &.selected {
    background-color: rgba(0, 0, 0, 0.05);
  }
}
</style>