package audio_service

import (
	audio_api_v1 "github.com/Sensory-Cloud/go-sdk/pkg/generated/v1/audio"
)

// FullTranscriptAggregator aggregates and stores transcription responses.
// This class can be used to maintain the full transcript returned from the server.
type FullTranscriptAggregator struct {
	//  Hold onto the entire transcript as an array of words
	currentWordList []*audio_api_v1.TranscribeWord
}

// ProcessResponse processes a single, sliding-window response from the server
func (a *FullTranscriptAggregator) ProcessResponse(response *audio_api_v1.TranscribeWordResponse) {
	// If nothing is returned, do nothing
	if response == nil || len(response.Words) == 0 {
		return
	}

	// Get the expected transcript size from the index of the last word.
	responseSize := response.GetLastWordIndex() + 1

	// Grow the word buffer if the incoming transcript is larger.
	if responseSize > uint64(len(a.currentWordList)) {

		// Compute size difference
		cacheSizeDifference := responseSize - uint64(len(a.currentWordList))

		// Expand the cached word list by the number of incoming items by adding empty records
		cacheExtension := make([]*audio_api_v1.TranscribeWord, cacheSizeDifference)
		a.currentWordList = append(a.currentWordList, cacheExtension...)
	}

	// Loop through returned words and set the returned value at the specified index in currentWordList
	for idx, word := range response.GetWords() {
		a.currentWordList[word.GetWordIndex()] = response.GetWords()[idx]
	}

	// // Check if the transcript is smaller than our currentWordList
	if responseSize < uint64(len(a.currentWordList)) {
		// Remove trailing elements from the array
		a.currentWordList = a.currentWordList[:responseSize]
	}
}

// GetCurrentWordList returns the cache of words returned from the server
func (a *FullTranscriptAggregator) GetCurrentWordList() []*audio_api_v1.TranscribeWord {
	return a.currentWordList
}

// GetCurrentTranscript returns the full transcript as computed from the current word list
func (a *FullTranscriptAggregator) GetCurrentTranscript(delimiter string) string {
	if len(a.currentWordList) == 0 {
		return ""
	}

	if delimiter == "" {
		delimiter = " "
	}

	transcript := ""
	for _, word := range a.currentWordList {
		transcript += delimiter + word.GetWord()
	}

	return transcript[len(delimiter):]
}
