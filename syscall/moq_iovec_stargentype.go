// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that syscall.Iovec_starGenType is
// mocked completely
var _ Iovec_starGenType = (*MoqIovec_starGenType_mock)(nil)

// Iovec_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type Iovec_starGenType interface {
	SetLen(length int)
}

// MoqIovec_starGenType holds the state of a moq of the Iovec_starGenType type
type MoqIovec_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqIovec_starGenType_mock

	ResultsByParams_SetLen []MoqIovec_starGenType_SetLen_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			SetLen struct {
				Length moq.ParamIndexing
			}
		}
	}
	// MoqIovec_starGenType_mock isolates the mock interface of the
}

// Iovec_starGenType type
type MoqIovec_starGenType_mock struct {
	Moq *MoqIovec_starGenType
}

// MoqIovec_starGenType_recorder isolates the recorder interface of the
// Iovec_starGenType type
type MoqIovec_starGenType_recorder struct {
	Moq *MoqIovec_starGenType
}

// MoqIovec_starGenType_SetLen_params holds the params of the Iovec_starGenType
// type
type MoqIovec_starGenType_SetLen_params struct{ Length int }

// MoqIovec_starGenType_SetLen_paramsKey holds the map key params of the
// Iovec_starGenType type
type MoqIovec_starGenType_SetLen_paramsKey struct {
	Params struct{ Length int }
	Hashes struct{ Length hash.Hash }
}

// MoqIovec_starGenType_SetLen_resultsByParams contains the results for a given
// set of parameters for the Iovec_starGenType type
type MoqIovec_starGenType_SetLen_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqIovec_starGenType_SetLen_paramsKey]*MoqIovec_starGenType_SetLen_results
}

// MoqIovec_starGenType_SetLen_doFn defines the type of function needed when
// calling AndDo for the Iovec_starGenType type
type MoqIovec_starGenType_SetLen_doFn func(length int)

// MoqIovec_starGenType_SetLen_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Iovec_starGenType type
type MoqIovec_starGenType_SetLen_doReturnFn func(length int)

