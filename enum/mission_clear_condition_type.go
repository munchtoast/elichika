package enum

const (
	MissionClearConditionTypeCountEventMarathonPoint               int32 = 0x00000001 // total point in event marathon
	MissionClearConditionTypeCountEventPlay                        int32 = 0x00000002 // unused
	MissionClearConditionTypeEventRank                             int32 = 0x00000003 // unused
	MissionClearConditionTypeCountLoveLevel                        int32 = 0x00000004 // # of member at a certain love level
	MissionClearConditionTypeCountCard                             int32 = 0x00000005 // # of time we gain a card, not # of different card
	MissionClearConditionTypeCountSpecificMember                   int32 = 0x00000006 // # of different card a member has
	MissionClearConditionTypeCountCardType                         int32 = 0x00000007 // unused
	MissionClearConditionTypeClearedStory                          int32 = 0x00000008 // unused
	MissionClearConditionTypeClearedEpisode                        int32 = 0x00000009 // # of different bond episode (member story) read
	MissionClearConditionTypeCountLinkSifid                        int32 = 0x0000000a // unused
	MissionClearConditionTypeUserRank                              int32 = 0x0000000b // current user rank
	MissionClearConditionTypeCountEditProfile                      int32 = 0x0000000c // NOT SUPPORTED
	MissionClearConditionTypeCountPlayLive                         int32 = 0x0000000d // play live, presumablly doesn't require actually clearing it
	MissionClearConditionTypeCountClearedLive                      int32 = 0x0000000e // actually clearing started live, skip ticket works
	MissionClearConditionTypeCountClearedLiveDifficultyType        int32 = 0x0000000f // clear live with certain difficulty
	MissionClearConditionTypeCountClearedLiveResultEvaluationType  int32 = 0x00000010 // unused
	MissionClearConditionTypeCountFullCombo                        int32 = 0x00000011 // unused
	MissionClearConditionTypeCountPerfectFullCombo                 int32 = 0x00000012 // get an all perfect
	MissionClearConditionTypeAppealVoltage                         int32 = 0x00000013 // get a tap score of at least some thing
	MissionClearConditionTypeCountClearedNoDamage                  int32 = 0x00000014 // unused
	MissionClearConditionTypeCountClearedSpecificLiveDeck          int32 = 0x00000015 // unused
	MissionClearConditionTypeCountClearedSpecificGroup             int32 = 0x00000016 // clear with only member of a group (muse), no skip ticket
	MissionClearConditionTypeCountClearedSpecificUnit              int32 = 0x00000017 // clear with only member of a subunit (lily white), no skip ticket
	MissionClearConditionTypeCountClearedSpecificMember            int32 = 0x00000018 // clear with a member in the formation, no skip ticket
	MissionClearConditionTypeCountClearedSpecificOnlyMember        int32 = 0x00000019 // unused
	MissionClearConditionTypeCountClearedSpecificOnlyCardType      int32 = 0x0000001a // unused
	MissionClearConditionTypeCountClearedSpecificOnlySchoolYear    int32 = 0x0000001b // clear with only a school year
	MissionClearConditionTypeCountUseSkill                         int32 = 0x0000001c // # of skill usage
	MissionClearConditionTypeCountUseSpecialSkill                  int32 = 0x0000001d // # of SP skill usage
	MissionClearConditionTypeCountUseVoltageSkill                  int32 = 0x0000001e // # of skill that give +voltage
	MissionClearConditionTypeCountUseRecoverySkill                 int32 = 0x0000001f // # of skill that recover stamina
	MissionClearConditionTypeCountUseDebuffSkill                   int32 = 0x00000020 // unused
	MissionClearConditionTypeCountUseBuffSkill                     int32 = 0x00000021 // # of buff skill
	MissionClearConditionTypeCountClearedNotUseVoltageSkill        int32 = 0x00000022 // unused
	MissionClearConditionTypeCountClearedNotUseRecoverySkill       int32 = 0x00000023 // unused
	MissionClearConditionTypeCountClearedNotUseDebuffSkill         int32 = 0x00000024 // unused
	MissionClearConditionTypeCountClearedNotUseBuffSkill           int32 = 0x00000025 // unused
	MissionClearConditionTypeCountClearedLowDamage                 int32 = 0x00000026 // unused
	MissionClearConditionTypeClearedUnderUniqueMember              int32 = 0x00000028 // clear using less than # of unique member
	MissionClearConditionTypeClearedOverUniqueMember               int32 = 0x00000029 // unused
	MissionClearConditionTypeAppealRecoveryValue                   int32 = 0x0000002a // restore # of stamina at once
	MissionClearConditionTypeAppealShieldValue                     int32 = 0x0000002b // gain # of shield at once
	MissionClearConditionTypeLiveRecoveryValue                     int32 = 0x0000002c // unused
	MissionClearConditionTypeCountLiveOverShield                   int32 = 0x0000002d // unused
	MissionClearConditionTypeCountLiveContinuingTaps               int32 = 0x0000002e // unused
	MissionClearConditionTypeLiveAppealAscendingValue              int32 = 0x0000002f // unused
	MissionClearConditionTypeLiveStaminaDamageDown                 int32 = 0x00000030 // unused
	MissionClearConditionTypeLiveStaminaDamageDownAndNoDamage      int32 = 0x00000031 // unused
	MissionClearConditionTypeCountUseSpecificMemberSkill           int32 = 0x00000032 // user a member's skill
	MissionClearConditionTypeCountUseSpecificMemberSpecialSkill    int32 = 0x00000033 // use an sp skill with a member in it
	MissionClearConditionTypeCountClearedSpecificMemberAndPosition int32 = 0x00000034 // clear # of live show with a member at a specific position, in practice the position = 1
	MissionClearConditionTypeCountLogin                            int32 = 0x00000035 // login for a certain amount of day
	MissionClearConditionTypeCompleteDaily                         int32 = 0x00000036 // clear a daily mission
	MissionClearConditionTypeCompleteWeekly                        int32 = 0x00000037 // clear a weekly mission
	MissionClearConditionTypeCountClearedMission                   int32 = 0x00000038 // unused
	MissionClearConditionTypeCountPassiveSkill                     int32 = 0x00000039 // equip insight skill (different number of insight across all cards)
	MissionClearConditionTypeCountLesson                           int32 = 0x0000003a // perform lesson
	MissionClearConditionTypeCountGradeEmblem                      int32 = 0x0000003b // gain a title with a certain grade or higher
	MissionClearConditionTypeCountActiveEmblem                     int32 = 0x0000003c // change the title
	MissionClearConditionTypeCountTrainingTree                     int32 = 0x0000003d // unused
	MissionClearConditionTypeCountEnhancementSkill                 int32 = 0x0000003e // unused
	MissionClearConditionTypeCountSpecificRarityMakeLevelMax       int32 = 0x0000003f // raise # of card of specfic rarity to a certain max level
	MissionClearConditionTypeCountSchoolIdolAwakening              int32 = 0x00000040 // Idolize idols
	MissionClearConditionTypeCountVoice                            int32 = 0x00000041 // get # of voice line
	MissionClearConditionTypeCountSuit                             int32 = 0x00000042 // get # of costume
	MissionClearConditionTypeCountEditLiveDeck                     int32 = 0x00000043 // edit live deck
	MissionClearConditionTypeCountEditLessonDeck                   int32 = 0x00000044 // edit lesson deck
	MissionClearConditionTypeCountFriend                           int32 = 0x00000045 // get a friend (old)
	MissionClearConditionTypeCountSchoolIdolGrade                  int32 = 0x00000046 // perform a limit break
	MissionClearConditionTypeCompleteEvent                         int32 = 0x00000047 // unused
	MissionClearConditionTypeCountTrainingTreeCellType             int32 = 0x00000048 // unlock training tree of a certain tile type
	MissionClearConditionTypeCountLiveMission                      int32 = 0x00000049 // clear live show's trophy
	MissionClearConditionTypeEditLivePartyMember                   int32 = 0x0000004a // change live party member
	MissionClearConditionTypeEditLivePartyAccessory                int32 = 0x0000004b // change accessory
	MissionClearConditionTypeCountChangeLiveParty                  int32 = 0x0000004c // unused
	MissionClearConditionTypeCountAccessoryType                    int32 = 0x0000004d // get different type of accessory
	MissionClearConditionTypeCountAccessory                        int32 = 0x0000004e // get accessory
	MissionClearConditionTypeCountAccessoryLevelUp                 int32 = 0x0000004f // level up accessory x times
	MissionClearConditionTypeCountAccessoryGradeUp                 int32 = 0x00000050 // grade up accessory x times
	MissionClearConditionTypeCountAccessoryRarityUp                int32 = 0x00000051 // rarity up accessory x times
	MissionClearConditionTypeCountAccessoryMelt                    int32 = 0x00000052 // melt accessorty x times
	MissionClearConditionTypeCompleteBeginner                      int32 = 0x00000053 // complete beginner mission (old)
	MissionClearConditionTypeDataLink                              int32 = 0x00000054 // unused
	MissionClearConditionTypeReadReferenceBook                     int32 = 0x00000055 // read reference book for beginner mission
	MissionClearConditionTypeCountPlayLiveDailyMusic               int32 = 0x00000056 // play daily song
	MissionClearConditionTypeCountClearedLiveDailyMusic            int32 = 0x00000057 // clear daily song
	MissionClearConditionTypeCountEventMiningPoint                 int32 = 0x00000058 // unused
	MissionClearConditionTypeTargetVoltage                         int32 = 0x00000059 // unused
	MissionClearConditionTypeSpSkillVoltage                        int32 = 0x0000005a // unused
	MissionClearConditionTypeCountStoryLiveMission                 int32 = 0x0000005b // unused
	MissionClearConditionTypeCountStoryMainStill                   int32 = 0x0000005c // unused
	MissionClearConditionTypeCountEventCoopAward                   int32 = 0x0000005d // get SBL awards
	MissionClearConditionTypeMemberLovePanelCell                   int32 = 0x0000005e // unlock love panel cell
	MissionClearConditionTypeMemberLovePanel                       int32 = 0x0000005f // fully unlock a member's bond board
	MissionClearConditionTypeTowerPlayLive                         int32 = 0x00000060 // play a live in tower
	MissionClearConditionTypeTowerClearLiveStage                   int32 = 0x00000061 // complete dlp stage
	MissionClearConditionTypeTowerRecoverPpForPersonal             int32 = 0x00000062 // recover pp for member
	MissionClearConditionTypeTowerRecoverPpForFree                 int32 = 0x00000063 // recover pp for all member
	MissionClearConditionTypeClearedStorySide                      int32 = 0x00000064 // unused
	MissionClearConditionTypeClearedStoryMainChapter               int32 = 0x00000065 // clear story chapter
	MissionClearConditionTypeCountReadDailyTheater                 int32 = 0x00000066 // # of different daily theater read
	MissionClearConditionTypeClearedSRank                          int32 = 0x00000067 // clear a live difficulty on S rank
	MissionClearConditionTypeCountCompleteTraining                 int32 = 0x00000068 // unused
)
