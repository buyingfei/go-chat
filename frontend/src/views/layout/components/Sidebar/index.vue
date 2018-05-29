<template>
  <scroll-bar>
    <el-row>
      <el-col :span="24">
        <div class="gonggao" style="">
          系统消息：虎牙依法对直播内容进行24小时巡查，禁止传播违法违规、封建迷信、
          暴力血腥、低俗色情、招嫖诈骗、违禁品等不良信息，坚决维护青少年群体精神文明健康。
          请勿轻信各类招聘征婚、代练代抽、刷钻、购买礼包码、游戏币、电商贩卖等广告信息，以免上当受骗。
        </div>
      </el-col>
      <message-log></message-log>
      <el-col :span="24">
        <el-form :inline="true"  size="small" :model="formInline" class="demo-form-inline">
          <el-form-item>
            <el-input v-model="mess" placeholder="请文明发言"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit">发送</el-button>
          </el-form-item>
        </el-form>
      </el-col>
      <websocket></websocket>
    </el-row>
  </scroll-bar>
</template>

<script>
import { mapGetters } from 'vuex'
import SidebarItem from './SidebarItem'
import MessageLog from './messageLog'
import ScrollBar from '@/components/ScrollBar'
import Websocket from '@/components/websocket'
import EventBus from '@/events/event-bus.js'

export default {
  data() {
    return {
      formInline: {},
      mess: ''
    }
  },
  components: { SidebarItem, ScrollBar, Websocket, MessageLog },
  methods: {
    onSubmit() {
      EventBus.$emit('send-message', { action: 'sendMessage', 'token': this.$store.getters.token, message: this.mess })
      EventBus.$emit('add-message', { token: this.$store.getters.token, message: this.mess })
      this.mess = ''
    }
  },
  computed: {
    ...mapGetters([
      'sidebar'
    ]),
    routes() {
      return this.$router.options.routes
    },
    isCollapse() {
      return !this.sidebar.opened
    }
  }
}
</script>
<style rel="stylesheet/scss" lang="scss" scoped>
  .gonggao {
    background-color: rgba(103,194,58,.1);
    display: inline-block;
    padding: 0 10px;
    line-height: 30px;
    font-size: 12px;
    color: #409eff;
    border-radius: 4px;
    box-sizing: border-box;
    border: 1px solid rgba(64,158,255,.2);
    border-color: rgba(103,194,58,.2);
    color: #67c23a;
  }
</style>
