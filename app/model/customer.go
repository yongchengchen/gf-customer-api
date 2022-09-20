package model

import (
	"database/sql"
	"time"
)

type Customer struct {
	CustomersID               int            `db:"customers_id"`
	CustomerSeg               int            `db:"customer_seg"`
	IsGuest                   sql.NullInt64  `db:"is_guest"`
	CustomersGender           string         `db:"customers_gender"`
	CustomersFirstname        string         `db:"customers_firstname"`
	CustomersLastname         string         `db:"customers_lastname"`
	CustomersDob              time.Time      `db:"customers_dob"`
	CustomersEmailAddress     sql.NullString `db:"customers_email_address"`
	CustomersDefaultAddressID int            `db:"customers_default_address_id"`
	CustomersTelephone        string         `db:"customers_telephone"`
	CustomersFax              sql.NullString `db:"customers_fax"`
	CustomersPassword         string         `db:"customers_password"`
	CustomersNewsletter       sql.NullString `db:"customers_newsletter"`
	CustomersComment          sql.NullString `db:"customers_comment"`
	InternalID                sql.NullInt64  `db:"internalId"`
	ErpModified               sql.NullTime   `db:"erp_modified"`
	UniqueEmailID             sql.NullString `db:"unique_email_id"`
	Sns                       sql.NullInt64  `db:"sns"`
	DefaultDeliveryAddressID  sql.NullInt64  `db:"default_delivery_address_id"`
	EmailVerifiedAt           sql.NullTime   `db:"email_verified_at"`
}
