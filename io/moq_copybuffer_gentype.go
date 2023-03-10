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

// CopyBuffer_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type CopyBuffer_genType func(dst io.Writer, src io.Reader, buf []byte) (written int64, err error)

// MoqCopyBuffer_genType holds the state of a moq of the CopyBuffer_genType
// type
type MoqCopyBuffer_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqCopyBuffer_genType_mock

	ResultsByParams []MoqCopyBuffer_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Dst moq.ParamIndexing
			Src moq.ParamIndexing
			Buf moq.ParamIndexing
		}
	}
}

// MoqCopyBuffer_genType_mock isolates the mock interface of the
// CopyBuffer_genType type
type MoqCopyBuffer_genType_mock struct {
	Moq *MoqCopyBuffer_genType
}

// MoqCopyBuffer_genType_params holds the params of the CopyBuffer_genType type
type MoqCopyBuffer_genType_params struct {
	Dst io.Writer
	Src io.Reader
	Buf []byte
}

// MoqCopyBuffer_genType_paramsKey holds the map key params of the
// CopyBuffer_genType type
type MoqCopyBuffer_genType_paramsKey struct {
	Params struct {
		Dst io.Writer
		Src io.Reader
	}
	Hashes struct {
		Dst hash.Hash
		Src hash.Hash
		Buf hash.Hash
	}
}

// MoqCopyBuffer_genType_resultsByParams contains the results for a given set
// of parameters for the CopyBuffer_genType type
type MoqCopyBuffer_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqCopyBuffer_genType_paramsKey]*MoqCopyBuffer_genType_results
}

// MoqCopyBuffer_genType_doFn defines the type of function needed when calling
// AndDo for the CopyBuffer_genType type
type MoqCopyBuffer_genType_doFn func(dst io.Writer, src io.Reader, buf []byte)

// MoqCopyBuffer_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the CopyBuffer_genType type
type MoqCopyBuffer_genType_doReturnFn func(dst io.Writer, src io.Reader, buf []byte) (written int64, err error)

