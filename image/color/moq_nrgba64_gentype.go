// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package color

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that color.NRGBA64_genType is mocked
// completely
var _ NRGBA64_genType = (*MoqNRGBA64_genType_mock)(nil)

// NRGBA64_genType is the fabricated implementation type of this mock (emitted
// when mocking a collections of methods directly and not from an interface
// type)
type NRGBA64_genType interface {
	RGBA() (r, g, b, a uint32)
}

// MoqNRGBA64_genType holds the state of a moq of the NRGBA64_genType type
type MoqNRGBA64_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNRGBA64_genType_mock

	ResultsByParams_RGBA []MoqNRGBA64_genType_RGBA_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			RGBA struct{}
		}
	}
}

// MoqNRGBA64_genType_mock isolates the mock interface of the NRGBA64_genType
// type
type MoqNRGBA64_genType_mock struct {
	Moq *MoqNRGBA64_genType
}

// MoqNRGBA64_genType_recorder isolates the recorder interface of the
// NRGBA64_genType type
type MoqNRGBA64_genType_recorder struct {
	Moq *MoqNRGBA64_genType
}

// MoqNRGBA64_genType_RGBA_params holds the params of the NRGBA64_genType type
type MoqNRGBA64_genType_RGBA_params struct{}

// MoqNRGBA64_genType_RGBA_paramsKey holds the map key params of the
// NRGBA64_genType type
type MoqNRGBA64_genType_RGBA_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqNRGBA64_genType_RGBA_resultsByParams contains the results for a given set
// of parameters for the NRGBA64_genType type
type MoqNRGBA64_genType_RGBA_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNRGBA64_genType_RGBA_paramsKey]*MoqNRGBA64_genType_RGBA_results
}

// MoqNRGBA64_genType_RGBA_doFn defines the type of function needed when
// calling AndDo for the NRGBA64_genType type
type MoqNRGBA64_genType_RGBA_doFn func()

// MoqNRGBA64_genType_RGBA_doReturnFn defines the type of function needed when
// calling DoReturnResults for the NRGBA64_genType type
type MoqNRGBA64_genType_RGBA_doReturnFn func() (r, g, b, a uint32)

