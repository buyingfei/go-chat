<template>
  <div class="menu-wrapper">
    <template v-for="item in routes" v-if="!item.hidden&&item.children">
      <el-submenu  :index="item.name||item.path" :key="item.name">
        <template slot="title">
          <svg-icon :el-type="true" v-if="item.meta&&item.meta.icon" :icon-class="item.meta.icon"></svg-icon>
          <span v-if="item.meta&&item.meta.title">{{item.meta.title}}</span>
        </template>

        <template v-for="child in item.children" v-if="!child.hidden">
          <sidebar-item :is-nest="true" class="nest-menu" v-if="child.children&&child.children.length>0" :routes="[child]" :key="child.path"></sidebar-item>

          <template v-else>
            
            <!--判断是否是/ 目录-->
            <router-link v-if="item.path=='/'" :to="child.path" :key="child.name">
              <el-menu-item :index="child.path">
                <svg-icon  v-if="child.meta&&child.meta.icon" :icon-class="child.meta.icon"></svg-icon>
                <span v-if="child.meta&&child.meta.title">{{child.meta.title}}</span>
              </el-menu-item>
            </router-link>
  
            <router-link v-else :to="item.path+'/'+child.path" :key="child.name">
              <el-menu-item :index="item.path+'/'+child.path">
                <svg-icon  v-if="child.meta&&child.meta.icon" :icon-class="child.meta.icon"></svg-icon>
                <span v-if="child.meta&&child.meta.title">{{child.meta.title}}</span>
              </el-menu-item>
            </router-link>
            
          </template>
          
        </template>
      </el-submenu>

    </template>
  </div>
</template>

<script>
export default {
  name: 'SidebarItem',
  props: {
    routes: {
      type: Array
    },
    isNest: {
      type: Boolean,
      default: false
    }
  },
  watch: {
  },
  computed: {
  }
}
</script>
