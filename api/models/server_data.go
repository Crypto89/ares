// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServerData server data
//
// swagger:model ServerData
type ServerData struct {

	// acs
	Acs bool `json:"acs,omitempty"`

	// bash limit
	BashLimit int64 `json:"bash_limit,omitempty"`

	// bonus fields
	BonusFields int64 `json:"bonus_fields,omitempty"`

	// cargo hyperspace tech multiplier
	CargoHyperspaceTechMultiplier int64 `json:"cargo_hyperspace_tech_multiplier,omitempty"`

	// character classes enabled
	CharacterClassesEnabled bool `json:"characterClassesEnabled,omitempty"`

	// combat debris field limit
	CombatDebrisFieldLimit float64 `json:"combatDebrisFieldLimit,omitempty"`

	// dark matter new acount
	DarkMatterNewAcount int64 `json:"dark_matter_new_acount,omitempty"`

	// debris factor
	DebrisFactor float64 `json:"debris_factor,omitempty"`

	// debris factor def
	DebrisFactorDef float64 `json:"debris_factor_def,omitempty"`

	// def to t f
	DefTotF bool `json:"def_to_tF,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`

	// donut galaxy
	DonutGalaxy bool `json:"donut_galaxy,omitempty"`

	// donut system
	DonutSystem bool `json:"donut_system,omitempty"`

	// explorer bonus increased expedition outcome
	ExplorerBonusIncreasedExpeditionOutcome float64 `json:"explorerBonusIncreasedExpeditionOutcome,omitempty"`

	// explorer bonus increased research speed
	ExplorerBonusIncreasedResearchSpeed float64 `json:"explorerBonusIncreasedResearchSpeed,omitempty"`

	// explorer bonus larger planets
	ExplorerBonusLargerPlanets float64 `json:"explorerBonusLargerPlanets,omitempty"`

	// explorer unit items per day
	ExplorerUnitItemsPerDay int64 `json:"explorerUnitItemsPerDay,omitempty"`

	// galaxies
	Galaxies int64 `json:"galaxies,omitempty"`

	// global deuterium save factor
	GlobalDeuteriumSaveFactor float64 `json:"global_deuterium_save_factor,omitempty"`

	// language
	Language string `json:"language,omitempty"`

	// miner bonus additional fleet slots
	MinerBonusAdditionalFleetSlots int64 `json:"minerBonusAdditionalFleetSlots,omitempty"`

	// miner bonus faster trading ships
	MinerBonusFasterTradingShips bool `json:"minerBonusFasterTradingShips,omitempty"`

	// miner bonus increased cargo capacity for trading ships
	MinerBonusIncreasedCargoCapacityForTradingShips float64 `json:"minerBonusIncreasedCargoCapacityForTradingShips,omitempty"`

	// miner bonus resource production
	MinerBonusResourceProduction float64 `json:"minerBonusResourceProduction,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// newbie protection high
	NewbieProtectionHigh int64 `json:"newbie_protection_high,omitempty"`

	// newbie protection limit
	NewbieProtectionLimit int64 `json:"newbie_protection_limit,omitempty"`

	// number
	Number int64 `json:"number,omitempty"`

	// probe cargo
	ProbeCargo bool `json:"probe_cargo,omitempty"`

	// rapid fire
	RapidFire bool `json:"rapid_fire,omitempty"`

	// repair factor
	RepairFactor float64 `json:"repair_factor,omitempty"`

	// research duration divisor
	ResearchDurationDivisor int64 `json:"research_duration_divisor,omitempty"`

	// resource buggy energy consumption per unit
	ResourceBuggyEnergyConsumptionPerUnit int64 `json:"resourceBuggyEnergyConsumptionPerUnit,omitempty"`

	// resource buggy max production boost
	ResourceBuggyMaxProductionBoost float64 `json:"resourceBuggyMaxProductionBoost,omitempty"`

	// resource buggy production boost
	ResourceBuggyProductionBoost float64 `json:"resourceBuggyProductionBoost,omitempty"`

	// speed
	Speed int64 `json:"speed,omitempty"`

	// speed fleet
	SpeedFleet int64 `json:"speed_fleet,omitempty"`

	// systems
	Systems int64 `json:"systems,omitempty"`

	// timezone
	Timezone string `json:"timezone,omitempty"`

	// timezone offset
	TimezoneOffset string `json:"timezone_offset,omitempty"`

	// top score
	TopScore float64 `json:"top_score,omitempty"`

	// version
	Version string `json:"version,omitempty"`

	// warrior bonus faster combat ships
	WarriorBonusFasterCombatShips bool `json:"warriorBonusFasterCombatShips,omitempty"`

	// warrior bonus faster recyclers
	WarriorBonusFasterRecyclers bool `json:"warriorBonusFasterRecyclers,omitempty"`

	// warrior bonus recycler fuel consumption
	WarriorBonusRecyclerFuelConsumption float64 `json:"warriorBonusRecyclerFuelConsumption,omitempty"`

	// wf basic percent repairable
	WfBasicPercentRepairable int64 `json:"wf_basic_percent_repairable,omitempty"`

	// wf enabled
	WfEnabled bool `json:"wf_enabled,omitempty"`

	// wf min loss percent
	WfMinLossPercent int64 `json:"wf_min_loss_percent,omitempty"`

	// wf min res lost
	WfMinResLost int64 `json:"wf_min_res_lost,omitempty"`
}

// Validate validates this server data
func (m *ServerData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this server data based on context it is used
func (m *ServerData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerData) UnmarshalBinary(b []byte) error {
	var res ServerData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
