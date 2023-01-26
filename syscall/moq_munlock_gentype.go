// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Munlock_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Munlock_genType func(b []byte) (err error)

// MoqMunlock_genType holds the state of a moq of the Munlock_genType type
type MoqMunlock_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqMunlock_genType_mock

	ResultsByParams []MoqMunlock_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			B moq.ParamIndexing
		}
	}
}

// MoqMunlock_genType_mock isolates the mock interface of the Munlock_genType
// type
type MoqMunlock_genType_mock struct {
	Moq *MoqMunlock_genType
}

// MoqMunlock_genType_params holds the params of the Munlock_genType type
type MoqMunlock_genType_params struct{ B []byte }

// MoqMunlock_genType_paramsKey holds the map key params of the Munlock_genType
// type
type MoqMunlock_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ B hash.Hash }
}

// MoqMunlock_genType_resultsByParams contains the results for a given set of
// parameters for the Munlock_genType type
type MoqMunlock_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqMunlock_genType_paramsKey]*MoqMunlock_genType_results
}

// MoqMunlock_genType_doFn defines the type of function needed when calling
// AndDo for the Munlock_genType type
type MoqMunlock_genType_doFn func(b []byte)

// MoqMunlock_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Munlock_genType type
type MoqMunlock_genType_doReturnFn func(b []byte) (err error)

// MoqMunlock_genType_results holds the results of the Munlock_genType type
type MoqMunlock_genType_results struct {
	Params  MoqMunlock_genType_params
	Results []struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqMunlock_genType_doFn
		DoReturnFn MoqMunlock_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqMunlock_genType_fnRecorder routes recorded function calls to the
// MoqMunlock_genType moq
type MoqMunlock_genType_fnRecorder struct {
	Params    MoqMunlock_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqMunlock_genType_results
	Moq       *MoqMunlock_genType
}

// MoqMunlock_genType_anyParams isolates the any params functions of the
// Munlock_genType type
type MoqMunlock_genType_anyParams struct {
	Recorder *MoqMunlock_genType_fnRecorder
}

// NewMoqMunlock_genType creates a new moq of the Munlock_genType type
func NewMoqMunlock_genType(scene *moq.Scene, config *moq.Config) *MoqMunlock_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqMunlock_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqMunlock_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				B moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			B moq.ParamIndexing
		}{
			B: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Munlock_genType type
func (m *MoqMunlock_genType) Mock() Munlock_genType {
	return func(b []byte) (_ error) {
		m.Scene.T.Helper()
		moq := &MoqMunlock_genType_mock{Moq: m}
		return moq.Fn(b)
	}
}

func (m *MoqMunlock_genType_mock) Fn(b []byte) (err error) {
	m.Moq.Scene.T.Helper()
	params := MoqMunlock_genType_params{
		B: b,
	}
	var results *MoqMunlock_genType_results
	for _, resultsByParams := range m.Moq.ResultsByParams {
		paramsKey := m.Moq.ParamsKey(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(b)
	}

	if result.Values != nil {
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		err = result.DoReturnFn(b)
	}
	return
}

func (m *MoqMunlock_genType) OnCall(b []byte) *MoqMunlock_genType_fnRecorder {
	return &MoqMunlock_genType_fnRecorder{
		Params: MoqMunlock_genType_params{
			B: b,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqMunlock_genType_fnRecorder) Any() *MoqMunlock_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqMunlock_genType_anyParams{Recorder: r}
}

func (a *MoqMunlock_genType_anyParams) B() *MoqMunlock_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqMunlock_genType_fnRecorder) Seq() *MoqMunlock_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqMunlock_genType_fnRecorder) NoSeq() *MoqMunlock_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqMunlock_genType_fnRecorder) ReturnResults(err error) *MoqMunlock_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqMunlock_genType_doFn
		DoReturnFn MoqMunlock_genType_doReturnFn
	}{
		Values: &struct{ Err error }{
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqMunlock_genType_fnRecorder) AndDo(fn MoqMunlock_genType_doFn) *MoqMunlock_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqMunlock_genType_fnRecorder) DoReturnResults(fn MoqMunlock_genType_doReturnFn) *MoqMunlock_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqMunlock_genType_doFn
		DoReturnFn MoqMunlock_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqMunlock_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqMunlock_genType_resultsByParams
	for n, res := range r.Moq.ResultsByParams {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqMunlock_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqMunlock_genType_paramsKey]*MoqMunlock_genType_results{},
		}
		r.Moq.ResultsByParams = append(r.Moq.ResultsByParams, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams) {
			copy(r.Moq.ResultsByParams[insertAt+1:], r.Moq.ResultsByParams[insertAt:0])
			r.Moq.ResultsByParams[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqMunlock_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqMunlock_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqMunlock_genType_fnRecorder {
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
				Values     *struct{ Err error }
				Sequence   uint32
				DoFn       MoqMunlock_genType_doFn
				DoReturnFn MoqMunlock_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqMunlock_genType) PrettyParams(params MoqMunlock_genType_params) string {
	return fmt.Sprintf("Munlock_genType(%#v)", params.B)
}

func (m *MoqMunlock_genType) ParamsKey(params MoqMunlock_genType_params, anyParams uint64) MoqMunlock_genType_paramsKey {
	m.Scene.T.Helper()
	var bUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.B == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The b parameter can't be indexed by value")
		}
		bUsedHash = hash.DeepHash(params.B)
	}
	return MoqMunlock_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ B hash.Hash }{
			B: bUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqMunlock_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqMunlock_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams(results.Params))
			}
		}
	}
}