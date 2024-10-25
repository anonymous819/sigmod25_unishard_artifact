//line :1
package log

import (
	"fmt"
	"io"
	stdlog "log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type EyiXh9_pYL int32

const (
	DEBUG EyiXh9_pYL = iota
	INFO
	WARNING
	ERROR
)

var sDh_lp6e5v6b = []string{
	DEBUG:   "DEBUG",
	INFO:    "INFO",
	WARNING: "WARNING",
	ERROR:   "ERROR",
}

type Ab8vTP byte

func (jSMHhmANvA *EyiXh9_pYL) Get() interface{} {
	return *jSMHhmANvA
}

func (jSMHhmANvA *EyiXh9_pYL) Set(cAF5bqkAnK string) error {
	ndmrjodO1P33 := INFO
	cAF5bqkAnK = /*line am7_nnm.go:1*/ strings.ToUpper(cAF5bqkAnK)
	for fMYWheYh3qci, cPPhfs6dVFA1 := range sDh_lp6e5v6b {
		if cPPhfs6dVFA1 == cAF5bqkAnK {
			ndmrjodO1P33 = /*line Y0Y1YBq2yRd.go:1*/ EyiXh9_pYL(fMYWheYh3qci)
		}
	}
	*jSMHhmANvA = ndmrjodO1P33
	return nil
}

func (jSMHhmANvA *EyiXh9_pYL) String() string {
	return sDh_lp6e5v6b[ /*line UaTBRz_F.go:1*/ int(*jSMHhmANvA)]
}

type dGVlx7zMoY struct {
	sync.Mutex

	vXwBFYlJM    *stdlog.Logger
	aHyvdWG      *stdlog.Logger
	dkBaQ_u      *stdlog.Logger
	cvsU_KTv     *stdlog.Logger
	bvGgkI0LJ_MX *stdlog.Logger
	p3bGPiJaFSz  *stdlog.Logger
	sKAJXeY      *stdlog.Logger
	cTrSsYtJqSLk *stdlog.Logger
	caDdgW       *stdlog.Logger
	p57_bcl8     *stdlog.Logger

	OKYtPP6TP EyiXh9_pYL
	jI29Ghqn_ string
}

var wS34PTi_MS dGVlx7zMoY

func VhjKf8sJCql9() {
	zZgt2OJXoJ3A := stdlog.Ldate | stdlog.Ltime | stdlog.Lmicroseconds | stdlog.Lshortfile
	_B_jZWlB7KF := /*line M479HGfFTU.go:1*/ fmt.Sprintf(func() /*line I9Hs1e.go:1*/ string {
		data := /*line I9Hs1e.go:1*/ make([]byte, 0, 18)
		i := 5
		decryptKey := 58
		for counter := 0; i != 1; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 3:
				i = 7
				data = /*line I9Hs1e.go:1*/ append(data, 120)
			case 2:
				for y := range data {
					data[y] = data[y] - /*line I9Hs1e.go:1*/ byte(decryptKey^y)
				}
				i = 1
			case 0:
				i = 4
				data = /*line I9Hs1e.go:1*/ append(data, "\x8a\x8e"...,
				)
			case 9:
				data = /*line I9Hs1e.go:1*/ append(data, 56)
				i = 3
			case 6:
				data = /*line I9Hs1e.go:1*/ append(data, 126)
				i = 10
			case 8:
				data = /*line I9Hs1e.go:1*/ append(data, 140)
				i = 6
			case 10:
				data = /*line I9Hs1e.go:1*/ append(data, "\x91|\x95K"...,
				)
				i = 0
			case 4:
				i = 2
				data = /*line I9Hs1e.go:1*/ append(data, 103)
			case 5:
				data = /*line I9Hs1e.go:1*/ append(data, "5\x84@"...,
				)
				i = 9
			case 7:
				i = 8
				data = /*line I9Hs1e.go:1*/ append(data, "C}x"...,
				)
			}
		}
		return /*line I9Hs1e.go:1*/ string(data)
	}(), /*line VKtml1.go:1*/ filepath.Base(os.Args[0]) /*line vn7jx0.go:1*/, os.Getpid())
	oKKPL9KO34_, oSrzLqmu0S := /*line vkFSPT.go:1*/ os.Create( /*line pxow4muHiXl.go:1*/ filepath.Join(wS34PTi_MS.jI29Ghqn_, _B_jZWlB7KF))
	if oSrzLqmu0S != nil {
		/*line tORfiuFv.go:1*/ stdlog.Fatal(oSrzLqmu0S)
	}
	lfn_G63RV := /*line HvC8omEg.go:1*/ fmt.Sprintf(func() /*line Z3ha3mg6At.go:1*/ string {
		data := [] /*line Z3ha3mg6At.go:1*/ byte("\x9f\xaa\xe4Wd,f`\xdarda\xbfa\xa9\v֞ަ֭")
		positions := [...]byte{6, 8, 17, 18, 15, 8, 5, 2, 5, 2, 20, 16, 3, 0, 1, 11, 21, 19, 14, 7, 1, 12, 2, 0, 14, 8}
		for i := 0; i < 26; i += 2 {
			localKey := /*line Z3ha3mg6At.go:1*/ byte(i) + /*line Z3ha3mg6At.go:1*/ byte(positions[i]^positions[i+1]) + 171
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line Z3ha3mg6At.go:1*/ string(data)
	}(), /*line HEEtV97Qs7m.go:1*/ filepath.Base(os.Args[0]) /*line JVkhE4i9Wj.go:1*/, os.Getpid())
	e81FhXHb, oSrzLqmu0S := /*line mDSxaBD7.go:1*/ os.Create( /*line GqdvPsYnJ.go:1*/ filepath.Join(wS34PTi_MS.jI29Ghqn_, lfn_G63RV))
	if oSrzLqmu0S != nil {
		/*line TAWfMpJFg2_.go:1*/ stdlog.Fatal(oSrzLqmu0S)
	}
	if9dfv := /*line DwYab4KKgnZc.go:1*/ fmt.Sprintf(func() /*line _AZ5Xx.go:1*/ string {
		key := [] /*line _AZ5Xx.go:1*/ byte("\xdf%S\xb8\xe6\x19Ǿ.")
		data := [] /*line _AZ5Xx.go:1*/ byte("\xfaV}\x9d\x827\xab\xd1I")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line _AZ5Xx.go:1*/ string(data)
	}(), /*line vl9k2OAlpEZt.go:1*/ filepath.Base(os.Args[0]) /*line RkWqRKRaeZ.go:1*/, os.Getpid())
	aLWMHdYk95gm, oSrzLqmu0S := /*line YLXCjlPJv.go:1*/ os.Create( /*line AhCuayuHsYac.go:1*/ filepath.Join(wS34PTi_MS.jI29Ghqn_, if9dfv))
	if oSrzLqmu0S != nil {
		/*line FIlZarmkNpQ.go:1*/ stdlog.Fatal(oSrzLqmu0S)
	}
	z2_IunfxA := /*line FQR4c9NVz.go:1*/ fmt.Sprintf(func() /*line YFZI3m_.go:1*/ string {
		seed := /*line YFZI3m_.go:1*/ byte(111)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line YFZI3m_.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line YFZI3m_.go:1*/ fnc(148)(118)(167)(69)(201)(92)(253)(239)(215)(191)(112)(207)(191)(107)(233)(199)(152)(49)(87)(168)(27)(116)(235)(206)
		return /*line YFZI3m_.go:1*/ string(data)
	}(), /*line cxQIYe2wdvBE.go:1*/ filepath.Base(os.Args[0]) /*line RRCCxbEZ.go:1*/, os.Getpid())
	xkGT29Kv, oSrzLqmu0S := /*line Hi_EM59koD.go:1*/ os.Create( /*line w6df8aN1by.go:1*/ filepath.Join(wS34PTi_MS.jI29Ghqn_, z2_IunfxA))
	if oSrzLqmu0S != nil {
		/*line iQfh2B.go:1*/ stdlog.Fatal(oSrzLqmu0S)
	}
	hMwkMiAW9yLg := /*line oEZmFgTfa.go:1*/ fmt.Sprintf(func() /*line hkZlUKq.go:1*/ string {
		key := [] /*line hkZlUKq.go:1*/ byte("\x8b\xbb\xef)GaCŭ\xe4\"\x93`\xa2z\x11\xe1\xb7\xd4\x03\x13\xe6\tH\xed\xc6")
		data := [] /*line hkZlUKq.go:1*/ byte("\xb0.\x1dN\xab\x8f\xa8=\x1dI\x94\xfc\xcd\a\xe8\x853\x1cGx\x7fZ7\xab`<")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line hkZlUKq.go:1*/ string(data)
	}(), /*line Cru3EQj4.go:1*/ filepath.Base(os.Args[0]) /*line cO4mekz4Xb.go:1*/, os.Getpid())
	tJexkT5a, oSrzLqmu0S := /*line HCNx9XaE.go:1*/ os.Create( /*line rbTakFWeg.go:1*/ filepath.Join(wS34PTi_MS.jI29Ghqn_, hMwkMiAW9yLg))
	if oSrzLqmu0S != nil {
		/*line pjMJvIrSC.go:1*/ stdlog.Fatal(oSrzLqmu0S)
	}

	wS34PTi_MS.vXwBFYlJM = /*line VCextt7jog.go:1*/ stdlog.New(oKKPL9KO34_, func() /*line dAGxoUouC.go:1*/ string {
		data := /*line dAGxoUouC.go:1*/ make([]byte, 0, 10)
		i := 2
		decryptKey := 191
		for counter := 0; i != 1; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 10:
				data = /*line dAGxoUouC.go:1*/ append(data, 159)
				i = 9
			case 8:
				i = 0
				data = /*line dAGxoUouC.go:1*/ append(data, 241)
			case 9:
				i = 5
				data = /*line dAGxoUouC.go:1*/ append(data, 137)
			case 7:
				i = 8
				data = /*line dAGxoUouC.go:1*/ append(data, 131)
			case 5:
				i = 7
				data = /*line dAGxoUouC.go:1*/ append(data, 152)
			case 2:
				data = /*line dAGxoUouC.go:1*/ append(data, 130)
				i = 6
			case 3:
				i = 4
				data = /*line dAGxoUouC.go:1*/ append(data, 159)
			case 0:
				for y := range data {
					data[y] = data[y] ^ /*line dAGxoUouC.go:1*/ byte(decryptKey^y)
				}
				i = 1
			case 4:
				i = 10
				data = /*line dAGxoUouC.go:1*/ append(data, 159)
			case 6:
				i = 3
				data = /*line dAGxoUouC.go:1*/ append(data, 159)
			}
		}
		return /*line dAGxoUouC.go:1*/ string(data)
	}(), zZgt2OJXoJ3A)
	wS34PTi_MS.aHyvdWG = /*line G2IO1T3O.go:1*/ stdlog.New(oKKPL9KO34_, func() /*line P262J4Ddl.go:1*/ string {
		data := /*line P262J4Ddl.go:1*/ make([]byte, 0, 9)
		i := 4
		decryptKey := 63
		for counter := 0; i != 0; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 5:
				i = 2
				data = /*line P262J4Ddl.go:1*/ append(data, 92)
			case 3:
				data = /*line P262J4Ddl.go:1*/ append(data, 57)
				i = 1
			case 7:
				data = /*line P262J4Ddl.go:1*/ append(data, 69)
				i = 3
			case 6:
				data = /*line P262J4Ddl.go:1*/ append(data, 83)
				i = 5
			case 4:
				data = /*line P262J4Ddl.go:1*/ append(data, 69)
				i = 8
			case 8:
				data = /*line P262J4Ddl.go:1*/ append(data, 88)
				i = 9
			case 2:
				i = 7
				data = /*line P262J4Ddl.go:1*/ append(data, 84)
			case 1:
				for y := range data {
					data[y] = data[y] ^ /*line P262J4Ddl.go:1*/ byte(decryptKey^y)
				}
				i = 0
			case 9:
				i = 6
				data = /*line P262J4Ddl.go:1*/ append(data, 85)
			}
		}
		return /*line P262J4Ddl.go:1*/ string(data)
	}(), zZgt2OJXoJ3A)
	wS34PTi_MS.dkBaQ_u = /*line yc8OhBu.go:1*/ stdlog.New(e81FhXHb, func() /*line A7BOY2JGL.go:1*/ string {
		seed := /*line A7BOY2JGL.go:1*/ byte(45)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line A7BOY2JGL.go:1*/ append(data, x^seed); seed += x; return fnc }
		/*line A7BOY2JGL.go:1*/ fnc(118)(224)(199)(15)(27)(33)(210)(58)(129)
		return /*line A7BOY2JGL.go:1*/ string(data)
	}(), zZgt2OJXoJ3A)
	wS34PTi_MS.cvsU_KTv = /*line UPZC91x.go:1*/ stdlog.New(e81FhXHb, func() /*line Ra8mvhNWqE5n.go:1*/ string {
		data := /*line Ra8mvhNWqE5n.go:1*/ make([]byte, 0, 9)
		i := 8
		decryptKey := 184
		for counter := 0; i != 7; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				i = 3
				data = /*line Ra8mvhNWqE5n.go:1*/ append(data, 253)
			case 6:
				for y := range data {
					data[y] = data[y] ^ /*line Ra8mvhNWqE5n.go:1*/ byte(decryptKey^y)
				}
				i = 7
			case 9:
				i = 5
				data = /*line Ra8mvhNWqE5n.go:1*/ append(data, 244)
			case 4:
				data = /*line Ra8mvhNWqE5n.go:1*/ append(data, 244)
				i = 1
			case 2:
				i = 0
				data = /*line Ra8mvhNWqE5n.go:1*/ append(data, 237)
			case 5:
				data = /*line Ra8mvhNWqE5n.go:1*/ append(data, 252)
				i = 2
			case 0:
				data = /*line Ra8mvhNWqE5n.go:1*/ append(data, 145)
				i = 6
			case 3:
				i = 9
				data = /*line Ra8mvhNWqE5n.go:1*/ append(data, 251)
			case 8:
				i = 4
				data = /*line Ra8mvhNWqE5n.go:1*/ append(data, 237)
			}
		}
		return /*line Ra8mvhNWqE5n.go:1*/ string(data)
	}(), zZgt2OJXoJ3A)
	wS34PTi_MS.bvGgkI0LJ_MX = /*line YZvWHtl957yb.go:1*/ stdlog.New(aLWMHdYk95gm, func() /*line kMLN_Ve.go:1*/ string {
		seed := /*line kMLN_Ve.go:1*/ byte(254)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line kMLN_Ve.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line kMLN_Ve.go:1*/ fnc(89)(155)(55)(107)(233)(196)(158)(255)
		return /*line kMLN_Ve.go:1*/ string(data)
	}(), zZgt2OJXoJ3A)
	wS34PTi_MS.p3bGPiJaFSz = /*line qXlafKOk_.go:1*/ stdlog.New(aLWMHdYk95gm, "[INFO] ", zZgt2OJXoJ3A)
	yviakf := /*line wNL8XPxi2.go:1*/ io.MultiWriter(aLWMHdYk95gm, os.Stderr)
	wS34PTi_MS.sKAJXeY = /*line nETfoQ.go:1*/ stdlog.New(yviakf, func() /*line aeqKOKdvtZI.go:1*/ string {
		data := [] /*line aeqKOKdvtZI.go:1*/ byte("XW3\xf7NINP\xfaD")
		positions := [...]byte{7, 2, 3, 8, 8, 2, 2, 8, 9, 0, 9, 7, 0, 2}
		for i := 0; i < 14; i += 2 {
			localKey := /*line aeqKOKdvtZI.go:1*/ byte(i) + /*line aeqKOKdvtZI.go:1*/ byte(positions[i]^positions[i+1]) + 155
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line aeqKOKdvtZI.go:1*/ string(data)
	}(), zZgt2OJXoJ3A)
	wS34PTi_MS.cTrSsYtJqSLk = /*line J4i64YCde.go:1*/ stdlog.New(yviakf, func() /*line JbAEIXVK.go:1*/ string {
		seed := /*line JbAEIXVK.go:1*/ byte(57)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line JbAEIXVK.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line JbAEIXVK.go:1*/ fnc(34)(234)(13)(0)(253)(3)(11)(195)
		return /*line JbAEIXVK.go:1*/ string(data)
	}(), zZgt2OJXoJ3A)
	wS34PTi_MS.caDdgW = /*line BnipQ8Of.go:1*/ stdlog.New(xkGT29Kv, func() /*line FjmMMtpa.go:1*/ string {
		seed := /*line FjmMMtpa.go:1*/ byte(45)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line FjmMMtpa.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line FjmMMtpa.go:1*/ fnc(136)(8)(5)(3)(23)(32)(45)(124)(249)(223)(209)(151)(56)(113)(215)(168)(106)
		return /*line FjmMMtpa.go:1*/ string(data)
	}(), zZgt2OJXoJ3A)
	wS34PTi_MS.p57_bcl8 = /*line X8iJ0SfaYv.go:1*/ stdlog.New(tJexkT5a, "", 0)
}

