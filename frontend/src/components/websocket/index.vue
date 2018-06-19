<template>
  <div>
  </div>
</template>

<script>
import config from '@/config/index'
import EventBus from '@/events/event-bus.js'

export default {
  data() {
    return {
      token: '',
      uri: this.$route.path,
      wenzi_message_websocket: undefined,
      dialogShow: false,
      loadingInstance: null
    }
  },

  created() {
    this.connect()
    this.registerEvent()
  },

  destroyed() {
    this.destroyEvent()
  },

  methods: {
    sendWebSocketData(data) {
      if (!this.wenzi_message_websocket || this.wenzi_message_websocket.readyState === 3) {
        return
      }
      // websocket 请求加上token 作为唯一标识
      data.token = this.token
      data = JSON.stringify(data)
      this.wenzi_message_websocket.send(data)
    },
    connect() {
      // 参考链接：swoole 快速起步 swoole 部分： https://wiki.swoole.com/wiki/page/479.html
      if (window.wenzi_message_websocket) {
        this.wenzi_message_websocket = window.wenzi_message_websocket
        return
      }

      console.log('开始连接')
      var wenzi_message_websocket = new WebSocket(config.wenzi_message_websocket_uri)

      wenzi_message_websocket.onopen = (evt) => {
        var data = {
          action: 'open'
        }

        this.sendWebSocketData(data)
      }

      wenzi_message_websocket.onmessage = (evt) => {
        var res = JSON.parse(evt.data)
        if (res.action === 'open') {
          // 没有token ，设置token
          if (this.token === '') {
            this.token = res.data.token
            this.$store.dispatch('modifyToken', res.data.token)
          }
          // 重发open，保存会还
          var data = {
            action: 'open'
          }

          this.sendWebSocketData(data)
        }
        if (res.action === 'close') {
          alert(res.data.message)
          this.closeWebPage()
        }

        if (res.action === 'replyMessage') {
          EventBus.$emit('add-message', { token: res.data.token, message: res.data.message })
        }
      }

      wenzi_message_websocket.onclose = (evt) => {
        console.log(evt)
        console.log('断开连接')
        window.wenzi_message_websocket = null
        this.wenzi_message_websocket = null
        console.log('reconnection')
        this.connect()
      }

      wenzi_message_websocket.onerror = function(evt, e) {
        console.log('Error occured: ' + evt.data)
      }

      window.wenzi_message_websocket = wenzi_message_websocket
      this.wenzi_message_websocket = window.wenzi_message_websocket
    },

    registerEvent() {
      EventBus.$on('send-message', (data) => {
        this.sendWebSocketData(data)
      })
    },

    destroyEvent() {
      EventBus.$off('send-message')
    }
  }
}
</script>

<style scoped>
</style>

