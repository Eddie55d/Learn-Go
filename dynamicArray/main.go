package main

import (
	"fmt"
)

const (
	superSmall int = 2 << iota
	small
	medium
	large
)

type Array struct {
	superSmallArr *[superSmall]int
	superSmallIdx int

	smallArr *[small]int
	smallIdx int

	mediumArr *[medium]int
	mediumIdx int

	largeArr *[large]int
	largeIdx int

	currentLen int
	currentCap int

	// что-, где ты будешь хранить текущий размер и текущее capacity
}

func (a *Array) Add(v int) int {

	switch {
	case a.currentLen < superSmall || a.currentLen == 0:
		fmt.Println("superSmallArr")
		if a.superSmallIdx == 0 {
			a.superSmallArr = &[superSmall]int{}
			a.superSmallArr[0] = v
			a.superSmallIdx++
			a.currentLen++
			a.currentCap = len(a.superSmallArr)
			return 0
		}
		for idx, val := range a.superSmallArr {
			if val == 0 {
				a.superSmallArr[idx] = v
				a.superSmallIdx++
				a.currentLen++

				return idx
			}
		}
	case a.currentLen < small:
		fmt.Println("smallArr")
		if a.smallIdx == 0 {
			a.smallArr = &[small]int{}
			for idx, val := range a.superSmallArr {
				a.smallArr[idx] = val
				a.smallIdx++
			}
			a.currentCap = len(a.smallArr)
			a.superSmallArr = nil
			a.superSmallIdx = 0
		}
		for idx, val := range a.smallArr {
			if val == 0 {
				a.smallArr[idx] = v
				a.smallIdx++
				a.currentLen++

				return idx
			}
		}

	case a.currentLen < medium:
		fmt.Println("mediumArr")
		if a.mediumIdx == 0 {
			a.mediumArr = &[medium]int{}
			for idx, val := range a.smallArr {
				a.mediumArr[idx] = val
				a.mediumIdx++
			}
			a.currentCap = len(a.mediumArr)
			a.smallArr = nil
			a.smallIdx = 0
		}
		for idx, val := range a.mediumArr {
			if val == 0 {
				a.mediumArr[idx] = v
				a.mediumIdx++
				a.currentLen++

				return idx
			}
		}

	case a.currentLen < large:
		fmt.Println("largeArr")
		if a.largeIdx == 0 {
			a.largeArr = &[large]int{}
			for idx, val := range a.mediumArr {
				a.largeArr[idx] = val
				a.largeIdx++
			}
			a.currentCap = len(a.largeArr)
			a.mediumArr = nil
			a.mediumIdx = 0
		}
		for idx, val := range a.largeArr {
			if val == 0 {
				a.largeArr[idx] = v
				a.largeIdx++
				a.currentLen++

				return idx
			}
		}
	}
	return 0
}

func (a *Array) Get(idx int) int {
	switch {
	case a.superSmallArr != nil:
		for i, v := range a.superSmallArr {
			if i == idx {
				return v
			}
		}
	case a.smallArr != nil:
		for i, v := range a.smallArr {
			if i == idx {
				return v
			}
		}
	case a.mediumArr != nil:
		for i, v := range a.mediumArr {
			if i == idx {
				return v
			}
		}
	case a.largeArr != nil:
		for i, v := range a.largeArr {
			if i == idx {
				return v
			}
		}
	}
	return 0
}

