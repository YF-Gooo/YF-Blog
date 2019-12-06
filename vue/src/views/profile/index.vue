<template>
  <div class="app-container" style="margin: 0 10%;">
    <div v-if="user">
      <el-row :gutter="20">

        <el-col :span="6" :xs="24">
          <user-card :user="user" />
        </el-col>

        <el-col :span="18" :xs="24">
          <el-card>
            <el-tabs v-model="activeTab">
              <el-tab-pane label="Resume" name="resume">
                <resume />
              </el-tab-pane>
              <el-tab-pane label="Timeline" name="timeline">
                <timeline />
              </el-tab-pane>
            </el-tabs>
          </el-card>
        </el-col>

      </el-row>
    </div>
  </div>
</template>

<script>

import UserCard from './components/UserCard'
import Timeline from './components/Timeline'
import Resume from './components/Resume'
export default {
  name: 'Profile',
  components: { UserCard, Timeline,Resume},
  data() {
    return {
      user: {},
      activeTab: 'resume',
    }
  },
  created() {
    this.getUser()
  },
  methods: {
    getUser() {
      this.user = {
        name: "YF",
        role: "DataScientist",
        email: 'admin@test.com',
        avatar: ""
      }
    }
  }
}
</script>
