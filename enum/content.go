package enum

const (
	// from m_content_setting, maybe this can be loaded by DB but meh
	ContentTypeSnsCoin               = 1
	ContentTypeCard                  = 3
	ContentTypeCardExp               = 4
	ContentTypeGachaPoint            = 5
	ContentTypeLessonEnhancingItem   = 6
	ContentTypeSuit                  = 7
	ContentTypeVoice                 = 8
	ContentTypeGachaTicket           = 9
	ContentTypeGameMoney             = 10
	ContentTypeTrainingMaterial      = 12
	ContentTypeCardExchange          = 13 // grade up items
	ContentTypeEmblem                = 15
	ContentTypeSheetRecoveryAP       = 16
	ContentTypeRecoveryLP            = 17
	ContentTypeStorySide             = 19
	ContentTypeStoryMember           = 20
	ContentTypeExchangeEventPoint    = 21
	ContentTypeAccessory             = 23
	ContentTypeAccessoryLevelUpItem  = 24
	ContentTypeAccessoryRarityUpItem = 25
	ContentTypeCustomBackground      = 26
	ContentTypeEventMarathonBooster  = 27
	ContentTypeSkipTicket            = 28
	ContentTypeStoryEventUnlock      = 30
	// 31 32 33 don't have clear definition but they can rerecovered from other tables / context
	ContentTypeTowerRecoveryItem      = 31 // found in m_mission_reward by cross referencing content_id (known to be 24001)
	ContentTypeSubscriptionCoin       = 32 // found in m_trade as currency
	ContentTypeMemberGuildSupportItem = 33 // found in m_mission_reward by cross referencing content_id (known to be 27002)
)