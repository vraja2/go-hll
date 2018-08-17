package hll

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestNewHLL(t *testing.T) {
  numRegisters := 5
  hllInstance := NewHLL(numRegisters)
  expectedRegisters := make([]int, numRegisters)
  assert.Equal(t, expectedRegisters, hllInstance.registers, "Registers should be equal")
}

func TestAdd(t *testing.T) {
}
