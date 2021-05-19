package models


type Collection struct {
	//ID       primitive.ObjectID `bson:"_id,omitempty"`
	//User uuid.UUID  `bson:"user,omitempty"`
	Name string  `bson:"name,omitempty"`
	//SavedPosts []uuid.UUID  `bson:"collection,omitempty"`
}

