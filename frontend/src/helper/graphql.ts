import type { IMap } from '@/entities';
import { ApolloClient, HttpLink, InMemoryCache, from, split, type NormalizedCacheObject } from '@apollo/client/core';
import { loadDevMessages, loadErrorMessages } from '@apollo/client/dev';
import { setContext } from '@apollo/client/link/context';
import { ErrorLink } from '@apollo/client/link/error';
import { GraphQLWsLink } from '@apollo/client/link/subscriptions';
import { getMainDefinition } from '@apollo/client/utilities';
import { createClient } from 'graphql-ws';

let _client: ApolloClient<NormalizedCacheObject>

const onError = new ErrorLink((err) => {
  let msg = err?.networkError?.message
  if ((err?.graphQLErrors ?? []).length > 0) {
    msg = err.graphQLErrors![0].message
  }

  return err.forward(err.operation)
})

const injectTokens = setContext(async (_, { skip, headers }) => {
  const newHeaders = { ...headers }

  const authEmail = localStorage.getItem('auth-email')
  if (authEmail) {
    newHeaders['X-Email'] = authEmail
  }

  const authRoom = localStorage.getItem('auth-room')
  if (authRoom) {
    newHeaders['X-Room'] = authRoom
  }

  return {
    headers: newHeaders
  }
})

export const createGqlClient = ({
  uri,
  headers
}: {
  uri: string,
  headers?: IMap,
}) => {
  if (!_client && uri) {
    const httpLink = new HttpLink({ uri })
    const wsLink = new GraphQLWsLink(createClient({
      url: import.meta.env.VITE_GQL_WS_URL,
    }))

    const splitLink = split(
      ({ query }) => {
        const definition = getMainDefinition(query);
        return (
          definition.kind === 'OperationDefinition' &&
          definition.operation === 'subscription'
        );
      },
      wsLink,
      httpLink,
    );

    _client = new ApolloClient({
      cache: new InMemoryCache({}),
      headers,
      link: from([
        injectTokens,
        onError,
        splitLink,
      ]),
      defaultOptions: {
        watchQuery: {
          fetchPolicy: 'no-cache',
          errorPolicy: 'ignore',
        },
        query: {
          fetchPolicy: 'no-cache',
          errorPolicy: 'all',
        },
      }
    })
  }

  loadDevMessages()
  loadErrorMessages()

  return _client
}

export const apolloClient = () => createGqlClient({
  uri: import.meta.env.VITE_GQL_URL
})
