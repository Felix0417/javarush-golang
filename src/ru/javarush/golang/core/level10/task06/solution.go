package main

import (
	"fmt"
	"net/http"
)

func main() {
	// TODO: Сформируйте строку из http.MethodGet и http.StatusOK, разделив их пробелом, и выведите ровно одну строку с переводом строки.

	fmt.Println(http.MethodGet, http.StatusOK)
}
