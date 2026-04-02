# Usecase

Usecase is where you write the business flow.

## Guidelines

- One file = one business flow.
- Call external data through `repo` interface.
- Put reusable formulas/conversions in `internal/logic`.

## Example

```go
type renderPageUsecase struct {
	repo iRenderPageRepo
}

type iRenderPageRepo interface {
	GetCartFromDB(userID int64) (model.CartData, error)
	GetUserInfo(userID int64) (model.UserData, error)
}

func (uc renderPageUsecase) HttpGenericHandler(ctx context.Context, input CheckoutPageRequest) (CheckoutPageResponse, error) {
	cart, err := uc.repo.GetCartFromDB(input.UserID)
	if err != nil {
		return CheckoutPageResponse{}, err
	}
	_ = cart
	return CheckoutPageResponse{}, nil
}
```