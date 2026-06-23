package main

import (
	"fmt"
	"github.com/haimiyahya/tsqlparser"
)

func main() {
	sql := `SELECT *, ROW_NUMBER() OVER (PARTITION BY ExtRefNo ORDER BY VersionID DESC) AS Rank FROM DSM_QuotaAcctTracking`
	_, errs := tsqlparser.Parse(sql)
	if len(errs) > 0 {
		fmt.Println("Errors:")
		for _, e := range errs {
			fmt.Println("  ", e)
		}
	} else {
		fmt.Println("Success!")
	}
}
