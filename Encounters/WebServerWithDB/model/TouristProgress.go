package model

type TouristProgress struct {
	Id     int64 `bson:"_id,omitempty" json:"id"`
	UserId int64 `bson:"userid,omitempty" json:"userid"`
	Xp     int   `bson:"xp,omitempty" json:"xp"`
	Level  int   `bson:"level,omitempty" json:"level"`
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
