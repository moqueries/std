// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package macho

import (
	"debug/macho"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Open_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Open_genType func(name string) (*macho.File, error)

// MoqOpen_genType holds the state of a moq of the Open_genType type
type MoqOpen_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqOpen_genType_mock

	ResultsByParams []MoqOpen_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Name moq.ParamIndexing
		}
	}
}

// MoqOpen_genType_mock isolates the mock interface of the Open_genType type
type MoqOpen_genType_mock struct {
	Moq *MoqOpen_genType
}

// MoqOpen_genType_params holds the params of the Open_genType type
type MoqOpen_genType_params struct{ Name string }

// MoqOpen_genType_paramsKey holds the map key params of the Open_genType type
type MoqOpen_genType_paramsKey struct {
	Params struct{ Name string }
	Hashes struct{ Name hash.Hash }
}

// MoqOpen_genType_resultsByParams contains the results for a given set of
// parameters for the Open_genType type
type MoqOpen_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqOpen_genType_paramsKey]*MoqOpen_genType_results
}

// MoqOpen_genType_doFn defines the type of function needed when calling AndDo
// for the Open_genType type
type MoqOpen_genType_doFn func(name string)

// MoqOpen_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Open_genType type
type MoqOpen_genType_doReturnFn func(name string) (*macho.File, error)

// MoqOpen_genType_results holds the results of the Open_genType type
type MoqOpen_genType_results struct {
	Params  MoqOpen_genType_params
	Results []struct {
		Values *struct {
			Result1 *macho.File
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqOpen_genType_doFn
		DoReturnFn MoqOpen_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqOpen_genType_fnRecorder routes recorded function calls to the
// MoqOpen_genType moq
type MoqOpen_genType_fnRecorder struct {
	Params    MoqOpen_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqOpen_genType_results
	Moq       *MoqOpen_genType
}

// MoqOpen_genType_anyParams isolates the any params functions of the
// Open_genType type
type MoqOpen_genType_anyParams struct {
	Recorder *MoqOpen_genType_fnRecorder
}

// NewMoqOpen_genType creates a new moq of the Open_genType type
func NewMoqOpen_genType(scene *moq.Scene, config *moq.Config) *MoqOpen_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqOpen_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqOpen_genType_mock{},

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

// Mock returns the moq implementation of the Open_genType type
func (m *MoqOpen_genType) Mock() Open_genType {
	return func(name string) (*macho.File, error) {
		m.Scene.T.Helper()
		moq := &MoqOpen_genType_mock{Moq: m}
		return moq.Fn(name)
	}
}

func (m *MoqOpen_genType_mock) Fn(name string) (result1 *macho.File, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqOpen_genType_params{
		Name: name,
	}
	var results *MoqOpen_genType_results
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

func (m *MoqOpen_genType) OnCall(name string) *MoqOpen_genType_fnRecorder {
	return &MoqOpen_genType_fnRecorder{
		Params: MoqOpen_genType_params{
			Name: name,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqOpen_genType_fnRecorder) Any() *MoqOpen_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqOpen_genType_anyParams{Recorder: r}
}

func (a *MoqOpen_genType_anyParams) Name() *MoqOpen_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqOpen_genType_fnRecorder) Seq() *MoqOpen_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqOpen_genType_fnRecorder) NoSeq() *MoqOpen_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqOpen_genType_fnRecorder) ReturnResults(result1 *macho.File, result2 error) *MoqOpen_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *macho.File
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqOpen_genType_doFn
		DoReturnFn MoqOpen_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *macho.File
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqOpen_genType_fnRecorder) AndDo(fn MoqOpen_genType_doFn) *MoqOpen_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqOpen_genType_fnRecorder) DoReturnResults(fn MoqOpen_genType_doReturnFn) *MoqOpen_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *macho.File
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqOpen_genType_doFn
		DoReturnFn MoqOpen_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqOpen_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqOpen_genType_resultsByParams
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
		results = &MoqOpen_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqOpen_genType_paramsKey]*MoqOpen_genType_results{},
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
		r.Results = &MoqOpen_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqOpen_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqOpen_genType_fnRecorder {
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
					Result1 *macho.File
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqOpen_genType_doFn
				DoReturnFn MoqOpen_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqOpen_genType) PrettyParams(params MoqOpen_genType_params) string {
	return fmt.Sprintf("Open_genType(%#v)", params.Name)
}

func (m *MoqOpen_genType) ParamsKey(params MoqOpen_genType_params, anyParams uint64) MoqOpen_genType_paramsKey {
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
	return MoqOpen_genType_paramsKey{
		Params: struct{ Name string }{
			Name: nameUsed,
		},
		Hashes: struct{ Name hash.Hash }{
			Name: nameUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqOpen_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqOpen_genType) AssertExpectationsMet() {
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
