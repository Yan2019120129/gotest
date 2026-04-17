package main

import (
	"fmt"
	crawling2 "gotest/practice/crawling_t/crawling"
	"gotest/practice/crawling_t/crawling/amazon"
	"net/url"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//var addrs = []string{
//	"https://www.amazon.com/s?k=Essentials%2CFEAR+OF+GOD+ESSENTIALS%2CBoxy+Hoodie&__mk_zh_CN=%E4%BA%9A%E9%A9%AC%E9%80%8A%E7%BD%91%E7%AB%99&crid=IFA6BWCH8JRH&sprefix=essentials%2Cfear+of+god+essentials%2Cboxy+hoodie%2Caps%2C6355&ref=nb_sb_noss",
//	"https://www.amazon.com/dp/B072YVWBXH/?mr_donotredirect=undefined&th=1",
//	"https://www.amazon.com/s?i=fashion-mens-intl-ship&bbn=16225019011&rh=n%3A7141123011%2Cn%3A16225019011%2Cn%3A7147441011%2Cn%3A1040658&dc&language=zh&ds=v1%3A0qMeFWgElNp7zKRHk7%2FFniqSLbOE3bo%2FlTElqN3vB3s&qid=1716639236&ref=sr_ex_n_1",
//	"https://www.amazon.com/s?k=outdoor+sports&i=fashion-mens&__mk_es_US=%C3%85M%C3%85%C5%BD%C3%95%C3%91&crid=21PNQMXB0BHU0&sprefix=outdoor+sports+exercise+pedal%2Cfashion-mens%2C1163&ref=nb_sb_noss_1",
//	"https://www.amazon.com/-/en/Wrangler-Authentics-Pantalones-el%C3%A1sticos-polvorienta/dp/B07YP9LTYF/?_encoding=UTF8&pd_rd_w=dV7JV&content-id=amzn1.sym.9929d3ab-edb7-4ef5-a232-26d90f828fa5&pf_rd_p=9929d3ab-edb7-4ef5-a232-26d90f828fa5&pf_rd_r=ET1V3PPQ0DE8ZJV0PAWR&pd_rd_wg=NZmEe&pd_rd_r=9b6df6db-31d4-4ee1-981a-c2c59a9f8e3f&ref_=pd_hp_d_btf_crs_zg_bs_7141123011",
//	"https://www.amazon.com/-/en/interactivo-mascotas-electr%C3%B3nicas-virtuales-reaccionan/dp/B0BW1B31SR/ref=sr_1_5?_encoding=UTF8&content-id=amzn1.sym.44da4965-9668-4613-bec2-a3a75f0c2ad4&dib=eyJ2IjoiMSJ9.b8n9r6_6RGcDrh2OyrQ8rq9KqGl45pWLWgfqQLxww4g6NMQf59WWhrc3edi2tgBpAzI-KqtZSctfqBqqgIlRsx2SErh1Y4n2mWaQKdREXS2sHn_RNFIbIl78zmd6k05Cx9nmTtGag-_eCJNZvINdx7wdXC_MPz9eIDaZMbNMH5jdldvLyvtB74dNGB2XTdqFXcNg1m7kIZWw7KR4RbLOJsTaflpTZGpB6vWP5rAHFr9Yhs2lGm34m-7dysAz-18FHTXWikcUmHPLrBfRbLvyt5IkNVeLvk1V5r9r6yR2Vc8.iKFYnrMBMKtFMx0lTFIGSeUQIv8wdIEkTIs-IoDIW8E&dib_tag=se&keywords=toys&pd_rd_r=cfbeebeb-5643-4e4f-acf8-452d0bc8a5bf&pd_rd_w=xb5Ae&pd_rd_wg=GxofG&pf_rd_p=44da4965-9668-4613-bec2-a3a75f0c2ad4&pf_rd_r=WREZYKEXD9249MDRSWDR&qid=1717228485&refinements=p_36%3A-2500&sr=8-5",
//	"https://www.amazon.com/-/en/dp/B0BJZ8WJ86/ref=sr_1_1?dib=eyJ2IjoiMSJ9.h4wcHhGqJDYWDNbwqcxOaz0HfEV_he2cxauBO-GVLAupv6jgmKgPWbg2enOIDIbnnzBsAJrhqMu6chMxmP9cVv30njqzaRfCJwWO3Xk4jqdvKCVCSuAB4IqtpCZ-Gj5poN3P9maYDkf0GoOemTEln_d7aeoPYZV4dFLvSj94fIL-EleIHnDQUmGOrAdL0YyEZXPqxiYr-UIz2qQi8eFPOSrRFgbaprQP0tySywNmUa75W-G9E87mOeTxZA2uBcm3ykg8e48UKLc7OPtlsZGGu6n-92DrJxkBwzRtrUq3cjQ.q7Y70ZhT_g7iM-DL4xs4KSxtwigzIdKsqEmJmSUPhBI&dib_tag=se&qid=1716542042&s=fashion-mens-intl-ship&sr=1-1",
//	"https://www.amazon.com/SAMSUNG-Smartphone-Unlocked-Android-Processor/dp/B0CMDPRN7M/ref=sr_1_7?crid=18UMAKQ6QSB0V&dib=eyJ2IjoiMSJ9.THp_WLIyFX2-n1VQorYio1deNbWqTeg2EwxjBOqjecRZ08NJmUJDoW-QxnEmw2g8J-xDVG-AW_FVkvH3S2RW8i70jOZGEdUoYX8Jf5BZ-oDlbxVE4mg3ceH4h-u1Dfx6HNvqasykovjQGHHCHxPtTvoWZv4sd23mbtqBO2pMHZ0djbWDlyzHnXFliBzG5eBg3A3H6877Pz-YnSNO4NSIbu0-Gzctv0KcWer0heJ-jm8.gLDA4ma_UGO5TrVK2e_OuDNlrH40VwVrA9nTj5sY9QQ&dib_tag=se&keywords=phone&qid=1716515618&sprefix=phon%2Caps%2C721&sr=8-7&th=1",
//	"https://www.amazon.com/-/en/dp/B0CHN7H2SL/ref=sr_1_1?dib=eyJ2IjoiMSJ9.mr4o2X0z_1tgbQNezvkZPK5kcImaQnJIikB2oylCs2bsxGnFxs8kDOkIHIjnw0M3_CRMN3kJNKZnrz5_5vlvGVx7Fo1oMJrjN7EOHzla2Gg-93ucVwaOs-qAQZ6gHlNMgSApX3ThIh1y3ERk4vqD3yYQKM-0Y8YB-SoJu8i5RrdMeaBCbcuG7UrO1zKwNvRFbZO2JlPVyU20zpsrF8XAix8SstWJ06sRHrePz-94yYk.KZfp_Ca7FoWZ71fZG3valXYBt9AeU95PmhvznZtE9SU&dib_tag=se&qid=1716543489&s=software-intl-ship&sr=1-1",
//	"https://www.amazon.com/dp/B09MKNL9M3/?mr_donotredirect=undefined&th=1&psc=1",
//	"https://www.amazon.com/-/en/dp/B07WN1FKS3/ref=sr_1_4?dib=eyJ2IjoiMSJ9.iMl2wKNffcNoehjuH_QlWAik428wpZp697-BzRqRti713zOuTmkQQEKvMk2cpGOQTdIGEz9skxbMYHyVW3ob8tzIc93bywcspo3Sx73K18nS-f39Tkkt7-GyRu2oKC4arh_Tc4hyIshVZGnmRbYPOFS0KmQ9Re7wpW27TUm9EV5PFtDh57-2LQMJG9caDt2gil_6MeV_JBLWrjNtVRjKCLeTEoqzJFr1nk-LZwfeRABQUTIbstrCckO4zyaHIoXEWLsmPfgT-EJ4GilkmwTgpenHSmk0dTYyVhUbuZXSbnE.mtkVzBWqHXWIFjy22rl6tZ_gWeWGQ77OB96sYlFoW5c&dib_tag=se&qid=1716553753&s=apparel&sr=1-4&th=1",
//	"https://www.amazon.com/-/en/dp/B08152D6DL?ref_=Oct_DLandingS_D_948c90c7_0&th=1",
//	"https://www.amazon.com/-/en/dp/B072YVWBXH/ref=sr_1_6?dib=eyJ2IjoiMSJ9.ugc4h7Cjckige0OnPnfm0Q9c4L70fU-d3NXpOEc2bpHAsBlRuerqp0-crUeqlOcogxgdQAFeWH18BxkUxUVMqW6n8i4E-XnEraYiaGGUCBJOuUW69SQXKpN7bCjsyR0RK4bUTgQlZ6SWuf41CT15-rcezjYpzA8nSPX94oEfl97DQ_ycLDjvDUduWj-cFra2.Ia_a0Me33aSZ9ir5fj8gPnNNE-LSh-PqYcDEPZmYfPg&dib_tag=se&qid=1716557092&s=hpc-intl-ship&sr=1-6",
//	"https://www.amazon.com/-/en/dp/B09P1DV7D8/ref=sr_1_7?_encoding=UTF8&content-id=amzn1.sym.36b3bd24-30cc-4394-a99c-32175deb1058&dib=eyJ2IjoiMSJ9.dhz8NIOr3fJI_Y_BZn2dio9iMwsxkxh92SLmNsJQ6BQbkcnabGX1H0kNpIfLWKVMx9z7vcLaUipPjPCx1kJipr-TKIdLX8qtObfswj2DMsoT4cxLB10Pjd1U1Aj-2UDJdRGn9CEYedsHlPWE98uvaaVf-0irxC0nLtQDHvlUmKYYV8FU9bkcgt8kH57BAR5n4YxaW1uNF6eCtdNf_8K1FCVNJY94kXMtObD5h0UmXNpdPJuY8oDqUdOl5oeGkKwoFI4xiZEaoEhyiknrDpF9_Sp4Zco3ecRuRClbpfs5m5A.Xt1scjg5sprI3h4pMUhmBm-Jv3bueGe2LDnTHB2kaEk&dib_tag=se&keywords=tech%2Bgifts%2Bfor%2Bmen&pd_rd_r=bea00079-f6b2-4e29-8ca3-6b322d8d4b85&pd_rd_w=z4Vk7&pd_rd_wg=LCOET&pf_rd_p=36b3bd24-30cc-4394-a99c-32175deb1058&pf_rd_r=F03EF7M4YHKFSFD6W28B&qid=1716628148&sr=8-7&th=1",
//}

func main() {
	addrsMap := make(map[string]string)
	filePath := `C:\Users\15564\Downloads\商品库_拆分 2.xlsx`

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		panic(err)
	}

	// 获取第一个 sheet
	sheetName := f.GetSheetName(0)

	rows, err := f.GetRows(sheetName)
	if err != nil {
		panic(err)
	}

	// 找到“关键词”列索引
	var keywordIndex int = -1

	header := rows[0]
	for i, col := range header {
		if col == "关键词" {
			keywordIndex = i
			break
		}
	}

	if keywordIndex == -1 {
		panic("没找到【关键词】列")
	}

	// 遍历数据行
	for _, row := range rows[1:] {
		if keywordIndex >= len(row) {
			continue
		}

		keyword := row[keywordIndex]
		keyword = strings.TrimSpace(keyword)
		if keyword == "" {
			continue
		}
		fmt.Println(keyword)
		//urlPatch := buildAmazonURL(keyword)
		//addrsMap[keyword] = urlPatch
		//fmt.Println(urlPatch)
		//fmt.Printf("第%d行: %s\n%s\n\n", i+2, keyword, urlPatch)
		//if i == 4 {
		//	break
		//}
	}
	//fmt.Println(addrsMap)
	return
	err = crawling2.InitRedis("8.138.57.34:6379", "", 1)
	crawl, err := crawling2.NewCrawling()
	if err != nil {
		return
	}

	crawl.SetRequestHeaders(amazon.Headers()).
		SetOnHtml(amazon.ObtainDetails(0, "#dp")).
		SetOnHtml(amazon.SearchClassLink("#search")).
		SetOnScraped(func(r *colly.Response) {
			if data := r.Ctx.GetAny("message"); data != nil {
				if p, ok := data.(crawling2.ProductAttr); ok {
					fmt.Printf("product：%+v", p)
				}
			}
		})

	errPath := make([]string, 0)
	//urlTmp := strings.Split(addrs[0], "\n")
	// https://www.amazon.com/s?k=Amiri+Vintage+Line+Cargo+Pant

	for k, u := range addrsMap {
		err = crawl.Run(u)
		if err != nil {
			zap.L().Error("Crawling", zap.Error(err))
			errPath = append(errPath, k)
		}
	}
	crawl.Wait()
	crawl.Clone()
	time.Sleep(5 * time.Second)

	fmt.Println(errPath)
}

