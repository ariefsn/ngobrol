import type { MessageData } from '@/graphql.generated'
import gql from '@/services'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useMessageStore = defineStore('message', () => {
  const newMessage = ref<MessageData | null>()
  const socketNewMessage = ref()
  const messages = ref<MessageData[]>([])

  async function getMessages() {
    const res = await gql.message.getMessage({ limit: 1000, skip: 0 })

    const items = res?.data?.getMessages?.items ?? []
    messages.value = items.reverse()
  }

  async function sendMessage(message: string) {
    const res = await gql.message.sendMessage({ message })
  }

  function subsNewMessage(code: string) {
    socketNewMessage.value = gql.message.subNewMessage(code).subscribe((res) => {
      const data = res?.data
      newMessage.value = data?.subNewMessage
      if (newMessage.value) {
        const isNew = messages.value.filter((item) => item.id === data?.subNewMessage?.id).length === 0
        if (isNew) {
          messages.value = [...messages.value, data!.subNewMessage]
        }
      }
    })
  }

  function unsubNewMessage() {
    socketNewMessage.value?.unsubscribe()
  }

  return { subsNewMessage, unsubNewMessage, newMessage, sendMessage, getMessages, messages }
})
