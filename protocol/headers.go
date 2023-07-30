package protocol

const (
	PetStatusUpdateHeader = 1907 //error 404

	CfhTopicsMessageHeader                = 325
	FavoriteRoomsCountHeader              = 151
	UserCurrencyHeader                    = 2018
	RedeemVoucherOKHeader                 = 3336
	RoomUserShoutHeader                   = 1036
	RoomUserStatusHeader                  = 1640
	RoomUserDataHeader                    = 3920
	RoomAddRightsListHeader               = 2088
	RoomRemoveRightsListHeader            = 1327
	RoomRightsListHeader                  = 1284
	RoomUserHandItemHeader                = 1474
	RoomUsersHeader                       = 374
	FriendRequestHeader                   = 2219
	GuildBoughtHeader                     = 2808
	AddUserBadgeHeader                    = 2493
	RecyclerCompleteHeader                = 468
	GuildBuyRoomsHeader                   = 2159
	FriendsHeader                         = 3130
	StalkErrorHeader                      = 3048
	TradeCloseWindowHeader                = 1001
	RemoveFloorItemHeader                 = 2703
	InventoryPetsHeader                   = 3522
	UserCreditsHeader                     = 3475
	WiredTriggerDataHeader                = 383
	TradeStoppedHeader                    = 1373
	ModToolUserChatlogHeader              = 3377
	GuildInfoHeader                       = 1702
	UserPermissionsHeader                 = 411
	PetNameErrorHeader                    = 1503
	TradeStartFailHeader                  = 217
	AddHabboItemHeader                    = 2103
	InventoryBotsHeader                   = 3086
	CanCreateRoomHeader                   = 378
	MarketplaceBuyErrorHeader             = 2032
	BonusRareHeader                       = 1533
	HotelViewHeader                       = 122
	UpdateFriendHeader                    = 2800
	FloorItemUpdateHeader                 = 3776
	RoomAccessDeniedHeader                = 878
	GuildFurniWidgetHeader                = 3293
	GiftConfigurationHeader               = 2234
	UserClubHeader                        = 954
	InventoryBadgesHeader                 = 717
	RoomUserTypingHeader                  = 1717
	GuildJoinErrorHeader                  = 762
	RoomCategoriesHeader                  = 1562
	InventoryAchievementsHeader           = 2501
	MarketplaceItemInfoHeader             = 725
	RoomRelativeMapHeader                 = 2753
	ModToolHeaderTwo                      = 2335
	ModToolHeaderOne                      = 3192
	RoomRightsHeader                      = 780
	ObjectOnRollerHeader                  = 3207
	PollStartHeader                       = 3785
	GuildRefreshMembersListHeader         = 2445
	UserPerksHeader                       = 2586
	UserCitizinShipHeader                 = 1203
	PublicRoomsHeader                     = -1 //error 404
	MarketplaceOffersHeader               = 680
	ModToolHeader                         = 2696
	UserBadgesHeader                      = 1087
	GuildManageHeader                     = 3965
	RemoveFriendHeader                    = -1 //error 404
	InitDiffieHandshakeHeader             = 1347
	CompleteDiffieHandshakeHeader         = 3885
	UserDataHeader                        = 2725
	UserSearchResultHeader                = 973
	ModToolUserRoomVisitsHeader           = 1752
	RoomUserRespectHeader                 = 2815
	RoomChatSettingsHeader                = 1191
	RemoveHabboItemHeader                 = 159
	RoomUserRemoveHeader                  = 2661
	RoomHeightMapHeader                   = 1301
	RoomPetHorseFigureHeader              = 1924
	PetErrorHeader                        = 2913
	TradeUpdateHeader                     = 2024
	PrivateRoomsHeader                    = 52
	RoomModelHeader                       = 2031
	RoomScoreHeader                       = 482
	DoorbellAddUserHeader                 = 2309
	SecureLoginOKHeader                   = 2491
	AvailabilityStatusMessageHeader       = 2033
	GuildMemberUpdateHeader               = 265
	RoomFloorItemsHeader                  = 1778
	InventoryItemsHeader                  = 994
	RoomUserTalkHeader                    = 1446
	TradeStartHeader                      = 2505
	InventoryItemUpdateHeader             = 104
	ModToolIssueUpdateHeader              = 3150
	MeMenuSettingsHeader                  = 513
	ModToolRoomInfoHeader                 = 1333
	GuildListHeader                       = 420
	RecyclerLogicHeader                   = 3164
	UserHomeRoomHeader                    = 2875
	RoomUserDanceHeader                   = 2233
	RoomSettingsUpdatedHeader             = 3297
	AlertPurchaseFailedHeader             = 1404
	RoomDataHeader                        = 687
	TagsHeader                            = 2012
	InventoryRefreshHeader                = 3151 // PRODUCTION-201611291003-338511768
	RemovePetHeader                       = 3253 // PRODUCTION-201611291003-338511768
	RemoveWallItemHeader                  = 3208 // PRODUCTION-201611291003-338511768
	TradeCompleteHeader                   = 2369 // PRODUCTION-201611291003-338511768
	NewsWidgetsHeader                     = 286  // PRODUCTION-201611291003-338511768
	WiredEffectDataHeader                 = 1434 // PRODUCTION-201611291003-338511768
	BubbleAlertHeader                     = 1992 // PRODUCTION-201611291003-338511768
	ReloadRecyclerHeader                  = 3433 // PRODUCTION-201611291003-338511768
	MoodLightDataHeader                   = 2710 // PRODUCTION-201611291003-338511768
	WiredRewardAlertHeader                = 178  // PRODUCTION-201611291003-338511768
	CatalogPageHeader                     = 804  // PRODUCTION-201611291003-338511768
	CatalogModeHeader                     = 3828 // PRODUCTION-201611291003-338511768
	ChangeNameUpdateHeader                = 118  // PRODUCTION-201611291003-338511768
	AddFloorItemHeader                    = 1534 // PRODUCTION-201611291003-338511768
	EnableNotificationsHeader             = 3284 // PRODUCTION-201611291003-338511768
	HallOfFameHeader                      = 3005 // PRODUCTION-201611291003-338511768
	WiredSavedHeader                      = 1155 // PRODUCTION-201611291003-338511768
	RoomPaintHeader                       = 2454 // PRODUCTION-201611291003-338511768
	MarketplaceConfigHeader               = 1823 // PRODUCTION-201611291003-338511768
	AddBotHeader                          = 1352 // PRODUCTION-201611291003-338511768
	FriendRequestErrorHeader              = 892  // PRODUCTION-201611291003-338511768
	GuildMembersHeader                    = 1200 // PRODUCTION-201611291003-338511768
	RoomOpenHeader                        = 758  // PRODUCTION-201611291003-338511768
	ModToolRoomChatlogHeader              = 3434 // PRODUCTION-201611291003-338511768
	DiscountHeader                        = 2347 // PRODUCTION-201611291003-338511768
	MarketplaceCancelSaleHeader           = 3264 // PRODUCTION-201611291003-338511768
	RoomPetRespectHeader                  = 2788 // PRODUCTION-201611291003-338511768
	RoomSettingsHeader                    = 1498 // PRODUCTION-201611291003-338511768
	TalentTrackHeader                     = 3406 // PRODUCTION-201611291003-338511768
	CatalogPagesListHeader                = 1032 // PRODUCTION-201611291003-338511768
	AlertLimitedSoldOutHeader             = 377  // PRODUCTION-201611291003-338511768
	CatalogUpdatedHeader                  = 1866 // PRODUCTION-201611291003-338511768
	PurchaseOKHeader                      = 869  // PRODUCTION-201611291003-338511768
	WallItemUpdateHeader                  = 2009 // PRODUCTION-201611291003-338511768
	TradeAcceptedHeader                   = 2568 // PRODUCTION-201611291003-338511768
	AddWallItemHeader                     = 2187 // PRODUCTION-201611291003-338511768
	RoomEntryInfoHeader                   = -1   // PRODUCTION-201611291003-338511768
	HotelViewDataHeader                   = 1745 // PRODUCTION-201611291003-338511768
	PresentItemOpenedHeader               = 56   // PRODUCTION-201611291003-338511768
	RoomUserRemoveRightsHeader            = 84   // PRODUCTION-201611291003-338511768
	UserBCLimitsHeader                    = -1   // PRODUCTION-201611291003-338511768
	PetTrainingPanelHeader                = 1164 // PRODUCTION-201611291003-338511768
	RoomPaneHeader                        = 749  // PRODUCTION-201611291003-338511768
	RedeemVoucherErrorHeader              = 714  // PRODUCTION-201611291003-338511768
	RoomCreatedHeader                     = 1304 // PRODUCTION-201611291003-338511768
	GenericAlertHeader                    = 3801 // PRODUCTION-201611291003-338511768
	GroupPartsHeader                      = 2238 // PRODUCTION-201611291003-338511768
	ModToolIssueInfoHeader                = 3609 // PRODUCTION-201611291003-338511768
	RoomUserWhisperHeader                 = 2704 // PRODUCTION-201611291003-338511768
	BotErrorHeader                        = 639  // PRODUCTION-201611291003-338511768
	FreezeLivesHeader                     = 2324 // PRODUCTION-201611291003-338511768
	LoadFriendRequestsHeader              = 280  // PRODUCTION-201611291003-338511768
	MarketplaceSellItemHeader             = 54   // PRODUCTION-201611291003-338511768
	ClubDataHeader                        = 2405 // PRODUCTION-201611291003-338511768
	ProfileFriendsHeader                  = 2016 // PRODUCTION-201611291003-338511768
	MarketplaceOwnItemsHeader             = 3884 // PRODUCTION-201611291003-338511768
	RoomOwnerHeader                       = 339  // PRODUCTION-201611291003-338511768
	WiredConditionDataHeader              = 1108 // PRODUCTION-201611291003-338511768
	ModToolUserInfoHeader                 = 2866 // PRODUCTION-201611291003-338511768
	UserWardrobeHeader                    = 3315 // PRODUCTION-201611291003-338511768
	RoomPetExperienceHeader               = 2156 // PRODUCTION-201611291003-338511768
	FriendChatMessageHeader               = 1587 // PRODUCTION-201611291003-338511768
	PetInformationHeader                  = 2901 // PRODUCTION-201611291003-338511768
	RoomThicknessHeader                   = 3547 // PRODUCTION-201611291003-338511768
	AddPetHeader                          = 2101 // PRODUCTION-201611291003-338511768
	UpdateStackHeightHeader               = 558  // PRODUCTION-201611291003-338511768
	RemoveBotHeader                       = 233  // PRODUCTION-201611291003-338511768
	RoomEnterErrorHeader                  = 899  // PRODUCTION-201611291003-338511768
	PollQuestionsHeader                   = 2997 // PRODUCTION-201611291003-338511768
	GenericErrorMessages                  = 1600 // PRODUCTION-201611291003-338511768
	RoomWallItemsHeader                   = 1369 // PRODUCTION-201611291003-338511768
	RoomUserEffectHeader                  = 1167 // PRODUCTION-201611291003-338511768
	PetBreedsHeader                       = 3331 // PRODUCTION-201611291003-338511768
	ModToolIssueChatlogHeader             = 607  // PRODUCTION-201611291003-338511768
	RoomUserActionHeader                  = 1631 // PRODUCTION-201611291003-338511768
	BotSettingsHeader                     = 1618 // PRODUCTION-201611291003-338511768
	UserProfileHeader                     = 3898 // PRODUCTION-201611291003-338511768
	MinimailCountHeader                   = 2803 // PRODUCTION-201611291003-338511768
	UserAchievementScoreHeader            = 1968 // PRODUCTION-201611291003-338511768
	PetLevelUpHeader                      = 859  // PRODUCTION-201611291003-338511768
	UserPointsHeader                      = 2275 // PRODUCTION-201611291003-338511768
	ReportRoomFormHeader                  = 1121 // PRODUCTION-201611291003-338511768
	ModToolIssueHandledHeader             = 934  // PRODUCTION-201611291003-338511768
	FloodCounterHeader                    = 566  // PRODUCTION-201611291003-338511768
	UpdateFailedHeader                    = 156  // PRODUCTION-201611291003-338511768
	FloorPlanEditorDoorSettingsHeader     = 1664 // PRODUCTION-201611291003-338511768
	FloorPlanEditorBlockedTilesHeader     = 3990 // PRODUCTION-201611291003-338511768
	BuildersClubExpiredHeader             = 1452 // PRODUCTION-201611291003-338511768
	RoomSettingsSavedHeader               = 948  // PRODUCTION-201611291003-338511768
	MessengerInitHeader                   = 1605 // PRODUCTION-201611291003-338511768
	UserClothesHeader                     = 1450 // PRODUCTION-201611291003-338511768
	UserEffectsListHeader                 = 340  // PRODUCTION-201611291003-338511768
	NewUserIdentityHeader                 = 3738 // PRODUCTION-201611291003-338511768
	NewNavigatorEventCategoriesHeader     = 3244 // PRODUCTION-201611291003-338511768
	NewNavigatorCollapsedCategoriesHeader = 1543 // PRODUCTION-201611291003-338511768
	NewNavigatorLiftedRoomsHeader         = 3104 // PRODUCTION-201611291003-338511768
	NewNavigatorSavedSearchesHeader       = 3984 // PRODUCTION-201611291003-338511768
	PostItDataHeader                      = 2202 // PRODUCTION-201611291003-338511768
	ModToolReportReceivedAlertHeader      = 3635 // PRODUCTION-201611291003-338511768
	ModToolIssueResponseAlertHeader       = 3796 // PRODUCTION-201611291003-338511768
	AchievementListHeader                 = 305  // PRODUCTION-201611291003-338511768
	AchievementProgressHeader             = 2107 // PRODUCTION-201611291003-338511768
	AchievementUnlockedHeader             = 806  // PRODUCTION-201611291003-338511768
	ClubGiftsHeader                       = 619  // PRODUCTION-201611291003-338511768
	MachineIDHeader                       = 1488 // PRODUCTION-201611291003-338511768
	PongHeader                            = 10   // PRODUCTION-201611291003-338511768
	ModToolIssueHandlerDimensionsHeader   = 1576 // PRODUCTION-201611291003-338511768

	//Uknown but work
	IsFirstLoginOfDayHeader         = 793  // PRODUCTION-201611291003-338511768 //Quest Engine
	MysteryBoxKeysHeader            = 2833 // PRODUCTION-201611291003-338511768 //Mysterbox
	IgnoredUsersHeader              = 126  // PRODUCTION-201611291003-338511768
	NewNavigatorMetaDataHeader      = 3052 // PRODUCTION-201611291003-338511768
	NewNavigatorSearchResultsHeader = 2690 // PRODUCTION-201611291003-338511768
	MysticBoxStartOpenHeader        = 3201 // PRODUCTION-201611291003-338511768
	MysticBoxCloseHeader            = 596  // PRODUCTION-201611291003-338511768
	MysticBoxPrizeHeader            = 3712 // PRODUCTION-201611291003-338511768
	RentableSpaceInfoHeader         = 3559 // PRODUCTION-201611291003-338511768
	RentableSpaceUnknownHeader      = 2046 // PRODUCTION-201611291003-338511768
	RentableSpaceUnknown2Header     = 1868 // PRODUCTION-201611291003-338511768
	GuildConfirmRemoveMemberHeader  = 1876 // PRODUCTION-201611291003-338511768

	HotelViewBadgeButtonConfigHeader = 2998 // PRODUCTION-201611291003-338511768
	EpicPopupFrameHeader             = 3945 // PRODUCTION-201611291003-338511768
	BaseJumpLoadGameURLHeader        = 2624 // PRODUCTION-201611291003-338511768
	RoomUserTagsHeader               = 1255 // PRODUCTION-201611291003-338511768
	RoomInviteErrorHeader            = 462  // PRODUCTION-201611291003-338511768
	PostItStickyPoleOpenHeader       = 2366 // PRODUCTION-201611291003-338511768
	NewYearResolutionProgressHeader  = 3370 // PRODUCTION-201611291003-338511768
	ClubGiftReceivedHeader           = 659  // PRODUCTION-201611291003-338511768
	ItemStateHeader                  = 2376 // PRODUCTION-201611291003-338511768
	ItemExtraDataHeader              = 2547 // PRODUCTION-201611291003-338511768
	PostUpdateMessageHeader          = 324  // PRODUCTION-201611291003-338511768
	//NotSure Needs Testing
	QuestionInfoHeader             = -1   // PRODUCTION-201611291003-338511768
	TalentTrackEmailVerifiedHeader = 612  // PRODUCTION-201611291003-338511768
	TalentTrackEmailFailedHeader   = 1815 // PRODUCTION-201611291003-338511768
	UnknownAvatarEditorHeader      = 3473 // PRODUCTION-201611291003-338511768

	GuildMembershipRequestedHeader = 1180 // PRODUCTION-201611291003-338511768

	GuildForumsUnreadMessagesCountHeader = 2379 // PRODUCTION-201611291003-338511768
	GuildForumThreadMessagesHeader       = 1862 // PRODUCTION-201611291003-338511768
	GuildForumAddCommentHeader           = 2049 // PRODUCTION-201611291003-338511768
	GuildForumDataHeader                 = 3011 // PRODUCTION-201611291003-338511768
	GuildForumCommentsHeader             = 509  // PRODUCTION-201611291003-338511768
	UnknownGuildForumHeader6             = -1   // PRODUCTION-201611291003-338511768
	UnknownGuildForumHeader7             = -1   // PRODUCTION-201611291003-338511768
	GuildForumThreadsHeader              = 1073 // PRODUCTION-201611291003-338511768
	GuildForumListHeader                 = 3001 // PRODUCTION-201611291003-338511768
	ThreadUpdateMessageHeader            = 2528
	GuideSessionAttachedHeader           = 1591 // PRODUCTION-201611291003-338511768
	GuideSessionDetachedHeader           = 138  // PRODUCTION-201611291003-338511768
	GuideSessionStartedHeader            = 3209 // PRODUCTION-201611291003-338511768
	GuideSessionEndedHeader              = 1456 // PRODUCTION-201611291003-338511768
	GuideSessionErrorHeader              = 673  // PRODUCTION-201611291003-338511768
	GuideSessionMessageHeader            = 841  // PRODUCTION-201611291003-338511768
	GuideSessionRequesterRoomHeader      = 1847 // PRODUCTION-201611291003-338511768
	GuideSessionInvitedToGuideRoomHeader = 219  // PRODUCTION-201611291003-338511768
	GuideSessionPartnerIsTypingHeader    = 1016 // PRODUCTION-201611291003-338511768

	GuideToolsHeader                = 1548 // PRODUCTION-201611291003-338511768
	GuardianNewReportReceivedHeader = 735  // PRODUCTION-201611291003-338511768
	GuardianVotingRequestedHeader   = 143  // PRODUCTION-201611291003-338511768
	GuardianVotingVotesHeader       = 1829 // PRODUCTION-201611291003-338511768
	GuardianVotingResultHeader      = 3276 // PRODUCTION-201611291003-338511768
	GuardianVotingTimeEnded         = 30   // PRODUCTION-201611291003-338511768

	RoomMutedHeader = 2533 // PRODUCTION-201611291003-338511768

	HideDoorbellHeader     = 3783 // PRODUCTION-201611291003-338511768
	RoomQueueStatusMessage = 2208 // PRODUCTION-201611291003-338511768
	RoomUnknown3Header     = 1033 // PRODUCTION-201611291003-338511768

	EffectsListRemoveHeader = 2228 // PRODUCTION-201611291003-338511768

	OldPublicRoomsHeader = 2726 // PRODUCTION-201611291003-338511768
	ItemStateHeader2     = 3431 // PRODUCTION-201611291003-338511768

	HotelWillCloseInMinutesHeader           = 1050 // PRODUCTION-201611291003-338511768
	HotelWillCloseInMinutesAndBackInHeader  = 1350 // PRODUCTION-201611291003-338511768
	HotelClosesAndWillOpenAtHeader          = 2771 // PRODUCTION-201611291003-338511768
	HotelClosedAndOpensHeader               = 3728 // PRODUCTION-201611291003-338511768
	StaffAlertAndOpenHabboWayHeader         = 1683 // PRODUCTION-201611291003-338511768
	StaffAlertWithLinkHeader                = 2030 // PRODUCTION-201611291003-338511768
	StaffAlertWIthLinkAndOpenHabboWayHeader = 1890 // PRODUCTION-201611291003-338511768

	RoomMessagesPostedCountHeader          = 1634 // PRODUCTION-201611291003-338511768
	CantScratchPetNotOldEnoughHeader       = 1130 // PRODUCTION-201611291003-338511768
	PetBoughtNotificationHeader            = 1111 // PRODUCTION-201611291003-338511768
	MessagesForYouHeader                   = 2035 // PRODUCTION-201611291003-338511768
	UnknownStatusHeader                    = 1243 // PRODUCTION-201611291003-338511768
	CloseWebPageHeader                     = 426  // PRODUCTION-201611291003-338511768
	PickMonthlyClubGiftNotificationHeader  = 2188 // PRODUCTION-201611291003-338511768
	RemoveGuildFromRoomHeader              = 3129 // PRODUCTION-201611291003-338511768
	RoomBannedUsersHeader                  = 1869 // PRODUCTION-201611291003-338511768
	OpenRoomCreationWindowHeader           = 2064 // PRODUCTION-201611291003-338511768
	ItemsDataUpdateHeader                  = 1453 // PRODUCTION-201611291003-338511768
	WelcomeGiftHeader                      = 2707 // PRODUCTION-201611291003-338511768
	SimplePollStartHeader                  = 2665 // PRODUCTION-201611291003-338511768
	RoomNoRightsHeader                     = 2392 // PRODUCTION-201611291003-338511768
	GuildEditFailHeader                    = 3988 // PRODUCTION-201611291003-338511768
	MinimailNewMessageHeader               = 1911 // PRODUCTION-201611291003-338511768
	RoomFilterWordsHeader                  = 2937 // PRODUCTION-201611291003-338511768
	VerifyMobileNumberHeader               = 3639 // PRODUCTION-201611291003-338511768
	NewUserGiftHeader                      = 3575 // PRODUCTION-201611291003-338511768
	UpdateUserLookHeader                   = 2429 // PRODUCTION-201611291003-338511768
	RoomUserIgnoredHeader                  = 207  // PRODUCTION-201611291003-338511768
	PetBreedingFailedHeader                = 1625 // PRODUCTION-201611291003-338511768
	RoomUserNameChangedHeader              = 2182 // PRODUCTION-201611291003-338511768
	LoveLockFurniStartHeader               = 3753 // PRODUCTION-201611291003-338511768
	LoveLockFurniFriendConfirmedHeader     = 382  // PRODUCTION-201611291003-338511768
	LoveLockFurniFinishedHeader            = 770  // PRODUCTION-201611291003-338511768
	PetPackageNameValidationHeader         = 546  // PRODUCTION-201611291003-338511768
	GameCenterFeaturedPlayersHeader        = 3097 // PRODUCTION-201611291003-338511768
	HabboMallHeader                        = 1237 // PRODUCTION-201611291003-338511768
	TargetedOfferHeader                    = 119  // PRODUCTION-201611291003-338511768
	LeprechaunStarterBundleHeader          = 2380 // PRODUCTION-201611291003-338511768
	VerifyMobilePhoneWindowHeader          = 2890 // PRODUCTION-201611291003-338511768
	VerifyMobilePhoneCodeWindowHeader      = 800  // PRODUCTION-201611291003-338511768
	VerifyMobilePhoneDoneHeader            = 91   // PRODUCTION-201611291003-338511768
	RoomUserReceivedHandItemHeader         = 354  // PRODUCTION-201611291003-338511768
	MutedWhisperHeader                     = 826  // PRODUCTION-201611291003-338511768
	UnknownHintHeader                      = 1787 // PRODUCTION-201611291003-338511768
	BullyReportClosedHeader                = 2674 // PRODUCTION-201611291003-338511768
	PromoteOwnRoomsListHeader              = 2468 // PRODUCTION-201611291003-338511768
	NotEnoughPointsTypeHeader              = 3914 // PRODUCTION-201611291003-338511768
	WatchAndEarnRewardHeader               = 2125 // PRODUCTION-201611291003-338511768
	NewYearResolutionHeader                = 66   // PRODUCTION-201611291003-338511768
	WelcomeGiftErrorHeader                 = 2293 // PRODUCTION-201611291003-338511768
	RentableItemBuyOutPriceHeader          = 35   // PRODUCTION-201611291003-338511768
	VipTutorialsStartHeader                = 2278 // PRODUCTION-201611291003-338511768
	NewNavigatorCategoryUserCountHeader    = 1455 // PRODUCTION-201611291003-338511768
	CameraRoomThumbnailSavedHeader         = 3595 // PRODUCTION-201611291003-338511768
	RoomEditSettingsErrorHeader            = 1555 // PRODUCTION-201611291003-338511768
	GuildAcceptMemberErrorHeader           = 818  // PRODUCTION-201611291003-338511768
	MostUselessErrorAlertHeader            = 662  // PRODUCTION-201611291003-338511768
	AchievementsConfigurationHeader        = 1689 // PRODUCTION-201611291003-338511768
	PetBreedingResultHeader                = 634  // PRODUCTION-201611291003-338511768
	RoomUserQuestionAnsweredHeader         = -1   // PRODUCTION-201611291003-338511768
	PetBreedingStartHeader                 = 1746 // PRODUCTION-201611291003-338511768
	CustomNotificationHeader               = 909  // PRODUCTION-201611291003-338511768
	UpdateStackHeightTileHeightHeader      = 2816 // PRODUCTION-201611291003-338511768
	HotelViewCustomTimerHeader             = -1   // PRODUCTION-201611291003-338511768
	MarketplaceItemPostedHeader            = 1359 // PRODUCTION-201611291003-338511768
	HabboWayQuizHeader2                    = 2927 // PRODUCTION-201611291003-338511768
	GuildFavoriteRoomUserUpdateHeader      = 3403 // PRODUCTION-201611291003-338511768
	RoomAdErrorHeader                      = 1759 // PRODUCTION-201611291003-338511768
	NewNavigatorSettingsHeader             = 518  // PRODUCTION-201611291003-338511768
	CameraPublishWaitMessageHeader         = 2057 // PRODUCTION-201611291003-338511768
	RoomInviteHeader                       = 3870 // PRODUCTION-201611291003-338511768
	BullyReportRequestHeader               = 3463 // PRODUCTION-201611291003-338511768
	UnknownHelperHeader                    = 77   // PRODUCTION-201611291003-338511768
	HelperRequestDisabledHeader            = 1651 // PRODUCTION-201611291003-338511768
	RoomUnitIdleHeader                     = 1797 // PRODUCTION-201611291003-338511768
	QuestCompletedHeader                   = 949  // PRODUCTION-201611291003-338511768
	UnkownPetPackageHeader                 = 1723 // PRODUCTION-201611291003-338511768
	ForwardToRoomHeader                    = 160  // PRODUCTION-201611291003-338511768
	EffectsListEffectEnableHeader          = 1959 // PRODUCTION-201611291003-338511768
	CompetitionEntrySubmitResultHeader     = 1177 // PRODUCTION-201611291003-338511768
	ExtendClubMessageHeader                = 3964 // PRODUCTION-201611291003-338511768
	HotelViewConcurrentUsersHeader         = 2737 // PRODUCTION-201611291003-338511768
	InventoryAddEffectHeader               = -1   //error 404
	TalentLevelUpdateHeader                = 638  // PRODUCTION-201611291003-338511768
	BullyReportedMessageHeader             = 3285 // PRODUCTION-201611291003-338511768
	UnknownQuestHeader3                    = 1122 // PRODUCTION-201611291003-338511768
	FriendToolbarNotificationHeader        = 3082 // PRODUCTION-201611291003-338511768
	MessengerErrorHeader                   = 896  // PRODUCTION-201611291003-338511768
	CameraPriceHeader                      = 3878 // PRODUCTION-201611291003-338511768
	PetBreedingCompleted                   = 2527 // PRODUCTION-201611291003-338511768
	RoomUserUnbannedHeader                 = 3429 // PRODUCTION-201611291003-338511768
	HotelViewCommunityGoalHeader           = 2525 // PRODUCTION-201611291003-338511768
	UserClassificationHeader               = 966  // PRODUCTION-201611291003-338511768
	CanCreateEventHeader                   = 2599 // PRODUCTION-201611291003-338511768
	UnknownGuild2Header                    = 1459 // PRODUCTION-201611291003-338511768
	YoutubeDisplayListHeader               = 1112 // PRODUCTION-201611291003-338511768
	YoutubeMessageHeader2                  = 1411 // PRODUCTION-201611291003-338511768
	YoutubeMessageHeader3                  = 1554 // PRODUCTION-201611291003-338511768
	RoomCategoryUpdateMessageHeader        = 3896 // PRODUCTION-201611291003-338511768
	QuestsHeader                           = 3625 // PRODUCTION-201611291003-338511768
	GiftReceiverNotFoundHeader             = 1517 // PRODUCTION-201611291003-338511768
	ConvertedForwardToRoomHeader           = 1331 // PRODUCTION-201611291003-338511768
	FavoriteRoomChangedHeader              = 2524 // PRODUCTION-201611291003-338511768
	AlertPurchaseUnavailableHeader         = 3770 // PRODUCTION-201611291003-338511768
	PetBreedingStartFailedHeader           = 2621 // PRODUCTION-201611291003-338511768
	DailyQuestHeader                       = 1878 // PRODUCTION-201611291003-338511768
	HotelViewHideCommunityVoteButtonHeader = 1435 // PRODUCTION-201611291003-338511768
	CatalogSearchResultHeader              = 3388 // PRODUCTION-201611291003-338511768
	FriendFindingRoomHeader                = 1210 // PRODUCTION-201611291003-338511768
	QuestHeader                            = 230  // PRODUCTION-201611291003-338511768
	ModToolSanctionDataHeader              = 2782 // PRODUCTION-201611291003-338511768
	RoomEventMessageHeader                 = 1840

	JukeBoxMySongsHeader           = 2602 // PRODUCTION-201611291003-338511768
	JukeBoxNowPlayingMessageHeader = 469  // PRODUCTION-201611291003-338511768
	JukeBoxPlaylistFullHeader      = 105  // PRODUCTION-201611291003-338511768
	JukeBoxPlayListUpdatedHeader   = 1748 // PRODUCTION-201611291003-338511768
	JukeBoxPlayListAddSongHeader   = 1140 // PRODUCTION-201611291003-338511768
	JukeBoxPlayListHeader          = 34   // PRODUCTION-201611291003-338511768
	JukeBoxTrackDataHeader         = 3365 // PRODUCTION-201611291003-338511768
	JukeBoxTrackCodeHeader         = 1381 // PRODUCTION-201611291003-338511768

	CraftableProductsHeader = 1000 // PRODUCTION-201611291003-338511768
	CraftingRecipeHeader    = 2774 // PRODUCTION-201611291003-338511768
	CraftingResultHeader    = 618  // PRODUCTION-201611291003-338511768
	CraftingHeaderFour      = 2124 // PRODUCTION-201611291003-338511768

	UnknownHeader_100                  = 1553 // PRODUCTION-201611291003-338511768 //PetBReedingResult
	ConnectionErrorHeader              = 1004 // PRODUCTION-201611291003-338511768
	BotForceOpenContextMenuHeader      = 296  // PRODUCTION-201611291003-338511768
	UnknownHeader_1111                 = 1551 // PRODUCTION-201611291003-338511768
	Game2WeeklyLeaderboardHeader       = 2196 // PRODUCTION-201611291003-338511768
	UnknownHeader_1165                 = 904  // PRODUCTION-201611291003-338511768
	EffectsListAddHeader               = 2867 // PRODUCTION-201611291003-338511768
	UnknownHeader_1188                 = 1437 // PRODUCTION-201611291003-338511768
	SubmitCompetitionRoomHeader        = 3841 // PRODUCTION-201611291003-338511768
	GameAchievementsListHeader         = 2265 // PRODUCTION-201611291003-338511768
	OtherTradingDisabledHeader         = 1254 // PRODUCTION-201611291003-338511768
	BaseJumpUnloadGameHeader           = 1715 // PRODUCTION-201611291003-338511768
	UnknownHeader_137                  = 2897 // PRODUCTION-201611291003-338511768
	GameCenterAccountInfoHeader        = 2893 // PRODUCTION-201611291003-338511768
	UnknowHeader_1390                  = 2270 // PRODUCTION-201611291003-338511768
	BaseJumpLoadGameHeader             = 3654 // PRODUCTION-201611291003-338511768
	UnknowHeader_1427                  = 3319 // PRODUCTION-201611291003-338511768
	AdventCalendarDataHeader           = 2531 // PRODUCTION-201611291003-338511768
	UnknownHeader_152                  = 3954 // PRODUCTION-201611291003-338511768
	UnknownHeader_1577                 = 2641 // PRODUCTION-201611291003-338511768
	NewYearResolutionCompletedHeader   = 740  // PRODUCTION-201611291003-338511768
	UnknownHeader_1741                 = 2246 // PRODUCTION-201611291003-338511768
	UnknownHeader_1744                 = 2873 // PRODUCTION-201611291003-338511768
	AdventCalendarProductHeader        = 2551 // PRODUCTION-201611291003-338511768
	ModToolSanctionInfoHeader          = 2221 // PRODUCTION-201611291003-338511768
	UnknownHeader_1965                 = 3292 // PRODUCTION-201611291003-338511768
	GuideSessionPartnerIsPlayingHeader = 448  // PRODUCTION-201611291003-338511768
	BaseJumpLeaveQueueHeader           = 1477 // PRODUCTION-201611291003-338511768
	Game2WeeklySmallLeaderboardHeader  = 3512 // PRODUCTION-201611291003-338511768
	GameCenterGameListHeader           = 222  // PRODUCTION-201611291003-338511768
	RoomUsersGuildBadgesHeader         = 2402 // PRODUCTION-201611291003-338511768
	UnknownHeader_2563                 = 1774 // PRODUCTION-201611291003-338511768
	UnknownHeader_2601                 = 1663 // PRODUCTION-201611291003-338511768
	UnknownHeader_2621                 = 1927 // PRODUCTION-201611291003-338511768
	UnknownHeader_2698                 = 563  // PRODUCTION-201611291003-338511768
	//    2699
	//    2748
	//    2773
	//    2964
	//    3020
	//    3024
	//    3046
	//    3124
	//    3179
	//    3189
	//    328
	//    3291
	//    3334
	HabboWayQuizHeader1 = 3379
	//    3391
	//    3427
	//    347
	//    3509
	//    3519
	//    3581
	//    3684
	YouTradingDisabledHeader = 3058 // PRODUCTION-201611291003-338511768
	//    3707
	//    3745
	//    3759
	//    3782
	RoomFloorThicknessUpdatedHeader = 3786
	//    3822
	CameraPurchaseSuccesfullHeader = 2783 // PRODUCTION-201611291003-338511768
	CameraCompetitionStatusHeader  = 133  // PRODUCTION-201611291003-338511768
	//    3986
	//    467
	//    549
	CameraURLHeader = 3696 // PRODUCTION-201611291003-338511768
	//    608
	//    624
	//    675
	HotelViewCatalogPageExpiringHeader = 690
	//    749
	//    812
	//    843
	//    947
	//    982

	SimplePollAnswerHeader = 2589

	PingHeader                  = 3928
	TradingWaitingConfirmHeader = 2720
	BaseJumpJoinQueueHeader     = 2260
	ClubCenterDataHeader        = 3277

	SimplePollAnswersHeader               = 1066
	UnknownFurniModelHeader               = 1501
	UnknownAdManagerHeader                = 1808
	WiredOpenHeader                       = 1830
	UnknownCatalogPageOfferHeader         = 1889
	NuxAlertHeader                        = 2023
	HotelViewExpiringCatalogPageCommposer = 2515
	UnknownHabboWayQuizHeader             = 2772
	PetLevelUpdatedHeader                 = 2824
	QuestExpiredHeader                    = 3027
	UnknownTradeHeader                    = 3128
	UnknownMessengerErrorHeader           = 3359
	UnknownHeader8                        = 3441
	RemoveRoomEventHeader                 = 3479
	UnknownCompetitionHeader              = 3506
	UnknownRoomViewerHeader               = 3523
	ErrorLoginHeader                      = 4000
	HotelViewNextLTDAvailableHeader       = 44
	HotelViewSecondsUntilHeader           = 3926
	UnknownRoomDesktopHeader              = 69
	UnknownGuildHeader3                   = 876

	GameCenterGameHeader = 3805

	SnowStormGameStartedHeader            = 5000
	SnowStormQuePositionHeader            = 5001
	SnowStormStartBlockTickerHeader       = 5002
	SnowStormStartLobbyCounterHeader      = 5003
	SnowStormUnusedAlertGenericHeader     = 5004
	SnowStormLongDataHeader               = 5005
	SnowStormGameEndedHeader              = 5006
	SnowStormQuePlayerAddedHeader         = 5008
	SnowStormPlayAgainHeader              = 5009
	SnowStormGamesLeftHeader              = 5010
	SnowStormQuePlayerRemovedHeader       = 5011
	SnowStormGamesInformationHeader       = 5012
	SnowStormLongData2Header              = 5013
	UNUSED_SNOWSTORM_5014                 = 5014
	SnowStormGameStatusHeader             = 5015
	SnowStormFullGameStatusHeader         = 5016
	SnowStormOnStageStartHeader           = 5017
	SnowStormIntializeGameArenaViewHeader = 5018
	SnowStormRejoinPreviousRoomHeader     = 5019
	UNKNOWN_SNOWSTORM_5020                = 5020
	SnowStormLevelDataHeader              = 5021
	SnowStormOnGameEndingHeader           = 5022
	SnowStormUserChatMessageHeader        = 5023
	SnowStormOnStageRunningHeader         = 5024
	SnowStormOnStageEndingHeader          = 5025
	SnowStormIntializedPlayersHeader      = 5026
	SnowStormOnPlayerExitedArenaHeader    = 5027
	SnowStormGenericErrorHeader           = 5028
	SnowStormUserRematchedHeader          = 5029
)
