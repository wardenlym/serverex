package csv

type MonsterData struct {
	// 怪物ID
	Id int

	// 击杀点
	Killsvalue int

	// 怪物类型
	Monster_type int
}

var Table_MonsterData = map[int]MonsterData{}

func Load_Table_MonsterData() {
	Table_MonsterData = map[int]MonsterData{
		10000001: MonsterData{
			Id:           10000001,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000002: MonsterData{
			Id:           10000002,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000003: MonsterData{
			Id:           10000003,
			Killsvalue:   5,
			Monster_type: 3,
		},
		10000004: MonsterData{
			Id:           10000004,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000005: MonsterData{
			Id:           10000005,
			Killsvalue:   4,
			Monster_type: 2,
		},
		10000006: MonsterData{
			Id:           10000006,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000007: MonsterData{
			Id:           10000007,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000008: MonsterData{
			Id:           10000008,
			Killsvalue:   5,
			Monster_type: 3,
		},
		10000009: MonsterData{
			Id:           10000009,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000010: MonsterData{
			Id:           10000010,
			Killsvalue:   4,
			Monster_type: 2,
		},
		10000011: MonsterData{
			Id:           10000011,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000012: MonsterData{
			Id:           10000012,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000013: MonsterData{
			Id:           10000013,
			Killsvalue:   5,
			Monster_type: 3,
		},
		10000014: MonsterData{
			Id:           10000014,
			Killsvalue:   10,
			Monster_type: 3,
		},
		10000015: MonsterData{
			Id:           10000015,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000016: MonsterData{
			Id:           10000016,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000017: MonsterData{
			Id:           10000017,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000018: MonsterData{
			Id:           10000018,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000019: MonsterData{
			Id:           10000019,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000020: MonsterData{
			Id:           10000020,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000021: MonsterData{
			Id:           10000021,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000022: MonsterData{
			Id:           10000022,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000023: MonsterData{
			Id:           10000023,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000024: MonsterData{
			Id:           10000024,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000025: MonsterData{
			Id:           10000025,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000026: MonsterData{
			Id:           10000026,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000027: MonsterData{
			Id:           10000027,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000028: MonsterData{
			Id:           10000028,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000029: MonsterData{
			Id:           10000029,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000030: MonsterData{
			Id:           10000030,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000031: MonsterData{
			Id:           10000031,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000032: MonsterData{
			Id:           10000032,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000033: MonsterData{
			Id:           10000033,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000034: MonsterData{
			Id:           10000034,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000035: MonsterData{
			Id:           10000035,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000036: MonsterData{
			Id:           10000036,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000037: MonsterData{
			Id:           10000037,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000038: MonsterData{
			Id:           10000038,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000039: MonsterData{
			Id:           10000039,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000040: MonsterData{
			Id:           10000040,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000041: MonsterData{
			Id:           10000041,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000042: MonsterData{
			Id:           10000042,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000043: MonsterData{
			Id:           10000043,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000044: MonsterData{
			Id:           10000044,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000045: MonsterData{
			Id:           10000045,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000046: MonsterData{
			Id:           10000046,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000047: MonsterData{
			Id:           10000047,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000048: MonsterData{
			Id:           10000048,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000049: MonsterData{
			Id:           10000049,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000050: MonsterData{
			Id:           10000050,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000051: MonsterData{
			Id:           10000051,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000052: MonsterData{
			Id:           10000052,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000053: MonsterData{
			Id:           10000053,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000054: MonsterData{
			Id:           10000054,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000055: MonsterData{
			Id:           10000055,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000056: MonsterData{
			Id:           10000056,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000057: MonsterData{
			Id:           10000057,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000058: MonsterData{
			Id:           10000058,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000059: MonsterData{
			Id:           10000059,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000060: MonsterData{
			Id:           10000060,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000061: MonsterData{
			Id:           10000061,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000062: MonsterData{
			Id:           10000062,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000063: MonsterData{
			Id:           10000063,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000064: MonsterData{
			Id:           10000064,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000065: MonsterData{
			Id:           10000065,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000066: MonsterData{
			Id:           10000066,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000067: MonsterData{
			Id:           10000067,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000068: MonsterData{
			Id:           10000068,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000069: MonsterData{
			Id:           10000069,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000070: MonsterData{
			Id:           10000070,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000071: MonsterData{
			Id:           10000071,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000072: MonsterData{
			Id:           10000072,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000073: MonsterData{
			Id:           10000073,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000074: MonsterData{
			Id:           10000074,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000075: MonsterData{
			Id:           10000075,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000076: MonsterData{
			Id:           10000076,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000077: MonsterData{
			Id:           10000077,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000078: MonsterData{
			Id:           10000078,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000079: MonsterData{
			Id:           10000079,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000080: MonsterData{
			Id:           10000080,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000081: MonsterData{
			Id:           10000081,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000082: MonsterData{
			Id:           10000082,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000083: MonsterData{
			Id:           10000083,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000084: MonsterData{
			Id:           10000084,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000085: MonsterData{
			Id:           10000085,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000086: MonsterData{
			Id:           10000086,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000087: MonsterData{
			Id:           10000087,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000088: MonsterData{
			Id:           10000088,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000089: MonsterData{
			Id:           10000089,
			Killsvalue:   4,
			Monster_type: 4,
		},
		10000090: MonsterData{
			Id:           10000090,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000091: MonsterData{
			Id:           10000091,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000092: MonsterData{
			Id:           10000092,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000093: MonsterData{
			Id:           10000093,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000094: MonsterData{
			Id:           10000094,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000095: MonsterData{
			Id:           10000095,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000096: MonsterData{
			Id:           10000096,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000097: MonsterData{
			Id:           10000097,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000098: MonsterData{
			Id:           10000098,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000099: MonsterData{
			Id:           10000099,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000100: MonsterData{
			Id:           10000100,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000101: MonsterData{
			Id:           10000101,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000102: MonsterData{
			Id:           10000102,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000103: MonsterData{
			Id:           10000103,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000104: MonsterData{
			Id:           10000104,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000105: MonsterData{
			Id:           10000105,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000106: MonsterData{
			Id:           10000106,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000107: MonsterData{
			Id:           10000107,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000108: MonsterData{
			Id:           10000108,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000109: MonsterData{
			Id:           10000109,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000110: MonsterData{
			Id:           10000110,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000111: MonsterData{
			Id:           10000111,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000112: MonsterData{
			Id:           10000112,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000113: MonsterData{
			Id:           10000113,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000114: MonsterData{
			Id:           10000114,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000115: MonsterData{
			Id:           10000115,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000116: MonsterData{
			Id:           10000116,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000117: MonsterData{
			Id:           10000117,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000118: MonsterData{
			Id:           10000118,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000119: MonsterData{
			Id:           10000119,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000120: MonsterData{
			Id:           10000120,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000121: MonsterData{
			Id:           10000121,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000122: MonsterData{
			Id:           10000122,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000123: MonsterData{
			Id:           10000123,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000124: MonsterData{
			Id:           10000124,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000125: MonsterData{
			Id:           10000125,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000126: MonsterData{
			Id:           10000126,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000127: MonsterData{
			Id:           10000127,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000128: MonsterData{
			Id:           10000128,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000129: MonsterData{
			Id:           10000129,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000130: MonsterData{
			Id:           10000130,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000131: MonsterData{
			Id:           10000131,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000132: MonsterData{
			Id:           10000132,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000133: MonsterData{
			Id:           10000133,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000134: MonsterData{
			Id:           10000134,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000135: MonsterData{
			Id:           10000135,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000136: MonsterData{
			Id:           10000136,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000137: MonsterData{
			Id:           10000137,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000138: MonsterData{
			Id:           10000138,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000139: MonsterData{
			Id:           10000139,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000140: MonsterData{
			Id:           10000140,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000141: MonsterData{
			Id:           10000141,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000142: MonsterData{
			Id:           10000142,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000143: MonsterData{
			Id:           10000143,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000144: MonsterData{
			Id:           10000144,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000145: MonsterData{
			Id:           10000145,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000146: MonsterData{
			Id:           10000146,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000147: MonsterData{
			Id:           10000147,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000148: MonsterData{
			Id:           10000148,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000149: MonsterData{
			Id:           10000149,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000150: MonsterData{
			Id:           10000150,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000151: MonsterData{
			Id:           10000151,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000152: MonsterData{
			Id:           10000152,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000153: MonsterData{
			Id:           10000153,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000154: MonsterData{
			Id:           10000154,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000155: MonsterData{
			Id:           10000155,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000156: MonsterData{
			Id:           10000156,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000157: MonsterData{
			Id:           10000157,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000158: MonsterData{
			Id:           10000158,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000159: MonsterData{
			Id:           10000159,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000160: MonsterData{
			Id:           10000160,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000161: MonsterData{
			Id:           10000161,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000162: MonsterData{
			Id:           10000162,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000163: MonsterData{
			Id:           10000163,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000164: MonsterData{
			Id:           10000164,
			Killsvalue:   4,
			Monster_type: 4,
		},
		10000165: MonsterData{
			Id:           10000165,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000166: MonsterData{
			Id:           10000166,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000167: MonsterData{
			Id:           10000167,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000168: MonsterData{
			Id:           10000168,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000169: MonsterData{
			Id:           10000169,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000170: MonsterData{
			Id:           10000170,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000171: MonsterData{
			Id:           10000171,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000172: MonsterData{
			Id:           10000172,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000173: MonsterData{
			Id:           10000173,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000174: MonsterData{
			Id:           10000174,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000175: MonsterData{
			Id:           10000175,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000176: MonsterData{
			Id:           10000176,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000177: MonsterData{
			Id:           10000177,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000178: MonsterData{
			Id:           10000178,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000179: MonsterData{
			Id:           10000179,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000180: MonsterData{
			Id:           10000180,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000181: MonsterData{
			Id:           10000181,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000182: MonsterData{
			Id:           10000182,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000183: MonsterData{
			Id:           10000183,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000184: MonsterData{
			Id:           10000184,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000185: MonsterData{
			Id:           10000185,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000186: MonsterData{
			Id:           10000186,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000187: MonsterData{
			Id:           10000187,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000188: MonsterData{
			Id:           10000188,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000189: MonsterData{
			Id:           10000189,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000190: MonsterData{
			Id:           10000190,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000191: MonsterData{
			Id:           10000191,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000192: MonsterData{
			Id:           10000192,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000193: MonsterData{
			Id:           10000193,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000194: MonsterData{
			Id:           10000194,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000195: MonsterData{
			Id:           10000195,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000196: MonsterData{
			Id:           10000196,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000197: MonsterData{
			Id:           10000197,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000198: MonsterData{
			Id:           10000198,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000199: MonsterData{
			Id:           10000199,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000200: MonsterData{
			Id:           10000200,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000201: MonsterData{
			Id:           10000201,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000202: MonsterData{
			Id:           10000202,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000203: MonsterData{
			Id:           10000203,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000204: MonsterData{
			Id:           10000204,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000205: MonsterData{
			Id:           10000205,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000206: MonsterData{
			Id:           10000206,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000207: MonsterData{
			Id:           10000207,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000208: MonsterData{
			Id:           10000208,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000209: MonsterData{
			Id:           10000209,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000210: MonsterData{
			Id:           10000210,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000211: MonsterData{
			Id:           10000211,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000212: MonsterData{
			Id:           10000212,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000213: MonsterData{
			Id:           10000213,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000214: MonsterData{
			Id:           10000214,
			Killsvalue:   4,
			Monster_type: 4,
		},
		10000215: MonsterData{
			Id:           10000215,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000216: MonsterData{
			Id:           10000216,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000217: MonsterData{
			Id:           10000217,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000218: MonsterData{
			Id:           10000218,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000219: MonsterData{
			Id:           10000219,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000220: MonsterData{
			Id:           10000220,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000221: MonsterData{
			Id:           10000221,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000222: MonsterData{
			Id:           10000222,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000223: MonsterData{
			Id:           10000223,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000224: MonsterData{
			Id:           10000224,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000225: MonsterData{
			Id:           10000225,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000226: MonsterData{
			Id:           10000226,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000227: MonsterData{
			Id:           10000227,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000228: MonsterData{
			Id:           10000228,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000229: MonsterData{
			Id:           10000229,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000230: MonsterData{
			Id:           10000230,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000231: MonsterData{
			Id:           10000231,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000232: MonsterData{
			Id:           10000232,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000233: MonsterData{
			Id:           10000233,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000234: MonsterData{
			Id:           10000234,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000235: MonsterData{
			Id:           10000235,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000236: MonsterData{
			Id:           10000236,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000237: MonsterData{
			Id:           10000237,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000238: MonsterData{
			Id:           10000238,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000239: MonsterData{
			Id:           10000239,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000240: MonsterData{
			Id:           10000240,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000241: MonsterData{
			Id:           10000241,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000242: MonsterData{
			Id:           10000242,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000243: MonsterData{
			Id:           10000243,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000244: MonsterData{
			Id:           10000244,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000245: MonsterData{
			Id:           10000245,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000246: MonsterData{
			Id:           10000246,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000247: MonsterData{
			Id:           10000247,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000248: MonsterData{
			Id:           10000248,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000249: MonsterData{
			Id:           10000249,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000250: MonsterData{
			Id:           10000250,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000251: MonsterData{
			Id:           10000251,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000252: MonsterData{
			Id:           10000252,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000253: MonsterData{
			Id:           10000253,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000254: MonsterData{
			Id:           10000254,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000255: MonsterData{
			Id:           10000255,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000256: MonsterData{
			Id:           10000256,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000257: MonsterData{
			Id:           10000257,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000258: MonsterData{
			Id:           10000258,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000259: MonsterData{
			Id:           10000259,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000260: MonsterData{
			Id:           10000260,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000261: MonsterData{
			Id:           10000261,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000262: MonsterData{
			Id:           10000262,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000263: MonsterData{
			Id:           10000263,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000264: MonsterData{
			Id:           10000264,
			Killsvalue:   4,
			Monster_type: 4,
		},
		10000265: MonsterData{
			Id:           10000265,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000266: MonsterData{
			Id:           10000266,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000267: MonsterData{
			Id:           10000267,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000268: MonsterData{
			Id:           10000268,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000269: MonsterData{
			Id:           10000269,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000270: MonsterData{
			Id:           10000270,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000271: MonsterData{
			Id:           10000271,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000272: MonsterData{
			Id:           10000272,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000273: MonsterData{
			Id:           10000273,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000274: MonsterData{
			Id:           10000274,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000275: MonsterData{
			Id:           10000275,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000276: MonsterData{
			Id:           10000276,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000277: MonsterData{
			Id:           10000277,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000278: MonsterData{
			Id:           10000278,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000279: MonsterData{
			Id:           10000279,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000280: MonsterData{
			Id:           10000280,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000281: MonsterData{
			Id:           10000281,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000282: MonsterData{
			Id:           10000282,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000283: MonsterData{
			Id:           10000283,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000284: MonsterData{
			Id:           10000284,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000285: MonsterData{
			Id:           10000285,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000286: MonsterData{
			Id:           10000286,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000287: MonsterData{
			Id:           10000287,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000288: MonsterData{
			Id:           10000288,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000289: MonsterData{
			Id:           10000289,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000290: MonsterData{
			Id:           10000290,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000291: MonsterData{
			Id:           10000291,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000292: MonsterData{
			Id:           10000292,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000293: MonsterData{
			Id:           10000293,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000294: MonsterData{
			Id:           10000294,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000295: MonsterData{
			Id:           10000295,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000296: MonsterData{
			Id:           10000296,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000297: MonsterData{
			Id:           10000297,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000298: MonsterData{
			Id:           10000298,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000299: MonsterData{
			Id:           10000299,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000300: MonsterData{
			Id:           10000300,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000301: MonsterData{
			Id:           10000301,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000302: MonsterData{
			Id:           10000302,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000303: MonsterData{
			Id:           10000303,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000304: MonsterData{
			Id:           10000304,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000305: MonsterData{
			Id:           10000305,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000306: MonsterData{
			Id:           10000306,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000307: MonsterData{
			Id:           10000307,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000308: MonsterData{
			Id:           10000308,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000309: MonsterData{
			Id:           10000309,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000310: MonsterData{
			Id:           10000310,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000311: MonsterData{
			Id:           10000311,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000312: MonsterData{
			Id:           10000312,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000313: MonsterData{
			Id:           10000313,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000314: MonsterData{
			Id:           10000314,
			Killsvalue:   4,
			Monster_type: 4,
		},
		10000315: MonsterData{
			Id:           10000315,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000316: MonsterData{
			Id:           10000316,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000317: MonsterData{
			Id:           10000317,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000318: MonsterData{
			Id:           10000318,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000319: MonsterData{
			Id:           10000319,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000320: MonsterData{
			Id:           10000320,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000321: MonsterData{
			Id:           10000321,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000322: MonsterData{
			Id:           10000322,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000323: MonsterData{
			Id:           10000323,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000324: MonsterData{
			Id:           10000324,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000325: MonsterData{
			Id:           10000325,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000326: MonsterData{
			Id:           10000326,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000327: MonsterData{
			Id:           10000327,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000328: MonsterData{
			Id:           10000328,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000329: MonsterData{
			Id:           10000329,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000330: MonsterData{
			Id:           10000330,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000331: MonsterData{
			Id:           10000331,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000332: MonsterData{
			Id:           10000332,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000333: MonsterData{
			Id:           10000333,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000334: MonsterData{
			Id:           10000334,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000335: MonsterData{
			Id:           10000335,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000336: MonsterData{
			Id:           10000336,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000337: MonsterData{
			Id:           10000337,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000338: MonsterData{
			Id:           10000338,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000339: MonsterData{
			Id:           10000339,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000340: MonsterData{
			Id:           10000340,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000341: MonsterData{
			Id:           10000341,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000342: MonsterData{
			Id:           10000342,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000343: MonsterData{
			Id:           10000343,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000344: MonsterData{
			Id:           10000344,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000345: MonsterData{
			Id:           10000345,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000346: MonsterData{
			Id:           10000346,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000347: MonsterData{
			Id:           10000347,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000348: MonsterData{
			Id:           10000348,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000349: MonsterData{
			Id:           10000349,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000350: MonsterData{
			Id:           10000350,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000351: MonsterData{
			Id:           10000351,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000352: MonsterData{
			Id:           10000352,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000353: MonsterData{
			Id:           10000353,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000354: MonsterData{
			Id:           10000354,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000355: MonsterData{
			Id:           10000355,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000356: MonsterData{
			Id:           10000356,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000357: MonsterData{
			Id:           10000357,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000358: MonsterData{
			Id:           10000358,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000359: MonsterData{
			Id:           10000359,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000360: MonsterData{
			Id:           10000360,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000361: MonsterData{
			Id:           10000361,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000362: MonsterData{
			Id:           10000362,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000363: MonsterData{
			Id:           10000363,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000364: MonsterData{
			Id:           10000364,
			Killsvalue:   4,
			Monster_type: 4,
		},
		10000365: MonsterData{
			Id:           10000365,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000366: MonsterData{
			Id:           10000366,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000367: MonsterData{
			Id:           10000367,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000368: MonsterData{
			Id:           10000368,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000369: MonsterData{
			Id:           10000369,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000370: MonsterData{
			Id:           10000370,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000371: MonsterData{
			Id:           10000371,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000372: MonsterData{
			Id:           10000372,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000373: MonsterData{
			Id:           10000373,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000374: MonsterData{
			Id:           10000374,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000375: MonsterData{
			Id:           10000375,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000376: MonsterData{
			Id:           10000376,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000377: MonsterData{
			Id:           10000377,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000378: MonsterData{
			Id:           10000378,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000379: MonsterData{
			Id:           10000379,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000380: MonsterData{
			Id:           10000380,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000381: MonsterData{
			Id:           10000381,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000382: MonsterData{
			Id:           10000382,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000383: MonsterData{
			Id:           10000383,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000384: MonsterData{
			Id:           10000384,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000385: MonsterData{
			Id:           10000385,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000386: MonsterData{
			Id:           10000386,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000387: MonsterData{
			Id:           10000387,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000388: MonsterData{
			Id:           10000388,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000389: MonsterData{
			Id:           10000389,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000390: MonsterData{
			Id:           10000390,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000391: MonsterData{
			Id:           10000391,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000392: MonsterData{
			Id:           10000392,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000393: MonsterData{
			Id:           10000393,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000394: MonsterData{
			Id:           10000394,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000395: MonsterData{
			Id:           10000395,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000396: MonsterData{
			Id:           10000396,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000397: MonsterData{
			Id:           10000397,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000398: MonsterData{
			Id:           10000398,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000399: MonsterData{
			Id:           10000399,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000400: MonsterData{
			Id:           10000400,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000401: MonsterData{
			Id:           10000401,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000402: MonsterData{
			Id:           10000402,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000403: MonsterData{
			Id:           10000403,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000404: MonsterData{
			Id:           10000404,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000405: MonsterData{
			Id:           10000405,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000406: MonsterData{
			Id:           10000406,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000407: MonsterData{
			Id:           10000407,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000408: MonsterData{
			Id:           10000408,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000409: MonsterData{
			Id:           10000409,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000410: MonsterData{
			Id:           10000410,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000411: MonsterData{
			Id:           10000411,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000412: MonsterData{
			Id:           10000412,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000413: MonsterData{
			Id:           10000413,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000414: MonsterData{
			Id:           10000414,
			Killsvalue:   4,
			Monster_type: 4,
		},
		10000415: MonsterData{
			Id:           10000415,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000416: MonsterData{
			Id:           10000416,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000417: MonsterData{
			Id:           10000417,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000418: MonsterData{
			Id:           10000418,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000419: MonsterData{
			Id:           10000419,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000420: MonsterData{
			Id:           10000420,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000421: MonsterData{
			Id:           10000421,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000422: MonsterData{
			Id:           10000422,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000423: MonsterData{
			Id:           10000423,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000424: MonsterData{
			Id:           10000424,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000425: MonsterData{
			Id:           10000425,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000426: MonsterData{
			Id:           10000426,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000427: MonsterData{
			Id:           10000427,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000428: MonsterData{
			Id:           10000428,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000429: MonsterData{
			Id:           10000429,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000430: MonsterData{
			Id:           10000430,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000431: MonsterData{
			Id:           10000431,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000432: MonsterData{
			Id:           10000432,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000433: MonsterData{
			Id:           10000433,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000434: MonsterData{
			Id:           10000434,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000435: MonsterData{
			Id:           10000435,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000436: MonsterData{
			Id:           10000436,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000437: MonsterData{
			Id:           10000437,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000438: MonsterData{
			Id:           10000438,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000439: MonsterData{
			Id:           10000439,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000440: MonsterData{
			Id:           10000440,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000441: MonsterData{
			Id:           10000441,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000442: MonsterData{
			Id:           10000442,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000443: MonsterData{
			Id:           10000443,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000444: MonsterData{
			Id:           10000444,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000445: MonsterData{
			Id:           10000445,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000446: MonsterData{
			Id:           10000446,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000447: MonsterData{
			Id:           10000447,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000448: MonsterData{
			Id:           10000448,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000449: MonsterData{
			Id:           10000449,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000450: MonsterData{
			Id:           10000450,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000451: MonsterData{
			Id:           10000451,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000452: MonsterData{
			Id:           10000452,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000453: MonsterData{
			Id:           10000453,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000454: MonsterData{
			Id:           10000454,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000455: MonsterData{
			Id:           10000455,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000456: MonsterData{
			Id:           10000456,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000457: MonsterData{
			Id:           10000457,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000458: MonsterData{
			Id:           10000458,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000459: MonsterData{
			Id:           10000459,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000460: MonsterData{
			Id:           10000460,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000461: MonsterData{
			Id:           10000461,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000462: MonsterData{
			Id:           10000462,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000463: MonsterData{
			Id:           10000463,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000464: MonsterData{
			Id:           10000464,
			Killsvalue:   4,
			Monster_type: 4,
		},
		10000465: MonsterData{
			Id:           10000465,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000466: MonsterData{
			Id:           10000466,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000467: MonsterData{
			Id:           10000467,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000468: MonsterData{
			Id:           10000468,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000469: MonsterData{
			Id:           10000469,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000470: MonsterData{
			Id:           10000470,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000471: MonsterData{
			Id:           10000471,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000472: MonsterData{
			Id:           10000472,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000473: MonsterData{
			Id:           10000473,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000474: MonsterData{
			Id:           10000474,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000475: MonsterData{
			Id:           10000475,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000476: MonsterData{
			Id:           10000476,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000477: MonsterData{
			Id:           10000477,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000478: MonsterData{
			Id:           10000478,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000479: MonsterData{
			Id:           10000479,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000480: MonsterData{
			Id:           10000480,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000481: MonsterData{
			Id:           10000481,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000482: MonsterData{
			Id:           10000482,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000483: MonsterData{
			Id:           10000483,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000484: MonsterData{
			Id:           10000484,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000485: MonsterData{
			Id:           10000485,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000486: MonsterData{
			Id:           10000486,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000487: MonsterData{
			Id:           10000487,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000488: MonsterData{
			Id:           10000488,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000489: MonsterData{
			Id:           10000489,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000490: MonsterData{
			Id:           10000490,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000491: MonsterData{
			Id:           10000491,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000492: MonsterData{
			Id:           10000492,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000493: MonsterData{
			Id:           10000493,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000494: MonsterData{
			Id:           10000494,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000495: MonsterData{
			Id:           10000495,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000496: MonsterData{
			Id:           10000496,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000497: MonsterData{
			Id:           10000497,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000498: MonsterData{
			Id:           10000498,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000499: MonsterData{
			Id:           10000499,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000500: MonsterData{
			Id:           10000500,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000501: MonsterData{
			Id:           10000501,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000502: MonsterData{
			Id:           10000502,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000503: MonsterData{
			Id:           10000503,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000504: MonsterData{
			Id:           10000504,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000505: MonsterData{
			Id:           10000505,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000506: MonsterData{
			Id:           10000506,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000507: MonsterData{
			Id:           10000507,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000508: MonsterData{
			Id:           10000508,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000509: MonsterData{
			Id:           10000509,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000510: MonsterData{
			Id:           10000510,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000511: MonsterData{
			Id:           10000511,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000512: MonsterData{
			Id:           10000512,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000513: MonsterData{
			Id:           10000513,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000514: MonsterData{
			Id:           10000514,
			Killsvalue:   4,
			Monster_type: 4,
		},
		10000515: MonsterData{
			Id:           10000515,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000516: MonsterData{
			Id:           10000516,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000517: MonsterData{
			Id:           10000517,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000518: MonsterData{
			Id:           10000518,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000519: MonsterData{
			Id:           10000519,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000520: MonsterData{
			Id:           10000520,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000521: MonsterData{
			Id:           10000521,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000522: MonsterData{
			Id:           10000522,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000523: MonsterData{
			Id:           10000523,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000524: MonsterData{
			Id:           10000524,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000525: MonsterData{
			Id:           10000525,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000526: MonsterData{
			Id:           10000526,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000527: MonsterData{
			Id:           10000527,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000528: MonsterData{
			Id:           10000528,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000529: MonsterData{
			Id:           10000529,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000530: MonsterData{
			Id:           10000530,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000531: MonsterData{
			Id:           10000531,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000532: MonsterData{
			Id:           10000532,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000533: MonsterData{
			Id:           10000533,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000534: MonsterData{
			Id:           10000534,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000535: MonsterData{
			Id:           10000535,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000536: MonsterData{
			Id:           10000536,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000537: MonsterData{
			Id:           10000537,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000538: MonsterData{
			Id:           10000538,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000539: MonsterData{
			Id:           10000539,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000540: MonsterData{
			Id:           10000540,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000541: MonsterData{
			Id:           10000541,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000542: MonsterData{
			Id:           10000542,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000543: MonsterData{
			Id:           10000543,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000544: MonsterData{
			Id:           10000544,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000545: MonsterData{
			Id:           10000545,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000546: MonsterData{
			Id:           10000546,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000547: MonsterData{
			Id:           10000547,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000548: MonsterData{
			Id:           10000548,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000549: MonsterData{
			Id:           10000549,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000550: MonsterData{
			Id:           10000550,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000551: MonsterData{
			Id:           10000551,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000552: MonsterData{
			Id:           10000552,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000553: MonsterData{
			Id:           10000553,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000554: MonsterData{
			Id:           10000554,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000555: MonsterData{
			Id:           10000555,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000556: MonsterData{
			Id:           10000556,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000557: MonsterData{
			Id:           10000557,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000558: MonsterData{
			Id:           10000558,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000559: MonsterData{
			Id:           10000559,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000560: MonsterData{
			Id:           10000560,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000561: MonsterData{
			Id:           10000561,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000562: MonsterData{
			Id:           10000562,
			Killsvalue:   1,
			Monster_type: 1,
		},
		10000563: MonsterData{
			Id:           10000563,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000564: MonsterData{
			Id:           10000564,
			Killsvalue:   4,
			Monster_type: 4,
		},
		10000565: MonsterData{
			Id:           10000565,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000566: MonsterData{
			Id:           10000566,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000567: MonsterData{
			Id:           10000567,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000568: MonsterData{
			Id:           10000568,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000569: MonsterData{
			Id:           10000569,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000570: MonsterData{
			Id:           10000570,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000571: MonsterData{
			Id:           10000571,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000572: MonsterData{
			Id:           10000572,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000573: MonsterData{
			Id:           10000573,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000574: MonsterData{
			Id:           10000574,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000575: MonsterData{
			Id:           10000575,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000576: MonsterData{
			Id:           10000576,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000577: MonsterData{
			Id:           10000577,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000578: MonsterData{
			Id:           10000578,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000579: MonsterData{
			Id:           10000579,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000580: MonsterData{
			Id:           10000580,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000581: MonsterData{
			Id:           10000581,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000582: MonsterData{
			Id:           10000582,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000583: MonsterData{
			Id:           10000583,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000584: MonsterData{
			Id:           10000584,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000585: MonsterData{
			Id:           10000585,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000586: MonsterData{
			Id:           10000586,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000587: MonsterData{
			Id:           10000587,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000588: MonsterData{
			Id:           10000588,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000589: MonsterData{
			Id:           10000589,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000590: MonsterData{
			Id:           10000590,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000591: MonsterData{
			Id:           10000591,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000592: MonsterData{
			Id:           10000592,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000593: MonsterData{
			Id:           10000593,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000594: MonsterData{
			Id:           10000594,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000595: MonsterData{
			Id:           10000595,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000596: MonsterData{
			Id:           10000596,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000597: MonsterData{
			Id:           10000597,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000598: MonsterData{
			Id:           10000598,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000599: MonsterData{
			Id:           10000599,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000600: MonsterData{
			Id:           10000600,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000601: MonsterData{
			Id:           10000601,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000602: MonsterData{
			Id:           10000602,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000603: MonsterData{
			Id:           10000603,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000604: MonsterData{
			Id:           10000604,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000605: MonsterData{
			Id:           10000605,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000606: MonsterData{
			Id:           10000606,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000607: MonsterData{
			Id:           10000607,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000608: MonsterData{
			Id:           10000608,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000609: MonsterData{
			Id:           10000609,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000610: MonsterData{
			Id:           10000610,
			Killsvalue:   2,
			Monster_type: 2,
		},
		10000611: MonsterData{
			Id:           10000611,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000612: MonsterData{
			Id:           10000612,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000613: MonsterData{
			Id:           10000613,
			Killsvalue:   3,
			Monster_type: 3,
		},
		10000614: MonsterData{
			Id:           10000614,
			Killsvalue:   4,
			Monster_type: 4,
		},
		10000615: MonsterData{
			Id:           10000615,
			Killsvalue:   0,
			Monster_type: 5,
		},
		10000616: MonsterData{
			Id:           10000616,
			Killsvalue:   0,
			Monster_type: 6,
		},
		10000617: MonsterData{
			Id:           10000617,
			Killsvalue:   0,
			Monster_type: 7,
		},
		10000618: MonsterData{
			Id:           10000618,
			Killsvalue:   0,
			Monster_type: 8,
		},
		10000619: MonsterData{
			Id:           10000619,
			Killsvalue:   20,
			Monster_type: 1,
		},
		10000620: MonsterData{
			Id:           10000620,
			Killsvalue:   20,
			Monster_type: 1,
		},
		10000621: MonsterData{
			Id:           10000621,
			Killsvalue:   4,
			Monster_type: 4,
		},
	}
}
