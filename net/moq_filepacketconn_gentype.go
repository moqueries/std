// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package net

import (
	"fmt"
	"math/bits"
	"net"
	"os"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// FilePacketConn_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type FilePacketConn_genType func(f *os.File) (c net.PacketConn, err error)

// MoqFilePacketConn_genType holds the state of a moq of the
// FilePacketConn_genType type
type MoqFilePacketConn_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqFilePacketConn_genType_mock

	ResultsByParams []MoqFilePacketConn_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			F moq.ParamIndexing
		}
	}
}

// MoqFilePacketConn_genType_mock isolates the mock interface of the
// FilePacketConn_genType type
type MoqFilePacketConn_genType_mock struct {
	Moq *MoqFilePacketConn_genType
}

// MoqFilePacketConn_genType_params holds the params of the
// FilePacketConn_genType type
type MoqFilePacketConn_genType_params struct{ F *os.File }

// MoqFilePacketConn_genType_paramsKey holds the map key params of the
// FilePacketConn_genType type
type MoqFilePacketConn_genType_paramsKey struct {
	Params struct{ F *os.File }
	Hashes struct{ F hash.Hash }
}

// MoqFilePacketConn_genType_resultsByParams contains the results for a given
// set of parameters for the FilePacketConn_genType type
type MoqFilePacketConn_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFilePacketConn_genType_paramsKey]*MoqFilePacketConn_genType_results
}

// MoqFilePacketConn_genType_doFn defines the type of function needed when
// calling AndDo for the FilePacketConn_genType type
type MoqFilePacketConn_genType_doFn func(f *os.File)

// MoqFilePacketConn_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the FilePacketConn_genType type
type MoqFilePacketConn_genType_doReturnFn func(f *os.File) (c net.PacketConn, err error)

// MoqFilePacketConn_genType_results holds the results of the
// FilePacketConn_genType type
type MoqFilePacketConn_genType_results struct {
	Params  MoqFilePacketConn_genType_params
	Results []struct {
		Values *struct {
			C   net.PacketConn
			Err error
		}
		Sequence   uint32
		DoFn       MoqFilePacketConn_genType_doFn
		DoReturnFn MoqFilePacketConn_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFilePacketConn_genType_fnRecorder routes recorded function calls to the
// MoqFilePacketConn_genType moq
type MoqFilePacketConn_genType_fnRecorder struct {
	Params    MoqFilePacketConn_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFilePacketConn_genType_results
	Moq       *MoqFilePacketConn_genType
}

// MoqFilePacketConn_genType_anyParams isolates the any params functions of the
// FilePacketConn_genType type
type MoqFilePacketConn_genType_anyParams struct {
	Recorder *MoqFilePacketConn_genType_fnRecorder
}

// NewMoqFilePacketConn_genType creates a new moq of the FilePacketConn_genType
// type
func NewMoqFilePacketConn_genType(scene *moq.Scene, config *moq.Config) *MoqFilePacketConn_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqFilePacketConn_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqFilePacketConn_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				F moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			F moq.ParamIndexing
		}{
			F: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the FilePacketConn_genType type
func (m *MoqFilePacketConn_genType) Mock() FilePacketConn_genType {
	return func(f *os.File) (_ net.PacketConn, _ error) {
		m.Scene.T.Helper()
		moq := &MoqFilePacketConn_genType_mock{Moq: m}
		return moq.Fn(f)
	}
}

func (m *MoqFilePacketConn_genType_mock) Fn(f *os.File) (c net.PacketConn, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqFilePacketConn_genType_params{
		F: f,
	}
	var results *MoqFilePacketConn_genType_results
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
		result.DoFn(f)
	}

	if result.Values != nil {
		c = result.Values.C
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		c, err = result.DoReturnFn(f)
	}
	return
}

func (m *MoqFilePacketConn_genType) OnCall(f *os.File) *MoqFilePacketConn_genType_fnRecorder {
	return &MoqFilePacketConn_genType_fnRecorder{
		Params: MoqFilePacketConn_genType_params{
			F: f,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqFilePacketConn_genType_fnRecorder) Any() *MoqFilePacketConn_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqFilePacketConn_genType_anyParams{Recorder: r}
}

func (a *MoqFilePacketConn_genType_anyParams) F() *MoqFilePacketConn_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqFilePacketConn_genType_fnRecorder) Seq() *MoqFilePacketConn_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFilePacketConn_genType_fnRecorder) NoSeq() *MoqFilePacketConn_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFilePacketConn_genType_fnRecorder) ReturnResults(c net.PacketConn, err error) *MoqFilePacketConn_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			C   net.PacketConn
			Err error
		}
		Sequence   uint32
		DoFn       MoqFilePacketConn_genType_doFn
		DoReturnFn MoqFilePacketConn_genType_doReturnFn
	}{
		Values: &struct {
			C   net.PacketConn
			Err error
		}{
			C:   c,
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqFilePacketConn_genType_fnRecorder) AndDo(fn MoqFilePacketConn_genType_doFn) *MoqFilePacketConn_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFilePacketConn_genType_fnRecorder) DoReturnResults(fn MoqFilePacketConn_genType_doReturnFn) *MoqFilePacketConn_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			C   net.PacketConn
			Err error
		}
		Sequence   uint32
		DoFn       MoqFilePacketConn_genType_doFn
		DoReturnFn MoqFilePacketConn_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFilePacketConn_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFilePacketConn_genType_resultsByParams
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
		results = &MoqFilePacketConn_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFilePacketConn_genType_paramsKey]*MoqFilePacketConn_genType_results{},
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
		r.Results = &MoqFilePacketConn_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFilePacketConn_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFilePacketConn_genType_fnRecorder {
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
					C   net.PacketConn
					Err error
				}
				Sequence   uint32
				DoFn       MoqFilePacketConn_genType_doFn
				DoReturnFn MoqFilePacketConn_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFilePacketConn_genType) PrettyParams(params MoqFilePacketConn_genType_params) string {
	return fmt.Sprintf("FilePacketConn_genType(%#v)", params.F)
}

func (m *MoqFilePacketConn_genType) ParamsKey(params MoqFilePacketConn_genType_params, anyParams uint64) MoqFilePacketConn_genType_paramsKey {
	m.Scene.T.Helper()
	var fUsed *os.File
	var fUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.F == moq.ParamIndexByValue {
			fUsed = params.F
		} else {
			fUsedHash = hash.DeepHash(params.F)
		}
	}
	return MoqFilePacketConn_genType_paramsKey{
		Params: struct{ F *os.File }{
			F: fUsed,
		},
		Hashes: struct{ F hash.Hash }{
			F: fUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqFilePacketConn_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqFilePacketConn_genType) AssertExpectationsMet() {
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
