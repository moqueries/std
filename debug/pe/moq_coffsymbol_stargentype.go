// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package pe

import (
	"debug/pe"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that pe.COFFSymbol_starGenType is
// mocked completely
var _ COFFSymbol_starGenType = (*MoqCOFFSymbol_starGenType_mock)(nil)

// COFFSymbol_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type COFFSymbol_starGenType interface {
	FullName(st pe.StringTable) (string, error)
}

// MoqCOFFSymbol_starGenType holds the state of a moq of the
// COFFSymbol_starGenType type
type MoqCOFFSymbol_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqCOFFSymbol_starGenType_mock

	ResultsByParams_FullName []MoqCOFFSymbol_starGenType_FullName_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			FullName struct {
				St moq.ParamIndexing
			}
		}
	}
	// MoqCOFFSymbol_starGenType_mock isolates the mock interface of the
}

// COFFSymbol_starGenType type
type MoqCOFFSymbol_starGenType_mock struct {
	Moq *MoqCOFFSymbol_starGenType
}

// MoqCOFFSymbol_starGenType_recorder isolates the recorder interface of the
// COFFSymbol_starGenType type
type MoqCOFFSymbol_starGenType_recorder struct {
	Moq *MoqCOFFSymbol_starGenType
}

// MoqCOFFSymbol_starGenType_FullName_params holds the params of the
// COFFSymbol_starGenType type
type MoqCOFFSymbol_starGenType_FullName_params struct{ St pe.StringTable }

// MoqCOFFSymbol_starGenType_FullName_paramsKey holds the map key params of the
// COFFSymbol_starGenType type
type MoqCOFFSymbol_starGenType_FullName_paramsKey struct {
	Params struct{}
	Hashes struct{ St hash.Hash }
}

// MoqCOFFSymbol_starGenType_FullName_resultsByParams contains the results for
// a given set of parameters for the COFFSymbol_starGenType type
type MoqCOFFSymbol_starGenType_FullName_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqCOFFSymbol_starGenType_FullName_paramsKey]*MoqCOFFSymbol_starGenType_FullName_results
}

// MoqCOFFSymbol_starGenType_FullName_doFn defines the type of function needed
// when calling AndDo for the COFFSymbol_starGenType type
type MoqCOFFSymbol_starGenType_FullName_doFn func(st pe.StringTable)

// MoqCOFFSymbol_starGenType_FullName_doReturnFn defines the type of function
// needed when calling DoReturnResults for the COFFSymbol_starGenType type
type MoqCOFFSymbol_starGenType_FullName_doReturnFn func(st pe.StringTable) (string, error)

