/*
Infra API

Infra REST API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ProviderRequest struct for ProviderRequest
type ProviderRequest struct {
	Kind         ProviderKind  `json:"kind"`
	Domain       string        `json:"domain" binding:"fqdn,required"`
	ClientID     string        `json:"clientID"`
	ClientSecret string        `json:"clientSecret"`
	Okta         *ProviderOkta `json:"okta,omitempty"`
}

// NewProviderRequest instantiates a new ProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderRequest(kind ProviderKind, domain string, clientID string, clientSecret string) *ProviderRequest {
	this := ProviderRequest{}
	this.Kind = kind
	this.Domain = domain
	this.ClientID = clientID
	this.ClientSecret = clientSecret
	return &this
}

// NewProviderRequestWithDefaults instantiates a new ProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderRequestWithDefaults() *ProviderRequest {
	this := ProviderRequest{}
	return &this
}

// GetKind returns the Kind field value
func (o *ProviderRequest) GetKind() ProviderKind {
	if o == nil {
		var ret ProviderKind
		return ret
	}

	return o.Kind
}

// GetKindOK returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ProviderRequest) GetKindOK() (*ProviderKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ProviderRequest) SetKind(v ProviderKind) {
	o.Kind = v
}

// GetDomain returns the Domain field value
func (o *ProviderRequest) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOK returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *ProviderRequest) GetDomainOK() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *ProviderRequest) SetDomain(v string) {
	o.Domain = v
}

// GetClientID returns the ClientID field value
func (o *ProviderRequest) GetClientID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientID
}

// GetClientIDOK returns a tuple with the ClientID field value
// and a boolean to check if the value has been set.
func (o *ProviderRequest) GetClientIDOK() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientID, true
}

// SetClientID sets field value
func (o *ProviderRequest) SetClientID(v string) {
	o.ClientID = v
}

// GetClientSecret returns the ClientSecret field value
func (o *ProviderRequest) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOK returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *ProviderRequest) GetClientSecretOK() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *ProviderRequest) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetOkta returns the Okta field value if set, zero value otherwise.
func (o *ProviderRequest) GetOkta() ProviderOkta {
	if o == nil || o.Okta == nil {
		var ret ProviderOkta
		return ret
	}
	return *o.Okta
}

// GetOktaOK returns a tuple with the Okta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderRequest) GetOktaOK() (*ProviderOkta, bool) {
	if o == nil || o.Okta == nil {
		return nil, false
	}
	return o.Okta, true
}

// HasOkta returns a boolean if a field has been set.
func (o *ProviderRequest) HasOkta() bool {
	if o != nil && o.Okta != nil {
		return true
	}

	return false
}

// SetOkta gets a reference to the given ProviderOkta and assigns it to the Okta field.
func (o *ProviderRequest) SetOkta(v ProviderOkta) {
	o.Okta = &v
}

func (o ProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if true {
		toSerialize["clientID"] = o.ClientID
	}
	if true {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if o.Okta != nil {
		toSerialize["okta"] = o.Okta
	}
	return json.Marshal(toSerialize)
}

type NullableProviderRequest struct {
	value *ProviderRequest
	isSet bool
}

func (v NullableProviderRequest) Get() *ProviderRequest {
	return v.value
}

func (v *NullableProviderRequest) Set(val *ProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderRequest(val *ProviderRequest) *NullableProviderRequest {
	return &NullableProviderRequest{value: val, isSet: true}
}

func (v NullableProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}