// MoqCopyBuffer_genType_results holds the results of the CopyBuffer_genType
// type
type MoqCopyBuffer_genType_results struct {
	Params  MoqCopyBuffer_genType_params
	Results []struct {
		Values *struct {
			Written int64
			Err     error
		}
		Sequence   uint32
		DoFn       MoqCopyBuffer_genType_doFn
		DoReturnFn MoqCopyBuffer_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqCopyBuffer_genType_fnRecorder routes recorded function calls to the
// MoqCopyBuffer_genType moq
type MoqCopyBuffer_genType_fnRecorder struct {
	Params    MoqCopyBuffer_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqCopyBuffer_genType_results
	Moq       *MoqCopyBuffer_genType
}

// MoqCopyBuffer_genType_anyParams isolates the any params functions of the
// CopyBuffer_genType type
type MoqCopyBuffer_genType_anyParams struct {
	Recorder *MoqCopyBuffer_genType_fnRecorder
}

// NewMoqCopyBuffer_genType creates a new moq of the CopyBuffer_genType type
func NewMoqCopyBuffer_genType(scene *moq.Scene, config *moq.Config) *MoqCopyBuffer_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqCopyBuffer_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqCopyBuffer_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Dst moq.ParamIndexing
				Src moq.ParamIndexing
				Buf moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Dst moq.ParamIndexing
			Src moq.ParamIndexing
			Buf moq.ParamIndexing
		}{
			Dst: moq.ParamIndexByHash,
			Src: moq.ParamIndexByHash,
			Buf: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the CopyBuffer_genType type
func (m *MoqCopyBuffer_genType) Mock() CopyBuffer_genType {
	return func(dst io.Writer, src io.Reader, buf []byte) (_ int64, _ error) {
		m.Scene.T.Helper()
		moq := &MoqCopyBuffer_genType_mock{Moq: m}
		return moq.Fn(dst, src, buf)
	}
}

func (m *MoqCopyBuffer_genType_mock) Fn(dst io.Writer, src io.Reader, buf []byte) (written int64, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqCopyBuffer_genType_params{
		Dst: dst,
		Src: src,
		Buf: buf,
	}
	var results *MoqCopyBuffer_genType_results
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
		result.DoFn(dst, src, buf)
	}

	if result.Values != nil {
		written = result.Values.Written
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		written, err = result.DoReturnFn(dst, src, buf)
	}
	return
}

func (m *MoqCopyBuffer_genType) OnCall(dst io.Writer, src io.Reader, buf []byte) *MoqCopyBuffer_genType_fnRecorder {
	return &MoqCopyBuffer_genType_fnRecorder{
		Params: MoqCopyBuffer_genType_params{
			Dst: dst,
			Src: src,
			Buf: buf,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqCopyBuffer_genType_fnRecorder) Any() *MoqCopyBuffer_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqCopyBuffer_genType_anyParams{Recorder: r}
}

func (a *MoqCopyBuffer_genType_anyParams) Dst() *MoqCopyBuffer_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqCopyBuffer_genType_anyParams) Src() *MoqCopyBuffer_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqCopyBuffer_genType_anyParams) Buf() *MoqCopyBuffer_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqCopyBuffer_genType_fnRecorder) Seq() *MoqCopyBuffer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqCopyBuffer_genType_fnRecorder) NoSeq() *MoqCopyBuffer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqCopyBuffer_genType_fnRecorder) ReturnResults(written int64, err error) *MoqCopyBuffer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Written int64
			Err     error
		}
		Sequence   uint32
		DoFn       MoqCopyBuffer_genType_doFn
		DoReturnFn MoqCopyBuffer_genType_doReturnFn
	}{
		Values: &struct {
			Written int64
			Err     error
		}{
			Written: written,
			Err:     err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqCopyBuffer_genType_fnRecorder) AndDo(fn MoqCopyBuffer_genType_doFn) *MoqCopyBuffer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqCopyBuffer_genType_fnRecorder) DoReturnResults(fn MoqCopyBuffer_genType_doReturnFn) *MoqCopyBuffer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Written int64
			Err     error
		}
		Sequence   uint32
		DoFn       MoqCopyBuffer_genType_doFn
		DoReturnFn MoqCopyBuffer_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqCopyBuffer_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqCopyBuffer_genType_resultsByParams
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
		results = &MoqCopyBuffer_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqCopyBuffer_genType_paramsKey]*MoqCopyBuffer_genType_results{},
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
		r.Results = &MoqCopyBuffer_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqCopyBuffer_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqCopyBuffer_genType_fnRecorder {
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
					Written int64
					Err     error
				}
				Sequence   uint32
				DoFn       MoqCopyBuffer_genType_doFn
				DoReturnFn MoqCopyBuffer_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqCopyBuffer_genType) PrettyParams(params MoqCopyBuffer_genType_params) string {
	return fmt.Sprintf("CopyBuffer_genType(%#v, %#v, %#v)", params.Dst, params.Src, params.Buf)
}

func (m *MoqCopyBuffer_genType) ParamsKey(params MoqCopyBuffer_genType_params, anyParams uint64) MoqCopyBuffer_genType_paramsKey {
	m.Scene.T.Helper()
	var dstUsed io.Writer
	var dstUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Dst == moq.ParamIndexByValue {
			dstUsed = params.Dst
		} else {
			dstUsedHash = hash.DeepHash(params.Dst)
		}
	}
	var srcUsed io.Reader
	var srcUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Src == moq.ParamIndexByValue {
			srcUsed = params.Src
		} else {
			srcUsedHash = hash.DeepHash(params.Src)
		}
	}
	var bufUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Buf == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The buf parameter can't be indexed by value")
		}
		bufUsedHash = hash.DeepHash(params.Buf)
	}
	return MoqCopyBuffer_genType_paramsKey{
		Params: struct {
			Dst io.Writer
			Src io.Reader
		}{
			Dst: dstUsed,
			Src: srcUsed,
		},
		Hashes: struct {
			Dst hash.Hash
			Src hash.Hash
			Buf hash.Hash
		}{
			Dst: dstUsedHash,
			Src: srcUsedHash,
			Buf: bufUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqCopyBuffer_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqCopyBuffer_genType) AssertExpectationsMet() {
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
