// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package reflect

import (
	"fmt"
	"math/bits"
	"reflect"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Select_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Select_genType func(cases []reflect.SelectCase) (chosen int, recv reflect.Value, recvOK bool)

// MoqSelect_genType holds the state of a moq of the Select_genType type
type MoqSelect_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSelect_genType_mock

	ResultsByParams []MoqSelect_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Cases moq.ParamIndexing
		}
	}
}

// MoqSelect_genType_mock isolates the mock interface of the Select_genType
// type
type MoqSelect_genType_mock struct {
	Moq *MoqSelect_genType
}

// MoqSelect_genType_params holds the params of the Select_genType type
type MoqSelect_genType_params struct{ Cases []reflect.SelectCase }

// MoqSelect_genType_paramsKey holds the map key params of the Select_genType
// type
type MoqSelect_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ Cases hash.Hash }
}

// MoqSelect_genType_resultsByParams contains the results for a given set of
// parameters for the Select_genType type
type MoqSelect_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSelect_genType_paramsKey]*MoqSelect_genType_results
}

// MoqSelect_genType_doFn defines the type of function needed when calling
// AndDo for the Select_genType type
type MoqSelect_genType_doFn func(cases []reflect.SelectCase)

// MoqSelect_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Select_genType type
type MoqSelect_genType_doReturnFn func(cases []reflect.SelectCase) (chosen int, recv reflect.Value, recvOK bool)

// MoqSelect_genType_results holds the results of the Select_genType type
type MoqSelect_genType_results struct {
	Params  MoqSelect_genType_params
	Results []struct {
		Values *struct {
			Chosen int
			Recv   reflect.Value
			RecvOK bool
		}
		Sequence   uint32
		DoFn       MoqSelect_genType_doFn
		DoReturnFn MoqSelect_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSelect_genType_fnRecorder routes recorded function calls to the
// MoqSelect_genType moq
type MoqSelect_genType_fnRecorder struct {
	Params    MoqSelect_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSelect_genType_results
	Moq       *MoqSelect_genType
}

// MoqSelect_genType_anyParams isolates the any params functions of the
// Select_genType type
type MoqSelect_genType_anyParams struct {
	Recorder *MoqSelect_genType_fnRecorder
}

// NewMoqSelect_genType creates a new moq of the Select_genType type
func NewMoqSelect_genType(scene *moq.Scene, config *moq.Config) *MoqSelect_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSelect_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSelect_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Cases moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Cases moq.ParamIndexing
		}{
			Cases: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Select_genType type
func (m *MoqSelect_genType) Mock() Select_genType {
	return func(cases []reflect.SelectCase) (_ int, _ reflect.Value, _ bool) {
		m.Scene.T.Helper()
		moq := &MoqSelect_genType_mock{Moq: m}
		return moq.Fn(cases)
	}
}

func (m *MoqSelect_genType_mock) Fn(cases []reflect.SelectCase) (chosen int, recv reflect.Value, recvOK bool) {
	m.Moq.Scene.T.Helper()
	params := MoqSelect_genType_params{
		Cases: cases,
	}
	var results *MoqSelect_genType_results
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
		result.DoFn(cases)
	}

	if result.Values != nil {
		chosen = result.Values.Chosen
		recv = result.Values.Recv
		recvOK = result.Values.RecvOK
	}
	if result.DoReturnFn != nil {
		chosen, recv, recvOK = result.DoReturnFn(cases)
	}
	return
}

func (m *MoqSelect_genType) OnCall(cases []reflect.SelectCase) *MoqSelect_genType_fnRecorder {
	return &MoqSelect_genType_fnRecorder{
		Params: MoqSelect_genType_params{
			Cases: cases,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSelect_genType_fnRecorder) Any() *MoqSelect_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSelect_genType_anyParams{Recorder: r}
}

func (a *MoqSelect_genType_anyParams) Cases() *MoqSelect_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqSelect_genType_fnRecorder) Seq() *MoqSelect_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSelect_genType_fnRecorder) NoSeq() *MoqSelect_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSelect_genType_fnRecorder) ReturnResults(chosen int, recv reflect.Value, recvOK bool) *MoqSelect_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Chosen int
			Recv   reflect.Value
			RecvOK bool
		}
		Sequence   uint32
		DoFn       MoqSelect_genType_doFn
		DoReturnFn MoqSelect_genType_doReturnFn
	}{
		Values: &struct {
			Chosen int
			Recv   reflect.Value
			RecvOK bool
		}{
			Chosen: chosen,
			Recv:   recv,
			RecvOK: recvOK,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSelect_genType_fnRecorder) AndDo(fn MoqSelect_genType_doFn) *MoqSelect_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSelect_genType_fnRecorder) DoReturnResults(fn MoqSelect_genType_doReturnFn) *MoqSelect_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Chosen int
			Recv   reflect.Value
			RecvOK bool
		}
		Sequence   uint32
		DoFn       MoqSelect_genType_doFn
		DoReturnFn MoqSelect_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSelect_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSelect_genType_resultsByParams
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
		results = &MoqSelect_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSelect_genType_paramsKey]*MoqSelect_genType_results{},
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
		r.Results = &MoqSelect_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSelect_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSelect_genType_fnRecorder {
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
					Chosen int
					Recv   reflect.Value
					RecvOK bool
				}
				Sequence   uint32
				DoFn       MoqSelect_genType_doFn
				DoReturnFn MoqSelect_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSelect_genType) PrettyParams(params MoqSelect_genType_params) string {
	return fmt.Sprintf("Select_genType(%#v)", params.Cases)
}

func (m *MoqSelect_genType) ParamsKey(params MoqSelect_genType_params, anyParams uint64) MoqSelect_genType_paramsKey {
	m.Scene.T.Helper()
	var casesUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Cases == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The cases parameter can't be indexed by value")
		}
		casesUsedHash = hash.DeepHash(params.Cases)
	}
	return MoqSelect_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ Cases hash.Hash }{
			Cases: casesUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSelect_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSelect_genType) AssertExpectationsMet() {
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
