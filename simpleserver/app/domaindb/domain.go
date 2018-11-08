package domaindb

import (
	"encoding/csv"
	"github.com/karimarttila/go/simpleserver/app/util"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// DomainDB singleton.
var myDomainDB = initDomainDb()

type ProductGroups struct {
	ProductGroupsMap map[string]string `json:"product-groups"`
}

type RawProduct struct {
	PgId             int
	PId              int
	Title            string
	Price            float64
	AuthorOrDirector string
	Year             int
	Country          string
	GenreOrLanguage  string
}

type Product struct {
	PgId             int
	PId              int
	Title            string
	Price            float64
}

type RawProducts struct {
	RawProductsList []RawProduct `json:"raw-product-groups"`
}

type Products struct {
	ProductsList []Product `json:"product-groups"`
	Ret string `json:"ret"`
}

type DomainDb struct {
	productGroups ProductGroups
	rawProductsMap map[int]RawProducts
	productsMap map[int]Products
}

func readCsvFile(csvFileName string) [][]string {
	util.LogEnter()
	var lines [][]string
	dir, _ := os.Getwd()
	util.LogDebug("dir: " + dir)
	fileName := []string{"../../resources/" + csvFileName}
	_, dirName, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirName), strings.Join(fileName, ""))
	csvFile, err := os.Open(filePath)
	if err != nil {
		util.LogError("Failed to open csv file: " + filePath)
	} else {
		reader := csv.NewReader(csvFile)
		reader.Comma = '\t'
		lines, err = reader.ReadAll()
		if err != nil {
			util.LogError("Failed to read csv file: " + filePath)
		}
	}
	util.LogExit()
	return lines
}

func readProductGroups() ProductGroups {
	util.LogEnter()
	lines := readCsvFile("product-groups.csv")
	myPG := make(map[string]string)
	for _, line := range lines {
		myPG[line[0]] = line[1]
	}
	productGroups := ProductGroups{myPG}
	util.LogExit()
	return productGroups
}

func readProducts(pgId int) (RawProducts, Products) {
	util.LogEnter()
	lines := readCsvFile("pg-" + strconv.Itoa(pgId) + "-productsproduct-groups.csv")
	count := len(lines)
	rawProductsList := make([]RawProduct, count)
	productsList := make([]Product, count)
	i := 0
	for _, line := range lines {
		pgId, _ := strconv.Atoi(line[0])
		pId, _  := strconv.Atoi(line[1])
		title := line[2]
		price, _ := strconv.ParseFloat(line[3], 64)
		authorOrDirector := line[4]
		year, _ := strconv.Atoi(line[5])
		country := line[6]
		genreOrLanguage := line[7]
		rawProductsList[i] = RawProduct{pgId, pId, title, price, authorOrDirector, year, country, genreOrLanguage }
		productsList[i] = Product{pgId, pId, title, price}
		i++
	}
	rawProducts := RawProducts{ rawProductsList}
	products := Products{ productsList, "ok"}
	util.LogExit()
	return rawProducts, products
}


func initDomainDb() DomainDb {
	util.LogEnter()
	myProductGroups := readProductGroups()
	pgMap := myProductGroups.ProductGroupsMap
	pgKeys := make([]string, len(pgMap))
	rawProductsMap := make(map[int]RawProducts)
	productsMap := make(map[int]Products)
	for i := range pgKeys {
		pgId, _ := strconv.Atoi(pgKeys[i])
		rawProducts, products := readProducts(pgId)
		rawProductsMap[i] = rawProducts
		productsMap[i] = products
	}
	ret := DomainDb{
		productGroups: myProductGroups, rawProductsMap: rawProductsMap, productsMap: productsMap}
	util.LogExit()
	return ret
}

// Gets product groups.
func GetProductGroups() ProductGroups {
	util.LogEnter()
	ret := myDomainDB.productGroups
	util.LogExit()
	return ret
}

// Gets products
func GetProducts(pgId int) Products {
	util.LogEnter()
	ret := myDomainDB.productsMap[pgId]
	util.LogExit()
	return ret
}
