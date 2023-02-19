// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package flag

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// IntVar_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type IntVar_genType func(p *int, name string, value int, usage string)

// MoqIntVar_genType holds the state of a moq of the IntVar_genType type
type MoqIntVar_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqIntVar_genType_mock

	ResultsByParams []MoqIntVar_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			P     moq.ParamIndexing
			Name  moq.ParamIndexing
			Value moq.ParamIndexing
			Usage moq.ParamIndexing
		}
	}
}

// MoqIntVar_genType_mock isolates the mock interface of the IntVar_genType
// type
type MoqIntVar_genType_mock struct {
	Moq *MoqIntVar_genType
}

// MoqIntVar_genType_params holds the params of the IntVar_genType type
type MoqIntVar_genType_params struct {
	P     *int
	Name  string
	Value int
	Usage string
}

// MoqIntVar_genType_paramsKey holds the map key params of the IntVar_genType
// type
type MoqIntVar_genType_paramsKey struct {
	Params struct {
		P     *int
		Name  string
		Value int
		Usage string
	}
	Hashes struct {
		P     hash.Hash
		Name  hash.Hash
		Value hash.Hash
		Usage hash.Hash
	}
}

// MoqIntVar_genType_resultsByParams contains the results for a given set of
// parameters for the IntVar_genType type
type MoqIntVar_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqIntVar_genType_paramsKey]*MoqIntVar_genType_results
}

// MoqIntVar_genType_doFn defines the type of function needed when calling
// AndDo for the IntVar_genType type
type MoqIntVar_genType_doFn func(p *int, name string, value int, usage string)

// MoqIntVar_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the IntVar_genType type
type MoqIntVar_genType_doReturnFn func(p *int, name string, value int, usage string)

// MoqIntVar_genType_results holds the results of the IntVar_genType type
type MoqIntVar_genType_results struct {
	Params  MoqIntVar_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqIntVar_genType_doFn
		DoReturnFn MoqIntVar_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqIntVar_genType_fnRecorder routes recorded function calls to the
// MoqIntVar_genType moq
type MoqIntVar_genType_fnRecorder struct {
	Params    MoqIntVar_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqIntVar_genType_results
	Moq       *MoqIntVar_genType
}

// MoqIntVar_genType_anyParams isolates the any params functions of the
// IntVar_genType type
type MoqIntVar_genType_anyParams struct {
	Recorder *MoqIntVar_genType_fnRecorder
}

// NewMoqIntVar_genType creates a new moq of the IntVar_genType type
func NewMoqIntVar_genType(scene *moq.Scene, config *moq.Config) *MoqIntVar_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqIntVar_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqIntVar_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				P     moq.ParamIndexing
				Name  moq.ParamIndexing
				Value moq.ParamIndexing
				Usage moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			P     moq.ParamIndexing
			Name  moq.ParamIndexing
			Value moq.ParamIndexing
			Usage moq.ParamIndexing
		}{
			P:     moq.ParamIndexByHash,
			Name:  moq.ParamIndexByValue,
			Value: moq.ParamIndexByValue,
			Usage: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the IntVar_genType type
func (m *MoqIntVar_genType) Mock() IntVar_genType {
	return func(p *int, name string, value int, usage string) {
		m.Scene.T.Helper()
		moq := &MoqIntVar_genType_mock{Moq: m}
		moq.Fn(p, name, value, usage)
	}
}

func (m *MoqIntVar_genType_mock) Fn(p *int, name string, value int, usage string) {
	m.Moq.Scene.T.Helper()
	params := MoqIntVar_genType_params{
		P:     p,
		Name:  name,
		Value: value,
		Usage: usage,
	}
	var results *MoqIntVar_genType_results
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
		result.DoFn(p, name, value, usage)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(p, name, value, usage)
	}
	return
}

func (m *MoqIntVar_genType) OnCall(p *int, name string, value int, usage string) *MoqIntVar_genType_fnRecorder {
	return &MoqIntVar_genType_fnRecorder{
		Params: MoqIntVar_genType_params{
			P:     p,
			Name:  name,
			Value: value,
			Usage: usage,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqIntVar_genType_fnRecorder) Any() *MoqIntVar_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqIntVar_genType_anyParams{Recorder: r}
}

func (a *MoqIntVar_genType_anyParams) P() *MoqIntVar_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqIntVar_genType_anyParams) Name() *MoqIntVar_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqIntVar_genType_anyParams) Value() *MoqIntVar_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (a *MoqIntVar_genType_anyParams) Usage() *MoqIntVar_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 3
	return a.Recorder
}

func (r *MoqIntVar_genType_fnRecorder) Seq() *MoqIntVar_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqIntVar_genType_fnRecorder) NoSeq() *MoqIntVar_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqIntVar_genType_fnRecorder) ReturnResults() *MoqIntVar_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqIntVar_genType_doFn
		DoReturnFn MoqIntVar_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqIntVar_genType_fnRecorder) AndDo(fn MoqIntVar_genType_doFn) *MoqIntVar_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqIntVar_genType_fnRecorder) DoReturnResults(fn MoqIntVar_genType_doReturnFn) *MoqIntVar_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqIntVar_genType_doFn
		DoReturnFn MoqIntVar_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqIntVar_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqIntVar_genType_resultsByParams
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
		results = &MoqIntVar_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqIntVar_genType_paramsKey]*MoqIntVar_genType_results{},
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
		r.Results = &MoqIntVar_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqIntVar_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqIntVar_genType_fnRecorder {
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
				DoFn       MoqIntVar_genType_doFn
				DoReturnFn MoqIntVar_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqIntVar_genType) PrettyParams(params MoqIntVar_genType_params) string {
	return fmt.Sprintf("IntVar_genType(%#v, %#v, %#v, %#v)", params.P, params.Name, params.Value, params.Usage)
}

func (m *MoqIntVar_genType) ParamsKey(params MoqIntVar_genType_params, anyParams uint64) MoqIntVar_genType_paramsKey {
	m.Scene.T.Helper()
	var pUsed *int
	var pUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.P == moq.ParamIndexByValue {
			pUsed = params.P
		} else {
			pUsedHash = hash.DeepHash(params.P)
		}
	}
	var nameUsed string
	var nameUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Name == moq.ParamIndexByValue {
			nameUsed = params.Name
		} else {
			nameUsedHash = hash.DeepHash(params.Name)
		}
	}
	var valueUsed int
	var valueUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Value == moq.ParamIndexByValue {
			valueUsed = params.Value
		} else {
			valueUsedHash = hash.DeepHash(params.Value)
		}
	}
	var usageUsed string
	var usageUsedHash hash.Hash
	if anyParams&(1<<3) == 0 {
		if m.Runtime.ParameterIndexing.Usage == moq.ParamIndexByValue {
			usageUsed = params.Usage
		} else {
			usageUsedHash = hash.DeepHash(params.Usage)
		}
	}
	return MoqIntVar_genType_paramsKey{
		Params: struct {
			P     *int
			Name  string
			Value int
			Usage string
		}{
			P:     pUsed,
			Name:  nameUsed,
			Value: valueUsed,
			Usage: usageUsed,
		},
		Hashes: struct {
			P     hash.Hash
			Name  hash.Hash
			Value hash.Hash
			Usage hash.Hash
		}{
			P:     pUsedHash,
			Name:  nameUsedHash,
			Value: valueUsedHash,
			Usage: usageUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqIntVar_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqIntVar_genType) AssertExpectationsMet() {
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
