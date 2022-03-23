package trailheads

// Non-exhaustive Trailhead Object
type Trailhead struct {
	Name        string     `json:"name" bson:"name"`
	Description string     `json:"description" bson:"description"`
	Directions  Directions `json:"directions" bson:"directions"`
	Elevation   int64      `json:"elevation" bson:"elevation"`
	Parking     Parking    `json:"parking" bson:"parking"`
	Permits     []Permit   `json:"permits" bson:"permits"`
	Facilities  []Facility `json:"facilities" bson:"facilities"`
	Notes       []string   `json:"notes" bson:"notes"`
}

// Coordinate Pair for Site :: Decimal Degrees
type Coordinates struct {
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}

// Driving Directions to Site including Hwy. Exit (where applicable)
type Directions struct {
	Exit    int      `json:"exit" bson:"exit"`
	Summary string   `json:"summary" bson:"summary"`
	Steps   []string `json:"steps" bson:"steps"`
}

// Parking Accommodations at Site
type Parking struct {
	Type   string `json:"type" bson:"type"`
	Amount int64  `json:"amount" bson:"amount"`
}

// Permits Required for Parking at Site
type Permit struct {
	Type   string `json:"type" bson:"type"`
	Agency string `json:"agency" bson:"agency"`
}

// Facilities Available at Site
type Facility struct {
	Type   string `json:"type" bson:"type"`
	Amount int64  `json:"amount" bson:"amount"`
}
