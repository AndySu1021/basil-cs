<template>
  <div class="app-container">
    <el-row :gutter="14">
      <el-col :span="6">
        <RoomList />
      </el-col>
      <el-col :span="10">
        <ChatBox />
      </el-col>
      <el-col :span="8">
        <ToolBoard />
      </el-col>
    </el-row>
  </div>
</template>

<script>
import RoomList from "@/views/cs-system/components/RoomList";
import ChatBox from "@/views/cs-system/components/ChatBox";
import ToolBoard from "@/views/cs-system/components/ToolBoard";
import {mapGetters} from "vuex";
export default {
  name: "index",
  components: {ToolBoard, ChatBox, RoomList}
}
</script>

<style lang="scss" scoped>

</style>