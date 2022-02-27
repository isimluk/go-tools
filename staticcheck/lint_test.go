package staticcheck

import (
	"testing"

	"honnef.co/go/tools/analysis/lint/testutil"
)

func TestAll(t *testing.T) {
	checks := map[string][]testutil.Test{
		"SA1000": {{Dir: "CheckRegexps"}},
		"SA1001": {{Dir: "CheckTemplate"}},
		"SA1002": {{Dir: "CheckTimeParse"}},
		"SA1003": {{Dir: "CheckEncodingBinary"}, {Dir: "CheckEncodingBinary_go17", Version: "1.7"}, {Dir: "CheckEncodingBinary_go18", Version: "1.8"}},
		"SA1004": {{Dir: "CheckTimeSleepConstant"}},
		"SA1005": {{Dir: "CheckExec"}},
		"SA1006": {{Dir: "CheckUnsafePrintf"}},
		"SA1007": {{Dir: "CheckURLs"}},
		"SA1008": {{Dir: "CheckCanonicalHeaderKey"}},
		"SA1010": {{Dir: "checkStdlibUsageRegexpFindAll"}},
		"SA1011": {{Dir: "checkStdlibUsageUTF8Cutset"}},
		"SA1012": {{Dir: "checkStdlibUsageNilContext"}},
		"SA1013": {{Dir: "checkStdlibUsageSeeker"}},
		"SA1014": {{Dir: "CheckUnmarshalPointer"}},
		"SA1015": {{Dir: "CheckLeakyTimeTick"}, {Dir: "CheckLeakyTimeTick-main"}},
		"SA1016": {{Dir: "CheckUntrappableSignal"}},
		"SA1017": {{Dir: "CheckUnbufferedSignalChan"}},
		"SA1018": {{Dir: "CheckStringsReplaceZero"}},
		"SA1019": {
			{Dir: "CheckDeprecated"},
			{Dir: "CheckDeprecated_go13", Version: "1.3"},
			{Dir: "CheckDeprecated_go14", Version: "1.4"},
			{Dir: "CheckDeprecated_go18", Version: "1.8"},
		},
		"SA1020": {{Dir: "CheckListenAddress"}},
		"SA1021": {{Dir: "CheckBytesEqualIP"}},
		"SA1023": {{Dir: "CheckWriterBufferModified"}},
		"SA1024": {{Dir: "CheckNonUniqueCutset"}},
		"SA1025": {{Dir: "CheckTimerResetReturnValue"}},
		"SA1026": {{Dir: "CheckUnsupportedMarshal"}},
		"SA1027": {{Dir: "CheckAtomicAlignment"}},
		"SA1028": {{Dir: "CheckSortSlice"}},
		"SA1029": {{Dir: "CheckWithValueKey"}},
		"SA1030": {
			{Dir: "CheckStrconv"},
			{Dir: "CheckStrconv_go115", Version: "1.15"},
		},
		"SA2000": {{Dir: "CheckWaitgroupAdd"}},
		"SA2001": {{Dir: "CheckEmptyCriticalSection"}},
		"SA2002": {{Dir: "CheckConcurrentTesting"}},
		"SA2003": {{Dir: "CheckDeferLock"}},
		"SA3000": {
			{Dir: "CheckTestMainExit-1_go14", Version: "1.4"},
			{Dir: "CheckTestMainExit-2_go14", Version: "1.4"},
			{Dir: "CheckTestMainExit-3_go14", Version: "1.4"},
			{Dir: "CheckTestMainExit-4_go14", Version: "1.4"},
			{Dir: "CheckTestMainExit-5_go14", Version: "1.4"},
			{Dir: "CheckTestMainExit-1_go115", Version: "1.15"},
		},
		"SA3001": {{Dir: "CheckBenchmarkN"}},
		"SA4000": {{Dir: "CheckLhsRhsIdentical"}},
		"SA4001": {{Dir: "CheckIneffectiveCopy"}},
		"SA4003": {{Dir: "CheckExtremeComparison"}},
		"SA4004": {{Dir: "CheckIneffectiveLoop"}, {Dir: "CheckIneffectiveLoop_generics", Version: "1.18"}},
		"SA4005": {{Dir: "CheckIneffectiveFieldAssignments"}},
		"SA4006": {{Dir: "CheckUnreadVariableValues"}},
		"SA4008": {{Dir: "CheckLoopCondition"}},
		"SA4009": {{Dir: "CheckArgOverwritten"}},
		"SA4010": {{Dir: "CheckIneffectiveAppend"}},
		"SA4011": {{Dir: "CheckScopedBreak"}},
		"SA4012": {{Dir: "CheckNaNComparison"}},
		"SA4013": {{Dir: "CheckDoubleNegation"}},
		"SA4014": {{Dir: "CheckRepeatedIfElse"}},
		"SA4015": {{Dir: "CheckMathInt"}},
		"SA4016": {{Dir: "CheckSillyBitwiseOps"}, {Dir: "CheckSillyBitwiseOps_shadowedIota"}, {Dir: "CheckSillyBitwiseOps_dotImport"}},
		"SA4017": {{Dir: "CheckPureFunctions"}},
		"SA4018": {{Dir: "CheckSelfAssignment"}},
		"SA4019": {{Dir: "CheckDuplicateBuildConstraints"}},
		"SA4020": {{Dir: "CheckUnreachableTypeCases"}},
		"SA4021": {{Dir: "CheckSingleArgAppend"}},
		"SA4022": {{Dir: "CheckAddressIsNil"}},
		"SA4023": {
			{Dir: "CheckTypedNilInterface"},
			{Dir: "CheckTypedNilInterface/i26000"},
			{Dir: "CheckTypedNilInterface/i27815"},
			{Dir: "CheckTypedNilInterface/i28241"},
			{Dir: "CheckTypedNilInterface/i31873"},
			{Dir: "CheckTypedNilInterface/i33965"},
			{Dir: "CheckTypedNilInterface/i33994"},
			{Dir: "CheckTypedNilInterface/i35217"},
		},
		"SA4024": {{Dir: "CheckBuiltinZeroComparison"}},
		"SA4025": {{Dir: "CheckIntegerDivisionEqualsZero"}},
		"SA4026": {{Dir: "CheckNegativeZeroFloat"}},
		"SA4027": {{Dir: "CheckIneffectiveURLQueryModification"}},
		"SA4028": {{Dir: "CheckModuloOne"}},
		"SA4029": {{Dir: "CheckIneffectiveSort"}},
		"SA4030": {{Dir: "CheckIneffectiveRandInt"}},
		"SA4031": {{Dir: "CheckAllocationNilCheck"}},
		"SA5000": {{Dir: "CheckNilMaps"}},
		"SA5001": {{Dir: "CheckEarlyDefer"}},
		"SA5002": {{Dir: "CheckInfiniteEmptyLoop"}},
		"SA5003": {{Dir: "CheckDeferInInfiniteLoop"}},
		"SA5004": {{Dir: "CheckLoopEmptyDefault"}},
		"SA5005": {{Dir: "CheckCyclicFinalizer"}},
		"SA5007": {{Dir: "CheckInfiniteRecursion"}},
		"SA5008": {{Dir: "CheckStructTags"}, {Dir: "CheckStructTags2"}, {Dir: "CheckStructTags3"}},
		"SA5009": {{Dir: "CheckPrintf"}},
		"SA5010": {{Dir: "CheckImpossibleTypeAssertion"}},
		"SA5011": {{Dir: "CheckMaybeNil"}},
		"SA5012": {{Dir: "CheckEvenSliceLength"}},
		"SA6000": {{Dir: "CheckRegexpMatchLoop"}},
		"SA6001": {{Dir: "CheckMapBytesKey"}},
		"SA6002": {{Dir: "CheckSyncPoolValue"}},
		"SA6003": {{Dir: "CheckRangeStringRunes"}},
		"SA6005": {{Dir: "CheckToLowerToUpperComparison"}},
		"SA9001": {{Dir: "CheckDubiousDeferInChannelRangeLoop"}},
		"SA9002": {{Dir: "CheckNonOctalFileMode"}},
		"SA9003": {{Dir: "CheckEmptyBranch"}},
		"SA9004": {{Dir: "CheckMissingEnumTypesInDeclaration"}},
		"SA9005": {{Dir: "CheckNoopMarshal"}},
		"SA9006": {{Dir: "CheckStaticBitShift"}},
		"SA9007": {{Dir: "CheckBadRemoveAll"}},
		"SA9008": {{Dir: "CheckTypeAssertionShadowingElse"}},
	}

	testutil.Run(t, Analyzers, checks)
}
