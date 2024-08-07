package models

type UpdateProviderTableInput struct {
	Provider    int
	Start       int
	End         int
	RequestId   int
	UserId      int
	IsAvailable int
}
