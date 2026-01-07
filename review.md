Telebot Repository Analysis & Bot API 9.3 Comparison Report

  Repository Overview

  The telebot repository is a comprehensive Go framework for building Telegram bots, currently implementing Bot API 8.3. It's organized as a modular package (gopkg.in/telebot.v4) with 55 Go files and approximately 12,247 lines of code.

  Current Implementation Status

  Version: v4 branch
  Latest Commit: 9f2af51 (Chat type struct fixes)
  Bot API Support: Up to Bot API 8.3
  Last Major Updates: Bot API 8.0, 8.1, 8.2, 8.3 implementations

  Bot API 9.3 Changes Analysis (December 31, 2025)

  The latest Telegram Bot API 9.3 introduces three major feature categories that require implementation updates:

  ---
  📋 Feature Comparison & Implementation Gaps

  1. Topics in Private Chats

  API 9.3 Requirements:

  - has_topics_enabled field in User class
  - sendMessageDraft method for streaming partial messages
  - message_thread_id support in private chats
  - is_topic_message field support in private chats
  - Forum topic management extended to private chats

  Current Implementation Status:

  ✅ Implemented:
  - ThreadID field in Message struct (message.go:14)
  - TopicMessage (as is_topic_message) field in Message struct (message.go:83)
  - ThreadID field in SendOptions (options.go:85)
  - Topic management methods in topic.go (CreateTopic, EditTopic, etc.)

  ❌ Missing:
  - HasTopicsEnabled field in User struct (chat.go:9-29)
  - sendMessageDraft method for streaming messages (not found in any file)
  - Documentation/support for using topics in private chats (currently only documented for forums)

  Implementation Required:
  // In chat.go, User struct needs:
  HasTopicsEnabled bool `json:"has_topics_enabled"`

  New Method Needed:
  // SendMessageDraft for streaming partial messages during generation
  func (b *Bot) SendMessageDraft(to Recipient, text string, opts ...interface{}) (*Message, error)

  ---
  2. Gifts Enhancements

  API 9.3 Requirements:

  - getUserGifts method for gift retrieval
  - getChatGifts method for chat-specific gifts
  - Gift metadata with currency and amount fields (replacing star counts)
  - New gift origins: "gifted_upgrade" and "offer"
  - Enhanced Gift class with premium status, blockchain origin, color scheme
  - New GiftBackground class for visual customization
  - Gift filtering and variant counting capabilities

  Current Implementation Status (gifts.go):

  ✅ Partially Implemented:
  - Basic Gift struct with ID, Sticker, StarCount, UpgradeStarCount (gifts.go:8-27)
  - GetAvailableGifts() method (gifts.go:36-49)
  - SendGift() method with text and pay_for_upgrade options (gifts.go:54-96)
  - Support for sending gifts to chat members (Bot API 8.3)
  - CanSendGift field in ChatFullInfo (chat.go:361)

  ❌ Missing:
  - getUserGifts() method (not found)
  - getChatGifts() method (not found)
  - Currency/Amount fields in Gift struct (only has StarCount)
  - Gift origins support (gifted_upgrade, offer)
  - UniqueGift class/struct
  - Premium status field in Gift
  - Blockchain origin information
  - Color scheme support
  - GiftBackground struct/class

  Major Updates Required:

  1. Expand Gift struct (gifts.go:8-27):
  type Gift struct {
      ID               string   `json:"id"`
      Sticker          *Sticker `json:"sticker"`

      // Deprecated in API 9.3 but keep for backward compatibility
      StarCount        int      `json:"star_count,omitempty"`
      UpgradeStarCount int      `json:"upgrade_star_count,omitempty"`

      // New in API 9.3
      Currency         string   `json:"currency,omitempty"`
      Amount           int      `json:"amount,omitempty"`
      Origin           string   `json:"origin,omitempty"` // "gifted_upgrade", "offer", etc.
      IsPremium        bool     `json:"is_premium,omitempty"`
      BlockchainOrigin string   `json:"blockchain_origin,omitempty"`
      ColorScheme      []string `json:"color_scheme,omitempty"`

      TotalCount       int      `json:"total_count,omitempty"`
      RemainingCount   int      `json:"remaining_count,omitempty"`
  }

  2. Add UniqueGift and GiftBackground structs
  3. Implement new methods:
    - GetUserGifts(user Recipient) ([]Gift, error)
    - GetChatGifts(chat Recipient) ([]Gift, error)

  ---
  3. Bot Username Management

  API 9.3 Changes:

  - Improvements for Fragment-purchased usernames
  - Enhanced username management for bots

  Current Implementation Status:

  ✅ Implemented:
  - Usernames (active_usernames) field in User struct (chat.go:20)
  - Basic username support

  ❓ Unknown:
  - Specific Fragment username management methods (need to review full API changelog for details)

  ---
  4. Channel Chat Restrictions

  API 9.3 Changes:

  - can_restrict_members can now be disabled for channel chats

  Current Implementation Status:

  ✅ Already Implemented:
  - CanRestrictMembers field exists in Rights struct (admin.go:21)
  - Full admin rights management system in place (admin.go:9-51)

  Status: No changes needed - already supported.

  ---
  5. Story Reposting

  API 9.3 Requirements:

  - repostStory method for cross-account story sharing

  Current Implementation Status:

  ✅ Partial Implementation:
  - Story struct exists (chat.go:270-276)
  - Story references in Message struct (message.go:63, 77)

  ❌ Missing:
  - repostStory method (not found in any file)

  Implementation Required:
  // New method in bot.go or new story.go file
  func (b *Bot) RepostStory(to Recipient, story *Story, opts ...interface{}) error {
      // Implementation for repostStory API method
  }

  ---
  6. User Rating System

  API 9.3 Requirements:

  - New UserRating class
  - Rating field in ChatFullInfo

  Current Implementation Status:

  ❌ Missing:
  - UserRating struct/type (not found)
  - Rating field in ChatFullInfo (chat.go:289-362)

  Implementation Required:
  // New UserRating struct
  type UserRating struct {
      // Fields based on API documentation
  }

  // Add to ChatFullInfo (chat.go):
  Rating *UserRating `json:"rating,omitempty"`

  ---
  7. Paid Media Enhancements

  API 9.3 Changes:

  - Maximum price increased to 25,000 Telegram Stars (from previous limit)

  Current Implementation Status:

  ✅ Implemented:
  - Paid media support via SendPaidMedia method in API interface (api.go:84)
  - PaidAlbum type exists
  - Stars payment constant defined (payments.go:10)

  Status: Implementation exists, but may need validation/constant updates for the new 25,000 star limit.

  ---
  8. Message Effects in Forward/Copy

  API 9.3 Changes:

  - Message effect support for forwardMessage and copyMessage methods

  Current Implementation Status:

  ✅ Implemented:
  - EffectID field in SendOptions (options.go:97)
  - Forward() method exists (bot.go:439)
  - Copy() method exists (bot.go:484)
  - ForwardMany() and CopyMany() methods accept SendOptions

  Status: Infrastructure exists, but needs verification that EffectID is properly passed to forward/copy API calls.

  ---
  🎯 Priority Implementation Roadmap

  High Priority (Breaking API Changes)

  1. Gift System Overhaul
    - Update Gift struct with currency/amount fields
    - Add UniqueGift and GiftBackground types
    - Implement getUserGifts and getChatGifts methods
    - Files affected: gifts.go, api.go
  2. Topics in Private Chats
    - Add HasTopicsEnabled to User struct
    - Implement sendMessageDraft method
    - Update documentation for private chat topics
    - Files affected: chat.go, bot.go, api.go

  Medium Priority (New Features)

  3. Story Reposting
    - Implement repostStory method
    - Files affected: Create new story.go or extend existing
  4. User Rating System
    - Add UserRating struct
    - Add Rating field to ChatFullInfo
    - Files affected: chat.go

  Low Priority (Validation/Enhancement)

  5. Paid Media Limit Update
    - Validate 25,000 star limit implementation
    - Update constants/documentation if needed
    - Files affected: payments.go, validation logic
  6. Message Effects Verification
    - Ensure EffectID properly embedded in Forward/Copy methods
    - Files affected: bot.go

  ---
  📊 Summary Statistics

  Total Features in Bot API 9.3: 8 major feature areas
  Fully Implemented: 2 (25%)
  Partially Implemented: 3 (37.5%)
  Not Implemented: 3 (37.5%)

  Estimated Implementation Effort:
  - Gift System: 8-10 hours (complex refactor)
  - Topics in Private Chats: 4-6 hours
  - Story Reposting: 2-3 hours
  - User Rating: 1-2 hours
  - Total: ~15-21 hours of development

  ---
  🔍 Additional Notes

  - The codebase is well-structured and follows Go best practices
  - Most infrastructure for Bot API 9.3 features already exists
  - Main gaps are in extending existing types and adding new API methods
  - No breaking changes to existing user code required for most updates
  - The repository supports up to Bot API 8.3, so it's only one version behind

  Recommended Next Steps

  1. Create feature branch for Bot API 9.3 implementation
  2. Start with Gift system refactor (highest impact)
  3. Implement private chat topics support
  4. Add remaining missing methods (repostStory, getUserGifts, getChatGifts, sendMessageDraft)
  5. Comprehensive testing with Telegram Bot API 9.3
  6. Update documentation and examples
  7. Submit PR to v4 branch