// MoqCOFFSymbol_starGenType_FullName_results holds the results of the
// COFFSymbol_starGenType type
type MoqCOFFSymbol_starGenType_FullName_results struct {
	Params  MoqCOFFSymbol_starGenType_FullName_params
	Results []struct {
		Values *struct {
			Result1 string
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqCOFFSymbol_starGenType_FullName_doFn
		DoReturnFn MoqCOFFSymbol_starGenType_FullName_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqCOFFSymbol_starGenType_FullName_fnRecorder routes recorded function calls
// to the MoqCOFFSymbol_starGenType moq
type MoqCOFFSymbol_starGenType_FullName_fnRecorder struct {
	Params    MoqCOFFSymbol_starGenType_FullName_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqCOFFSymbol_starGenType_FullName_results
	Moq       *MoqCOFFSymbol_starGenType
}

// MoqCOFFSymbol_starGenType_FullName_anyParams isolates the any params
// functions of the COFFSymbol_starGenType type
type MoqCOFFSymbol_starGenType_FullName_anyParams struct {
	Recorder *MoqCOFFSymbol_starGenType_FullName_fnRecorder
}

// NewMoqCOFFSymbol_starGenType creates a new moq of the COFFSymbol_starGenType
// type
func NewMoqCOFFSymbol_starGenType(scene *moq.Scene, config *moq.Config) *MoqCOFFSymbol_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqCOFFSymbol_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqCOFFSymbol_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				FullName struct {
					St moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			FullName struct {
				St moq.ParamIndexing
			}
		}{
			FullName: struct {
				St moq.ParamIndexing
			}{
				St: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the COFFSymbol_starGenType type
func (m *MoqCOFFSymbol_starGenType) Mock() *MoqCOFFSymbol_starGenType_mock { return m.Moq }

func (m *MoqCOFFSymbol_starGenType_mock) FullName(st pe.StringTable) (result1 string, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqCOFFSymbol_starGenType_FullName_params{
		St: st,
	}
	var results *MoqCOFFSymbol_starGenType_FullName_results
	for _, resultsByParams := range m.Moq.ResultsByParams_FullName {
		paramsKey := m.Moq.ParamsKey_FullName(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_FullName(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_FullName(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_FullName(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(st)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(st)
	}
	return
}

// OnCall returns the recorder implementation of the COFFSymbol_starGenType
// type
func (m *MoqCOFFSymbol_starGenType) OnCall() *MoqCOFFSymbol_starGenType_recorder {
	return &MoqCOFFSymbol_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqCOFFSymbol_starGenType_recorder) FullName(st pe.StringTable) *MoqCOFFSymbol_starGenType_FullName_fnRecorder {
	return &MoqCOFFSymbol_starGenType_FullName_fnRecorder{
		Params: MoqCOFFSymbol_starGenType_FullName_params{
			St: st,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqCOFFSymbol_starGenType_FullName_fnRecorder) Any() *MoqCOFFSymbol_starGenType_FullName_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_FullName(r.Params))
		return nil
	}
	return &MoqCOFFSymbol_starGenType_FullName_anyParams{Recorder: r}
}

func (a *MoqCOFFSymbol_starGenType_FullName_anyParams) St() *MoqCOFFSymbol_starGenType_FullName_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqCOFFSymbol_starGenType_FullName_fnRecorder) Seq() *MoqCOFFSymbol_starGenType_FullName_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_FullName(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqCOFFSymbol_starGenType_FullName_fnRecorder) NoSeq() *MoqCOFFSymbol_starGenType_FullName_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_FullName(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqCOFFSymbol_starGenType_FullName_fnRecorder) ReturnResults(result1 string, result2 error) *MoqCOFFSymbol_starGenType_FullName_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqCOFFSymbol_starGenType_FullName_doFn
		DoReturnFn MoqCOFFSymbol_starGenType_FullName_doReturnFn
	}{
		Values: &struct {
			Result1 string
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqCOFFSymbol_starGenType_FullName_fnRecorder) AndDo(fn MoqCOFFSymbol_starGenType_FullName_doFn) *MoqCOFFSymbol_starGenType_FullName_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqCOFFSymbol_starGenType_FullName_fnRecorder) DoReturnResults(fn MoqCOFFSymbol_starGenType_FullName_doReturnFn) *MoqCOFFSymbol_starGenType_FullName_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqCOFFSymbol_starGenType_FullName_doFn
		DoReturnFn MoqCOFFSymbol_starGenType_FullName_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqCOFFSymbol_starGenType_FullName_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqCOFFSymbol_starGenType_FullName_resultsByParams
	for n, res := range r.Moq.ResultsByParams_FullName {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqCOFFSymbol_starGenType_FullName_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqCOFFSymbol_starGenType_FullName_paramsKey]*MoqCOFFSymbol_starGenType_FullName_results{},
		}
		r.Moq.ResultsByParams_FullName = append(r.Moq.ResultsByParams_FullName, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_FullName) {
			copy(r.Moq.ResultsByParams_FullName[insertAt+1:], r.Moq.ResultsByParams_FullName[insertAt:0])
			r.Moq.ResultsByParams_FullName[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_FullName(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqCOFFSymbol_starGenType_FullName_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqCOFFSymbol_starGenType_FullName_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqCOFFSymbol_starGenType_FullName_fnRecorder {
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
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqCOFFSymbol_starGenType_FullName_doFn
				DoReturnFn MoqCOFFSymbol_starGenType_FullName_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqCOFFSymbol_starGenType) PrettyParams_FullName(params MoqCOFFSymbol_starGenType_FullName_params) string {
	return fmt.Sprintf("FullName(%#v)", params.St)
}

func (m *MoqCOFFSymbol_starGenType) ParamsKey_FullName(params MoqCOFFSymbol_starGenType_FullName_params, anyParams uint64) MoqCOFFSymbol_starGenType_FullName_paramsKey {
	m.Scene.T.Helper()
	var stUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.FullName.St == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The st parameter of the FullName function can't be indexed by value")
		}
		stUsedHash = hash.DeepHash(params.St)
	}
	return MoqCOFFSymbol_starGenType_FullName_paramsKey{
		Params: struct{}{},
		Hashes: struct{ St hash.Hash }{
			St: stUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqCOFFSymbol_starGenType) Reset() { m.ResultsByParams_FullName = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqCOFFSymbol_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_FullName {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_FullName(results.Params))
			}
		}
	}
}
