package audio_service_test

import (
	"testing"

	"github.com/Sensory-Cloud/go-sdk/pkg/generated/v1/audio"
	audio_service "github.com/Sensory-Cloud/go-sdk/pkg/service/audio"
	test_util "github.com/Sensory-Cloud/go-sdk/pkg/util/test"
)

func TestFullTranscriptAggregator_ProcessResponse(t *testing.T) {
	aggregator := audio_service.FullTranscriptAggregator{}

	test_util.AssertEquals(t, aggregator.GetCurrentTranscript(" "), "")
	test_util.AssertEquals(t, len(aggregator.GetCurrentWordList()), 0)

	test := audio.TranscribeResponse{}
	test.GetWordList()

	aggregator.ProcessResponse(&audio.TranscribeWordResponse{
		Words: []*audio.TranscribeWord{
			{
				WordIndex: 0,
				Word:      "Foo",
			},
			{
				WordIndex: 1,
				Word:      "Bar",
			},
		},
		LastWordIndex: 1,
	})

	test_util.AssertEquals(t, aggregator.GetCurrentTranscript(""), "Foo Bar")
	test_util.AssertEquals(t, len(aggregator.GetCurrentWordList()), 2)

	aggregator.ProcessResponse(&audio.TranscribeWordResponse{
		Words: []*audio.TranscribeWord{
			{
				WordIndex: 1,
				Word:      "Baz",
			},
		},
		LastWordIndex: 1,
	})

	test_util.AssertEquals(t, aggregator.GetCurrentTranscript(""), "Foo Baz")
	test_util.AssertEquals(t, len(aggregator.GetCurrentWordList()), 2)

	aggregator.ProcessResponse(&audio.TranscribeWordResponse{
		Words: []*audio.TranscribeWord{
			{
				WordIndex: 0,
				Word:      "Foo",
			},
		},
		LastWordIndex: 0,
	})

	test_util.AssertEquals(t, aggregator.GetCurrentTranscript(" "), "Foo")
	test_util.AssertEquals(t, len(aggregator.GetCurrentWordList()), 1)

	aggregator.ProcessResponse(nil)

	test_util.AssertEquals(t, aggregator.GetCurrentTranscript(" "), "Foo")
	test_util.AssertEquals(t, len(aggregator.GetCurrentWordList()), 1)
}
