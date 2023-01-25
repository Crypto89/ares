package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/crypto89/ares/api/models"
	"github.com/crypto89/ares/api/restapi"
	"github.com/crypto89/ares/api/restapi/operations"
	"github.com/crypto89/go-ogetit"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"

	player_ops "github.com/crypto89/ares/api/restapi/operations/player"
)

type API struct {
	apiKey string
	// backend ares.Backend
	Handler http.Handler
	client  ogetit.Client
}

func NewAPI(apiKey string) (*API, error) {
	client, err := ogetit.NewCachedClient(apiKey, 30*time.Minute, 24*time.Hour)
	if err != nil {
		return nil, err
	}
	api := &API{
		client: client,
	}

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return nil, err
	}

	openAPI := operations.NewAresAPI(swaggerSpec)
	openAPI.Middleware = func(b middleware.Builder) http.Handler {
		return middleware.Spec("", swaggerSpec.Raw(), openAPI.Context().RoutesHandler(b))
	}

	// add the handlers here
	openAPI.PlayerGetPlayerHandler = player_ops.GetPlayerHandlerFunc(api.getPlayerHandler)

	api.Handler = openAPI.Serve(nil)

	return api, nil
}

func (api *API) getPlayerHandler(params player_ops.GetPlayerParams) middleware.Responder {
	start := time.Now()

	_, community, universeID, _, err := ogetit.ParseKey(*params.Body.Key)
	if err != nil {
		logrus.WithError(err).Warn("bad report key")
		return nil
	}

	serverData, err := api.client.GetServerData(universeID, community)
	if err != nil {
		logrus.WithError(err).Warn("failed to get server data")
		return player_ops.NewGetPlayerServiceUnavailable().WithPayload("OGame API Unavailable")
	}

	report, err := api.client.GetSpyReport(*params.Body.Key)
	if err != nil {
		logrus.WithError(err).Warn("failed to get report data")
		return player_ops.NewGetPlayerServiceUnavailable().WithPayload("OGame API Unavailable")
	}

	if report.ResultCode == 6001 {
		return player_ops.NewGetPlayerNotFound()
	}

	if report.ResultCode != 1000 {
		logrus.Warnf("failed to get report: %d", report.ResultCode)
		return player_ops.NewGetPlayerServiceUnavailable().WithPayload("OGame API Unavailable")
	}

	spyReport := models.PlayerResponse{
		Fleet:  *params.Body.Fleet,
		Party:  *params.Body.Party,
		Data:   spyReportToOpenAPI(report),
		Server: serverDataToOpenAPI(serverData),
		Time:   time.Since(start).Seconds(),
	}

	return player_ops.NewGetPlayerCreated().WithPayload(spyReport)
}

func spyReportToOpenAPI(sr *ogetit.SpyReport) *models.SpyReport {
	attackerGalaxy, attackerSystem, attackerPosition := splitCoordinates(sr.Data.Generic.AttackerPlanetCoordinates)
	attacker := &models.PlayerData{
		Alliance: &models.PlayerDataAlliance{
			Name: sr.Data.Generic.AttackerAllianceName,
			Tag:  sr.Data.Generic.AttackerAllianceTag,
		},
		Class: sr.Data.Generic.AttackerCharacterClassID,
		ID:    sr.Data.Generic.AttackerUserID,
		Name:  sr.Data.Generic.AttackerName,
		Planet: &models.Planet{
			Coordinates: sr.Data.Generic.AttackerPlanetCoordinates,
			Galaxy:      attackerGalaxy,
			System:      attackerSystem,
			Position:    attackerPosition,
			Name:        sr.Data.Generic.AttackerPlanetName,
			Type:        sr.Data.Generic.AttackerPlanetType,
		},
	}

	defenderGalaxy, defenderSystem, defenderPosition := splitCoordinates(sr.Data.Generic.DefenderPlanetCoordinates)
	defender := &models.PlayerData{
		Alliance: &models.PlayerDataAlliance{
			Name: sr.Data.Generic.DefenderAllianceName,
			Tag:  sr.Data.Generic.DefenderAllianceTag,
		},
		Class: sr.Data.Generic.DefenderCharacterClassID,
		ID:    sr.Data.Generic.DefenderUserID,
		Name:  sr.Data.Generic.DefenderName,
		Planet: &models.Planet{
			Coordinates: sr.Data.Generic.DefenderPlanetCoordinates,
			Galaxy:      defenderGalaxy,
			System:      defenderSystem,
			Position:    defenderPosition,
			Name:        sr.Data.Generic.DefenderPlanetName,
			Type:        sr.Data.Generic.DefenderPlanetType,
		},
		Buildings: buildingsToOpenAPI(sr.Data.Details.Buildings),
		Research:  researchToOpenAPI(sr.Data.Details.Research),
		Defence:   defenceToOpenAPI(sr.Data.Details.Defence),
		Ships:     shipsToOpenAPI(sr.Data.Details.Ships),
		Resources: &models.PlayerDataResources{
			Metal:     sr.Data.Details.Resources.Metal,
			Crystal:   sr.Data.Details.Resources.Crystal,
			Deuterium: sr.Data.Details.Resources.Deuterium,
			Energy:    0,
		},
	}

	return &models.SpyReport{
		Activity:          sr.Data.Generic.Activity,
		Attacker:          attacker,
		Defender:          defender,
		ID:                sr.Data.Generic.ID,
		LootPercentage:    sr.Data.Generic.LootPercentage,
		SpyFailChance:     sr.Data.Generic.SpyChanceFail,
		Time:              sr.Data.Generic.EventTime.Format(time.RFC3339),
		Timestamp:         sr.Data.Generic.EventTimestamp,
		TotalDefenceCount: sr.Data.Generic.TotalDefenceCount,
		TotalShipCount:    sr.Data.Generic.TotalShipCount,
	}
}