// MoqNRGBA64_genType_RGBA_results holds the results of the NRGBA64_genType
// type
type MoqNRGBA64_genType_RGBA_results struct {
	Params  MoqNRGBA64_genType_RGBA_params
	Results []struct {
		Values     *struct{ Result1, G, B, A uint32 }
		Sequence   uint32
		DoFn       MoqNRGBA64_genType_RGBA_doFn
		DoReturnFn MoqNRGBA64_genType_RGBA_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNRGBA64_genType_RGBA_fnRecorder routes recorded function calls to the
// MoqNRGBA64_genType moq
type MoqNRGBA64_genType_RGBA_fnRecorder struct {
	Params    MoqNRGBA64_genType_RGBA_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNRGBA64_genType_RGBA_results
	Moq       *MoqNRGBA64_genType
}

// MoqNRGBA64_genType_RGBA_anyParams isolates the any params functions of the
// NRGBA64_genType type
type MoqNRGBA64_genType_RGBA_anyParams struct {
	Recorder *MoqNRGBA64_genType_RGBA_fnRecorder
}

// NewMoqNRGBA64_genType creates a new moq of the NRGBA64_genType type
func NewMoqNRGBA64_genType(scene *moq.Scene, config *moq.Config) *MoqNRGBA64_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNRGBA64_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNRGBA64_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				RGBA struct{}
			}
		}{ParameterIndexing: struct {
			RGBA struct{}
		}{
			RGBA: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the NRGBA64_genType type
func (m *MoqNRGBA64_genType) Mock() *MoqNRGBA64_genType_mock { return m.Moq }

func (m *MoqNRGBA64_genType_mock) RGBA() (result1, g, b, a uint32) {
	m.Moq.Scene.T.Helper()
	params := MoqNRGBA64_genType_RGBA_params{}
	var results *MoqNRGBA64_genType_RGBA_results
	for _, resultsByParams := range m.Moq.ResultsByParams_RGBA {
		paramsKey := m.Moq.ParamsKey_RGBA(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_RGBA(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_RGBA(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_RGBA(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		g = result.Values.G
		b = result.Values.B
		a = result.Values.A
	}
	if result.DoReturnFn != nil {
		result1, g, b, a = result.DoReturnFn()
	}
	return
}

// OnCall returns the recorder implementation of the NRGBA64_genType type
func (m *MoqNRGBA64_genType) OnCall() *MoqNRGBA64_genType_recorder {
	return &MoqNRGBA64_genType_recorder{
		Moq: m,
	}
}

func (m *MoqNRGBA64_genType_recorder) RGBA() *MoqNRGBA64_genType_RGBA_fnRecorder {
	return &MoqNRGBA64_genType_RGBA_fnRecorder{
		Params:   MoqNRGBA64_genType_RGBA_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqNRGBA64_genType_RGBA_fnRecorder) Any() *MoqNRGBA64_genType_RGBA_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_RGBA(r.Params))
		return nil
	}
	return &MoqNRGBA64_genType_RGBA_anyParams{Recorder: r}
}

func (r *MoqNRGBA64_genType_RGBA_fnRecorder) Seq() *MoqNRGBA64_genType_RGBA_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_RGBA(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNRGBA64_genType_RGBA_fnRecorder) NoSeq() *MoqNRGBA64_genType_RGBA_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_RGBA(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNRGBA64_genType_RGBA_fnRecorder) ReturnResults(result1, g, b, a uint32) *MoqNRGBA64_genType_RGBA_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Result1, G, B, A uint32 }
		Sequence   uint32
		DoFn       MoqNRGBA64_genType_RGBA_doFn
		DoReturnFn MoqNRGBA64_genType_RGBA_doReturnFn
	}{
		Values: &struct{ Result1, G, B, A uint32 }{
			Result1: result1,
			G:       g,
			B:       b,
			A:       a,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNRGBA64_genType_RGBA_fnRecorder) AndDo(fn MoqNRGBA64_genType_RGBA_doFn) *MoqNRGBA64_genType_RGBA_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNRGBA64_genType_RGBA_fnRecorder) DoReturnResults(fn MoqNRGBA64_genType_RGBA_doReturnFn) *MoqNRGBA64_genType_RGBA_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Result1, G, B, A uint32 }
		Sequence   uint32
		DoFn       MoqNRGBA64_genType_RGBA_doFn
		DoReturnFn MoqNRGBA64_genType_RGBA_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNRGBA64_genType_RGBA_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNRGBA64_genType_RGBA_resultsByParams
	for n, res := range r.Moq.ResultsByParams_RGBA {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqNRGBA64_genType_RGBA_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNRGBA64_genType_RGBA_paramsKey]*MoqNRGBA64_genType_RGBA_results{},
		}
		r.Moq.ResultsByParams_RGBA = append(r.Moq.ResultsByParams_RGBA, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_RGBA) {
			copy(r.Moq.ResultsByParams_RGBA[insertAt+1:], r.Moq.ResultsByParams_RGBA[insertAt:0])
			r.Moq.ResultsByParams_RGBA[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_RGBA(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqNRGBA64_genType_RGBA_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNRGBA64_genType_RGBA_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNRGBA64_genType_RGBA_fnRecorder {
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
				Values     *struct{ Result1, G, B, A uint32 }
				Sequence   uint32
				DoFn       MoqNRGBA64_genType_RGBA_doFn
				DoReturnFn MoqNRGBA64_genType_RGBA_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNRGBA64_genType) PrettyParams_RGBA(params MoqNRGBA64_genType_RGBA_params) string {
	return fmt.Sprintf("RGBA()")
}

func (m *MoqNRGBA64_genType) ParamsKey_RGBA(params MoqNRGBA64_genType_RGBA_params, anyParams uint64) MoqNRGBA64_genType_RGBA_paramsKey {
	m.Scene.T.Helper()
	return MoqNRGBA64_genType_RGBA_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqNRGBA64_genType) Reset() { m.ResultsByParams_RGBA = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNRGBA64_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_RGBA {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_RGBA(results.Params))
			}
		}
	}
}
