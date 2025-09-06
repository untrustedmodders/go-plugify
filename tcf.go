package plugify

import "C"
import "fmt"

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception any

func Throw(up Exception) {
	panicker(up)
}

func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

func panicker(v any) {
	msg := fmt.Sprintf("%v", v)
	stack := plugify.fnPluginPanicCallback()
	if len(stack) > 0 {
		msg += fmt.Sprintf("\nStack Trace: \n%s", stack)
	}
	C.Plugify_PrintException(msg)
	panic(v)
}
