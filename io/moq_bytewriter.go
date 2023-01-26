// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package io

import (
	"fmt"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that io.ByteWriter is mocked completely
var _ io.ByteWriter = (*MoqByteWriter_mock)(nil)

// MoqByteWriter holds the state of a moq of the ByteWriter type
type MoqByteWriter struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqByteWriter_mock

	ResultsByParams_WriteByte []MoqByteWriter_WriteByte_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			WriteByte struct {
				C moq.ParamIndexing
			}
		}
	}
	// MoqByteWriter_mock isolates the mock interface of the ByteWriter type
}

type MoqByteWriter_mock struct {
	Moq *MoqByteWriter
}

// MoqByteWriter_recorder isolates the recorder interface of the ByteWriter
// type
type MoqByteWriter_recorder struct {
	Moq *MoqByteWriter
}

// MoqByteWriter_WriteByte_params holds the params of the ByteWriter type
type MoqByteWriter_WriteByte_params struct{ C byte }

// MoqByteWriter_WriteByte_paramsKey holds the map key params of the ByteWriter
// type
type MoqByteWriter_WriteByte_paramsKey struct {
	Params struct{ C byte }
	Hashes struct{ C hash.Hash }
}

// MoqByteWriter_WriteByte_resultsByParams contains the results for a given set
// of parameters for the ByteWriter type
type MoqByteWriter_WriteByte_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqByteWriter_WriteByte_paramsKey]*MoqByteWriter_WriteByte_results
}

// MoqByteWriter_WriteByte_doFn defines the type of function needed when
// calling AndDo for the ByteWriter type
type MoqByteWriter_WriteByte_doFn func(c byte)

// MoqByteWriter_WriteByte_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ByteWriter type
type MoqByteWriter_WriteByte_doReturnFn func(c byte) error

// MoqByteWriter_WriteByte_results holds the results of the ByteWriter type
type MoqByteWriter_WriteByte_results struct {
	Params  MoqByteWriter_WriteByte_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqByteWriter_WriteByte_doFn
		DoReturnFn MoqByteWriter_WriteByte_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqByteWriter_WriteByte_fnRecorder routes recorded function calls to the
// MoqByteWriter moq
type MoqByteWriter_WriteByte_fnRecorder struct {
	Params    MoqByteWriter_WriteByte_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqByteWriter_WriteByte_results
	Moq       *MoqByteWriter
}

// MoqByteWriter_WriteByte_anyParams isolates the any params functions of the
// ByteWriter type
type MoqByteWriter_WriteByte_anyParams struct {
	Recorder *MoqByteWriter_WriteByte_fnRecorder
}

// NewMoqByteWriter creates a new moq of the ByteWriter type
func NewMoqByteWriter(scene *moq.Scene, config *moq.Config) *MoqByteWriter {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqByteWriter{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqByteWriter_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				WriteByte struct {
					C moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			WriteByte struct {
				C moq.ParamIndexing
			}
		}{
			WriteByte: struct {
				C moq.ParamIndexing
			}{
				C: moq.ParamIndexByValue,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the ByteWriter type
func (m *MoqByteWriter) Mock() *MoqByteWriter_mock { return m.Moq }

func (m *MoqByteWriter_mock) WriteByte(c byte) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqByteWriter_WriteByte_params{
		C: c,
	}
	var results *MoqByteWriter_WriteByte_results
	for _, resultsByParams := range m.Moq.ResultsByParams_WriteByte {
		paramsKey := m.Moq.ParamsKey_WriteByte(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_WriteByte(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_WriteByte(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_WriteByte(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(c)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(c)
	}
	return
}

// OnCall returns the recorder implementation of the ByteWriter type
func (m *MoqByteWriter) OnCall() *MoqByteWriter_recorder {
	return &MoqByteWriter_recorder{
		Moq: m,
	}
}

func (m *MoqByteWriter_recorder) WriteByte(c byte) *MoqByteWriter_WriteByte_fnRecorder {
	return &MoqByteWriter_WriteByte_fnRecorder{
		Params: MoqByteWriter_WriteByte_params{
			C: c,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqByteWriter_WriteByte_fnRecorder) Any() *MoqByteWriter_WriteByte_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_WriteByte(r.Params))
		return nil
	}
	return &MoqByteWriter_WriteByte_anyParams{Recorder: r}
}

func (a *MoqByteWriter_WriteByte_anyParams) C() *MoqByteWriter_WriteByte_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqByteWriter_WriteByte_fnRecorder) Seq() *MoqByteWriter_WriteByte_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_WriteByte(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqByteWriter_WriteByte_fnRecorder) NoSeq() *MoqByteWriter_WriteByte_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_WriteByte(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqByteWriter_WriteByte_fnRecorder) ReturnResults(result1 error) *MoqByteWriter_WriteByte_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqByteWriter_WriteByte_doFn
		DoReturnFn MoqByteWriter_WriteByte_doReturnFn
	}{
		Values: &struct {
			Result1 error
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqByteWriter_WriteByte_fnRecorder) AndDo(fn MoqByteWriter_WriteByte_doFn) *MoqByteWriter_WriteByte_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqByteWriter_WriteByte_fnRecorder) DoReturnResults(fn MoqByteWriter_WriteByte_doReturnFn) *MoqByteWriter_WriteByte_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqByteWriter_WriteByte_doFn
		DoReturnFn MoqByteWriter_WriteByte_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqByteWriter_WriteByte_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqByteWriter_WriteByte_resultsByParams
	for n, res := range r.Moq.ResultsByParams_WriteByte {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqByteWriter_WriteByte_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqByteWriter_WriteByte_paramsKey]*MoqByteWriter_WriteByte_results{},
		}
		r.Moq.ResultsByParams_WriteByte = append(r.Moq.ResultsByParams_WriteByte, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_WriteByte) {
			copy(r.Moq.ResultsByParams_WriteByte[insertAt+1:], r.Moq.ResultsByParams_WriteByte[insertAt:0])
			r.Moq.ResultsByParams_WriteByte[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_WriteByte(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqByteWriter_WriteByte_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqByteWriter_WriteByte_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqByteWriter_WriteByte_fnRecorder {
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
					Result1 error
				}
				Sequence   uint32
				DoFn       MoqByteWriter_WriteByte_doFn
				DoReturnFn MoqByteWriter_WriteByte_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqByteWriter) PrettyParams_WriteByte(params MoqByteWriter_WriteByte_params) string {
	return fmt.Sprintf("WriteByte(%#v)", params.C)
}

func (m *MoqByteWriter) ParamsKey_WriteByte(params MoqByteWriter_WriteByte_params, anyParams uint64) MoqByteWriter_WriteByte_paramsKey {
	m.Scene.T.Helper()
	var cUsed byte
	var cUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.WriteByte.C == moq.ParamIndexByValue {
			cUsed = params.C
		} else {
			cUsedHash = hash.DeepHash(params.C)
		}
	}
	return MoqByteWriter_WriteByte_paramsKey{
		Params: struct{ C byte }{
			C: cUsed,
		},
		Hashes: struct{ C hash.Hash }{
			C: cUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqByteWriter) Reset() { m.ResultsByParams_WriteByte = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqByteWriter) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_WriteByte {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_WriteByte(results.Params))
			}
		}
	}
}
