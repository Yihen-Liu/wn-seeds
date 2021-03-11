#!/usr/bin/env bash
# at first, we need to install abigen and solc, solc version must match the seeds.sol;
# in this implement, we use version v0.8.2 of solc
abigen --sol $1/$1.sol --out $1/$1.go --pkg $1