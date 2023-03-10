// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package quick

import (
	"fmt"
	"math/bits"
	"math/rand"
	"reflect"
	"sync/atomic"
	"testing/quick"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that quick.Generator is mocked
// completely
var _ quick.Generator = (*MoqGenerator_mock)(nil)

// MoqGenerator holds the state of a moq of the Generator type
type MoqGenerator struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqGenerator_mock

	ResultsByParams_Generate []MoqGenerator_Generate_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Generate struct {
				Rand moq.ParamIndexing
				Size moq.ParamIndexing
			}
		}
	}
	// MoqGenerator_mock isolates the mock interface of the Generator type
}

type MoqGenerator_mock struct {
	Moq *MoqGenerator
}

// MoqGenerator_recorder isolates the recorder interface of the Generator type
type MoqGenerator_recorder struct {
	Moq *MoqGenerator
}

// MoqGenerator_Generate_params holds the params of the Generator type
type MoqGenerator_Generate_params struct {
	Rand *rand.Rand
	Size int
}

// MoqGenerator_Generate_paramsKey holds the map key params of the Generator
// type
type MoqGenerator_Generate_paramsKey struct {
	Params struct {
		Rand *rand.Rand
		Size int
	}
	Hashes struct {
		Rand hash.Hash
		Size hash.Hash
	}
}

// MoqGenerator_Generate_resultsByParams contains the results for a given set
// of parameters for the Generator type
type MoqGenerator_Generate_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqGenerator_Generate_paramsKey]*MoqGenerator_Generate_results
}

// MoqGenerator_Generate_doFn defines the type of function needed when calling
// AndDo for the Generator type
type MoqGenerator_Generate_doFn func(rand *rand.Rand, size int)

// MoqGenerator_Generate_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Generator type
type MoqGenerator_Generate_doReturnFn func(rand *rand.Rand, size int) reflect.Value

