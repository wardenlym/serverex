/// Code generated by protocol.def.go; DO NOT EDIT.

package msg

type Code = int

const (
	C_GetUserInfo                    = 1
	C_EnterStage                     = 2
	C_ExitStage                      = 3
	C_BattleStart                    = 4
	C_BattleEnd                      = 5
	C_EnterDungeon                   = 6
	C_ExitDungeon                    = 7
	C_Reborn                         = 8
	C_OpenBox                        = 9
	C_PickupLoot                     = 10
	C_Equip                          = 11
	C_Unequip                        = 12
	C_Craft                          = 13
	C_Decompose                      = 14
	C_Enhance                        = 15
	C_ConsumItem                     = 16
	C_ListOnSale                     = 17
	C_RefreshOnSale                  = 18
	C_Purchase                       = 19
	C_ListDiamondPrice               = 20
	C_SubmitOrder                    = 21
	C_QueryOrderStatus               = 22
	C_LevelUp                        = 23
	C_JobUpgrade                     = 24
	C_RuneEquip                      = 25
	C_RuneUnequip                    = 26
	C_ExpandBagCell                  = 27
	C_ExpandStashCell                = 28
	C_MoveItemsBagToStash            = 29
	C_MoveItemsStashToBag            = 30
	C_MoveGoldBagToStash             = 31
	C_MoveGoldStashToBag             = 32
	C_DestroyItem                    = 33
	C_DestroyMultipleItem            = 34
	C_DestroyItemsInStash            = 35
	C_SortItemsInBag                 = 36
	C_SortItemsInStash               = 37
	C_RefreshDiamond                 = 38
	C_GetUserStatInfo                = 39
	C_UserRename                     = 40
	C_PresentRedeemCode              = 41
	C_ListNoticeInfo                 = 42
	C_ConfirmNoticeInfo              = 43
	C_DeleteNoticeInfo               = 44
	C_MarkAsReadNoticeInfo           = 45
	C_GetExploreInfo                 = 46
	C_AddExploreChargeCount          = 47
	C_ReduceExploreTime              = 48
	C_ExploreStart                   = 49
	C_ExploreViewAward               = 50
	C_ExploreEnd                     = 51
	C_PuzzleStart                    = 52
	C_PuzzleUpdateProcess            = 53
	C_PuzzleEnd                      = 54
	C_ActiveNewCharacter             = 55
	C_KillCharacter                  = 56
	C_GetDungeonRankingInfo          = 57
	C_GetDungeonRankingList          = 58
	C_GetPersonalRankingInfo         = 59
	C_GetPersonalRankingList         = 60
	C_GetThreadComment               = 61
	C_GetThreadHotComment            = 62
	C_GetHeroThreadId                = 63
	C_PostComment                    = 64
	C_PraiseThread                   = 65
	C_PraiseComment                  = 66
	C_GetUserProperties              = 67
	C_SetUserProperties              = 68
	C_GetAchievementAward            = 69
	C_GetOperationActivityInfo       = 70
	C_GetActivtyAwardFisrtPay        = 71
	C_GetActivtyAwardSignInFirst7Day = 72
	C_GetActivtyAwardSignIn30Day     = 73
	C_GetDungeonRule1                = 74
	C_GetDungeonRule2                = 75
	C_UseItem                        = 76
	C_GetVipInfo                     = 77
	C_VipYueKaMaterialGetAward       = 78
	C_PayGetCallbackUrl              = 101
	C_PayPrepare                     = 102
	C_PayVerifyReceipt               = 103
	C_PayAppleVerifyReceipt          = 104
	C_PayVerifyOrder                 = 105
	C_PayListOrder                   = 106
	C_PayGooglePurchase              = 107
	C_PayGoogleConsume               = 108
	C_AccountLogin                   = 151
	C_AccountLogout                  = 152
	C_GM_ActivateUser                = 201
	C_GM_ResetUser                   = 202
	C_GM_CreateItem                  = 203
	C_GM_SetMyMoney                  = 204
	C_GM_ResetJob                    = 205
	C_GM_SetVipInfo                  = 206
)

