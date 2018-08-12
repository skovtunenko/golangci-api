// Code generated by go-queryset. DO NOT EDIT.
package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set RepoAnalysisStatusQuerySet

// RepoAnalysisStatusQuerySet is an queryset type for RepoAnalysisStatus
type RepoAnalysisStatusQuerySet struct {
	db *gorm.DB
}

// NewRepoAnalysisStatusQuerySet constructs new RepoAnalysisStatusQuerySet
func NewRepoAnalysisStatusQuerySet(db *gorm.DB) RepoAnalysisStatusQuerySet {
	return RepoAnalysisStatusQuerySet{
		db: db.Model(&RepoAnalysisStatus{}),
	}
}

func (qs RepoAnalysisStatusQuerySet) w(db *gorm.DB) RepoAnalysisStatusQuerySet {
	return NewRepoAnalysisStatusQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) All(ret *[]RepoAnalysisStatus) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *RepoAnalysisStatus) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) CreatedAtEq(createdAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) CreatedAtGt(createdAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) CreatedAtGte(createdAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) CreatedAtLt(createdAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) CreatedAtLte(createdAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) CreatedAtNe(createdAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *RepoAnalysisStatus) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) Delete() error {
	return qs.db.Delete(RepoAnalysisStatus{}).Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) DeletedAtEq(deletedAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) DeletedAtGt(deletedAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) DeletedAtGte(deletedAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) DeletedAtIsNotNull() RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) DeletedAtIsNull() RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) DeletedAtLt(deletedAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) DeletedAtLte(deletedAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) DeletedAtNe(deletedAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) GetUpdater() RepoAnalysisStatusUpdater {
	return NewRepoAnalysisStatusUpdater(qs.db)
}

// HasPendingChangesEq is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) HasPendingChangesEq(hasPendingChanges bool) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("has_pending_changes = ?", hasPendingChanges))
}

// HasPendingChangesIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) HasPendingChangesIn(hasPendingChanges ...bool) RepoAnalysisStatusQuerySet {
	if len(hasPendingChanges) == 0 {
		qs.db.AddError(errors.New("must at least pass one hasPendingChanges in HasPendingChangesIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("has_pending_changes IN (?)", hasPendingChanges))
}

// HasPendingChangesNe is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) HasPendingChangesNe(hasPendingChanges bool) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("has_pending_changes != ?", hasPendingChanges))
}

// HasPendingChangesNotIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) HasPendingChangesNotIn(hasPendingChanges ...bool) RepoAnalysisStatusQuerySet {
	if len(hasPendingChanges) == 0 {
		qs.db.AddError(errors.New("must at least pass one hasPendingChanges in HasPendingChangesNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("has_pending_changes NOT IN (?)", hasPendingChanges))
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) IDEq(ID uint) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) IDGt(ID uint) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) IDGte(ID uint) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) IDIn(ID ...uint) RepoAnalysisStatusQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) IDLt(ID uint) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) IDLte(ID uint) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) IDNe(ID uint) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) IDNotIn(ID ...uint) RepoAnalysisStatusQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// LastAnalyzedAtEq is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) LastAnalyzedAtEq(lastAnalyzedAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("last_analyzed_at = ?", lastAnalyzedAt))
}

// LastAnalyzedAtGt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) LastAnalyzedAtGt(lastAnalyzedAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("last_analyzed_at > ?", lastAnalyzedAt))
}

// LastAnalyzedAtGte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) LastAnalyzedAtGte(lastAnalyzedAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("last_analyzed_at >= ?", lastAnalyzedAt))
}

// LastAnalyzedAtLt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) LastAnalyzedAtLt(lastAnalyzedAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("last_analyzed_at < ?", lastAnalyzedAt))
}

// LastAnalyzedAtLte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) LastAnalyzedAtLte(lastAnalyzedAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("last_analyzed_at <= ?", lastAnalyzedAt))
}

