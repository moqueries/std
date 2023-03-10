// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package flag

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Set_genType is the fabricated implementation type of this mock (emitted when
// mocking functions directly and not from a function type)
type Set_genType func(name, value string) error

// MoqSet_genType holds the state of a moq of the Set_genType type
type MoqSet_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSet_genType_mock

	ResultsByParams []MoqSet_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Name  moq.ParamIndexing
			Value moq.ParamIndexing
		}
	}
}

// MoqSet_genType_mock isolates the mock interface of the Set_genType type
type MoqSet_genType_mock struct {
	Moq *MoqSet_genType
}

// MoqSet_genType_params holds the params of the Set_genType type
type MoqSet_genType_params struct{ Name, Value string }

// MoqSet_genType_paramsKey holds the map key params of the Set_genType type
type MoqSet_genType_paramsKey struct {
	Params struct{ Name, Value string }
	Hashes struct{ Name, Value hash.Hash }
}

// MoqSet_genType_resultsByParams contains the results for a given set of
// parameters for the Set_genType type
type MoqSet_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSet_genType_paramsKey]*MoqSet_genType_results
}

// MoqSet_genType_doFn defines the type of function needed when calling AndDo
// for the Set_genType type
type MoqSet_genType_doFn func(name, value string)

// MoqSet_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Set_genType type
type MoqSet_genType_doReturnFn func(name, value string) error

// MoqSet_genType_results holds the results of the Set_genType type
type MoqSet_genType_results struct {
	Params  MoqSet_genType_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqSet_genType_doFn
		DoReturnFn MoqSet_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSet_genType_fnRecorder routes recorded function calls to the
// MoqSet_genType moq
type MoqSet_genType_fnRecorder struct {
	Params    MoqSet_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSet_genType_results
	Moq       *MoqSet_genType
}

// MoqSet_genType_anyParams isolates the any params functions of the
// Set_genType type
type MoqSet_genType_anyParams struct {
	Recorder *MoqSet_genType_fnRecorder
}

// NewMoqSet_genType creates a new moq of the Set_genType type
func NewMoqSet_genType(scene *moq.Scene, config *moq.Config) *MoqSet_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSet_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSet_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Name  moq.ParamIndexing
				Value moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Name  moq.ParamIndexing
			Value moq.ParamIndexing
		}{
			Name:  moq.ParamIndexByValue,
			Value: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Set_genType type
func (m *MoqSet_genType) Mock() Set_genType {
	return func(name, value string) error {
		m.Scene.T.Helper()
		moq := &MoqSet_genType_mock{Moq: m}
		return moq.Fn(name, value)
	}
}

func (m *MoqSet_genType_mock) Fn(name, value string) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqSet_genType_params{
		Name:  name,
		Value: value,
	}
	var results *MoqSet_genType_results
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
		result.DoFn(name, value)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(name, value)
	}
	return
}

func (m *MoqSet_genType) OnCall(name, value string) *MoqSet_genType_fnRecorder {
	return &MoqSet_genType_fnRecorder{
		Params: MoqSet_genType_params{
			Name:  name,
			Value: value,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSet_genType_fnRecorder) Any() *MoqSet_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSet_genType_anyParams{Recorder: r}
}

func (a *MoqSet_genType_anyParams) Name() *MoqSet_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqSet_genType_anyParams) Value() *MoqSet_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqSet_genType_fnRecorder) Seq() *MoqSet_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSet_genType_fnRecorder) NoSeq() *MoqSet_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSet_genType_fnRecorder) ReturnResults(result1 error) *MoqSet_genType_fnRecorder {
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
		DoFn       MoqSet_genType_doFn
		DoReturnFn MoqSet_genType_doReturnFn
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

func (r *MoqSet_genType_fnRecorder) AndDo(fn MoqSet_genType_doFn) *MoqSet_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSet_genType_fnRecorder) DoReturnResults(fn MoqSet_genType_doReturnFn) *MoqSet_genType_fnRecorder {
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
		DoFn       MoqSet_genType_doFn
		DoReturnFn MoqSet_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSet_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSet_genType_resultsByParams
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
		results = &MoqSet_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSet_genType_paramsKey]*MoqSet_genType_results{},
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
		r.Results = &MoqSet_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSet_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSet_genType_fnRecorder {
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
				DoFn       MoqSet_genType_doFn
				DoReturnFn MoqSet_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSet_genType) PrettyParams(params MoqSet_genType_params) string {
	return fmt.Sprintf("Set_genType(%#v, %#v)", params.Name, params.Value)
}

func (m *MoqSet_genType) ParamsKey(params MoqSet_genType_params, anyParams uint64) MoqSet_genType_paramsKey {
	m.Scene.T.Helper()
	var nameUsed string
	var nameUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Name == moq.ParamIndexByValue {
			nameUsed = params.Name
		} else {
			nameUsedHash = hash.DeepHash(params.Name)
		}
	}
	var valueUsed string
	var valueUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Value == moq.ParamIndexByValue {
			valueUsed = params.Value
		} else {
			valueUsedHash = hash.DeepHash(params.Value)
		}
	}
	return MoqSet_genType_paramsKey{
		Params: struct{ Name, Value string }{
			Name:  nameUsed,
			Value: valueUsed,
		},
		Hashes: struct{ Name, Value hash.Hash }{
			Name:  nameUsedHash,
			Value: valueUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSet_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSet_genType) AssertExpectationsMet() {
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
