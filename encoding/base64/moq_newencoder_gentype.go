// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package base64

import (
	"encoding/base64"
	"fmt"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// NewEncoder_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type NewEncoder_genType func(enc *base64.Encoding, w io.Writer) io.WriteCloser

// MoqNewEncoder_genType holds the state of a moq of the NewEncoder_genType
// type
type MoqNewEncoder_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNewEncoder_genType_mock

	ResultsByParams []MoqNewEncoder_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Enc moq.ParamIndexing
			W   moq.ParamIndexing
		}
	}
}

// MoqNewEncoder_genType_mock isolates the mock interface of the
// NewEncoder_genType type
type MoqNewEncoder_genType_mock struct {
	Moq *MoqNewEncoder_genType
}

// MoqNewEncoder_genType_params holds the params of the NewEncoder_genType type
type MoqNewEncoder_genType_params struct {
	Enc *base64.Encoding
	W   io.Writer
}

// MoqNewEncoder_genType_paramsKey holds the map key params of the
// NewEncoder_genType type
type MoqNewEncoder_genType_paramsKey struct {
	Params struct {
		Enc *base64.Encoding
		W   io.Writer
	}
	Hashes struct {
		Enc hash.Hash
		W   hash.Hash
	}
}

// MoqNewEncoder_genType_resultsByParams contains the results for a given set
// of parameters for the NewEncoder_genType type
type MoqNewEncoder_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNewEncoder_genType_paramsKey]*MoqNewEncoder_genType_results
}

// MoqNewEncoder_genType_doFn defines the type of function needed when calling
// AndDo for the NewEncoder_genType type
type MoqNewEncoder_genType_doFn func(enc *base64.Encoding, w io.Writer)

// MoqNewEncoder_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the NewEncoder_genType type
type MoqNewEncoder_genType_doReturnFn func(enc *base64.Encoding, w io.Writer) io.WriteCloser

// MoqNewEncoder_genType_results holds the results of the NewEncoder_genType
// type
type MoqNewEncoder_genType_results struct {
	Params  MoqNewEncoder_genType_params
	Results []struct {
		Values *struct {
			Result1 io.WriteCloser
		}
		Sequence   uint32
		DoFn       MoqNewEncoder_genType_doFn
		DoReturnFn MoqNewEncoder_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNewEncoder_genType_fnRecorder routes recorded function calls to the
// MoqNewEncoder_genType moq
type MoqNewEncoder_genType_fnRecorder struct {
	Params    MoqNewEncoder_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNewEncoder_genType_results
	Moq       *MoqNewEncoder_genType
}

// MoqNewEncoder_genType_anyParams isolates the any params functions of the
// NewEncoder_genType type
type MoqNewEncoder_genType_anyParams struct {
	Recorder *MoqNewEncoder_genType_fnRecorder
}

// NewMoqNewEncoder_genType creates a new moq of the NewEncoder_genType type
func NewMoqNewEncoder_genType(scene *moq.Scene, config *moq.Config) *MoqNewEncoder_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNewEncoder_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNewEncoder_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Enc moq.ParamIndexing
				W   moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Enc moq.ParamIndexing
			W   moq.ParamIndexing
		}{
			Enc: moq.ParamIndexByHash,
			W:   moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the NewEncoder_genType type
func (m *MoqNewEncoder_genType) Mock() NewEncoder_genType {
	return func(enc *base64.Encoding, w io.Writer) io.WriteCloser {
		m.Scene.T.Helper()
		moq := &MoqNewEncoder_genType_mock{Moq: m}
		return moq.Fn(enc, w)
	}
}

func (m *MoqNewEncoder_genType_mock) Fn(enc *base64.Encoding, w io.Writer) (result1 io.WriteCloser) {
	m.Moq.Scene.T.Helper()
	params := MoqNewEncoder_genType_params{
		Enc: enc,
		W:   w,
	}
	var results *MoqNewEncoder_genType_results
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
		result.DoFn(enc, w)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(enc, w)
	}
	return
}

func (m *MoqNewEncoder_genType) OnCall(enc *base64.Encoding, w io.Writer) *MoqNewEncoder_genType_fnRecorder {
	return &MoqNewEncoder_genType_fnRecorder{
		Params: MoqNewEncoder_genType_params{
			Enc: enc,
			W:   w,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNewEncoder_genType_fnRecorder) Any() *MoqNewEncoder_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNewEncoder_genType_anyParams{Recorder: r}
}

func (a *MoqNewEncoder_genType_anyParams) Enc() *MoqNewEncoder_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqNewEncoder_genType_anyParams) W() *MoqNewEncoder_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqNewEncoder_genType_fnRecorder) Seq() *MoqNewEncoder_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNewEncoder_genType_fnRecorder) NoSeq() *MoqNewEncoder_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNewEncoder_genType_fnRecorder) ReturnResults(result1 io.WriteCloser) *MoqNewEncoder_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 io.WriteCloser
		}
		Sequence   uint32
		DoFn       MoqNewEncoder_genType_doFn
		DoReturnFn MoqNewEncoder_genType_doReturnFn
	}{
		Values: &struct {
			Result1 io.WriteCloser
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNewEncoder_genType_fnRecorder) AndDo(fn MoqNewEncoder_genType_doFn) *MoqNewEncoder_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNewEncoder_genType_fnRecorder) DoReturnResults(fn MoqNewEncoder_genType_doReturnFn) *MoqNewEncoder_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 io.WriteCloser
		}
		Sequence   uint32
		DoFn       MoqNewEncoder_genType_doFn
		DoReturnFn MoqNewEncoder_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNewEncoder_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNewEncoder_genType_resultsByParams
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
		results = &MoqNewEncoder_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNewEncoder_genType_paramsKey]*MoqNewEncoder_genType_results{},
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
		r.Results = &MoqNewEncoder_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNewEncoder_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNewEncoder_genType_fnRecorder {
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
					Result1 io.WriteCloser
				}
				Sequence   uint32
				DoFn       MoqNewEncoder_genType_doFn
				DoReturnFn MoqNewEncoder_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNewEncoder_genType) PrettyParams(params MoqNewEncoder_genType_params) string {
	return fmt.Sprintf("NewEncoder_genType(%#v, %#v)", params.Enc, params.W)
}

func (m *MoqNewEncoder_genType) ParamsKey(params MoqNewEncoder_genType_params, anyParams uint64) MoqNewEncoder_genType_paramsKey {
	m.Scene.T.Helper()
	var encUsed *base64.Encoding
	var encUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Enc == moq.ParamIndexByValue {
			encUsed = params.Enc
		} else {
			encUsedHash = hash.DeepHash(params.Enc)
		}
	}
	var wUsed io.Writer
	var wUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.W == moq.ParamIndexByValue {
			wUsed = params.W
		} else {
			wUsedHash = hash.DeepHash(params.W)
		}
	}
	return MoqNewEncoder_genType_paramsKey{
		Params: struct {
			Enc *base64.Encoding
			W   io.Writer
		}{
			Enc: encUsed,
			W:   wUsed,
		},
		Hashes: struct {
			Enc hash.Hash
			W   hash.Hash
		}{
			Enc: encUsedHash,
			W:   wUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNewEncoder_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNewEncoder_genType) AssertExpectationsMet() {
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