func FPmo4wB6Q(dY3VLLs ...interface{}) {
	if wS34PTi_MS.OKYtPP6TP == DEBUG {
		_ = /*line o4O4xtNo6A.go:1*/ wS34PTi_MS.vXwBFYlJM.Output(2 /*line MF3y3c4I.go:1*/, fmt.Sprint(dY3VLLs...))
	}
}

func Eo2LyO(zZgt2OJXoJ3A string, dY3VLLs ...interface{}) {
	if wS34PTi_MS.OKYtPP6TP == DEBUG {
		_ = /*line Ausxxe2CG.go:1*/ wS34PTi_MS.vXwBFYlJM.Output(2 /*line bHmqdl4.go:1*/, fmt.Sprintf(zZgt2OJXoJ3A, dY3VLLs...))
	}
}

func Pu1JbYyDl(dY3VLLs ...interface{}) {
	if wS34PTi_MS.OKYtPP6TP <= INFO {
		_ = /*line vafMwK.go:1*/ wS34PTi_MS.aHyvdWG.Output(2 /*line ECQBmoZGsQ.go:1*/, fmt.Sprint(dY3VLLs...))
	}
}

func W3Uo65mg_F(zZgt2OJXoJ3A string, dY3VLLs ...interface{}) {
	if wS34PTi_MS.OKYtPP6TP <= INFO {
		_ = /*line GDrZ7iUdb.go:1*/ wS34PTi_MS.aHyvdWG.Output(2 /*line o3aq5pLT.go:1*/, fmt.Sprintf(zZgt2OJXoJ3A, dY3VLLs...))
	}
}

