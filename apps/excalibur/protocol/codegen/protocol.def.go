package codegen

import "gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"

// namespace msg
const ( // enum Msg from 1
	_ int = iota
	GetUserInfo

	EnterStage
	ExitStage
	BattleStart
	BattleEnd
	EnterDungeon
	ExitDungeon
	Reborn
	OpenBox
	PickupLoot

	Equip
	Unequip

	Craft
	Decompose
	Enhance
	ConsumItem

	ListOnSale
	RefreshOnSale
	Purchase

	ListDiamondPrice
	SubmitOrder
	QueryOrderStatus

	LevelUp
	JobUpgrade
	RuneEquip
	RuneUnequip

	ExpandBagCell
	ExpandStashCell

	MoveItemsBagToStash
	MoveItemsStashToBag
	MoveGoldBagToStash
	MoveGoldStashToBag
	DestroyItem
	DestroyMultipleItem
	DestroyItemsInStash
	SortItemsInBag
	SortItemsInStash

	RefreshDiamond

	GetUserStatInfo
	UserRename
	PresentRedeemCode

	ListNoticeInfo
	ConfirmNoticeInfo
	DeleteNoticeInfo
	MarkAsReadNoticeInfo

	GetExploreInfo
	AddExploreChargeCount
	ReduceExploreTime
	ExploreStart
	ExploreViewAward
	ExploreEnd

	PuzzleStart
	PuzzleUpdateProcess
	PuzzleEnd

	ActiveNewCharacter
	KillCharacter

	GetDungeonRankingInfo
	GetDungeonRankingList
	GetPersonalRankingInfo
	GetPersonalRankingList

	GetThreadComment
	GetThreadHotComment
	GetHeroThreadId
	PostComment
	PraiseThread
	PraiseComment

	GetUserProperties
	SetUserProperties

	GetAchievementAward

	GetOperationActivityInfo
	GetActivtyAwardFisrtPay
	GetActivtyAwardSignInFirst7Day
	GetActivtyAwardSignIn30Day

	GetDungeonRule1
	GetDungeonRule2

	UseItem

	GetVipInfo
	VipYueKaMaterialGetAward
)

const ( // enum Msg from 101
	_ int = iota
	PayGetCallbackUrl
	PayPrepare
	PayVerifyReceipt
	PayAppleVerifyReceipt
	PayVerifyOrder
	PayListOrder

	PayGooglePurchase
	PayGoogleConsume
)

const ( // enum Msg from 151
	_ int = iota
	AccountLogin
	AccountLogout
)

const ( // enum Msg from 201
	_ int = iota
	GM_ActivateUser
	GM_ResetUser
	GM_CreateItem
	GM_SetMyMoney
	GM_ResetJob
	GM_SetVipInfo
)

const ( // enum Err from 0
	_                   int = -1 + iota
	Success                 // 成功的默认值为0
	ErrNotImplemented       // 服务端无法理解的消息
	ErrSendFailure          // 客户端发送失败
	ErrRecvFailure          // 客户端接收失败
	ErrRpcTimeout           // 调用超时
	ErrServerNotExist       // 服务器无法访问
	ErrDecoding             // 服务器解析协议失败
	ErrEncoding             // 客户端编码消息失败
	ErrDataInconsistent     // 未知的数据一致性出错
)

const ( // enum Err from 1000
	_                         int = iota
	ErrImpossiableLogic           // 进入无法处理的逻辑分支,暗示一个服务器的bug
	ErrInvalidUserId              // 无效用户
	ErrInvalidCharacterId         // 无效角色
	ErrInvalidStageId             // 无效关卡
	ErrInvalidCellIndex           // 无效下标操作
	ErrInvalidNoticeId            // 无效公告
	ErrInvalidDungeonState        // 无效地下城状态
	ErrInvalidUserInput           // 无效用户输入
	ErrInvalidChargeCount         // 无效充能数额
	ErrInvalidExploreId           // 无效冒险派遣
	ErrInvalidGoodsId             // 无效商品
	ErrInvalidItemId              // 无效物品
	ErrInvalidHeroId              // 无效英雄
	ErrInvalidRuneTypeId          // 无效符文类型
	ErrInvalidAchievementId       // 无效成就
	ErrInvalidActivityProcess     // 无效运营活动进度
	ErrInvalidActivityTime        // 无效运营活动时间
	ErrInvalidDungeonRule         // 无效的地下城规则
)

