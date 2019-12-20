package payment

// Taken from `github.com/ascepanovic/receiptvalidator`

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"gitlab.com/megatech/serverex/lib/log"
)

// Receipt is information returned by Apple
// Documentation: https://developer.apple.com/library/ios/releasenotes/General/ValidateAppStoreReceipt/Chapters/ReceiptFields.html#//apple_ref/doc/uid/TP40010573-CH106-SW10
// type Receipt struct {
// 	ReceiptType                string            `json:"receipt_type"`
// 	AdamID                     int               `json:"adam_id"`
// 	AppItemID                  int               `json:"app_item_id"`
// 	BundleID                   string            `json:"bundle_id"`
// 	ApplicationVersion         string            `json:"application_version"`
// 	DownloadID                 int               `json:"download_id"`
// 	VersionExternalIdentifier  string            `json:"version_external_identifier"` // int ? string?
// 	ReceiptCreationDate        string            `json:"receipt_creation_date"`
// 	ReceiptCreationDateMs      string            `json:"receipt_creation_date_ms"`
// 	ReceiptCreationDatePst     string            `json:"receipt_creation_date_pst"`
// 	OriginalPurchaseDate       string            `json:"original_purchase_date"`
// 	OriginalPurchaseDateMs     string            `json:"original_purchase_date_ms"`
// 	OriginalPurchaseDatePst    string            `json:"original_purchase_date_pst"`
// 	OriginalApplicationVersion string            `json:"original_application_version"`
// 	InApp                      []PurchaseReceipt `json:"in_app"`
// }

//PurchaseReceipt CHECK OUT what are required fields and so on
// type PurchaseReceipt struct {
// 	Quantity                string `json:"quantity"`
// 	ProductID               string `json:"product_id"`
// 	TransactionID           string `json:"transaction_id"`
// 	OriginalTransactionID   string `json:"original_transaction_id"`
// 	PurchaseDate            string `json:"purchase_date"`
// 	PurchaseDateMs          string `json:"purchase_date_ms"`
// 	PurchaseDatePst         string `json:"purchase_date_pst"`
// 	OriginalPurchaseDate    string `json:"original_purchase_date"`
// 	OriginalPurchaseDateMs  string `json:"original_purchase_date_ms"`
// 	OriginalPurchaseDatePst string `json:"original_purchase_date_pst"`
// 	IsTrialPeriod           string `json:"is_trial_period"`
// }

type receiptRequestData struct {
	Receiptdata string `json:"receipt-data"`
}

const (
	appleSandboxURL    string = "https://sandbox.itunes.apple.com/verifyReceipt"
	appleProductionURL string = "https://buy.itunes.apple.com/verifyReceipt"
)

//Error basic struct
type Error struct {
	error
}

//VerifyReceipt Given receiptData (base64 encoded) it tries to connect to either the sandbox (useSandbox true) or apples ordinary service (useSandbox false) to validate the receipt. Returns either a receipt struct or an error.
func VerifyReceipt(receiptData string, useSandbox bool) (*Receipt, error) {
	receipt, err := sendReceiptToApple(receiptData, verificationURL(useSandbox))
	return receipt, err
}

// Selects the proper url to use when talking to apple based on if we should use the sandbox environment or not
func verificationURL(useSandbox bool) string {

	if useSandbox {
		return appleSandboxURL
	}
	return appleProductionURL
}

// 完全是根据返回的结果推断的结构。因为苹果的文档对不上。
type Receipt struct {
	A             string `json:"original_purchase_date_pst"`
	B             string `json:"purchase_date_ms"`
	C             string `json:"unique_identifier"`
	D             string `json:"original_transaction_id"`
	E             string `json:"bvrs"`
	TransactionID string `json:"transaction_id"`
	G             string `json:"quantity"`
	H             string `json:"unique_vendor_identifier"`
	I             string `json:"item_id"`
	J             string `json:"version_external_identifier"`
	K             string `json:"bid"`
	L             string `json:"is_in_intro_offer_period"`
	ProductID     string `json:"product_id"`
	N             string `json:"purchase_date"`
	O             string `json:"is_trial_period"`
	P             string `json:"purchase_date_pst"`
	Q             string `json:"original_purchase_date"`
	R             string `json:"original_purchase_date_ms"`
}

// Sends the receipt to apple, returns the receipt or an error upon completion
func sendReceiptToApple(receiptData, url string) (*Receipt, error) {
	requestData, err := json.Marshal(receiptRequestData{receiptData})

	if err != nil {
		return nil, err
	}

	toSend := bytes.NewBuffer(requestData)

	resp, err := http.Post(url, "application/json", toSend)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	var responseData struct {
		Status         float64  `json:"status"`
		ReceiptContent *Receipt `json:"receipt"`
	}

	responseData.ReceiptContent = new(Receipt)

	log.Info(string(body))

	err = json.Unmarshal(body, &responseData)

	if err != nil {
		return nil, err
	}

	if responseData.Status != 0 {
		return nil, verificationError(responseData.Status)
	}

	return responseData.ReceiptContent, nil
}

// Generates the correct error based on a status error code
func verificationError(errCode float64) error {
	var errorMessage string

	switch errCode {
	case 21000:
		errorMessage = "The App Store could not read the JSON object you provided."
		break
	case 21002:
		errorMessage = "The data in the receipt-data property was malformed."
		break

	case 21003:
		errorMessage = "The receipt could not be authenticated."
		break

	case 21004:
		errorMessage = "The shared secret you provided does not match the shared secret on file for your account."
		break

	case 21005:
		errorMessage = "The receipt server is not currently available."
		break
	case 21006:
		errorMessage = "This receipt is valid but the subscription has expired. When this status code is returned to your server, " +
			"the receipt data is also decoded and returned as part of the response."
		break
	case 21007:
		errorMessage = "This receipt is a sandbox receipt, but it was sent to the production service for verification."
		break
	case 21008:
		errorMessage = "This receipt is a production receipt, but it was sent to the sandbox service for verification."
		break
	default:
		errorMessage = "An unknown error ocurred."
		break
	}

	return &Error{errors.New(errorMessage)}
}
