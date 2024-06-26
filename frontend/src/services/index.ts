import { apolloClient } from "@/helper/graphql";
import { messageGql } from "./message";
import { roomGql } from "./room";
import { userGql } from "./user";

export default {
  user: userGql(apolloClient()),
  room: roomGql(apolloClient()),
  message: messageGql(apolloClient()),
}