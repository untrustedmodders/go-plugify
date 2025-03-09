package plugify

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