// MoqIovec_starGenType_SetLen_results holds the results of the
// Iovec_starGenType type
type MoqIovec_starGenType_SetLen_results struct {
	Params  MoqIovec_starGenType_SetLen_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqIovec_starGenType_SetLen_doFn
		DoReturnFn MoqIovec_starGenType_SetLen_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqIovec_starGenType_SetLen_fnRecorder routes recorded function calls to the
// MoqIovec_starGenType moq
type MoqIovec_starGenType_SetLen_fnRecorder struct {
	Params    MoqIovec_starGenType_SetLen_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqIovec_starGenType_SetLen_results
	Moq       *MoqIovec_starGenType
}

// MoqIovec_starGenType_SetLen_anyParams isolates the any params functions of
// the Iovec_starGenType type
type MoqIovec_starGenType_SetLen_anyParams struct {
	Recorder *MoqIovec_starGenType_SetLen_fnRecorder
}

// NewMoqIovec_starGenType creates a new moq of the Iovec_starGenType type
func NewMoqIovec_starGenType(scene *moq.Scene, config *moq.Config) *MoqIovec_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqIovec_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqIovec_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				SetLen struct {
					Length moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			SetLen struct {
				Length moq.ParamIndexing
			}
		}{
			SetLen: struct {
				Length moq.ParamIndexing
			}{
				Length: moq.ParamIndexByValue,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Iovec_starGenType type
func (m *MoqIovec_starGenType) Mock() *MoqIovec_starGenType_mock { return m.Moq }

func (m *MoqIovec_starGenType_mock) SetLen(length int) {
	m.Moq.Scene.T.Helper()
	params := MoqIovec_starGenType_SetLen_params{
		Length: length,
	}
	var results *MoqIovec_starGenType_SetLen_results
	for _, resultsByParams := range m.Moq.ResultsByParams_SetLen {
		paramsKey := m.Moq.ParamsKey_SetLen(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_SetLen(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_SetLen(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_SetLen(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(length)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(length)
	}
	return
}

// OnCall returns the recorder implementation of the Iovec_starGenType type
func (m *MoqIovec_starGenType) OnCall() *MoqIovec_starGenType_recorder {
	return &MoqIovec_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqIovec_starGenType_recorder) SetLen(length int) *MoqIovec_starGenType_SetLen_fnRecorder {
	return &MoqIovec_starGenType_SetLen_fnRecorder{
		Params: MoqIovec_starGenType_SetLen_params{
			Length: length,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqIovec_starGenType_SetLen_fnRecorder) Any() *MoqIovec_starGenType_SetLen_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_SetLen(r.Params))
		return nil
	}
	return &MoqIovec_starGenType_SetLen_anyParams{Recorder: r}
}

func (a *MoqIovec_starGenType_SetLen_anyParams) Length() *MoqIovec_starGenType_SetLen_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqIovec_starGenType_SetLen_fnRecorder) Seq() *MoqIovec_starGenType_SetLen_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_SetLen(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqIovec_starGenType_SetLen_fnRecorder) NoSeq() *MoqIovec_starGenType_SetLen_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_SetLen(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqIovec_starGenType_SetLen_fnRecorder) ReturnResults() *MoqIovec_starGenType_SetLen_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqIovec_starGenType_SetLen_doFn
		DoReturnFn MoqIovec_starGenType_SetLen_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqIovec_starGenType_SetLen_fnRecorder) AndDo(fn MoqIovec_starGenType_SetLen_doFn) *MoqIovec_starGenType_SetLen_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqIovec_starGenType_SetLen_fnRecorder) DoReturnResults(fn MoqIovec_starGenType_SetLen_doReturnFn) *MoqIovec_starGenType_SetLen_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqIovec_starGenType_SetLen_doFn
		DoReturnFn MoqIovec_starGenType_SetLen_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqIovec_starGenType_SetLen_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqIovec_starGenType_SetLen_resultsByParams
	for n, res := range r.Moq.ResultsByParams_SetLen {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqIovec_starGenType_SetLen_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqIovec_starGenType_SetLen_paramsKey]*MoqIovec_starGenType_SetLen_results{},
		}
		r.Moq.ResultsByParams_SetLen = append(r.Moq.ResultsByParams_SetLen, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_SetLen) {
			copy(r.Moq.ResultsByParams_SetLen[insertAt+1:], r.Moq.ResultsByParams_SetLen[insertAt:0])
			r.Moq.ResultsByParams_SetLen[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_SetLen(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqIovec_starGenType_SetLen_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqIovec_starGenType_SetLen_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqIovec_starGenType_SetLen_fnRecorder {
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
				Values     *struct{}
				Sequence   uint32
				DoFn       MoqIovec_starGenType_SetLen_doFn
				DoReturnFn MoqIovec_starGenType_SetLen_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqIovec_starGenType) PrettyParams_SetLen(params MoqIovec_starGenType_SetLen_params) string {
	return fmt.Sprintf("SetLen(%#v)", params.Length)
}

func (m *MoqIovec_starGenType) ParamsKey_SetLen(params MoqIovec_starGenType_SetLen_params, anyParams uint64) MoqIovec_starGenType_SetLen_paramsKey {
	m.Scene.T.Helper()
	var lengthUsed int
	var lengthUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.SetLen.Length == moq.ParamIndexByValue {
			lengthUsed = params.Length
		} else {
			lengthUsedHash = hash.DeepHash(params.Length)
		}
	}
	return MoqIovec_starGenType_SetLen_paramsKey{
		Params: struct{ Length int }{
			Length: lengthUsed,
		},
		Hashes: struct{ Length hash.Hash }{
			Length: lengthUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqIovec_starGenType) Reset() { m.ResultsByParams_SetLen = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqIovec_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_SetLen {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_SetLen(results.Params))
			}
		}
	}
}
