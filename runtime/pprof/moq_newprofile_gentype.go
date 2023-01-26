// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package pprof

import (
	"fmt"
	"math/bits"
	"runtime/pprof"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// NewProfile_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type NewProfile_genType func(name string) *pprof.Profile

// MoqNewProfile_genType holds the state of a moq of the NewProfile_genType
// type
type MoqNewProfile_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNewProfile_genType_mock

	ResultsByParams []MoqNewProfile_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Name moq.ParamIndexing
		}
	}
}

// MoqNewProfile_genType_mock isolates the mock interface of the
// NewProfile_genType type
type MoqNewProfile_genType_mock struct {
	Moq *MoqNewProfile_genType
}

// MoqNewProfile_genType_params holds the params of the NewProfile_genType type
type MoqNewProfile_genType_params struct{ Name string }

// MoqNewProfile_genType_paramsKey holds the map key params of the
// NewProfile_genType type
type MoqNewProfile_genType_paramsKey struct {
	Params struct{ Name string }
	Hashes struct{ Name hash.Hash }
}

// MoqNewProfile_genType_resultsByParams contains the results for a given set
// of parameters for the NewProfile_genType type
type MoqNewProfile_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNewProfile_genType_paramsKey]*MoqNewProfile_genType_results
}

// MoqNewProfile_genType_doFn defines the type of function needed when calling
// AndDo for the NewProfile_genType type
type MoqNewProfile_genType_doFn func(name string)

// MoqNewProfile_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the NewProfile_genType type
type MoqNewProfile_genType_doReturnFn func(name string) *pprof.Profile

// MoqNewProfile_genType_results holds the results of the NewProfile_genType
// type
type MoqNewProfile_genType_results struct {
	Params  MoqNewProfile_genType_params
	Results []struct {
		Values *struct {
			Result1 *pprof.Profile
		}
		Sequence   uint32
		DoFn       MoqNewProfile_genType_doFn
		DoReturnFn MoqNewProfile_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNewProfile_genType_fnRecorder routes recorded function calls to the
// MoqNewProfile_genType moq
type MoqNewProfile_genType_fnRecorder struct {
	Params    MoqNewProfile_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNewProfile_genType_results
	Moq       *MoqNewProfile_genType
}

// MoqNewProfile_genType_anyParams isolates the any params functions of the
// NewProfile_genType type
type MoqNewProfile_genType_anyParams struct {
	Recorder *MoqNewProfile_genType_fnRecorder
}

// NewMoqNewProfile_genType creates a new moq of the NewProfile_genType type
func NewMoqNewProfile_genType(scene *moq.Scene, config *moq.Config) *MoqNewProfile_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNewProfile_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNewProfile_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Name moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Name moq.ParamIndexing
		}{
			Name: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the NewProfile_genType type
func (m *MoqNewProfile_genType) Mock() NewProfile_genType {
	return func(name string) *pprof.Profile {
		m.Scene.T.Helper()
		moq := &MoqNewProfile_genType_mock{Moq: m}
		return moq.Fn(name)
	}
}

func (m *MoqNewProfile_genType_mock) Fn(name string) (result1 *pprof.Profile) {
	m.Moq.Scene.T.Helper()
	params := MoqNewProfile_genType_params{
		Name: name,
	}
	var results *MoqNewProfile_genType_results
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
		result.DoFn(name)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(name)
	}
	return
}

func (m *MoqNewProfile_genType) OnCall(name string) *MoqNewProfile_genType_fnRecorder {
	return &MoqNewProfile_genType_fnRecorder{
		Params: MoqNewProfile_genType_params{
			Name: name,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNewProfile_genType_fnRecorder) Any() *MoqNewProfile_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNewProfile_genType_anyParams{Recorder: r}
}

func (a *MoqNewProfile_genType_anyParams) Name() *MoqNewProfile_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqNewProfile_genType_fnRecorder) Seq() *MoqNewProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNewProfile_genType_fnRecorder) NoSeq() *MoqNewProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNewProfile_genType_fnRecorder) ReturnResults(result1 *pprof.Profile) *MoqNewProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *pprof.Profile
		}
		Sequence   uint32
		DoFn       MoqNewProfile_genType_doFn
		DoReturnFn MoqNewProfile_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *pprof.Profile
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNewProfile_genType_fnRecorder) AndDo(fn MoqNewProfile_genType_doFn) *MoqNewProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNewProfile_genType_fnRecorder) DoReturnResults(fn MoqNewProfile_genType_doReturnFn) *MoqNewProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *pprof.Profile
		}
		Sequence   uint32
		DoFn       MoqNewProfile_genType_doFn
		DoReturnFn MoqNewProfile_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNewProfile_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNewProfile_genType_resultsByParams
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
		results = &MoqNewProfile_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNewProfile_genType_paramsKey]*MoqNewProfile_genType_results{},
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
		r.Results = &MoqNewProfile_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNewProfile_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNewProfile_genType_fnRecorder {
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
					Result1 *pprof.Profile
				}
				Sequence   uint32
				DoFn       MoqNewProfile_genType_doFn
				DoReturnFn MoqNewProfile_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNewProfile_genType) PrettyParams(params MoqNewProfile_genType_params) string {
	return fmt.Sprintf("NewProfile_genType(%#v)", params.Name)
}

func (m *MoqNewProfile_genType) ParamsKey(params MoqNewProfile_genType_params, anyParams uint64) MoqNewProfile_genType_paramsKey {
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
	return MoqNewProfile_genType_paramsKey{
		Params: struct{ Name string }{
			Name: nameUsed,
		},
		Hashes: struct{ Name hash.Hash }{
			Name: nameUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNewProfile_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNewProfile_genType) AssertExpectationsMet() {
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