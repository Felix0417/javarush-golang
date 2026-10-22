package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type JobRes struct {
	closeFlag string
}

func (r JobRes) Close() error {
	// TODO: Вернуть ошибку с текстом "close failed", если r.closeFlag == "fail".
	if r.closeFlag == "fail" {
		return errors.New("close failed")
	}
	return nil
}

func runJob(i int, workFlag string, closeFlag string) (err error) {
	_ = i // сигнатуру менять нельзя

	res := JobRes{closeFlag: closeFlag}

	// TODO: Закрывать ресурс через defer внутри runJob (а не в main), чтобы он закрывался для каждого job отдельно.
	defer func() {
		if cerr := res.Close(); cerr != nil {
			// TODO: Объединить err и cerr через errors.Join, чтобы не потерять ошибки работы и закрытия.
			if cerr != nil {
				err = errors.Join(err, cerr)
			}
		}
	}()

	// TODO: Реализовать "основную работу": при workFlag == "fail" вернуть ошибку с текстом "work failed".
	if workFlag == "fail" {
		return errors.New("work failed")
	}

	return nil
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i <= n; i++ {
		var workTok, closeTok string
		fmt.Fscan(in, &workTok, &closeTok)

		workFlag := strings.TrimPrefix(workTok, "work=")
		closeFlag := strings.TrimPrefix(closeTok, "close=")

		err := runJob(i, workFlag, closeFlag)
		fmt.Printf("job %d: %v\n", i, err)
	}
}