func (a *Array) Delete(idx int) int {
	switch {
	case a.superSmallArr != nil:
		if idx > len(a.superSmallArr)-1 {
			return 0
		}
		if a.superSmallArr[idx] == 0 {
			return 0
		}
		if idx == len(a.superSmallArr)-1 {
			a.superSmallArr[idx] = 0
			a.superSmallIdx--
			a.currentLen--
			return idx
		}
		for i := 0; i <= len(a.superSmallArr)-1; i++ {
			if i >= idx && i != len(a.superSmallArr)-1 {
				a.superSmallArr[i] = a.superSmallArr[i+1]
			}
		}
		if a.superSmallArr[superSmall-1] != 0 {
			a.superSmallArr[superSmall-1] = 0
		}

		a.superSmallIdx--
		a.currentLen--

		return idx

	case a.smallArr != nil:
		if idx > len(a.smallArr)-1 {
			return 0
		}
		if a.smallArr[idx] == 0 {
			return 0
		}
		if idx == len(a.smallArr)-1 {
			a.smallArr[idx] = 0
			a.smallIdx--
			a.currentLen--
			return idx
		}
		for i := 0; i <= len(a.smallArr)-1; i++ {
			if i >= idx && i != len(a.smallArr)-1 {
				a.smallArr[i] = a.smallArr[i+1]
			}

		}
		if a.smallArr[small-1] != 0 {
			a.smallArr[small-1] = 0
		}
		a.smallIdx--
		a.currentLen--
		if a.smallIdx == superSmall {
			a.superSmallArr = &[superSmall]int{}

			for i, v := range a.smallArr {
				if v != 0 {
					a.superSmallArr[i] = v
					a.superSmallIdx++
				}
			}
			a.currentCap = superSmall
			a.smallIdx = 0
			a.smallArr = nil
		}
		return idx

	case a.mediumArr != nil:
		if idx > len(a.mediumArr)-1 {
			return 0
		}
		if a.mediumArr[idx] == 0 {
			return 0
		}
		if idx == medium-1 {
			a.mediumArr[idx] = 0
			a.mediumIdx--
			a.currentLen--
			return idx
		}
		for i := 0; i <= len(a.mediumArr)-1; i++ {
			if i >= idx && i != medium-1 {
				a.mediumArr[i] = a.mediumArr[i+1]
			}
		}
		if a.mediumArr[medium-1] != 0 {
			a.mediumArr[medium-1] = 0
		}
		a.mediumIdx--
		a.currentLen--
		if a.mediumIdx == small {
			a.smallArr = &[small]int{}
			for i, v := range a.mediumArr {
				if v != 0 {
					a.smallArr[i] = v
					a.smallIdx++
				}
			}
			a.currentCap = small
			a.mediumIdx = 0
			a.mediumArr = nil
		}
		return idx
	case a.largeArr != nil:
		if idx > len(a.largeArr)-1 {
			return 0
		}
		if a.largeArr[idx] == 0 {
			return 0
		}
		if idx == len(a.largeArr)-1 {
			a.largeArr[idx] = 0
			a.largeIdx--
			a.currentLen--
			return idx
		}
		for i := 0; i <= len(a.largeArr)-1; i++ {
			if i >= idx && i != len(a.largeArr)-1 {
				a.largeArr[i] = a.largeArr[i+1]
			}

		}
		if a.largeArr[large-1] != 0 {
			a.largeArr[large-1] = 0
		}
		a.largeIdx--
		a.currentLen--
		if a.largeIdx == medium {
			a.mediumArr = &[medium]int{}

			for i, v := range a.largeArr {
				if v != 0 {
					a.mediumArr[i] = v
					a.mediumIdx++
				}
			}

			a.currentCap = medium
			a.largeIdx = 0
			a.largeArr = nil
		}
		return idx
	}

	return 0

}

func main() {

	arr := new(Array)

	fmt.Println(arr.currentCap)
	fmt.Println(arr.currentLen)
	fmt.Println(arr.Add(1))
	fmt.Println(arr.Add(2))
	fmt.Println(arr.currentCap)
	fmt.Println(arr.currentLen)
	fmt.Println(arr.Add(3))
	fmt.Println(arr.Add(4))
	fmt.Println(arr.currentCap)
	fmt.Println(arr.currentLen)
	fmt.Println(arr.Delete(3))
	fmt.Println(arr.Delete(2))
	fmt.Println(arr.Get(1))
	fmt.Println(arr.currentCap)
	fmt.Println(arr.currentLen)
	fmt.Println(arr.Delete(1))
	fmt.Println(arr.currentCap)
	fmt.Println(arr.currentLen)
	fmt.Println(arr.Delete(1))
	fmt.Println(arr.currentCap)
	fmt.Println(arr.currentLen)
}
