package isvalidsubsequence

func IsValidSubsequence(subSequence, sequence []int) bool {
	secIdx, subSeqIdx := 0, 0
	sequenceLength, subSequenceLength := len(sequence), len(subSequence)

	for secIdx < sequenceLength && subSeqIdx < subSequenceLength {
		if sequence[secIdx] == subSequence[subSeqIdx] {
			subSeqIdx++
		}

		secIdx++
	}

	return subSeqIdx == len(subSequence)
}

func IsValidSubsequence2(subSequence, sequence []int) bool {
	subSeqIdx := 0
	for _, value := range sequence {
		if subSeqIdx == len(subSequence) {
			break
		}

		if value == subSequence[subSeqIdx] {
			subSeqIdx++
		}
	}

	return subSeqIdx == len(subSequence)
}
