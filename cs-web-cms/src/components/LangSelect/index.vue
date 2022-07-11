<template>
  <el-dropdown trigger="click" class="international" @command="handleSetLanguage">
    <div>
      <svg-icon name="language" icon-class="language" />
      {{ languageTitle }}
    </div>
    <el-dropdown-menu slot="dropdown">
      <el-dropdown-item :disabled="language === 'zh'" command="zh">
        中文
      </el-dropdown-item>
      <el-dropdown-item :disabled="language === 'en'" command="en">
        English
      </el-dropdown-item>
    </el-dropdown-menu>
  </el-dropdown>
</template>

<script>
export default {
  name: 'LangSelect',
  computed: {
    language() {
      return this.$store.state.language;
    },
    languageTitle() {
      let languages = {
        zh: '中文',
        en: 'English',
      };
      let curLanguage = this.$store.state.language;
      return languages[curLanguage];
    },
  },
  methods: {
    handleSetLanguage(lang) {
      this.$i18n.locale = lang;
      this.$store.dispatch('app/SetLanguage', lang);
    },
  },
};
</script>

<style lang="scss" scoped>
.international {
  width: 50px;
  text-align: center;
  font-size: 20px;
  cursor: pointer;
  transition: all 0.1s ease-in-out;
  &:hover {
    color: #409eff;
  }
}
</style>
