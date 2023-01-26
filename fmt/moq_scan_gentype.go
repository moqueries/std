// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package fmt

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Scan_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Scan_genType func(a ...interface{}) (n int, err error)

// MoqScan_genType holds the state of a moq of the Scan_genType type
type MoqScan_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqScan_genType_mock

	ResultsByParams []MoqScan_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			A moq.ParamIndexing
		}
	}
}

// MoqScan_genType_mock isolates the mock interface of the Scan_genType type
type MoqScan_genType_mock struct {
	Moq *MoqScan_genType
}

// MoqScan_genType_params holds the params of the Scan_genType type
type MoqScan_genType_params struct{ A []interface{} }

// MoqScan_genType_paramsKey holds the map key params of the Scan_genType type
type MoqScan_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ A hash.Hash }
}

// MoqScan_genType_resultsByParams contains the results for a given set of
// parameters for the Scan_genType type
type MoqScan_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqScan_genType_paramsKey]*MoqScan_genType_results
}

// MoqScan_genType_doFn defines the type of function needed when calling AndDo
// for the Scan_genType type
type MoqScan_genType_doFn func(a ...interface{})

// MoqScan_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Scan_genType type
type MoqScan_genType_doReturnFn func(a ...interface{}) (n int, err error)

// MoqScan_genType_results holds the results of the Scan_genType type
type MoqScan_genType_results struct {
	Params  MoqScan_genType_params
	Results []struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqScan_genType_doFn
		DoReturnFn MoqScan_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqScan_genType_fnRecorder routes recorded function calls to the
// MoqScan_genType moq
type MoqScan_genType_fnRecorder struct {
	Params    MoqScan_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqScan_genType_results
	Moq       *MoqScan_genType
}

// MoqScan_genType_anyParams isolates the any params functions of the
// Scan_genType type
type MoqScan_genType_anyParams struct {
	Recorder *MoqScan_genType_fnRecorder
}

// NewMoqScan_genType creates a new moq of the Scan_genType type
func NewMoqScan_genType(scene *moq.Scene, config *moq.Config) *MoqScan_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqScan_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqScan_genType_mock{},

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

// Mock returns the moq implementation of the Scan_genType type
func (m *MoqScan_genType) Mock() Scan_genType {
	return func(a ...interface{}) (_ int, _ error) {
		m.Scene.T.Helper()
		moq := &MoqScan_genType_mock{Moq: m}
		return moq.Fn(a...)
	}
}

func (m *MoqScan_genType_mock) Fn(a ...interface{}) (n int, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqScan_genType_params{
		A: a,
	}
	var results *MoqScan_genType_results
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

func (m *MoqScan_genType) OnCall(a ...interface{}) *MoqScan_genType_fnRecorder {
	return &MoqScan_genType_fnRecorder{
		Params: MoqScan_genType_params{
			A: a,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqScan_genType_fnRecorder) Any() *MoqScan_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqScan_genType_anyParams{Recorder: r}
}

func (a *MoqScan_genType_anyParams) A() *MoqScan_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqScan_genType_fnRecorder) Seq() *MoqScan_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqScan_genType_fnRecorder) NoSeq() *MoqScan_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqScan_genType_fnRecorder) ReturnResults(n int, err error) *MoqScan_genType_fnRecorder {
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
		DoFn       MoqScan_genType_doFn
		DoReturnFn MoqScan_genType_doReturnFn
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

func (r *MoqScan_genType_fnRecorder) AndDo(fn MoqScan_genType_doFn) *MoqScan_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqScan_genType_fnRecorder) DoReturnResults(fn MoqScan_genType_doReturnFn) *MoqScan_genType_fnRecorder {
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
		DoFn       MoqScan_genType_doFn
		DoReturnFn MoqScan_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqScan_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqScan_genType_resultsByParams
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
		results = &MoqScan_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqScan_genType_paramsKey]*MoqScan_genType_results{},
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
		r.Results = &MoqScan_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqScan_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqScan_genType_fnRecorder {
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
				DoFn       MoqScan_genType_doFn
				DoReturnFn MoqScan_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqScan_genType) PrettyParams(params MoqScan_genType_params) string {
	return fmt.Sprintf("Scan_genType(%#v)", params.A)
}

func (m *MoqScan_genType) ParamsKey(params MoqScan_genType_params, anyParams uint64) MoqScan_genType_paramsKey {
	m.Scene.T.Helper()
	var aUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.A == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The a parameter can't be indexed by value")
		}
		aUsedHash = hash.DeepHash(params.A)
	}
	return MoqScan_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ A hash.Hash }{
			A: aUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqScan_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqScan_genType) AssertExpectationsMet() {
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
