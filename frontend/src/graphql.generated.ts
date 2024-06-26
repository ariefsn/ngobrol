export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Map: { input: any; output: any; }
  Time: { input: any; output: any; }
};

export type Audit = {
  __typename?: 'Audit';
  createdAt: Scalars['Time']['output'];
  createdBy?: Maybe<Scalars['String']['output']>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['String']['output']>;
};

export type MessageCreatePayload = {
  message: Scalars['String']['input'];
};

export type MessageData = {
  __typename?: 'MessageData';
  audit: Audit;
  fromId: Scalars['String']['output'];
  id: Scalars['String']['output'];
  isNew: Scalars['Boolean']['output'];
  message: Scalars['String']['output'];
  roomId: Scalars['String']['output'];
};

export type MessageSearchPayload = {
  limit?: InputMaybe<Scalars['Int']['input']>;
  skip?: InputMaybe<Scalars['Int']['input']>;
};

export type MessageSearchResponse = {
  __typename?: 'MessageSearchResponse';
  items: Array<MessageData>;
  total?: Maybe<Scalars['Int']['output']>;
};

export type Mutation = {
  __typename?: 'Mutation';
  roomCreate: RoomDataDetails;
  sendMessage: MessageData;
  userLogin: UserLoginResponse;
  userLogout: Scalars['Boolean']['output'];
  userUpdateProfile: UserData;
};


export type MutationRoomCreateArgs = {
  input: RoomCreatePayload;
};


export type MutationSendMessageArgs = {
  input: MessageCreatePayload;
};


export type MutationUserLoginArgs = {
  input: UserLoginPayload;
};


export type MutationUserUpdateProfileArgs = {
  input: UserUpdatePayload;
};

export type Query = {
  __typename?: 'Query';
  getMessages: MessageSearchResponse;
  getUsers: UserSearchResponse;
  roomDetails: RoomDataDetails;
  userProfile: UserData;
};


export type QueryGetMessagesArgs = {
  input: MessageSearchPayload;
};


export type QueryGetUsersArgs = {
  input: UserSearchPayload;
};

export type RoomCreatePayload = {
  roomCode: Scalars['String']['input'];
  userId: Scalars['String']['input'];
};

export type RoomData = {
  __typename?: 'RoomData';
  audit: Audit;
  code: Scalars['String']['output'];
  id: Scalars['String']['output'];
  image: Scalars['String']['output'];
  users: Array<Scalars['String']['output']>;
};

export type RoomDataDetails = {
  __typename?: 'RoomDataDetails';
  audit: Audit;
  code: Scalars['String']['output'];
  id: Scalars['String']['output'];
  image: Scalars['String']['output'];
  users: Array<UserData>;
};

export enum ServiceProvider {
  Email = 'email',
  Whatsapp = 'whatsapp'
}

export type Subscription = {
  __typename?: 'Subscription';
  subNewMessage: MessageData;
  subRoomDetails: RoomDataDetails;
};


export type SubscriptionSubNewMessageArgs = {
  code: Scalars['String']['input'];
};


export type SubscriptionSubRoomDetailsArgs = {
  code: Scalars['String']['input'];
};

export type UserData = {
  __typename?: 'UserData';
  audit: Audit;
  email: Scalars['String']['output'];
  firstName: Scalars['String']['output'];
  id: Scalars['String']['output'];
  image: Scalars['String']['output'];
  lastName: Scalars['String']['output'];
};

export type UserLoginPayload = {
  email: Scalars['String']['input'];
  roomCode: Scalars['String']['input'];
};

export type UserLoginResponse = {
  __typename?: 'UserLoginResponse';
  profile: UserData;
};

export type UserSearchPayload = {
  email: Scalars['String']['input'];
  firstName: Scalars['String']['input'];
  lastName: Scalars['String']['input'];
  limit?: InputMaybe<Scalars['Int']['input']>;
  skip?: InputMaybe<Scalars['Int']['input']>;
};

