package pwapi

import (
	"context"
	"fmt"

	"pathwar.land/go/pkg/errcode"
	"pathwar.land/go/pkg/pwdb"
)

func (svc *service) CouponValidate(ctx context.Context, in *CouponValidate_Input) (*CouponValidate_Output, error) {
	// validation
	if in == nil || in.Hash == "" || in.TeamID == 0 {
		return nil, errcode.ErrMissingInput
	}

	userID, err := userIDFromContext(ctx, svc.db)
	if err != nil {
		return nil, fmt.Errorf("get userid from context: %w", err)
	}

	// check if user belongs to team
	// FIXME: or is admin
	var team pwdb.Team
	err = svc.db.
		Joins("JOIN team_member ON team_member.team_id = team.id AND team_member.user_id = ?", userID).
		Preload("Members").
		First(&team, in.TeamID).
		Error
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	}

	// fetch coupon
	var coupon pwdb.Coupon
	err = svc.db.
		Where(pwdb.Coupon{Hash: in.Hash}).
		First(&coupon).
		Error
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	}

	var validations int32
	err = svc.db.
		Model(&pwdb.CouponValidation{}).
		Where(&pwdb.CouponValidation{CouponID: coupon.ID}).
		Count(&validations).
		Error
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	}
	if validations >= coupon.MaxValidationCount {
		return nil, errcode.TODO.Wrap(err)
	}

	// FIXME: validate team
	// FIXME: already validated by this team
	// FIXME: inacitve user/team

	// create validation
	validation := pwdb.CouponValidation{
		Comment:  "xxx",
		AuthorID: userID,
		TeamID:   team.ID,
		CouponID: coupon.ID,
	}
	err = svc.db.Create(&validation).Error
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	}

	// FIXME: increment team.cash

	// load it again with preload
	err = svc.db.
		Preload("Team").
		Preload("Author").
		Preload("Coupon").
		First(&validation, validation.ID).
		Error
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	}

	ret := CouponValidate_Output{
		CouponValidation: &validation,
	}
	return &ret, nil
}