const ( // enum Err from 2001
	_                          int = iota
	ErrGoldNotEnough               // 金币不足
	ErrDiamondNotEnough            // 钻石不足
	ErrRequireItemNotEnough        // 所需指定物品数量不足
	ErrBagCellNotEnough            // 背包空间不足
	ErrStashCellNotEnough          // 仓库空间不足
	ErrAvalibleCountNotEnough      // 可用数量不足
	ErrLevelNotEnough              // 等级不足
	ErrOutOfRangeLimit             // 超出范围
	ErrVipYueKaGetAwardExpire      // 材料月卡过期
	ErrVipYueKaGetAwardFailure     // 材料月卡领取失败
)

const ( // enum Err from 3001
	_                         int = iota
	ErrVerifyReceiptFailure       // 充值收据验证失败
	ErrInvalidProductId           // 商品id不存在
	ErrTransactionIdReceipted     // 已经处理过的交易
)

// namespace odm
type User struct {
	Uid                  int64
	Nickname             string
	CreatedAt            string
	UpdatedAt            string
	CreateAtUnix         int64
	UpdatedAtUnix        int64
	CharacterIds         []string
	Characters           map[string]Character
	Stash                Stash
	Diamond              uint
	VipInfo              VipInfo
	OnSaleInfos          []OnSaleInfo
	OnSaleLastUpdateTime int64
	RankingInfo          RankingInfo
	StatInfo             StatInfo
	ExploreInfo          ExploreInfo
	BestDungeonRanking   []DungeonRankingInfo
	OperationActivity    OperationActivity
}

type VipInfo struct {
	IsVip_Godbless                 bool
	Vip_YueKa_Material_ExpireTime  int64
	Vip_YueKa_Material_LastGetTime int64
}

type Character struct {
	HeroId          int
	Level           int
	Gold            uint
	BattleAttribute Attribute
	RankingInfo     RankingInfo
	DungeonStatus   DungeonStatus
	RuneTree        RuneTree
	Bag             Bag
}

type Attribute struct {
	Hp         int
	Mp         int
	KillValue  uint
	JobTree    JobTree
	LuckyValue uint
}

type JobTree struct {
	JobStatus []JobStatus
}

type JobStatus struct {
	JobId   int
	JobStar uint
}

type RuneTree struct {
	Cells []RuneCell
}

type RuneCell struct {
	Item      *Item
	SlotColor int
	Enable    bool
}

type Bag struct {
	Equipments []Cell
	Cells      []Cell
}

type Stash struct {
	Gold  uint
	Cells []Cell
}

type Cell struct {
	Num  uint
	Item *Item
}

type Item struct {
	Guid            int64
	TypeId          int
	StarLevel       int
	StarUpgradeInfo []int
	Fresh           bool
}

type StageInfo struct {
	StageId      int
	MapId        int
	EntranceType int
	Seed         int
	NpcList      []Npc
	BoxList      []Box
	PuzzleInfo   PuzzleInfo
}

type Npc struct {
	NpcId     int
	NpcTypeId int
	IsDead    bool
}

type Box struct {
	Id       int
	TypeId   int
	IsOpened bool
}

type DungeonStatus struct {
	EnterAt           int64
	State             string
	StageProgress     []StageInfo
	CurrentStageInfo  StageInfo
	CurrentBattleNpc  Npc
	KilledNpcList     []Npc
	OpenBoxList       []Npc
	UnpickedLoot      []Cell
	LootTotal         []Cell
	GoldTotal         uint
	GoldExtTotal      uint
	AttributeSnapshot DungeonAttributeSnapshot
	StageStat         DungeonStageStat
}

type OnSaleInfo struct {
	Guid             int64
	GoodsId          int
	TypeId           int
	NumTotal         uint
	NumAvailable     uint
	NumSQ            uint
	AvailableStageId int
	CostGold         uint
	CostDiamond      uint
	Discount         uint
}

type DiamondPrice struct {
	GoodsId  int
	Num      uint
	PriceRMB uint
	Discount uint
}

