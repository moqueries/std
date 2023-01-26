// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package cipher

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that cipher.StreamWriter_genType is
// mocked completely
var _ StreamWriter_genType = (*MoqStreamWriter_genType_mock)(nil)

// StreamWriter_genType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type StreamWriter_genType interface {
	Write(src []byte) (n int, err error)
	Close() error
}

// MoqStreamWriter_genType holds the state of a moq of the StreamWriter_genType
// type
type MoqStreamWriter_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqStreamWriter_genType_mock

	ResultsByParams_Write []MoqStreamWriter_genType_Write_resultsByParams
	ResultsByParams_Close []MoqStreamWriter_genType_Close_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Write struct {
				Src moq.ParamIndexing
			}
			Close struct{}
		}
	}
}

// MoqStreamWriter_genType_mock isolates the mock interface of the
// StreamWriter_genType type
type MoqStreamWriter_genType_mock struct {
	Moq *MoqStreamWriter_genType
}

// MoqStreamWriter_genType_recorder isolates the recorder interface of the
// StreamWriter_genType type
type MoqStreamWriter_genType_recorder struct {
	Moq *MoqStreamWriter_genType
}

// MoqStreamWriter_genType_Write_params holds the params of the
// StreamWriter_genType type
type MoqStreamWriter_genType_Write_params struct{ Src []byte }

// MoqStreamWriter_genType_Write_paramsKey holds the map key params of the
// StreamWriter_genType type
type MoqStreamWriter_genType_Write_paramsKey struct {
	Params struct{}
	Hashes struct{ Src hash.Hash }
}

// MoqStreamWriter_genType_Write_resultsByParams contains the results for a
// given set of parameters for the StreamWriter_genType type
type MoqStreamWriter_genType_Write_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqStreamWriter_genType_Write_paramsKey]*MoqStreamWriter_genType_Write_results
}

// MoqStreamWriter_genType_Write_doFn defines the type of function needed when
// calling AndDo for the StreamWriter_genType type
type MoqStreamWriter_genType_Write_doFn func(src []byte)

// MoqStreamWriter_genType_Write_doReturnFn defines the type of function needed
// when calling DoReturnResults for the StreamWriter_genType type
type MoqStreamWriter_genType_Write_doReturnFn func(src []byte) (n int, err error)

// MoqStreamWriter_genType_Write_results holds the results of the
// StreamWriter_genType type
type MoqStreamWriter_genType_Write_results struct {
	Params  MoqStreamWriter_genType_Write_params
	Results []struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqStreamWriter_genType_Write_doFn
		DoReturnFn MoqStreamWriter_genType_Write_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqStreamWriter_genType_Write_fnRecorder routes recorded function calls to
// the MoqStreamWriter_genType moq
type MoqStreamWriter_genType_Write_fnRecorder struct {
	Params    MoqStreamWriter_genType_Write_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqStreamWriter_genType_Write_results
	Moq       *MoqStreamWriter_genType
}

// MoqStreamWriter_genType_Write_anyParams isolates the any params functions of
// the StreamWriter_genType type
type MoqStreamWriter_genType_Write_anyParams struct {
	Recorder *MoqStreamWriter_genType_Write_fnRecorder
}

// MoqStreamWriter_genType_Close_params holds the params of the
// StreamWriter_genType type
type MoqStreamWriter_genType_Close_params struct{}

// MoqStreamWriter_genType_Close_paramsKey holds the map key params of the
// StreamWriter_genType type
type MoqStreamWriter_genType_Close_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqStreamWriter_genType_Close_resultsByParams contains the results for a
// given set of parameters for the StreamWriter_genType type
type MoqStreamWriter_genType_Close_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqStreamWriter_genType_Close_paramsKey]*MoqStreamWriter_genType_Close_results
}

// MoqStreamWriter_genType_Close_doFn defines the type of function needed when
// calling AndDo for the StreamWriter_genType type
type MoqStreamWriter_genType_Close_doFn func()

// MoqStreamWriter_genType_Close_doReturnFn defines the type of function needed
// when calling DoReturnResults for the StreamWriter_genType type
type MoqStreamWriter_genType_Close_doReturnFn func() error