func CDebpj(dY3VLLs ...interface{}) {
	if wS34PTi_MS.OKYtPP6TP == DEBUG {
		_ = /*line bBDEe4X.go:1*/ wS34PTi_MS.dkBaQ_u.Output(2 /*line yNZgtGo.go:1*/, fmt.Sprint(dY3VLLs...))
	}
}

func ViJSfja(zZgt2OJXoJ3A string, dY3VLLs ...interface{}) {
	if wS34PTi_MS.OKYtPP6TP == DEBUG {
		_ = /*line UFKohJyezTNS.go:1*/ wS34PTi_MS.dkBaQ_u.Output(2 /*line wdhM0HRsWuDV.go:1*/, fmt.Sprintf(zZgt2OJXoJ3A, dY3VLLs...))
	}
}

func NPdnXsO(dY3VLLs ...interface{}) {
	if wS34PTi_MS.OKYtPP6TP <= INFO {
		_ = /*line Lwpx21Hzd.go:1*/ wS34PTi_MS.cvsU_KTv.Output(2 /*line Okk3pVur.go:1*/, fmt.Sprint(dY3VLLs...))
	}
}

func XWmhQ9(zZgt2OJXoJ3A string, dY3VLLs ...interface{}) {
	if wS34PTi_MS.OKYtPP6TP <= INFO {
		_ = /*line GQKBQd.go:1*/ wS34PTi_MS.cvsU_KTv.Output(2 /*line Ieuwhl6xKb.go:1*/, fmt.Sprintf(zZgt2OJXoJ3A, dY3VLLs...))
	}
}

