// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package scanner

import (
	"fmt"
	"go/scanner"
	"go/token"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that scanner.Scanner_starGenType is
// mocked completely
var _ Scanner_starGenType = (*MoqScanner_starGenType_mock)(nil)

// Scanner_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type Scanner_starGenType interface {
	Init(file *token.File, src []byte, err scanner.ErrorHandler, mode scanner.Mode)
	Scan() (pos token.Pos, tok token.Token, lit string)
}

// MoqScanner_starGenType holds the state of a moq of the Scanner_starGenType
// type
type MoqScanner_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqScanner_starGenType_mock

	ResultsByParams_Init []MoqScanner_starGenType_Init_resultsByParams
	ResultsByParams_Scan []MoqScanner_starGenType_Scan_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Init struct {
				File moq.ParamIndexing
				Src  moq.ParamIndexing
				Err  moq.ParamIndexing
				Mode moq.ParamIndexing
			}
			Scan struct{}
		}
	}
}

// MoqScanner_starGenType_mock isolates the mock interface of the
// Scanner_starGenType type
type MoqScanner_starGenType_mock struct {
	Moq *MoqScanner_starGenType
}

// MoqScanner_starGenType_recorder isolates the recorder interface of the
// Scanner_starGenType type
type MoqScanner_starGenType_recorder struct {
	Moq *MoqScanner_starGenType
}

// MoqScanner_starGenType_Init_params holds the params of the
// Scanner_starGenType type
type MoqScanner_starGenType_Init_params struct {
	File *token.File
	Src  []byte
	Err  scanner.ErrorHandler
	Mode scanner.Mode
}

// MoqScanner_starGenType_Init_paramsKey holds the map key params of the
// Scanner_starGenType type
type MoqScanner_starGenType_Init_paramsKey struct {
	Params struct {
		File *token.File
		Mode scanner.Mode
	}
	Hashes struct {
		File hash.Hash
		Src  hash.Hash
		Err  hash.Hash
		Mode hash.Hash
	}
}

// MoqScanner_starGenType_Init_resultsByParams contains the results for a given
// set of parameters for the Scanner_starGenType type
type MoqScanner_starGenType_Init_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqScanner_starGenType_Init_paramsKey]*MoqScanner_starGenType_Init_results
}

// MoqScanner_starGenType_Init_doFn defines the type of function needed when
// calling AndDo for the Scanner_starGenType type
type MoqScanner_starGenType_Init_doFn func(file *token.File, src []byte, err scanner.ErrorHandler, mode scanner.Mode)

// MoqScanner_starGenType_Init_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Scanner_starGenType type
type MoqScanner_starGenType_Init_doReturnFn func(file *token.File, src []byte, err scanner.ErrorHandler, mode scanner.Mode)

// MoqScanner_starGenType_Init_results holds the results of the
// Scanner_starGenType type
type MoqScanner_starGenType_Init_results struct {
	Params  MoqScanner_starGenType_Init_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqScanner_starGenType_Init_doFn
		DoReturnFn MoqScanner_starGenType_Init_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqScanner_starGenType_Init_fnRecorder routes recorded function calls to the
// MoqScanner_starGenType moq
type MoqScanner_starGenType_Init_fnRecorder struct {
	Params    MoqScanner_starGenType_Init_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqScanner_starGenType_Init_results
	Moq       *MoqScanner_starGenType
}

// MoqScanner_starGenType_Init_anyParams isolates the any params functions of
// the Scanner_starGenType type
type MoqScanner_starGenType_Init_anyParams struct {
	Recorder *MoqScanner_starGenType_Init_fnRecorder
}

// MoqScanner_starGenType_Scan_params holds the params of the
// Scanner_starGenType type
type MoqScanner_starGenType_Scan_params struct{}

// MoqScanner_starGenType_Scan_paramsKey holds the map key params of the
// Scanner_starGenType type
type MoqScanner_starGenType_Scan_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqScanner_starGenType_Scan_resultsByParams contains the results for a given
// set of parameters for the Scanner_starGenType type
type MoqScanner_starGenType_Scan_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqScanner_starGenType_Scan_paramsKey]*MoqScanner_starGenType_Scan_results
}

