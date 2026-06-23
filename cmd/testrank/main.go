package main

import (
	"fmt"
	"os"
	tsqlparser "github.com/haimiyahya/tsqlparser"
)

func main() {
	sql := `CREATE OR ALTER PROCEDURE [dbo].[test]
	AS
	BEGIN
		SET NOCOUNT ON;
		WITH LatestRecords AS (
			SELECT *,
				   ROW_NUMBER() OVER (
					   PARTITION BY ExtRefNo
					   ORDER BY VersionID DESC
				   ) AS Rank
			FROM [dbo].[DSM_QuotaAcctTracking]
			WHERE LogTimestamp BETWEEN @sz_StartDate AND @sz_EndDate
		)
		SELECT * FROM LatestRecords WHERE Rank = 1;
	END`

	stmt, errs := tsqlparser.Parse(sql)
	if len(errs) > 0 {
		fmt.Println("Errors:")
		for _, e := range errs {
			fmt.Println("  ", e)
		}
		os.Exit(1)
	} else {
		fmt.Println("Success!")
		fmt.Printf("Statement type: %T\n", stmt)
	}
}