func Debug(dY3VLLs ...interface{}) {
	if wS34PTi_MS.OKYtPP6TP == DEBUG {
		_ = /*line anH_kjX.go:1*/ wS34PTi_MS.bvGgkI0LJ_MX.Output(2 /*line zqhZR0MCfbiF.go:1*/, fmt.Sprint(dY3VLLs...))
	}
}

func Debugf(zZgt2OJXoJ3A string, dY3VLLs ...interface{}) {
	if wS34PTi_MS.OKYtPP6TP == DEBUG {
		_ = /*line UB72gUN.go:1*/ wS34PTi_MS.bvGgkI0LJ_MX.Output(2 /*line oOJpZzVxFc0C.go:1*/, fmt.Sprintf(zZgt2OJXoJ3A, dY3VLLs...))
	}
}

func ZWwzboe(dY3VLLs ...interface{}) {
	if wS34PTi_MS.OKYtPP6TP <= INFO {
		_ = /*line wND7gHkoIo6.go:1*/ wS34PTi_MS.p3bGPiJaFSz.Output(2 /*line vWb47sbo2U.go:1*/, fmt.Sprint(dY3VLLs...))
	}
}

func Sa4H8el(zZgt2OJXoJ3A string, dY3VLLs ...interface{}) {
	if wS34PTi_MS.OKYtPP6TP <= INFO {
		_ = /*line HqL5n3.go:1*/ wS34PTi_MS.p3bGPiJaFSz.Output(2 /*line zDH3wm.go:1*/, fmt.Sprintf(zZgt2OJXoJ3A, dY3VLLs...))
	}
}