// MoqScanner_starGenType_Scan_doFn defines the type of function needed when
// calling AndDo for the Scanner_starGenType type
type MoqScanner_starGenType_Scan_doFn func()

// MoqScanner_starGenType_Scan_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Scanner_starGenType type
type MoqScanner_starGenType_Scan_doReturnFn func() (pos token.Pos, tok token.Token, lit string)

// MoqScanner_starGenType_Scan_results holds the results of the
// Scanner_starGenType type
type MoqScanner_starGenType_Scan_results struct {
	Params  MoqScanner_starGenType_Scan_params
	Results []struct {
		Values *struct {
			Pos token.Pos
			Tok token.Token
			Lit string
		}
		Sequence   uint32
		DoFn       MoqScanner_starGenType_Scan_doFn
		DoReturnFn MoqScanner_starGenType_Scan_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqScanner_starGenType_Scan_fnRecorder routes recorded function calls to the
// MoqScanner_starGenType moq
type MoqScanner_starGenType_Scan_fnRecorder struct {
	Params    MoqScanner_starGenType_Scan_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqScanner_starGenType_Scan_results
	Moq       *MoqScanner_starGenType
}

// MoqScanner_starGenType_Scan_anyParams isolates the any params functions of
// the Scanner_starGenType type
type MoqScanner_starGenType_Scan_anyParams struct {
	Recorder *MoqScanner_starGenType_Scan_fnRecorder
}

// NewMoqScanner_starGenType creates a new moq of the Scanner_starGenType type
func NewMoqScanner_starGenType(scene *moq.Scene, config *moq.Config) *MoqScanner_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqScanner_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqScanner_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Init struct {
					File moq.ParamIndexing
					Src  moq.ParamIndexing
					Err  moq.ParamIndexing
					Mode moq.ParamIndexing
				}
				Scan struct{}
			}
		}{ParameterIndexing: struct {
			Init struct {
				File moq.ParamIndexing
				Src  moq.ParamIndexing
				Err  moq.ParamIndexing
				Mode moq.ParamIndexing
			}
			Scan struct{}
		}{
			Init: struct {
				File moq.ParamIndexing
				Src  moq.ParamIndexing
				Err  moq.ParamIndexing
				Mode moq.ParamIndexing
			}{
				File: moq.ParamIndexByHash,
				Src:  moq.ParamIndexByHash,
				Err:  moq.ParamIndexByHash,
				Mode: moq.ParamIndexByValue,
			},
			Scan: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Scanner_starGenType type
func (m *MoqScanner_starGenType) Mock() *MoqScanner_starGenType_mock { return m.Moq }

func (m *MoqScanner_starGenType_mock) Init(file *token.File, src []byte, err scanner.ErrorHandler, mode scanner.Mode) {
	m.Moq.Scene.T.Helper()
	params := MoqScanner_starGenType_Init_params{
		File: file,
		Src:  src,
		Err:  err,
		Mode: mode,
	}
	var results *MoqScanner_starGenType_Init_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Init {
		paramsKey := m.Moq.ParamsKey_Init(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Init(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Init(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Init(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(file, src, err, mode)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(file, src, err, mode)
	}
	return
}

func (m *MoqScanner_starGenType_mock) Scan() (pos token.Pos, tok token.Token, lit string) {
	m.Moq.Scene.T.Helper()
	params := MoqScanner_starGenType_Scan_params{}
	var results *MoqScanner_starGenType_Scan_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Scan {
		paramsKey := m.Moq.ParamsKey_Scan(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Scan(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Scan(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Scan(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		pos = result.Values.Pos
		tok = result.Values.Tok
		lit = result.Values.Lit
	}
	if result.DoReturnFn != nil {
		pos, tok, lit = result.DoReturnFn()
	}
	return
}

// OnCall returns the recorder implementation of the Scanner_starGenType type
func (m *MoqScanner_starGenType) OnCall() *MoqScanner_starGenType_recorder {
	return &MoqScanner_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqScanner_starGenType_recorder) Init(file *token.File, src []byte, err scanner.ErrorHandler, mode scanner.Mode) *MoqScanner_starGenType_Init_fnRecorder {
	return &MoqScanner_starGenType_Init_fnRecorder{
		Params: MoqScanner_starGenType_Init_params{
			File: file,
			Src:  src,
			Err:  err,
			Mode: mode,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqScanner_starGenType_Init_fnRecorder) Any() *MoqScanner_starGenType_Init_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Init(r.Params))
		return nil
	}
	return &MoqScanner_starGenType_Init_anyParams{Recorder: r}
}

func (a *MoqScanner_starGenType_Init_anyParams) File() *MoqScanner_starGenType_Init_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqScanner_starGenType_Init_anyParams) Src() *MoqScanner_starGenType_Init_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqScanner_starGenType_Init_anyParams) Err() *MoqScanner_starGenType_Init_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (a *MoqScanner_starGenType_Init_anyParams) Mode() *MoqScanner_starGenType_Init_fnRecorder {
	a.Recorder.AnyParams |= 1 << 3
	return a.Recorder
}

func (r *MoqScanner_starGenType_Init_fnRecorder) Seq() *MoqScanner_starGenType_Init_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Init(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqScanner_starGenType_Init_fnRecorder) NoSeq() *MoqScanner_starGenType_Init_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Init(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqScanner_starGenType_Init_fnRecorder) ReturnResults() *MoqScanner_starGenType_Init_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqScanner_starGenType_Init_doFn
		DoReturnFn MoqScanner_starGenType_Init_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqScanner_starGenType_Init_fnRecorder) AndDo(fn MoqScanner_starGenType_Init_doFn) *MoqScanner_starGenType_Init_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqScanner_starGenType_Init_fnRecorder) DoReturnResults(fn MoqScanner_starGenType_Init_doReturnFn) *MoqScanner_starGenType_Init_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqScanner_starGenType_Init_doFn
		DoReturnFn MoqScanner_starGenType_Init_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqScanner_starGenType_Init_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqScanner_starGenType_Init_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Init {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqScanner_starGenType_Init_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqScanner_starGenType_Init_paramsKey]*MoqScanner_starGenType_Init_results{},
		}
		r.Moq.ResultsByParams_Init = append(r.Moq.ResultsByParams_Init, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Init) {
			copy(r.Moq.ResultsByParams_Init[insertAt+1:], r.Moq.ResultsByParams_Init[insertAt:0])
			r.Moq.ResultsByParams_Init[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Init(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqScanner_starGenType_Init_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqScanner_starGenType_Init_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqScanner_starGenType_Init_fnRecorder {
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
				DoFn       MoqScanner_starGenType_Init_doFn
				DoReturnFn MoqScanner_starGenType_Init_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqScanner_starGenType) PrettyParams_Init(params MoqScanner_starGenType_Init_params) string {
	return fmt.Sprintf("Init(%#v, %#v, %#v, %#v)", params.File, params.Src, params.Err, params.Mode)
}

func (m *MoqScanner_starGenType) ParamsKey_Init(params MoqScanner_starGenType_Init_params, anyParams uint64) MoqScanner_starGenType_Init_paramsKey {
	m.Scene.T.Helper()
	var fileUsed *token.File
	var fileUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Init.File == moq.ParamIndexByValue {
			fileUsed = params.File
		} else {
			fileUsedHash = hash.DeepHash(params.File)
		}
	}
	var srcUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Init.Src == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The src parameter of the Init function can't be indexed by value")
		}
		srcUsedHash = hash.DeepHash(params.Src)
	}
	var errUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Init.Err == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The err parameter of the Init function can't be indexed by value")
		}
		errUsedHash = hash.DeepHash(params.Err)
	}
	var modeUsed scanner.Mode
	var modeUsedHash hash.Hash
	if anyParams&(1<<3) == 0 {
		if m.Runtime.ParameterIndexing.Init.Mode == moq.ParamIndexByValue {
			modeUsed = params.Mode
		} else {
			modeUsedHash = hash.DeepHash(params.Mode)
		}
	}
	return MoqScanner_starGenType_Init_paramsKey{
		Params: struct {
			File *token.File
			Mode scanner.Mode
		}{
			File: fileUsed,
			Mode: modeUsed,
		},
		Hashes: struct {
			File hash.Hash
			Src  hash.Hash
			Err  hash.Hash
			Mode hash.Hash
		}{
			File: fileUsedHash,
			Src:  srcUsedHash,
			Err:  errUsedHash,
			Mode: modeUsedHash,
		},
	}
}

func (m *MoqScanner_starGenType_recorder) Scan() *MoqScanner_starGenType_Scan_fnRecorder {
	return &MoqScanner_starGenType_Scan_fnRecorder{
		Params:   MoqScanner_starGenType_Scan_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqScanner_starGenType_Scan_fnRecorder) Any() *MoqScanner_starGenType_Scan_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Scan(r.Params))
		return nil
	}
	return &MoqScanner_starGenType_Scan_anyParams{Recorder: r}
}

