package q

func Assert(truth bool, msg string) {
	if !truth {
		Boom("Assertion failed:" + msg)
	}
}

func AssertNoError(err error) {
	AssertNil(err, err.Error())
}

func AssertNotNil(item any, msg string) {
	Assert(item != nil, msg)
}

func AssertNil(item any, msg string) {
	Assert(item == nil, msg)
}