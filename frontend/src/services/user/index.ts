import type { UserLoginMutation, UserLoginPayload, UserProfileQuery, UserSearchPayload, UsersQuery, UserUpdateMutation, UserUpdatePayload } from "@/graphql.generated"
import { ApolloClient, type FetchResult, type NormalizedCacheObject } from "@apollo/client/core"
import { GET_USERS, LOGIN, PROFILE, UPDATE_PROFILE } from "./operation.gql"

export function userGql(client: ApolloClient<NormalizedCacheObject>) {
  return {
    async list(payload: UserSearchPayload): Promise<FetchResult<UsersQuery> | undefined> {
      return client?.query<UsersQuery>({
        query: GET_USERS,
        variables: {
          payload
        },
      })
    },
    async login(payload: UserLoginPayload): Promise<FetchResult<UserLoginMutation> | undefined> {
      return client?.query<UserLoginMutation>({
        query: LOGIN,
        variables: {
          payload
        },
      })
    },
    async profile(): Promise<FetchResult<UserProfileQuery> | undefined> {
      return client?.query<UserProfileQuery>({
        query: PROFILE,
      })
    },
    async update(payload: UserUpdatePayload): Promise<FetchResult<UserUpdateMutation> | undefined> {
      return client?.mutate<UserUpdateMutation>({
        mutation: UPDATE_PROFILE,
        variables: {
          payload
        }
      })
    }
  }
}
