import type { SubRoomDetailsSubscription } from "@/graphql.generated"
import { ApolloClient, Observable, type FetchResult, type NormalizedCacheObject, type OperationVariables } from "@apollo/client/core"
import { SUB_ROOM_DETAILS } from "./operation.gql"

export function roomGql(client: ApolloClient<NormalizedCacheObject>) {
  return {
    subRoomDetails(code: string): Observable<FetchResult<SubRoomDetailsSubscription>> {
      return client?.subscribe<SubRoomDetailsSubscription, OperationVariables>({
        query: SUB_ROOM_DETAILS,
        variables: {
          code
        },
      })
    },
  }
}
