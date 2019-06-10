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
	GetRaspBlipspotResults struct {
		Mapinfo string  `json:"mapinfo"`
		Region  string  `json:"region"`
		Grid    string  `json:"grid"`
		GridI   int     `json:"grid-i"`
		GridJ   int     `json:"grid-j"`
		Lat     float64 `json:"Lat"`
		Lon     float64 `json:"Lon"`
		Created string  `json:"Created"`
		Results struct {
			W struct {
				Values []struct {
					Zero600Lst string `json:"0600lst,omitempty"`
					FcstPd     string `json:"Fcst Pd"`
					Zero630Lst string `json:"0630lst,omitempty"`
					Zero700Lst int    `json:"0700lst,omitempty"`
					Zero730Lst int    `json:"0730lst,omitempty"`
					Zero800Lst int    `json:"0800lst,omitempty"`
					Zero830Lst int    `json:"0830lst,omitempty"`
					Zero900Lst int    `json:"0900lst,omitempty"`
					Zero930Lst int    `json:"0930lst,omitempty"`
					One000Lst  int    `json:"1000lst,omitempty"`
					One030Lst  int    `json:"1030lst,omitempty"`
					One100Lst  int    `json:"1100lst,omitempty"`
					One130Lst  int    `json:"1130lst,omitempty"`
					One200Lst  int    `json:"1200lst,omitempty"`
					One230Lst  int    `json:"1230lst,omitempty"`
					One300Lst  int    `json:"1300lst,omitempty"`
					One330Lst  int    `json:"1330lst,omitempty"`
					One400Lst  int    `json:"1400lst,omitempty"`
					One430Lst  int    `json:"1430lst,omitempty"`
					One500Lst  int    `json:"1500lst,omitempty"`
					One530Lst  int    `json:"1530lst,omitempty"`
					One600Lst  int    `json:"1600lst,omitempty"`
					One630Lst  int    `json:"1630lst,omitempty"`
					One700Lst  int    `json:"1700lst,omitempty"`
					One730Lst  int    `json:"1730lst,omitempty"`
					One800Lst  int    `json:"1800lst,omitempty"`
					One830Lst  int    `json:"1830lst,omitempty"`
					One900Lst  int    `json:"1900lst,omitempty"`
				} `json:"values"`
			} `json:"W*"`
			BLTop struct {
				Values []struct {
					Zero600Lst string `json:"0600lst,omitempty"`
					FcstPd     string `json:"Fcst Pd"`
					Zero630Lst string `json:"0630lst,omitempty"`
					Zero700Lst int    `json:"0700lst,omitempty"`
					Zero730Lst int    `json:"0730lst,omitempty"`
					Zero800Lst int    `json:"0800lst,omitempty"`
					Zero830Lst int    `json:"0830lst,omitempty"`
					Zero900Lst int    `json:"0900lst,omitempty"`
					Zero930Lst int    `json:"0930lst,omitempty"`
					One000Lst  int    `json:"1000lst,omitempty"`
					One030Lst  int    `json:"1030lst,omitempty"`
					One100Lst  int    `json:"1100lst,omitempty"`
					One130Lst  int    `json:"1130lst,omitempty"`
					One200Lst  int    `json:"1200lst,omitempty"`
					One230Lst  int    `json:"1230lst,omitempty"`
					One300Lst  int    `json:"1300lst,omitempty"`
					One330Lst  int    `json:"1330lst,omitempty"`
					One400Lst  int    `json:"1400lst,omitempty"`
					One430Lst  int    `json:"1430lst,omitempty"`
					One500Lst  int    `json:"1500lst,omitempty"`
					One530Lst  int    `json:"1530lst,omitempty"`
					One600Lst  int    `json:"1600lst,omitempty"`
					One630Lst  int    `json:"1630lst,omitempty"`
					One700Lst  int    `json:"1700lst,omitempty"`
					One730Lst  int    `json:"1730lst,omitempty"`
					One800Lst  int    `json:"1800lst,omitempty"`
					One830Lst  int    `json:"1830lst,omitempty"`
					One900Lst  int    `json:"1900lst,omitempty"`
				} `json:"values"`
			} `json:"BL Top"`
			ThmlHt struct {
				Values []struct {
					Zero600Lst string `json:"0600lst,omitempty"`
					FcstPd     string `json:"Fcst Pd"`
					Zero630Lst string `json:"0630lst,omitempty"`
					Zero700Lst int    `json:"0700lst,omitempty"`
					Zero730Lst int    `json:"0730lst,omitempty"`
					Zero800Lst int    `json:"0800lst,omitempty"`
					Zero830Lst int    `json:"0830lst,omitempty"`
					Zero900Lst int    `json:"0900lst,omitempty"`
					Zero930Lst int    `json:"0930lst,omitempty"`
					One000Lst  int    `json:"1000lst,omitempty"`
					One030Lst  int    `json:"1030lst,omitempty"`
					One100Lst  int    `json:"1100lst,omitempty"`
					One130Lst  int    `json:"1130lst,omitempty"`
					One200Lst  int    `json:"1200lst,omitempty"`
					One230Lst  int    `json:"1230lst,omitempty"`
					One300Lst  int    `json:"1300lst,omitempty"`
					One330Lst  int    `json:"1330lst,omitempty"`
					One400Lst  int    `json:"1400lst,omitempty"`
					One430Lst  int    `json:"1430lst,omitempty"`
					One500Lst  int    `json:"1500lst,omitempty"`
					One530Lst  int    `json:"1530lst,omitempty"`
					One600Lst  int    `json:"1600lst,omitempty"`
					One630Lst  int    `json:"1630lst,omitempty"`
					One700Lst  int    `json:"1700lst,omitempty"`
					One730Lst  int    `json:"1730lst,omitempty"`
					One800Lst  int    `json:"1800lst,omitempty"`
					One830Lst  int    `json:"1830lst,omitempty"`
					One900Lst  int    `json:"1900lst,omitempty"`
				} `json:"values"`
			} `json:"Thml Ht"`
			Hcrit175 struct {
				Values []struct {
					Zero600Lst string `json:"0600lst,omitempty"`
					FcstPd     string `json:"Fcst Pd"`
					Zero630Lst string `json:"0630lst,omitempty"`
					Zero700Lst int    `json:"0700lst,omitempty"`
					Zero730Lst int    `json:"0730lst,omitempty"`
					Zero800Lst int    `json:"0800lst,omitempty"`
					Zero830Lst int    `json:"0830lst,omitempty"`
					Zero900Lst int    `json:"0900lst,omitempty"`
					Zero930Lst int    `json:"0930lst,omitempty"`
					One000Lst  int    `json:"1000lst,omitempty"`
					One030Lst  int    `json:"1030lst,omitempty"`
					One100Lst  int    `json:"1100lst,omitempty"`
					One130Lst  int    `json:"1130lst,omitempty"`
					One200Lst  int    `json:"1200lst,omitempty"`
					One230Lst  int    `json:"1230lst,omitempty"`
					One300Lst  int    `json:"1300lst,omitempty"`
					One330Lst  int    `json:"1330lst,omitempty"`
					One400Lst  int    `json:"1400lst,omitempty"`
					One430Lst  int    `json:"1430lst,omitempty"`
					One500Lst  int    `json:"1500lst,omitempty"`
					One530Lst  int    `json:"1530lst,omitempty"`
					One600Lst  int    `json:"1600lst,omitempty"`
					One630Lst  int    `json:"1630lst,omitempty"`
					One700Lst  int    `json:"1700lst,omitempty"`
					One730Lst  int    `json:"1730lst,omitempty"`
					One800Lst  int    `json:"1800lst,omitempty"`
					One830Lst  int    `json:"1830lst,omitempty"`
					One900Lst  int    `json:"1900lst,omitempty"`
				} `json:"values"`
			} `json:"Hcrit(175)"`
			SfcSun struct {
				Values []struct {
					Zero600Lst string  `json:"0600lst,omitempty"`
					FcstPd     string  `json:"Fcst Pd"`
					Zero630Lst string  `json:"0630lst,omitempty"`
					Zero700Lst float64 `json:"0700lst,omitempty"`
					Zero730Lst float64 `json:"0730lst,omitempty"`
					Zero800Lst float64 `json:"0800lst,omitempty"`
					Zero830Lst float64 `json:"0830lst,omitempty"`
					Zero900Lst float64 `json:"0900lst,omitempty"`
					Zero930Lst float64 `json:"0930lst,omitempty"`
					One000Lst  float64 `json:"1000lst,omitempty"`
					One030Lst  float64 `json:"1030lst,omitempty"`
					One100Lst  float64 `json:"1100lst,omitempty"`
					One130Lst  float64 `json:"1130lst,omitempty"`
					One200Lst  float64 `json:"1200lst,omitempty"`
					One230Lst  float64 `json:"1230lst,omitempty"`
					One300Lst  float64 `json:"1300lst,omitempty"`
					One330Lst  float64 `json:"1330lst,omitempty"`
					One400Lst  float64 `json:"1400lst,omitempty"`
					One430Lst  float64 `json:"1430lst,omitempty"`
					One500Lst  float64 `json:"1500lst,omitempty"`
					One530Lst  float64 `json:"1530lst,omitempty"`
					One600Lst  float64 `json:"1600lst,omitempty"`
					One630Lst  float64 `json:"1630lst,omitempty"`
					One700Lst  float64 `json:"1700lst,omitempty"`
					One730Lst  float64 `json:"1730lst,omitempty"`
					One800Lst  float64 `json:"1800lst,omitempty"`
					One830Lst  float64 `json:"1830lst,omitempty"`
					One900Lst  float64 `json:"1900lst,omitempty"`
				} `json:"values"`
			} `json:"Sfc. Sun %"`
			Temp2M struct {
				Values []struct {
					Zero600Lst string  `json:"0600lst,omitempty"`
					FcstPd     string  `json:"Fcst Pd"`
					Zero630Lst string  `json:"0630lst,omitempty"`
					Zero700Lst float64 `json:"0700lst,omitempty"`
					Zero730Lst float64 `json:"0730lst,omitempty"`
					Zero800Lst float64 `json:"0800lst,omitempty"`
					Zero830Lst float64 `json:"0830lst,omitempty"`
					Zero900Lst float64 `json:"0900lst,omitempty"`
					Zero930Lst float64 `json:"0930lst,omitempty"`
					One000Lst  float64 `json:"1000lst,omitempty"`
					One030Lst  float64 `json:"1030lst,omitempty"`
					One100Lst  float64 `json:"1100lst,omitempty"`
					One130Lst  float64 `json:"1130lst,omitempty"`
					One200Lst  float64 `json:"1200lst,omitempty"`
					One230Lst  float64 `json:"1230lst,omitempty"`
					One300Lst  float64 `json:"1300lst,omitempty"`
					One330Lst  float64 `json:"1330lst,omitempty"`
					One400Lst  float64 `json:"1400lst,omitempty"`
					One430Lst  float64 `json:"1430lst,omitempty"`
					One500Lst  float64 `json:"1500lst,omitempty"`
					One530Lst  float64 `json:"1530lst,omitempty"`
					One600Lst  float64 `json:"1600lst,omitempty"`
					One630Lst  float64 `json:"1630lst,omitempty"`
					One700Lst  float64 `json:"1700lst,omitempty"`
					One730Lst  float64 `json:"1730lst,omitempty"`
					One800Lst  float64 `json:"1800lst,omitempty"`
					One830Lst  float64 `json:"1830lst,omitempty"`
					One900Lst  float64 `json:"1900lst,omitempty"`
				} `json:"values"`
			} `json:"Temp@2m"`
			DewPt2M struct {
				Values []struct {
					Zero600Lst string  `json:"0600lst,omitempty"`
					FcstPd     string  `json:"Fcst Pd"`
					Zero630Lst string  `json:"0630lst,omitempty"`
					Zero700Lst float64 `json:"0700lst,omitempty"`
					Zero730Lst float64 `json:"0730lst,omitempty"`
					Zero800Lst float64 `json:"0800lst,omitempty"`
					Zero830Lst float64 `json:"0830lst,omitempty"`
					Zero900Lst float64 `json:"0900lst,omitempty"`
					Zero930Lst float64 `json:"0930lst,omitempty"`
					One000Lst  float64 `json:"1000lst,omitempty"`
					One030Lst  float64 `json:"1030lst,omitempty"`
					One100Lst  float64 `json:"1100lst,omitempty"`
					One130Lst  float64 `json:"1130lst,omitempty"`
					One200Lst  float64 `json:"1200lst,omitempty"`
					One230Lst  float64 `json:"1230lst,omitempty"`
					One300Lst  float64 `json:"1300lst,omitempty"`
					One330Lst  float64 `json:"1330lst,omitempty"`
					One400Lst  float64 `json:"1400lst,omitempty"`
					One430Lst  float64 `json:"1430lst,omitempty"`
					One500Lst  float64 `json:"1500lst,omitempty"`
					One530Lst  float64 `json:"1530lst,omitempty"`
					One600Lst  float64 `json:"1600lst,omitempty"`
					One630Lst  float64 `json:"1630lst,omitempty"`
					One700Lst  float64 `json:"1700lst,omitempty"`
					One730Lst  float64 `json:"1730lst,omitempty"`
					One800Lst  float64 `json:"1800lst,omitempty"`
					One830Lst  float64 `json:"1830lst,omitempty"`
					One900Lst  float64 `json:"1900lst,omitempty"`
				} `json:"values"`
			} `json:"DewPt@2m"`
			MSLPress struct {
				Values []struct {
					Zero600Lst string  `json:"0600lst,omitempty"`
					FcstPd     string  `json:"Fcst Pd"`
					Zero630Lst string  `json:"0630lst,omitempty"`
					Zero700Lst float64 `json:"0700lst,omitempty"`
					Zero730Lst float64 `json:"0730lst,omitempty"`
					Zero800Lst float64 `json:"0800lst,omitempty"`
					Zero830Lst float64 `json:"0830lst,omitempty"`
					Zero900Lst float64 `json:"0900lst,omitempty"`
					Zero930Lst float64 `json:"0930lst,omitempty"`
					One000Lst  float64 `json:"1000lst,omitempty"`
					One030Lst  float64 `json:"1030lst,omitempty"`
					One100Lst  float64 `json:"1100lst,omitempty"`
					One130Lst  float64 `json:"1130lst,omitempty"`
					One200Lst  float64 `json:"1200lst,omitempty"`
					One230Lst  float64 `json:"1230lst,omitempty"`
					One300Lst  float64 `json:"1300lst,omitempty"`
					One330Lst  float64 `json:"1330lst,omitempty"`
					One400Lst  float64 `json:"1400lst,omitempty"`
					One430Lst  float64 `json:"1430lst,omitempty"`
					One500Lst  float64 `json:"1500lst,omitempty"`
					One530Lst  float64 `json:"1530lst,omitempty"`
					One600Lst  float64 `json:"1600lst,omitempty"`
					One630Lst  float64 `json:"1630lst,omitempty"`
					One700Lst  float64 `json:"1700lst,omitempty"`
					One730Lst  float64 `json:"1730lst,omitempty"`
					One800Lst  float64 `json:"1800lst,omitempty"`
					One830Lst  float64 `json:"1830lst,omitempty"`
					One900Lst  float64 `json:"1900lst,omitempty"`
				} `json:"values"`
			} `json:"MSL Press"`
			SfcWDir struct {
				Values []struct {
					Zero600Lst string `json:"0600lst,omitempty"`
					FcstPd     string `json:"Fcst Pd"`
					Zero630Lst string `json:"0630lst,omitempty"`
					Zero700Lst int    `json:"0700lst,omitempty"`
					Zero730Lst int    `json:"0730lst,omitempty"`
					Zero800Lst int    `json:"0800lst,omitempty"`
					Zero830Lst int    `json:"0830lst,omitempty"`
					Zero900Lst int    `json:"0900lst,omitempty"`
					Zero930Lst int    `json:"0930lst,omitempty"`
					One000Lst  int    `json:"1000lst,omitempty"`
					One030Lst  int    `json:"1030lst,omitempty"`
					One100Lst  int    `json:"1100lst,omitempty"`
					One130Lst  int    `json:"1130lst,omitempty"`
					One200Lst  int    `json:"1200lst,omitempty"`
					One230Lst  int    `json:"1230lst,omitempty"`
					One300Lst  int    `json:"1300lst,omitempty"`
					One330Lst  int    `json:"1330lst,omitempty"`
					One400Lst  int    `json:"1400lst,omitempty"`
					One430Lst  int    `json:"1430lst,omitempty"`
					One500Lst  int    `json:"1500lst,omitempty"`
					One530Lst  int    `json:"1530lst,omitempty"`
					One600Lst  int    `json:"1600lst,omitempty"`
					One630Lst  int    `json:"1630lst,omitempty"`
					One700Lst  int    `json:"1700lst,omitempty"`
					One730Lst  int    `json:"1730lst,omitempty"`
					One800Lst  int    `json:"1800lst,omitempty"`
					One830Lst  int    `json:"1830lst,omitempty"`
					One900Lst  int    `json:"1900lst,omitempty"`
				} `json:"values"`
			} `json:"Sfc. W.Dir"`
			SfcWSpd struct {
				Values []struct {
					Zero600Lst string `json:"0600lst,omitempty"`
					FcstPd     string `json:"Fcst Pd"`
					Zero630Lst string `json:"0630lst,omitempty"`
					Zero700Lst int    `json:"0700lst,omitempty"`
					Zero730Lst int    `json:"0730lst,omitempty"`
					Zero800Lst int    `json:"0800lst,omitempty"`
					Zero830Lst int    `json:"0830lst,omitempty"`
					Zero900Lst int    `json:"0900lst,omitempty"`
					Zero930Lst int    `json:"0930lst,omitempty"`
					One000Lst  int    `json:"1000lst,omitempty"`
					One030Lst  int    `json:"1030lst,omitempty"`
					One100Lst  int    `json:"1100lst,omitempty"`
					One130Lst  int    `json:"1130lst,omitempty"`
					One200Lst  int    `json:"1200lst,omitempty"`
					One230Lst  int    `json:"1230lst,omitempty"`
					One300Lst  int    `json:"1300lst,omitempty"`
					One330Lst  int    `json:"1330lst,omitempty"`
					One400Lst  int    `json:"1400lst,omitempty"`
					One430Lst  int    `json:"1430lst,omitempty"`
					One500Lst  int    `json:"1500lst,omitempty"`
					One530Lst  int    `json:"1530lst,omitempty"`
					One600Lst  int    `json:"1600lst,omitempty"`
					One630Lst  int    `json:"1630lst,omitempty"`
					One700Lst  int    `json:"1700lst,omitempty"`
					One730Lst  int    `json:"1730lst,omitempty"`
					One800Lst  int    `json:"1800lst,omitempty"`
					One830Lst  int    `json:"1830lst,omitempty"`
					One900Lst  int    `json:"1900lst,omitempty"`
				} `json:"values"`
			} `json:"Sfc. W.Spd"`
			BLWindSpd struct {
				Values []struct {
					Zero600Lst string `json:"0600lst,omitempty"`
					FcstPd     string `json:"Fcst Pd"`
					Zero630Lst string `json:"0630lst,omitempty"`
					Zero700Lst int    `json:"0700lst,omitempty"`
					Zero730Lst int    `json:"0730lst,omitempty"`
					Zero800Lst int    `json:"0800lst,omitempty"`
					Zero830Lst int    `json:"0830lst,omitempty"`
					Zero900Lst int    `json:"0900lst,omitempty"`
					Zero930Lst int    `json:"0930lst,omitempty"`
					One000Lst  int    `json:"1000lst,omitempty"`
					One030Lst  int    `json:"1030lst,omitempty"`
					One100Lst  int    `json:"1100lst,omitempty"`
					One130Lst  int    `json:"1130lst,omitempty"`
					One200Lst  int    `json:"1200lst,omitempty"`
					One230Lst  int    `json:"1230lst,omitempty"`
					One300Lst  int    `json:"1300lst,omitempty"`
					One330Lst  int    `json:"1330lst,omitempty"`
					One400Lst  int    `json:"1400lst,omitempty"`
					One430Lst  int    `json:"1430lst,omitempty"`
					One500Lst  int    `json:"1500lst,omitempty"`
					One530Lst  int    `json:"1530lst,omitempty"`
					One600Lst  int    `json:"1600lst,omitempty"`
					One630Lst  int    `json:"1630lst,omitempty"`
					One700Lst  int    `json:"1700lst,omitempty"`
					One730Lst  int    `json:"1730lst,omitempty"`
					One800Lst  int    `json:"1800lst,omitempty"`
					One830Lst  int    `json:"1830lst,omitempty"`
					One900Lst  int    `json:"1900lst,omitempty"`
				} `json:"values"`
			} `json:"BL Wind Spd"`
			BlWindDir struct {
				Values []struct {
					Zero600Lst string `json:"0600lst,omitempty"`
					FcstPd     string `json:"Fcst Pd"`
					Zero630Lst string `json:"0630lst,omitempty"`
					Zero700Lst int    `json:"0700lst,omitempty"`
					Zero730Lst int    `json:"0730lst,omitempty"`
					Zero800Lst int    `json:"0800lst,omitempty"`
					Zero830Lst int    `json:"0830lst,omitempty"`
					Zero900Lst int    `json:"0900lst,omitempty"`
					Zero930Lst int    `json:"0930lst,omitempty"`
					One000Lst  int    `json:"1000lst,omitempty"`
					One030Lst  int    `json:"1030lst,omitempty"`
					One100Lst  int    `json:"1100lst,omitempty"`
					One130Lst  int    `json:"1130lst,omitempty"`
					One200Lst  int    `json:"1200lst,omitempty"`
					One230Lst  int    `json:"1230lst,omitempty"`
					One300Lst  int    `json:"1300lst,omitempty"`
					One330Lst  int    `json:"1330lst,omitempty"`
					One400Lst  int    `json:"1400lst,omitempty"`
					One430Lst  int    `json:"1430lst,omitempty"`
					One500Lst  int    `json:"1500lst,omitempty"`
					One530Lst  int    `json:"1530lst,omitempty"`
					One600Lst  int    `json:"1600lst,omitempty"`
					One630Lst  int    `json:"1630lst,omitempty"`
					One700Lst  int    `json:"1700lst,omitempty"`
					One730Lst  int    `json:"1730lst,omitempty"`
					One800Lst  int    `json:"1800lst,omitempty"`
					One830Lst  int    `json:"1830lst,omitempty"`
					One900Lst  int    `json:"1900lst,omitempty"`
				} `json:"values"`
			} `json:"Bl Wind Dir"`
			MaxConverg struct {
				Values []struct {
					Zero600Lst string `json:"0600lst,omitempty"`
					FcstPd     string `json:"Fcst Pd"`
					Zero630Lst string `json:"0630lst,omitempty"`
					Zero700Lst int    `json:"0700lst,omitempty"`
					Zero730Lst int    `json:"0730lst,omitempty"`
					Zero800Lst int    `json:"0800lst,omitempty"`
					Zero830Lst int    `json:"0830lst,omitempty"`
					Zero900Lst int    `json:"0900lst,omitempty"`
					Zero930Lst int    `json:"0930lst,omitempty"`
					One000Lst  int    `json:"1000lst,omitempty"`
					One030Lst  int    `json:"1030lst,omitempty"`
					One100Lst  int    `json:"1100lst,omitempty"`
					One130Lst  int    `json:"1130lst,omitempty"`
					One200Lst  int    `json:"1200lst,omitempty"`
					One230Lst  int    `json:"1230lst,omitempty"`
					One300Lst  int    `json:"1300lst,omitempty"`
					One330Lst  int    `json:"1330lst,omitempty"`
					One400Lst  int    `json:"1400lst,omitempty"`
					One430Lst  int    `json:"1430lst,omitempty"`
					One500Lst  int    `json:"1500lst,omitempty"`
					One530Lst  int    `json:"1530lst,omitempty"`
					One600Lst  int    `json:"1600lst,omitempty"`
					One630Lst  int    `json:"1630lst,omitempty"`
					One700Lst  int    `json:"1700lst,omitempty"`
					One730Lst  int    `json:"1730lst,omitempty"`
					One800Lst  int    `json:"1800lst,omitempty"`
					One830Lst  int    `json:"1830lst,omitempty"`
					One900Lst  int    `json:"1900lst,omitempty"`
				} `json:"values"`
			} `json:"Max.Converg"`
			CUpot struct {
				Values []struct {
					Zero600Lst string `json:"0600lst,omitempty"`
					FcstPd     string `json:"Fcst Pd"`
					Zero630Lst string `json:"0630lst,omitempty"`
					Zero700Lst int    `json:"0700lst,omitempty"`
					Zero730Lst int    `json:"0730lst,omitempty"`
					Zero800Lst int    `json:"0800lst,omitempty"`
					Zero830Lst int    `json:"0830lst,omitempty"`
					Zero900Lst int    `json:"0900lst,omitempty"`
					Zero930Lst int    `json:"0930lst,omitempty"`
					One000Lst  int    `json:"1000lst,omitempty"`
					One030Lst  int    `json:"1030lst,omitempty"`
					One100Lst  int    `json:"1100lst,omitempty"`
					One130Lst  int    `json:"1130lst,omitempty"`
					One200Lst  int    `json:"1200lst,omitempty"`
					One230Lst  int    `json:"1230lst,omitempty"`
					One300Lst  int    `json:"1300lst,omitempty"`
					One330Lst  int    `json:"1330lst,omitempty"`
					One400Lst  int    `json:"1400lst,omitempty"`
					One430Lst  int    `json:"1430lst,omitempty"`
					One500Lst  int    `json:"1500lst,omitempty"`
					One530Lst  int    `json:"1530lst,omitempty"`
					One600Lst  int    `json:"1600lst,omitempty"`
					One630Lst  int    `json:"1630lst,omitempty"`
					One700Lst  int    `json:"1700lst,omitempty"`
					One730Lst  int    `json:"1730lst,omitempty"`
					One800Lst  int    `json:"1800lst,omitempty"`
					One830Lst  int    `json:"1830lst,omitempty"`
					One900Lst  int    `json:"1900lst,omitempty"`
				} `json:"values"`
			} `json:"CUpot"`
			OneHrRain struct {
				Values []struct {
					Zero600Lst string  `json:"0600lst,omitempty"`
					FcstPd     string  `json:"Fcst Pd"`
					Zero630Lst string  `json:"0630lst,omitempty"`
					Zero700Lst float64 `json:"0700lst,omitempty"`
					Zero730Lst float64 `json:"0730lst,omitempty"`
					Zero800Lst float64 `json:"0800lst,omitempty"`
					Zero830Lst float64 `json:"0830lst,omitempty"`
					Zero900Lst float64 `json:"0900lst,omitempty"`
					Zero930Lst float64 `json:"0930lst,omitempty"`
					One000Lst  float64 `json:"1000lst,omitempty"`
					One030Lst  float64 `json:"1030lst,omitempty"`
					One100Lst  float64 `json:"1100lst,omitempty"`
					One130Lst  float64 `json:"1130lst,omitempty"`
					One200Lst  float64 `json:"1200lst,omitempty"`
					One230Lst  float64 `json:"1230lst,omitempty"`
					One300Lst  float64 `json:"1300lst,omitempty"`
					One330Lst  float64 `json:"1330lst,omitempty"`
					One400Lst  float64 `json:"1400lst,omitempty"`
					One430Lst  float64 `json:"1430lst,omitempty"`
					One500Lst  float64 `json:"1500lst,omitempty"`
					One530Lst  float64 `json:"1530lst,omitempty"`
					One600Lst  float64 `json:"1600lst,omitempty"`
					One630Lst  float64 `json:"1630lst,omitempty"`
					One700Lst  float64 `json:"1700lst,omitempty"`
					One730Lst  float64 `json:"1730lst,omitempty"`
					One800Lst  float64 `json:"1800lst,omitempty"`
					One830Lst  float64 `json:"1830lst,omitempty"`
					One900Lst  float64 `json:"1900lst,omitempty"`
				} `json:"values"`
			} `json:"1hr Rain"`
			Stars struct {
				Values []struct {
					Zero600Lst string  `json:"0600lst,omitempty"`
					FcstPd     string  `json:"Fcst Pd"`
					Zero630Lst string  `json:"0630lst,omitempty"`
					Zero700Lst float64 `json:"0700lst,omitempty"`
					Zero730Lst float64 `json:"0730lst,omitempty"`
					Zero800Lst float64 `json:"0800lst,omitempty"`
					Zero830Lst float64 `json:"0830lst,omitempty"`
					Zero900Lst float64 `json:"0900lst,omitempty"`
					Zero930Lst float64 `json:"0930lst,omitempty"`
					One000Lst  float64 `json:"1000lst,omitempty"`
					One030Lst  float64 `json:"1030lst,omitempty"`
					One100Lst  float64 `json:"1100lst,omitempty"`
					One130Lst  float64 `json:"1130lst,omitempty"`
					One200Lst  float64 `json:"1200lst,omitempty"`
					One230Lst  float64 `json:"1230lst,omitempty"`
					One300Lst  float64 `json:"1300lst,omitempty"`
					One330Lst  float64 `json:"1330lst,omitempty"`
					One400Lst  float64 `json:"1400lst,omitempty"`
					One430Lst  float64 `json:"1430lst,omitempty"`
					One500Lst  float64 `json:"1500lst,omitempty"`
					One530Lst  float64 `json:"1530lst,omitempty"`
					One600Lst  float64 `json:"1600lst,omitempty"`
					One630Lst  float64 `json:"1630lst,omitempty"`
					One700Lst  float64 `json:"1700lst,omitempty"`
					One730Lst  float64 `json:"1730lst,omitempty"`
					One800Lst  float64 `json:"1800lst,omitempty"`
					One830Lst  float64 `json:"1830lst,omitempty"`
					One900Lst  float64 `json:"1900lst,omitempty"`
				} `json:"values"`
			} `json:"Stars"`
		} `json:"Results"`
	} `json:"get_rasp_blipspot_results"`
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

	fmt.Printf("%+v", record)

}
