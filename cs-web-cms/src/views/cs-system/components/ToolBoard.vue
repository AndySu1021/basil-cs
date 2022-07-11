<template>
<div class="tool-board-container">
  <el-tabs v-model="activeTab" lazy stretch>
    <el-tab-pane class="tab-panel-container" label="快捷回覆" name="fm">
      <el-input
          placeholder="請輸入標題"
          prefix-icon="el-icon-search"
          style="margin-bottom: 20px;"
          v-model="titleSearch"
          @change="handleSearch"
      >
      </el-input>
      <el-collapse class="group-collapse" v-if="tmpGroup.length > 0" v-model="activeItems">
        <el-collapse-item v-for="(groupItem,idx) in tmpGroup" :key="idx" :title="groupItem.category.name" :name="groupItem.category.id">
          <el-table
              :data="groupItem.items"
              border
              style="width: 100%;border-radius: 8px;"
              :cell-style="{padding:'7px'}"
              :header-cell-style="{padding:'5px'}"
          >
            <el-table-column
              fixed
              prop="title"
              label="標題"
              width="100"
            >
            </el-table-column>
            <el-table-column
              fixed
              prop="content"
              label="內容"
            >
            </el-table-column>
            <el-table-column
              label="操作"
              width="90"
              align="center"
            >
              <template slot-scope="scope">
                <el-button type="primary" size="small" @click="handleSend(scope.row.content)">發送</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-collapse-item>
      </el-collapse>
    </el-tab-pane>
    <el-tab-pane class="tab-panel-container" label="詳細資訊" name="ci">
      <el-descriptions title="訪問資訊" border :column="1">
        <el-descriptions-item label="來源站點">{{roomInfo.site_name}}</el-descriptions-item>
        <el-descriptions-item label="裝置來源">
          <el-tag v-if="roomInfo.source === 1" size="small">WEB</el-tag>
          <el-tag v-else-if="roomInfo.source === 2" size="small">APP</el-tag>
          <el-tag v-else size="small">未知</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="訪問 IP">{{roomInfo.ip}}</el-descriptions-item>
        <el-descriptions-item label="首次訪問時間">{{getTime(roomInfo.created_at)}}</el-descriptions-item>
        <el-descriptions-item label="瀏覽器">{{roomInfo.user_agent}}</el-descriptions-item>
      </el-descriptions>
      <div style="margin-top: 30px;">
        <span style="display: block;margin-bottom: 20px;font-weight: bold;font-size: 16px;color: #303133;">會員資料</span>
        <el-form label-width="80px" label-position="left">
          <el-form-item label="姓名">
            <el-input v-model="roomInfo.real_name" size="small"></el-input>
          </el-form-item>
          <el-form-item label="電子信箱">
            <el-input v-model="roomInfo.email" size="small"></el-input>
          </el-form-item>
          <el-form-item label="手機號">
            <el-input v-model="roomInfo.mobile" size="small"></el-input>
          </el-form-item>
          <el-form-item label="備註">
            <el-input
                style="width: 300px"
                v-model="roomInfo.note"
                placeholder="請輸入內容"
                :rows="5"
                type="textarea"
                resize="none"
                maxlength="200"
            />
          </el-form-item>
          <el-button type="primary" size="small" @click="submitForm">儲存</el-button>
        </el-form>
      </div>
    </el-tab-pane>
  </el-tabs>
</div>
</template>

<script>
import {apiGetFastReplyGroup} from "@/api/fast-reply";
import {sendSocketMessage} from "@/utils/ws";
import {mapGetters} from "vuex";
import {deepCopy} from "@/utils";
import {apiUpdateMemberInfo} from "@/api/member";
import moment from "moment";
import {encrypt} from "@/utils/crypto";

export default {
  name: "ToolBoard",
  created() {
    this.fetchFastReplyGroup()
  },
  computed: {
    ...mapGetters({
      activeRoomId: 'cs/activeRoomId',
      roomInfo: 'cs/roomInfo',
    }),
  },
  data() {
    return {
      activeTab: "fm",
      activeItems: ['1'],
      group: [],
      tmpGroup: [],
      titleSearch: '',
    }
  },
  methods: {
    getTime(time) {
      if (time === "") {
        return ""
      }
      return moment(time).format('YYYY-MM-DD HH:mm:ss')
    },
    async fetchFastReplyGroup() {
      try {
        this.loading = true
        const { data } = await apiGetFastReplyGroup()
        this.group = data
        this.tmpGroup = data
      } catch (err) {
        console.log(err)
      } finally {
        this.loading = false
      }
    },
    async submitForm() {
      try {
        await apiUpdateMemberInfo(this.roomInfo.member_id, {
            real_name: this.roomInfo.real_name,
            email: this.roomInfo.email,
            mobile: this.roomInfo.mobile,
            note: this.roomInfo.note,
        })
        this.$showSuccessMessage("操作成功", this.afterSubmit)
      } catch (error) {
        console.log(error)
      }
    },
    handleSend(message) {
      message = message.trim()
      message = message.replace(/\r\n|\n/g,"");
      if(message !== "" && this.activeRoomId !== "") {
        sendSocketMessage({
          op_type: 2,
          room_id: this.activeRoomId,
          content_type: 1,
          content: encrypt(message),
        })
      }
    },
    handleSearch(title) {
      this.tmpGroup = deepCopy(this.group)
      for(let i = 0; i < this.tmpGroup.length; i++){
        this.tmpGroup[i].items = this.tmpGroup[i].items.filter(function (item) {
          return item.title.includes(title)
        })
      }
      for(let i = 0; i < this.tmpGroup.length; i++){
        this.tmpGroup[i].items = this.tmpGroup[i].items.filter(function (item) {
          return item.title.includes(title)
        })
      }
      this.tmpGroup = this.tmpGroup.filter(function (group) {
        return group.items.length > 0
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.tool-board-container {
  width: 100%;
  height: 90vh;
  border-radius: 6px;
  background-color: white;
  .tab-panel-container {
    padding: 10px 20px;
    overflow-y: scroll;
    height: 82vh;
    .group-collapse {
      ::v-deep .el-collapse-item__header {
        padding-left: 10px;
        font-size: 14px;
        color: #333;
        font-weight: bold;
      }
    }
  }
}
::v-deep .el-tabs {
  ::v-deep .el-tabs__header {
    ::v-deep .el-tabs__item {
      padding: 0 !important;
    }
  }
}
</style>