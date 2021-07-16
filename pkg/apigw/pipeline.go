package apigw

import (
	"context"
	"fmt"
	"net/http"
)

type (
	Execer interface {
		Exec(context.Context, *scp) error
	}

	Sorter interface {
		Weight() int
	}

	ErrorHandler interface {
		Exec(context.Context, *scp, error)
	}

	Worker interface {
		Execer
		// Sorter
	}

	workers []Worker

	pl struct {
		w   workers
		err ErrorHandler
	}

	scp map[string]interface{}
)

func (s scp) Request() *http.Request {
	if _, ok := s["request"]; ok {
		return s["request"].(*http.Request)
	}

	return nil
}

func (s scp) Writer() http.ResponseWriter {
	if _, ok := s["writer"]; ok {
		return s["writer"].(http.ResponseWriter)
	}

	return nil
}

func (s scp) Set(k string, v interface{}) {
	s[k] = v
}

func (s scp) Get(k string) (v interface{}, err error) {
	var ok bool

	if v, ok = s[k]; !ok {
		err = fmt.Errorf("could not get key on index: %s", k)
		return
	}

	return
}

// Exec takes care of error handling and main
// functionality that takes place in worker
func (pp *pl) Exec(ctx context.Context, scope *scp) (err error) {
	for _, w := range pp.w {
		err = w.Exec(ctx, scope)

		if err != nil {
			if pp.err != nil {
				// call the error handler
				pp.err.Exec(ctx, scope, err)
			}

			return
		}
	}

	return
}

// Add registers a new worker with parameters
// fetched from store
func (pp *pl) Add(ff Worker) {
	pp.w = append(pp.w, ff)
	// sort.Sort(pp.w)
}

// add error handler
func (pp *pl) ErrorHandler(ff ErrorHandler) {
	pp.err = ff
}

// func (a workers) Len() int { return len(a) }
// func (a workers) Less(i, j int) bool {
// 	return a[i].worker.Weight() < a[j].worker.Weight()
// }
// func (a workers) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
