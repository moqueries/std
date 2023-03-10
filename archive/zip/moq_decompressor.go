// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// MoqDecompressor holds the state of a moq of the Decompressor type
type MoqDecompressor struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqDecompressor_mock

	ResultsByParams []MoqDecompressor_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Param1 moq.ParamIndexing
		}
	}
}

// MoqDecompressor_mock isolates the mock interface of the Decompressor type
type MoqDecompressor_mock struct {
	Moq *MoqDecompressor
}

// MoqDecompressor_params holds the params of the Decompressor type
type MoqDecompressor_params struct{ Param1 io.Reader }

// MoqDecompressor_paramsKey holds the map key params of the Decompressor type
type MoqDecompressor_paramsKey struct {
	Params struct{ Param1 io.Reader }
	Hashes struct{ Param1 hash.Hash }
}

// MoqDecompressor_resultsByParams contains the results for a given set of
// parameters for the Decompressor type
type MoqDecompressor_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqDecompressor_paramsKey]*MoqDecompressor_results
}

// MoqDecompressor_doFn defines the type of function needed when calling AndDo
// for the Decompressor type
type MoqDecompressor_doFn func(r io.Reader)

// MoqDecompressor_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Decompressor type
type MoqDecompressor_doReturnFn func(r io.Reader) io.ReadCloser

// MoqDecompressor_results holds the results of the Decompressor type
type MoqDecompressor_results struct {
	Params  MoqDecompressor_params
	Results []struct {
		Values *struct {
			Result1 io.ReadCloser
		}
		Sequence   uint32
		DoFn       MoqDecompressor_doFn
		DoReturnFn MoqDecompressor_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqDecompressor_fnRecorder routes recorded function calls to the
// MoqDecompressor moq
type MoqDecompressor_fnRecorder struct {
	Params    MoqDecompressor_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqDecompressor_results
	Moq       *MoqDecompressor
}

// MoqDecompressor_anyParams isolates the any params functions of the
// Decompressor type
type MoqDecompressor_anyParams struct {
	Recorder *MoqDecompressor_fnRecorder
}

// NewMoqDecompressor creates a new moq of the Decompressor type
func NewMoqDecompressor(scene *moq.Scene, config *moq.Config) *MoqDecompressor {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqDecompressor{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqDecompressor_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Param1 moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Param1 moq.ParamIndexing
		}{
			Param1: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Decompressor type
func (m *MoqDecompressor) Mock() zip.Decompressor {
	return func(param1 io.Reader) io.ReadCloser {
		m.Scene.T.Helper()
		moq := &MoqDecompressor_mock{Moq: m}
		return moq.Fn(param1)
	}
}

func (m *MoqDecompressor_mock) Fn(param1 io.Reader) (result1 io.ReadCloser) {
	m.Moq.Scene.T.Helper()
	params := MoqDecompressor_params{
		Param1: param1,
	}
	var results *MoqDecompressor_results
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
		result.DoFn(param1)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(param1)
	}
	return
}

func (m *MoqDecompressor) OnCall(param1 io.Reader) *MoqDecompressor_fnRecorder {
	return &MoqDecompressor_fnRecorder{
		Params: MoqDecompressor_params{
			Param1: param1,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqDecompressor_fnRecorder) Any() *MoqDecompressor_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqDecompressor_anyParams{Recorder: r}
}

func (a *MoqDecompressor_anyParams) Param1() *MoqDecompressor_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqDecompressor_fnRecorder) Seq() *MoqDecompressor_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqDecompressor_fnRecorder) NoSeq() *MoqDecompressor_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqDecompressor_fnRecorder) ReturnResults(result1 io.ReadCloser) *MoqDecompressor_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 io.ReadCloser
		}
		Sequence   uint32
		DoFn       MoqDecompressor_doFn
		DoReturnFn MoqDecompressor_doReturnFn
	}{
		Values: &struct {
			Result1 io.ReadCloser
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqDecompressor_fnRecorder) AndDo(fn MoqDecompressor_doFn) *MoqDecompressor_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqDecompressor_fnRecorder) DoReturnResults(fn MoqDecompressor_doReturnFn) *MoqDecompressor_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 io.ReadCloser
		}
		Sequence   uint32
		DoFn       MoqDecompressor_doFn
		DoReturnFn MoqDecompressor_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqDecompressor_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqDecompressor_resultsByParams
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
		results = &MoqDecompressor_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqDecompressor_paramsKey]*MoqDecompressor_results{},
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
		r.Results = &MoqDecompressor_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqDecompressor_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqDecompressor_fnRecorder {
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
					Result1 io.ReadCloser
				}
				Sequence   uint32
				DoFn       MoqDecompressor_doFn
				DoReturnFn MoqDecompressor_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqDecompressor) PrettyParams(params MoqDecompressor_params) string {
	return fmt.Sprintf("Decompressor(%#v)", params.Param1)
}

func (m *MoqDecompressor) ParamsKey(params MoqDecompressor_params, anyParams uint64) MoqDecompressor_paramsKey {
	m.Scene.T.Helper()
	var param1Used io.Reader
	var param1UsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Param1 == moq.ParamIndexByValue {
			param1Used = params.Param1
		} else {
			param1UsedHash = hash.DeepHash(params.Param1)
		}
	}
	return MoqDecompressor_paramsKey{
		Params: struct{ Param1 io.Reader }{
			Param1: param1Used,
		},
		Hashes: struct{ Param1 hash.Hash }{
			Param1: param1UsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqDecompressor) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqDecompressor) AssertExpectationsMet() {
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
