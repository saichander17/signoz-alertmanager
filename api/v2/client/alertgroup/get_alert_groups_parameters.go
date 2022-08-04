// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package alertgroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetAlertGroupsParams creates a new GetAlertGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAlertGroupsParams() *GetAlertGroupsParams {
	return &GetAlertGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlertGroupsParamsWithTimeout creates a new GetAlertGroupsParams object
// with the ability to set a timeout on a request.
func NewGetAlertGroupsParamsWithTimeout(timeout time.Duration) *GetAlertGroupsParams {
	return &GetAlertGroupsParams{
		timeout: timeout,
	}
}

// NewGetAlertGroupsParamsWithContext creates a new GetAlertGroupsParams object
// with the ability to set a context for a request.
func NewGetAlertGroupsParamsWithContext(ctx context.Context) *GetAlertGroupsParams {
	return &GetAlertGroupsParams{
		Context: ctx,
	}
}

// NewGetAlertGroupsParamsWithHTTPClient creates a new GetAlertGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAlertGroupsParamsWithHTTPClient(client *http.Client) *GetAlertGroupsParams {
	return &GetAlertGroupsParams{
		HTTPClient: client,
	}
}

/* GetAlertGroupsParams contains all the parameters to send to the API endpoint
   for the get alert groups operation.

   Typically these are written to a http.Request.
*/
type GetAlertGroupsParams struct {

	/* Active.

	   Show active alerts

	   Default: true
	*/
	Active *bool

	/* Filter.

	   A list of matchers to filter alerts by
	*/
	Filter []string

	/* Inhibited.

	   Show inhibited alerts

	   Default: true
	*/
	Inhibited *bool

	/* Receiver.

	   A regex matching receivers to filter alerts by
	*/
	Receiver *string

	/* Silenced.

	   Show silenced alerts

	   Default: true
	*/
	Silenced *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get alert groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertGroupsParams) WithDefaults() *GetAlertGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get alert groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertGroupsParams) SetDefaults() {
	var (
		activeDefault = bool(true)

		inhibitedDefault = bool(true)

		silencedDefault = bool(true)
	)

	val := GetAlertGroupsParams{
		Active:    &activeDefault,
		Inhibited: &inhibitedDefault,
		Silenced:  &silencedDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get alert groups params
func (o *GetAlertGroupsParams) WithTimeout(timeout time.Duration) *GetAlertGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alert groups params
func (o *GetAlertGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alert groups params
func (o *GetAlertGroupsParams) WithContext(ctx context.Context) *GetAlertGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alert groups params
func (o *GetAlertGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alert groups params
func (o *GetAlertGroupsParams) WithHTTPClient(client *http.Client) *GetAlertGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alert groups params
func (o *GetAlertGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActive adds the active to the get alert groups params
func (o *GetAlertGroupsParams) WithActive(active *bool) *GetAlertGroupsParams {
	o.SetActive(active)
	return o
}

// SetActive adds the active to the get alert groups params
func (o *GetAlertGroupsParams) SetActive(active *bool) {
	o.Active = active
}

// WithFilter adds the filter to the get alert groups params
func (o *GetAlertGroupsParams) WithFilter(filter []string) *GetAlertGroupsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get alert groups params
func (o *GetAlertGroupsParams) SetFilter(filter []string) {
	o.Filter = filter
}

// WithInhibited adds the inhibited to the get alert groups params
func (o *GetAlertGroupsParams) WithInhibited(inhibited *bool) *GetAlertGroupsParams {
	o.SetInhibited(inhibited)
	return o
}

// SetInhibited adds the inhibited to the get alert groups params
func (o *GetAlertGroupsParams) SetInhibited(inhibited *bool) {
	o.Inhibited = inhibited
}

// WithReceiver adds the receiver to the get alert groups params
func (o *GetAlertGroupsParams) WithReceiver(receiver *string) *GetAlertGroupsParams {
	o.SetReceiver(receiver)
	return o
}

// SetReceiver adds the receiver to the get alert groups params
func (o *GetAlertGroupsParams) SetReceiver(receiver *string) {
	o.Receiver = receiver
}

// WithSilenced adds the silenced to the get alert groups params
func (o *GetAlertGroupsParams) WithSilenced(silenced *bool) *GetAlertGroupsParams {
	o.SetSilenced(silenced)
	return o
}

// SetSilenced adds the silenced to the get alert groups params
func (o *GetAlertGroupsParams) SetSilenced(silenced *bool) {
	o.Silenced = silenced
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlertGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Active != nil {

		// query param active
		var qrActive bool

		if o.Active != nil {
			qrActive = *o.Active
		}
		qActive := swag.FormatBool(qrActive)
		if qActive != "" {

			if err := r.SetQueryParam("active", qActive); err != nil {
				return err
			}
		}
	}

	if o.Filter != nil {

		// binding items for filter
		joinedFilter := o.bindParamFilter(reg)

		// query array param filter
		if err := r.SetQueryParam("filter", joinedFilter...); err != nil {
			return err
		}
	}

	if o.Inhibited != nil {

		// query param inhibited
		var qrInhibited bool

		if o.Inhibited != nil {
			qrInhibited = *o.Inhibited
		}
		qInhibited := swag.FormatBool(qrInhibited)
		if qInhibited != "" {

			if err := r.SetQueryParam("inhibited", qInhibited); err != nil {
				return err
			}
		}
	}

	if o.Receiver != nil {

		// query param receiver
		var qrReceiver string

		if o.Receiver != nil {
			qrReceiver = *o.Receiver
		}
		qReceiver := qrReceiver
		if qReceiver != "" {

			if err := r.SetQueryParam("receiver", qReceiver); err != nil {
				return err
			}
		}
	}

	if o.Silenced != nil {

		// query param silenced
		var qrSilenced bool

		if o.Silenced != nil {
			qrSilenced = *o.Silenced
		}
		qSilenced := swag.FormatBool(qrSilenced)
		if qSilenced != "" {

			if err := r.SetQueryParam("silenced", qSilenced); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetAlertGroups binds the parameter filter
func (o *GetAlertGroupsParams) bindParamFilter(formats strfmt.Registry) []string {
	filterIR := o.Filter

	var filterIC []string
	for _, filterIIR := range filterIR { // explode []string

		filterIIV := filterIIR // string as string
		filterIC = append(filterIC, filterIIV)
	}

	// items.CollectionFormat: "multi"
	filterIS := swag.JoinByFormat(filterIC, "multi")

	return filterIS
}