type OrderInfo struct {
	OrderId         string
	GoodsTypeId     int
	Amount          uint
	PayToken        string
	TradeSubmitTime int64
	TradeDealTime   int64
}

type PlainCell struct {
	TypeId int
	Num    uint
}

type LocatedCell struct {
	Index int
	Cell  Cell
}

type RankingInfo struct {
	BestStageId int
}

type NoticeInfo struct {
	Id        int
	Read      bool
	Confirmed bool
	Deleted   bool
	SendTime  int64
	Title     string
	Sender    string
	Content   string
	Gold      uint
	Diamond   uint
	Gift      []Cell
}

type UserNoticeInfo struct {
	Uid         int64
	NoticeInfos []NoticeInfo
}

type StatInfo struct {
	CreateTime          int64
	LoginDaysTotal      uint
	LastLoginTime       int64
	DeadTimes           uint
	AchievementScore    uint
	AwardItemCount      uint
	KillMonsterCount    uint
	HuoDeChengZhuangShu uint
}

type ExploreInfo struct {
	LastRefreshChargeTime int64
	ChargeAvailable       int
	ChargeTotal           int
	Status                []*ExploreStatus
}

type ExploreStatus struct {
	CharacterId    string
	HeroId         int
	IsExploring    bool
	ExploreStageId int
	DispatchedAt   int64
	CompleteAt     int64
}

type PuzzleInfo struct {
	RandomSeed int
	Total      int
	Process    int
	Selected   int
}

type DungeonStageStat struct {
	Damage  int
	Healing int
	Kill    uint
}

type DungeonAttributeSnapshot struct {
	Hp        uint
	Mp        uint
	MinAttack uint
	MaxAttack uint
	MinDef    uint
	MaxDef    uint
	HeroLevel uint
	GearScore int
}

type DungeonScore struct {
	BestStageId      int
	OpenBoxScore     int
	KillMonsterScore int
	HealthScore      int
	GearScore        int
	KillBossScore    int
	AwardItemScore   int
	TimeCost         int64
}

type DungeonRankingInfo struct {
	Uid               int64
	CharacterId       string
	Nickname          string
	Time              int64
	Date              int64
	Rank              int
	ChapterId         int
	AttributeSnapshot DungeonAttributeSnapshot
	StageStat         DungeonStageStat
	Equipments        []Cell
	JobTree           JobTree
	ThreadId          int64
}

type PersonalRankingInfo struct {
	Uid         int64
	Nickname    string
	Date        int64
	Rank        int
	Gold        uint
	Kill        uint
	Achievement uint
	ThreadId    int64
}

type Comment struct {
	ThreadId  int64
	CommentId int64
	Uid       int64
	Date      int64
	Nickname  string
	Praise    int
	Content   string
}

type OperationActivity struct {
	FisrtPay        ActivityProcessInfo
	SignInFirst7Day ActivityProcessInfo
	SignIn30Day     ActivityProcessInfo
}

type ActivityProcessInfo struct {
	Process     int
	LastGetTime int64
}

// namespace rpc
type AccountLoginRequest struct {
}
type AccountLoginResponse struct {
}

type AccountLogoutRequest struct {
}
type AccountLogoutResponse struct {
	Message string
}

type GetUserInfoRequest struct {
}
type GetUserInfoResponse struct {
	User *odm.User
}

type EnterDungeonRequest struct {
	CharacterId string
}
type EnterDungeonResponse struct {
	State string
}

type ExitDungeonRequest struct {
	CharacterId string
	IsSafeExit  bool
}
type ExitDungeonResponse struct {
	State          string
	Gold           uint
	StashGold      uint
	DungeonGold    uint
	DungeonExtGold uint
	NpcList        []odm.Npc
	ItemList       []odm.Cell
	ItemPlainList  []odm.PlainCell
	EquipmentsDiff []odm.LocatedCell
	BagDiff        []odm.LocatedCell
	DungeonScore   odm.DungeonScore
}

type RebornRequest struct {
	CharacterId string
	Diamond     uint
}
type RebornResponse struct {
	State   string
	Diamond uint
}

type EnterStageRequest struct {
	CharacterId string
	StageId     int
}
type EnterStageResponse struct {
	State     string
	StageInfo odm.StageInfo
}

