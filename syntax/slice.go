/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package syntax

import "fmt"

type Slice struct {
}

func (a *Slice)SliceRise(s []int) {
	s = append(s, 0)
	for i, _:= range s {
		s[i]++
	}
}


func (a *Slice)Boot() {
	var s1 = []int{1,2}
	s2 := s1
	s2 = append(s2, 3)
	a.SliceRise(s1)
	a.SliceRise(s2)
	fmt.Println(s1)
	fmt.Println(s2)
}
