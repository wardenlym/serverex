/// Code generated by protocol.def.go; DO NOT EDIT.
namespace Network.sdk.Excalibur.Protocol
{
	public sealed partial class Msg
	{
		public enum Code : int
		{
            GetUserInfo                    = 1,
            EnterStage                     = 2,
            ExitStage                      = 3,
            BattleStart                    = 4,
            BattleEnd                      = 5,
            EnterDungeon                   = 6,
            ExitDungeon                    = 7,
            Reborn                         = 8,
            OpenBox                        = 9,
            PickupLoot                     = 10,
            Equip                          = 11,
            Unequip                        = 12,
            Craft                          = 13,
            Decompose                      = 14,
            Enhance                        = 15,
            ConsumItem                     = 16,
            ListOnSale                     = 17,
            RefreshOnSale                  = 18,
            Purchase                       = 19,
            ListDiamondPrice               = 20,
            SubmitOrder                    = 21,
            QueryOrderStatus               = 22,
            LevelUp                        = 23,
            JobUpgrade                     = 24,
            RuneEquip                      = 25,
            RuneUnequip                    = 26,
            ExpandBagCell                  = 27,
            ExpandStashCell                = 28,
            MoveItemsBagToStash            = 29,
            MoveItemsStashToBag            = 30,
            MoveGoldBagToStash             = 31,
            MoveGoldStashToBag             = 32,
            DestroyItem                    = 33,
            DestroyMultipleItem            = 34,
            DestroyItemsInStash            = 35,
            SortItemsInBag                 = 36,
            SortItemsInStash               = 37,
            RefreshDiamond                 = 38,
            GetUserStatInfo                = 39,
            UserRename                     = 40,
            PresentRedeemCode              = 41,
            ListNoticeInfo                 = 42,
            ConfirmNoticeInfo              = 43,
            DeleteNoticeInfo               = 44,
            MarkAsReadNoticeInfo           = 45,
            GetExploreInfo                 = 46,
            AddExploreChargeCount          = 47,
            ReduceExploreTime              = 48,
            ExploreStart                   = 49,
            ExploreViewAward               = 50,
            ExploreEnd                     = 51,
            PuzzleStart                    = 52,
            PuzzleUpdateProcess            = 53,
            PuzzleEnd                      = 54,
            ActiveNewCharacter             = 55,
            KillCharacter                  = 56,
            GetDungeonRankingInfo          = 57,
            GetDungeonRankingList          = 58,
            GetPersonalRankingInfo         = 59,
            GetPersonalRankingList         = 60,
            GetThreadComment               = 61,
            GetThreadHotComment            = 62,
            GetHeroThreadId                = 63,
            PostComment                    = 64,
            PraiseThread                   = 65,
            PraiseComment                  = 66,
            GetUserProperties              = 67,
            SetUserProperties              = 68,
            GetAchievementAward            = 69,
            GetOperationActivityInfo       = 70,
            GetActivtyAwardFisrtPay        = 71,
            GetActivtyAwardSignInFirst7Day = 72,
            GetActivtyAwardSignIn30Day     = 73,
            GetDungeonRule1                = 74,
            GetDungeonRule2                = 75,
            UseItem                        = 76,
            GetVipInfo                     = 77,
            VipYueKaMaterialGetAward       = 78,

            PayGetCallbackUrl              = 101,
            PayPrepare                     = 102,
            PayVerifyReceipt               = 103,
            PayAppleVerifyReceipt          = 104,
            PayVerifyOrder                 = 105,
            PayListOrder                   = 106,
            PayGooglePurchase              = 107,
            PayGoogleConsume               = 108,

            AccountLogin                   = 151,
            AccountLogout                  = 152,

            GM_ActivateUser                = 201,
            GM_ResetUser                   = 202,
            GM_CreateItem                  = 203,
            GM_SetMyMoney                  = 204,
            GM_ResetJob                    = 205,
            GM_SetVipInfo                  = 206,
        }

		public enum Err : int
		{
            Success                    = 0, // 成功的默认值为0
            ErrNotImplemented          = 1, // 服务端无法理解的消息
            ErrSendFailure             = 2, // 客户端发送失败
            ErrRecvFailure             = 3, // 客户端接收失败
            ErrRpcTimeout              = 4, // 调用超时
            ErrServerNotExist          = 5, // 服务器无法访问
            ErrDecoding                = 6, // 服务器解析协议失败
            ErrEncoding                = 7, // 客户端编码消息失败
            ErrDataInconsistent        = 8, // 未知的数据一致性出错

            ErrImpossiableLogic        = 1000, // 进入无法处理的逻辑分支,暗示一个服务器的bug
            ErrInvalidUserId           = 1001, // 无效用户
            ErrInvalidCharacterId      = 1002, // 无效角色
            ErrInvalidStageId          = 1003, // 无效关卡
            ErrInvalidCellIndex        = 1004, // 无效下标操作
            ErrInvalidNoticeId         = 1005, // 无效公告
            ErrInvalidDungeonState     = 1006, // 无效地下城状态
            ErrInvalidUserInput        = 1007, // 无效用户输入
            ErrInvalidChargeCount      = 1008, // 无效充能数额
            ErrInvalidExploreId        = 1009, // 无效冒险派遣
            ErrInvalidGoodsId          = 1010, // 无效商品
            ErrInvalidItemId           = 1011, // 无效物品
            ErrInvalidHeroId           = 1012, // 无效英雄
            ErrInvalidRuneTypeId       = 1013, // 无效符文类型
            ErrInvalidAchievementId    = 1014, // 无效成就
            ErrInvalidActivityProcess  = 1015, // 无效运营活动进度
            ErrInvalidActivityTime     = 1016, // 无效运营活动时间
            ErrInvalidDungeonRule      = 1017, // 无效的地下城规则

            ErrGoldNotEnough           = 2001, // 金币不足
            ErrDiamondNotEnough        = 2002, // 钻石不足
            ErrRequireItemNotEnough    = 2003, // 所需指定物品数量不足
            ErrBagCellNotEnough        = 2004, // 背包空间不足
            ErrStashCellNotEnough      = 2005, // 仓库空间不足
            ErrAvalibleCountNotEnough  = 2006, // 可用数量不足
            ErrLevelNotEnough          = 2007, // 等级不足
            ErrOutOfRangeLimit         = 2008, // 超出范围
            ErrVipYueKaGetAwardExpire  = 2009, // 材料月卡过期
            ErrVipYueKaGetAwardFailure = 2010, // 材料月卡领取失败

            ErrVerifyReceiptFailure    = 3001, // 充值收据验证失败
            ErrInvalidProductId        = 3002, // 商品id不存在
            ErrTransactionIdReceipted  = 3003, // 已经处理过的交易
        }
	}
}
