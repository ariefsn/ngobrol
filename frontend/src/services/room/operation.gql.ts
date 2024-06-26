import { gql } from "@apollo/client/core";

export const SUB_ROOM_DETAILS = gql`
  subscription subRoomDetails($code: String!) {
    subRoomDetails (code: $code){
      id,
      code,
      image,
      users {
        id
        firstName
        lastName
        email
        image
        audit {
          createdAt
          createdBy
          updatedAt
          updatedBy
        }
      }
      audit {
        createdAt
        createdBy
        updatedAt
        updatedBy
      }
    }
  }
`
