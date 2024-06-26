import { gql } from "@apollo/client/core";

export const GET_USERS = gql`
  query Users($payload: UserSearchPayload!) {
    getUsers(input:$payload) {
      total
      items {
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
    }
  }
`

export const LOGIN = gql`
  mutation UserLogin($payload:UserLoginPayload!) {
    userLogin(input: $payload) {
      profile {
        id,
        email,
        firstName,
        lastName,
        image,
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

export const PROFILE = gql`
  query UserProfile {
    userProfile {
      id,
      firstName,
      lastName,
      email,
      image,
      audit {
        createdAt
        createdBy
        updatedAt
        updatedBy
      }
    }
  }
`

export const UPDATE_PROFILE = gql`
  mutation UserUpdate($payload: UserUpdatePayload!) {
    userUpdateProfile(input:$payload) {
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
  }
`