type ExitStageRequest struct {
	CharacterId       string
	StageId           int
	StageClear        bool
	IsChapterEnding   bool
	StageStat         odm.DungeonStageStat
	AttributeSnapshot odm.DungeonAttributeSnapshot
}
type ExitStageResponse struct {
	State       string
	BestStageId int
}

type BattleStartRequest struct {
	CharacterId string
	NpcId       int
	NpcTypeId   int
	LuckyValue  uint
}
type BattleStartResponse struct {
	ErrMsg string
	State  string
}

type BattleEndRequest struct {
	CharacterId  string
	NpcId        int
	NpcType      int
	BattleResult string
}
type BattleEndResponse struct {
	State              string
	KillValueIncreased uint
	KillValue          uint
	GoldIncreased      uint
	GoldExtIncreased   uint
	Gold               uint
	UnpickupItems      []odm.Cell
	AllAwardItems      []odm.Cell
	BagDiff            []odm.LocatedCell
}

type DestroyItemRequest struct {
	CharacterId string
	CellIndex   int
}
type DestroyItemResponse struct {
}

type ConsumItemRequest struct {
	CharacterId string
	CellIndex   int
	Num         uint
}
type ConsumItemResponse struct {
	CellNewStatus odm.Cell
}

type UseItemRequest struct {
	CharacterId string
	TypeId      int
	Num         uint
}
type UseItemResponse struct {
	Diamond   uint
	Gold      uint
	StashGold uint
	BagDiff   []odm.LocatedCell
}

type EquipRequest struct {
	CharacterId   string
	FromCellIndex int
	ToCellIndex   int
}
type EquipResponse struct {
}

type UnequipRequest struct {
	CharacterId   string
	FromCellIndex int
	ToCellIndex   int
}
type UnequipResponse struct {
}

type ListOnSaleRequest struct {
	CharacterId string
}
type ListOnSaleResponse struct {
	LastTime        int64
	CurrentTime     int64
	NextRefreshTime int64
	OnSaleInfos     []odm.OnSaleInfo
}

type RefreshOnSaleRequest struct {
	CharacterId string
	CostGold    uint
	CostDiamond uint
}
type RefreshOnSaleResponse struct {
	Diamond uint
}

type PurchaseRequest struct {
	CharacterId string
	Guid        int64
	GoodsId     int
	TypeId      int
	Num         uint
	CostGold    uint
	CostDiamond uint
}
type PurchaseResponse struct {
	Gold      uint
	StashGold uint
	Diamond   uint
	BagDiff   []odm.LocatedCell
}

type ListDiamondPriceRequest struct {
	CharacterId string
}
type ListDiamondPriceResponse struct {
	DiamondPrices []odm.DiamondPrice
}

type SubmitOrderRequest struct {
	CharacterId string
	GoodsTypeId int
	Amount      uint
	PayToken    string
}
type SubmitOrderResponse struct {
	OrderInfo odm.OrderInfo
}

type QueryOrderStatusRequest struct {
	OrderId string
}
type QueryOrderStatusResponse struct {
	OrderInfo odm.OrderInfo
}

type CraftRequest struct {
	CharacterId   string
	ToCraftTypeId int
	Requires      []odm.PlainCell
	CostGold      uint
	CostDiamond   uint
}
type CraftResponse struct {
	CraftedItems odm.Cell
	Gold         uint
	StashGold    uint
	Diamond      uint
	BagDiff      []odm.LocatedCell
}

type DecomposeRequest struct {
	CharacterId         string
	DecomposedCellIndex uint
	CostGold            uint
	CostDiamond         uint
}
type DecomposeResponse struct {
	TypeId    int
	Num       uint
	Gold      uint
	StashGold uint
	Diamond   uint
	BagDiff   []odm.LocatedCell
}

type EnhanceRequest struct {
	CharacterId    string
	ToEnhanceIndex int
	CostGold       uint
	CostDiamond    uint
	Requires       []odm.PlainCell
}
type EnhanceResponse struct {
	EnhancedItem odm.Cell
	Gold         uint
	StashGold    uint
	Diamond      uint
	BagDiff      []odm.LocatedCell
}