func buildingsToOpenAPI(buildings []ogetit.Building) models.LevelableMap {
	result := make(map[string]models.Levelable)
	for _, v := range buildings {
		result[fmt.Sprintf("%d", v.BuildingType)] = models.Levelable{Level: v.Level}
	}

	return result
}

func defenceToOpenAPI(defence []ogetit.Defence) models.CountableMap {
	result := make(map[string]models.Countable)
	for _, v := range defence {
		result[fmt.Sprintf("%d", v.DefenceType)] = models.Countable{Count: v.Count}
	}

	return result
}

func researchToOpenAPI(research []ogetit.Research) models.LevelableMap {
	result := make(map[string]models.Levelable)
	for _, v := range research {
		result[fmt.Sprintf("%d", v.ResearchType)] = models.Levelable{Level: v.Level}
	}

	return result
}

func shipsToOpenAPI(ships []ogetit.Ship) models.CountableMap {
	result := make(map[string]models.Countable)
	for _, v := range ships {
		result[fmt.Sprintf("%d", v.ShipType)] = models.Countable{Count: int64(v.Count)}
	}

	return result
}

func splitCoordinates(coordinates string) (int64, int64, int64) {
	coords := strings.SplitN(coordinates, ":", 3)

	galaxy, _ := strconv.Atoi(coords[0])
	system, _ := strconv.Atoi(coords[1])
	position, _ := strconv.Atoi(coords[2])

	return int64(galaxy), int64(system), int64(position)
}

func serverDataToOpenAPI(sd *ogetit.ServerData) *models.ServerData {
	result := &models.ServerData{
		Name:                          sd.Name,
		Number:                        sd.Number,
		Language:                      sd.Language,
		Timezone:                      sd.Timezone,
		TimezoneOffset:                sd.TimezoneOffset,
		Domain:                        sd.Domain,
		Version:                       sd.Version,
		Speed:                         sd.Speed,
		SpeedFleet:                    sd.SpeedFleetWar,
		Galaxies:                      sd.Galaxies,
		Systems:                       sd.Systems,
		Acs:                           sd.ACS,
		RapidFire:                     sd.RapidFire,
		DefTotF:                       sd.DefToTF,
		DebrisFactor:                  sd.DebrisFactor,
		DebrisFactorDef:               sd.DebrisFactorDef,
		RepairFactor:                  sd.RepairFactor,
		NewbieProtectionLimit:         sd.NewbieProtectionLimit,
		NewbieProtectionHigh:          sd.NewbieProtectionHigh,
		TopScore:                      sd.TopScore,
		BonusFields:                   sd.BonusFields,
		DonutGalaxy:                   sd.DonutGalaxy,
		DonutSystem:                   sd.DonutSystem,
		WfEnabled:                     sd.WFEnabled,
		WfMinResLost:                  sd.WFMinimumRessLost,
		WfMinLossPercent:              sd.WFMinimumLossPercentage,
		WfBasicPercentRepairable:      sd.WFBasicPercentageRepairable,
		GlobalDeuteriumSaveFactor:     sd.GlobalDeuteriumSaveFactor,
		BashLimit:                     sd.BashLimit,
		ProbeCargo:                    sd.ProbeCargo,
		ResearchDurationDivisor:       sd.ResearchDurationDivisor,
		DarkMatterNewAcount:           sd.DarkMatterNewAcount,
		CargoHyperspaceTechMultiplier: sd.CargoHyperspaceTechMultiplier,
		CharacterClassesEnabled:       sd.CharacterClassesEnabled,
		MinerBonusResourceProduction:  sd.MinerBonusResourceProduction,
		MinerBonusFasterTradingShips:  sd.MinerBonusFasterTradingShips,
		MinerBonusIncreasedCargoCapacityForTradingShips: sd.MinerBonusIncreasedCargoCapacityForTradingShips,
		MinerBonusAdditionalFleetSlots:                  sd.MinerBonusAdditionalFleetSlots,
		ResourceBuggyProductionBoost:                    sd.ResourceBuggyProductionBoost,
		ResourceBuggyMaxProductionBoost:                 sd.ResourceBuggyMaxProductionBoost,
		ResourceBuggyEnergyConsumptionPerUnit:           sd.ResourceBuggyEnergyConsumptionPerUnit,
		WarriorBonusFasterCombatShips:                   sd.WarriorBonusFasterCombatShips,
		WarriorBonusFasterRecyclers:                     sd.WarriorBonusFasterRecyclers,
		WarriorBonusRecyclerFuelConsumption:             sd.WarriorBonusFuelConsumption,
		CombatDebrisFieldLimit:                          sd.CombatDebrisFieldLimit,
		ExplorerBonusIncreasedResearchSpeed:             sd.ExplorerBonusIncreasedResearchSpeed,
		ExplorerBonusIncreasedExpeditionOutcome:         sd.ExplorerBonusIncreasedExpeditionOutcome,
		ExplorerBonusLargerPlanets:                      sd.ExplorerBonusLargerPlanets,
		ExplorerUnitItemsPerDay:                         sd.ExplorerUnitItemsPerDay,
	}

	return result
}
