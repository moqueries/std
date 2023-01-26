// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package pprof

import (
	"fmt"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// WriteHeapProfile_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type WriteHeapProfile_genType func(w io.Writer) error

// MoqWriteHeapProfile_genType holds the state of a moq of the
// WriteHeapProfile_genType type
type MoqWriteHeapProfile_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqWriteHeapProfile_genType_mock

	ResultsByParams []MoqWriteHeapProfile_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			W moq.ParamIndexing
		}
	}
}

// MoqWriteHeapProfile_genType_mock isolates the mock interface of the
// WriteHeapProfile_genType type
type MoqWriteHeapProfile_genType_mock struct {
	Moq *MoqWriteHeapProfile_genType
}

// MoqWriteHeapProfile_genType_params holds the params of the
// WriteHeapProfile_genType type
type MoqWriteHeapProfile_genType_params struct{ W io.Writer }

// MoqWriteHeapProfile_genType_paramsKey holds the map key params of the
// WriteHeapProfile_genType type
type MoqWriteHeapProfile_genType_paramsKey struct {
	Params struct{ W io.Writer }
	Hashes struct{ W hash.Hash }
}

// MoqWriteHeapProfile_genType_resultsByParams contains the results for a given
// set of parameters for the WriteHeapProfile_genType type
type MoqWriteHeapProfile_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqWriteHeapProfile_genType_paramsKey]*MoqWriteHeapProfile_genType_results
}

// MoqWriteHeapProfile_genType_doFn defines the type of function needed when
// calling AndDo for the WriteHeapProfile_genType type
type MoqWriteHeapProfile_genType_doFn func(w io.Writer)

// MoqWriteHeapProfile_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the WriteHeapProfile_genType type
type MoqWriteHeapProfile_genType_doReturnFn func(w io.Writer) error

// MoqWriteHeapProfile_genType_results holds the results of the
// WriteHeapProfile_genType type
type MoqWriteHeapProfile_genType_results struct {
	Params  MoqWriteHeapProfile_genType_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqWriteHeapProfile_genType_doFn
		DoReturnFn MoqWriteHeapProfile_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqWriteHeapProfile_genType_fnRecorder routes recorded function calls to the
// MoqWriteHeapProfile_genType moq
type MoqWriteHeapProfile_genType_fnRecorder struct {
	Params    MoqWriteHeapProfile_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqWriteHeapProfile_genType_results
	Moq       *MoqWriteHeapProfile_genType
}

// MoqWriteHeapProfile_genType_anyParams isolates the any params functions of
// the WriteHeapProfile_genType type
type MoqWriteHeapProfile_genType_anyParams struct {
	Recorder *MoqWriteHeapProfile_genType_fnRecorder
}

// NewMoqWriteHeapProfile_genType creates a new moq of the
// WriteHeapProfile_genType type
func NewMoqWriteHeapProfile_genType(scene *moq.Scene, config *moq.Config) *MoqWriteHeapProfile_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqWriteHeapProfile_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqWriteHeapProfile_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				W moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			W moq.ParamIndexing
		}{
			W: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the WriteHeapProfile_genType type
func (m *MoqWriteHeapProfile_genType) Mock() WriteHeapProfile_genType {
	return func(w io.Writer) error {
		m.Scene.T.Helper()
		moq := &MoqWriteHeapProfile_genType_mock{Moq: m}
		return moq.Fn(w)
	}
}

func (m *MoqWriteHeapProfile_genType_mock) Fn(w io.Writer) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqWriteHeapProfile_genType_params{
		W: w,
	}
	var results *MoqWriteHeapProfile_genType_results
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
		result.DoFn(w)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(w)
	}
	return
}

func (m *MoqWriteHeapProfile_genType) OnCall(w io.Writer) *MoqWriteHeapProfile_genType_fnRecorder {
	return &MoqWriteHeapProfile_genType_fnRecorder{
		Params: MoqWriteHeapProfile_genType_params{
			W: w,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqWriteHeapProfile_genType_fnRecorder) Any() *MoqWriteHeapProfile_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqWriteHeapProfile_genType_anyParams{Recorder: r}
}

func (a *MoqWriteHeapProfile_genType_anyParams) W() *MoqWriteHeapProfile_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqWriteHeapProfile_genType_fnRecorder) Seq() *MoqWriteHeapProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqWriteHeapProfile_genType_fnRecorder) NoSeq() *MoqWriteHeapProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqWriteHeapProfile_genType_fnRecorder) ReturnResults(result1 error) *MoqWriteHeapProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqWriteHeapProfile_genType_doFn
		DoReturnFn MoqWriteHeapProfile_genType_doReturnFn
	}{
		Values: &struct {
			Result1 error
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqWriteHeapProfile_genType_fnRecorder) AndDo(fn MoqWriteHeapProfile_genType_doFn) *MoqWriteHeapProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqWriteHeapProfile_genType_fnRecorder) DoReturnResults(fn MoqWriteHeapProfile_genType_doReturnFn) *MoqWriteHeapProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqWriteHeapProfile_genType_doFn
		DoReturnFn MoqWriteHeapProfile_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqWriteHeapProfile_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqWriteHeapProfile_genType_resultsByParams
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
		results = &MoqWriteHeapProfile_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqWriteHeapProfile_genType_paramsKey]*MoqWriteHeapProfile_genType_results{},
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
		r.Results = &MoqWriteHeapProfile_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqWriteHeapProfile_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqWriteHeapProfile_genType_fnRecorder {
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
					Result1 error
				}
				Sequence   uint32
				DoFn       MoqWriteHeapProfile_genType_doFn
				DoReturnFn MoqWriteHeapProfile_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqWriteHeapProfile_genType) PrettyParams(params MoqWriteHeapProfile_genType_params) string {
	return fmt.Sprintf("WriteHeapProfile_genType(%#v)", params.W)
}

func (m *MoqWriteHeapProfile_genType) ParamsKey(params MoqWriteHeapProfile_genType_params, anyParams uint64) MoqWriteHeapProfile_genType_paramsKey {
	m.Scene.T.Helper()
	var wUsed io.Writer
	var wUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.W == moq.ParamIndexByValue {
			wUsed = params.W
		} else {
			wUsedHash = hash.DeepHash(params.W)
		}
	}
	return MoqWriteHeapProfile_genType_paramsKey{
		Params: struct{ W io.Writer }{
			W: wUsed,
		},
		Hashes: struct{ W hash.Hash }{
			W: wUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqWriteHeapProfile_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqWriteHeapProfile_genType) AssertExpectationsMet() {
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
