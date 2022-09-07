// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
	"toucan-leaderboard/ent/schema"
	"toucan-leaderboard/ent/tgocache"
	"toucan-leaderboard/ent/tgoens"
	"toucan-leaderboard/ent/tgoretirement"
	"toucan-leaderboard/ent/tuser"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	tgocacheFields := schema.TGoCache{}.Fields()
	_ = tgocacheFields
	// tgocacheDescCacheValue is the schema descriptor for cache_value field.
	tgocacheDescCacheValue := tgocacheFields[2].Descriptor()
	// tgocache.DefaultCacheValue holds the default value on creation for the cache_value field.
	tgocache.DefaultCacheValue = tgocacheDescCacheValue.Default.(string)
	// tgocacheDescMtime is the schema descriptor for mtime field.
	tgocacheDescMtime := tgocacheFields[3].Descriptor()
	// tgocache.DefaultMtime holds the default value on creation for the mtime field.
	tgocache.DefaultMtime = tgocacheDescMtime.Default.(time.Time)
	// tgocache.UpdateDefaultMtime holds the default value on update for the mtime field.
	tgocache.UpdateDefaultMtime = tgocacheDescMtime.UpdateDefault.(func() time.Time)
	// tgocacheDescCtime is the schema descriptor for ctime field.
	tgocacheDescCtime := tgocacheFields[4].Descriptor()
	// tgocache.DefaultCtime holds the default value on creation for the ctime field.
	tgocache.DefaultCtime = tgocacheDescCtime.Default.(time.Time)
	tgoensFields := schema.TGoEns{}.Fields()
	_ = tgoensFields
	// tgoensDescMtime is the schema descriptor for mtime field.
	tgoensDescMtime := tgoensFields[3].Descriptor()
	// tgoens.DefaultMtime holds the default value on creation for the mtime field.
	tgoens.DefaultMtime = tgoensDescMtime.Default.(time.Time)
	// tgoens.UpdateDefaultMtime holds the default value on update for the mtime field.
	tgoens.UpdateDefaultMtime = tgoensDescMtime.UpdateDefault.(func() time.Time)
	// tgoensDescCtime is the schema descriptor for ctime field.
	tgoensDescCtime := tgoensFields[4].Descriptor()
	// tgoens.DefaultCtime holds the default value on creation for the ctime field.
	tgoens.DefaultCtime = tgoensDescCtime.Default.(time.Time)
	tgoretirementFields := schema.TGoRetirement{}.Fields()
	_ = tgoretirementFields
	// tgoretirementDescCreationTx is the schema descriptor for creation_tx field.
	tgoretirementDescCreationTx := tgoretirementFields[1].Descriptor()
	// tgoretirement.DefaultCreationTx holds the default value on creation for the creation_tx field.
	tgoretirement.DefaultCreationTx = tgoretirementDescCreationTx.Default.(string)
	// tgoretirementDescCreatorAddress is the schema descriptor for creator_address field.
	tgoretirementDescCreatorAddress := tgoretirementFields[2].Descriptor()
	// tgoretirement.DefaultCreatorAddress holds the default value on creation for the creator_address field.
	tgoretirement.DefaultCreatorAddress = tgoretirementDescCreatorAddress.Default.(string)
	// tgoretirementDescBeneficiaryAddress is the schema descriptor for beneficiary_address field.
	tgoretirementDescBeneficiaryAddress := tgoretirementFields[3].Descriptor()
	// tgoretirement.DefaultBeneficiaryAddress holds the default value on creation for the beneficiary_address field.
	tgoretirement.DefaultBeneficiaryAddress = tgoretirementDescBeneficiaryAddress.Default.(string)
	// tgoretirementDescAmount is the schema descriptor for amount field.
	tgoretirementDescAmount := tgoretirementFields[4].Descriptor()
	// tgoretirement.DefaultAmount holds the default value on creation for the amount field.
	tgoretirement.DefaultAmount = tgoretirementDescAmount.Default.(float64)
	// tgoretirementDescTokenAddress is the schema descriptor for token_address field.
	tgoretirementDescTokenAddress := tgoretirementFields[5].Descriptor()
	// tgoretirement.DefaultTokenAddress holds the default value on creation for the token_address field.
	tgoretirement.DefaultTokenAddress = tgoretirementDescTokenAddress.Default.(string)
	// tgoretirementDescTokenName is the schema descriptor for token_name field.
	tgoretirementDescTokenName := tgoretirementFields[6].Descriptor()
	// tgoretirement.DefaultTokenName holds the default value on creation for the token_name field.
	tgoretirement.DefaultTokenName = tgoretirementDescTokenName.Default.(string)
	// tgoretirementDescTokenType is the schema descriptor for token_type field.
	tgoretirementDescTokenType := tgoretirementFields[7].Descriptor()
	// tgoretirement.DefaultTokenType holds the default value on creation for the token_type field.
	tgoretirement.DefaultTokenType = tgoretirementDescTokenType.Default.(string)
	// tgoretirementDescMtime is the schema descriptor for mtime field.
	tgoretirementDescMtime := tgoretirementFields[9].Descriptor()
	// tgoretirement.DefaultMtime holds the default value on creation for the mtime field.
	tgoretirement.DefaultMtime = tgoretirementDescMtime.Default.(time.Time)
	// tgoretirement.UpdateDefaultMtime holds the default value on update for the mtime field.
	tgoretirement.UpdateDefaultMtime = tgoretirementDescMtime.UpdateDefault.(func() time.Time)
	// tgoretirementDescCtime is the schema descriptor for ctime field.
	tgoretirementDescCtime := tgoretirementFields[10].Descriptor()
	// tgoretirement.DefaultCtime holds the default value on creation for the ctime field.
	tgoretirement.DefaultCtime = tgoretirementDescCtime.Default.(time.Time)
	tuserFields := schema.TUser{}.Fields()
	_ = tuserFields
	// tuserDescLastLoginTime is the schema descriptor for last_login_time field.
	tuserDescLastLoginTime := tuserFields[9].Descriptor()
	// tuser.DefaultLastLoginTime holds the default value on creation for the last_login_time field.
	tuser.DefaultLastLoginTime = tuserDescLastLoginTime.Default.(time.Time)
	// tuser.UpdateDefaultLastLoginTime holds the default value on update for the last_login_time field.
	tuser.UpdateDefaultLastLoginTime = tuserDescLastLoginTime.UpdateDefault.(func() time.Time)
	// tuserDescMtime is the schema descriptor for mtime field.
	tuserDescMtime := tuserFields[10].Descriptor()
	// tuser.DefaultMtime holds the default value on creation for the mtime field.
	tuser.DefaultMtime = tuserDescMtime.Default.(time.Time)
	// tuser.UpdateDefaultMtime holds the default value on update for the mtime field.
	tuser.UpdateDefaultMtime = tuserDescMtime.UpdateDefault.(func() time.Time)
	// tuserDescCtime is the schema descriptor for ctime field.
	tuserDescCtime := tuserFields[11].Descriptor()
	// tuser.DefaultCtime holds the default value on creation for the ctime field.
	tuser.DefaultCtime = tuserDescCtime.Default.(time.Time)
}
