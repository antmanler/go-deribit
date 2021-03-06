// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPrivateGetOrderHistoryByInstrumentParams creates a new GetPrivateGetOrderHistoryByInstrumentParams object
// with the default values initialized.
func NewGetPrivateGetOrderHistoryByInstrumentParams() *GetPrivateGetOrderHistoryByInstrumentParams {
	var ()
	return &GetPrivateGetOrderHistoryByInstrumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateGetOrderHistoryByInstrumentParamsWithTimeout creates a new GetPrivateGetOrderHistoryByInstrumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateGetOrderHistoryByInstrumentParamsWithTimeout(timeout time.Duration) *GetPrivateGetOrderHistoryByInstrumentParams {
	var ()
	return &GetPrivateGetOrderHistoryByInstrumentParams{

		timeout: timeout,
	}
}

// NewGetPrivateGetOrderHistoryByInstrumentParamsWithContext creates a new GetPrivateGetOrderHistoryByInstrumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateGetOrderHistoryByInstrumentParamsWithContext(ctx context.Context) *GetPrivateGetOrderHistoryByInstrumentParams {
	var ()
	return &GetPrivateGetOrderHistoryByInstrumentParams{

		Context: ctx,
	}
}

// NewGetPrivateGetOrderHistoryByInstrumentParamsWithHTTPClient creates a new GetPrivateGetOrderHistoryByInstrumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateGetOrderHistoryByInstrumentParamsWithHTTPClient(client *http.Client) *GetPrivateGetOrderHistoryByInstrumentParams {
	var ()
	return &GetPrivateGetOrderHistoryByInstrumentParams{
		HTTPClient: client,
	}
}

/*GetPrivateGetOrderHistoryByInstrumentParams contains all the parameters to send to the API endpoint
for the get private get order history by instrument operation typically these are written to a http.Request
*/
type GetPrivateGetOrderHistoryByInstrumentParams struct {

	/*Count
	  Number of requested items, default - `20`

	*/
	Count *int64
	/*IncludeOld
	  Include in result orders older than 2 days, default - `false`

	*/
	IncludeOld *bool
	/*IncludeUnfilled
	  Include in result fully unfilled closed orders, default - `false`

	*/
	IncludeUnfilled *bool
	/*InstrumentName
	  Instrument name

	*/
	InstrumentName string
	/*Offset
	  The offset for pagination, default - `0`

	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private get order history by instrument params
func (o *GetPrivateGetOrderHistoryByInstrumentParams) WithTimeout(timeout time.Duration) *GetPrivateGetOrderHistoryByInstrumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private get order history by instrument params
func (o *GetPrivateGetOrderHistoryByInstrumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private get order history by instrument params
func (o *GetPrivateGetOrderHistoryByInstrumentParams) WithContext(ctx context.Context) *GetPrivateGetOrderHistoryByInstrumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private get order history by instrument params
func (o *GetPrivateGetOrderHistoryByInstrumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private get order history by instrument params
func (o *GetPrivateGetOrderHistoryByInstrumentParams) WithHTTPClient(client *http.Client) *GetPrivateGetOrderHistoryByInstrumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private get order history by instrument params
func (o *GetPrivateGetOrderHistoryByInstrumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get private get order history by instrument params
func (o *GetPrivateGetOrderHistoryByInstrumentParams) WithCount(count *int64) *GetPrivateGetOrderHistoryByInstrumentParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get private get order history by instrument params
func (o *GetPrivateGetOrderHistoryByInstrumentParams) SetCount(count *int64) {
	o.Count = count
}

// WithIncludeOld adds the includeOld to the get private get order history by instrument params
func (o *GetPrivateGetOrderHistoryByInstrumentParams) WithIncludeOld(includeOld *bool) *GetPrivateGetOrderHistoryByInstrumentParams {
	o.SetIncludeOld(includeOld)
	return o
}

// SetIncludeOld adds the includeOld to the get private get order history by instrument params
func (o *GetPrivateGetOrderHistoryByInstrumentParams) SetIncludeOld(includeOld *bool) {
	o.IncludeOld = includeOld
}

// WithIncludeUnfilled adds the includeUnfilled to the get private get order history by instrument params
func (o *GetPrivateGetOrderHistoryByInstrumentParams) WithIncludeUnfilled(includeUnfilled *bool) *GetPrivateGetOrderHistoryByInstrumentParams {
	o.SetIncludeUnfilled(includeUnfilled)
	return o
}

// SetIncludeUnfilled adds the includeUnfilled to the get private get order history by instrument params
func (o *GetPrivateGetOrderHistoryByInstrumentParams) SetIncludeUnfilled(includeUnfilled *bool) {
	o.IncludeUnfilled = includeUnfilled
}

// WithInstrumentName adds the instrumentName to the get private get order history by instrument params
func (o *GetPrivateGetOrderHistoryByInstrumentParams) WithInstrumentName(instrumentName string) *GetPrivateGetOrderHistoryByInstrumentParams {
	o.SetInstrumentName(instrumentName)
	return o
}

// SetInstrumentName adds the instrumentName to the get private get order history by instrument params
func (o *GetPrivateGetOrderHistoryByInstrumentParams) SetInstrumentName(instrumentName string) {
	o.InstrumentName = instrumentName
}

// WithOffset adds the offset to the get private get order history by instrument params
func (o *GetPrivateGetOrderHistoryByInstrumentParams) WithOffset(offset *int64) *GetPrivateGetOrderHistoryByInstrumentParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get private get order history by instrument params
func (o *GetPrivateGetOrderHistoryByInstrumentParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateGetOrderHistoryByInstrumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int64
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	if o.IncludeOld != nil {

		// query param include_old
		var qrIncludeOld bool
		if o.IncludeOld != nil {
			qrIncludeOld = *o.IncludeOld
		}
		qIncludeOld := swag.FormatBool(qrIncludeOld)
		if qIncludeOld != "" {
			if err := r.SetQueryParam("include_old", qIncludeOld); err != nil {
				return err
			}
		}

	}

	if o.IncludeUnfilled != nil {

		// query param include_unfilled
		var qrIncludeUnfilled bool
		if o.IncludeUnfilled != nil {
			qrIncludeUnfilled = *o.IncludeUnfilled
		}
		qIncludeUnfilled := swag.FormatBool(qrIncludeUnfilled)
		if qIncludeUnfilled != "" {
			if err := r.SetQueryParam("include_unfilled", qIncludeUnfilled); err != nil {
				return err
			}
		}

	}

	// query param instrument_name
	qrInstrumentName := o.InstrumentName
	qInstrumentName := qrInstrumentName
	if qInstrumentName != "" {
		if err := r.SetQueryParam("instrument_name", qInstrumentName); err != nil {
			return err
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
