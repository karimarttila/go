package domaindb

import (
	"github.com/karimarttila/go/simpleserver/app/util"
	"testing"
)

func TestGetProductGroups(t *testing.T) {
	util.LogEnter()
	myProductGroups := GetProductGroups()
	myPGMap := myProductGroups.ProductGroupsMap
	if len(myPGMap) != 2 {
		t.Errorf("There should be exactly two product groups, got: %d", len(myPGMap))
	}
	if myPGMap["1"] != "Books" || myPGMap["2"] != "Movies" {
		t.Errorf("There were wrong values for product groups: %s", myPGMap)
	}
}

//func TestGetProducts(t *testing.T) {
//	util.LogEnter()
//	myProducts := GetProducts(1)
//	myProductsList := myProducts.ProductsList
//	if len(myProductsList) != 2 {
//		t.Errorf("There should be exactly 2 products in product group 1, got: %d", len(myProductsList))
//	}
//}

