package payment

import (
	"context"

	jsoniter "github.com/json-iterator/go"
	"gitlab.com/megatech/serverex/apps/excalibur/payment/playstore"
	"gitlab.com/megatech/serverex/lib/log"
)

type GooglePurchaseData struct {
	OrderId       string `json:"orderId"`
	PackageName   string `json:"packageName"`
	ProductId     string `json:"productId"`
	PurchaseTime  int64  `json:"purchaseTime"`
	PurchaseState int    `json:"purchaseState"`
	PurchaseToken string `json:"purchaseToken"`
}

func UnmarshalGooglePurchaseData(s string) (error, GooglePurchaseData) {

	var data GooglePurchaseData
	err := jsoniter.UnmarshalFromString(s, &data)
	return err, data
}

const key = `{
	"type": "service_account",
	"project_id": "api-5646349018931174637-672120",
	"private_key_id": "8ccee400ee7199735a1a78e3bf10cf01d1933151",
	"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDD5uPfs/aMrI9F\n0od+zcOtRYuMP3pWrrWQ9vLOtGAIDYLEUBduvFrpMAwWW/dgP5fY+gXVRqye+78y\ngO/pSBCxudNPQsJj7M10rZmlAILrjnZY0F+qa4g4+S3e/fC8uVKFErWbbHGOx34n\n5xu2VUh6OuSobSzE1e0OAU2wY/G/R151hHBPPzc916wjUG+4aXKSshrjmAasx2fH\nO0RzibxNYkApfSl35nlOTVkW4jTXWZRk17d8Uzl56/aNrhR0QIpAKdyhTINp9SxT\ney3X235h3gjBjONZtW/B4YHvguHZQrloKDtvJJmg4SFXabzvAvaDGqlPtFhpdShW\nxRu13QlDAgMBAAECggEAQAVRPBa6gJ+WFOnn6OtCwvoC0FsgVnqzEIR7tkyhQ5By\nx0grWD+K8EEI+jek4Nn3fGu9A0Wc/TnlcpJLlS+9cO97HKCDGZfWTwa6uv8vSd7d\nU0oJsY6Bzqi0MndHuM1Rp/aEzh3DALe2RUEUP+KgFH3xow9Gva7x0Q1g/7+xzeHo\nYzHijybAHKb4+jed8Qk3tMn9U7wubMtpt1q7jxVLIqF9/E8if8BjRWV5vtkFXCLM\n2Cd9cWUkYpz720elFoiKHHdurDlaHDemKAaeCVKf98UgQ8sFvM5xEF+U1x7v3mI1\nvIWk+FWyepNWVezSlhjugfvGxGeRXk/xjQ1jhGCT7QKBgQD3lK+q5hdyoHoW09IO\nnrFrDrQ8jAM72gt7YFn0H9TgwjA6z893MwwVRcAwAFMt+4wKhiNHmzQSI4DS+O1v\nqIVoAGQ1Fa8OPPc69OJ31eG5jZh9xo5BatCEgg431rYas5mcAckop6t8vxeavuxc\nDDrjnoZAGWfLgpZet+ChRWpb3QKBgQDKkFBDzj+wQROfnJIIqmBRBoa9N9WZQtFc\nWVgaJyXmJA2D+aFHl1jT4PRlP7cSgsCWf6hmPByXp8bv6vA/XKPGtUxCfoIr3j5Y\npV20fASqmA1XGBitbAIpwkhsu9nuoSaEQ2YW+xn1XlynhEH3ew33f3w0F21DFqds\ndEycUP+3nwKBgGRPWNFcQlnODUlcbzo/L9EZdwslC1ZfpNFLG4cesoiSBSp+8ibw\nhs/Dn/eS2iDCT14gNQEZMAK9Yb+Y3dKkq9CUgAVGENQq/VxkxbEZ9kfQx90F6rM0\njOFzNEGJe+TbqbwcuR0Auilj3zNycZ9l0SJSI0CrSnOZKEWeKKQO1OX9AoGAN3JH\nNjg3gLRffboq0wpn8OXbbeHhquD6U8/06Lu8iQFnXa+v8NvBqcWSwEAviXaW5tss\nqkdcADshnjxt5Loj7llj9XORiBZ+dF9XhfbIhDW4uRVUc+vzr6CJmJHW9pXP6DZW\ntIWpNlvM+uGd+PMLAKs5wTyEYO/LhklTSxp01LsCgYBPHEbRZpkWxPKwaHLk2V3I\ncrZPVUlHnfPiA9AYNpqgIWN6rqXtccmDuQ5Af2bV5L8Hf3QDvWleGZThjDMw9tr/\nd38+ece+vHGpQQkde3Jk6GKoocrtX/dAVGbqmf4F7fvdacgUV+ll0nDkwzqbxcw9\ngHnzJyX1IEmrcf/DT6ZN5A==\n-----END PRIVATE KEY-----\n",
	"client_email": "test1-86@api-5646349018931174637-672120.iam.gserviceaccount.com",
	"client_id": "113838427568685064693",
	"auth_uri": "https://accounts.google.com/o/oauth2/auth",
	"token_uri": "https://accounts.google.com/o/oauth2/token",
	"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
	"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/test1-86%40api-5646349018931174637-672120.iam.gserviceaccount.com"
  }

`

func GoogleValidateToken(packageName string, productID string, token string) error {
	client, err := playstore.New([]byte(key))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	resp, err := client.VerifyProduct(ctx, packageName, productID, token)

	if err != nil {
		log.Error(err)
	}
	log.Info(resp)
	return err
}
