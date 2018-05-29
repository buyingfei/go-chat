<template>
  <i v-if="elType" :class="iconName"></i>
  <svg v-else-if="!elType" :class="svgClass" aria-hidden="true">
    <use :xlink:href="iconName"></use>
  </svg>
</template>

<script>
export default {
  name: 'svg-icon',
  props: {
    iconClass: {
      type: String,
      required: true
    },
    elType: {
      type: Boolean,
      require: false
    },
    className: {
      type: String
    }
  },
  computed: {
    iconName() {
      if (this.elType) {
        return `el-icon-${this.iconClass}`
      } else {
        return `#icon-${this.iconClass}`
      }
    },
    svgClass() {
      if (this.className) {
        return 'svg-icon ' + this.className
      } else {
        return 'svg-icon'
      }
    }
  }
}
</script>

<style scoped>
.svg-icon {
  width: 1em;
  height: 1em;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}
</style>
