// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package user

import (
	"fmt"
	"math/bits"
	"os/user"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// LookupGroup_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type LookupGroup_genType func(name string) (*user.Group, error)

// MoqLookupGroup_genType holds the state of a moq of the LookupGroup_genType
// type
type MoqLookupGroup_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqLookupGroup_genType_mock

	ResultsByParams []MoqLookupGroup_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Name moq.ParamIndexing
		}
	}
}

// MoqLookupGroup_genType_mock isolates the mock interface of the
// LookupGroup_genType type
type MoqLookupGroup_genType_mock struct {
	Moq *MoqLookupGroup_genType
}

// MoqLookupGroup_genType_params holds the params of the LookupGroup_genType
// type
type MoqLookupGroup_genType_params struct{ Name string }

// MoqLookupGroup_genType_paramsKey holds the map key params of the
// LookupGroup_genType type
type MoqLookupGroup_genType_paramsKey struct {
	Params struct{ Name string }
	Hashes struct{ Name hash.Hash }
}

// MoqLookupGroup_genType_resultsByParams contains the results for a given set
// of parameters for the LookupGroup_genType type
type MoqLookupGroup_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqLookupGroup_genType_paramsKey]*MoqLookupGroup_genType_results
}

// MoqLookupGroup_genType_doFn defines the type of function needed when calling
// AndDo for the LookupGroup_genType type
type MoqLookupGroup_genType_doFn func(name string)

// MoqLookupGroup_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the LookupGroup_genType type
type MoqLookupGroup_genType_doReturnFn func(name string) (*user.Group, error)

// MoqLookupGroup_genType_results holds the results of the LookupGroup_genType
// type
type MoqLookupGroup_genType_results struct {
	Params  MoqLookupGroup_genType_params
	Results []struct {
		Values *struct {
			Result1 *user.Group
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqLookupGroup_genType_doFn
		DoReturnFn MoqLookupGroup_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqLookupGroup_genType_fnRecorder routes recorded function calls to the
// MoqLookupGroup_genType moq
type MoqLookupGroup_genType_fnRecorder struct {
	Params    MoqLookupGroup_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqLookupGroup_genType_results
	Moq       *MoqLookupGroup_genType
}

// MoqLookupGroup_genType_anyParams isolates the any params functions of the
// LookupGroup_genType type
type MoqLookupGroup_genType_anyParams struct {
	Recorder *MoqLookupGroup_genType_fnRecorder
}

// NewMoqLookupGroup_genType creates a new moq of the LookupGroup_genType type
func NewMoqLookupGroup_genType(scene *moq.Scene, config *moq.Config) *MoqLookupGroup_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqLookupGroup_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqLookupGroup_genType_mock{},

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

// Mock returns the moq implementation of the LookupGroup_genType type
func (m *MoqLookupGroup_genType) Mock() LookupGroup_genType {
	return func(name string) (*user.Group, error) {
		m.Scene.T.Helper()
		moq := &MoqLookupGroup_genType_mock{Moq: m}
		return moq.Fn(name)
	}
}

func (m *MoqLookupGroup_genType_mock) Fn(name string) (result1 *user.Group, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqLookupGroup_genType_params{
		Name: name,
	}
	var results *MoqLookupGroup_genType_results
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
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(name)
	}
	return
}

func (m *MoqLookupGroup_genType) OnCall(name string) *MoqLookupGroup_genType_fnRecorder {
	return &MoqLookupGroup_genType_fnRecorder{
		Params: MoqLookupGroup_genType_params{
			Name: name,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqLookupGroup_genType_fnRecorder) Any() *MoqLookupGroup_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqLookupGroup_genType_anyParams{Recorder: r}
}

func (a *MoqLookupGroup_genType_anyParams) Name() *MoqLookupGroup_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqLookupGroup_genType_fnRecorder) Seq() *MoqLookupGroup_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqLookupGroup_genType_fnRecorder) NoSeq() *MoqLookupGroup_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqLookupGroup_genType_fnRecorder) ReturnResults(result1 *user.Group, result2 error) *MoqLookupGroup_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *user.Group
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqLookupGroup_genType_doFn
		DoReturnFn MoqLookupGroup_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *user.Group
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqLookupGroup_genType_fnRecorder) AndDo(fn MoqLookupGroup_genType_doFn) *MoqLookupGroup_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqLookupGroup_genType_fnRecorder) DoReturnResults(fn MoqLookupGroup_genType_doReturnFn) *MoqLookupGroup_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *user.Group
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqLookupGroup_genType_doFn
		DoReturnFn MoqLookupGroup_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqLookupGroup_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqLookupGroup_genType_resultsByParams
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
		results = &MoqLookupGroup_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqLookupGroup_genType_paramsKey]*MoqLookupGroup_genType_results{},
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
		r.Results = &MoqLookupGroup_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqLookupGroup_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqLookupGroup_genType_fnRecorder {
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
					Result1 *user.Group
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqLookupGroup_genType_doFn
				DoReturnFn MoqLookupGroup_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqLookupGroup_genType) PrettyParams(params MoqLookupGroup_genType_params) string {
	return fmt.Sprintf("LookupGroup_genType(%#v)", params.Name)
}

func (m *MoqLookupGroup_genType) ParamsKey(params MoqLookupGroup_genType_params, anyParams uint64) MoqLookupGroup_genType_paramsKey {
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
	return MoqLookupGroup_genType_paramsKey{
		Params: struct{ Name string }{
			Name: nameUsed,
		},
		Hashes: struct{ Name hash.Hash }{
			Name: nameUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqLookupGroup_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqLookupGroup_genType) AssertExpectationsMet() {
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