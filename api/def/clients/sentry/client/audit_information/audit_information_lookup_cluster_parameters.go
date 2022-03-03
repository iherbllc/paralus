// Code generated by go-swagger; DO NOT EDIT.

package audit_information

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
)

// NewAuditInformationLookupClusterParams creates a new AuditInformationLookupClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuditInformationLookupClusterParams() *AuditInformationLookupClusterParams {
	return &AuditInformationLookupClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuditInformationLookupClusterParamsWithTimeout creates a new AuditInformationLookupClusterParams object
// with the ability to set a timeout on a request.
func NewAuditInformationLookupClusterParamsWithTimeout(timeout time.Duration) *AuditInformationLookupClusterParams {
	return &AuditInformationLookupClusterParams{
		timeout: timeout,
	}
}

// NewAuditInformationLookupClusterParamsWithContext creates a new AuditInformationLookupClusterParams object
// with the ability to set a context for a request.
func NewAuditInformationLookupClusterParamsWithContext(ctx context.Context) *AuditInformationLookupClusterParams {
	return &AuditInformationLookupClusterParams{
		Context: ctx,
	}
}

// NewAuditInformationLookupClusterParamsWithHTTPClient creates a new AuditInformationLookupClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuditInformationLookupClusterParamsWithHTTPClient(client *http.Client) *AuditInformationLookupClusterParams {
	return &AuditInformationLookupClusterParams{
		HTTPClient: client,
	}
}

/* AuditInformationLookupClusterParams contains all the parameters to send to the API endpoint
   for the audit information lookup cluster operation.

   Typically these are written to a http.Request.
*/
type AuditInformationLookupClusterParams struct {

	// ClusterSNI.
	ClusterSNI *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the audit information lookup cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditInformationLookupClusterParams) WithDefaults() *AuditInformationLookupClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the audit information lookup cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditInformationLookupClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the audit information lookup cluster params
func (o *AuditInformationLookupClusterParams) WithTimeout(timeout time.Duration) *AuditInformationLookupClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the audit information lookup cluster params
func (o *AuditInformationLookupClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the audit information lookup cluster params
func (o *AuditInformationLookupClusterParams) WithContext(ctx context.Context) *AuditInformationLookupClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the audit information lookup cluster params
func (o *AuditInformationLookupClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the audit information lookup cluster params
func (o *AuditInformationLookupClusterParams) WithHTTPClient(client *http.Client) *AuditInformationLookupClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the audit information lookup cluster params
func (o *AuditInformationLookupClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterSNI adds the clusterSNI to the audit information lookup cluster params
func (o *AuditInformationLookupClusterParams) WithClusterSNI(clusterSNI *string) *AuditInformationLookupClusterParams {
	o.SetClusterSNI(clusterSNI)
	return o
}

// SetClusterSNI adds the clusterSNI to the audit information lookup cluster params
func (o *AuditInformationLookupClusterParams) SetClusterSNI(clusterSNI *string) {
	o.ClusterSNI = clusterSNI
}

// WriteToRequest writes these params to a swagger request
func (o *AuditInformationLookupClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClusterSNI != nil {

		// query param clusterSNI
		var qrClusterSNI string

		if o.ClusterSNI != nil {
			qrClusterSNI = *o.ClusterSNI
		}
		qClusterSNI := qrClusterSNI
		if qClusterSNI != "" {

			if err := r.SetQueryParam("clusterSNI", qClusterSNI); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}