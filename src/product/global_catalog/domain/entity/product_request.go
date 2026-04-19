package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type RequestStatus string

const (
	RequestStatusPending   RequestStatus = "pending"
	RequestStatusApproved  RequestStatus = "approved"
	RequestStatusRejected  RequestStatus = "rejected"
	RequestStatusFulfilled RequestStatus = "fulfilled"
)

// ProductRequest representa una solicitud de un tenant para agregar un producto al catálogo global.
type ProductRequest struct {
	id              uuid.UUID
	tenantID        string
	name            string
	brand           *string
	category        *string
	description     *string
	businessType    *string
	status          RequestStatus
	adminNotes      *string
	globalProductID *uuid.UUID
	createdAt       time.Time
	updatedAt       time.Time
}

func NewProductRequest(tenantID, name string, brand, category, description, businessType *string) (*ProductRequest, error) {
	if tenantID == "" {
		return nil, errors.New("tenant_id es requerido")
	}
	if name == "" {
		return nil, errors.New("el nombre del producto es requerido")
	}

	return &ProductRequest{
		id:           uuid.New(),
		tenantID:     tenantID,
		name:         name,
		brand:        brand,
		category:     category,
		description:  description,
		businessType: businessType,
		status:       RequestStatusPending,
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
	}, nil
}

func NewProductRequestFromRepository(
	id uuid.UUID, tenantID, name string,
	brand, category, description, businessType *string,
	status RequestStatus, adminNotes *string, globalProductID *uuid.UUID,
	createdAt, updatedAt time.Time,
) *ProductRequest {
	return &ProductRequest{
		id: id, tenantID: tenantID, name: name,
		brand: brand, category: category, description: description,
		businessType: businessType, status: status, adminNotes: adminNotes,
		globalProductID: globalProductID, createdAt: createdAt, updatedAt: updatedAt,
	}
}

func (r *ProductRequest) ID() uuid.UUID            { return r.id }
func (r *ProductRequest) IDString() string          { return r.id.String() }
func (r *ProductRequest) TenantID() string          { return r.tenantID }
func (r *ProductRequest) Name() string              { return r.name }
func (r *ProductRequest) Brand() *string             { return r.brand }
func (r *ProductRequest) Category() *string          { return r.category }
func (r *ProductRequest) Description() *string       { return r.description }
func (r *ProductRequest) BusinessType() *string      { return r.businessType }
func (r *ProductRequest) Status() RequestStatus      { return r.status }
func (r *ProductRequest) AdminNotes() *string        { return r.adminNotes }
func (r *ProductRequest) GlobalProductID() *uuid.UUID { return r.globalProductID }
func (r *ProductRequest) CreatedAt() time.Time       { return r.createdAt }
func (r *ProductRequest) UpdatedAt() time.Time       { return r.updatedAt }

func (r *ProductRequest) Approve(notes string) {
	r.status = RequestStatusApproved
	r.adminNotes = &notes
	r.updatedAt = time.Now()
}

func (r *ProductRequest) Reject(notes string) {
	r.status = RequestStatusRejected
	r.adminNotes = &notes
	r.updatedAt = time.Now()
}

func (r *ProductRequest) Fulfill(globalProductID uuid.UUID) {
	r.status = RequestStatusFulfilled
	r.globalProductID = &globalProductID
	r.updatedAt = time.Now()
}