func B4DhS7wNt0JY(zZgt2OJXoJ3A string, dY3VLLs ...interface{}) {
	if wS34PTi_MS.OKYtPP6TP <= INFO {
		_ = /*line pncALhOfP.go:1*/ wS34PTi_MS.caDdgW.Output(2 /*line rWgFXLwxA.go:1*/, fmt.Sprintf(zZgt2OJXoJ3A, dY3VLLs...))
	}
}

func Bq9scv(dY3VLLs ...interface{}) {
	_ = /*line vUX8mI1XqD.go:1*/ wS34PTi_MS.p57_bcl8.Output(2 /*line v6Fmg24.go:1*/, fmt.Sprint(dY3VLLs...))
}

func Eal9ypd6VP(zZgt2OJXoJ3A string, dY3VLLs ...interface{}) {
	_ = /*line ntWIqc_aF9ec.go:1*/ wS34PTi_MS.p57_bcl8.Output(2 /*line eA_ZlkaFHloa.go:1*/, fmt.Sprintf(zZgt2OJXoJ3A, dY3VLLs...))
}

func ITw4kRdXNVBY(dY3VLLs ...interface{}) {
	if wS34PTi_MS.OKYtPP6TP <= WARNING {
		_ = /*line _5yfSHlZRLN.go:1*/ wS34PTi_MS.sKAJXeY.Output(2 /*line Tn0Gh6.go:1*/, fmt.Sprint(dY3VLLs...))
	}
}

func WPP4KjwN(zZgt2OJXoJ3A string, dY3VLLs ...interface{}) {
	if wS34PTi_MS.OKYtPP6TP <= WARNING {
		_ = /*line EJKjABoza.go:1*/ wS34PTi_MS.sKAJXeY.Output(2 /*line i1CfsjhZpBI.go:1*/, fmt.Sprintf(zZgt2OJXoJ3A, dY3VLLs...))
	}
}

func CIfHzNc72c(dY3VLLs ...interface{}) {
	_ = /*line mvFSvQ.go:1*/ wS34PTi_MS.cTrSsYtJqSLk.Output(2 /*line m6Igu0kx1.go:1*/, fmt.Sprint(dY3VLLs...))
}

func Jp980o6YjM(zZgt2OJXoJ3A string, dY3VLLs ...interface{}) {
	_ = /*line _MrcTGR.go:1*/ wS34PTi_MS.cTrSsYtJqSLk.Output(2 /*line SggFzs.go:1*/, fmt.Sprintf(zZgt2OJXoJ3A, dY3VLLs...))
}

func FB5S6xdVix(dY3VLLs ...interface{}) {
	_ = /*line KiRxfDNk2ayt.go:1*/ wS34PTi_MS.cTrSsYtJqSLk.Output(2 /*line J1EOG8avrVy.go:1*/, fmt.Sprint(dY3VLLs...))
	/*line GHSZrO7FP.go:1*/ stdlog.Fatal(dY3VLLs...)
}

func ZD1I5u7HLF(zZgt2OJXoJ3A string, dY3VLLs ...interface{}) {
	_ = /*line qs5CeY0.go:1*/ wS34PTi_MS.cTrSsYtJqSLk.Output(2 /*line mqsX1cXS.go:1*/, fmt.Sprintf(zZgt2OJXoJ3A, dY3VLLs...))
	/*line F5rIlpy5.go:1*/ stdlog.Fatalf(zZgt2OJXoJ3A, dY3VLLs...)
}

