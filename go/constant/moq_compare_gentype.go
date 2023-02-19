// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package constant

import (
	"fmt"
	"go/constant"
	"go/token"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Compare_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Compare_genType func(x_ constant.Value, op token.Token, y_ constant.Value) bool

// MoqCompare_genType holds the state of a moq of the Compare_genType type
type MoqCompare_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqCompare_genType_mock

	ResultsByParams []MoqCompare_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			X_ moq.ParamIndexing
			Op moq.ParamIndexing
			Y_ moq.ParamIndexing
		}
	}
}

// MoqCompare_genType_mock isolates the mock interface of the Compare_genType
// type
type MoqCompare_genType_mock struct {
	Moq *MoqCompare_genType
}

// MoqCompare_genType_params holds the params of the Compare_genType type
type MoqCompare_genType_params struct {
	X_ constant.Value
	Op token.Token
	Y_ constant.Value
}

// MoqCompare_genType_paramsKey holds the map key params of the Compare_genType
// type
type MoqCompare_genType_paramsKey struct {
	Params struct {
		X_ constant.Value
		Op token.Token
		Y_ constant.Value
	}
	Hashes struct {
		X_ hash.Hash
		Op hash.Hash
		Y_ hash.Hash
	}
}

// MoqCompare_genType_resultsByParams contains the results for a given set of
// parameters for the Compare_genType type
type MoqCompare_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqCompare_genType_paramsKey]*MoqCompare_genType_results
}

// MoqCompare_genType_doFn defines the type of function needed when calling
// AndDo for the Compare_genType type
type MoqCompare_genType_doFn func(x_ constant.Value, op token.Token, y_ constant.Value)

// MoqCompare_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Compare_genType type
type MoqCompare_genType_doReturnFn func(x_ constant.Value, op token.Token, y_ constant.Value) bool

// MoqCompare_genType_results holds the results of the Compare_genType type
type MoqCompare_genType_results struct {
	Params  MoqCompare_genType_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqCompare_genType_doFn
		DoReturnFn MoqCompare_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqCompare_genType_fnRecorder routes recorded function calls to the
// MoqCompare_genType moq
type MoqCompare_genType_fnRecorder struct {
	Params    MoqCompare_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqCompare_genType_results
	Moq       *MoqCompare_genType
}

// MoqCompare_genType_anyParams isolates the any params functions of the
// Compare_genType type
type MoqCompare_genType_anyParams struct {
	Recorder *MoqCompare_genType_fnRecorder
}

// NewMoqCompare_genType creates a new moq of the Compare_genType type
func NewMoqCompare_genType(scene *moq.Scene, config *moq.Config) *MoqCompare_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqCompare_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqCompare_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				X_ moq.ParamIndexing
				Op moq.ParamIndexing
				Y_ moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			X_ moq.ParamIndexing
			Op moq.ParamIndexing
			Y_ moq.ParamIndexing
		}{
			X_: moq.ParamIndexByHash,
			Op: moq.ParamIndexByValue,
			Y_: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Compare_genType type
func (m *MoqCompare_genType) Mock() Compare_genType {
	return func(x_ constant.Value, op token.Token, y_ constant.Value) bool {
		m.Scene.T.Helper()
		moq := &MoqCompare_genType_mock{Moq: m}
		return moq.Fn(x_, op, y_)
	}
}

func (m *MoqCompare_genType_mock) Fn(x_ constant.Value, op token.Token, y_ constant.Value) (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqCompare_genType_params{
		X_: x_,
		Op: op,
		Y_: y_,
	}
	var results *MoqCompare_genType_results
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
		result.DoFn(x_, op, y_)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(x_, op, y_)
	}
	return
}

func (m *MoqCompare_genType) OnCall(x_ constant.Value, op token.Token, y_ constant.Value) *MoqCompare_genType_fnRecorder {
	return &MoqCompare_genType_fnRecorder{
		Params: MoqCompare_genType_params{
			X_: x_,
			Op: op,
			Y_: y_,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqCompare_genType_fnRecorder) Any() *MoqCompare_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqCompare_genType_anyParams{Recorder: r}
}

func (a *MoqCompare_genType_anyParams) X_() *MoqCompare_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqCompare_genType_anyParams) Op() *MoqCompare_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqCompare_genType_anyParams) Y_() *MoqCompare_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqCompare_genType_fnRecorder) Seq() *MoqCompare_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqCompare_genType_fnRecorder) NoSeq() *MoqCompare_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqCompare_genType_fnRecorder) ReturnResults(result1 bool) *MoqCompare_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqCompare_genType_doFn
		DoReturnFn MoqCompare_genType_doReturnFn
	}{
		Values: &struct {
			Result1 bool
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqCompare_genType_fnRecorder) AndDo(fn MoqCompare_genType_doFn) *MoqCompare_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqCompare_genType_fnRecorder) DoReturnResults(fn MoqCompare_genType_doReturnFn) *MoqCompare_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqCompare_genType_doFn
		DoReturnFn MoqCompare_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqCompare_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqCompare_genType_resultsByParams
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
		results = &MoqCompare_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqCompare_genType_paramsKey]*MoqCompare_genType_results{},
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
		r.Results = &MoqCompare_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqCompare_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqCompare_genType_fnRecorder {
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
					Result1 bool
				}
				Sequence   uint32
				DoFn       MoqCompare_genType_doFn
				DoReturnFn MoqCompare_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqCompare_genType) PrettyParams(params MoqCompare_genType_params) string {
	return fmt.Sprintf("Compare_genType(%#v, %#v, %#v)", params.X_, params.Op, params.Y_)
}

func (m *MoqCompare_genType) ParamsKey(params MoqCompare_genType_params, anyParams uint64) MoqCompare_genType_paramsKey {
	m.Scene.T.Helper()
	var x_Used constant.Value
	var x_UsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.X_ == moq.ParamIndexByValue {
			x_Used = params.X_
		} else {
			x_UsedHash = hash.DeepHash(params.X_)
		}
	}
	var opUsed token.Token
	var opUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Op == moq.ParamIndexByValue {
			opUsed = params.Op
		} else {
			opUsedHash = hash.DeepHash(params.Op)
		}
	}
	var y_Used constant.Value
	var y_UsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Y_ == moq.ParamIndexByValue {
			y_Used = params.Y_
		} else {
			y_UsedHash = hash.DeepHash(params.Y_)
		}
	}
	return MoqCompare_genType_paramsKey{
		Params: struct {
			X_ constant.Value
			Op token.Token
			Y_ constant.Value
		}{
			X_: x_Used,
			Op: opUsed,
			Y_: y_Used,
		},
		Hashes: struct {
			X_ hash.Hash
			Op hash.Hash
			Y_ hash.Hash
		}{
			X_: x_UsedHash,
			Op: opUsedHash,
			Y_: y_UsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqCompare_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqCompare_genType) AssertExpectationsMet() {
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
