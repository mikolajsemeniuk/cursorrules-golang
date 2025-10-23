package main

import (
	"errors"
	"fmt"
)

var ErrProductNotFound = errors.New("product not found")

func main() {
	productID := "12345"

	err1 := fmt.Errorf("%w, productID: %s", ErrProductNotFound, productID)

	fmt.Println("=== EXAMPLE 1: fmt.Errorf(\"%w ...\") ===")
	fmt.Println("Error message:", err1)
	fmt.Println("errors.Is(err1, ErrProductNotFound):", errors.Is(err1, ErrProductNotFound))
	fmt.Println()

	// --- Example 2: errors.Join (joining multiple errors)
	err2 := errors.Join(ErrProductNotFound, fmt.Errorf("productID: %s", productID))

	fmt.Println("=== EXAMPLE 2: errors.Join(...) ===")
	fmt.Println("Error message:", err2)
	fmt.Println("errors.Is(err2, ErrProductNotFound):", errors.Is(err2, ErrProductNotFound))
	fmt.Println()

	// --- Regular error
	err3 := fmt.Errorf("product not found (productID=%s)", productID)
	fmt.Println("=== EXAMPLE 3: zwykły error bez wrapa ===")
	fmt.Println("Error message:", err3)
	fmt.Println("errors.Is(err3, ErrProductNotFound):", errors.Is(err3, ErrProductNotFound))
	fmt.Println()

	fmt.Println("=== EXAMPLE 4: errors.Unwrap(...) ===")
	deep := fmt.Errorf("level3: %w", errors.New("root"))
	mid := fmt.Errorf("level2: %w", deep)
	top := fmt.Errorf("level1: %w", mid)

	fmt.Println("top:", top)
	u1 := errors.Unwrap(top)
	fmt.Println("u1:", u1)
	u2 := errors.Unwrap(u1)
	fmt.Println("u2:", u2)
	u3 := errors.Unwrap(u2)
	fmt.Println("u3:", u3)
}
