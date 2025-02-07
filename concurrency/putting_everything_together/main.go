package main

import (
	"context"
	"errors"
	"time"
)

type abProcessor struct {
	outA chan aOut
	outB chan bOut
	errs chan error
}

func newABProcessor() *abProcessor {
	return &abProcessor{
		outA: make(chan aOut, 1),
		outB: make(chan bOut, 1),
		errs: make(chan error, 2),
	}
}

type cProcessor struct {
	outC chan cOut
	err  chan error
}

func newCProcessor() *cProcessor {
	return &cProcessor{
		outC: make(chan cOut, 1),
		err:  make(chan error, 1),
	}
}

func (abp *abProcessor) start(ctx context.Context, data Input) {
	go func() {
		aOut, err := getResultA(ctx, data.A)
		if err != nil {
			abp.errs <- err
			return
		}
		abp.outA <- aOut
	}()

	go func() {
		bOut, err := getResultB(ctx, data.B)
		if err != nil {
			abp.errs <- err
			return
		}
		abp.outB <- bOut
	}()
}

func (cp *cProcessor) start(ctx context.Context, data Input) {
	go func() {
		cOut, err := getResultC(ctx, inputC)
		if err != nil {
			p.errs <- err
			return
		}
		p.outC <- cOut
	}()
}

func (abp *abProcessor) wait(ctx context.Context) (cIn, error) {
	var cData cIn
	for count := 0; count < 2; count++ {
		select {
		case <-ctx.Done():
			return cIn{}, ctx.Err()
		case a := <-abp.outA:
			cData.A = a
		case b := <-abp.outB:
			cData.B = b
		case err := <-abp.errs:
			return cIn{}, err
		}
	}
}

func (cp *cProcessor) wait(ctx context.Context) {
	select {
	case out := <-p.outC:
		return out, nil
	case err := <-p.errs:
		return COut{}, err
	case <-ctx.Done():
		return COut{}, ctx.Err()
	}
}

func GatherAndProcess(ctx context.Context, data Input) (COut, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	ab := newABProcessor()
	ab.start(ctx, data)
	inputC, err := ab.wait(ctx)
	if err != nil {
		return COut{}, err
	}

	c := newCProcessor()
	c.start(ctx, inputC)
	out, err := c.wait(ctx)
	if err != nil {
		return COut{}, err
	}

	return out, nil
}
