// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that syscall.Msghdr_starGenType is
// mocked completely
var _ Msghdr_starGenType = (*MoqMsghdr_starGenType_mock)(nil)

// Msghdr_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type Msghdr_starGenType interface {
	SetControllen(length int)
}

// MoqMsghdr_starGenType holds the state of a moq of the Msghdr_starGenType
// type
type MoqMsghdr_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqMsghdr_starGenType_mock

	ResultsByParams_SetControllen []MoqMsghdr_starGenType_SetControllen_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			SetControllen struct {
				Length moq.ParamIndexing
			}
		}
	}
	// MoqMsghdr_starGenType_mock isolates the mock interface of the
}

// Msghdr_starGenType type
type MoqMsghdr_starGenType_mock struct {
	Moq *MoqMsghdr_starGenType
}

// MoqMsghdr_starGenType_recorder isolates the recorder interface of the
// Msghdr_starGenType type
type MoqMsghdr_starGenType_recorder struct {
	Moq *MoqMsghdr_starGenType
}

// MoqMsghdr_starGenType_SetControllen_params holds the params of the
// Msghdr_starGenType type
type MoqMsghdr_starGenType_SetControllen_params struct{ Length int }

// MoqMsghdr_starGenType_SetControllen_paramsKey holds the map key params of
// the Msghdr_starGenType type
type MoqMsghdr_starGenType_SetControllen_paramsKey struct {
	Params struct{ Length int }
	Hashes struct{ Length hash.Hash }
}

// MoqMsghdr_starGenType_SetControllen_resultsByParams contains the results for
// a given set of parameters for the Msghdr_starGenType type
type MoqMsghdr_starGenType_SetControllen_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqMsghdr_starGenType_SetControllen_paramsKey]*MoqMsghdr_starGenType_SetControllen_results
}

// MoqMsghdr_starGenType_SetControllen_doFn defines the type of function needed
// when calling AndDo for the Msghdr_starGenType type
type MoqMsghdr_starGenType_SetControllen_doFn func(length int)

// MoqMsghdr_starGenType_SetControllen_doReturnFn defines the type of function
// needed when calling DoReturnResults for the Msghdr_starGenType type
type MoqMsghdr_starGenType_SetControllen_doReturnFn func(length int)

