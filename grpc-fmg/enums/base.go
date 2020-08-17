package enumsbase


type EnumBaseInterface interface {
	Has(int, ...int) bool
	Mixin(...int) int
}

type EnumBase struct {
	Enums []int
}

func (e EnumBase) Has(enum int, enums ...int) bool {
	var mixin int
	if len(enums) > 0 {
		mixin = e.Mixin(enums...)
	} else {
		mixin = e.Mixin(e.Enums...)
	}
	if mixin & enum == enum {
		return true
	}
	return false
}

func (e EnumBase) Mixin(enums ...int) int {
	mixin := enums[0]
	for i := 1; i < len(enums); i++ {
		mixin = mixin | enums[i]
	}
	return mixin
}