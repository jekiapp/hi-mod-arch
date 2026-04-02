# Model

Model is shared data shape used by usecase, logic, and repository.

## Guidelines

- Mostly plain structs.
- No business flow in model.
- Put formula/conversion in `internal/logic`.

## Example

```go
type CartItem struct {
	ProductID int64
	Quantity  int64
}

type CheckoutItem struct {
	Product  ProductData
	Quantity int64
	Subtotal float64
}
```