type LevelUpRequest struct {
	CharacterId  string
	CurrentLevel int
	CostGold     uint
	Requires     []odm.PlainCell
}
type LevelUpResponse struct {
	CurrentLevel int
	Gold         uint
	StashGold    uint
	BagDiff      []odm.LocatedCell
	RuneTree     odm.RuneTree
}

type JobUpgradeRequest struct {
	CharacterId   string
	CostKillValue uint
	JobStatus     odm.JobStatus
}
type JobUpgradeResponse struct {
	JobTree   odm.JobTree
	KillValue uint
}

type RuneEquipRequest struct {
	CharacterId   string
	FromCellIndex int
	ToCellIndex   int
}
type RuneEquipResponse struct {
	RuneTree odm.RuneTree
	BagDiff  []odm.LocatedCell
}

type RuneUnequipRequest struct {
	CharacterId   string
	FromCellIndex int
}
type RuneUnequipResponse struct {
	CellIndex int
	RuneTree  odm.RuneTree
	BagDiff   []odm.LocatedCell
}

type ExpandBagCellRequest struct {
	CharacterId string
	GoodsId     int
}
type ExpandBagCellResponse struct {
	Start   int
	Offset  int
	Diamond uint
}

type PickupLootRequest struct {
	CharacterId string
	From        []int
}
type PickupLootResponse struct {
	BagDiff []odm.LocatedCell
}

type DestroyMultipleItemRequest struct {
	CharacterId string
	CellIndices []int
}
type DestroyMultipleItemResponse struct {
	BagDiff []odm.LocatedCell
}

type RefreshDiamondRequest struct {
}
type RefreshDiamondResponse struct {
	Diamond uint
}

type ListNoticeInfoRequest struct {
	CharacterId string
}
type ListNoticeInfoResponse struct {
	NoticeInfos []odm.NoticeInfo
}

type ConfirmNoticeInfoRequest struct {
	CharacterId  string
	NoticeInfoId int
}
type ConfirmNoticeInfoResponse struct {
	BagGold   uint
	StashGold uint
	Diamond   uint
	BagDiff   []odm.LocatedCell
}

type DeleteNoticeInfoRequest struct {
	CharacterId  string
	NoticeInfoId int
}
type DeleteNoticeInfoResponse struct {
}

type GetUserStatInfoRequest struct {
}
type GetUserStatInfoResponse struct {
	StatInfo odm.StatInfo
}

type UserRenameRequest struct {
	NewName string
}
type UserRenameResponse struct {
	Nickname string
	Diamond  uint
}

type PresentRedeemCodeRequest struct {
	CharacterId string
	RedeemCode  string
}
type PresentRedeemCodeResponse struct {
	AwardItems []odm.Cell
	BagDiff    []odm.LocatedCell
}

type MoveItemsStashToBagRequest struct {
	CharacterId string
	Indices     []int
}
type MoveItemsStashToBagResponse struct {
	StashDiff []odm.LocatedCell
	BagDiff   []odm.LocatedCell
}

type MoveItemsBagToStashRequest struct {
	CharacterId string
	Indices     []int
}
type MoveItemsBagToStashResponse struct {
	StashDiff []odm.LocatedCell
	BagDiff   []odm.LocatedCell
}

type MoveGoldStashToBagRequest struct {
	CharacterId string
	Gold        uint
}
type MoveGoldStashToBagResponse struct {
	BagGold   uint
	StashGold uint
}

type MoveGoldBagToStashRequest struct {
	CharacterId string
	Gold        uint
}
type MoveGoldBagToStashResponse struct {
	BagGold   uint
	StashGold uint
}

type DestroyItemsInStashRequest struct {
	CharacterId string
	Indices     []int
}
type DestroyItemsInStashResponse struct {
	StashDiff []odm.LocatedCell
}

type SortItemsInStashRequest struct {
}
type SortItemsInStashResponse struct {
	StashDiff []odm.LocatedCell
}

type ExpandStashCellRequest struct {
}
type ExpandStashCellResponse struct {
	Start   int
	Offset  int
	Diamond uint
}

type OpenBoxRequest struct {
	CharacterId string
	TypeId      int
}
type OpenBoxResponse struct {
	UnpickupItems []odm.Cell
	AllAwardItems []odm.Cell
	BagDiff       []odm.LocatedCell
}

