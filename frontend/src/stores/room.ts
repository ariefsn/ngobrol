import type { RoomDataDetails } from '@/graphql.generated'
import gql from '@/services'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useRoomStore = defineStore('room', () => {
  const details = ref<RoomDataDetails | null>()
  const socketDetails = ref()

  function subsDetails(code: string) {
    socketDetails.value = gql.room.subRoomDetails(code).subscribe((res) => {
      const data = res?.data
      details.value = data?.subRoomDetails
    })
  }

  function unsubDetails() {
    socketDetails.value?.unsubscribe()
  }

  const users = () => details.value?.users ?? []

  return { subsDetails, unsubDetails, details, users }
})
