package monitor

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	APIKey     string
	SecretKey  string
	Passphrase string
}

func NewClient(apiKey, secret, passphrase string) *Client {
	return &Client{
		BaseURL:    "https://www.okx.com",
		APIKey:     apiKey,
		SecretKey:  secret,
		Passphrase: passphrase,
	}
}

func (c *Client) sign(timestamp, method, path, body string) string {
	message := timestamp + method + path + body

	mac := hmac.New(sha256.New, []byte(c.SecretKey))
	mac.Write([]byte(message))

	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

type BalanceResponse struct {
	Code string `json:"code"`
	Data []struct {
		AdjEq              string `json:"adjEq"`
		AvailEq            string `json:"availEq"`
		BorrowFroz         string `json:"borrowFroz"`
		Delta              string `json:"delta"`
		DeltaLever         string `json:"deltaLever"`
		DeltaNeutralStatus string `json:"deltaNeutralStatus"`
		Details            []struct {
			AutoLendStatus        string `json:"autoLendStatus"`
			AutoLendMtAmt         string `json:"autoLendMtAmt"`
			AvailBal              string `json:"availBal"`
			AvailEq               string `json:"availEq"`
			BorrowFroz            string `json:"borrowFroz"`
			CashBal               string `json:"cashBal"`
			Ccy                   string `json:"ccy"`
			CrossLiab             string `json:"crossLiab"`
			ColRes                string `json:"colRes"`
			CollateralEnabled     bool   `json:"collateralEnabled"`
			CollateralRestrict    bool   `json:"collateralRestrict"`
			ColBorrAutoConversion string `json:"colBorrAutoConversion"`
			DisEq                 string `json:"disEq"`
			Eq                    string `json:"eq"`
			EqUsd                 string `json:"eqUsd"`
			SmtSyncEq             string `json:"smtSyncEq"`
			SpotCopyTradingEq     string `json:"spotCopyTradingEq"`
			FixedBal              string `json:"fixedBal"`
			FrozenBal             string `json:"frozenBal"`
			FrpType               string `json:"frpType"`
			Imr                   string `json:"imr"`
			Interest              string `json:"interest"`
			IsoEq                 string `json:"isoEq"`
			IsoLiab               string `json:"isoLiab"`
			IsoUpl                string `json:"isoUpl"`
			Liab                  string `json:"liab"`
			MaxLoan               string `json:"maxLoan"`
			MgnRatio              string `json:"mgnRatio"`
			Mmr                   string `json:"mmr"`
			NotionalLever         string `json:"notionalLever"`
			OrdFrozen             string `json:"ordFrozen"`
			RewardBal             string `json:"rewardBal"`
			SpotInUseAmt          string `json:"spotInUseAmt"`
			ClSpotInUseAmt        string `json:"clSpotInUseAmt"`
			MaxSpotInUse          string `json:"maxSpotInUse"`
			SpotIsoBal            string `json:"spotIsoBal"`
			StgyEq                string `json:"stgyEq"`
			Twap                  string `json:"twap"`
			UTime                 string `json:"uTime"`
			Upl                   string `json:"upl"`
			UplLiab               string `json:"uplLiab"`
			SpotBal               string `json:"spotBal"`
			OpenAvgPx             string `json:"openAvgPx"`
			AccAvgPx              string `json:"accAvgPx"`
			SpotUpl               string `json:"spotUpl"`
			SpotUplRatio          string `json:"spotUplRatio"`
			TotalPnl              string `json:"totalPnl"`
			TotalPnlRatio         string `json:"totalPnlRatio"`
		} `json:"details"`
		Imr                   string `json:"imr"`
		IsoEq                 string `json:"isoEq"`
		MgnRatio              string `json:"mgnRatio"`
		Mmr                   string `json:"mmr"`
		NotionalUsd           string `json:"notionalUsd"`
		NotionalUsdForBorrow  string `json:"notionalUsdForBorrow"`
		NotionalUsdForFutures string `json:"notionalUsdForFutures"`
		NotionalUsdForOption  string `json:"notionalUsdForOption"`
		NotionalUsdForSwap    string `json:"notionalUsdForSwap"`
		OrdFroz               string `json:"ordFroz"`
		TotalEq               string `json:"totalEq"`
		UTime                 string `json:"uTime"`
		Upl                   string `json:"upl"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func getTimeNow() string {
	return time.Now().
		UTC().
		Format("2006-01-02T15:04:05.000Z")
}
func (c *Client) GetAccountBalance(ccy string) (*BalanceResponse, error) {

	path := "/api/v5/account/balance"

	// query参数
	if ccy != "" {
		path = fmt.Sprintf("%s?ccy=%s", path, ccy)
	}
	timestamp := getTimeNow()
	method := "GET"
	body := ""

	sign := c.sign(timestamp, method, path, body)

	req, err := http.NewRequest(method, c.BaseURL+path, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("OK-ACCESS-KEY", c.APIKey)
	req.Header.Set("OK-ACCESS-SIGN", sign)
	req.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("OK-ACCESS-PASSPHRASE", c.Passphrase)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)

	var result BalanceResponse
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