type SortItemsInBagRequest struct {
	CharacterId string
}
type SortItemsInBagResponse struct {
	BagDiff []odm.LocatedCell
}

type MarkAsReadNoticeInfoRequest struct {
	CharacterId  string
	NoticeInfoId int
}
type MarkAsReadNoticeInfoResponse struct {
}

type GetExploreInfoRequest struct {
}
type GetExploreInfoResponse struct {
	ServerTime  int64
	ExploreInfo odm.ExploreInfo
}

type AddExploreChargeCountRequest struct {
	Count int
}
type AddExploreChargeCountResponse struct {
	ServerTime  int64
	Diamond     uint
	ExploreInfo odm.ExploreInfo
}

type ReduceExploreTimeRequest struct {
	CharacterId    string
	ExploreStageId int
	ReduceTime     int64
}
type ReduceExploreTimeResponse struct {
	ServerTime  int64
	Diamond     uint
	ExploreInfo odm.ExploreInfo
}

type ExploreStartRequest struct {
	CharacterId    string
	ExploreStageId int
}
type ExploreStartResponse struct {
	ExploreInfo odm.ExploreInfo
}

type ExploreViewAwardRequest struct {
	CharacterId    string
	ExploreStageId int
}
type ExploreViewAwardResponse struct {
	Award []odm.PlainCell
}

type ExploreEndRequest struct {
	CharacterId        string
	ExploreCharacterId string
	ExploreStageId     int
}
type ExploreEndResponse struct {
	ExploreInfo odm.ExploreInfo
	AwardItems  []odm.Cell
	BagDiff     []odm.LocatedCell
}

type PuzzleStartRequest struct {
	CharacterId string
	StageId     int
	RandomSeed  int
	Total       int
}
type PuzzleStartResponse struct {
}

type PuzzleUpdateProcessRequest struct {
	CharacterId string
	StageId     int
	Process     int
}
type PuzzleUpdateProcessResponse struct {
}

type PuzzleEndRequest struct {
	CharacterId string
	StageId     int
	Selected    int
}
type PuzzleEndResponse struct {
}

type ActiveNewCharacterRequest struct {
	CharacterId string
	HeroId      int
}
type ActiveNewCharacterResponse struct {
	CharacterId string
	Diamond     uint
	Gold        uint
	StashGold   uint
}

type KillCharacterRequest struct {
	CharacterId string
}
type KillCharacterResponse struct {
	State   string
	BagDiff []odm.LocatedCell
}

type GetDungeonRankingInfoRequest struct {
	ChapterId int
}
type GetDungeonRankingInfoResponse struct {
	Info odm.DungeonRankingInfo
}

type GetDungeonRankingListRequest struct {
	ChapterId int
	IndexFrom int
	Count     int
}
type GetDungeonRankingListResponse struct {
	Info []odm.DungeonRankingInfo
}

type GetPersonalRankingInfoRequest struct {
	RankingType string
}
type GetPersonalRankingInfoResponse struct {
	Info odm.PersonalRankingInfo
}

type GetPersonalRankingListRequest struct {
	RankingType string
	IndexFrom   int
	Count       int
}
type GetPersonalRankingListResponse struct {
	Info []odm.PersonalRankingInfo
}

type GetThreadCommentRequest struct {
	ThreadId  int64
	IndexFrom int
	Count     int
}
type GetThreadCommentResponse struct {
	Comments []odm.Comment
}

type GetThreadHotCommentRequest struct {
	ThreadId int64
	TopN     int
}
type GetThreadHotCommentResponse struct {
	Comments []odm.Comment
}

type GetHeroThreadIdRequest struct {
	HeroId int
}
type GetHeroThreadIdResponse struct {
	ThreadId int64
}

type PostCommentRequest struct {
	ThreadId int64
	Content  string
}
type PostCommentResponse struct {
}

type PraiseThreadRequest struct {
	ThreadId int64
}
type PraiseThreadResponse struct {
}

type PraiseCommentRequest struct {
	ThreadId  int64
	CommentId int64
}
type PraiseCommentResponse struct {
}

type GetUserPropertiesRequest struct {
	Key string
}

type GetUserPropertiesResponse struct {
	Doc string
}

type SetUserPropertiesRequest struct {
	Key string
	Doc string
}

