package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://www.dhgate.com/product/design-ggity-socks-for-women-sexy-letter/926012915.html"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "www.dhgate.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "csrfToken=VV_lxhF1iyXuAxZsoAwDLuPp; language=en; vid=820A13ACD9DFEE65D15DB24E02FC1004; b2b_ip_country=US; b2b_ship_country=US; last_choice=0; ref_df=direct; odvid=rBMKgmXu39lOsl3RBBD8Ag==")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
