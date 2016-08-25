package api

// Message represents a message
type Message struct {
    // ID is unique message identifier
    ID int32 `json:"message_id"`
    
    // User	is a sender, can be empty for messages sent to channels
    // Optional
    From User `json:"from"`
    
    // Date of the message was sent in Unix time
    Date int32 `json:"date"`
    
    // Chat is a conversation the message belongs to
    Chat Chat `json:"chat"`
    
    // ForwardFrom for forwarded messages, sender of the original message
    // Optional
    ForwardFrom User `json:"forward_from"`
    
    // ForwardFromChat for messages forwarded from a channel, information about the original channel
    // Optional
    ForwardFromChat Chat `json:"forward_from_chat"`
    
    // ForwardDate for forwarded messages, date the original message was sent in Unix time
    // Optional
    ForwardDate int32 `json:"forward_date"`
    
    // ReplyToMessage is the original message for replies.
    // Note that the Message object in this field will not contain further reply_to_message fields
    // even if it itself is a reply.
    // Optional
    ReplyToMessage Message `json:"reply_to_message"`
    
    // EditDate the message was last edited in Unix time
    // Optional
    EditDate int32 `json:"edit_date"`
    
    // Text for text messages, the actual UTF-8 text of the message, 0-4096 characters.
    // Optional
    Text string `json:"text"`
    
    // Entities for text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
    // Optional
    Entities []MessageEntity `json:"entities"`
    
    // Audio is a message with an audio file, information about the file
    // Optional
    Audio Audio `json:"audio"`
    
    // Document is a message with a general file, information about the file
    // Optional
    Document Document `json:"document"`
    
    // Photo is a photo message, available sizes of the photo
    // Optional
    Photo []PhotoSize `json:"photo"`
    
    // Sticker is a sticker message, information about the sticker
    // Optional
    Sticker Sticker `json:"sticker"`
    
    // Video is a video message, information about the video
    // Optional
    Video Video `json:"video"`
    
    // Voice is a voice message, information about the file
    // Optional
    Voice Voice `json:"voice"`
    
    // Caption for the document, photo or video, 0-200 characters
    // Optional
    Caption string `json:"caption"`
    
    // Contact is a shared contact message, information about the contact
    // Optional
    Contact Contact `json:"contact"`
    
    // Location is a shared location message, information about the location
    // Optional
    Location Location `json:"location"`
    
    // Venue is a venue message, information about the venue
    // Optional
    Venue Venue `json:"venue"`
    
    // NewChatMember was added to the group, information about them (this member may be the bot itself)
    // Optional
    NewChatMember User `json:"new_chat_member"`
    
    // LeftChatMemeber was removed from the group, information about them (this member may be the bot itself)
    // Optional
    LeftChatMemeber User `json:"left_chat_member"`
    
    // NewChatTitle is a chat title that was changed to
    // Optional
    NewChatTitle string `json:"new_chat_title"`
    
    // NewChatPhoto is a chat photo was change to
    // Optional
    NewChatPhoto []PhotoSize `json:"new_chat_photo"`
    
    // DeleteChatPhoto is True if exist.
    // Service message: the chat photo was deleted
    // Optional
    DeleteChatPhoto bool `json:"delete_chat_photo"`
    
    // GroupChatCreated is True if exist.
    // Service message: the group has been created
    // Optional
    GroupChatCreated bool `json:"group_chat_created"`
    
    // SuperGroupChatCreated is True if exist.
    // Service message: the supergroup has been created
    // Optional
    GroupChatCreated bool `json:"supergroup_chat_created"`
    
    // ChannelChatCreated is True if exist.
    // Service message: the channel has been created
    // Optional
    ChannelChatCreated bool `json:"channel_chat_created"`
    
    // MigrateToChatID The group has been migrated to a supergroup with the specified identifier.
    // Optional
    MigrateToChatID int64 `json:"migrate_to_chat_id"`
    
    // MigrateFromChatID The supergroup has been migrated from a group with the specified identifier
    // Optional
    MigrateFromChatID int64 `json:"migrate_from_chat_id"`
    
    // PinnedMessage specified message was pinned.
    // Note that the Message object in this field will not contain
    // further reply_to_message fields even if it is itself a reply.
    // Optional
    PinnedMessage Message `json:"pinned_message"`
}