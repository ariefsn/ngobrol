<script setup lang="ts">
import router from '@/router'
import gql from '@/services'
import { ref } from 'vue'

const model = ref({
  email: '',
  roomCode: ''
})

const onSubmit = async () => {
  try {
    const res = await gql.user.login(model.value)
    const profile = res?.data?.userLogin?.profile
    if (profile) {
      localStorage.setItem('auth-email', profile.email)
      localStorage.setItem('auth-room', model.value.roomCode)
      router.replace('/')
    } else {
      localStorage.clear()
      router.replace('/login')
    }
  } catch (error) {
    console.log(error)
  }
}
</script>

<template>
  <v-container
    fluid
  >
    <v-row
      align="center"
      no-gutters
      class="h-screen"
    >
      <v-col
        cols="12"
        lg="4"
        offset="4"
      >
        <v-sheet class="pa-2 ma-2">
          <v-text-field label="Email" v-model="model.email" />
          <v-text-field label="Room Code" v-model="model.roomCode" />
          <v-btn color="primary" @click="onSubmit">Join</v-btn>
        </v-sheet>
      </v-col>
    </v-row>
  </v-container>
</template>
