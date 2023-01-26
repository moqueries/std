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

// Sort_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Sort_genType func(data sort.Interface)

// MoqSort_genType holds the state of a moq of the Sort_genType type
type MoqSort_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSort_genType_mock

	ResultsByParams []MoqSort_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Data moq.ParamIndexing
		}
	}
}

// MoqSort_genType_mock isolates the mock interface of the Sort_genType type
type MoqSort_genType_mock struct {
	Moq *MoqSort_genType
}

// MoqSort_genType_params holds the params of the Sort_genType type
type MoqSort_genType_params struct{ Data sort.Interface }

// MoqSort_genType_paramsKey holds the map key params of the Sort_genType type
type MoqSort_genType_paramsKey struct {
	Params struct{ Data sort.Interface }
	Hashes struct{ Data hash.Hash }
}

// MoqSort_genType_resultsByParams contains the results for a given set of
// parameters for the Sort_genType type
type MoqSort_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSort_genType_paramsKey]*MoqSort_genType_results
}

// MoqSort_genType_doFn defines the type of function needed when calling AndDo
// for the Sort_genType type
type MoqSort_genType_doFn func(data sort.Interface)

// MoqSort_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Sort_genType type
type MoqSort_genType_doReturnFn func(data sort.Interface)

// MoqSort_genType_results holds the results of the Sort_genType type
type MoqSort_genType_results struct {
	Params  MoqSort_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSort_genType_doFn
		DoReturnFn MoqSort_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSort_genType_fnRecorder routes recorded function calls to the
// MoqSort_genType moq
type MoqSort_genType_fnRecorder struct {
	Params    MoqSort_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSort_genType_results
	Moq       *MoqSort_genType
}

// MoqSort_genType_anyParams isolates the any params functions of the
// Sort_genType type
type MoqSort_genType_anyParams struct {
	Recorder *MoqSort_genType_fnRecorder
}

// NewMoqSort_genType creates a new moq of the Sort_genType type
func NewMoqSort_genType(scene *moq.Scene, config *moq.Config) *MoqSort_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSort_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSort_genType_mock{},

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

// Mock returns the moq implementation of the Sort_genType type
func (m *MoqSort_genType) Mock() Sort_genType {
	return func(data sort.Interface) { m.Scene.T.Helper(); moq := &MoqSort_genType_mock{Moq: m}; moq.Fn(data) }
}

func (m *MoqSort_genType_mock) Fn(data sort.Interface) {
	m.Moq.Scene.T.Helper()
	params := MoqSort_genType_params{
		Data: data,
	}
	var results *MoqSort_genType_results
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

	if result.DoReturnFn != nil {
		result.DoReturnFn(data)
	}
	return
}

func (m *MoqSort_genType) OnCall(data sort.Interface) *MoqSort_genType_fnRecorder {
	return &MoqSort_genType_fnRecorder{
		Params: MoqSort_genType_params{
			Data: data,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSort_genType_fnRecorder) Any() *MoqSort_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSort_genType_anyParams{Recorder: r}
}

func (a *MoqSort_genType_anyParams) Data() *MoqSort_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqSort_genType_fnRecorder) Seq() *MoqSort_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSort_genType_fnRecorder) NoSeq() *MoqSort_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSort_genType_fnRecorder) ReturnResults() *MoqSort_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSort_genType_doFn
		DoReturnFn MoqSort_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSort_genType_fnRecorder) AndDo(fn MoqSort_genType_doFn) *MoqSort_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSort_genType_fnRecorder) DoReturnResults(fn MoqSort_genType_doReturnFn) *MoqSort_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSort_genType_doFn
		DoReturnFn MoqSort_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSort_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSort_genType_resultsByParams
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
		results = &MoqSort_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSort_genType_paramsKey]*MoqSort_genType_results{},
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
		r.Results = &MoqSort_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSort_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSort_genType_fnRecorder {
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
				DoFn       MoqSort_genType_doFn
				DoReturnFn MoqSort_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSort_genType) PrettyParams(params MoqSort_genType_params) string {
	return fmt.Sprintf("Sort_genType(%#v)", params.Data)
}

func (m *MoqSort_genType) ParamsKey(params MoqSort_genType_params, anyParams uint64) MoqSort_genType_paramsKey {
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
	return MoqSort_genType_paramsKey{
		Params: struct{ Data sort.Interface }{
			Data: dataUsed,
		},
		Hashes: struct{ Data hash.Hash }{
			Data: dataUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSort_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSort_genType) AssertExpectationsMet() {
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