var Code_To_String = map[Code]string{
	C_GetUserInfo:                    "GetUserInfo",
	C_EnterStage:                     "EnterStage",
	C_ExitStage:                      "ExitStage",
	C_BattleStart:                    "BattleStart",
	C_BattleEnd:                      "BattleEnd",
	C_EnterDungeon:                   "EnterDungeon",
	C_ExitDungeon:                    "ExitDungeon",
	C_Reborn:                         "Reborn",
	C_OpenBox:                        "OpenBox",
	C_PickupLoot:                     "PickupLoot",
	C_Equip:                          "Equip",
	C_Unequip:                        "Unequip",
	C_Craft:                          "Craft",
	C_Decompose:                      "Decompose",
	C_Enhance:                        "Enhance",
	C_ConsumItem:                     "ConsumItem",
	C_ListOnSale:                     "ListOnSale",
	C_RefreshOnSale:                  "RefreshOnSale",
	C_Purchase:                       "Purchase",
	C_ListDiamondPrice:               "ListDiamondPrice",
	C_SubmitOrder:                    "SubmitOrder",
	C_QueryOrderStatus:               "QueryOrderStatus",
	C_LevelUp:                        "LevelUp",
	C_JobUpgrade:                     "JobUpgrade",
	C_RuneEquip:                      "RuneEquip",
	C_RuneUnequip:                    "RuneUnequip",
	C_ExpandBagCell:                  "ExpandBagCell",
	C_ExpandStashCell:                "ExpandStashCell",
	C_MoveItemsBagToStash:            "MoveItemsBagToStash",
	C_MoveItemsStashToBag:            "MoveItemsStashToBag",
	C_MoveGoldBagToStash:             "MoveGoldBagToStash",
	C_MoveGoldStashToBag:             "MoveGoldStashToBag",
	C_DestroyItem:                    "DestroyItem",
	C_DestroyMultipleItem:            "DestroyMultipleItem",
	C_DestroyItemsInStash:            "DestroyItemsInStash",
	C_SortItemsInBag:                 "SortItemsInBag",
	C_SortItemsInStash:               "SortItemsInStash",
	C_RefreshDiamond:                 "RefreshDiamond",
	C_GetUserStatInfo:                "GetUserStatInfo",
	C_UserRename:                     "UserRename",
	C_PresentRedeemCode:              "PresentRedeemCode",
	C_ListNoticeInfo:                 "ListNoticeInfo",
	C_ConfirmNoticeInfo:              "ConfirmNoticeInfo",
	C_DeleteNoticeInfo:               "DeleteNoticeInfo",
	C_MarkAsReadNoticeInfo:           "MarkAsReadNoticeInfo",
	C_GetExploreInfo:                 "GetExploreInfo",
	C_AddExploreChargeCount:          "AddExploreChargeCount",
	C_ReduceExploreTime:              "ReduceExploreTime",
	C_ExploreStart:                   "ExploreStart",
	C_ExploreViewAward:               "ExploreViewAward",
	C_ExploreEnd:                     "ExploreEnd",
	C_PuzzleStart:                    "PuzzleStart",
	C_PuzzleUpdateProcess:            "PuzzleUpdateProcess",
	C_PuzzleEnd:                      "PuzzleEnd",
	C_ActiveNewCharacter:             "ActiveNewCharacter",
	C_KillCharacter:                  "KillCharacter",
	C_GetDungeonRankingInfo:          "GetDungeonRankingInfo",
	C_GetDungeonRankingList:          "GetDungeonRankingList",
	C_GetPersonalRankingInfo:         "GetPersonalRankingInfo",
	C_GetPersonalRankingList:         "GetPersonalRankingList",
	C_GetThreadComment:               "GetThreadComment",
	C_GetThreadHotComment:            "GetThreadHotComment",
	C_GetHeroThreadId:                "GetHeroThreadId",
	C_PostComment:                    "PostComment",
	C_PraiseThread:                   "PraiseThread",
	C_PraiseComment:                  "PraiseComment",
	C_GetUserProperties:              "GetUserProperties",
	C_SetUserProperties:              "SetUserProperties",
	C_GetAchievementAward:            "GetAchievementAward",
	C_GetOperationActivityInfo:       "GetOperationActivityInfo",
	C_GetActivtyAwardFisrtPay:        "GetActivtyAwardFisrtPay",
	C_GetActivtyAwardSignInFirst7Day: "GetActivtyAwardSignInFirst7Day",
	C_GetActivtyAwardSignIn30Day:     "GetActivtyAwardSignIn30Day",
	C_GetDungeonRule1:                "GetDungeonRule1",
	C_GetDungeonRule2:                "GetDungeonRule2",
	C_UseItem:                        "UseItem",
	C_GetVipInfo:                     "GetVipInfo",
	C_VipYueKaMaterialGetAward:       "VipYueKaMaterialGetAward",
	C_PayGetCallbackUrl:              "PayGetCallbackUrl",
	C_PayPrepare:                     "PayPrepare",
	C_PayVerifyReceipt:               "PayVerifyReceipt",
	C_PayAppleVerifyReceipt:          "PayAppleVerifyReceipt",
	C_PayVerifyOrder:                 "PayVerifyOrder",
	C_PayListOrder:                   "PayListOrder",
	C_PayGooglePurchase:              "PayGooglePurchase",
	C_PayGoogleConsume:               "PayGoogleConsume",
	C_AccountLogin:                   "AccountLogin",
	C_AccountLogout:                  "AccountLogout",
	C_GM_ActivateUser:                "GM_ActivateUser",
	C_GM_ResetUser:                   "GM_ResetUser",
	C_GM_CreateItem:                  "GM_CreateItem",
	C_GM_SetMyMoney:                  "GM_SetMyMoney",
	C_GM_ResetJob:                    "GM_ResetJob",
	C_GM_SetVipInfo:                  "GM_SetVipInfo",
}