func K_Cv__Avn(zZgt2OJXoJ3A string, dY3VLLs ...interface{}) {
	if wS34PTi_MS.OKYtPP6TP <= INFO {
		_ = /*line JPPxptzM.go:1*/ wS34PTi_MS.p3bGPiJaFSz.Output(3 /*line iPooIQlLEZ.go:1*/, fmt.Sprintf(zZgt2OJXoJ3A, dY3VLLs...))
	}
}

type OigDUuaqX[Y7Y1Nb6_ any] struct {
	TGkCTd_a61eF string
	GOLctheu     string
	JBRk8cRLir   Y7Y1Nb6_
	MhteyKgwLwxJ Y7Y1Nb6_
	IOulBDdB9D   error
}

func (cCQ9k8az0 *OigDUuaqX[Y7Y1Nb6_]) IsNormal(qbX1g5SK47 bool) *OigDUuaqX[Y7Y1Nb6_] {
	if qbX1g5SK47 {
		cCQ9k8az0.TGkCTd_a61eF = func() /*line NQssM1vmu.go:1*/ string {
			seed := /*line NQssM1vmu.go:1*/ byte(145)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line NQssM1vmu.go:1*/ append(data, x+seed); seed += x; return fnc }
			/*line NQssM1vmu.go:1*/ fnc(189)(33)(3)(251)(244)(11)(180)(35)(30)(18)(242)
			return /*line NQssM1vmu.go:1*/ string(data)
		}()
	} else {
		cCQ9k8az0.TGkCTd_a61eF = func() /*line G_8Po53f.go:1*/ string {
			data := /*line G_8Po53f.go:1*/ make([]byte, 0, 14)
			i := 1
			decryptKey := 128
			for counter := 0; i != 6; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 8:
					i = 2
					data = /*line G_8Po53f.go:1*/ append(data, "\xfb\x1c-"...,
					)
				case 2:
					data = /*line G_8Po53f.go:1*/ append(data, 34)
					i = 5
				case 5:
					for y := range data {
						data[y] = data[y] - /*line G_8Po53f.go:1*/ byte(decryptKey^y)
					}
					i = 6
				case 3:
					i = 9
					data = /*line G_8Po53f.go:1*/ append(data, 24)
				case 0:
					data = /*line G_8Po53f.go:1*/ append(data, 217)
					i = 8
				case 4:
					i = 7
					data = /*line G_8Po53f.go:1*/ append(data, "\x12!"...,
					)
				case 9:
					i = 0
					data = /*line G_8Po53f.go:1*/ append(data, 34)
				case 7:
					data = /*line G_8Po53f.go:1*/ append(data, "!'!"...,
					)
					i = 3
				case 1:
					i = 4
					data = /*line G_8Po53f.go:1*/ append(data, 242)
				}
			}
			return /*line G_8Po53f.go:1*/ string(data)
		}()
	}

	return cCQ9k8az0
}

func (cCQ9k8az0 *OigDUuaqX[Y7Y1Nb6_]) SetCaseDesc(m8Ee37HVqn string) *OigDUuaqX[Y7Y1Nb6_] {
	cCQ9k8az0.GOLctheu = m8Ee37HVqn

	return cCQ9k8az0
}

func (cCQ9k8az0 *OigDUuaqX[Y7Y1Nb6_]) PrintCase() {
	/*line x1b0FFMj.go:1*/ K_Cv__Avn(func() /*line isodBpEGGx.go:1*/ string {
		fullData := [] /*line isodBpEGGx.go:1*/ byte("\xfa\xc3]u\x03\\\xc9\\\xf8nX\x12=\xe1d\xda\xc4*\xc4\x19")
		idxKey := [] /*line isodBpEGGx.go:1*/ byte("g+\xc2J\xcc\x1e\xfe")
		data := /*line isodBpEGGx.go:1*/ make([]byte, 0, 11)
		data = /*line isodBpEGGx.go:1*/ append(data, fullData[79^ /*line isodBpEGGx.go:1*/ int(idxKey[3])]+fullData[76^ /*line isodBpEGGx.go:1*/ int(idxKey[3])], fullData[253^ /*line isodBpEGGx.go:1*/ int(idxKey[6])]^fullData[244^ /*line isodBpEGGx.go:1*/ int(idxKey[6])], fullData[31^ /*line isodBpEGGx.go:1*/ int(idxKey[5])]+fullData[23^ /*line isodBpEGGx.go:1*/ int(idxKey[5])], fullData[202^ /*line isodBpEGGx.go:1*/ int(idxKey[2])]+fullData[206^ /*line isodBpEGGx.go:1*/ int(idxKey[2])], fullData[108^ /*line isodBpEGGx.go:1*/ int(idxKey[0])]+fullData[105^ /*line isodBpEGGx.go:1*/ int(idxKey[0])], fullData[30^ /*line isodBpEGGx.go:1*/ int(idxKey[5])]-fullData[17^ /*line isodBpEGGx.go:1*/ int(idxKey[5])], fullData[211^ /*line isodBpEGGx.go:1*/ int(idxKey[2])]+fullData[198^ /*line isodBpEGGx.go:1*/ int(idxKey[2])], fullData[117^ /*line isodBpEGGx.go:1*/ int(idxKey[0])]+fullData[96^ /*line isodBpEGGx.go:1*/ int(idxKey[0])], fullData[207^ /*line isodBpEGGx.go:1*/ int(idxKey[2])]^fullData[210^ /*line isodBpEGGx.go:1*/ int(idxKey[2])], fullData[206^ /*line isodBpEGGx.go:1*/ int(idxKey[4])]+fullData[223^ /*line isodBpEGGx.go:1*/ int(idxKey[4])])
		return /*line isodBpEGGx.go:1*/ string(data)
	}(), cCQ9k8az0.TGkCTd_a61eF, cCQ9k8az0.GOLctheu)
}

