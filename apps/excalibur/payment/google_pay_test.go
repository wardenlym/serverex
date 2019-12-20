package payment

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"testing"

	"gitlab.com/megatech/serverex/lib/log"

	"gitlab.com/megatech/serverex/apps/excalibur/payment/playstore"
)

// Google Play Android Developer
// 客户端 ID
// 939560867272-ttvgnuc2uupvd66bleiq5lhut3gk171c.apps.googleusercontent.com
// 客户端密钥
// VlCwNQ1mYsVgS_3zfI05NVqE

const base64JsonKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAscnqXtxQSBjwoMZRusJMz8MhgtGCLhjO/yhZmdPeGoRIsefO1C6HsKFVM71vIxcGhoeGyrziUYGP6O3BfpWBcvHn9VlubgrJ6tqx0hYRedO5dwg+YU15Km/D4TrptgqDp9U4dhWvUqxfbjyeiUJY/qE13FEdTlmX65yqb4Gz029O7UCTx4P696QMJ/h/xBP9PGec1KEW3wJutx0qTPMfS/U9TgTOQlXe/h/v3T4OwPsko5TU2j200ZFP4z9wOLIgdveW2Xo6mrX8lESVf4MKU1w54EbikqbYnjWswCjtmjuMK+Ud1+PN2r7Ye1eX6xBK91QLWE5+YGoSmNIqmB6PcwIDAQAB"

var jsonKey []byte

const PackageName = "com.sdbean.excalibur"
const ProductID = "android.test.purchased"
const PurchaseToken = "inapp:com.sdbean.excalibur:android.test.purchased"

func init() {
	f, err := base64.StdEncoding.DecodeString(base64JsonKey)
	if err != nil {
		panic(err)
	}
	jsonKey = f

}

func Test_b(t *testing.T) {

	jsonKey, err := ioutil.ReadFile("Google Play Android Developer-8ccee400ee71.json")
	if err != nil {
		log.Fatal(err)
	}
	client, err := playstore.New(jsonKey)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	token := "eicpikhhpmeiendabogkcdfh.AO-J1OzUntXBM45Zw2mqPgeVeBsGCnQZI8D76y2AM7XfhV64iclYxme81D-YHAV05xwf23sK6_t3Ay-6nMnLH0fogZqq1UUzdjrryDWeLf3TnXEq2Oxf6OozTFeWoMUYX167qerPRYBx"
	resp, err := client.VerifyProduct(ctx, PackageName, "com.sdbean.excalibur.test", token)

	if err != nil {
		log.Error(err)
	}
	log.Info(resp)
}