func (r *MoqScanner_starGenType_Scan_fnRecorder) Seq() *MoqScanner_starGenType_Scan_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Scan(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqScanner_starGenType_Scan_fnRecorder) NoSeq() *MoqScanner_starGenType_Scan_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Scan(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqScanner_starGenType_Scan_fnRecorder) ReturnResults(pos token.Pos, tok token.Token, lit string) *MoqScanner_starGenType_Scan_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Pos token.Pos
			Tok token.Token
			Lit string
		}
		Sequence   uint32
		DoFn       MoqScanner_starGenType_Scan_doFn
		DoReturnFn MoqScanner_starGenType_Scan_doReturnFn
	}{
		Values: &struct {
			Pos token.Pos
			Tok token.Token
			Lit string
		}{
			Pos: pos,
			Tok: tok,
			Lit: lit,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqScanner_starGenType_Scan_fnRecorder) AndDo(fn MoqScanner_starGenType_Scan_doFn) *MoqScanner_starGenType_Scan_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqScanner_starGenType_Scan_fnRecorder) DoReturnResults(fn MoqScanner_starGenType_Scan_doReturnFn) *MoqScanner_starGenType_Scan_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Pos token.Pos
			Tok token.Token
			Lit string
		}
		Sequence   uint32
		DoFn       MoqScanner_starGenType_Scan_doFn
		DoReturnFn MoqScanner_starGenType_Scan_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqScanner_starGenType_Scan_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqScanner_starGenType_Scan_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Scan {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqScanner_starGenType_Scan_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqScanner_starGenType_Scan_paramsKey]*MoqScanner_starGenType_Scan_results{},
		}
		r.Moq.ResultsByParams_Scan = append(r.Moq.ResultsByParams_Scan, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Scan) {
			copy(r.Moq.ResultsByParams_Scan[insertAt+1:], r.Moq.ResultsByParams_Scan[insertAt:0])
			r.Moq.ResultsByParams_Scan[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Scan(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqScanner_starGenType_Scan_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqScanner_starGenType_Scan_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqScanner_starGenType_Scan_fnRecorder {
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
					Pos token.Pos
					Tok token.Token
					Lit string
				}
				Sequence   uint32
				DoFn       MoqScanner_starGenType_Scan_doFn
				DoReturnFn MoqScanner_starGenType_Scan_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqScanner_starGenType) PrettyParams_Scan(params MoqScanner_starGenType_Scan_params) string {
	return fmt.Sprintf("Scan()")
}

func (m *MoqScanner_starGenType) ParamsKey_Scan(params MoqScanner_starGenType_Scan_params, anyParams uint64) MoqScanner_starGenType_Scan_paramsKey {
	m.Scene.T.Helper()
	return MoqScanner_starGenType_Scan_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqScanner_starGenType) Reset() { m.ResultsByParams_Init = nil; m.ResultsByParams_Scan = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqScanner_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Init {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Init(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Scan {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Scan(results.Params))
			}
		}
	}
}