// MoqMsghdr_starGenType_SetControllen_results holds the results of the
// Msghdr_starGenType type
type MoqMsghdr_starGenType_SetControllen_results struct {
	Params  MoqMsghdr_starGenType_SetControllen_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqMsghdr_starGenType_SetControllen_doFn
		DoReturnFn MoqMsghdr_starGenType_SetControllen_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqMsghdr_starGenType_SetControllen_fnRecorder routes recorded function
// calls to the MoqMsghdr_starGenType moq
type MoqMsghdr_starGenType_SetControllen_fnRecorder struct {
	Params    MoqMsghdr_starGenType_SetControllen_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqMsghdr_starGenType_SetControllen_results
	Moq       *MoqMsghdr_starGenType
}

// MoqMsghdr_starGenType_SetControllen_anyParams isolates the any params
// functions of the Msghdr_starGenType type
type MoqMsghdr_starGenType_SetControllen_anyParams struct {
	Recorder *MoqMsghdr_starGenType_SetControllen_fnRecorder
}

// NewMoqMsghdr_starGenType creates a new moq of the Msghdr_starGenType type
func NewMoqMsghdr_starGenType(scene *moq.Scene, config *moq.Config) *MoqMsghdr_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqMsghdr_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqMsghdr_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				SetControllen struct {
					Length moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			SetControllen struct {
				Length moq.ParamIndexing
			}
		}{
			SetControllen: struct {
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

// Mock returns the mock implementation of the Msghdr_starGenType type
func (m *MoqMsghdr_starGenType) Mock() *MoqMsghdr_starGenType_mock { return m.Moq }

func (m *MoqMsghdr_starGenType_mock) SetControllen(length int) {
	m.Moq.Scene.T.Helper()
	params := MoqMsghdr_starGenType_SetControllen_params{
		Length: length,
	}
	var results *MoqMsghdr_starGenType_SetControllen_results
	for _, resultsByParams := range m.Moq.ResultsByParams_SetControllen {
		paramsKey := m.Moq.ParamsKey_SetControllen(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_SetControllen(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_SetControllen(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_SetControllen(params))
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

// OnCall returns the recorder implementation of the Msghdr_starGenType type
func (m *MoqMsghdr_starGenType) OnCall() *MoqMsghdr_starGenType_recorder {
	return &MoqMsghdr_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqMsghdr_starGenType_recorder) SetControllen(length int) *MoqMsghdr_starGenType_SetControllen_fnRecorder {
	return &MoqMsghdr_starGenType_SetControllen_fnRecorder{
		Params: MoqMsghdr_starGenType_SetControllen_params{
			Length: length,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqMsghdr_starGenType_SetControllen_fnRecorder) Any() *MoqMsghdr_starGenType_SetControllen_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_SetControllen(r.Params))
		return nil
	}
	return &MoqMsghdr_starGenType_SetControllen_anyParams{Recorder: r}
}

func (a *MoqMsghdr_starGenType_SetControllen_anyParams) Length() *MoqMsghdr_starGenType_SetControllen_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqMsghdr_starGenType_SetControllen_fnRecorder) Seq() *MoqMsghdr_starGenType_SetControllen_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_SetControllen(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqMsghdr_starGenType_SetControllen_fnRecorder) NoSeq() *MoqMsghdr_starGenType_SetControllen_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_SetControllen(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqMsghdr_starGenType_SetControllen_fnRecorder) ReturnResults() *MoqMsghdr_starGenType_SetControllen_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqMsghdr_starGenType_SetControllen_doFn
		DoReturnFn MoqMsghdr_starGenType_SetControllen_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqMsghdr_starGenType_SetControllen_fnRecorder) AndDo(fn MoqMsghdr_starGenType_SetControllen_doFn) *MoqMsghdr_starGenType_SetControllen_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqMsghdr_starGenType_SetControllen_fnRecorder) DoReturnResults(fn MoqMsghdr_starGenType_SetControllen_doReturnFn) *MoqMsghdr_starGenType_SetControllen_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqMsghdr_starGenType_SetControllen_doFn
		DoReturnFn MoqMsghdr_starGenType_SetControllen_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqMsghdr_starGenType_SetControllen_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqMsghdr_starGenType_SetControllen_resultsByParams
	for n, res := range r.Moq.ResultsByParams_SetControllen {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqMsghdr_starGenType_SetControllen_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqMsghdr_starGenType_SetControllen_paramsKey]*MoqMsghdr_starGenType_SetControllen_results{},
		}
		r.Moq.ResultsByParams_SetControllen = append(r.Moq.ResultsByParams_SetControllen, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_SetControllen) {
			copy(r.Moq.ResultsByParams_SetControllen[insertAt+1:], r.Moq.ResultsByParams_SetControllen[insertAt:0])
			r.Moq.ResultsByParams_SetControllen[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_SetControllen(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqMsghdr_starGenType_SetControllen_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqMsghdr_starGenType_SetControllen_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqMsghdr_starGenType_SetControllen_fnRecorder {
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
				DoFn       MoqMsghdr_starGenType_SetControllen_doFn
				DoReturnFn MoqMsghdr_starGenType_SetControllen_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqMsghdr_starGenType) PrettyParams_SetControllen(params MoqMsghdr_starGenType_SetControllen_params) string {
	return fmt.Sprintf("SetControllen(%#v)", params.Length)
}

func (m *MoqMsghdr_starGenType) ParamsKey_SetControllen(params MoqMsghdr_starGenType_SetControllen_params, anyParams uint64) MoqMsghdr_starGenType_SetControllen_paramsKey {
	m.Scene.T.Helper()
	var lengthUsed int
	var lengthUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.SetControllen.Length == moq.ParamIndexByValue {
			lengthUsed = params.Length
		} else {
			lengthUsedHash = hash.DeepHash(params.Length)
		}
	}
	return MoqMsghdr_starGenType_SetControllen_paramsKey{
		Params: struct{ Length int }{
			Length: lengthUsed,
		},
		Hashes: struct{ Length hash.Hash }{
			Length: lengthUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqMsghdr_starGenType) Reset() { m.ResultsByParams_SetControllen = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqMsghdr_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_SetControllen {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_SetControllen(results.Params))
			}
		}
	}
}