package array

func Slice[GA ~[]A, A any](low, high int) func(as GA) GA {
	return func(as GA) GA {
		return as[low:high]
	}
}

func IsEmpty[GA ~[]A, A any](as GA) bool {
	return len(as) == 0
}

func IsNil[GA ~[]A, A any](as GA) bool {
	return as == nil
}

func IsNonNil[GA ~[]A, A any](as GA) bool {
	return as != nil
}

func Reduce[GA ~[]A, A, B any](fa GA, f func(B, A) B, initial B) B {
	current := initial
	count := len(fa)
	for i := 0; i < count; i++ {
		current = f(current, fa[i])
	}
	return current
}

func Append[GA ~[]A, A any](as GA, a A) GA {
	return append(as, a)
}

func Empty[GA ~[]A, A any]() GA {
	return make(GA, 0)
}

func upsertAt[GA ~[]A, A any](fa GA, a A) GA {
	buf := make(GA, len(fa)+1)
	buf[copy(buf, fa)] = a
	return buf
}

func UpsertAt[GA ~[]A, A any](a A) func(GA) GA {
	return func(ma GA) GA {
		return upsertAt(ma, a)
	}
}

func MonadMap[GA ~[]A, GB ~[]B, A, B any](as GA, f func(a A) B) GB {
	count := len(as)
	bs := make(GB, count)
	for i := count - 1; i >= 0; i-- {
		bs[i] = f(as[i])
	}
	return bs
}

func ConstNil[GA ~[]A, A any]() GA {
	return (GA)(nil)
}