// LastAnalyzedAtNe is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) LastAnalyzedAtNe(lastAnalyzedAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("last_analyzed_at != ?", lastAnalyzedAt))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) Limit(limit int) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// NameEq is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) NameEq(name string) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("name = ?", name))
}

// NameIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) NameIn(name ...string) RepoAnalysisStatusQuerySet {
	if len(name) == 0 {
		qs.db.AddError(errors.New("must at least pass one name in NameIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("name IN (?)", name))
}

// NameNe is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) NameNe(name string) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("name != ?", name))
}

// NameNotIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) NameNotIn(name ...string) RepoAnalysisStatusQuerySet {
	if len(name) == 0 {
		qs.db.AddError(errors.New("must at least pass one name in NameNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("name NOT IN (?)", name))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) Offset(offset int) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs RepoAnalysisStatusQuerySet) One(ret *RepoAnalysisStatus) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) OrderAscByCreatedAt() RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) OrderAscByDeletedAt() RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) OrderAscByID() RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByLastAnalyzedAt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) OrderAscByLastAnalyzedAt() RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Order("last_analyzed_at ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) OrderAscByUpdatedAt() RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderAscByVersion is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) OrderAscByVersion() RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Order("version ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) OrderDescByCreatedAt() RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) OrderDescByDeletedAt() RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) OrderDescByID() RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByLastAnalyzedAt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) OrderDescByLastAnalyzedAt() RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Order("last_analyzed_at DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) OrderDescByUpdatedAt() RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// OrderDescByVersion is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) OrderDescByVersion() RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Order("version DESC"))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u RepoAnalysisStatusUpdater) SetCreatedAt(createdAt time.Time) RepoAnalysisStatusUpdater {
	u.fields[string(RepoAnalysisStatusDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDeletedAt is an autogenerated method
// nolint: dupl
func (u RepoAnalysisStatusUpdater) SetDeletedAt(deletedAt *time.Time) RepoAnalysisStatusUpdater {
	u.fields[string(RepoAnalysisStatusDBSchema.DeletedAt)] = deletedAt
	return u
}

// SetHasPendingChanges is an autogenerated method
// nolint: dupl
func (u RepoAnalysisStatusUpdater) SetHasPendingChanges(hasPendingChanges bool) RepoAnalysisStatusUpdater {
	u.fields[string(RepoAnalysisStatusDBSchema.HasPendingChanges)] = hasPendingChanges
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u RepoAnalysisStatusUpdater) SetID(ID uint) RepoAnalysisStatusUpdater {
	u.fields[string(RepoAnalysisStatusDBSchema.ID)] = ID
	return u
}

// SetLastAnalyzedAt is an autogenerated method
// nolint: dupl
func (u RepoAnalysisStatusUpdater) SetLastAnalyzedAt(lastAnalyzedAt time.Time) RepoAnalysisStatusUpdater {
	u.fields[string(RepoAnalysisStatusDBSchema.LastAnalyzedAt)] = lastAnalyzedAt
	return u
}

// SetName is an autogenerated method
// nolint: dupl
func (u RepoAnalysisStatusUpdater) SetName(name string) RepoAnalysisStatusUpdater {
	u.fields[string(RepoAnalysisStatusDBSchema.Name)] = name
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u RepoAnalysisStatusUpdater) SetUpdatedAt(updatedAt time.Time) RepoAnalysisStatusUpdater {
	u.fields[string(RepoAnalysisStatusDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SetVersion is an autogenerated method
// nolint: dupl
func (u RepoAnalysisStatusUpdater) SetVersion(version int) RepoAnalysisStatusUpdater {
	u.fields[string(RepoAnalysisStatusDBSchema.Version)] = version
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u RepoAnalysisStatusUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u RepoAnalysisStatusUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) UpdatedAtEq(updatedAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) UpdatedAtGt(updatedAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) UpdatedAtGte(updatedAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) UpdatedAtLt(updatedAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) UpdatedAtLte(updatedAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) UpdatedAtNe(updatedAt time.Time) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// VersionEq is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) VersionEq(version int) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("version = ?", version))
}

// VersionGt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) VersionGt(version int) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("version > ?", version))
}

// VersionGte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) VersionGte(version int) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("version >= ?", version))
}

// VersionIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) VersionIn(version ...int) RepoAnalysisStatusQuerySet {
	if len(version) == 0 {
		qs.db.AddError(errors.New("must at least pass one version in VersionIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("version IN (?)", version))
}

// VersionLt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) VersionLt(version int) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("version < ?", version))
}

// VersionLte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) VersionLte(version int) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("version <= ?", version))
}

// VersionNe is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) VersionNe(version int) RepoAnalysisStatusQuerySet {
	return qs.w(qs.db.Where("version != ?", version))
}

// VersionNotIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisStatusQuerySet) VersionNotIn(version ...int) RepoAnalysisStatusQuerySet {
	if len(version) == 0 {
		qs.db.AddError(errors.New("must at least pass one version in VersionNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("version NOT IN (?)", version))
}

// ===== END of query set RepoAnalysisStatusQuerySet

// ===== BEGIN of RepoAnalysisStatus modifiers

// RepoAnalysisStatusDBSchemaField describes database schema field. It requires for method 'Update'
type RepoAnalysisStatusDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f RepoAnalysisStatusDBSchemaField) String() string {
	return string(f)
}

// RepoAnalysisStatusDBSchema stores db field names of RepoAnalysisStatus
var RepoAnalysisStatusDBSchema = struct {
	ID                RepoAnalysisStatusDBSchemaField
	CreatedAt         RepoAnalysisStatusDBSchemaField
	UpdatedAt         RepoAnalysisStatusDBSchemaField
	DeletedAt         RepoAnalysisStatusDBSchemaField
	Name              RepoAnalysisStatusDBSchemaField
	LastAnalyzedAt    RepoAnalysisStatusDBSchemaField
	HasPendingChanges RepoAnalysisStatusDBSchemaField
	Version           RepoAnalysisStatusDBSchemaField
}{

	ID:                RepoAnalysisStatusDBSchemaField("id"),
	CreatedAt:         RepoAnalysisStatusDBSchemaField("created_at"),
	UpdatedAt:         RepoAnalysisStatusDBSchemaField("updated_at"),
	DeletedAt:         RepoAnalysisStatusDBSchemaField("deleted_at"),
	Name:              RepoAnalysisStatusDBSchemaField("name"),
	LastAnalyzedAt:    RepoAnalysisStatusDBSchemaField("last_analyzed_at"),
	HasPendingChanges: RepoAnalysisStatusDBSchemaField("has_pending_changes"),
	Version:           RepoAnalysisStatusDBSchemaField("version"),
}

// Update updates RepoAnalysisStatus fields by primary key
// nolint: dupl
func (o *RepoAnalysisStatus) Update(db *gorm.DB, fields ...RepoAnalysisStatusDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":                  o.ID,
		"created_at":          o.CreatedAt,
		"updated_at":          o.UpdatedAt,
		"deleted_at":          o.DeletedAt,
		"name":                o.Name,
		"last_analyzed_at":    o.LastAnalyzedAt,
		"has_pending_changes": o.HasPendingChanges,
		"version":             o.Version,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update RepoAnalysisStatus %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// RepoAnalysisStatusUpdater is an RepoAnalysisStatus updates manager
type RepoAnalysisStatusUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewRepoAnalysisStatusUpdater creates new RepoAnalysisStatus updater
// nolint: dupl
func NewRepoAnalysisStatusUpdater(db *gorm.DB) RepoAnalysisStatusUpdater {
	return RepoAnalysisStatusUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&RepoAnalysisStatus{}),
	}
}

// ===== END of RepoAnalysisStatus modifiers

// ===== END of all query sets
