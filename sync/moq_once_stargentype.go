// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package sync

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that sync.Once_starGenType is mocked
// completely
var _ Once_starGenType = (*MoqOnce_starGenType_mock)(nil)

// Once_starGenType is the fabricated implementation type of this mock (emitted
// when mocking a collections of methods directly and not from an interface
// type)
type Once_starGenType interface {
	Do(f func())
}

// MoqOnce_starGenType holds the state of a moq of the Once_starGenType type
type MoqOnce_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqOnce_starGenType_mock

	ResultsByParams_Do []MoqOnce_starGenType_Do_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Do struct {
				F moq.ParamIndexing
			}
		}
	}
	// MoqOnce_starGenType_mock isolates the mock interface of the Once_starGenType
}

// type
type MoqOnce_starGenType_mock struct {
	Moq *MoqOnce_starGenType
}

// MoqOnce_starGenType_recorder isolates the recorder interface of the
// Once_starGenType type
type MoqOnce_starGenType_recorder struct {
	Moq *MoqOnce_starGenType
}

// MoqOnce_starGenType_Do_params holds the params of the Once_starGenType type
type MoqOnce_starGenType_Do_params struct{ F func() }

// MoqOnce_starGenType_Do_paramsKey holds the map key params of the
// Once_starGenType type
type MoqOnce_starGenType_Do_paramsKey struct {
	Params struct{}
	Hashes struct{ F hash.Hash }
}

// MoqOnce_starGenType_Do_resultsByParams contains the results for a given set
// of parameters for the Once_starGenType type
type MoqOnce_starGenType_Do_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqOnce_starGenType_Do_paramsKey]*MoqOnce_starGenType_Do_results
}

// MoqOnce_starGenType_Do_doFn defines the type of function needed when calling
// AndDo for the Once_starGenType type
type MoqOnce_starGenType_Do_doFn func(f func())

// MoqOnce_starGenType_Do_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Once_starGenType type
type MoqOnce_starGenType_Do_doReturnFn func(f func())

// MoqOnce_starGenType_Do_results holds the results of the Once_starGenType
// type
type MoqOnce_starGenType_Do_results struct {
	Params  MoqOnce_starGenType_Do_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqOnce_starGenType_Do_doFn
		DoReturnFn MoqOnce_starGenType_Do_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqOnce_starGenType_Do_fnRecorder routes recorded function calls to the
// MoqOnce_starGenType moq
type MoqOnce_starGenType_Do_fnRecorder struct {
	Params    MoqOnce_starGenType_Do_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqOnce_starGenType_Do_results
	Moq       *MoqOnce_starGenType
}

// MoqOnce_starGenType_Do_anyParams isolates the any params functions of the
// Once_starGenType type
type MoqOnce_starGenType_Do_anyParams struct {
	Recorder *MoqOnce_starGenType_Do_fnRecorder
}

// NewMoqOnce_starGenType creates a new moq of the Once_starGenType type
func NewMoqOnce_starGenType(scene *moq.Scene, config *moq.Config) *MoqOnce_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqOnce_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqOnce_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Do struct {
					F moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Do struct {
				F moq.ParamIndexing
			}
		}{
			Do: struct {
				F moq.ParamIndexing
			}{
				F: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Once_starGenType type
func (m *MoqOnce_starGenType) Mock() *MoqOnce_starGenType_mock { return m.Moq }

func (m *MoqOnce_starGenType_mock) Do(f func()) {
	m.Moq.Scene.T.Helper()
	params := MoqOnce_starGenType_Do_params{
		F: f,
	}
	var results *MoqOnce_starGenType_Do_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Do {
		paramsKey := m.Moq.ParamsKey_Do(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Do(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Do(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Do(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(f)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(f)
	}
	return
}

// OnCall returns the recorder implementation of the Once_starGenType type
func (m *MoqOnce_starGenType) OnCall() *MoqOnce_starGenType_recorder {
	return &MoqOnce_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqOnce_starGenType_recorder) Do(f func()) *MoqOnce_starGenType_Do_fnRecorder {
	return &MoqOnce_starGenType_Do_fnRecorder{
		Params: MoqOnce_starGenType_Do_params{
			F: f,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqOnce_starGenType_Do_fnRecorder) Any() *MoqOnce_starGenType_Do_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Do(r.Params))
		return nil
	}
	return &MoqOnce_starGenType_Do_anyParams{Recorder: r}
}

func (a *MoqOnce_starGenType_Do_anyParams) F() *MoqOnce_starGenType_Do_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqOnce_starGenType_Do_fnRecorder) Seq() *MoqOnce_starGenType_Do_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Do(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqOnce_starGenType_Do_fnRecorder) NoSeq() *MoqOnce_starGenType_Do_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Do(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqOnce_starGenType_Do_fnRecorder) ReturnResults() *MoqOnce_starGenType_Do_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqOnce_starGenType_Do_doFn
		DoReturnFn MoqOnce_starGenType_Do_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqOnce_starGenType_Do_fnRecorder) AndDo(fn MoqOnce_starGenType_Do_doFn) *MoqOnce_starGenType_Do_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqOnce_starGenType_Do_fnRecorder) DoReturnResults(fn MoqOnce_starGenType_Do_doReturnFn) *MoqOnce_starGenType_Do_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqOnce_starGenType_Do_doFn
		DoReturnFn MoqOnce_starGenType_Do_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqOnce_starGenType_Do_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqOnce_starGenType_Do_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Do {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqOnce_starGenType_Do_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqOnce_starGenType_Do_paramsKey]*MoqOnce_starGenType_Do_results{},
		}
		r.Moq.ResultsByParams_Do = append(r.Moq.ResultsByParams_Do, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Do) {
			copy(r.Moq.ResultsByParams_Do[insertAt+1:], r.Moq.ResultsByParams_Do[insertAt:0])
			r.Moq.ResultsByParams_Do[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Do(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqOnce_starGenType_Do_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqOnce_starGenType_Do_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqOnce_starGenType_Do_fnRecorder {
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
				DoFn       MoqOnce_starGenType_Do_doFn
				DoReturnFn MoqOnce_starGenType_Do_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqOnce_starGenType) PrettyParams_Do(params MoqOnce_starGenType_Do_params) string {
	return fmt.Sprintf("Do(%#v)", params.F)
}

func (m *MoqOnce_starGenType) ParamsKey_Do(params MoqOnce_starGenType_Do_params, anyParams uint64) MoqOnce_starGenType_Do_paramsKey {
	m.Scene.T.Helper()
	var fUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Do.F == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The f parameter of the Do function can't be indexed by value")
		}
		fUsedHash = hash.DeepHash(params.F)
	}
	return MoqOnce_starGenType_Do_paramsKey{
		Params: struct{}{},
		Hashes: struct{ F hash.Hash }{
			F: fUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqOnce_starGenType) Reset() { m.ResultsByParams_Do = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqOnce_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Do {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Do(results.Params))
			}
		}
	}
}
