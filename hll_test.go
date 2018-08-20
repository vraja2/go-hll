package hll

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestNewHLL(t *testing.T) {
  numRegisterBits := 6
  numRegisters := 64
  hllInstance := NewHLL(numRegisterBits)
  assert.Equal(t, numRegisterBits, hllInstance.numRegisterBits, "Num register bits should be equal")
  expectedRegisters := make([]int, numRegisters)
  assert.Equal(t, expectedRegisters, hllInstance.registers, "Registers should be equal")
}

func TestAdd(t *testing.T) {
  numRegisterBits := 6
  hllInstance := NewHLL(numRegisterBits)
  hllInstance.Add("hello")
}