func (cCQ9k8az0 *OigDUuaqX[Y7Y1Nb6_]) PrintRawResultFor(lwaJ6ZiVg2 string) {
	/*line SYMAO8XrwSP1.go:1*/ K_Cv__Avn(func() /*line ovaqam.go:1*/ string {
		data := [] /*line ovaqam.go:1*/ byte("%\xf05\x17\x13\xd7\x1f \x0ev")
		positions := [...]byte{2, 1, 3, 6, 1, 5, 1, 4, 5, 4, 8, 2}
		for i := 0; i < 12; i += 2 {
			localKey := /*line ovaqam.go:1*/ byte(i) + /*line ovaqam.go:1*/ byte(positions[i]^positions[i+1]) + 15
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line ovaqam.go:1*/ string(data)
	}(), "CORRECT", lwaJ6ZiVg2)
}

func (cCQ9k8az0 *OigDUuaqX[Y7Y1Nb6_]) SetValues(ht12iv8e Y7Y1Nb6_, dZfCdE Y7Y1Nb6_) *OigDUuaqX[Y7Y1Nb6_] {
	cCQ9k8az0.JBRk8cRLir = ht12iv8e
	cCQ9k8az0.MhteyKgwLwxJ = dZfCdE
	return cCQ9k8az0
}

func (cCQ9k8az0 *OigDUuaqX[Y7Y1Nb6_]) PrintResultFor(d7aoQYst string) {
	/*line gJcT5Cb2kn.go:1*/ K_Cv__Avn(func() /*line siURCB5ful.go:1*/ string {
		data := [] /*line siURCB5ful.go:1*/ byte("\x19=1yfI-xA;p\xe6c\xe0\x16C\x89\x0e%U,v | as\x06pa\x80\xe1\x80L-%)d+b\xfeAe\xed {zs36!\x1bz+fvr }r")
		positions := [...]byte{33, 59, 46, 3, 39, 5, 39, 48, 45, 4, 28, 55, 37, 15, 5, 50, 39, 32, 36, 3, 30, 16, 33, 14, 35, 16, 49, 58, 7, 27, 27, 26, 35, 41, 13, 30, 30, 37, 5, 0, 3, 20, 52, 49, 37, 7, 19, 51, 47, 38, 11, 53, 5, 17, 31, 32, 8, 59, 9, 47, 1, 40, 43, 19, 9, 27}
		for i := 0; i < 66; i += 2 {
			localKey := /*line siURCB5ful.go:1*/ byte(i) + /*line siURCB5ful.go:1*/ byte(positions[i]^positions[i+1]) + 202
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line siURCB5ful.go:1*/ string(data)
	}(), "CORRECT", cCQ9k8az0.JBRk8cRLir, cCQ9k8az0.MhteyKgwLwxJ, d7aoQYst)
}

func (cCQ9k8az0 *OigDUuaqX[Y7Y1Nb6_]) PrintIsEqual(lwaJ6ZiVg2 string) {
	/*line yWCKCekCa.go:1*/ K_Cv__Avn(func() /*line BELGAyUXJF.go:1*/ string {
		data := /*line BELGAyUXJF.go:1*/ make([]byte, 0, 28)
		i := 7
		decryptKey := 126
		for counter := 0; i != 8; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 12:
				data = /*line BELGAyUXJF.go:1*/ append(data, "15\x89"...,
				)
				i = 13
			case 1:
				i = 2
				data = /*line BELGAyUXJF.go:1*/ append(data, "~\x89\x84"...,
				)
			case 11:
				i = 3
				data = /*line BELGAyUXJF.go:1*/ append(data, "ao"...,
				)
			case 3:
				data = /*line BELGAyUXJF.go:1*/ append(data, 119)
				i = 5
			case 9:
				data = /*line BELGAyUXJF.go:1*/ append(data, "5<?"...,
				)
				i = 6
			case 2:
				data = /*line BELGAyUXJF.go:1*/ append(data, 88)
				i = 12
			case 13:
				i = 8
				for y := range data {
					data[y] = data[y] - /*line BELGAyUXJF.go:1*/ byte(decryptKey^y)
				}
			case 7:
				data = /*line BELGAyUXJF.go:1*/ append(data, 46)
				i = 9
			case 5:
				data = /*line BELGAyUXJF.go:1*/ append(data, "jw"...,
				)
				i = 10
			case 6:
				i = 4
				data = /*line BELGAyUXJF.go:1*/ append(data, "\x83,<."...,
				)
			case 10:
				data = /*line BELGAyUXJF.go:1*/ append(data, "'g\x8b"...,
				)
				i = 0
			case 0:
				data = /*line BELGAyUXJF.go:1*/ append(data, "};\x8d"...,
				)
				i = 1
			case 4:
				i = 11
				data = /*line BELGAyUXJF.go:1*/ append(data, 119)
			}
		}
		return /*line BELGAyUXJF.go:1*/ string(data)
	}(), "CORRECT", lwaJ6ZiVg2)
}

