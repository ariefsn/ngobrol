<script setup lang="ts">
import type { MessageData, UserData } from '@/graphql.generated'
import { useMessageStore } from '@/stores/message'
import { useRoomStore } from '@/stores/room'
import { useUserStore } from '@/stores/user'
import { onMounted, ref } from 'vue'

const userStore = useUserStore()
const roomStore = useRoomStore()
const messageStore = useMessageStore()

const user = ref<UserData>({
  email: '',
  firstName: '',
  lastName: ''
} as UserData)

const newMessage = ref('')

onMounted(async () => {
  await Promise.all([
    getProfile(),
    messageStore.getMessages(),
  ])
  const roomCode = localStorage.getItem('auth-room') ?? ''
  roomStore.subsDetails(roomCode)
  messageStore.subsNewMessage(roomCode)
})

const getProfile = async () => {
  await userStore.getProfile()
  if (userStore?.profile) {
    user.value = {...userStore!.profile}
  }
}

const onUpdate = async () => {
  await userStore.updateProfile({
    firstName: user.value.firstName,
    lastName: user.value.lastName,
    image: user.value.image
  })
  if (userStore?.profile) {
    user.value = {...userStore!.profile}
  }
}

const onLogout = async () => {
  localStorage.clear()
  location.href = '/login'
}

const onSendMessage = async () => {
  await messageStore.sendMessage(newMessage.value)
  newMessage.value = ''
}

const isMyMessage = (message: MessageData): boolean => {
  return localStorage.getItem('auth-email') === message.fromId
}

const userName = (message: MessageData): { image: string, name: string } => {
  if (roomStore.users().length === 0) {
    return {
      image: '',
      name: ''
    }
  }
  const user = roomStore.users().find((user: UserData) => user.email === message.fromId)
  let fullName = user?.email

  if (user?.firstName) {
    const names = [user?.firstName]

    if (user?.lastName) {
      names.push(user?.lastName)
    }

    fullName = names.join(' ')
  }

  return {
    image: user?.image ?? '',
    name: fullName ?? ''
  }
}

const userInitials = (name: string): string => {
  return name
    .split(' ')
    .map((n) => n[0])
    .join('')
    .toUpperCase()
}

</script>

<template>
  <main>
    <v-container fluid>
      <v-row align="center" class="h-screen">
        <v-col
          cols="4"
        >
          <v-sheet elevation="3" rounded>
            <v-sheet class="pa-2 ma-2">
              <v-text-field label="Email" v-model="user.email" readonly />
              <v-text-field label="First Name" v-model="user.firstName" />
              <v-text-field label="Last Name" v-model="user.lastName" />
              <v-btn color="primary" @click="onUpdate">Update</v-btn>
              <v-btn color="secondary" class="ml-3" @click="onLogout">Logout</v-btn>
            </v-sheet>
          </v-sheet>
        </v-col>
        <v-col
          cols="8"
        >
          <v-sheet elevation="3" rounded class="pa-3 relative" style="min-height: 80vh;">
            <div style="max-height: 70vh; overflow-y: auto;">
              <div v-for="message in messageStore.messages" :key="message.id">
                <div v-if="isMyMessage(message)" class="flex w-full justify-end">
                  <v-sheet style="width: fit-content;" elevation="3" class="mb-2 mr-1 mt-1 p-2">
                    <span>
                      {{ message.message }}
                    </span>
                    <v-avatar color="surface-variant" class="ml-2">
                      <span>{{ userInitials(userName(message).name) }}</span>
                    </v-avatar>
                  </v-sheet>
                </div>
                <div v-else>
                  <v-sheet style="width: fit-content;" elevation="3" class="mb-2 ml-1 mt-1 p-2">
                    <v-avatar color="surface-variant" class="mr-2">
                      <span>{{ userInitials(userName(message).name) }}</span>
                    </v-avatar>
                    <span>
                      {{ message.message }}
                    </span>
                  </v-sheet>
                </div>
              </div>
            </div>
            <div>
              <v-text-field label="Message" v-model="newMessage" class="absolute w-11/12 bottom-0">
                <template #append>
                  <div>
                    <v-btn color="primary" @click="onSendMessage">Send</v-btn>
                  </div>
                </template>
              </v-text-field>
            </div>
          </v-sheet>
        </v-col>
      </v-row>
    </v-container>
  </main>
</template>
