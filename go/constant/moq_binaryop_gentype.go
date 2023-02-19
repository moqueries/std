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

// BinaryOp_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type BinaryOp_genType func(x_ constant.Value, op token.Token, y_ constant.Value) constant.Value

// MoqBinaryOp_genType holds the state of a moq of the BinaryOp_genType type
type MoqBinaryOp_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqBinaryOp_genType_mock

	ResultsByParams []MoqBinaryOp_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			X_ moq.ParamIndexing
			Op moq.ParamIndexing
			Y_ moq.ParamIndexing
		}
	}
}

// MoqBinaryOp_genType_mock isolates the mock interface of the BinaryOp_genType
// type
type MoqBinaryOp_genType_mock struct {
	Moq *MoqBinaryOp_genType
}

// MoqBinaryOp_genType_params holds the params of the BinaryOp_genType type
type MoqBinaryOp_genType_params struct {
	X_ constant.Value
	Op token.Token
	Y_ constant.Value
}

// MoqBinaryOp_genType_paramsKey holds the map key params of the
// BinaryOp_genType type
type MoqBinaryOp_genType_paramsKey struct {
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

// MoqBinaryOp_genType_resultsByParams contains the results for a given set of
// parameters for the BinaryOp_genType type
type MoqBinaryOp_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqBinaryOp_genType_paramsKey]*MoqBinaryOp_genType_results
}

// MoqBinaryOp_genType_doFn defines the type of function needed when calling
// AndDo for the BinaryOp_genType type
type MoqBinaryOp_genType_doFn func(x_ constant.Value, op token.Token, y_ constant.Value)

// MoqBinaryOp_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the BinaryOp_genType type
type MoqBinaryOp_genType_doReturnFn func(x_ constant.Value, op token.Token, y_ constant.Value) constant.Value

// MoqBinaryOp_genType_results holds the results of the BinaryOp_genType type
type MoqBinaryOp_genType_results struct {
	Params  MoqBinaryOp_genType_params
	Results []struct {
		Values *struct {
			Result1 constant.Value
		}
		Sequence   uint32
		DoFn       MoqBinaryOp_genType_doFn
		DoReturnFn MoqBinaryOp_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqBinaryOp_genType_fnRecorder routes recorded function calls to the
// MoqBinaryOp_genType moq
type MoqBinaryOp_genType_fnRecorder struct {
	Params    MoqBinaryOp_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqBinaryOp_genType_results
	Moq       *MoqBinaryOp_genType
}

// MoqBinaryOp_genType_anyParams isolates the any params functions of the
// BinaryOp_genType type
type MoqBinaryOp_genType_anyParams struct {
	Recorder *MoqBinaryOp_genType_fnRecorder
}

// NewMoqBinaryOp_genType creates a new moq of the BinaryOp_genType type
func NewMoqBinaryOp_genType(scene *moq.Scene, config *moq.Config) *MoqBinaryOp_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqBinaryOp_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqBinaryOp_genType_mock{},

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

// Mock returns the moq implementation of the BinaryOp_genType type
func (m *MoqBinaryOp_genType) Mock() BinaryOp_genType {
	return func(x_ constant.Value, op token.Token, y_ constant.Value) constant.Value {
		m.Scene.T.Helper()
		moq := &MoqBinaryOp_genType_mock{Moq: m}
		return moq.Fn(x_, op, y_)
	}
}

func (m *MoqBinaryOp_genType_mock) Fn(x_ constant.Value, op token.Token, y_ constant.Value) (result1 constant.Value) {
	m.Moq.Scene.T.Helper()
	params := MoqBinaryOp_genType_params{
		X_: x_,
		Op: op,
		Y_: y_,
	}
	var results *MoqBinaryOp_genType_results
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

func (m *MoqBinaryOp_genType) OnCall(x_ constant.Value, op token.Token, y_ constant.Value) *MoqBinaryOp_genType_fnRecorder {
	return &MoqBinaryOp_genType_fnRecorder{
		Params: MoqBinaryOp_genType_params{
			X_: x_,
			Op: op,
			Y_: y_,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqBinaryOp_genType_fnRecorder) Any() *MoqBinaryOp_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqBinaryOp_genType_anyParams{Recorder: r}
}

func (a *MoqBinaryOp_genType_anyParams) X_() *MoqBinaryOp_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqBinaryOp_genType_anyParams) Op() *MoqBinaryOp_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqBinaryOp_genType_anyParams) Y_() *MoqBinaryOp_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqBinaryOp_genType_fnRecorder) Seq() *MoqBinaryOp_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqBinaryOp_genType_fnRecorder) NoSeq() *MoqBinaryOp_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqBinaryOp_genType_fnRecorder) ReturnResults(result1 constant.Value) *MoqBinaryOp_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 constant.Value
		}
		Sequence   uint32
		DoFn       MoqBinaryOp_genType_doFn
		DoReturnFn MoqBinaryOp_genType_doReturnFn
	}{
		Values: &struct {
			Result1 constant.Value
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqBinaryOp_genType_fnRecorder) AndDo(fn MoqBinaryOp_genType_doFn) *MoqBinaryOp_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqBinaryOp_genType_fnRecorder) DoReturnResults(fn MoqBinaryOp_genType_doReturnFn) *MoqBinaryOp_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 constant.Value
		}
		Sequence   uint32
		DoFn       MoqBinaryOp_genType_doFn
		DoReturnFn MoqBinaryOp_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqBinaryOp_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqBinaryOp_genType_resultsByParams
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
		results = &MoqBinaryOp_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqBinaryOp_genType_paramsKey]*MoqBinaryOp_genType_results{},
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
		r.Results = &MoqBinaryOp_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqBinaryOp_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqBinaryOp_genType_fnRecorder {
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
					Result1 constant.Value
				}
				Sequence   uint32
				DoFn       MoqBinaryOp_genType_doFn
				DoReturnFn MoqBinaryOp_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqBinaryOp_genType) PrettyParams(params MoqBinaryOp_genType_params) string {
	return fmt.Sprintf("BinaryOp_genType(%#v, %#v, %#v)", params.X_, params.Op, params.Y_)
}

func (m *MoqBinaryOp_genType) ParamsKey(params MoqBinaryOp_genType_params, anyParams uint64) MoqBinaryOp_genType_paramsKey {
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
	return MoqBinaryOp_genType_paramsKey{
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
func (m *MoqBinaryOp_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqBinaryOp_genType) AssertExpectationsMet() {
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