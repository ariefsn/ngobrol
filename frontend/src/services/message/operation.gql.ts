import { gql } from "@apollo/client/core";

export const GET_MESSAGES = gql`
  query GetMessages($payload: MessageSearchPayload!) {
    getMessages(input: $payload) {
      total
      items {
        id
        roomId
        fromId
        message
        isNew
        audit {
          createdAt
          createdBy
          updatedAt
          updatedBy
        }
      }
    }
  }
`


export const SEND_MESSAGE = gql`
  mutation MessageSend($payload: MessageCreatePayload!) {
    sendMessage(input: $payload) {
      id
      roomId
      fromId
      message
      isNew
      audit {
        createdAt
        createdBy
        updatedAt
        updatedBy
      }
    }
  }
`

export const SUB_NEW_MESSAGE = gql`
  subscription SubNewMessage($code: String!) {
    subNewMessage(code: $code) {
      id
      roomId
      fromId
      message
      isNew
      audit {
        createdAt
        createdBy
        updatedAt
        updatedBy
      }
    }
  }
`
