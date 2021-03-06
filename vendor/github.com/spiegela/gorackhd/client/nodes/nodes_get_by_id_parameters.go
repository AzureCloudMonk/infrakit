package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNodesGetByIDParams creates a new NodesGetByIDParams object
// with the default values initialized.
func NewNodesGetByIDParams() *NodesGetByIDParams {
	var ()
	return &NodesGetByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNodesGetByIDParamsWithTimeout creates a new NodesGetByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodesGetByIDParamsWithTimeout(timeout time.Duration) *NodesGetByIDParams {
	var ()
	return &NodesGetByIDParams{

		timeout: timeout,
	}
}

// NewNodesGetByIDParamsWithContext creates a new NodesGetByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewNodesGetByIDParamsWithContext(ctx context.Context) *NodesGetByIDParams {
	var ()
	return &NodesGetByIDParams{

		Context: ctx,
	}
}

// NewNodesGetByIDParamsWithHTTPClient creates a new NodesGetByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNodesGetByIDParamsWithHTTPClient(client *http.Client) *NodesGetByIDParams {
	var ()
	return &NodesGetByIDParams{
		HTTPClient: client,
	}
}

/*NodesGetByIDParams contains all the parameters to send to the API endpoint
for the nodes get by Id operation typically these are written to a http.Request
*/
type NodesGetByIDParams struct {

	/*Identifier
	  The node identifier

	*/
	Identifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the nodes get by Id params
func (o *NodesGetByIDParams) WithTimeout(timeout time.Duration) *NodesGetByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nodes get by Id params
func (o *NodesGetByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nodes get by Id params
func (o *NodesGetByIDParams) WithContext(ctx context.Context) *NodesGetByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nodes get by Id params
func (o *NodesGetByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nodes get by Id params
func (o *NodesGetByIDParams) WithHTTPClient(client *http.Client) *NodesGetByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nodes get by Id params
func (o *NodesGetByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentifier adds the identifier to the nodes get by Id params
func (o *NodesGetByIDParams) WithIdentifier(identifier string) *NodesGetByIDParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the nodes get by Id params
func (o *NodesGetByIDParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WriteToRequest writes these params to a swagger request
func (o *NodesGetByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
