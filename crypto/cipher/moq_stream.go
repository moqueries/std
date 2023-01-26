// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package cipher

import (
	"crypto/cipher"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that cipher.Stream is mocked completely
var _ cipher.Stream = (*MoqStream_mock)(nil)

// MoqStream holds the state of a moq of the Stream type
type MoqStream struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqStream_mock

	ResultsByParams_XORKeyStream []MoqStream_XORKeyStream_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			XORKeyStream struct {
				Dst moq.ParamIndexing
				Src moq.ParamIndexing
			}
		}
	}
	// MoqStream_mock isolates the mock interface of the Stream type
}

type MoqStream_mock struct {
	Moq *MoqStream
}

// MoqStream_recorder isolates the recorder interface of the Stream type
type MoqStream_recorder struct {
	Moq *MoqStream
}

// MoqStream_XORKeyStream_params holds the params of the Stream type
type MoqStream_XORKeyStream_params struct{ Dst, Src []byte }

// MoqStream_XORKeyStream_paramsKey holds the map key params of the Stream type
type MoqStream_XORKeyStream_paramsKey struct {
	Params struct{}
	Hashes struct{ Dst, Src hash.Hash }
}

// MoqStream_XORKeyStream_resultsByParams contains the results for a given set
// of parameters for the Stream type
type MoqStream_XORKeyStream_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqStream_XORKeyStream_paramsKey]*MoqStream_XORKeyStream_results
}

// MoqStream_XORKeyStream_doFn defines the type of function needed when calling
// AndDo for the Stream type
type MoqStream_XORKeyStream_doFn func(dst, src []byte)

// MoqStream_XORKeyStream_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Stream type
type MoqStream_XORKeyStream_doReturnFn func(dst, src []byte)

// MoqStream_XORKeyStream_results holds the results of the Stream type
type MoqStream_XORKeyStream_results struct {
	Params  MoqStream_XORKeyStream_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqStream_XORKeyStream_doFn
		DoReturnFn MoqStream_XORKeyStream_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqStream_XORKeyStream_fnRecorder routes recorded function calls to the
// MoqStream moq
type MoqStream_XORKeyStream_fnRecorder struct {
	Params    MoqStream_XORKeyStream_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqStream_XORKeyStream_results
	Moq       *MoqStream
}

// MoqStream_XORKeyStream_anyParams isolates the any params functions of the
// Stream type
type MoqStream_XORKeyStream_anyParams struct {
	Recorder *MoqStream_XORKeyStream_fnRecorder
}

// NewMoqStream creates a new moq of the Stream type
func NewMoqStream(scene *moq.Scene, config *moq.Config) *MoqStream {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqStream{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqStream_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				XORKeyStream struct {
					Dst moq.ParamIndexing
					Src moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			XORKeyStream struct {
				Dst moq.ParamIndexing
				Src moq.ParamIndexing
			}
		}{
			XORKeyStream: struct {
				Dst moq.ParamIndexing
				Src moq.ParamIndexing
			}{
				Dst: moq.ParamIndexByHash,
				Src: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Stream type
func (m *MoqStream) Mock() *MoqStream_mock { return m.Moq }

func (m *MoqStream_mock) XORKeyStream(dst, src []byte) {
	m.Moq.Scene.T.Helper()
	params := MoqStream_XORKeyStream_params{
		Dst: dst,
		Src: src,
	}
	var results *MoqStream_XORKeyStream_results
	for _, resultsByParams := range m.Moq.ResultsByParams_XORKeyStream {
		paramsKey := m.Moq.ParamsKey_XORKeyStream(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_XORKeyStream(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_XORKeyStream(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_XORKeyStream(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(dst, src)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(dst, src)
	}
	return
}

// OnCall returns the recorder implementation of the Stream type
func (m *MoqStream) OnCall() *MoqStream_recorder {
	return &MoqStream_recorder{
		Moq: m,
	}
}

func (m *MoqStream_recorder) XORKeyStream(dst, src []byte) *MoqStream_XORKeyStream_fnRecorder {
	return &MoqStream_XORKeyStream_fnRecorder{
		Params: MoqStream_XORKeyStream_params{
			Dst: dst,
			Src: src,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqStream_XORKeyStream_fnRecorder) Any() *MoqStream_XORKeyStream_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_XORKeyStream(r.Params))
		return nil
	}
	return &MoqStream_XORKeyStream_anyParams{Recorder: r}
}

func (a *MoqStream_XORKeyStream_anyParams) Dst() *MoqStream_XORKeyStream_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqStream_XORKeyStream_anyParams) Src() *MoqStream_XORKeyStream_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqStream_XORKeyStream_fnRecorder) Seq() *MoqStream_XORKeyStream_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_XORKeyStream(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqStream_XORKeyStream_fnRecorder) NoSeq() *MoqStream_XORKeyStream_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_XORKeyStream(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqStream_XORKeyStream_fnRecorder) ReturnResults() *MoqStream_XORKeyStream_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqStream_XORKeyStream_doFn
		DoReturnFn MoqStream_XORKeyStream_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqStream_XORKeyStream_fnRecorder) AndDo(fn MoqStream_XORKeyStream_doFn) *MoqStream_XORKeyStream_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqStream_XORKeyStream_fnRecorder) DoReturnResults(fn MoqStream_XORKeyStream_doReturnFn) *MoqStream_XORKeyStream_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqStream_XORKeyStream_doFn
		DoReturnFn MoqStream_XORKeyStream_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqStream_XORKeyStream_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqStream_XORKeyStream_resultsByParams
	for n, res := range r.Moq.ResultsByParams_XORKeyStream {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqStream_XORKeyStream_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqStream_XORKeyStream_paramsKey]*MoqStream_XORKeyStream_results{},
		}
		r.Moq.ResultsByParams_XORKeyStream = append(r.Moq.ResultsByParams_XORKeyStream, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_XORKeyStream) {
			copy(r.Moq.ResultsByParams_XORKeyStream[insertAt+1:], r.Moq.ResultsByParams_XORKeyStream[insertAt:0])
			r.Moq.ResultsByParams_XORKeyStream[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_XORKeyStream(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqStream_XORKeyStream_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqStream_XORKeyStream_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqStream_XORKeyStream_fnRecorder {
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
				DoFn       MoqStream_XORKeyStream_doFn
				DoReturnFn MoqStream_XORKeyStream_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqStream) PrettyParams_XORKeyStream(params MoqStream_XORKeyStream_params) string {
	return fmt.Sprintf("XORKeyStream(%#v, %#v)", params.Dst, params.Src)
}

func (m *MoqStream) ParamsKey_XORKeyStream(params MoqStream_XORKeyStream_params, anyParams uint64) MoqStream_XORKeyStream_paramsKey {
	m.Scene.T.Helper()
	var dstUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.XORKeyStream.Dst == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The dst parameter of the XORKeyStream function can't be indexed by value")
		}
		dstUsedHash = hash.DeepHash(params.Dst)
	}
	var srcUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.XORKeyStream.Src == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The src parameter of the XORKeyStream function can't be indexed by value")
		}
		srcUsedHash = hash.DeepHash(params.Src)
	}
	return MoqStream_XORKeyStream_paramsKey{
		Params: struct{}{},
		Hashes: struct{ Dst, Src hash.Hash }{
			Dst: dstUsedHash,
			Src: srcUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqStream) Reset() { m.ResultsByParams_XORKeyStream = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqStream) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_XORKeyStream {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_XORKeyStream(results.Params))
			}
		}
	}
}