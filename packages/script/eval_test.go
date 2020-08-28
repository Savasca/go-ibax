/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package script

import (
	"fmt"
	"testing"
)

type TestComp struct {
	Input  string
	Output string
		{"56 == 56", "true"},
		{"37 != 37", "false"},
		{"!!(1-1)", "false"},
		{"!!$citizenId || $wallet_id", "true"},
		{"!789", "false"},
		{"$citizenId == 56780 + 9", `true`},
		{"qwerty(45)", `unknown identifier qwerty`},
		{"Multi(2, 5) > 36", "false"},
		{"789 63 == 63", "true"},
		{"+421", "stack is empty"},
		{"1256778+223445==1480223", "true"},
		{"(67-34789)*3 == -104166", "true"},
		{"(5+78)*(1563-527) == 85988", "true"},
		{"124 * (143-527", "there is not pair"},
		{"341 * 234/0", "divided by zero"},
		{"0 == ((15+82)*2 + 5)/2 - 99", "true"},
		{"Multi( (34+35)*2, Multi( $citizenId, 56))== 1 || Multi( (34+35)*2, Multi( $citizenId, 56))== 0", `false`},
		{"2+ Multi( (34+35)*2, Multi( $citizenId, 56)) /2 == 56972", `true`},
		{"$citizenId && 0", "false"},
		{"0|| ($citizenId + $wallet_id == 950240)", "true"},
	}
	vars := map[string]interface{}{
		`citizenId`: 56789,
		`wallet_id`: 893451,
	}
	vm := NewVM()
	vm.Extend(&ExtendData{map[string]interface{}{"Multi": Multi}, nil, nil})
	for _, item := range test {
		out, err := vm.EvalIf(item.Input, 0, &vars)
		if err != nil {
			if err.Error() != item.Output {
				t.Error(`error of ifeval ` + item.Input + ` ` + err.Error())
			}
		} else {
			if fmt.Sprint(out) != item.Output {
				t.Error(`error of ifeval ` + item.Input + ` Output:` + fmt.Sprint(out))
			}
		}
	}
}