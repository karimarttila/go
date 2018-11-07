package domaindb


import (
	"github.com/karimarttila/go/simpleserver/util"
	"testing"
)

func TestGetProductGroups(t *testing.T) {
	util.LogEnter()
	myProductGroups := GetProductGroups()
	myPGMap := myProductGroups.productGroups
	if len(myPGMap) != 2 {
		t.Errorf("There should be exactly two product groups, got: %d", len(myPGMap))
	}
}