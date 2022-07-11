<template>
<div class="score-dialog-container">
  <el-dialog
      v-model="showScoreDialog"
      ref:="dialog"
      title="客服評分"
      width="400px"
      top="200px"
      :show-close="false"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      :center="true"
      :destroy-on-close="true"
      open-delay="300"
  >
    <div class="rating">
      <el-rate
          v-model="score"
          :colors="scoreColors"
          size="large"
      />
    </div>
    <template #footer>
      <span class="dialog-footer">
        <el-button color="#138184" size="large" @click="handleSend">送出</el-button>
      </span>
    </template>
  </el-dialog>
</div>
</template>

<script>
import {ElButton, ElDialog, ElRate} from 'element-plus'
import {mapGetters} from "vuex";
import {apiUpdateRoomScore} from "@/apis/room";
import {sendSocketMessage} from "@/utils/ws";
export default {
  name: "ScoreDialog",
  components: {
    ElDialog,
    ElButton,
    ElRate,
  },
  computed: {
    ...mapGetters({
      showScoreDialog: 'ws/showScoreDialog',
    })
  },
  data() {
    return {
      dialogVisible: false,
      score: 0,
      scoreColors: ['#99A9BF', '#F7BA2A', '#FF9900'],
    }
  },
  methods: {
    async handleSend() {
      if (this.score > 0) {
        await apiUpdateRoomScore({
          score: this.score
        })
        this.$store.commit('ws/SET_SHOW_SCORE_DIALOG', false)
        sendSocketMessage({
          op_type: 4,
          content_type: 1,
          content: ""
        })
      }
    },
  }
}
</script>
<style lang="scss" scoped>
.score-dialog-container{
  .rating {
    text-align: center;
    .el-rate {
      --el-rate-icon-size: 30px;
      --el-rate-icon-margin: 8px;
    }
  }
  :deep(.el-dialog) {
    border-radius: 13px;
    .el-dialog__header {
      background-color: #138184;
      margin-right: 0;
      border-radius: 13px 13px 0 0;
      padding: 12px;
      .el-dialog__title {
        color: #f4f5f6;
        font-weight: bolder;
      }
    }
    .el-dialog__body {
      padding: 30px;
    }
    .el-dialog__footer {
      padding-top: 0;
    }
  }
}
</style>