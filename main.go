package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

var payloads = []string{}

func main() {
	for _, stringData := range payloads {
		byteString, _ := hex.DecodeString(stringData)
		parsedData, err := DecodeTCP(&byteString)
		if err != nil {
			log.Panicf("Error when decoding a byteString, %v\n", err)
		}
		//
		// // fmt.Printf("Decoded packet codec 8:\n%+v\n", parsedData)
		// for _, avl := range parsedData.Data {
		// 	fmt.Printf("\n%s\n", time.Unix(int64(avl.Utime), 0))
		// 	if avl.EventID != 0 {
		// 		fmt.Println("EventID: ", avl.EventID)
		// 	}
		// 	for _, element := range avl.Elements {
		// 		if element.IOID == 48 {
		// 			fmt.Printf("48 - %d %%\n", element.Value[0])
		// 		} else if element.IOID == 84 {
		// 			fmt.Printf("84 - %d Litres\n", binary.BigEndian.Uint32(element.Value))
		// 		} else if element.IOID == 390 {
		// 			fmt.Printf("390 - %.1f Litres\n", float32(binary.BigEndian.Uint32(element.Value))*0.1)
		// 		} else {
		// 			fmt.Printf("%d - %v\n", element.IOID, element.Value)
		// 		}
		// 	}
		// }

		// initialize a human decoder
		humanDecoder := HumanDecoder{}

		// loop over raw data
		for _, val := range parsedData.Data {
			// loop over Elements
			for _, ioel := range val.Elements {
				// decode to human readable format
				decoded, err := humanDecoder.Human(&ioel, "FMBXY") // second parameter - device family type ["FMBXY", "FM64"]
				if err != nil {
					// log.Panicf("Error when converting human, %v\n", err)
					continue
				}

				// get final decoded value to value which is specified in ./teltonikajson/ in paramether FinalConversion
				if val, err := (*decoded).GetFinalValue(); err != nil {
					// log.Panicf("Unable to GetFinalValue() %v", err)
				} else if val != nil {
					// print output
					fmt.Printf("ID: %v, Property Name: %v, Value: %v\n", decoded.Element.IOID, decoded.AvlEncodeKey.PropertyName, val)
				}
			}
			fmt.Println("")
		}
		// spew.Dump(parsedData)
	}
}
