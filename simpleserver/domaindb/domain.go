package domaindb

import (
	"encoding/csv"
	"github.com/karimarttila/go/simpleserver/util"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

type ProductGroups struct {
    productGroups   map[string]string   `json:"product-groups"`
}

var myDomainDB = initDomainDb()

type DomainDb struct {
	productGroups  ProductGroups
}

func readProductGroups() (ProductGroups) {
	util.LogEnter()
	var productGroups ProductGroups
	dir, _ := os.Getwd()
	util.LogDebug("dir: " + dir)
	fileName := []string{"../resources/product-groups.csv"}
	_, dirName, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirName), strings.Join(fileName, ""))
	csvFile, err := os.Open(filePath)
	if err != nil {
		util.LogError("Failed to open csv file: " + filePath)
	} else {
		reader := csv.NewReader(csvFile)
		reader.Comma = '\t'
		lines, err := reader.ReadAll()
		if err != nil {
			util.LogError("Failed to read csv file: " + filePath)
		} else {
			myPG := make(map[string]string)
			for _, row := range lines {
				myPG[row[0]] = row[1]
			}
			productGroups = ProductGroups{myPG}
		}
	}
	util.LogExit()
	return productGroups
}

func initDomainDb() (DomainDb){
	util.LogEnter()
	myProductGroups := readProductGroups()
	ret := DomainDb {
		productGroups: myProductGroups,
	}
	util.LogExit()
	return ret
}

func GetProductGroups() (ProductGroups) {
	util.LogEnter()
	ret := myDomainDB.productGroups
	util.LogExit()
	return ret
}