package srg

import "fmt"

func (s *Srg) QuotaEquivalence(){
	fmt.Println(" ", s.PivotQuota())
	fmt.Println("-", s.PivotProduct())
	fmt.Println(" ", s.PivotQuotaLeft())
}