var DB *gorm.DB

// 拼接 Amazon URL
func buildAmazonURL(keyword string) string {
	base := "https://www.amazon.com/s"

	// 👉 推荐：把逗号换成空格（更符合搜索）
	keyword = strings.ReplaceAll(keyword, ",", " ")

	params := url.Values{}
	params.Set("k", keyword)

	return base + "?" + params.Encode()
}

type CrawProduct struct {
	ID           int64   `gorm:"primaryKey"`
	SourceID     int64   `gorm:"column:source_id"`
	SourceSite   string  `gorm:"column:source_site"`
	Name         string  `gorm:"column:name"`
	Description  string  `gorm:"column:description"`
	Price        float64 `gorm:"column:price"`
	MainImageURL string  `gorm:"column:main_image_url"`
}

func (CrawProduct) TableName() string {
	return "craw_products"
}

//
//func SaveProduct(p crawling.ProductAttr) error {
//	product := CrawProduct{
//		SourceID:     p.SourceID,
//		SourceSite:   p.SourceSite,
//		Name:         p.Name,
//		Description:  p.Description,
//		Price:        p.Price,
//		MainImageURL: p.MainImage,
//	}
//
//	return DB.Clauses(clause.OnConflict{
//		Columns: []clause.Column{
//			{Name: "source_id"},
//			{Name: "source_site"},
//		},
//		DoUpdates: clause.AssignmentColumns([]string{
//			"name",
//			"description",
//			"price",
//			"main_image_url",
//		}),
//	}).Create(&product).Error
//}
