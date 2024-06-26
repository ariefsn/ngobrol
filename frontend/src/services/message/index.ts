import type { GetMessagesQuery, MessageCreatePayload, MessageSearchPayload, MessageSendMutation, SubNewMessageSubscription } from "@/graphql.generated"
import { ApolloClient, Observable, type FetchResult, type NormalizedCacheObject, type OperationVariables } from "@apollo/client/core"
import { GET_MESSAGES, SEND_MESSAGE, SUB_NEW_MESSAGE } from "./operation.gql"

export function messageGql(client: ApolloClient<NormalizedCacheObject>) {
  return {
    subNewMessage(code: string): Observable<FetchResult<SubNewMessageSubscription>> {
      return client?.subscribe<SubNewMessageSubscription, OperationVariables>({
        query: SUB_NEW_MESSAGE,
        variables: {
          code
        },
      })
    },
    sendMessage(payload: MessageCreatePayload): Promise<FetchResult<MessageSendMutation> | undefined> {
      return client?.mutate<MessageSendMutation>({
        mutation: SEND_MESSAGE,
        variables: {
          payload
        }
      })
    },
    getMessage(payload: MessageSearchPayload): Promise<FetchResult<GetMessagesQuery> | undefined> {
      return client?.query<GetMessagesQuery>({
        query: GET_MESSAGES,
        variables: {
          payload
        }
      })
    }
  }
}
