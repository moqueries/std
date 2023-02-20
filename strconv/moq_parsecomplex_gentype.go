// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package strconv

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// ParseComplex_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type ParseComplex_genType func(s string, bitSize int) (complex128, error)

// MoqParseComplex_genType holds the state of a moq of the ParseComplex_genType
// type
type MoqParseComplex_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqParseComplex_genType_mock

	ResultsByParams []MoqParseComplex_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			S       moq.ParamIndexing
			BitSize moq.ParamIndexing
		}
	}
}

// MoqParseComplex_genType_mock isolates the mock interface of the
// ParseComplex_genType type
type MoqParseComplex_genType_mock struct {
	Moq *MoqParseComplex_genType
}

// MoqParseComplex_genType_params holds the params of the ParseComplex_genType
// type
type MoqParseComplex_genType_params struct {
	S       string
	BitSize int
}

// MoqParseComplex_genType_paramsKey holds the map key params of the
// ParseComplex_genType type
type MoqParseComplex_genType_paramsKey struct {
	Params struct {
		S       string
		BitSize int
	}
	Hashes struct {
		S       hash.Hash
		BitSize hash.Hash
	}
}

// MoqParseComplex_genType_resultsByParams contains the results for a given set
// of parameters for the ParseComplex_genType type
type MoqParseComplex_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqParseComplex_genType_paramsKey]*MoqParseComplex_genType_results
}

// MoqParseComplex_genType_doFn defines the type of function needed when
// calling AndDo for the ParseComplex_genType type
type MoqParseComplex_genType_doFn func(s string, bitSize int)

// MoqParseComplex_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ParseComplex_genType type
type MoqParseComplex_genType_doReturnFn func(s string, bitSize int) (complex128, error)

// MoqParseComplex_genType_results holds the results of the
// ParseComplex_genType type
type MoqParseComplex_genType_results struct {
	Params  MoqParseComplex_genType_params
	Results []struct {
		Values *struct {
			Result1 complex128
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqParseComplex_genType_doFn
		DoReturnFn MoqParseComplex_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqParseComplex_genType_fnRecorder routes recorded function calls to the
// MoqParseComplex_genType moq
type MoqParseComplex_genType_fnRecorder struct {
	Params    MoqParseComplex_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqParseComplex_genType_results
	Moq       *MoqParseComplex_genType
}

// MoqParseComplex_genType_anyParams isolates the any params functions of the
// ParseComplex_genType type
type MoqParseComplex_genType_anyParams struct {
	Recorder *MoqParseComplex_genType_fnRecorder
}

// NewMoqParseComplex_genType creates a new moq of the ParseComplex_genType
// type
func NewMoqParseComplex_genType(scene *moq.Scene, config *moq.Config) *MoqParseComplex_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqParseComplex_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqParseComplex_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				S       moq.ParamIndexing
				BitSize moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			S       moq.ParamIndexing
			BitSize moq.ParamIndexing
		}{
			S:       moq.ParamIndexByValue,
			BitSize: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ParseComplex_genType type
func (m *MoqParseComplex_genType) Mock() ParseComplex_genType {
	return func(s string, bitSize int) (complex128, error) {
		m.Scene.T.Helper()
		moq := &MoqParseComplex_genType_mock{Moq: m}
		return moq.Fn(s, bitSize)
	}
}

func (m *MoqParseComplex_genType_mock) Fn(s string, bitSize int) (result1 complex128, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqParseComplex_genType_params{
		S:       s,
		BitSize: bitSize,
	}
	var results *MoqParseComplex_genType_results
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
		result.DoFn(s, bitSize)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(s, bitSize)
	}
	return
}

func (m *MoqParseComplex_genType) OnCall(s string, bitSize int) *MoqParseComplex_genType_fnRecorder {
	return &MoqParseComplex_genType_fnRecorder{
		Params: MoqParseComplex_genType_params{
			S:       s,
			BitSize: bitSize,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqParseComplex_genType_fnRecorder) Any() *MoqParseComplex_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqParseComplex_genType_anyParams{Recorder: r}
}

func (a *MoqParseComplex_genType_anyParams) S() *MoqParseComplex_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqParseComplex_genType_anyParams) BitSize() *MoqParseComplex_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqParseComplex_genType_fnRecorder) Seq() *MoqParseComplex_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqParseComplex_genType_fnRecorder) NoSeq() *MoqParseComplex_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqParseComplex_genType_fnRecorder) ReturnResults(result1 complex128, result2 error) *MoqParseComplex_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 complex128
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqParseComplex_genType_doFn
		DoReturnFn MoqParseComplex_genType_doReturnFn
	}{
		Values: &struct {
			Result1 complex128
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqParseComplex_genType_fnRecorder) AndDo(fn MoqParseComplex_genType_doFn) *MoqParseComplex_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqParseComplex_genType_fnRecorder) DoReturnResults(fn MoqParseComplex_genType_doReturnFn) *MoqParseComplex_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 complex128
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqParseComplex_genType_doFn
		DoReturnFn MoqParseComplex_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqParseComplex_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqParseComplex_genType_resultsByParams
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
		results = &MoqParseComplex_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqParseComplex_genType_paramsKey]*MoqParseComplex_genType_results{},
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
		r.Results = &MoqParseComplex_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqParseComplex_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqParseComplex_genType_fnRecorder {
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
					Result1 complex128
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqParseComplex_genType_doFn
				DoReturnFn MoqParseComplex_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqParseComplex_genType) PrettyParams(params MoqParseComplex_genType_params) string {
	return fmt.Sprintf("ParseComplex_genType(%#v, %#v)", params.S, params.BitSize)
}

func (m *MoqParseComplex_genType) ParamsKey(params MoqParseComplex_genType_params, anyParams uint64) MoqParseComplex_genType_paramsKey {
	m.Scene.T.Helper()
	var sUsed string
	var sUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.S == moq.ParamIndexByValue {
			sUsed = params.S
		} else {
			sUsedHash = hash.DeepHash(params.S)
		}
	}
	var bitSizeUsed int
	var bitSizeUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.BitSize == moq.ParamIndexByValue {
			bitSizeUsed = params.BitSize
		} else {
			bitSizeUsedHash = hash.DeepHash(params.BitSize)
		}
	}
	return MoqParseComplex_genType_paramsKey{
		Params: struct {
			S       string
			BitSize int
		}{
			S:       sUsed,
			BitSize: bitSizeUsed,
		},
		Hashes: struct {
			S       hash.Hash
			BitSize hash.Hash
		}{
			S:       sUsedHash,
			BitSize: bitSizeUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqParseComplex_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqParseComplex_genType) AssertExpectationsMet() {
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
