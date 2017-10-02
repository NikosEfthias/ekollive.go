package controllers

import (
	"../models"
	"../models/oddType"
	"../models/odd"
	"../models/oddfieldType"
	"../lib/db"
	"sync"
)

var oddsLock sync.Mutex

func UpsertOdds(match models.Match) {
	oddsLock.Lock()
	defer oddsLock.Unlock()
	for _, o := range match.Odds {
		//each odd
		od := &oddType.Oddtype{
			Subtype:      o.Subtype,
			Type:         o.Type,
			Typeid:       o.Typeid,
			Oddtypevalue: o.Freetext,
		}
		oddType.Model.Where(&oddType.Oddtype{
			Subtype: o.Subtype,
			Type:    o.Type,
			Typeid:  o.Typeid,
		}).FirstOrCreate(od)

		//insert oddFields
		func(od oddType.Oddtype, o models.Odd) {
			for _, of := range o.OddsField {
				//each oddfield
				odf := &oddfieldType.Oddfieldtype{
					Oddtypeid: od.Typeid,
					Type:      of.Type,
					Typeid:    of.Typeid,
				}
				oddfieldType.Model.Where(&oddfieldType.Oddfieldtype{
					Oddtypeid: od.Typeid,
					Typeid:    of.Typeid,
				}).FirstOrCreate(odf)

				//odd.Model.Where(&odd.Odd{
				//	Oddid:          o.Id,
				//	Matchid:        match.Matchid,
				//	OddFieldTypeId: odf.Typeid,
				//	OddTypeId:      od.Oddtypeid,
				//}).Assign(&odd.Odd{
				//	Oddid:          o.Id,
				//	Matchid:        match.Matchid,
				//	OddFieldTypeId: odf.Typeid,
				//	OddTypeId:      od.Oddtypeid,
				//	Odd:            of.InnerValue,
				//	Specialvalue:   o.Specialoddsvalue,
				//	Mostbalanced:   o.Mostbalanced,
				//	Active:         of.Active,
				//}).FirstOrCreate(&odd.Odd{})
				data := &odd.Odd{
					Oddid:          o.Id,
					Matchid:        match.Matchid,
					OddFieldTypeId: odf.Typeid,
					OddTypeId:      od.Oddtypeid,
					Odd:            of.InnerValue,
					Specialvalue:   o.Specialoddsvalue,
					Mostbalanced:   o.Mostbalanced,
					Active:         of.Active,
				}
				if *o.Active == 0 {
					data.Active = o.Active
					odd.Model.Where(&odd.Odd{
						Oddid: o.Id,
					}).Update("active", 0)
				}
				db.Upsert(db.DB.DB(), "odds", data)

			}
		}(*od, o)
	}
}
