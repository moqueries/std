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

// Copy_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Copy_genType func(dst io.Writer, src io.Reader) (written int64, err error)

// MoqCopy_genType holds the state of a moq of the Copy_genType type
type MoqCopy_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqCopy_genType_mock

	ResultsByParams []MoqCopy_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Dst moq.ParamIndexing
			Src moq.ParamIndexing
		}
	}
}

// MoqCopy_genType_mock isolates the mock interface of the Copy_genType type
type MoqCopy_genType_mock struct {
	Moq *MoqCopy_genType
}

// MoqCopy_genType_params holds the params of the Copy_genType type
type MoqCopy_genType_params struct {
	Dst io.Writer
	Src io.Reader
}

// MoqCopy_genType_paramsKey holds the map key params of the Copy_genType type
type MoqCopy_genType_paramsKey struct {
	Params struct {
		Dst io.Writer
		Src io.Reader
	}
	Hashes struct {
		Dst hash.Hash
		Src hash.Hash
	}
}

// MoqCopy_genType_resultsByParams contains the results for a given set of
// parameters for the Copy_genType type
type MoqCopy_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqCopy_genType_paramsKey]*MoqCopy_genType_results
}

// MoqCopy_genType_doFn defines the type of function needed when calling AndDo
// for the Copy_genType type
type MoqCopy_genType_doFn func(dst io.Writer, src io.Reader)

// MoqCopy_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Copy_genType type
type MoqCopy_genType_doReturnFn func(dst io.Writer, src io.Reader) (written int64, err error)

// MoqCopy_genType_results holds the results of the Copy_genType type
type MoqCopy_genType_results struct {
	Params  MoqCopy_genType_params
	Results []struct {
		Values *struct {
			Written int64
			Err     error
		}
		Sequence   uint32
		DoFn       MoqCopy_genType_doFn
		DoReturnFn MoqCopy_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqCopy_genType_fnRecorder routes recorded function calls to the
// MoqCopy_genType moq
type MoqCopy_genType_fnRecorder struct {
	Params    MoqCopy_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqCopy_genType_results
	Moq       *MoqCopy_genType
}

// MoqCopy_genType_anyParams isolates the any params functions of the
// Copy_genType type
type MoqCopy_genType_anyParams struct {
	Recorder *MoqCopy_genType_fnRecorder
}

// NewMoqCopy_genType creates a new moq of the Copy_genType type
func NewMoqCopy_genType(scene *moq.Scene, config *moq.Config) *MoqCopy_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqCopy_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqCopy_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Dst moq.ParamIndexing
				Src moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Dst moq.ParamIndexing
			Src moq.ParamIndexing
		}{
			Dst: moq.ParamIndexByHash,
			Src: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Copy_genType type
func (m *MoqCopy_genType) Mock() Copy_genType {
	return func(dst io.Writer, src io.Reader) (_ int64, _ error) {
		m.Scene.T.Helper()
		moq := &MoqCopy_genType_mock{Moq: m}
		return moq.Fn(dst, src)
	}
}

func (m *MoqCopy_genType_mock) Fn(dst io.Writer, src io.Reader) (written int64, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqCopy_genType_params{
		Dst: dst,
		Src: src,
	}
	var results *MoqCopy_genType_results
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
		result.DoFn(dst, src)
	}

	if result.Values != nil {
		written = result.Values.Written
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		written, err = result.DoReturnFn(dst, src)
	}
	return
}

func (m *MoqCopy_genType) OnCall(dst io.Writer, src io.Reader) *MoqCopy_genType_fnRecorder {
	return &MoqCopy_genType_fnRecorder{
		Params: MoqCopy_genType_params{
			Dst: dst,
			Src: src,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqCopy_genType_fnRecorder) Any() *MoqCopy_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqCopy_genType_anyParams{Recorder: r}
}

func (a *MoqCopy_genType_anyParams) Dst() *MoqCopy_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqCopy_genType_anyParams) Src() *MoqCopy_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqCopy_genType_fnRecorder) Seq() *MoqCopy_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqCopy_genType_fnRecorder) NoSeq() *MoqCopy_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqCopy_genType_fnRecorder) ReturnResults(written int64, err error) *MoqCopy_genType_fnRecorder {
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
		DoFn       MoqCopy_genType_doFn
		DoReturnFn MoqCopy_genType_doReturnFn
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

func (r *MoqCopy_genType_fnRecorder) AndDo(fn MoqCopy_genType_doFn) *MoqCopy_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqCopy_genType_fnRecorder) DoReturnResults(fn MoqCopy_genType_doReturnFn) *MoqCopy_genType_fnRecorder {
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
		DoFn       MoqCopy_genType_doFn
		DoReturnFn MoqCopy_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqCopy_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqCopy_genType_resultsByParams
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
		results = &MoqCopy_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqCopy_genType_paramsKey]*MoqCopy_genType_results{},
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
		r.Results = &MoqCopy_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqCopy_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqCopy_genType_fnRecorder {
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
				DoFn       MoqCopy_genType_doFn
				DoReturnFn MoqCopy_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqCopy_genType) PrettyParams(params MoqCopy_genType_params) string {
	return fmt.Sprintf("Copy_genType(%#v, %#v)", params.Dst, params.Src)
}

func (m *MoqCopy_genType) ParamsKey(params MoqCopy_genType_params, anyParams uint64) MoqCopy_genType_paramsKey {
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
	return MoqCopy_genType_paramsKey{
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
		}{
			Dst: dstUsedHash,
			Src: srcUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqCopy_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqCopy_genType) AssertExpectationsMet() {
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