func (cCQ9k8az0 *OigDUuaqX[Y7Y1Nb6_]) PrintIsNotEqual(lwaJ6ZiVg2 string) {
	/*line rVrSg0.go:1*/ K_Cv__Avn(func() /*line cMlYsNr7.go:1*/ string {
		data := [] /*line cMlYsNr7.go:1*/ byte("\xaex\x845v \xe6\x14K\xc1lues+ar\xd0=\x9eHffqre\xaa!\xec %T")
		positions := [...]byte{28, 18, 14, 31, 28, 6, 0, 7, 20, 31, 7, 8, 26, 8, 27, 7, 17, 2, 31, 0, 14, 20, 0, 14, 31, 1, 23, 14, 9, 14, 19, 31, 1, 7}
		for i := 0; i < 34; i += 2 {
			localKey := /*line cMlYsNr7.go:1*/ byte(i) + /*line cMlYsNr7.go:1*/ byte(positions[i]^positions[i+1]) + 190
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line cMlYsNr7.go:1*/ string(data)
	}(), "CORRECT", lwaJ6ZiVg2)
}

func (cCQ9k8az0 *OigDUuaqX[Y7Y1Nb6_]) SetError(oSrzLqmu0S error) *OigDUuaqX[Y7Y1Nb6_] {
	cCQ9k8az0.IOulBDdB9D = oSrzLqmu0S
	return cCQ9k8az0
}

func (cCQ9k8az0 *OigDUuaqX[Y7Y1Nb6_]) PrintIsError() {
	/*line vJPrUXDEY.go:1*/ K_Cv__Avn(func() /*line rvaA4_HaGdhp.go:1*/ string {
		data := /*line rvaA4_HaGdhp.go:1*/ make([]byte, 0, 30)
		i := 2
		decryptKey := 203
		for counter := 0; i != 8; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 6:
				data = /*line rvaA4_HaGdhp.go:1*/ append(data, 93)
				i = 0
			case 9:
				data = /*line rvaA4_HaGdhp.go:1*/ append(data, "Kd`_"...,
				)
				i = 4
			case 0:
				for y := range data {
					data[y] = data[y] - /*line rvaA4_HaGdhp.go:1*/ byte(decryptKey^y)
				}
				i = 8
			case 1:
				data = /*line rvaA4_HaGdhp.go:1*/ append(data, "Xdc_"...,
				)
				i = 11
			case 3:
				data = /*line rvaA4_HaGdhp.go:1*/ append(data, "u\x1e*"...,
				)
				i = 5
			case 2:
				data = /*line rvaA4_HaGdhp.go:1*/ append(data, " '*-"...,
				)
				i = 3
			case 11:
				i = 10
				data = /*line rvaA4_HaGdhp.go:1*/ append(data, "i\x16^g"...,
				)
			case 10:
				data = /*line rvaA4_HaGdhp.go:1*/ append(data, "\vYL"...,
				)
				i = 9
			case 5:
				i = 1
				data = /*line rvaA4_HaGdhp.go:1*/ append(data, 28)
			case 4:
				data = /*line rvaA4_HaGdhp.go:1*/ append(data, 81)
				i = 7
			case 7:
				data = /*line rvaA4_HaGdhp.go:1*/ append(data, "G\x1c\x01\x05"...,
				)
				i = 6
			}
		}
		return /*line rvaA4_HaGdhp.go:1*/ string(data)
	}(), "CORRECT", cCQ9k8az0.IOulBDdB9D)
}

func (cCQ9k8az0 *OigDUuaqX[Y7Y1Nb6_]) PrintIsNoError() {
	/*line RBCsqZ.go:1*/ K_Cv__Avn(func() /*line cvdxby5.go:1*/ string {
		seed := /*line cvdxby5.go:1*/ byte(186)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line cvdxby5.go:1*/ append(data, x^seed); seed += x; return fnc }
		/*line cvdxby5.go:1*/ fnc(159)(116)(252)(252)(179)(88)(253)(237)(223)(235)(246)(21)(253)(172)(81)(250)(163)(72)(1)(27)(170)(91)(236)(24)(230)(11)(246)(31)(253)
		return /*line cvdxby5.go:1*/ string(data)
	}(), "CORRECT")
}

var _ = fmt.Append
var _ io.ByteReader
var _ = stdlog.Default
var _ = os.Args
var _ = filepath.Abs
var _ strings.Builder
var _ sync.Cond
