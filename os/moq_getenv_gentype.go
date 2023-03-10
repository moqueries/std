// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package os

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Getenv_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Getenv_genType func(key string) string

// MoqGetenv_genType holds the state of a moq of the Getenv_genType type
type MoqGetenv_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqGetenv_genType_mock

	ResultsByParams []MoqGetenv_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Key moq.ParamIndexing
		}
	}
}

// MoqGetenv_genType_mock isolates the mock interface of the Getenv_genType
// type
type MoqGetenv_genType_mock struct {
	Moq *MoqGetenv_genType
}

// MoqGetenv_genType_params holds the params of the Getenv_genType type
type MoqGetenv_genType_params struct{ Key string }

// MoqGetenv_genType_paramsKey holds the map key params of the Getenv_genType
// type
type MoqGetenv_genType_paramsKey struct {
	Params struct{ Key string }
	Hashes struct{ Key hash.Hash }
}

// MoqGetenv_genType_resultsByParams contains the results for a given set of
// parameters for the Getenv_genType type
type MoqGetenv_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqGetenv_genType_paramsKey]*MoqGetenv_genType_results
}

// MoqGetenv_genType_doFn defines the type of function needed when calling
// AndDo for the Getenv_genType type
type MoqGetenv_genType_doFn func(key string)

// MoqGetenv_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Getenv_genType type
type MoqGetenv_genType_doReturnFn func(key string) string

// MoqGetenv_genType_results holds the results of the Getenv_genType type
type MoqGetenv_genType_results struct {
	Params  MoqGetenv_genType_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqGetenv_genType_doFn
		DoReturnFn MoqGetenv_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqGetenv_genType_fnRecorder routes recorded function calls to the
// MoqGetenv_genType moq
type MoqGetenv_genType_fnRecorder struct {
	Params    MoqGetenv_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqGetenv_genType_results
	Moq       *MoqGetenv_genType
}

// MoqGetenv_genType_anyParams isolates the any params functions of the
// Getenv_genType type
type MoqGetenv_genType_anyParams struct {
	Recorder *MoqGetenv_genType_fnRecorder
}

// NewMoqGetenv_genType creates a new moq of the Getenv_genType type
func NewMoqGetenv_genType(scene *moq.Scene, config *moq.Config) *MoqGetenv_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqGetenv_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqGetenv_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Key moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Key moq.ParamIndexing
		}{
			Key: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Getenv_genType type
func (m *MoqGetenv_genType) Mock() Getenv_genType {
	return func(key string) string {
		m.Scene.T.Helper()
		moq := &MoqGetenv_genType_mock{Moq: m}
		return moq.Fn(key)
	}
}

func (m *MoqGetenv_genType_mock) Fn(key string) (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqGetenv_genType_params{
		Key: key,
	}
	var results *MoqGetenv_genType_results
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
		result.DoFn(key)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(key)
	}
	return
}

func (m *MoqGetenv_genType) OnCall(key string) *MoqGetenv_genType_fnRecorder {
	return &MoqGetenv_genType_fnRecorder{
		Params: MoqGetenv_genType_params{
			Key: key,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqGetenv_genType_fnRecorder) Any() *MoqGetenv_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqGetenv_genType_anyParams{Recorder: r}
}

func (a *MoqGetenv_genType_anyParams) Key() *MoqGetenv_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqGetenv_genType_fnRecorder) Seq() *MoqGetenv_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqGetenv_genType_fnRecorder) NoSeq() *MoqGetenv_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqGetenv_genType_fnRecorder) ReturnResults(result1 string) *MoqGetenv_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqGetenv_genType_doFn
		DoReturnFn MoqGetenv_genType_doReturnFn
	}{
		Values: &struct {
			Result1 string
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqGetenv_genType_fnRecorder) AndDo(fn MoqGetenv_genType_doFn) *MoqGetenv_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqGetenv_genType_fnRecorder) DoReturnResults(fn MoqGetenv_genType_doReturnFn) *MoqGetenv_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqGetenv_genType_doFn
		DoReturnFn MoqGetenv_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqGetenv_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqGetenv_genType_resultsByParams
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
		results = &MoqGetenv_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqGetenv_genType_paramsKey]*MoqGetenv_genType_results{},
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
		r.Results = &MoqGetenv_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqGetenv_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqGetenv_genType_fnRecorder {
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
					Result1 string
				}
				Sequence   uint32
				DoFn       MoqGetenv_genType_doFn
				DoReturnFn MoqGetenv_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqGetenv_genType) PrettyParams(params MoqGetenv_genType_params) string {
	return fmt.Sprintf("Getenv_genType(%#v)", params.Key)
}

func (m *MoqGetenv_genType) ParamsKey(params MoqGetenv_genType_params, anyParams uint64) MoqGetenv_genType_paramsKey {
	m.Scene.T.Helper()
	var keyUsed string
	var keyUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Key == moq.ParamIndexByValue {
			keyUsed = params.Key
		} else {
			keyUsedHash = hash.DeepHash(params.Key)
		}
	}
	return MoqGetenv_genType_paramsKey{
		Params: struct{ Key string }{
			Key: keyUsed,
		},
		Hashes: struct{ Key hash.Hash }{
			Key: keyUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqGetenv_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqGetenv_genType) AssertExpectationsMet() {
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
