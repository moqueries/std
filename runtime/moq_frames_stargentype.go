// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package runtime

import (
	"fmt"
	"math/bits"
	"runtime"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that runtime.Frames_starGenType is
// mocked completely
var _ Frames_starGenType = (*MoqFrames_starGenType_mock)(nil)

// Frames_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type Frames_starGenType interface {
	Next() (frame runtime.Frame, more bool)
}

// MoqFrames_starGenType holds the state of a moq of the Frames_starGenType
// type
type MoqFrames_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqFrames_starGenType_mock

	ResultsByParams_Next []MoqFrames_starGenType_Next_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Next struct{}
		}
	}
}

// MoqFrames_starGenType_mock isolates the mock interface of the
// Frames_starGenType type
type MoqFrames_starGenType_mock struct {
	Moq *MoqFrames_starGenType
}

// MoqFrames_starGenType_recorder isolates the recorder interface of the
// Frames_starGenType type
type MoqFrames_starGenType_recorder struct {
	Moq *MoqFrames_starGenType
}

// MoqFrames_starGenType_Next_params holds the params of the Frames_starGenType
// type
type MoqFrames_starGenType_Next_params struct{}

// MoqFrames_starGenType_Next_paramsKey holds the map key params of the
// Frames_starGenType type
type MoqFrames_starGenType_Next_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqFrames_starGenType_Next_resultsByParams contains the results for a given
// set of parameters for the Frames_starGenType type
type MoqFrames_starGenType_Next_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFrames_starGenType_Next_paramsKey]*MoqFrames_starGenType_Next_results
}

// MoqFrames_starGenType_Next_doFn defines the type of function needed when
// calling AndDo for the Frames_starGenType type
type MoqFrames_starGenType_Next_doFn func()

// MoqFrames_starGenType_Next_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Frames_starGenType type
type MoqFrames_starGenType_Next_doReturnFn func() (frame runtime.Frame, more bool)

// MoqFrames_starGenType_Next_results holds the results of the
// Frames_starGenType type
type MoqFrames_starGenType_Next_results struct {
	Params  MoqFrames_starGenType_Next_params
	Results []struct {
		Values *struct {
			Frame runtime.Frame
			More  bool
		}
		Sequence   uint32
		DoFn       MoqFrames_starGenType_Next_doFn
		DoReturnFn MoqFrames_starGenType_Next_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFrames_starGenType_Next_fnRecorder routes recorded function calls to the
// MoqFrames_starGenType moq
type MoqFrames_starGenType_Next_fnRecorder struct {
	Params    MoqFrames_starGenType_Next_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFrames_starGenType_Next_results
	Moq       *MoqFrames_starGenType
}

// MoqFrames_starGenType_Next_anyParams isolates the any params functions of
// the Frames_starGenType type
type MoqFrames_starGenType_Next_anyParams struct {
	Recorder *MoqFrames_starGenType_Next_fnRecorder
}

// NewMoqFrames_starGenType creates a new moq of the Frames_starGenType type
func NewMoqFrames_starGenType(scene *moq.Scene, config *moq.Config) *MoqFrames_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqFrames_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqFrames_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Next struct{}
			}
		}{ParameterIndexing: struct {
			Next struct{}
		}{
			Next: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Frames_starGenType type
func (m *MoqFrames_starGenType) Mock() *MoqFrames_starGenType_mock { return m.Moq }

func (m *MoqFrames_starGenType_mock) Next() (frame runtime.Frame, more bool) {
	m.Moq.Scene.T.Helper()
	params := MoqFrames_starGenType_Next_params{}
	var results *MoqFrames_starGenType_Next_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Next {
		paramsKey := m.Moq.ParamsKey_Next(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Next(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Next(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Next(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		frame = result.Values.Frame
		more = result.Values.More
	}
	if result.DoReturnFn != nil {
		frame, more = result.DoReturnFn()
	}
	return
}

// OnCall returns the recorder implementation of the Frames_starGenType type
func (m *MoqFrames_starGenType) OnCall() *MoqFrames_starGenType_recorder {
	return &MoqFrames_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqFrames_starGenType_recorder) Next() *MoqFrames_starGenType_Next_fnRecorder {
	return &MoqFrames_starGenType_Next_fnRecorder{
		Params:   MoqFrames_starGenType_Next_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqFrames_starGenType_Next_fnRecorder) Any() *MoqFrames_starGenType_Next_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Next(r.Params))
		return nil
	}
	return &MoqFrames_starGenType_Next_anyParams{Recorder: r}
}

func (r *MoqFrames_starGenType_Next_fnRecorder) Seq() *MoqFrames_starGenType_Next_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Next(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFrames_starGenType_Next_fnRecorder) NoSeq() *MoqFrames_starGenType_Next_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Next(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFrames_starGenType_Next_fnRecorder) ReturnResults(frame runtime.Frame, more bool) *MoqFrames_starGenType_Next_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Frame runtime.Frame
			More  bool
		}
		Sequence   uint32
		DoFn       MoqFrames_starGenType_Next_doFn
		DoReturnFn MoqFrames_starGenType_Next_doReturnFn
	}{
		Values: &struct {
			Frame runtime.Frame
			More  bool
		}{
			Frame: frame,
			More:  more,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqFrames_starGenType_Next_fnRecorder) AndDo(fn MoqFrames_starGenType_Next_doFn) *MoqFrames_starGenType_Next_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFrames_starGenType_Next_fnRecorder) DoReturnResults(fn MoqFrames_starGenType_Next_doReturnFn) *MoqFrames_starGenType_Next_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Frame runtime.Frame
			More  bool
		}
		Sequence   uint32
		DoFn       MoqFrames_starGenType_Next_doFn
		DoReturnFn MoqFrames_starGenType_Next_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFrames_starGenType_Next_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFrames_starGenType_Next_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Next {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqFrames_starGenType_Next_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFrames_starGenType_Next_paramsKey]*MoqFrames_starGenType_Next_results{},
		}
		r.Moq.ResultsByParams_Next = append(r.Moq.ResultsByParams_Next, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Next) {
			copy(r.Moq.ResultsByParams_Next[insertAt+1:], r.Moq.ResultsByParams_Next[insertAt:0])
			r.Moq.ResultsByParams_Next[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Next(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqFrames_starGenType_Next_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFrames_starGenType_Next_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFrames_starGenType_Next_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults or DoReturnResults must be called before calling Repeat")
		return nil
	}
	r.Results.Repeat.Repeat(r.Moq.Scene.T, repeaters)
	last := r.Results.Results[len(r.Results.Results)-1]
	for n := 0; n < r.Results.Repeat.ResultCount-1; n++ {
		if r.Sequence {
			last = struct {
				Values *struct {
					Frame runtime.Frame
					More  bool
				}
				Sequence   uint32
				DoFn       MoqFrames_starGenType_Next_doFn
				DoReturnFn MoqFrames_starGenType_Next_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFrames_starGenType) PrettyParams_Next(params MoqFrames_starGenType_Next_params) string {
	return fmt.Sprintf("Next()")
}

func (m *MoqFrames_starGenType) ParamsKey_Next(params MoqFrames_starGenType_Next_params, anyParams uint64) MoqFrames_starGenType_Next_paramsKey {
	m.Scene.T.Helper()
	return MoqFrames_starGenType_Next_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqFrames_starGenType) Reset() { m.ResultsByParams_Next = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqFrames_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Next {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Next(results.Params))
			}
		}
	}
}