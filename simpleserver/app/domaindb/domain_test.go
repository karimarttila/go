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
	util.LogExit()
}

func TestGetProducts(t *testing.T) {
	util.LogEnter()
	myProductsPg_1 := GetProducts(1)
	myProductsPg_2 := GetProducts(2)
	myProductsListPg_1 := myProductsPg_1.ProductsList
	myProductsListPg_2 := myProductsPg_2.ProductsList
	if len(myProductsListPg_1) != 35 {
		t.Errorf("There should be exactly 35 products in product group 1, got: %d", len(myProductsListPg_1))
	}
	if len(myProductsListPg_2) != 169 {
		t.Errorf("There should be exactly 169 products in product group 2, got: %d", len(myProductsListPg_2))
	}
	util.LogExit()
}

func TestGetProduct(t *testing.T) {
	util.LogEnter()
	// What a coincidence! The chosen movie is the best western of all times!
	expectedTitle := "Once Upon a Time in the West"
	rawProduct := GetProduct(2, 49)
	if rawProduct.Title != expectedTitle {
		t.Errorf("Didn't find expected product: expected: %s, got: %s", expectedTitle, rawProduct.Title)
	}
	util.LogExit()
}
