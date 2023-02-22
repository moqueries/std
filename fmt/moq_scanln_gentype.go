// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package fmt

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Scanln_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Scanln_genType func(a ...any) (n int, err error)

// MoqScanln_genType holds the state of a moq of the Scanln_genType type
type MoqScanln_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqScanln_genType_mock

	ResultsByParams []MoqScanln_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			A moq.ParamIndexing
		}
	}
}

// MoqScanln_genType_mock isolates the mock interface of the Scanln_genType
// type
type MoqScanln_genType_mock struct {
	Moq *MoqScanln_genType
}

// MoqScanln_genType_params holds the params of the Scanln_genType type
type MoqScanln_genType_params struct{ A []any }

// MoqScanln_genType_paramsKey holds the map key params of the Scanln_genType
// type
type MoqScanln_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ A hash.Hash }
}

// MoqScanln_genType_resultsByParams contains the results for a given set of
// parameters for the Scanln_genType type
type MoqScanln_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqScanln_genType_paramsKey]*MoqScanln_genType_results
}

// MoqScanln_genType_doFn defines the type of function needed when calling
// AndDo for the Scanln_genType type
type MoqScanln_genType_doFn func(a ...any)

// MoqScanln_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Scanln_genType type
type MoqScanln_genType_doReturnFn func(a ...any) (n int, err error)

// MoqScanln_genType_results holds the results of the Scanln_genType type
type MoqScanln_genType_results struct {
	Params  MoqScanln_genType_params
	Results []struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqScanln_genType_doFn
		DoReturnFn MoqScanln_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqScanln_genType_fnRecorder routes recorded function calls to the
// MoqScanln_genType moq
type MoqScanln_genType_fnRecorder struct {
	Params    MoqScanln_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqScanln_genType_results
	Moq       *MoqScanln_genType
}

// MoqScanln_genType_anyParams isolates the any params functions of the
// Scanln_genType type
type MoqScanln_genType_anyParams struct {
	Recorder *MoqScanln_genType_fnRecorder
}

// NewMoqScanln_genType creates a new moq of the Scanln_genType type
func NewMoqScanln_genType(scene *moq.Scene, config *moq.Config) *MoqScanln_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqScanln_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqScanln_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				A moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			A moq.ParamIndexing
		}{
			A: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Scanln_genType type
func (m *MoqScanln_genType) Mock() Scanln_genType {
	return func(a ...any) (_ int, _ error) {
		m.Scene.T.Helper()
		moq := &MoqScanln_genType_mock{Moq: m}
		return moq.Fn(a...)
	}
}

func (m *MoqScanln_genType_mock) Fn(a ...any) (n int, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqScanln_genType_params{
		A: a,
	}
	var results *MoqScanln_genType_results
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
		result.DoFn(a...)
	}

	if result.Values != nil {
		n = result.Values.N
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		n, err = result.DoReturnFn(a...)
	}
	return
}

func (m *MoqScanln_genType) OnCall(a ...any) *MoqScanln_genType_fnRecorder {
	return &MoqScanln_genType_fnRecorder{
		Params: MoqScanln_genType_params{
			A: a,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqScanln_genType_fnRecorder) Any() *MoqScanln_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqScanln_genType_anyParams{Recorder: r}
}

func (a *MoqScanln_genType_anyParams) A() *MoqScanln_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqScanln_genType_fnRecorder) Seq() *MoqScanln_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqScanln_genType_fnRecorder) NoSeq() *MoqScanln_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqScanln_genType_fnRecorder) ReturnResults(n int, err error) *MoqScanln_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqScanln_genType_doFn
		DoReturnFn MoqScanln_genType_doReturnFn
	}{
		Values: &struct {
			N   int
			Err error
		}{
			N:   n,
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqScanln_genType_fnRecorder) AndDo(fn MoqScanln_genType_doFn) *MoqScanln_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqScanln_genType_fnRecorder) DoReturnResults(fn MoqScanln_genType_doReturnFn) *MoqScanln_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqScanln_genType_doFn
		DoReturnFn MoqScanln_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqScanln_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqScanln_genType_resultsByParams
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
		results = &MoqScanln_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqScanln_genType_paramsKey]*MoqScanln_genType_results{},
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
		r.Results = &MoqScanln_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqScanln_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqScanln_genType_fnRecorder {
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
					N   int
					Err error
				}
				Sequence   uint32
				DoFn       MoqScanln_genType_doFn
				DoReturnFn MoqScanln_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqScanln_genType) PrettyParams(params MoqScanln_genType_params) string {
	return fmt.Sprintf("Scanln_genType(%#v)", params.A)
}

func (m *MoqScanln_genType) ParamsKey(params MoqScanln_genType_params, anyParams uint64) MoqScanln_genType_paramsKey {
	m.Scene.T.Helper()
	var aUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.A == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The a parameter can't be indexed by value")
		}
		aUsedHash = hash.DeepHash(params.A)
	}
	return MoqScanln_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ A hash.Hash }{
			A: aUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqScanln_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqScanln_genType) AssertExpectationsMet() {
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
