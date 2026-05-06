// get the conversations paginator
const conversationsPaginator = await client.getSubscribedConversations();

// get conversations
const conversations = conversationsPaginator.items;
