package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var (
	testFilename4            = "M.1608388506.A.85D"
	testContentAll4          []byte
	testContent4             []byte
	testSignature4           []byte
	testComment4             []byte
	testFirstCommentsDBCS4   []byte
	testTheRestCommentsDBCS4 []byte
	testContent4Big5         [][]*types.Rune
	testContent4Utf8         [][]*types.Rune

	testFirstComments4 []*schema.Comment
)

func initTest4() {
	testContentAll4, testContent4, testSignature4, testComment4, testFirstCommentsDBCS4, testTheRestCommentsDBCS4 = loadTest(testFilename4)

	testContent4Big5 = [][]*types.Rune{
		{
			{
				Big5:   []byte("\xa7@\xaa\xcc: SYSOP () \xac\xdd\xaaO: WhoAmI"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{
			{
				Big5:   []byte("\xbc\xd0\xc3D: [\xb6\xa2\xb2\xe1] \xa9\xd2\xa5H\xafS\xae\xed\xa6r\xafu\xaa\xba\xacO\xa6\xb3\xba\xf1\xa6\xe2\xaa\xba\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{
			{
				Big5:   []byte("\xae\xc9\xb6\xa1: Sat Dec 19 22:35:04 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{},
		{
			{
				Big5:   []byte("\xb5M\xab\xe1 \\n \xa4\xa3\xb7|\xa6b big5 \xb5\xb2\xa7\xc0. \xa5i\xa5H\xa9\xf1\xa4\xdf\xaa\xbd\xb1\xb5\xa5\xce \\n \xc2_\xa6\xe6."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{
			{
				Big5:   []byte("\xa7\xda\xacO\xb3\\\xa5\\\xbb\\"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{},
		{
			{
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
	}

	testContent4Utf8 = [][]*types.Rune{
		{
			{
				Utf8:   "作者: SYSOP () 看板: WhoAmI",
				Big5:   []byte("\xa7@\xaa\xcc: SYSOP () \xac\xdd\xaaO: WhoAmI"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{
			{
				Utf8:   "標題: [閒聊] 所以特殊字真的是有綠色的～",
				Big5:   []byte("\xbc\xd0\xc3D: [\xb6\xa2\xb2\xe1] \xa9\xd2\xa5H\xafS\xae\xed\xa6r\xafu\xaa\xba\xacO\xa6\xb3\xba\xf1\xa6\xe2\xaa\xba\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{
			{
				Utf8:   "時間: Sat Dec 19 22:35:04 2020",
				Big5:   []byte("\xae\xc9\xb6\xa1: Sat Dec 19 22:35:04 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{},
		{
			{
				Utf8:   "然後 \\n 不會在 big5 結尾. 可以放心直接用 \\n 斷行.",
				Big5:   []byte("\xb5M\xab\xe1 \\n \xa4\xa3\xb7|\xa6b big5 \xb5\xb2\xa7\xc0. \xa5i\xa5H\xa9\xf1\xa4\xdf\xaa\xbd\xb1\xb5\xa5\xce \\n \xc2_\xa6\xe6."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{
			{
				Utf8:   "我是許功蓋",
				Big5:   []byte("\xa7\xda\xacO\xb3\\\xa5\\\xbb\\"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{},
		{
			{
				Utf8:   "--",
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
	}

	testFirstComments4 = []*schema.Comment{
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYDNAQlkyAA:3dK46zmOe5zmna12AC1gnQ"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("SYSOP"),
			CreateTime: types.NanoTS(1261233300000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "推",
						Big5:   []byte("\xb1\xc0                                                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "3dK46zmOe5zmna12AC1gnQ",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYDNDwGsIAA:FQaNH8WkdAbEGD7yp2Zkvg"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("chhsiao123"),
			CreateTime: types.NanoTS(1261233360000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "推",
						Big5:   []byte("\xb1\xc0                                                  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "FQaNH8WkdAbEGD7yp2Zkvg",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYDNHPnzeAA:cLGi8fC4fapuiBkTXHU2OA"),
			TheType:    types.COMMENT_TYPE_BOO,
			Owner:      bbs.UUserID("chhsiao123"),
			CreateTime: types.NanoTS(1261233420000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "噓～",
						Big5:   []byte("\xbcN\xa1\xe3                                                "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "cLGi8fC4fapuiBkTXHU2OA",
		},
	}
}
