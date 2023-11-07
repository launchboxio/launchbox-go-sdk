package launchbox

type GenericError struct{}

func (*GenericError) Error() string {
	return "An unexpected error occured"
}

type UnauthorizedError struct{}

func (*UnauthorizedError) Error() string {
	return "Unauthorized"
}

type ForbiddenError struct{}

func (*ForbiddenError) Error() string {
	return "Forbidden"
}

type ResourceNotFoundError struct{}

func (*ResourceNotFoundError) Error() string {
	return "Resource not found"
}
