package model

type TouristProgress struct {
	Id     int64 `bson:"_id,omitempty" json:"Id"`
	UserId int64 `bson:"userid,omitempty" json:"Userid"`
	Xp     int   `bson:"xp,omitempty" json:"Xp"`
	Level  int   `bson:"level,omitempty" json:"Level"`
}

type TouristProgressDto struct {
	Xp    int
	Level int
}

func AddXp(touristProgress *TouristProgress, xp int) *TouristProgress {
	touristProgress.Xp = touristProgress.Xp + xp
	touristProgress.Level = touristProgress.Xp/100 + 1
	return touristProgress
}
