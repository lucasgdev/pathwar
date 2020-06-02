package pwdb

import (
	"math/rand"

	"github.com/brianvoe/gofakeit"
	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"pathwar.land/v2/go/internal/randstring"
)

func GetInfo(db *gorm.DB, logger *zap.Logger) (*Info, error) {
	info := Info{
		TableRows: make(map[string]uint32),
	}
	for _, model := range All() {
		var count uint32
		tableName := db.NewScope(model).TableName()
		if err := db.Model(model).Count(&count).Error; err != nil {
			logger.Warn("get table rows", zap.String("table", tableName), zap.Error(err))
			continue
		}
		info.TableRows[tableName] = count
	}

	return &info, nil
}

func GetDump(db *gorm.DB) (*Dump, error) {
	dump := Dump{}
	if err := db.Find(&dump.Challenges).Error; err != nil {
		return nil, GormToErrcode(err)
	}
	if err := db.Find(&dump.ChallengeFlavors).Error; err != nil {
		return nil, GormToErrcode(err)
	}
	if err := db.Find(&dump.ChallengeInstances).Error; err != nil {
		return nil, GormToErrcode(err)
	}
	if err := db.Find(&dump.Agents).Error; err != nil {
		return nil, GormToErrcode(err)
	}
	if err := db.Find(&dump.Users).Error; err != nil {
		return nil, GormToErrcode(err)
	}
	if err := db.Find(&dump.Organizations).Error; err != nil {
		return nil, GormToErrcode(err)
	}
	if err := db.Find(&dump.OrganizationMembers).Error; err != nil {
		return nil, GormToErrcode(err)
	}
	if err := db.Find(&dump.Seasons).Error; err != nil {
		return nil, GormToErrcode(err)
	}
	if err := db.Find(&dump.Teams).Error; err != nil {
		return nil, GormToErrcode(err)
	}
	if err := db.Find(&dump.TeamMembers).Error; err != nil {
		return nil, GormToErrcode(err)
	}
	if err := db.Find(&dump.Coupons).Error; err != nil {
		return nil, GormToErrcode(err)
	}
	return &dump, nil
}

