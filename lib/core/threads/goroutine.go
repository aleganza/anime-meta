package threads

type ParallelFetchFunc struct {
	Name string
	Func func() (any, error)
}

type ParallelFetchResult struct {
	Name string
	Body any
	Err  error
}

func RunParallel(funcs ...ParallelFetchFunc) []ParallelFetchResult {
	ch := make(chan ParallelFetchResult, len(funcs))

	for _, fn := range funcs {
		go func(fn ParallelFetchFunc) {
			data, err := fn.Func()

			ch <- ParallelFetchResult{
				Name: fn.Name,
				Body: data,
				Err:  err,
			}
		}(fn)
	}

	results := make([]ParallelFetchResult, 0, len(funcs))

	for range funcs {
		results = append(results, <-ch)
	}

	return results
}