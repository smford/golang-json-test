package main

import (
	"encoding/json"
	"fmt"
	"log"
	//	 "net"
	//"net/http"
	// "net/url"
	"io/ioutil"
	"os"
	//"time"
)

type MyDownloadedData struct {
	GetRaspBlipspotResults GetRaspBlipspotResults `json:"get_rasp_blipspot_results"`
}

//=====
type Values struct {
	FcstPd   json.Number `json:"Fcst Pd"`
	t0600lst json.Number `json:"0600lst,omitempty"`
	t0630lst json.Number `json:"0630lst,omitempty"`
	t0700lst json.Number `json:"0700lst,omitempty"`
	t0730lst json.Number `json:"0730lst,omitempty"`
	t0800lst json.Number `json:"0800lst,omitempty"`
	t0830lst json.Number `json:"0830lst,omitempty"`
	t0900lst json.Number `json:"0900lst,omitempty"`
	t0930lst json.Number `json:"0930lst,omitempty"`
	t1000lst json.Number `json:"1000lst,omitempty"`
	t1030lst json.Number `json:"1030lst,omitempty"`
	t1100lst json.Number `json:"1100lst,omitempty"`
	t1130lst json.Number `json:"1130lst,omitempty"`
	t1200lst json.Number `json:"1200lst,omitempty"`
	t1230lst json.Number `json:"1230lst,omitempty"`
	t1300lst json.Number `json:"1300lst,omitempty"`
	t1330lst json.Number `json:"1330lst,omitempty"`
	t1400lst json.Number `json:"1400lst,omitempty"`
	t1430lst json.Number `json:"1430lst,omitempty"`
	t1500lst json.Number `json:"1500lst,omitempty"`
	t1530lst json.Number `json:"1530lst,omitempty"`
	t1600lst json.Number `json:"1600lst,omitempty"`
	t1630lst json.Number `json:"1630lst,omitempty"`
	t1700lst json.Number `json:"1700lst,omitempty"`
	t1730lst json.Number `json:"1730lst,omitempty"`
	t1800lst json.Number `json:"1800lst,omitempty"`
	t1830lst json.Number `json:"1830lst,omitempty"`
	t1900lst json.Number `json:"1900lst,omitempty"`
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
	Mapinfo string      `json:"mapinfo"`
	Region  string      `json:"region"`
	Grid    string      `json:"grid"`
	GridI   int         `json:"grid-i"`
	GridJ   int         `json:"grid-j"`
	Lat     json.Number `json:"Lat"`
	Lon     json.Number `json:"Lon"`
	Created string      `json:"Created"`
	Results Results     `json:"Results"`
}

func main() {
	file, err := os.Open("data-wstar.json")
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("error ioutil")
	}

	//fmt.Print(b)

	//==========
	var record MyDownloadedData

	err = json.Unmarshal([]byte(b), &record)

	if err != nil {
		log.Fatal("error unmarshalling")
	}

	fmt.Printf("%+v", record)

	fmt.Println("\nMapinfo:" + record.GetRaspBlipspotResults.Mapinfo)
	//fmt.Println("" + record.GetRaspBlipspotResults.)
	fmt.Println("Lat:" + string(record.GetRaspBlipspotResults.Lat))
	fmt.Println("Long:" + string(record.GetRaspBlipspotResults.Lon))
	fmt.Println("Created:" + record.GetRaspBlipspotResults.Created)
	//fmt.Println("TEST!:" + record.GetRaspBlipspotResults.Results.Hcrit175.values[1].One000Lst)
	//fmt.Println("1-0600lst:", record.GetRaspBlipspotResults.Results.Hcrit175.values[1].0600lst)
	fmt.Println("W-0-0730lst:" + record.GetRaspBlipspotResults.Results.W.Values[0].t0730lst)
	fmt.Println("W-1-0730lst:" + record.GetRaspBlipspotResults.Results.W.Values[1].t0730lst)
	fmt.Println("W-2-0730lst:" + record.GetRaspBlipspotResults.Results.W.Values[2].t0730lst)
	fmt.Println("--------------------------")
	fmt.Printf("%+v", record.GetRaspBlipspotResults.Results.W.Values)
}
