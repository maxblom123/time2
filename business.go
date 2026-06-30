package time2

func (t T) IsBusinessDay() bool { return t.IsWeekday() }

func (t T) NextBusinessDay() T {
	next := t.AddDays(1)
	for next.IsWeekend() {
		next = next.AddDays(1)
	}
	return next
}

func (t T) PrevBusinessDay() T {
	prev := t.AddDays(-1)
	for prev.IsWeekend() {
		prev = prev.AddDays(-1)
	}
	return prev
}

func (t T) AddBusinessDays(n int) T {
	result := t
	step := 1
	if n < 0 {
		step = -1
		n = -n
	}
	for added := 0; added < n; {
		result = result.AddDays(step)
		if result.IsBusinessDay() {
			added++
		}
	}
	return result
}

func (t T) BusinessDaysBetween(u T) int {
	start, end := t, u
	if t.After(u.Time) {
		start, end = u, t
	}
	count := 0
	for cur := start.AddDays(1); !cur.After(end.Time); cur = cur.AddDays(1) {
		if cur.IsBusinessDay() {
			count++
		}
	}
	return count
}