func GenerateFakeData(db *gorm.DB, sfn *snowflake.Node, logger *zap.Logger) error {
	//
	// agents
	//

	agents := []*Agent{}
	for i := 0; i < 3; i++ {
		agent := &Agent{
			Name:     gofakeit.HipsterWord(),
			Hostname: gofakeit.IPv4Address(),
			Status:   Agent_Active,
		}
		agents = append(agents, agent)
	}
	logger.Debug("Generating agents")
	for _, entity := range agents {
		if err := db.Set("gorm:association_autoupdate", true).Create(entity).Error; err != nil {
			return GormToErrcode(err)
		}
	}

	//
	// seasons
	//

	seasons := []*Season{}
	for i := 0; i < 3; i++ {
		season := &Season{
			Name:       gofakeit.HipsterWord(),
			Status:     Season_Started,
			Visibility: Season_Public,
			IsDefault:  false,
		}
		seasons = append(seasons, season)
	}
	seasons[0].IsDefault = true
	logger.Debug("Generating seasons")
	for _, entity := range seasons {
		if err := db.Set("gorm:association_autoupdate", true).Create(entity).Error; err != nil {
			return GormToErrcode(err)
		}
	}

	//
	// challenges
	//

	challenges := []*Challenge{}
	for i := 0; i < 5; i++ {
		challenge := &Challenge{
			Name:        gofakeit.HipsterWord(),
			Description: gofakeit.HipsterSentence(10),
			Author:      gofakeit.Name(),
			Locale:      "fr_FR",
			IsDraft:     false,
			Flavors:     []*ChallengeFlavor{},
		}
		for i := 0; i < 2; i++ {
			flavor := &ChallengeFlavor{
				Driver:           ChallengeFlavor_Docker,
				Version:          gofakeit.IPv4Address(),
				Changelog:        gofakeit.HipsterSentence(5),
				IsDraft:          false,
				IsLatest:         i == 0,
				SourceURL:        gofakeit.URL(),
				SeasonChallenges: []*SeasonChallenge{},
			}
			for j := 0; j < 2; j++ {
				seasonChallenge := &SeasonChallenge{
					SeasonID: seasons[rand.Int()%len(seasons)].ID,
				}
				flavor.SeasonChallenges = append(flavor.SeasonChallenges, seasonChallenge)
			}
			for j := 0; j < 2; j++ {
				instance := &ChallengeInstance{
					AgentID: agents[rand.Int()%len(agents)].ID,
					Status:  ChallengeInstance_Available,
				}
				flavor.Instances = append(flavor.Instances, instance)
			}
			challenge.Flavors = append(challenge.Flavors, flavor)
		}
		challenges = append(challenges, challenge)
	}

	logger.Debug("Generating challenges")
	for _, entity := range challenges {
		if err := db.Set("gorm:association_autoupdate", true).Create(entity).Error; err != nil {
			return GormToErrcode(err)
		}
	}

	//
	// organizations
	//

	organizations := []*Organization{}
	for i := 0; i < 5; i++ {
		organization := &Organization{
			Name:        gofakeit.HipsterWord(),
			GravatarURL: gofakeit.ImageURL(400, 400) + "?" + gofakeit.HipsterWord(),
			Locale:      "fr_FR",
		}
		organizations = append(organizations, organization)
	}

	logger.Debug("Generating organizations")
	for _, entity := range organizations {
		if err := db.Set("gorm:association_autoupdate", true).Create(entity).Error; err != nil {
			return GormToErrcode(err)
		}
	}

	//
	// users
	//

	users := []*User{}
	for i := 0; i < 10; i++ {
		user := &User{
			Username:                gofakeit.Name(),
			GravatarURL:             gofakeit.ImageURL(400, 400) + "?" + gofakeit.HipsterWord(),
			WebsiteURL:              gofakeit.URL(),
			Locale:                  "fr_FR",
			OrganizationMemberships: []*OrganizationMember{},
			OAuthSubject:            randstring.RandString(10),
		}
		users = append(users, user)
	}
	logger.Debug("Generating users")
	for _, entity := range users {
		if err := db.Set("gorm:association_autoupdate", true).Create(entity).Error; err != nil {
			return GormToErrcode(err)
		}
	}

	//
	// coupons
	//

	coupons := []*Coupon{}
	for i := 0; i < 3; i++ {
		coupon := &Coupon{
			Hash:               gofakeit.UUID(),
			MaxValidationCount: int32(rand.Int() % 5),
			Value:              int32(rand.Int() % 10),
			SeasonID:           seasons[rand.Int()%len(seasons)].ID,
		}
		coupons = append(coupons, coupon)
	}

	logger.Debug("Generating coupons")
	for _, entity := range coupons {
		if err := db.Set("gorm:association_autoupdate", true).Create(entity).Error; err != nil {
			return GormToErrcode(err)
		}
	}

	//
	// organization members
	//

	memberships := []*OrganizationMember{}
	for _, user := range users {
		for i := 0; i < 2; i++ {
			memberships = append(
				memberships,
				&OrganizationMember{
					OrganizationID: organizations[rand.Int()%len(organizations)].ID,
					UserID:         user.ID,
					Role:           OrganizationMember_Member, // or Owner
				},
			)
		}
	}

	logger.Debug("Generating memberships")
	for _, entity := range memberships {
		if err := db.Set("gorm:association_autoupdate", true).Create(entity).Error; err != nil {
			return GormToErrcode(err)
		}
	}

	return nil
}

func (m *SeasonChallenge) GetActiveSubscriptions() []*ChallengeSubscription {
	cs := make([]*ChallengeSubscription, 0)

	for _, subscription := range m.GetSubscriptions() {
		if subscription.GetStatus() == ChallengeSubscription_Active {
			cs = append(cs, subscription)
		}
	}

	return cs
}
