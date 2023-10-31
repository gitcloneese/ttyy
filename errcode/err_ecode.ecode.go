// Code generated by protoc-gen-ecode v0.7, DO NOT EDIT.
// source: err_ecode.proto

package v1

import (
	"gitlab.com/firerocksg/xy3-kratos/pkg/ecode"
)

// to suppressed 'imported but not used warning'
var _ ecode.Codes

// ErrCode ecode
var (
	DeserializeFailed                      = ecode.New(1)
	IllegalParams                          = ecode.New(2)
	RedisError                             = ecode.New(3)
	RedisLockTimeout                       = ecode.New(6)
	RedisLockFailed                        = ecode.New(7)
	Offlining                              = ecode.New(8)
	UserDisable                            = ecode.New(9)
	ForbidChatChannelErr                   = ecode.New(10)
	ForbidChat                             = ecode.New(11)
	AccountDisable                         = ecode.New(12)
	NoHandler                              = ecode.New(101)
	AuthorizationFailed                    = ecode.New(104)
	NotFindPlayer                          = ecode.New(105)
	SessionNotExist                        = ecode.New(106)
	PushFailed                             = ecode.New(107)
	ServiceHostNotMatched                  = ecode.New(108)
	NoActivityHost                         = ecode.New(201)
	PlayerAlreadyAllotted                  = ecode.New(202)
	CoordinatorError                       = ecode.New(301)
	PublisherError                         = ecode.New(302)
	CreateAccountError                     = ecode.New(401)
	SetTokenError                          = ecode.New(402)
	AccountNotExists                       = ecode.New(403)
	GetAccountError                        = ecode.New(404)
	PasswordMismatch                       = ecode.New(405)
	CreateVerifyCodeFailed                 = ecode.New(406)
	BeyondVerifyCodeMaxSendCount           = ecode.New(407)
	SaveVerifyCodeFailed                   = ecode.New(408)
	SendVerifyCodeFailed                   = ecode.New(409)
	AccountAlreadyExists                   = ecode.New(410)
	VerifyCodeError                        = ecode.New(411)
	RemoveVerifyCodeError                  = ecode.New(412)
	ThirdVerifyFailed                      = ecode.New(413)
	ThirdLoginFailed                       = ecode.New(414)
	LogoutFailed                           = ecode.New(415)
	VerifyAccessTokenFail                  = ecode.New(416)
	GenerateUnionIdError                   = ecode.New(417)
	LoginTypeInvalid                       = ecode.New(418)
	UnbindNotAllowed                       = ecode.New(419)
	PasswordCreationError                  = ecode.New(420)
	EmailCreationError                     = ecode.New(421)
	DeviceUidEmpty                         = ecode.New(422)
	GetServerListError                     = ecode.New(423)
	ConfigNotExist                         = ecode.New(601)
	DataException                          = ecode.New(602)
	SystemLocked                           = ecode.New(666)
	CrossSceneNil                          = ecode.New(670)
	CrossUserNotFind                       = ecode.New(671)
	CrossUnmarshalFailed                   = ecode.New(672)
	CrossMarshalFailed                     = ecode.New(673)
	GoodsNotExist                          = ecode.New(701)
	GoodsUseLvTooHigh                      = ecode.New(702)
	GoodsNotEnough                         = ecode.New(703)
	GoodsNotUse                            = ecode.New(704)
	GoodsNotConfig                         = ecode.New(705)
	GoodsInvalidArgs                       = ecode.New(706)
	GoodsNotSell                           = ecode.New(707)
	GoodsAwardFail                         = ecode.New(708)
	TokenInvalid                           = ecode.New(801)
	NoUser                                 = ecode.New(900)
	PlayerExistCreateFailed                = ecode.New(901)
	NoExistSelectFailed                    = ecode.New(902)
	DBErrSelectFailed                      = ecode.New(903)
	EnterGameFailed                        = ecode.New(904)
	DBErrGetFailed                         = ecode.New(905)
	UnknownError                           = ecode.New(10000)
	FormatActivityErr                      = ecode.New(10001)
	FormatChatErr                          = ecode.New(10002)
	FormatFriendErr                        = ecode.New(10003)
	FormatMailErr                          = ecode.New(10004)
	FormatTaskErr                          = ecode.New(10005)
	GuildNoUser                            = ecode.New(12350)
	GuildCreateFailed                      = ecode.New(12351)
	GuildRedisErr                          = ecode.New(12352)
	GuildExitForbid                        = ecode.New(12353)
	GuildExitErr                           = ecode.New(12354)
	GuildNotImpeach                        = ecode.New(12355)
	GuildNoChairAndVice                    = ecode.New(12356)
	GuildNoSameGuild                       = ecode.New(12357)
	GuildViceTooMore                       = ecode.New(12358)
	GuildNoticeIll                         = ecode.New(12360)
	GuildArchMaxLv                         = ecode.New(12361)
	GuildNoType                            = ecode.New(12362)
	GuildNoFund                            = ecode.New(12363)
	GuildMailErr                           = ecode.New(12364)
	GuildTextIll                           = ecode.New(12365)
	GuildNoGuild1                          = ecode.New(12366)
	GuildInGuil1                           = ecode.New(12367)
	GuildCooling1                          = ecode.New(12368)
	GuildNoChairman1                       = ecode.New(12369)
	GuildNoCfg1                            = ecode.New(12370)
	GuildHuntDiscardNoBattle               = ecode.New(12371)
	GuildArchUpMainLvNot                   = ecode.New(12372)
	GuildExitTimeErr                       = ecode.New(12373)
	MailNewPartFailed                      = ecode.New(12550)
	MailAlreadGeted                        = ecode.New(12551)
	MailNoReward                           = ecode.New(12552)
	MailNotExist                           = ecode.New(12553)
	MailUnread                             = ecode.New(12554)
	MailUnGet                              = ecode.New(12555)
	MailRewardFailed                       = ecode.New(12556)
	FriendMax                              = ecode.New(12701)
	FriendRequestNotList                   = ecode.New(12702)
	TargetFriendMax                        = ecode.New(12703)
	PlayerInfoNotExist                     = ecode.New(12704)
	FriendNotExist                         = ecode.New(12705)
	FriendPointNotExist                    = ecode.New(12706)
	PlayerLevelNotExist                    = ecode.New(12707)
	PlayerNotExist                         = ecode.New(12708)
	InBlackList                            = ecode.New(12709)
	GetPointCountMax                       = ecode.New(12710)
	GivePointCountMax                      = ecode.New(12711)
	SeachTooManyFriend                     = ecode.New(12712)
	LeaseHeroFightCountMax                 = ecode.New(12801)
	LeaseHeroRequestNotFriend              = ecode.New(12802)
	LeaseHeroSelfMax                       = ecode.New(12803)
	LeaseHeroTargetEmpty                   = ecode.New(12804)
	LeaseHeroTargetLease                   = ecode.New(12805)
	LeaseHeroTargetCD                      = ecode.New(12806)
	LeaseHeroNotSelf                       = ecode.New(12807)
	LeaseHeroRequestMax                    = ecode.New(12808)
	LeaseHeroNotOwner                      = ecode.New(12809)
	LeaseHeroTaskEmpty                     = ecode.New(12810)
	ChatImServerErr                        = ecode.New(12650)
	ChatUserNil                            = ecode.New(12651)
	ChatUserImid0                          = ecode.New(12652)
	ChatRoomImidNo                         = ecode.New(12653)
	ChatSendSignErr                        = ecode.New(12654)
	ChatUserErr                            = ecode.New(13050)
	ChatNoRole                             = ecode.New(13051)
	ChatChannelWorldErr                    = ecode.New(13052)
	ChatChannelPeerErr                     = ecode.New(13053)
	RedPacketRedisErr                      = ecode.New(13150)
	RedPacketNoCfg                         = ecode.New(13151)
	RedPacketFull                          = ecode.New(13152)
	RedPacketNoUser                        = ecode.New(13153)
	RedPacketNotInGuild                    = ecode.New(13154)
	RedPacketNotExist                      = ecode.New(13155)
	RedPacketOverTime                      = ecode.New(13156)
	RedPacketGetFinish                     = ecode.New(13157)
	RedPacketGeted                         = ecode.New(13158)
	RedPacketGiveMax                       = ecode.New(13159)
	RedPacketGetMax                        = ecode.New(13160)
	DevilConquerNoConfig                   = ecode.New(13250)
	DevilConquerRedisErr                   = ecode.New(13251)
	DevilConquerNoUser                     = ecode.New(13252)
	DevilConquerDayNot                     = ecode.New(13253)
	DevilConquerNoInTime                   = ecode.New(13254)
	DevilConquerNoRobbed                   = ecode.New(13255)
	AuctionNoConfig                        = ecode.New(13650)
	AuctionStageErr                        = ecode.New(13651)
	AuctionNoGoods                         = ecode.New(13652)
	AuctionHasFixedPrice                   = ecode.New(13653)
	AuctionRepeatBid                       = ecode.New(13654)
	AuctionBidPriceErr                     = ecode.New(13655)
	AuctionRedisErr                        = ecode.New(13656)
	AuctionNoTakein                        = ecode.New(13657)
	AuctionBonusGeted                      = ecode.New(13658)
	AuctionNowNoGoods                      = ecode.New(13659)
	LiveDRewardUnReach                     = ecode.New(30201)
	LiveDRewardGeted                       = ecode.New(30202)
	LiveDTBoxTooFew                        = ecode.New(30203)
	LiveDNoConfig                          = ecode.New(30204)
	LiveDRewardFailed                      = ecode.New(30205)
	GoodsAddEcoMax                         = ecode.New(30301)
	NewHandbookSetFailed                   = ecode.New(30401)
	VIPLvNoEnough                          = ecode.New(30402)
	PvpFightingRankErr                     = ecode.New(40101)
	PvpFightingRankDataErr                 = ecode.New(40102)
	PvpFightingChangeRankErr               = ecode.New(40103)
	PvpFightingReplaceErr                  = ecode.New(40104)
	PvpFightingReportNotExist              = ecode.New(40105)
	BattleInfoNotExist                     = ecode.New(40401)
	LineupInfoNotExist                     = ecode.New(40402)
	ZTPvpLevelRewardLimit                  = ecode.New(50000)
	ZTPvpLevelTaskNotExist                 = ecode.New(50001)
	ZTPvpSeasonNotJoin                     = ecode.New(50002)
	ZTPvpLevelNotEnough                    = ecode.New(50003)
	ZTPvpLevelTaskReceived                 = ecode.New(50004)
	ZTPvpTaskNotExist                      = ecode.New(50005)
	ZTPvpTaskNotEnough                     = ecode.New(50006)
	ZTPvpTaskReceived                      = ecode.New(50007)
	ZTPvpCheckBattleErr                    = ecode.New(50008)
	ZTPvpPlayerUpdateErr                   = ecode.New(50009)
	ZTPvpPlayerUpdateRankErr               = ecode.New(50010)
	ZTPvpGetPlayerRankErr                  = ecode.New(50011)
	ZTPvpUpdateLevelLimitErr               = ecode.New(50012)
	ZTPvpRecordUpdateErr                   = ecode.New(50013)
	TTPvpLevelRewardLimit                  = ecode.New(51000)
	TTPvpLevelTaskNotExist                 = ecode.New(51001)
	TTPvpSeasonNotJoin                     = ecode.New(51002)
	TTPvpLevelNotEnough                    = ecode.New(51003)
	TTPvpLevelTaskReceived                 = ecode.New(51004)
	TTPvpTaskNotExist                      = ecode.New(51005)
	TTPvpTaskNotEnough                     = ecode.New(51006)
	TTPvpTaskReceived                      = ecode.New(51007)
	TTPvpCheckBattleErr                    = ecode.New(51008)
	TTPvpPlayerUpdateErr                   = ecode.New(51009)
	TTPvpPlayerUpdateRankErr               = ecode.New(51010)
	TTPvpGetPlayerRankErr                  = ecode.New(51011)
	TTPvpUpdateLevelLimitErr               = ecode.New(51012)
	TTPvpRecordUpdateErr                   = ecode.New(51013)
	LineupGroupNotExist                    = ecode.New(10101)
	LineupIndexNotExist                    = ecode.New(10102)
	LineupGridNotExist                     = ecode.New(10104)
	LineupHeroNotExist                     = ecode.New(10105)
	LineupHeroAlreadyExist                 = ecode.New(10106)
	LineupGridNotUnlock                    = ecode.New(10107)
	LineupNotSave                          = ecode.New(10108)
	LineupDisabled                         = ecode.New(10109)
	LineupHeroNameAlreadyExist             = ecode.New(10110)
	LineupLeaseHeroAlreadyExist            = ecode.New(10111)
	LineupOtherLeaseAlreadyExist           = ecode.New(10112)
	LineupNoGroup                          = ecode.New(10113)
	LineupNoID                             = ecode.New(10114)
	EquipNotExist                          = ecode.New(10121)
	EquipQualityMax                        = ecode.New(10122)
	EquipLevelMax                          = ecode.New(10123)
	EquipQualityNotEnough                  = ecode.New(10124)
	EquipRefineLevelMax                    = ecode.New(10125)
	EquipSealLevelMax                      = ecode.New(10126)
	LineupTeamsIllegal                     = ecode.New(10141)
	LineupTeamsNotExist                    = ecode.New(10142)
	LineupTeamsMax                         = ecode.New(10143)
	LineupTeamsNameExist                   = ecode.New(10144)
	NameIsEmpty                            = ecode.New(10150)
	NameIsIllegal                          = ecode.New(10151)
	LineupNameLen                          = ecode.New(10152)
	LineupTeamNameLen                      = ecode.New(10153)
	HeroUpStarNotEnough                    = ecode.New(10301)
	HeroUpStarMax                          = ecode.New(10302)
	HeroUpQualityNotEnough                 = ecode.New(10303)
	HeroUpQualityMax                       = ecode.New(10304)
	HeroUpAwakenNotEnough                  = ecode.New(10305)
	HeroUpAwakenMax                        = ecode.New(10306)
	HeroNeedStarMax                        = ecode.New(10307)
	HeroNeedQualityMax                     = ecode.New(10308)
	HeroUpStarNotConfig                    = ecode.New(10309)
	HeroUpQualityNotConfig                 = ecode.New(10310)
	HeroUpAwakenNotConfig                  = ecode.New(10311)
	HeroActiveExist                        = ecode.New(10312)
	HeroActiveNotConfig                    = ecode.New(10313)
	HeroActiveNotEnough                    = ecode.New(10314)
	HeroLevelNotEnough                     = ecode.New(10315)
	HeroLevelMax                           = ecode.New(10316)
	HeroNotExist                           = ecode.New(10317)
	LeaderSkillNotExist                    = ecode.New(10401)
	LeaderSkillReplace                     = ecode.New(10402)
	LeaderTalentGoodsNotEnough             = ecode.New(10403)
	LeaderTalentMax                        = ecode.New(10404)
	LeaderFragmentNotEnough                = ecode.New(10405)
	LeaderFragmentUpNotEnough              = ecode.New(10406)
	LeaderForgeGoodsNotEnough              = ecode.New(10407)
	LeaderForgeNotExistType                = ecode.New(10408)
	LeaderForgeCountMax                    = ecode.New(10409)
	LeaderSkilLocked                       = ecode.New(10410)
	LeaderStarNotEnough                    = ecode.New(10411)
	LeaderUpQualityGoodsNotEnough          = ecode.New(10412)
	LeaderLevelNotEnough                   = ecode.New(10413)
	TouhouClickNotOpen                     = ecode.New(10501)
	TouhouClickNotEnough                   = ecode.New(10502)
	TouhouClickMax                         = ecode.New(10503)
	TouhouClickNotConfig                   = ecode.New(10504)
	StoryBattleIDNotExist                  = ecode.New(10701)
	StoryBattleNotOpen                     = ecode.New(10702)
	StoryBattleAlreadyPass                 = ecode.New(10703)
	StoryNoFighter                         = ecode.New(10704)
	StoryBattleNotExist                    = ecode.New(10705)
	StorySettleIDNotMatch                  = ecode.New(10706)
	StoryVitNotEnough                      = ecode.New(10710)
	StoryExploreLocked                     = ecode.New(10714)
	StoryDropinLocked                      = ecode.New(10715)
	StoryDropinMax                         = ecode.New(10716)
	StoryDropinCoolDown                    = ecode.New(10717)
	StoryExploreNotExist                   = ecode.New(10719)
	StoryExploreNotZero                    = ecode.New(10720)
	StoryFastPlaceLocked                   = ecode.New(10721)
	CoinNotEnough                          = ecode.New(10801)
	PurchaseTimeNotEnough                  = ecode.New(10802)
	GoodsNotPass                           = ecode.New(10803)
	GoodsSellOut                           = ecode.New(10804)
	RechargeNotExist                       = ecode.New(10805)
	GoodsAlreadyReceived                   = ecode.New(10806)
	GroceryConfigrefresh                   = ecode.New(10807)
	PackageConfigrefresh                   = ecode.New(10808)
	VipGiftNotExist                        = ecode.New(10821)
	VipGiftLevelNotEnough                  = ecode.New(10822)
	VipGiftNotGet                          = ecode.New(10823)
	VipCfgNotExist                         = ecode.New(10824)
	GroceryInfoNil                         = ecode.New(10825)
	GroceryCostCoinErr                     = ecode.New(10826)
	BableLocked                            = ecode.New(10901)
	BableNotOpen                           = ecode.New(10902)
	BableBattleCountNotEnough              = ecode.New(10903)
	BableLineupError                       = ecode.New(10904)
	BableLineupEmpty                       = ecode.New(10905)
	BableMaxFloor                          = ecode.New(10906)
	BableFloorNotExist                     = ecode.New(10907)
	BableBigRewardNotExist                 = ecode.New(10908)
	BableBigRewardReceived                 = ecode.New(10909)
	BableAssignHeroNotExist                = ecode.New(10910)
	BableAssignEquipNotExist               = ecode.New(10911)
	BableLeaderSkillNotExist               = ecode.New(10912)
	BableLineupNotMatch                    = ecode.New(10913)
	BableRankRewardReceived                = ecode.New(10914)
	BableRankRewardNotGet                  = ecode.New(10915)
	BableReplayNotExist                    = ecode.New(10916)
	BableAfkNotCheckout                    = ecode.New(10917)
	BableAfkLocked                         = ecode.New(10918)
	BableFastAfkMax                        = ecode.New(10919)
	BableFastAfkTicketNotEnough            = ecode.New(10920)
	LingBaoNotConfig                       = ecode.New(11010)
	LingBaoActiveAlready                   = ecode.New(11011)
	LingBaoActiveNotEnough                 = ecode.New(11012)
	LingBaoUpStarNotEnough                 = ecode.New(11020)
	LingBaoUpStarMax                       = ecode.New(11021)
	LingBaoUpStarNotConfig                 = ecode.New(11022)
	LingBaoUpAdvanceStarLow                = ecode.New(11030)
	LingBaoUpAdvanceNotConfig              = ecode.New(11031)
	LingBaoUpAdvanceNotEnough              = ecode.New(11032)
	LingBaoUpAdvanceMax                    = ecode.New(11033)
	LingBaoNoActive                        = ecode.New(11034)
	LingBaoUsingNumErr                     = ecode.New(11041)
	LingBaoUsingRepet                      = ecode.New(11042)
	FairyWareNotExist                      = ecode.New(11100)
	FairyWareNotConfig                     = ecode.New(11101)
	FairyWareTypeError                     = ecode.New(11102)
	FairyWareLevelMax                      = ecode.New(11103)
	FairyWareAlreadyLock                   = ecode.New(11104)
	FairyWareAlreadyEmploy                 = ecode.New(11105)
	RecruitBaseDayMax                      = ecode.New(11201)
	RecruitBaseGoodsNotEnough              = ecode.New(11202)
	RecruitWishNotEnough                   = ecode.New(11203)
	RecruitWishNotReset                    = ecode.New(11204)
	RecruitFriendDayMax                    = ecode.New(11205)
	RecruitFriendCoinNotEnough             = ecode.New(11206)
	RecruitGoodsDayMax                     = ecode.New(11207)
	RecruitGoodsGoodsNotEnough             = ecode.New(11208)
	RecruitGoodsWishEmpty                  = ecode.New(11209)
	RecruitRaceDayMax                      = ecode.New(11210)
	RecruitRaceNotEnough                   = ecode.New(11211)
	RecruitRaceOpened                      = ecode.New(11212)
	RecruitChangeRaceNotEnough             = ecode.New(11213)
	RecruitWishStateError                  = ecode.New(11214)
	RecruitWishLocked                      = ecode.New(11215)
	LinggenNotConfig                       = ecode.New(11301)
	LinggenHeroNotInLineup                 = ecode.New(11302)
	LinggenDotNotEnough                    = ecode.New(11303)
	LinggenReqLinggen                      = ecode.New(11304)
	LinggenAlreadyActive                   = ecode.New(11305)
	LinggenResetNoAlloc                    = ecode.New(11306)
	LinggenResetNotEnough                  = ecode.New(11307)
	LinggenNeedOrder                       = ecode.New(11308)
	LinggenDotUsedNotEnough                = ecode.New(11309)
	LinggenLevelError                      = ecode.New(11310)
	LinggenHeroTypeErr                     = ecode.New(11311)
	PvpFightingTargetNotExist              = ecode.New(11401)
	PvpFightingChallengeNotEnough          = ecode.New(11402)
	PvpFightingRefreshNotEnough            = ecode.New(11403)
	PvpFightingAddChallengeNotEnough       = ecode.New(11404)
	PvpFightingNotChallenge                = ecode.New(11405)
	PvpFightingChallengeNotList            = ecode.New(11406)
	PvpFightingPromoteLevelNotEnough       = ecode.New(11407)
	PvpFightingPromotePowerNotEnough       = ecode.New(11408)
	PvpFightingPromoteRankNotEnough        = ecode.New(11409)
	PvpFightingRankRewardExchange          = ecode.New(11410)
	PvpFightingRankRewardNotExchange       = ecode.New(11411)
	PvpFightingRankRewardNotExist          = ecode.New(11412)
	PvpFightingRankRewardGoodsNotEnough    = ecode.New(11413)
	PvpFightingPromoteMax                  = ecode.New(11414)
	PvpFightingLocked                      = ecode.New(11415)
	PvpFightingAddChallengeTicketNotEnough = ecode.New(11416)
	PvpFightingLineupLocked                = ecode.New(11417)
	EWNotHave                              = ecode.New(11500)
	EWNotEnough                            = ecode.New(11501)
	EWNotConfig                            = ecode.New(11502)
	EWAlreadMaked                          = ecode.New(11503)
	EWMaxStar                              = ecode.New(11504)
	EWMaxAwaken                            = ecode.New(11505)
	EWNoRune                               = ecode.New(11506)
	EWAlreadRune                           = ecode.New(11507)
	EWRunePlaceUnlock                      = ecode.New(11508)
	EWStageErrLow                          = ecode.New(11509)
	EWAlreadUnlock                         = ecode.New(11510)
	EWBarMax                               = ecode.New(11511)
	EWNoNeedConfirm                        = ecode.New(11512)
	EWHoleNoAwaken                         = ecode.New(11513)
	EWHoleMax                              = ecode.New(11514)
	EWNoHero                               = ecode.New(11515)
	EWAlreadLocked                         = ecode.New(11516)
	EWRandBarFailed                        = ecode.New(11517)
	EWNeedConfirm                          = ecode.New(11518)
	EWRuneNoComp                           = ecode.New(11519)
	EWRuneCompNotEn                        = ecode.New(11520)
	EWStageErrSame                         = ecode.New(11521)
	EWStageErr                             = ecode.New(11522)
	EWRandBarNeedConfirm                   = ecode.New(11523)
	EWAdvanceBarNeedConfirm                = ecode.New(11524)
	EWRandBarNoNeedConfirm                 = ecode.New(11525)
	EWHoleRoleLvLow                        = ecode.New(11526)
	EWAddNewFailed                         = ecode.New(11527)
	PengLaiTypeNotExist                    = ecode.New(11601)
	PengLaiIndexNotExist                   = ecode.New(11602)
	PengLaiIndexNotUnlock                  = ecode.New(11603)
	PengLaiIndexIsUnlock                   = ecode.New(11604)
	PengLaiHeroTypeNotAccord               = ecode.New(11605)
	PengLaiSysIsNotOpen                    = ecode.New(11606)
	ZhenFaisUnlock                         = ecode.New(11621)
	ZhenFaNotUnlock                        = ecode.New(11622)
	ZhenFaSystemNotUnlock                  = ecode.New(11623)
	ZhenFaMax                              = ecode.New(11624)
	ConstellationActivePropNotEnough       = ecode.New(11651)
	ConstellationActiveMaxLv               = ecode.New(11652)
	ConstellationAwakenPropNotEnough       = ecode.New(11653)
	ConstellationAwakenPointNotEnough      = ecode.New(11654)
	ConstellationAwakenMaxLv               = ecode.New(11655)
	ConstellationResetTimeLimit            = ecode.New(11656)
	ConstellationComposePropNotEnough      = ecode.New(11657)
	ConstellationCacluCost                 = ecode.New(11658)
	RoleNameIllegal                        = ecode.New(12001)
	RoleNameRepeat                         = ecode.New(12002)
	RoleNameLen                            = ecode.New(12003)
	RoleNameEmpty                          = ecode.New(12004)
	RoleHeadIconNoCfg                      = ecode.New(12005)
	RoleHeadIconNoActive                   = ecode.New(12006)
	SkinNoCfg                              = ecode.New(12007)
	SkinAlreadyOverTime                    = ecode.New(12008)
	SkinNoActive                           = ecode.New(12009)
	RoleNameNoCost                         = ecode.New(12010)
	RoleNameCharIllegal                    = ecode.New(12011)
	GamesNotOpen                           = ecode.New(12101)
	StarPalaceIDNotExist                   = ecode.New(12102)
	GamesEnd                               = ecode.New(12103)
	GuildRPCErr                            = ecode.New(12300)
	GuildInGuild                           = ecode.New(12301)
	GuildCreateNoEnough                    = ecode.New(12302)
	GuildCreateNoVip                       = ecode.New(12303)
	GuildNoCfg                             = ecode.New(12304)
	GuildWordIll                           = ecode.New(12305)
	GuildNotExist                          = ecode.New(12306)
	GuildApplyCool                         = ecode.New(12307)
	GuildNoChairman                        = ecode.New(12308)
	GuildNameNoEnough                      = ecode.New(12309)
	GuildCooling                           = ecode.New(12310)
	GuildNoGuild                           = ecode.New(12311)
	GuildHallMaxLv                         = ecode.New(12312)
	GuildHallFundNoEnough                  = ecode.New(12313)
	GuildHallRewardAlreadGet               = ecode.New(12314)
	GuildHallLvErr                         = ecode.New(12315)
	GuildHallLvRewardFailed                = ecode.New(12316)
	GuildNameLenErr                        = ecode.New(12317)
	GuildNoticeTooLong                     = ecode.New(12318)
	GuildFull                              = ecode.New(12319)
	GuildNoMoney                           = ecode.New(12320)
	GuildMaxDonateNum                      = ecode.New(12321)
	GuildRewardFailed                      = ecode.New(12322)
	GuildLowHuntNoCnt                      = ecode.New(12323)
	GuildHighHuntNoExist                   = ecode.New(12324)
	GuildHighHuntMax                       = ecode.New(12325)
	GuildBattleCheckFailed                 = ecode.New(12326)
	GuildSweepNoFirst                      = ecode.New(12327)
	GuildSweepHighErr                      = ecode.New(12328)
	GuildSkillPavilionLvLow                = ecode.New(12329)
	GuildSkillLeaderLvLow                  = ecode.New(12330)
	GuildSkillNeedOtherLv                  = ecode.New(12331)
	GuildSkillLvMax                        = ecode.New(12332)
	GuildSkillNoProp                       = ecode.New(12333)
	GuildHasEmptyChar                      = ecode.New(12335)
	GuildHasEnSymbol                       = ecode.New(12336)
	GuildTooLong                           = ecode.New(12337)
	GuildTooShort                          = ecode.New(12338)
	GuildNameRepeat                        = ecode.New(12339)
	GuildRedpacketRewardErr                = ecode.New(12340)
	GuildRedpacketHandSlow                 = ecode.New(12341)
	GuildFeatsNotEnough                    = ecode.New(12342)
	GuildMemberInfoNotExist                = ecode.New(12343)
	TaskRewardFailed                       = ecode.New(12501)
	TaskActiveRewardFailed                 = ecode.New(12502)
	TaskConnectionFailed                   = ecode.New(12503)
	BannerNo                               = ecode.New(12605)
	BannerNoCfg                            = ecode.New(12606)
	BannerNoEnough                         = ecode.New(12607)
	BannerFull                             = ecode.New(12608)
	BannerForgeCountExceeded               = ecode.New(12609)
	BannerGradeLow                         = ecode.New(12610)
	BannerLvStarMax                        = ecode.New(12611)
	BannerNoSubs                           = ecode.New(12612)
	BannerSubsMutex                        = ecode.New(12613)
	BannerSubsMax                          = ecode.New(12614)
	BannerNoRune                           = ecode.New(12615)
	BannerRuneTooLittle                    = ecode.New(12616)
	BannerMakeNoLevel                      = ecode.New(12617)
	BannerResolveReward                    = ecode.New(12618)
	BannerResolveNo                        = ecode.New(12619)
	BannerNoBannerRune                     = ecode.New(12620)
	BannerHoleLock                         = ecode.New(12621)
	BannerRuneMaxLv                        = ecode.New(12622)
	AchievementDataErr                     = ecode.New(12901)
	AchievementNotReach                    = ecode.New(12902)
	AchievementTypeErr                     = ecode.New(12903)
	AchievementPrecondition                = ecode.New(12904)
	AchievementNotExist                    = ecode.New(12905)
	ChatRPCErr                             = ecode.New(13001)
	ChatNoCfg                              = ecode.New(13002)
	ChatMarshalErr                         = ecode.New(13003)
	ChatCooling                            = ecode.New(13004)
	ChatNoMoney                            = ecode.New(13005)
	ChatNoGetIcon                          = ecode.New(13006)
	ChatNoSendMe                           = ecode.New(13007)
	ChatNoPeer                             = ecode.New(13008)
	ChatNoGuild                            = ecode.New(13009)
	ChatNoTeam                             = ecode.New(13010)
	ChatTextTooLong                        = ecode.New(13011)
	RedPacketTypeGetLimit                  = ecode.New(13105)
	RedPacketNoConfig                      = ecode.New(13106)
	DevilConquerRpcErr                     = ecode.New(13201)
	DevilConquerNoInDungeon                = ecode.New(13206)
	DevilConquerCooling                    = ecode.New(13207)
	DevilConquerSameGuild                  = ecode.New(13208)
	DevilConquerLineupErr                  = ecode.New(13209)
	DevilConquerNoBattleCk                 = ecode.New(13210)
	DevilConquerBattleFailed               = ecode.New(13211)
	DevilConquerNoRobScore0                = ecode.New(13212)
	MiscNoUser                             = ecode.New(13500)
	MiscIsSelf                             = ecode.New(13501)
	AuctionRPCErr                          = ecode.New(13607)
	AuctionNotEnoughGold                   = ecode.New(13608)
	ActivityNotOpen                        = ecode.New(13701)
	ActivityTypeErr                        = ecode.New(13702)
	ActivityInfoErr                        = ecode.New(13703)
	GoldJadeNumberNotEnough                = ecode.New(13710)
	GoldJadeRandomCoinErr                  = ecode.New(13711)
	RewardAlreadyReceive                   = ecode.New(13712)
	RewardNotTime                          = ecode.New(13713)
	GiftDoNotBuy                           = ecode.New(13714)
	GiftNotReceive                         = ecode.New(13715)
	GiftNotBuy                             = ecode.New(13716)
	GiftAlreadyBuy                         = ecode.New(13717)
	GiftAlreadyReceive                     = ecode.New(13718)
	ConditionNotEnough                     = ecode.New(13719)
	ActivityPeentoIDErr                    = ecode.New(13720)
	ThreeWorldRefreshCountNotEnough        = ecode.New(15001)
	ThreeWorldPassed                       = ecode.New(15002)
	ThreeWorldIslandNotExist               = ecode.New(15003)
	ThreeWorldLastIsland                   = ecode.New(15004)
	ThreeWorldEventNotExist                = ecode.New(15005)
	ThreeWorldNotChose                     = ecode.New(15006)
	ThreeWorldNotBattleEvent               = ecode.New(15007)
	ThreeWorldNotBattleLost                = ecode.New(15008)
	ThreeWorldShopItemNotExist             = ecode.New(15009)
	ThreeWorldNotHealingEvent              = ecode.New(15010)
	ThreeWorldEventPassed                  = ecode.New(15011)
	ThreeWorldNotBoxEvent                  = ecode.New(15012)
	ThreeWorldNotTheasureEvent             = ecode.New(15013)
	ThreeWorldTheasureNotExist             = ecode.New(15014)
	ThreeWorldNotHeroEvent                 = ecode.New(15015)
	ThreeWorldHeroNotExist                 = ecode.New(15016)
	ThreeWorldResetCountNotEnough          = ecode.New(15017)
	ThreeWorldEventStateErr                = ecode.New(15018)
	ThreeWorldRefreshIncompleteEvent       = ecode.New(15019)
	ReceiveManualRewardFail                = ecode.New(16000)
	ReceiveManualGrandTotalTaskRewardFail  = ecode.New(16001)
	FangcunLocked                          = ecode.New(18000)
	FangcunAreaNotOpen                     = ecode.New(18001)
	FangcunBattleCountNotEnough            = ecode.New(18002)
	FangcunLineupError                     = ecode.New(18003)
	FangcunLineupEmpty                     = ecode.New(18004)
	FangcunBattleMaxFloor                  = ecode.New(18005)
	FangcunBattleFloorNotExist             = ecode.New(18006)
	FangcunBattleFloorLocked               = ecode.New(18007)
	FangcunBigRewardNotExist               = ecode.New(18008)
	FangcunBigRewardReceived               = ecode.New(18009)
	FangcunBigRewardFloorMismatch          = ecode.New(18010)
	FangcunScoreRewardNotExist             = ecode.New(18011)
	FangcunScoreRewardReceived             = ecode.New(18012)
	FangcunScoreRewardScoreMismatch        = ecode.New(18013)
	FangcunAreaNotExist                    = ecode.New(18014)
	FangcunBadBattleResult                 = ecode.New(18015)
	SocialRewardAlready                    = ecode.New(19001)
	SocialNotExist                         = ecode.New(19002)
	DeleteAccountError                     = ecode.New(20000)
	DeleteAccounting                       = ecode.New(20001)
	RequestRepeated                        = ecode.New(20002)
)