type SetUserPropertiesResponse struct {
}

type GetAchievementAwardRequest struct {
	CharacterId   string
	AchievementId int
}
type GetAchievementAwardResponse struct {
	AchievementScore uint
	BagDiff          []odm.LocatedCell
	Diamond          uint
	Gold             uint
}

type GetOperationActivityInfoRequest struct {
}
type GetOperationActivityInfoResponse struct {
	ServerTime        int64
	OperationActivity odm.OperationActivity
}

type GetActivtyAwardFisrtPayRequest struct {
	CharacterId string
}
type GetActivtyAwardFisrtPayResponse struct {
	ActivityId  int
	ProcessInfo odm.ActivityProcessInfo
	BagDiff     []odm.LocatedCell
	Diamond     uint
	Gold        uint
}

type GetActivtyAwardSignInFirst7DayRequest struct {
	CharacterId string
}
type GetActivtyAwardSignInFirst7DayResponse struct {
	ActivityId  int
	ProcessInfo odm.ActivityProcessInfo
	BagDiff     []odm.LocatedCell
	Diamond     uint
	Gold        uint
}

type GetActivtyAwardSignIn30DayRequest struct {
	CharacterId string
}
type GetActivtyAwardSignIn30DayResponse struct {
	ActivityId  int
	ProcessInfo odm.ActivityProcessInfo
	BagDiff     []odm.LocatedCell
	Diamond     uint
	Gold        uint
}

type GetDungeonRule1Request struct {
	CharacterId string
}
type GetDungeonRule1Response struct {
	BagDiff []odm.LocatedCell
}

type GetDungeonRule2Request struct {
	CharacterId string
	Level       int
	IsEquiped   bool
	CellIndex   int
}
type GetDungeonRule2Response struct {
	BagDiff        []odm.LocatedCell
	EquipmentsDiff []odm.LocatedCell
}

type GetVipInfoRequest struct {
}
type GetVipInfoResponse struct {
	ServerTime int64
	VipInfo    odm.VipInfo
}

type VipYueKaMaterialGetAwardRequest struct {
	CharacterId string
}
type VipYueKaMaterialGetAwardResponse struct {
	Diamond uint
	BagDiff []odm.LocatedCell
}

// pay
type PayGetCallbackUrlRequest struct {
}
type PayGetCallbackUrlResponse struct {
}

type PayPrepareRequest struct {
}
type PayPrepareResponse struct {
}

type PayVerifyReceiptRequest struct {
}
type PayVerifyReceiptResponse struct {
}

type PayAppleVerifyReceiptRequest struct {
	UseSandbox  bool
	Receipt     string
	Uid         int64
	DeviceToken string
}
type PayAppleVerifyReceiptResponse struct {
	OrderId int
	Diamond uint
	SkuType string
}

type PayVerifyOrderRequest struct {
}
type PayVerifyOrderResponse struct {
}

type PayListOrderRequest struct {
}
type PayListOrderResponse struct {
}

type PayGooglePurchaseRequest struct {
	PurchaseInfo string
}
type PayGooglePurchaseResponse struct {
	GoogleErrMsg string
}

type PayGoogleConsumeRequest struct {
	OrderId       string
	PackageName   string
	ProductId     string
	PurchaseToken string
}
type PayGoogleConsumeResponse struct {
	GoogleErrMsg string
}

// GM
type GM_ActivateUserRequest struct {
}
type GM_ActivateUserResponse struct {
}

type GM_ResetUserRequest struct {
}
type GM_ResetUserResponse struct {
}

type GM_CreateItemRequest struct {
	CharacterId string
	TypeId      int
	StarLevel   int
	Num         uint
}
type GM_CreateItemResponse struct {
	NewItem   odm.Item
	Num       uint
	CellIndex int
}

type GM_SetMyMoneyRequest struct {
	Gold    uint
	Diamond uint
}
type GM_SetMyMoneyResponse struct {
}

type GM_ResetJobRequest struct {
	CharacterId string
	KillValue   uint
}
type GM_ResetJobResponse struct {
	JobTree   odm.JobTree
	KillValue uint
}

type GM_SetVipInfoRequest struct {
	VipInfo odm.VipInfo
}
type GM_SetVipInfoResponse struct {
}
