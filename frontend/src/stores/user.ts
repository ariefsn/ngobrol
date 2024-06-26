import type { UserData, UserUpdatePayload } from '@/graphql.generated'
import gql from '@/services'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserStore = defineStore('user', () => {
  const profile = ref<UserData | null>()

  async function getProfile() {
    await gql.user.profile().then((res) => {
      const data = res?.data
      profile.value = data?.userProfile
    })
  }

  async function updateProfile(newData: UserUpdatePayload) {
    await gql.user.update(newData).then((res) => {
      const data = res?.data
      profile.value = data?.userUpdateProfile
    })
  }

  return { profile, getProfile, updateProfile }
})
