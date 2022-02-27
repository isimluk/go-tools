package simple

import (
	"testing"

	"honnef.co/go/tools/analysis/lint/testutil"
)

func TestAll(t *testing.T) {
	checks := map[string][]testutil.Test{
		"S1000": {{Dir: "single-case-select"}},
		"S1001": {{Dir: "copy"}, {Dir: "copy_generics", Version: "1.18"}},
		"S1002": {{Dir: "bool-cmp"}, {Dir: "bool-cmp_generics", Version: "1.18"}},
		"S1003": {{Dir: "contains"}},
		"S1004": {{Dir: "compare"}},
		"S1005": {{Dir: "CheckBlankOK"}, {Dir: "receive-blank"}, {Dir: "range_go13", Version: "1.3"}, {Dir: "range_go14", Version: "1.4"}},
		"S1006": {{Dir: "for-true"}, {Dir: "generated"}},
		"S1007": {{Dir: "regexp-raw"}},
		"S1008": {{Dir: "if-return"}},
		"S1009": {{Dir: "nil-len"}, {Dir: "nil-len_generics", Version: "1.18"}},
		"S1010": {{Dir: "slicing"}},
		"S1011": {{Dir: "loop-append"}},
		"S1012": {{Dir: "time-since"}},
		"S1016": {
			{Dir: "convert"},
			{Dir: "convert_go17", Version: "1.7"},
			{Dir: "convert_go18", Version: "1.8"},
			{Dir: "convert_generics", Version: "1.18"},
		},
		"S1017": {{Dir: "trim"}},
		"S1018": {{Dir: "CheckLoopSlide"}},
		"S1019": {{Dir: "CheckMakeLenCap"}, {Dir: "CheckMakeLenCap_generics", Version: "1.18"}},
		"S1020": {{Dir: "CheckAssertNotNil"}},
		"S1021": {{Dir: "CheckDeclareAssign"}},
		"S1023": {{Dir: "CheckRedundantBreak"}, {Dir: "CheckRedundantReturn"}},
		"S1024": {{Dir: "CheckTimeUntil_go17", Version: "1.7"}, {Dir: "CheckTimeUntil_go18", Version: "1.8"}},
		"S1025": {{Dir: "CheckRedundantSprintf"}},
		"S1028": {{Dir: "CheckErrorsNewSprintf"}},
		"S1029": {{Dir: "CheckRangeStringRunes"}},
		"S1030": {{Dir: "CheckBytesBufferConversions"}},
		"S1031": {{Dir: "CheckNilCheckAroundRange"}, {Dir: "CheckNilCheckAroundRange_generics", Version: "1.18"}},
		"S1032": {{Dir: "CheckSortHelpers"}},
		"S1033": {{Dir: "CheckGuardedDelete"}},
		"S1034": {{Dir: "CheckSimplifyTypeSwitch"}},
		"S1035": {{Dir: "CheckRedundantCanonicalHeaderKey"}},
		"S1036": {{Dir: "CheckUnnecessaryGuard"}},
		"S1037": {{Dir: "CheckElaborateSleep"}},
		"S1038": {{Dir: "CheckPrintSprintf"}},
		"S1039": {{Dir: "CheckSprintLiteral"}},
		"S1040": {{Dir: "CheckSameTypeTypeAssertion"}},
	}

	testutil.Run(t, Analyzers, checks)
}