const (
	Success                    = 0
	ErrNotImplemented          = 1
	ErrSendFailure             = 2
	ErrRecvFailure             = 3
	ErrRpcTimeout              = 4
	ErrServerNotExist          = 5
	ErrDecoding                = 6
	ErrEncoding                = 7
	ErrDataInconsistent        = 8
	ErrImpossiableLogic        = 1000
	ErrInvalidUserId           = 1001
	ErrInvalidCharacterId      = 1002
	ErrInvalidStageId          = 1003
	ErrInvalidCellIndex        = 1004
	ErrInvalidNoticeId         = 1005
	ErrInvalidDungeonState     = 1006
	ErrInvalidUserInput        = 1007
	ErrInvalidChargeCount      = 1008
	ErrInvalidExploreId        = 1009
	ErrInvalidGoodsId          = 1010
	ErrInvalidItemId           = 1011
	ErrInvalidHeroId           = 1012
	ErrInvalidRuneTypeId       = 1013
	ErrInvalidAchievementId    = 1014
	ErrInvalidActivityProcess  = 1015
	ErrInvalidActivityTime     = 1016
	ErrInvalidDungeonRule      = 1017
	ErrGoldNotEnough           = 2001
	ErrDiamondNotEnough        = 2002
	ErrRequireItemNotEnough    = 2003
	ErrBagCellNotEnough        = 2004
	ErrStashCellNotEnough      = 2005
	ErrAvalibleCountNotEnough  = 2006
	ErrLevelNotEnough          = 2007
	ErrOutOfRangeLimit         = 2008
	ErrVipYueKaGetAwardExpire  = 2009
	ErrVipYueKaGetAwardFailure = 2010
	ErrVerifyReceiptFailure    = 3001
	ErrInvalidProductId        = 3002
	ErrTransactionIdReceipted  = 3003
)

var ErrCode_To_String = map[ErrCode]string{
	Success:                    "Success",
	ErrNotImplemented:          "ErrNotImplemented",
	ErrSendFailure:             "ErrSendFailure",
	ErrRecvFailure:             "ErrRecvFailure",
	ErrRpcTimeout:              "ErrRpcTimeout",
	ErrServerNotExist:          "ErrServerNotExist",
	ErrDecoding:                "ErrDecoding",
	ErrEncoding:                "ErrEncoding",
	ErrDataInconsistent:        "ErrDataInconsistent",
	ErrImpossiableLogic:        "ErrImpossiableLogic",
	ErrInvalidUserId:           "ErrInvalidUserId",
	ErrInvalidCharacterId:      "ErrInvalidCharacterId",
	ErrInvalidStageId:          "ErrInvalidStageId",
	ErrInvalidCellIndex:        "ErrInvalidCellIndex",
	ErrInvalidNoticeId:         "ErrInvalidNoticeId",
	ErrInvalidDungeonState:     "ErrInvalidDungeonState",
	ErrInvalidUserInput:        "ErrInvalidUserInput",
	ErrInvalidChargeCount:      "ErrInvalidChargeCount",
	ErrInvalidExploreId:        "ErrInvalidExploreId",
	ErrInvalidGoodsId:          "ErrInvalidGoodsId",
	ErrInvalidItemId:           "ErrInvalidItemId",
	ErrInvalidHeroId:           "ErrInvalidHeroId",
	ErrInvalidRuneTypeId:       "ErrInvalidRuneTypeId",
	ErrInvalidAchievementId:    "ErrInvalidAchievementId",
	ErrInvalidActivityProcess:  "ErrInvalidActivityProcess",
	ErrInvalidActivityTime:     "ErrInvalidActivityTime",
	ErrInvalidDungeonRule:      "ErrInvalidDungeonRule",
	ErrGoldNotEnough:           "ErrGoldNotEnough",
	ErrDiamondNotEnough:        "ErrDiamondNotEnough",
	ErrRequireItemNotEnough:    "ErrRequireItemNotEnough",
	ErrBagCellNotEnough:        "ErrBagCellNotEnough",
	ErrStashCellNotEnough:      "ErrStashCellNotEnough",
	ErrAvalibleCountNotEnough:  "ErrAvalibleCountNotEnough",
	ErrLevelNotEnough:          "ErrLevelNotEnough",
	ErrOutOfRangeLimit:         "ErrOutOfRangeLimit",
	ErrVipYueKaGetAwardExpire:  "ErrVipYueKaGetAwardExpire",
	ErrVipYueKaGetAwardFailure: "ErrVipYueKaGetAwardFailure",
	ErrVerifyReceiptFailure:    "ErrVerifyReceiptFailure",
	ErrInvalidProductId:        "ErrInvalidProductId",
	ErrTransactionIdReceipted:  "ErrTransactionIdReceipted",
}
