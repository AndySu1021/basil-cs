<template>
  <div>
    <ChatBox />
  </div>
</template>

<script>
import ChatBox from "@/components/ChatBox";
import { connectSocket } from "./utils/ws";
import {onMounted} from "vue";
import {setSID} from "@/utils/storage";
import {useStore} from 'vuex'

export default {
  name: 'App',
  components: {
    ChatBox,
  },
  setup() {
    const store = useStore()
    onMounted(() => {
      const params = new Proxy(new URLSearchParams(window.location.search), {
        get: (searchParams, prop) => searchParams.get(prop),
      });
      setSID(params.sid)
      store.commit('ws/SET_ROOM_ID', params.room_id)
      store.commit('ws/SET_MEMBER_NAME', params.name)
      connectSocket();
    })

    return {};
  },
}
</script>

<style>
* {
  box-sizing: border-box;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin-top: 30px;
  margin-bottom: 30px;
}
</style>