export type UserSearchResponse = {
  __typename?: 'UserSearchResponse';
  items: Array<UserData>;
  total?: Maybe<Scalars['Int']['output']>;
};

export type UserUpdatePayload = {
  firstName: Scalars['String']['input'];
  image: Scalars['String']['input'];
  lastName: Scalars['String']['input'];
};

export type GetMessagesQueryVariables = Exact<{
  payload: MessageSearchPayload;
}>;


export type GetMessagesQuery = { __typename?: 'Query', getMessages: { __typename?: 'MessageSearchResponse', total?: number | null, items: Array<{ __typename?: 'MessageData', id: string, roomId: string, fromId: string, message: string, isNew: boolean, audit: { __typename?: 'Audit', createdAt: any, createdBy?: string | null, updatedAt?: any | null, updatedBy?: string | null } }> } };

export type MessageSendMutationVariables = Exact<{
  payload: MessageCreatePayload;
}>;


export type MessageSendMutation = { __typename?: 'Mutation', sendMessage: { __typename?: 'MessageData', id: string, roomId: string, fromId: string, message: string, isNew: boolean, audit: { __typename?: 'Audit', createdAt: any, createdBy?: string | null, updatedAt?: any | null, updatedBy?: string | null } } };

export type SubNewMessageSubscriptionVariables = Exact<{
  code: Scalars['String']['input'];
}>;


export type SubNewMessageSubscription = { __typename?: 'Subscription', subNewMessage: { __typename?: 'MessageData', id: string, roomId: string, fromId: string, message: string, isNew: boolean, audit: { __typename?: 'Audit', createdAt: any, createdBy?: string | null, updatedAt?: any | null, updatedBy?: string | null } } };

export type SubRoomDetailsSubscriptionVariables = Exact<{
  code: Scalars['String']['input'];
}>;


export type SubRoomDetailsSubscription = { __typename?: 'Subscription', subRoomDetails: { __typename?: 'RoomDataDetails', id: string, code: string, image: string, users: Array<{ __typename?: 'UserData', id: string, firstName: string, lastName: string, email: string, image: string, audit: { __typename?: 'Audit', createdAt: any, createdBy?: string | null, updatedAt?: any | null, updatedBy?: string | null } }>, audit: { __typename?: 'Audit', createdAt: any, createdBy?: string | null, updatedAt?: any | null, updatedBy?: string | null } } };

export type UsersQueryVariables = Exact<{
  payload: UserSearchPayload;
}>;


export type UsersQuery = { __typename?: 'Query', getUsers: { __typename?: 'UserSearchResponse', total?: number | null, items: Array<{ __typename?: 'UserData', id: string, firstName: string, lastName: string, email: string, image: string, audit: { __typename?: 'Audit', createdAt: any, createdBy?: string | null, updatedAt?: any | null, updatedBy?: string | null } }> } };

export type UserLoginMutationVariables = Exact<{
  payload: UserLoginPayload;
}>;


export type UserLoginMutation = { __typename?: 'Mutation', userLogin: { __typename?: 'UserLoginResponse', profile: { __typename?: 'UserData', id: string, email: string, firstName: string, lastName: string, image: string, audit: { __typename?: 'Audit', createdAt: any, createdBy?: string | null, updatedAt?: any | null, updatedBy?: string | null } } } };

export type UserProfileQueryVariables = Exact<{ [key: string]: never; }>;


export type UserProfileQuery = { __typename?: 'Query', userProfile: { __typename?: 'UserData', id: string, firstName: string, lastName: string, email: string, image: string, audit: { __typename?: 'Audit', createdAt: any, createdBy?: string | null, updatedAt?: any | null, updatedBy?: string | null } } };

export type UserUpdateMutationVariables = Exact<{
  payload: UserUpdatePayload;
}>;


export type UserUpdateMutation = { __typename?: 'Mutation', userUpdateProfile: { __typename?: 'UserData', id: string, firstName: string, lastName: string, email: string, image: string, audit: { __typename?: 'Audit', createdAt: any, createdBy?: string | null, updatedAt?: any | null, updatedBy?: string | null } } };
