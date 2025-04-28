package telegram

type ChatFullInfo struct {
	ID                                 int64                 `json:"id"`
	Bot                                bool                  `json:"is_bot"`
	LanguageCode                       string                `json:"language_code"`
	Prmium                             bool                  `json:"is_premium,omitempty"`
	Type                               string                `json:"type"`
	Title                              *string               `json:"title,omitempty"`
	Username                           *string               `json:"username,omitempty"`
	FirstName                          *string               `json:"first_name,omitempty"`
	LastName                           *string               `json:"last_name,omitempty"`
	IsForum                            *bool                 `json:"is_forum,omitempty"`
	AccentColorID                      int                   `json:"accent_color_id"`
	MaxReactionCount                   int                   `json:"max_reaction_count"`
	Photo                              *ChatPhoto            `json:"photo,omitempty"`
	ActiveUsernames                    []string              `json:"active_usernames,omitempty"`
	Birthdate                          *Birthdate            `json:"birthdate,omitempty"`
	BusinessIntro                      *BusinessIntro        `json:"business_intro,omitempty"`
	BusinessLocation                   *BusinessLocation     `json:"business_location,omitempty"`
	BusinessOpeningHours               *BusinessOpeningHours `json:"business_opening_hours,omitempty"`
	PersonalChat                       *ChatFullInfo         `json:"personal_chat,omitempty"`
	AvailableReactions                 []ReactionType        `json:"available_reactions,omitempty"`
	BackgroundCustomEmojiID            *string               `json:"background_custom_emoji_id,omitempty"`
	ProfileAccentColorID               *int                  `json:"profile_accent_color_id,omitempty"`
	ProfileBackgroundCustomEmojiID     *string               `json:"profile_background_custom_emoji_id,omitempty"`
	EmojiStatusCustomEmojiID           *string               `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate          *int64                `json:"emoji_status_expiration_date,omitempty"`
	Bio                                *string               `json:"bio,omitempty"`
	HasPrivateForwards                 *bool                 `json:"has_private_forwards,omitempty"`
	HasRestrictedVoiceAndVideoMessages *bool                 `json:"has_restricted_voice_and_video_messages,omitempty"`
	JoinToSendMessages                 *bool                 `json:"join_to_send_messages,omitempty"`
	JoinByRequest                      *bool                 `json:"join_by_request,omitempty"`
	Description                        *string               `json:"description,omitempty"`
	InviteLink                         *string               `json:"invite_link,omitempty"`
	PinnedMessage                      *Message              `json:"pinned_message,omitempty"`
	AcceptedGiftTypes                  *AcceptedGiftTypes    `json:"accepted_gift_types,omitempty"`
	CanSendPaidMedia                   *bool                 `json:"can_send_paid_media,omitempty"`
	SlowModeDelay                      *int                  `json:"slow_mode_delay,omitempty"`
	UnrestrictBoostCount               *int                  `json:"unrestrict_boost_count,omitempty"`
	MessageAutoDeleteTime              *int                  `json:"message_auto_delete_time,omitempty"`
	HasAggressiveAntiSpamEnabled       *bool                 `json:"has_aggressive_anti_spam_enabled,omitempty"`
	HasHiddenMembers                   *bool                 `json:"has_hidden_members,omitempty"`
	HasProtectedContent                *bool                 `json:"has_protected_content,omitempty"`
	HasVisibleHistory                  *bool                 `json:"has_visible_history,omitempty"`
	StickerSetName                     *string               `json:"sticker_set_name,omitempty"`
	CanSetStickerSet                   *bool                 `json:"can_set_sticker_set,omitempty"`
	CustomEmojiStickerSetName          *string               `json:"custom_emoji_sticker_set_name,omitempty"`
	LinkedChatID                       *int64                `json:"linked_chat_id,omitempty"`
	Location                           *ChatLocation         `json:"location,omitempty"`
}

type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

type BusinessIntro struct {
	Title   *string  `json:"title,omitempty"`
	Message *string  `json:"message,omitempty"`
	Sticker *Sticker `json:"sticker,omitempty"`
}

type BusinessLocation struct {
	Address  string    `json:"address"`
	Location *Location `json:"location,omitempty"`
}

type BusinessOpeningHours struct {
	TimeZoneName string                         `json:"time_zone_name"`
	OpeningHours []BusinessOpeningHoursInterval `json:"opening_hours"`
}

type BusinessOpeningHoursInterval struct {
	OpeningMinute int `json:"opening_minute"`
	ClosingMinute int `json:"closing_minute"`
}

type ReactionType interface {
	isReactionType()
}

type ReactionTypeEmoji struct {
	Type  string `json:"type"` // "emoji"
	Emoji string `json:"emoji"`
}

func (ReactionTypeEmoji) isReactionType() {}

type ReactionTypeCustomEmoji struct {
	Type          string `json:"type"` // "custom_emoji"
	CustomEmojiID string `json:"custom_emoji_id"`
}

func (ReactionTypeCustomEmoji) isReactionType() {}

type ReactionTypePaid struct {
	Type string `json:"type"` // "paid"
}

func (ReactionTypePaid) isReactionType() {}

type ReactionCount struct {
	Type       ReactionType `json:"type"`
	TotalCount int          `json:"total_count"`
}

type MessageReactionUpdated struct {
	Chat        ChatFullInfo   `json:"chat"`
	MessageID   int            `json:"message_id"`
	User        *ChatFullInfo  `json:"user,omitempty"`
	ActorChat   *ChatFullInfo  `json:"actor_chat,omitempty"`
	Date        int64          `json:"date"`
	OldReaction []ReactionType `json:"old_reaction"`
	NewReaction []ReactionType `json:"new_reaction"`
}

type MessageReactionCountUpdated struct {
	Chat      ChatFullInfo    `json:"chat"`
	MessageID int             `json:"message_id"`
	Date      int64           `json:"date"`
	Reactions []ReactionCount `json:"reactions"`
}

type Sticker struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Type         string `json:"type"`
	Emoji        string `json:"emoji"`
	SetName      string `json:"set_name"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type AcceptedGiftTypes struct {
	GiftTypeIDs []int `json:"gift_type_ids"`
}

type ChatLocation struct {
	Location Location `json:"location"`
	Address  string   `json:"address"`
}
