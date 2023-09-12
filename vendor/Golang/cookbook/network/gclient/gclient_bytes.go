// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gclient

import (
	"context"
	"fmt"
)

// GetBytes sends a GET request, retrieves and returns the result content as bytes.
func (c *Client) GetBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, httpMethodGet, url, data...)
}

// PutBytes sends a PUT request, retrieves and returns the result content as bytes.
func (c *Client) PutBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, httpMethodPut, url, data...)
}

// PostBytes sends a POST request, retrieves and returns the result content as bytes.
func (c *Client) PostBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, httpMethodPost, url, data...)
}

// DeleteBytes sends a DELETE request, retrieves and returns the result content as bytes.
func (c *Client) DeleteBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, httpMethodDelete, url, data...)
}

// HeadBytes sends a HEAD request, retrieves and returns the result content as bytes.
func (c *Client) HeadBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, httpMethodHead, url, data...)
}

// PatchBytes sends a PATCH request, retrieves and returns the result content as bytes.
func (c *Client) PatchBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, httpMethodPatch, url, data...)
}

// ConnectBytes sends a CONNECT request, retrieves and returns the result content as bytes.
func (c *Client) ConnectBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, httpMethodConnect, url, data...)
}

// OptionsBytes sends a OPTIONS request, retrieves and returns the result content as bytes.
func (c *Client) OptionsBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, httpMethodOptions, url, data...)
}

// TraceBytes sends a TRACE request, retrieves and returns the result content as bytes.
func (c *Client) TraceBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, httpMethodTrace, url, data...)
}

// RequestBytes sends request using given HTTP method and data, retrieves returns the result
// as bytes. It reads and closes the response object internally automatically.
func (c *Client) RequestBytes(ctx context.Context, method string, url string, data ...interface{}) []byte {
	response, err := c.DoRequest(ctx, method, url, data...)
	if err != nil {
		fmt.Errorf(`%+v`, err)
		return nil
	}
	defer func() {
		if err = response.Close(); err != nil {
			fmt.Errorf(`%+v`, err)
		}
	}()
	return response.ReadAll()
}
