<template>
  <div class="navbar">
    <div class="hamburger-wrap">
      <hamburger :is-active="sidebar.opened" @toggleClick="toggleSideBar" />
    </div>
    <div class="breadcrumb-wrap">
      <breadcrumb />
    </div>
    <div class="info-wrap">
      <div>
        <el-tag
            :type="websocketStatus ? 'success' : 'danger'"
            effect="dark"
        >
          WS 連線{{websocketStatus ? '成功' : '失敗'}}
        </el-tag>
      </div>
      <div class="logout">
        <el-dropdown trigger="click" @command="handleCommand" placement="bottom">
          <div class="avatar">
            <img v-if="avatar !== ''" :src="require('@/assets/'+avatar)"/>
            <el-avatar v-else class="el-dropdown-link" icon="el-icon-user-solid"></el-avatar>
            <span class="dot1"></span>
            <span class="dot2" :class="command === 'Closed' ? 'red' : 'green'"></span>
          </div>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item command="Serving">
              <span class="green-point">開啟</span>
            </el-dropdown-item>
            <el-dropdown-item command="Closed">
              <span class="red-point">關閉</span>
            </el-dropdown-item>
            <el-dropdown-item icon="el-icon-picture-outline-round" command="Avatar" divided>設置大頭貼</el-dropdown-item>
            <el-dropdown-item icon="el-icon-switch-button" command="Logout" divided>登出</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </div>
    </div>
    <EditAvatarDialog ref="EditAvatarDialog" />
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Breadcrumb from '@/components/Breadcrumb'
import Hamburger from '@/components/Hamburger'
import LangSelect from '@/components/LangSelect'
import {apiUpdateStaffServingStatus} from "@/api/staff";
import EditAvatarDialog from "@/layout/components/EditAvatarDialog";

export default {
  components: {
    Breadcrumb,
    Hamburger,
    LangSelect,
    EditAvatarDialog,
  },
  computed: {
    ...mapGetters({
      sidebar: 'app/sidebar',
      websocketStatus: 'app/websocketStatus',
      avatar: 'user/avatar',
      servingStatus: 'user/servingStatus',
    }),
  },
  data() {
    return {
      command: ""
    }
  },
  mounted() {
    this.command = this.servingStatus
  },
  methods: {
    toggleSideBar() {
      this.$store.dispatch('app/toggleSideBar')
    },
    handleCommand(command) {
      if(command === "Logout") {
        this.logout()
      } else if(command === "Avatar") {
        this.$refs.EditAvatarDialog.show()
      } else {
        this.command = command
        switch (command) {
          case 'Closed':
            this.updateServingStatus(1)
            break
          case 'Serving':
            this.updateServingStatus(2)
            break
        }
      }
    },
    async updateServingStatus(status) {
      try {
        this.loading = true
        await apiUpdateStaffServingStatus({
          serving_status: status
        })
      } catch (err) {
        console.log(err)
      } finally {
        this.loading = false
      }
    },
    async logout() {
      try {
        await this.$confirmLogout()
        await this.$store.dispatch('user/logout')
        await this.$router.push('/login')
      } catch (err) {
        console.log(err)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.navbar {
  display: flex;
  height: 50px;
  overflow: hidden;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);

  .hamburger-wrap {
    width: 50px;
    padding-left: 5px;
  }
  .breadcrumb-wrap {
    flex-grow: 1;
  }

  .info-wrap {
    display: flex;
    align-items: center;
    //line-height: 50px;
    height: 50px;

    .info-item {
      padding: 0 15px;
    }
    .logout {
      padding: 0 10px;
      cursor: pointer;
      transition: all 0.1s ease-in-out;
      &:hover {
        color: #409eff;
      }
    }
  }
}
.el-dropdown-link {
  margin-top: 3px;
  &:hover {
    cursor: pointer;
  }
}
.el-dropdown-menu__item {
  .green-point::before{
    content: " ";
    display: inline-block;
    border: 4px solid #31A24C;/*设置红色*/
    border-radius:50%;
    z-index: 1000;
    right: 0;
    height: 2px;
    margin-right: 13px;
    margin-top: -5px;
  }
  .red-point::before{
    content: " ";
    display: inline-block;
    border: 4px solid #F56C6C;
    border-radius:50%;
    z-index: 1000;
    right: 0;
    height: 2px;
    margin-right: 13px;
    margin-top: -5px;
  }
}
.el-dropdown {
  .avatar {
    position: relative;
    img {
      width: 40px;
      height: 40px;
      border-radius: 50%;
    }
    .dot1 {
      display: inline-block;
      border-radius: 50%;
      position: absolute;
      top: 26px;
      right: -2px;
      border: 7px solid white;
    }
    .dot2 {
      display: inline-block;
      border-radius: 50%;
      position: absolute;
      top: 28px;
      right: 0px;
      &.red {
        border: 5px solid #F56C6C;
      }
      &.green {
        border: 5px solid #31A24C;
      }
    }
  }
}
</style>
