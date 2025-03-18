// This file was auto-generated by Fern from our API Definition.

package basistheory

import (
	json "encoding/json"
	core "github.com/Basis-Theory/go-sdk/core"
)

// Bad Request
type BadRequestError struct {
	*core.APIError
	Body *ValidationProblemDetails
}

func (b *BadRequestError) UnmarshalJSON(data []byte) error {
	var body *ValidationProblemDetails
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	b.StatusCode = 400
	b.Body = body
	return nil
}

func (b *BadRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Body)
}

func (b *BadRequestError) Unwrap() error {
	return b.APIError
}

// Conflict
type ConflictError struct {
	*core.APIError
	Body *ProblemDetails
}

func (c *ConflictError) UnmarshalJSON(data []byte) error {
	var body *ProblemDetails
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	c.StatusCode = 409
	c.Body = body
	return nil
}

func (c *ConflictError) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Body)
}

func (c *ConflictError) Unwrap() error {
	return c.APIError
}

// Forbidden
type ForbiddenError struct {
	*core.APIError
	Body *ProblemDetails
}

func (f *ForbiddenError) UnmarshalJSON(data []byte) error {
	var body *ProblemDetails
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	f.StatusCode = 403
	f.Body = body
	return nil
}

func (f *ForbiddenError) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Body)
}

func (f *ForbiddenError) Unwrap() error {
	return f.APIError
}

// Not Found
type NotFoundError struct {
	*core.APIError
	Body interface{}
}

func (n *NotFoundError) UnmarshalJSON(data []byte) error {
	var body interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	n.StatusCode = 404
	n.Body = body
	return nil
}

func (n *NotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Body)
}

func (n *NotFoundError) Unwrap() error {
	return n.APIError
}

// Server Error
type ServiceUnavailableError struct {
	*core.APIError
	Body *ProblemDetails
}

func (s *ServiceUnavailableError) UnmarshalJSON(data []byte) error {
	var body *ProblemDetails
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	s.StatusCode = 503
	s.Body = body
	return nil
}

func (s *ServiceUnavailableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Body)
}

func (s *ServiceUnavailableError) Unwrap() error {
	return s.APIError
}

// Unauthorized
type UnauthorizedError struct {
	*core.APIError
	Body *ProblemDetails
}

func (u *UnauthorizedError) UnmarshalJSON(data []byte) error {
	var body *ProblemDetails
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.StatusCode = 401
	u.Body = body
	return nil
}

func (u *UnauthorizedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}

func (u *UnauthorizedError) Unwrap() error {
	return u.APIError
}

// Client Error
type UnprocessableEntityError struct {
	*core.APIError
	Body *ProblemDetails
}

func (u *UnprocessableEntityError) UnmarshalJSON(data []byte) error {
	var body *ProblemDetails
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.StatusCode = 422
	u.Body = body
	return nil
}

func (u *UnprocessableEntityError) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}

func (u *UnprocessableEntityError) Unwrap() error {
	return u.APIError
}
