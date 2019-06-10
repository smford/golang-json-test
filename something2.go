package main

import (
	"encoding/json"
	"fmt"
	"log"
	//	 "net"
	"net/http"
	// "net/url"
	// "os"
	"time"
)

type MyDownloadedData struct {
	GetRaspBlipspotResults GetRaspBlipspotResults `json:"get_rasp_blipspot_results"`
}

//=====
type Values struct {
	Zero600Lst json.Number `json:"0600lst,omitempty"`
	FcstPd     json.Number `json:"Fcst Pd"`
	Zero630Lst json.Number `json:"0630lst,omitempty"`
	Zero700Lst json.Number `json:"0700lst,omitempty"`
	Zero730Lst json.Number `json:"0730lst,omitempty"`
	Zero800Lst json.Number `json:"0800lst,omitempty"`
	Zero830Lst json.Number `json:"0830lst,omitempty"`
	Zero900Lst json.Number `json:"0900lst,omitempty"`
	Zero930Lst json.Number `json:"0930lst,omitempty"`
	One000Lst  json.Number `json:"1000lst,omitempty"`
	One030Lst  json.Number `json:"1030lst,omitempty"`
	One100Lst  json.Number `json:"1100lst,omitempty"`
	One130Lst  json.Number `json:"1130lst,omitempty"`
	One200Lst  json.Number `json:"1200lst,omitempty"`
	One230Lst  json.Number `json:"1230lst,omitempty"`
	One300Lst  json.Number `json:"1300lst,omitempty"`
	One330Lst  json.Number `json:"1330lst,omitempty"`
	One400Lst  json.Number `json:"1400lst,omitempty"`
	One430Lst  json.Number `json:"1430lst,omitempty"`
	One500Lst  json.Number `json:"1500lst,omitempty"`
	One530Lst  json.Number `json:"1530lst,omitempty"`
	One600Lst  json.Number `json:"1600lst,omitempty"`
	One630Lst  json.Number `json:"1630lst,omitempty"`
	One700Lst  json.Number `json:"1700lst,omitempty"`
	One730Lst  json.Number `json:"1730lst,omitempty"`
	One800Lst  json.Number `json:"1800lst,omitempty"`
	One830Lst  json.Number `json:"1830lst,omitempty"`
	One900Lst  json.Number `json:"1900lst,omitempty"`
}

//===
type W struct {
	Values []Values `json:"values"`
}
type BLTop struct {
	Values []Values `json:"values"`
}
type ThmlHt struct {
	Values []Values `json:"values"`
}
type Hcrit175 struct {
	Values []Values `json:"values"`
}
type SfcSun struct {
	Values []Values `json:"values"`
}
type Temp2M struct {
	Values []Values `json:"values"`
}
type DewPt2M struct {
	Values []Values `json:"values"`
}
type MSLPress struct {
	Values []Values `json:"values"`
}
type SfcWDir struct {
	Values []Values `json:"values"`
}
type SfcWSpd struct {
	Values []Values `json:"values"`
}
type BLWindSpd struct {
	Values []Values `json:"values"`
}
type BlWindDir struct {
	Values []Values `json:"values"`
}
type MaxConverg struct {
	Values []Values `json:"values"`
}
type CUpot struct {
	Values []Values `json:"values"`
}
type OneHrRain struct {
	Values []Values `json:"values"`
}
type Stars struct {
	Values []Values `json:"values"`
}
type Results struct {
	W          W          `json:"W*"`
	BLTop      BLTop      `json:"BL Top"`
	ThmlHt     ThmlHt     `json:"Thml Ht"`
	Hcrit175   Hcrit175   `json:"Hcrit(175)"`
	SfcSun     SfcSun     `json:"Sfc. Sun %"`
	Temp2M     Temp2M     `json:"Temp@2m"`
	DewPt2M    DewPt2M    `json:"DewPt@2m"`
	MSLPress   MSLPress   `json:"MSL Press"`
	SfcWDir    SfcWDir    `json:"Sfc. W.Dir"`
	SfcWSpd    SfcWSpd    `json:"Sfc. W.Spd"`
	BLWindSpd  BLWindSpd  `json:"BL Wind Spd"`
	BlWindDir  BlWindDir  `json:"Bl Wind Dir"`
	MaxConverg MaxConverg `json:"Max.Converg"`
	CUpot      CUpot      `json:"CUpot"`
	OneHrRain  OneHrRain  `json:"1hr Rain"`
	Stars      Stars      `json:"Stars"`
}
type GetRaspBlipspotResults struct {
	Mapinfo string  `json:"mapinfo"`
	Region  string  `json:"region"`
	Grid    string  `json:"grid"`
	GridI   int     `json:"grid-i"`
	GridJ   int     `json:"grid-j"`
	Lat     float64 `json:"Lat"`
	Lon     float64 `json:"Lon"`
	Created string  `json:"Created"`
	Results Results `json:"Results"`
}

func main() {
	region := "UK12"
	grid := "d2"
	url := fmt.Sprintf("http://rasp.mrsap.org/cgi-bin/get_rasp_blipspot.cgi?region=%s&grid=%s&day=0&i=1332&k=1547&width=2000&height=2000&linfo=1&param=CUpot&format=json", region, grid)

	//==========
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}
	client.Timeout = time.Second * 15

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer resp.Body.Close()
	//==========
	var record MyDownloadedData

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	//fmt.Printf("%+v", record)

	fmt.Println("TEST!:" + record.GetRaspBlipspotResults.Results.Hcrit175.values[1].One000Lst)

}