// MoqGenerator_Generate_results holds the results of the Generator type
type MoqGenerator_Generate_results struct {
	Params  MoqGenerator_Generate_params
	Results []struct {
		Values *struct {
			Result1 reflect.Value
		}
		Sequence   uint32
		DoFn       MoqGenerator_Generate_doFn
		DoReturnFn MoqGenerator_Generate_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqGenerator_Generate_fnRecorder routes recorded function calls to the
// MoqGenerator moq
type MoqGenerator_Generate_fnRecorder struct {
	Params    MoqGenerator_Generate_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqGenerator_Generate_results
	Moq       *MoqGenerator
}

// MoqGenerator_Generate_anyParams isolates the any params functions of the
// Generator type
type MoqGenerator_Generate_anyParams struct {
	Recorder *MoqGenerator_Generate_fnRecorder
}

// NewMoqGenerator creates a new moq of the Generator type
func NewMoqGenerator(scene *moq.Scene, config *moq.Config) *MoqGenerator {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqGenerator{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqGenerator_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Generate struct {
					Rand moq.ParamIndexing
					Size moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Generate struct {
				Rand moq.ParamIndexing
				Size moq.ParamIndexing
			}
		}{
			Generate: struct {
				Rand moq.ParamIndexing
				Size moq.ParamIndexing
			}{
				Rand: moq.ParamIndexByHash,
				Size: moq.ParamIndexByValue,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Generator type
func (m *MoqGenerator) Mock() *MoqGenerator_mock { return m.Moq }

func (m *MoqGenerator_mock) Generate(rand *rand.Rand, size int) (result1 reflect.Value) {
	m.Moq.Scene.T.Helper()
	params := MoqGenerator_Generate_params{
		Rand: rand,
		Size: size,
	}
	var results *MoqGenerator_Generate_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Generate {
		paramsKey := m.Moq.ParamsKey_Generate(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Generate(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Generate(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Generate(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(rand, size)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(rand, size)
	}
	return
}

// OnCall returns the recorder implementation of the Generator type
func (m *MoqGenerator) OnCall() *MoqGenerator_recorder {
	return &MoqGenerator_recorder{
		Moq: m,
	}
}

func (m *MoqGenerator_recorder) Generate(rand *rand.Rand, size int) *MoqGenerator_Generate_fnRecorder {
	return &MoqGenerator_Generate_fnRecorder{
		Params: MoqGenerator_Generate_params{
			Rand: rand,
			Size: size,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqGenerator_Generate_fnRecorder) Any() *MoqGenerator_Generate_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Generate(r.Params))
		return nil
	}
	return &MoqGenerator_Generate_anyParams{Recorder: r}
}

func (a *MoqGenerator_Generate_anyParams) Rand() *MoqGenerator_Generate_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqGenerator_Generate_anyParams) Size() *MoqGenerator_Generate_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqGenerator_Generate_fnRecorder) Seq() *MoqGenerator_Generate_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Generate(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqGenerator_Generate_fnRecorder) NoSeq() *MoqGenerator_Generate_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Generate(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqGenerator_Generate_fnRecorder) ReturnResults(result1 reflect.Value) *MoqGenerator_Generate_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 reflect.Value
		}
		Sequence   uint32
		DoFn       MoqGenerator_Generate_doFn
		DoReturnFn MoqGenerator_Generate_doReturnFn
	}{
		Values: &struct {
			Result1 reflect.Value
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqGenerator_Generate_fnRecorder) AndDo(fn MoqGenerator_Generate_doFn) *MoqGenerator_Generate_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqGenerator_Generate_fnRecorder) DoReturnResults(fn MoqGenerator_Generate_doReturnFn) *MoqGenerator_Generate_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 reflect.Value
		}
		Sequence   uint32
		DoFn       MoqGenerator_Generate_doFn
		DoReturnFn MoqGenerator_Generate_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqGenerator_Generate_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqGenerator_Generate_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Generate {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqGenerator_Generate_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqGenerator_Generate_paramsKey]*MoqGenerator_Generate_results{},
		}
		r.Moq.ResultsByParams_Generate = append(r.Moq.ResultsByParams_Generate, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Generate) {
			copy(r.Moq.ResultsByParams_Generate[insertAt+1:], r.Moq.ResultsByParams_Generate[insertAt:0])
			r.Moq.ResultsByParams_Generate[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Generate(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqGenerator_Generate_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqGenerator_Generate_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqGenerator_Generate_fnRecorder {
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
					Result1 reflect.Value
				}
				Sequence   uint32
				DoFn       MoqGenerator_Generate_doFn
				DoReturnFn MoqGenerator_Generate_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqGenerator) PrettyParams_Generate(params MoqGenerator_Generate_params) string {
	return fmt.Sprintf("Generate(%#v, %#v)", params.Rand, params.Size)
}

func (m *MoqGenerator) ParamsKey_Generate(params MoqGenerator_Generate_params, anyParams uint64) MoqGenerator_Generate_paramsKey {
	m.Scene.T.Helper()
	var randUsed *rand.Rand
	var randUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Generate.Rand == moq.ParamIndexByValue {
			randUsed = params.Rand
		} else {
			randUsedHash = hash.DeepHash(params.Rand)
		}
	}
	var sizeUsed int
	var sizeUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Generate.Size == moq.ParamIndexByValue {
			sizeUsed = params.Size
		} else {
			sizeUsedHash = hash.DeepHash(params.Size)
		}
	}
	return MoqGenerator_Generate_paramsKey{
		Params: struct {
			Rand *rand.Rand
			Size int
		}{
			Rand: randUsed,
			Size: sizeUsed,
		},
		Hashes: struct {
			Rand hash.Hash
			Size hash.Hash
		}{
			Rand: randUsedHash,
			Size: sizeUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqGenerator) Reset() { m.ResultsByParams_Generate = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqGenerator) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Generate {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Generate(results.Params))
			}
		}
	}
}
