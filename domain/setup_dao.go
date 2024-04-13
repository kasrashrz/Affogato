package domain

import (
	"github.com/kasrashrz/Affogato/domain/dto_only"
	"github.com/kasrashrz/Affogato/domain/dto_only/levels"
	"github.com/kasrashrz/Affogato/utils/errors"
)

func (setup *Setup) Setup() *errors.RestErr {
	var posts []Post
	var statuses []dto_only.Status
	var tactics []dto_only.Tactic
	var paymentDetails []PaymentDetail
	var doctorLevels []levels.DoctorLevel
	var talentFinderLevels []levels.TalentFinderLevel
	var fitnessCoachLevel []levels.FitnessCoachLevel
	var transportationLevel []levels.TransportationLevel
	var shoppingLevel []levels.ShoppingLevel
	var restaurantLevel []levels.RestaurantLevel
	var parkingLevel []levels.ParkingLevel

	db.Table("posts").Scan(&posts)
	db.Table("statuses").Scan(&statuses)
	db.Table("tactics").Scan(&tactics)
	db.Table("payment_details").Scan(&paymentDetails)
	db.Table("doctor_levels").Scan(&doctorLevels)
	db.Table("talent_finder_levels").Scan(&talentFinderLevels)
	db.Table("fitness_coach_levels").Scan(&fitnessCoachLevel)
	db.Table("transportation_levels").Scan(&transportationLevel)
	db.Table("shopping_levels").Scan(&shoppingLevel)
	db.Table("restaurant_levels").Scan(&restaurantLevel)
	db.Table("restaurant_levels").Scan(&parkingLevel)

	if len(posts) >= 0 {
		p1 := Post{}
		p2 := Post{}
		p3 := Post{}
		p4 := Post{}
		p1.ID = 1
		p2.ID = 2
		p3.ID = 3
		p4.ID = 4
		err := db.Save(&p1).Error
		err = db.Save(&p2).Error
		err = db.Save(&p3).Error
		err = db.Save(&p4).Error
		if err != nil {
			return errors.InternalServerError("something went wrong, post creation")
		}
		err = db.Table("posts").Where("id = 1").Updates(map[string]interface{}{"post_name": "F"}).Error
		err = db.Table("posts").Where("id = 2").Updates(map[string]interface{}{"post_name": "GK"}).Error
		err = db.Table("posts").Where("id = 3").Updates(map[string]interface{}{"post_name": "D"}).Error
		err = db.Table("posts").Where("id = 4").Updates(map[string]interface{}{"post_name": "M"}).Error
		err = db.Exec("delete from posts where id > 4").Error
		if err != nil {
			return errors.InternalServerError("something went wrong")
		}
	}
	if len(statuses) >= 0 {
		s1 := dto_only.Status{}
		s2 := dto_only.Status{}
		s3 := dto_only.Status{}

		s1.ID = 1
		s2.ID = 2
		s3.ID = 3

		err := db.Save(&s1).Error
		err = db.Save(&s2).Error
		err = db.Save(&s3).Error
		if err != nil {
			return errors.InternalServerError("something went wrong, status creation")
		}
		err = db.Table("statuses").Where("id = 1").Updates(map[string]interface{}{"status_name": "Bad"}).Error
		err = db.Table("statuses").Where("id = 2").Updates(map[string]interface{}{"status_name": "Happy"}).Error
		err = db.Table("statuses").Where("id = 3").Updates(map[string]interface{}{"status_name": "Injured"}).Error
		err = db.Exec("delete from statuses where id > 3").Error
		if err != nil {
			return errors.InternalServerError("something went wrong")
		}
	}
	if len(tactics) >= 0 {
		t1 := dto_only.Tactic{}
		t2 := dto_only.Tactic{}
		t3 := dto_only.Tactic{}
		t4 := dto_only.Tactic{}

		t1.ID = 1
		t2.ID = 2
		t3.ID = 3
		t4.ID = 4

		err := db.Save(&t1).Error
		err = db.Save(&t2).Error
		err = db.Save(&t3).Error
		err = db.Save(&t4).Error
		if err != nil {
			return errors.InternalServerError("something went wrong, tactic creation")
		}
		err = db.Table("tactics").Where("id = 1").
			Updates(map[string]interface{}{"name": "Attack-Minded-Midfield", "tactic_format": "4-3-3"}).Error
		err = db.Table("tactics").Where("id = 2").
			Updates(map[string]interface{}{"name": "Defensive-Midfield", "tactic_format": "4-3-3"}).Error
		err = db.Table("tactics").Where("id = 3").
			Updates(map[string]interface{}{"name": "Diamond-Midfield", "tactic_format": "4-2-2"}).Error
		err = db.Table("tactics").Where("id = 4").
			Updates(map[string]interface{}{"name": "Flat-Midfield", "tactic_format": "4-4-2"}).Error
		err = db.Exec("delete from tactics where id > 4").Error
		if err != nil {
			return errors.InternalServerError("something went wrong")
		}
	}
	if len(paymentDetails) >= 0 {
		p1 := PaymentDetail{}
		p2 := PaymentDetail{}
		p3 := PaymentDetail{}
		p4 := PaymentDetail{}
		p5 := PaymentDetail{}
		p6 := PaymentDetail{}
		p7 := PaymentDetail{}

		p1.ID = 1
		p2.ID = 2
		p3.ID = 3
		p4.ID = 4
		p5.ID = 5
		p6.ID = 6
		p6.ID = 7

		err := db.Save(&p1).Error
		err = db.Save(&p2).Error
		err = db.Save(&p3).Error
		err = db.Save(&p4).Error
		err = db.Save(&p5).Error
		err = db.Save(&p6).Error
		err = db.Save(&p7).Error
		if err != nil {
			return errors.InternalServerError("something went wrong, status creation")
		}
		err = db.Table("payment_details").Where("id = 1").
			Updates(map[string]interface{}{"description": "buy_assistant_coach"}).Error
		err = db.Table("payment_details").Where("id = 2").
			Updates(map[string]interface{}{"description": "buy_doctor"}).Error
		err = db.Table("payment_details").Where("id = 3").
			Updates(map[string]interface{}{"description": "buy_fitness_coach"}).Error
		err = db.Table("payment_details").Where("id = 4").
			Updates(map[string]interface{}{"description": "buy_talent_finder"}).Error
		err = db.Table("payment_details").Where("id = 5").
			Updates(map[string]interface{}{"description": "buy_trainer"}).Error
		err = db.Table("payment_details").Where("id = 6").
			Updates(map[string]interface{}{"description": "buy_player"}).Error
		err = db.Table("payment_details").Where("id = 7").
			Updates(map[string]interface{}{"description": "paying salaries"}).Error
		err = db.Exec("delete from payment_details where id > 6").Error
		if err != nil {
			return errors.InternalServerError("something went wrong")
		}
	}
	if len(doctorLevels) >= 0 {
		dcl1 := levels.DoctorLevel{}
		dcl2 := levels.DoctorLevel{}
		dcl3 := levels.DoctorLevel{}
		dcl4 := levels.DoctorLevel{}
		dcl5 := levels.DoctorLevel{}
		dcl6 := levels.DoctorLevel{}

		dcl1.ID = 1
		dcl2.ID = 2
		dcl3.ID = 3
		dcl4.ID = 4
		dcl5.ID = 5
		dcl6.ID = 6

		err := db.Save(&dcl1).Error
		err = db.Save(&dcl2).Error
		err = db.Save(&dcl3).Error
		err = db.Save(&dcl4).Error
		err = db.Save(&dcl5).Error
		err = db.Save(&dcl6).Error
		if err != nil {
			return errors.InternalServerError("something went wrong, status creation")
		}
		err = db.Table("doctor_levels").Where("id = 1").
			Updates(map[string]interface{}{"healing": 10, "max_team_players": nil}).Error
		err = db.Table("doctor_levels").Where("id = 2").
			Updates(map[string]interface{}{"healing": 20, "max_team_players": nil}).Error
		err = db.Table("doctor_levels").Where("id = 3").
			Updates(map[string]interface{}{"healing": 30, "max_team_players": 31}).Error
		err = db.Table("doctor_levels").Where("id = 4").
			Updates(map[string]interface{}{"healing": 40, "max_team_players": 33}).Error
		err = db.Table("doctor_levels").Where("id = 5").
			Updates(map[string]interface{}{"healing": 50, "max_team_players": 35}).Error
		err = db.Table("doctor_levels").Where("id = 6").
			Updates(map[string]interface{}{"healing": 75, "max_team_players": 37}).Error
		err = db.Exec("delete from doctor_levels where id > 6").Error
		if err != nil {
			return errors.InternalServerError("something went wrong")
		}
	}
	if len(talentFinderLevels) >= 0 {
		tfl1 := levels.TalentFinderLevel{}
		tfl2 := levels.TalentFinderLevel{}
		tfl3 := levels.TalentFinderLevel{}
		tfl4 := levels.TalentFinderLevel{}
		tfl5 := levels.TalentFinderLevel{}
		tfl6 := levels.TalentFinderLevel{}
		tfl1.ID = 1
		tfl2.ID = 2
		tfl3.ID = 3
		tfl4.ID = 4
		tfl5.ID = 5
		tfl6.ID = 6
		err := db.Save(&tfl1).Error
		err = db.Save(&tfl2).Error
		err = db.Save(&tfl3).Error
		err = db.Save(&tfl4).Error
		err = db.Save(&tfl5).Error
		err = db.Save(&tfl6).Error
		if err != nil {
			return errors.InternalServerError("something went wrong, status creation")
		}
		err = db.Table("talent_finder_levels").Where("id = 1").
			Updates(map[string]interface{}{"observe_players": 1, "observation_time": 5}).Error
		err = db.Table("talent_finder_levels").Where("id = 2").
			Updates(map[string]interface{}{"observe_players": 2, "observation_time": 6, "max_foreigner_players": 2}).Error
		err = db.Table("talent_finder_levels").Where("id = 3").
			Updates(map[string]interface{}{"observe_players": 3, "observation_time": 6, "max_foreigner_players": 3,
				"weekly_transfer_capacity": 1}).Error
		err = db.Table("talent_finder_levels").Where("id = 4").
			Updates(map[string]interface{}{"observe_players": 4, "observation_time": 8, "max_foreigner_players": 5,
				"weekly_transfer_capacity": 2, "encryption_transfer": 1}).Error
		err = db.Table("talent_finder_levels").Where("id = 5").
			Updates(map[string]interface{}{"observe_players": 5, "observation_time": 9, "max_foreigner_players": 7,
				"weekly_transfer_capacity": 4, "encryption_transfer": 1}).Error
		err = db.Table("talent_finder_levels").Where("id = 6").
			Updates(map[string]interface{}{"observe_players": 5, "observation_time": 15, "max_foreigner_players": 9,
				"weekly_transfer_capacity": 6, "encryption_transfer": 1}).Error
		err = db.Exec("delete from talent_finder_levels where id > 6").Error
		if err != nil {
			return errors.InternalServerError("something went wrong")
		}
	}
	if len(fitnessCoachLevel) >= 0 {
		fcl1 := levels.FitnessCoachLevel{}
		fcl2 := levels.FitnessCoachLevel{}
		fcl3 := levels.FitnessCoachLevel{}
		fcl4 := levels.FitnessCoachLevel{}
		fcl5 := levels.FitnessCoachLevel{}
		fcl6 := levels.FitnessCoachLevel{}
		fcl1.ID = 1
		fcl2.ID = 2
		fcl3.ID = 3
		fcl4.ID = 4
		fcl5.ID = 5
		fcl6.ID = 6

		err := db.Save(&fcl1).Error
		err = db.Save(&fcl2).Error
		err = db.Save(&fcl3).Error
		err = db.Save(&fcl4).Error
		err = db.Save(&fcl5).Error
		err = db.Save(&fcl6).Error
		if err != nil {
			return errors.InternalServerError("something went wrong, status creation")
		}
		err = db.Table("fitness_coach_levels").Where("id = 1").
			Updates(map[string]interface{}{"power_increase_per_day": 2}).Error
		err = db.Table("fitness_coach_levels").Where("id = 2").
			Updates(map[string]interface{}{"power_increase_per_day": 3}).Error
		err = db.Table("fitness_coach_levels").Where("id = 3").
			Updates(map[string]interface{}{"power_increase_per_day": 4}).Error
		err = db.Table("fitness_coach_levels").Where("id = 4").
			Updates(map[string]interface{}{"power_increase_per_day": 5}).Error
		err = db.Table("fitness_coach_levels").Where("id = 5").
			Updates(map[string]interface{}{"power_increase_per_day": 6}).Error
		err = db.Table("fitness_coach_levels").Where("id = 6").
			Updates(map[string]interface{}{"power_increase_per_day": 7}).Error
		err = db.Exec("delete from talent_finder_levels where id > 6").Error
		if err != nil {
			return errors.InternalServerError("something went wrong")
		}
	}
	if len(transportationLevel) >= 0 {
		tpl1 := levels.TransportationLevel{}
		tpl2 := levels.TransportationLevel{}
		tpl3 := levels.TransportationLevel{}
		tpl4 := levels.TransportationLevel{}
		tpl1.ID = 1
		tpl2.ID = 2
		tpl3.ID = 3
		tpl4.ID = 4
		err := db.Save(&tpl1).Error
		err = db.Save(&tpl2).Error
		err = db.Save(&tpl3).Error
		err = db.Save(&tpl4).Error
		if err != nil {
			return errors.InternalServerError("something went wrong, status creation")
		}
		err = db.Table("transportation_levels").Where("id = 1").
			Updates(map[string]interface{}{"weekly_price": 399000, "weekly_price_type": 1}).Error
		err = db.Table("transportation_levels").Where("id = 2").
			Updates(map[string]interface{}{"weekly_price": 19, "weekly_price_type": 0}).Error
		err = db.Table("transportation_levels").Where("id = 3").
			Updates(map[string]interface{}{"weekly_price": 29, "weekly_price_type": 0}).Error
		err = db.Table("transportation_levels").Where("id = 4").
			Updates(map[string]interface{}{"weekly_price": 39, "weekly_price_type": 0}).Error
		err = db.Exec("delete from transportation_levels where id > 6").Error
		if err != nil {
			return errors.InternalServerError("something went wrong")
		}
	}
	if len(shoppingLevel) >= 0 {
		spl1 := levels.ShoppingLevel{}
		spl2 := levels.ShoppingLevel{}
		spl3 := levels.ShoppingLevel{}
		spl4 := levels.ShoppingLevel{}
		spl5 := levels.ShoppingLevel{}
		spl6 := levels.ShoppingLevel{}
		spl1.ID = 1
		spl2.ID = 2
		spl3.ID = 3
		spl4.ID = 4
		spl5.ID = 5
		spl6.ID = 6
		err := db.Save(&spl1).Error
		err = db.Save(&spl2).Error
		err = db.Save(&spl3).Error
		err = db.Save(&spl4).Error
		err = db.Save(&spl5).Error
		err = db.Save(&spl6).Error
		if err != nil {
			return errors.InternalServerError("something went wrong, status creation")
		}
		err = db.Table("shopping_levels").Where("id = 1").
			Updates(map[string]interface{}{"daily_income_price": 7500, "weekly_price": 5, "weekly_price_type": 0}).Error
		err = db.Table("shopping_levels").Where("id = 2").
			Updates(map[string]interface{}{"daily_income_price": 21000, "weekly_price": 12, "weekly_price_type": 0}).Error
		err = db.Table("shopping_levels").Where("id = 3").
			Updates(map[string]interface{}{"daily_income_price": 51000, "weekly_price": 24, "weekly_price_type": 0}).Error
		err = db.Table("shopping_levels").Where("id = 4").
			Updates(map[string]interface{}{"daily_income_price": 115000, "weekly_price": 39, "weekly_price_type": 0}).Error
		err = db.Table("shopping_levels").Where("id = 5").
			Updates(map[string]interface{}{"daily_income_price": 175000, "weekly_price": 59, "weekly_price_type": 0}).Error
		err = db.Table("shopping_levels").Where("id = 6").
			Updates(map[string]interface{}{"daily_income_price": 30000, "weekly_price": 84, "weekly_price_type": 0}).Error

		err = db.Exec("delete from shopping_levels where id > 6").Error
		if err != nil {
			return errors.InternalServerError("something went wrong")
		}
	}
	if len(restaurantLevel) >= 0 {
		rstl1 := levels.RestaurantLevel{}
		rstl2 := levels.RestaurantLevel{}
		rstl3 := levels.RestaurantLevel{}
		rstl4 := levels.RestaurantLevel{}
		rstl5 := levels.RestaurantLevel{}
		rstl6 := levels.RestaurantLevel{}
		rstl1.ID = 1
		rstl2.ID = 2
		rstl3.ID = 3
		rstl4.ID = 4
		rstl5.ID = 5
		rstl6.ID = 6
		err := db.Save(&rstl1).Error
		err = db.Save(&rstl2).Error
		err = db.Save(&rstl3).Error
		err = db.Save(&rstl4).Error
		err = db.Save(&rstl5).Error
		err = db.Save(&rstl6).Error
		if err != nil {
			return errors.InternalServerError("something went wrong, restaurant creation")
		}
		err = db.Table("restaurant_levels").Where("id = 1").
			Updates(map[string]interface{}{
				"daily_income_price":     1000,
				"weekly_price":           1,
				"weekly_price_type":      0,
				"minimum_score_required": 0,
				"pleasure":               8,
			}).Error
		err = db.Table("restaurant_levels").Where("id = 2").
			Updates(map[string]interface{}{
				"daily_income_price":     4500,
				"weekly_price":           5,
				"weekly_price_type":      0,
				"minimum_score_required": 50,
				"pleasure":               16,
			}).Error
		err = db.Table("restaurant_levels").Where("id = 3").
			Updates(map[string]interface{}{
				"daily_income_price":     12500,
				"weekly_price":           9,
				"weekly_price_type":      0,
				"minimum_score_required": 23,
				"pleasure":               390,
			}).Error
		err = db.Table("restaurant_levels").Where("id = 4").
			Updates(map[string]interface{}{
				"daily_income_price":     30000,
				"weekly_price":           19,
				"weekly_price_type":      0,
				"minimum_score_required": 30,
				"pleasure":               890,
			}).Error
		err = db.Table("restaurant_levels").Where("id = 5").
			Updates(map[string]interface{}{
				"daily_income_price":     60000,
				"weekly_price":           29,
				"weekly_price_type":      0,
				"minimum_score_required": 40,
				"pleasure":               1250,
			}).Error
		err = db.Table("restaurant_levels").Where("id = 6").
			Updates(map[string]interface{}{
				"daily_income_price":     95000,
				"weekly_price":           49,
				"weekly_price_type":      0,
				"minimum_score_required": 55,
				"pleasure":               1750,
			}).Error
		err = db.Exec("delete from restaurant_levels where id > 6").Error
		if err != nil {
			return errors.InternalServerError("something went wrong")
		}
	}
	if len(parkingLevel) >= 0 {
		prkl1 := levels.ParkingLevel{}
		prkl2 := levels.ParkingLevel{}
		prkl3 := levels.ParkingLevel{}
		prkl4 := levels.ParkingLevel{}
		prkl5 := levels.ParkingLevel{}
		prkl6 := levels.ParkingLevel{}
		prkl1.ID = 1
		prkl2.ID = 2
		prkl3.ID = 3
		prkl4.ID = 4
		prkl5.ID = 5
		prkl6.ID = 6

		err := db.Save(&prkl1).Error
		err = db.Save(&prkl2).Error
		err = db.Save(&prkl3).Error
		err = db.Save(&prkl4).Error
		err = db.Save(&prkl5).Error
		err = db.Save(&prkl6).Error
		if err != nil {
			return errors.InternalServerError("something went wrong, restaurant creation")
		}
		err = db.Table("parking_levels").Where("id = 1").
			Updates(map[string]interface{}{
				"daily_income_price":     1000,
				"weekly_price":           1,
				"weekly_price_type":      1,
				"minimum_score_required": 0,
				"capacity":               8,
			}).Error
		err = db.Table("parking_levels").Where("id = 2").
			Updates(map[string]interface{}{
				"daily_income_price":     4500,
				"weekly_price":           1,
				"weekly_price_type":      0,
				"minimum_score_required": 50,
				"capacity":               8,
			}).Error
		err = db.Table("parking_levels").Where("id = 3").
			Updates(map[string]interface{}{
				"daily_income_price":     12500,
				"weekly_price":           9,
				"weekly_price_type":      1,
				"minimum_score_required": 390,
				"capacity":               6500,
			}).Error
		err = db.Table("parking_levels").Where("id = 4").
			Updates(map[string]interface{}{
				"daily_income_price":     30000,
				"weekly_price":           19,
				"weekly_price_type":      1,
				"minimum_score_required": 890,
				"capacity":               12500,
			}).Error
		err = db.Table("parking_levels").Where("id = 5").
			Updates(map[string]interface{}{
				"daily_income_price":     60000,
				"weekly_price":           29,
				"weekly_price_type":      1,
				"minimum_score_required": 1250,
				"capacity":               19000,
			}).Error
		err = db.Table("parking_levels").Where("id = 6").
			Updates(map[string]interface{}{
				"daily_income_price":     95000,
				"weekly_price":           49,
				"weekly_price_type":      1,
				"minimum_score_required": 1750,
				"capacity":               42000,
			}).Error
		err = db.Exec("delete from parking_levels where id > 6").Error
		if err != nil {
			return errors.InternalServerError("something went wrong")
		}
	}

	return nil
}
