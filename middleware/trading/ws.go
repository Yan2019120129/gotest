package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gotest/common/utils"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/nacl/secretbox"
)

const (
	httpPath   = "https://live.tradingeconomics.com/socket.io"
	wsPath     = "wss://live.tradingeconomics.com/socket.io"
	teDecryptK = "j9ONifjoKzxt7kmfYTdKK/5vve0b9Y1UCj/n50jr8d8="
	teDecryptN = "Ipp9HNSfVBUntqFK7PrtofYaOPV312xy"
)

func Ws() {
	params := Params{
		Key:       "rain",
		Url:       "/",
		EIO:       "4",
		Transport: "websocket",
		T:         fmt.Sprintf("%d", time.Now().UnixNano()/1e6),
	}
	sid := GetSid(params)
	if sid == "" {
		panic("sid is null")
	}

	tmp := []any{"subscribe", map[string][]string{"s": {
		"jpyusd:cur",
	}}}

	param := url.Values{}
	param.Set("key", params.Key)
	param.Set("url", params.Url)
	param.Set("EIO", params.EIO)
	param.Set("transport", "websocket")
	param.Set("sid", sid)
	instance := utils.NewWs(wsPath + "/?" + param.Encode())
	err := instance.Run().Err
	if err != nil {
		fmt.Println("-------------", err)
		return
	}
	time.Sleep(2 * time.Second)
	instance.Send("2probe")
	instance.Read(func(bytes []byte) {
		if len(bytes) < 46 {
			numb := extractNumbers(bytes)
			if numb == 3 {
				instance.Send("5")
				instance.Send("40")
				fmt.Println(string(bytes))
				return
			}
			if numb == 2 {
				instance.Send("3")
				fmt.Println(string(bytes))
				return
			}

			if numb == 40 {
				instance.Send("42" + utils.ObjToString(tmp))
				fmt.Println(string(bytes))
				return
			}
			if numb == 451 {
				fmt.Println(string(bytes))
				return
			}
		}

		v, err := DecryptMessage(bytes, teDecryptK, teDecryptN)
		if err != nil {
			fmt.Println("-------------", err)
		}
		commodityData := CommodityData{}
		_ = json.Unmarshal(v, &commodityData)
		fmt.Printf("%+v:open:%v\n", commodityData, commodityData.P-commodityData.Nch)
	})
}

func GetSid(params Params) string {
	instance := utils.NewHttp()
	instance.AddParam("key", params.Key)
	instance.AddParam("url", params.Url)
	instance.AddParam("EIO", params.EIO)
	instance.AddParam("transport", "polling")
	instance.AddParam("t", params.T)
	val := instance.Get(httpPath + "/")
	response := WsResponse{}
	if len(val) > 0 {
		err := json.Unmarshal(val[1:], &response)
		if err != nil {
			panic(err)
		}
		return response.SID
	}

	return ""
}

// Base64ToBytes decodes a Base64-encoded string into a byte slice.
func Base64ToBytes(encoded string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// DecryptMessage decrypts a ciphertext using libsodium's secretbox and decompresses the result.
func DecryptMessage(ciphertext []byte, keyBase64 string, nonceBase64 string) ([]byte, error) {
	// Decode the key and nonce from Base64
	key, err := Base64ToBytes(keyBase64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode key")
	}
	nonce, err := Base64ToBytes(nonceBase64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode nonce")
	}

	// Convert the ciphertext into a uint8 array
	var nonceArray [24]byte
	copy(nonceArray[:], nonce)

	// Decrypt the message
	decrypted, ok := secretbox.Open(nil, ciphertext, &nonceArray, (*[32]byte)(key))
	if !ok {
		return nil, fmt.Errorf("decryption failed:%v", string(ciphertext))
	}

	// Decompress the decrypted plaintext
	decompressed, err := DecompressZlib(decrypted)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decompress plaintext")
	}

	return decompressed, nil
}

// DecompressZlib decompresses data using zlib.
func DecompressZlib(data []byte) ([]byte, error) {
	reader, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	var decompressed bytes.Buffer
	_, err = decompressed.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	return decompressed.Bytes(), nil
}

// extractNumbers 提取数字
func extractNumbers(b []byte) int {
	var val []byte
	for _, v := range b {
		if 48 <= v && v <= 57 {
			val = append(val, v)
			continue
		}
		break
	}
	parseInt, err := strconv.ParseInt(string(val), 10, 64)
	if err != nil {
		return -1
	}
	return int(parseInt)
}
