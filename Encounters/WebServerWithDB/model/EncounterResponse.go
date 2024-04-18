package model

type EncounterResponse struct {
	ID               int64     `bson:"_id"`
	Encounter        Encounter `bson:"encounter" `
	ChallengeDone    bool      `bson:"challengedone"`
	PictureLongitude float64   `bson:"picturelongitude"`
	PictureLatitude  float64   `bson:"picturelatitude"`
}
