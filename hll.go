package hll

type HLL struct {
  registers []int
}

func NewHLL(numRegisters int) HLL {
  registers := make([]int, numRegisters)
  hllInstance := HLL {registers}
  return hllInstance
}

func add() {

}

func count() {

}

func merge() {

}