// MoqStreamWriter_genType_Close_results holds the results of the
// StreamWriter_genType type
type MoqStreamWriter_genType_Close_results struct {
	Params  MoqStreamWriter_genType_Close_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqStreamWriter_genType_Close_doFn
		DoReturnFn MoqStreamWriter_genType_Close_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqStreamWriter_genType_Close_fnRecorder routes recorded function calls to
// the MoqStreamWriter_genType moq
type MoqStreamWriter_genType_Close_fnRecorder struct {
	Params    MoqStreamWriter_genType_Close_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqStreamWriter_genType_Close_results
	Moq       *MoqStreamWriter_genType
}

// MoqStreamWriter_genType_Close_anyParams isolates the any params functions of
// the StreamWriter_genType type
type MoqStreamWriter_genType_Close_anyParams struct {
	Recorder *MoqStreamWriter_genType_Close_fnRecorder
}

// NewMoqStreamWriter_genType creates a new moq of the StreamWriter_genType
// type
func NewMoqStreamWriter_genType(scene *moq.Scene, config *moq.Config) *MoqStreamWriter_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqStreamWriter_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqStreamWriter_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Write struct {
					Src moq.ParamIndexing
				}
				Close struct{}
			}
		}{ParameterIndexing: struct {
			Write struct {
				Src moq.ParamIndexing
			}
			Close struct{}
		}{
			Write: struct {
				Src moq.ParamIndexing
			}{
				Src: moq.ParamIndexByHash,
			},
			Close: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the StreamWriter_genType type
func (m *MoqStreamWriter_genType) Mock() *MoqStreamWriter_genType_mock { return m.Moq }

func (m *MoqStreamWriter_genType_mock) Write(src []byte) (n int, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqStreamWriter_genType_Write_params{
		Src: src,
	}
	var results *MoqStreamWriter_genType_Write_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Write {
		paramsKey := m.Moq.ParamsKey_Write(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Write(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Write(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Write(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(src)
	}

	if result.Values != nil {
		n = result.Values.N
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		n, err = result.DoReturnFn(src)
	}
	return
}

func (m *MoqStreamWriter_genType_mock) Close() (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqStreamWriter_genType_Close_params{}
	var results *MoqStreamWriter_genType_Close_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Close {
		paramsKey := m.Moq.ParamsKey_Close(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Close(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Close(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Close(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn()
	}
	return
}

// OnCall returns the recorder implementation of the StreamWriter_genType type
func (m *MoqStreamWriter_genType) OnCall() *MoqStreamWriter_genType_recorder {
	return &MoqStreamWriter_genType_recorder{
		Moq: m,
	}
}

func (m *MoqStreamWriter_genType_recorder) Write(src []byte) *MoqStreamWriter_genType_Write_fnRecorder {
	return &MoqStreamWriter_genType_Write_fnRecorder{
		Params: MoqStreamWriter_genType_Write_params{
			Src: src,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqStreamWriter_genType_Write_fnRecorder) Any() *MoqStreamWriter_genType_Write_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Write(r.Params))
		return nil
	}
	return &MoqStreamWriter_genType_Write_anyParams{Recorder: r}
}

func (a *MoqStreamWriter_genType_Write_anyParams) Src() *MoqStreamWriter_genType_Write_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqStreamWriter_genType_Write_fnRecorder) Seq() *MoqStreamWriter_genType_Write_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Write(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqStreamWriter_genType_Write_fnRecorder) NoSeq() *MoqStreamWriter_genType_Write_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Write(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqStreamWriter_genType_Write_fnRecorder) ReturnResults(n int, err error) *MoqStreamWriter_genType_Write_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqStreamWriter_genType_Write_doFn
		DoReturnFn MoqStreamWriter_genType_Write_doReturnFn
	}{
		Values: &struct {
			N   int
			Err error
		}{
			N:   n,
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqStreamWriter_genType_Write_fnRecorder) AndDo(fn MoqStreamWriter_genType_Write_doFn) *MoqStreamWriter_genType_Write_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqStreamWriter_genType_Write_fnRecorder) DoReturnResults(fn MoqStreamWriter_genType_Write_doReturnFn) *MoqStreamWriter_genType_Write_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqStreamWriter_genType_Write_doFn
		DoReturnFn MoqStreamWriter_genType_Write_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqStreamWriter_genType_Write_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqStreamWriter_genType_Write_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Write {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqStreamWriter_genType_Write_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqStreamWriter_genType_Write_paramsKey]*MoqStreamWriter_genType_Write_results{},
		}
		r.Moq.ResultsByParams_Write = append(r.Moq.ResultsByParams_Write, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Write) {
			copy(r.Moq.ResultsByParams_Write[insertAt+1:], r.Moq.ResultsByParams_Write[insertAt:0])
			r.Moq.ResultsByParams_Write[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Write(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqStreamWriter_genType_Write_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqStreamWriter_genType_Write_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqStreamWriter_genType_Write_fnRecorder {
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
					N   int
					Err error
				}
				Sequence   uint32
				DoFn       MoqStreamWriter_genType_Write_doFn
				DoReturnFn MoqStreamWriter_genType_Write_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqStreamWriter_genType) PrettyParams_Write(params MoqStreamWriter_genType_Write_params) string {
	return fmt.Sprintf("Write(%#v)", params.Src)
}

func (m *MoqStreamWriter_genType) ParamsKey_Write(params MoqStreamWriter_genType_Write_params, anyParams uint64) MoqStreamWriter_genType_Write_paramsKey {
	m.Scene.T.Helper()
	var srcUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Write.Src == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The src parameter of the Write function can't be indexed by value")
		}
		srcUsedHash = hash.DeepHash(params.Src)
	}
	return MoqStreamWriter_genType_Write_paramsKey{
		Params: struct{}{},
		Hashes: struct{ Src hash.Hash }{
			Src: srcUsedHash,
		},
	}
}

func (m *MoqStreamWriter_genType_recorder) Close() *MoqStreamWriter_genType_Close_fnRecorder {
	return &MoqStreamWriter_genType_Close_fnRecorder{
		Params:   MoqStreamWriter_genType_Close_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqStreamWriter_genType_Close_fnRecorder) Any() *MoqStreamWriter_genType_Close_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	return &MoqStreamWriter_genType_Close_anyParams{Recorder: r}
}

func (r *MoqStreamWriter_genType_Close_fnRecorder) Seq() *MoqStreamWriter_genType_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqStreamWriter_genType_Close_fnRecorder) NoSeq() *MoqStreamWriter_genType_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqStreamWriter_genType_Close_fnRecorder) ReturnResults(result1 error) *MoqStreamWriter_genType_Close_fnRecorder {
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
		DoFn       MoqStreamWriter_genType_Close_doFn
		DoReturnFn MoqStreamWriter_genType_Close_doReturnFn
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

func (r *MoqStreamWriter_genType_Close_fnRecorder) AndDo(fn MoqStreamWriter_genType_Close_doFn) *MoqStreamWriter_genType_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqStreamWriter_genType_Close_fnRecorder) DoReturnResults(fn MoqStreamWriter_genType_Close_doReturnFn) *MoqStreamWriter_genType_Close_fnRecorder {
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
		DoFn       MoqStreamWriter_genType_Close_doFn
		DoReturnFn MoqStreamWriter_genType_Close_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqStreamWriter_genType_Close_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqStreamWriter_genType_Close_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Close {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqStreamWriter_genType_Close_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqStreamWriter_genType_Close_paramsKey]*MoqStreamWriter_genType_Close_results{},
		}
		r.Moq.ResultsByParams_Close = append(r.Moq.ResultsByParams_Close, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Close) {
			copy(r.Moq.ResultsByParams_Close[insertAt+1:], r.Moq.ResultsByParams_Close[insertAt:0])
			r.Moq.ResultsByParams_Close[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Close(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqStreamWriter_genType_Close_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqStreamWriter_genType_Close_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqStreamWriter_genType_Close_fnRecorder {
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
				DoFn       MoqStreamWriter_genType_Close_doFn
				DoReturnFn MoqStreamWriter_genType_Close_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqStreamWriter_genType) PrettyParams_Close(params MoqStreamWriter_genType_Close_params) string {
	return fmt.Sprintf("Close()")
}

func (m *MoqStreamWriter_genType) ParamsKey_Close(params MoqStreamWriter_genType_Close_params, anyParams uint64) MoqStreamWriter_genType_Close_paramsKey {
	m.Scene.T.Helper()
	return MoqStreamWriter_genType_Close_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqStreamWriter_genType) Reset() {
	m.ResultsByParams_Write = nil
	m.ResultsByParams_Close = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqStreamWriter_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Write {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Write(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Close {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Close(results.Params))
			}
		}
	}
}