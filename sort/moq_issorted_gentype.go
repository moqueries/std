// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package sort

import (
	"fmt"
	"math/bits"
	"sort"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// IsSorted_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type IsSorted_genType func(data sort.Interface) bool

// MoqIsSorted_genType holds the state of a moq of the IsSorted_genType type
type MoqIsSorted_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqIsSorted_genType_mock

	ResultsByParams []MoqIsSorted_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Data moq.ParamIndexing
		}
	}
}

// MoqIsSorted_genType_mock isolates the mock interface of the IsSorted_genType
// type
type MoqIsSorted_genType_mock struct {
	Moq *MoqIsSorted_genType
}

// MoqIsSorted_genType_params holds the params of the IsSorted_genType type
type MoqIsSorted_genType_params struct{ Data sort.Interface }

// MoqIsSorted_genType_paramsKey holds the map key params of the
// IsSorted_genType type
type MoqIsSorted_genType_paramsKey struct {
	Params struct{ Data sort.Interface }
	Hashes struct{ Data hash.Hash }
}

// MoqIsSorted_genType_resultsByParams contains the results for a given set of
// parameters for the IsSorted_genType type
type MoqIsSorted_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqIsSorted_genType_paramsKey]*MoqIsSorted_genType_results
}

// MoqIsSorted_genType_doFn defines the type of function needed when calling
// AndDo for the IsSorted_genType type
type MoqIsSorted_genType_doFn func(data sort.Interface)

// MoqIsSorted_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the IsSorted_genType type
type MoqIsSorted_genType_doReturnFn func(data sort.Interface) bool

// MoqIsSorted_genType_results holds the results of the IsSorted_genType type
type MoqIsSorted_genType_results struct {
	Params  MoqIsSorted_genType_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqIsSorted_genType_doFn
		DoReturnFn MoqIsSorted_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqIsSorted_genType_fnRecorder routes recorded function calls to the
// MoqIsSorted_genType moq
type MoqIsSorted_genType_fnRecorder struct {
	Params    MoqIsSorted_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqIsSorted_genType_results
	Moq       *MoqIsSorted_genType
}

// MoqIsSorted_genType_anyParams isolates the any params functions of the
// IsSorted_genType type
type MoqIsSorted_genType_anyParams struct {
	Recorder *MoqIsSorted_genType_fnRecorder
}

// NewMoqIsSorted_genType creates a new moq of the IsSorted_genType type
func NewMoqIsSorted_genType(scene *moq.Scene, config *moq.Config) *MoqIsSorted_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqIsSorted_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqIsSorted_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Data moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Data moq.ParamIndexing
		}{
			Data: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the IsSorted_genType type
func (m *MoqIsSorted_genType) Mock() IsSorted_genType {
	return func(data sort.Interface) bool {
		m.Scene.T.Helper()
		moq := &MoqIsSorted_genType_mock{Moq: m}
		return moq.Fn(data)
	}
}

func (m *MoqIsSorted_genType_mock) Fn(data sort.Interface) (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqIsSorted_genType_params{
		Data: data,
	}
	var results *MoqIsSorted_genType_results
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
		result.DoFn(data)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(data)
	}
	return
}

func (m *MoqIsSorted_genType) OnCall(data sort.Interface) *MoqIsSorted_genType_fnRecorder {
	return &MoqIsSorted_genType_fnRecorder{
		Params: MoqIsSorted_genType_params{
			Data: data,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqIsSorted_genType_fnRecorder) Any() *MoqIsSorted_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqIsSorted_genType_anyParams{Recorder: r}
}

func (a *MoqIsSorted_genType_anyParams) Data() *MoqIsSorted_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqIsSorted_genType_fnRecorder) Seq() *MoqIsSorted_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqIsSorted_genType_fnRecorder) NoSeq() *MoqIsSorted_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqIsSorted_genType_fnRecorder) ReturnResults(result1 bool) *MoqIsSorted_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqIsSorted_genType_doFn
		DoReturnFn MoqIsSorted_genType_doReturnFn
	}{
		Values: &struct {
			Result1 bool
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqIsSorted_genType_fnRecorder) AndDo(fn MoqIsSorted_genType_doFn) *MoqIsSorted_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqIsSorted_genType_fnRecorder) DoReturnResults(fn MoqIsSorted_genType_doReturnFn) *MoqIsSorted_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqIsSorted_genType_doFn
		DoReturnFn MoqIsSorted_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqIsSorted_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqIsSorted_genType_resultsByParams
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
		results = &MoqIsSorted_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqIsSorted_genType_paramsKey]*MoqIsSorted_genType_results{},
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
		r.Results = &MoqIsSorted_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqIsSorted_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqIsSorted_genType_fnRecorder {
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
					Result1 bool
				}
				Sequence   uint32
				DoFn       MoqIsSorted_genType_doFn
				DoReturnFn MoqIsSorted_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqIsSorted_genType) PrettyParams(params MoqIsSorted_genType_params) string {
	return fmt.Sprintf("IsSorted_genType(%#v)", params.Data)
}

func (m *MoqIsSorted_genType) ParamsKey(params MoqIsSorted_genType_params, anyParams uint64) MoqIsSorted_genType_paramsKey {
	m.Scene.T.Helper()
	var dataUsed sort.Interface
	var dataUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Data == moq.ParamIndexByValue {
			dataUsed = params.Data
		} else {
			dataUsedHash = hash.DeepHash(params.Data)
		}
	}
	return MoqIsSorted_genType_paramsKey{
		Params: struct{ Data sort.Interface }{
			Data: dataUsed,
		},
		Hashes: struct{ Data hash.Hash }{
			Data: dataUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqIsSorted_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqIsSorted_genType) AssertExpectationsMet() {
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