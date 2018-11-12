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
	Flag             bool              `json:"-"` // Just to tell the whether we have initialized this struct or not (zero-value for bool is false, i.e. if the value is ready we know that we have initialized the struct).
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


type RawProducts struct {
	RawProductsList [][8]string `json:"raw-product-groups"`
}

type Products struct {
	ProductsList [][4]string `json:"products"`
	Ret          string    `json:"ret"`
}

type Product struct {
	Product [8]string `json:"product"`
	Ret          string    `json:"ret"`
}



type DomainDb struct {
	productGroups  ProductGroups
	rawProductsMap map[int]RawProducts
	productsMap    map[int]Products
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
	productGroups := ProductGroups{true, myPG}
	util.LogExit()
	return productGroups
}

// NOTE: We skip testing of all numeric/alpha conversions.
// In real production code we should handle all these error conditions, of course.
// But since this is an exercise, let's skip that part at least for now.
func readProducts(pgId int) (RawProducts, Products) {
	util.LogEnter()
	lines := readCsvFile("pg-" + strconv.Itoa(pgId) + "-products.csv")
	count := len(lines)
	util.LogTrace("count: " + strconv.Itoa(count))
	var rawProductsList [][8]string
	var productsList [][4]string
	//productsList := make([]Product, count)
	i := 0
	for _, line := range lines {
		// NOTE: Beware of shadowing pgId => that's why we have myPgId, not pgId (which is function parameter and the variable would shadow it, not a problem here but might be in certain cases).
		myPId := line[0]
		myPgId := line[1]
		myTitle := line[2]
		myPrice := line[3]
		myAuthorOrDirector := line[4]
		myYear := line[5]
		myCountry := line[6]
		myGenreOrLanguage := line[7]
		rawProductsList = append(rawProductsList, [8]string{myPId, myPgId, myTitle, myPrice, myAuthorOrDirector, myYear, myCountry, myGenreOrLanguage})
		productsList = append(productsList, [4]string{myPId, myPgId, myTitle, myPrice})
		i++
	}
	rawProducts := RawProducts{rawProductsList}
	products := Products{productsList, "ok"}
	util.LogExit()
	return rawProducts, products
}

func initDomainDb() DomainDb {
	util.LogEnter()
	myProductGroups := readProductGroups()
	pgMap := myProductGroups.ProductGroupsMap
	rawProductsMap := make(map[int]RawProducts)
	productsMap := make(map[int]Products)
	for key := range pgMap {
		pgId, _ := strconv.Atoi(key)
		rawProducts, products := readProducts(pgId)
		rawProductsMap[pgId] = rawProducts
		productsMap[pgId] = products
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

// Gets product
func GetProduct(pgId int, pId int) Product {
	util.LogEnter()
	rawProductsMap := myDomainDB.rawProductsMap
	rawProducts := rawProductsMap[pgId]
	rawProductsList := rawProducts.RawProductsList
	var found [8]string
	wantedPid := strconv.Itoa(pId)
	for _, product := range rawProductsList {
		if product[0] == wantedPid {
			found = product
			break
		}
	}
	ret := Product{found, "ok"}
	util.LogExit()
	return ret
}
