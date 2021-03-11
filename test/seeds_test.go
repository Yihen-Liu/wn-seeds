/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2021-03-10
 */
package test

import (
	"kangaroo-net/ethereum/contracts"
	"testing"
)

//go test -v -test.run TestAddSeed
func TestAddSeed(t *testing.T) {
	contracts.AddSeed("0x135e39D35b4a8740a215e81928E770A1Bbe2faCd")
}

func TestQuerySeed(t *testing.T) {
	contracts.Seeds(0)
}

func TestUpdateSeed(t *testing.T) {
	contracts.UpdateSeed(0,"0x135e39D35b4a8740a215e81928E770A1Bbe2faCC")
}

func TestCountSeeds(t *testing.T)  {
	contracts.SeedNums()
}

func TestDeleteSeed(t *testing.T)  {
	contracts.DeleteSeed(0)
}
