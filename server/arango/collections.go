package arango

// Document collections
type ArangoDocument string

const (
	// Application collections
	APPLICATION_CONFIG_COLLECTION ArangoDocument = "ApplicationConfig"
	// World collections
	WORLD_COLLECTION ArangoDocument = "World"
	// Location collections
	LOCATION_COLLECTION ArangoDocument = "Location"
	// Faction collections
	FACTION_COLLECTION ArangoDocument = "Faction"
	// Event collections
	EVENT_COLLECTION ArangoDocument = "Event"
	// Creature collections
	CREATURE_COLLECTION ArangoDocument = "Creature"
	// Character collections
	CHARACTER_COLLECTION ArangoDocument = "Character"
	// Item collections
	ITEM_COLLECTION ArangoDocument = "Item"
)

func (d ArangoDocument) String() string {
	return string(d)
}

// Edge collections
type ArangoEdge string

const (
	// Character edge collections
	CHARACTER_CHARACTER_EDGE ArangoEdge = "Character_Character"
	CHARACTER_CREATURE_EDGE  ArangoEdge = "Character_Creature"
	CHARACTER_EVENT_EDGE     ArangoEdge = "Character_Event"
	CHARACTER_FACTION_EDGE   ArangoEdge = "Character_Faction"
	CHARACTER_LOCATION_EDGE  ArangoEdge = "Character_Location"
	CHARACTER_ITEM_EDGE      ArangoEdge = "Character_Item"
	CHARACTER_WORLD_EDGE     ArangoEdge = "Character_World"
	// Creature edge collections
	CREATURE_CREATURE_EDGE ArangoEdge = "Creature_Creature"
	CREATURE_EVENT_EDGE    ArangoEdge = "Creature_Event"
	CREATURE_FACTION_EDGE  ArangoEdge = "Creature_Faction"
	CREATURE_LOCATION_EDGE ArangoEdge = "Creature_Location"
	CREATURE_ITEM_EDGE     ArangoEdge = "Creature_Item"
	CREATURE_WORLD_EDGE    ArangoEdge = "Creature_World"
	// Faction edge collections
	FACTION_FACTION_EDGE  ArangoEdge = "Faction_Faction"
	FACTION_EVENT_EDGE    ArangoEdge = "Faction_Event"
	FACTION_LOCATION_EDGE ArangoEdge = "Faction_Location"
	FACTION_ITEM_EDGE     ArangoEdge = "Faction_Item"
	FACTION_WORLD_EDGE    ArangoEdge = "Faction_World"
	// Location edge collections
	LOCATION_LOCATION_EDGE ArangoEdge = "Location_Location"
	LOCATION_EVENT_EDGE    ArangoEdge = "Location_Event"
	LOCATION_ITEM_EDGE     ArangoEdge = "Location_Item"
	LOCATION_WORLD_EDGE    ArangoEdge = "Location_World"
	// Item edge collections
	ITEM_ITEM_EDGE  ArangoEdge = "Item_Item"
	ITEM_EVENT_EDGE ArangoEdge = "Item_Event"
	ITEM_WORLD_EDGE ArangoEdge = "Item_World"
	// World edge collections
	WORLD_WORLD_EDGE ArangoEdge = "World_World"
	WORLD_EVENT_EDGE ArangoEdge = "World_Event"
	// Event edge collections
	EVENT_EVENT_EDGE ArangoEdge = "Event_Event"
)

func (e ArangoEdge) String() string {
	return string(e)
}

// Arrays of collections, views, and indexes
var DOCUMENT_COLLECTIONS = []ArangoDocument{
	// Application collections
	APPLICATION_CONFIG_COLLECTION,
	// World collections
	WORLD_COLLECTION,
	// Location collections
	LOCATION_COLLECTION,
	// Faction collections
	FACTION_COLLECTION,
	// Event collections
	EVENT_COLLECTION,
	// Creature collections
	CREATURE_COLLECTION,
	// Character collections
	CHARACTER_COLLECTION,
	// Item collections
	ITEM_COLLECTION,
}

var EDGE_COLLECTIONS = []ArangoEdge{
	// Character edge collections
	CHARACTER_CHARACTER_EDGE,
	CHARACTER_CREATURE_EDGE,
	CHARACTER_EVENT_EDGE,
	CHARACTER_FACTION_EDGE,
	CHARACTER_LOCATION_EDGE,
	CHARACTER_ITEM_EDGE,
	CHARACTER_WORLD_EDGE,
	// Creature edge collections
	CREATURE_CREATURE_EDGE,
	CREATURE_EVENT_EDGE,
	CREATURE_FACTION_EDGE,
	CREATURE_LOCATION_EDGE,
	CREATURE_ITEM_EDGE,
	CREATURE_WORLD_EDGE,
	// Faction edge collections
	FACTION_FACTION_EDGE,
	FACTION_EVENT_EDGE,
	FACTION_LOCATION_EDGE,
	FACTION_ITEM_EDGE,
	FACTION_WORLD_EDGE,
	// Location edge collections
	LOCATION_LOCATION_EDGE,
	LOCATION_EVENT_EDGE,
	LOCATION_ITEM_EDGE,
	LOCATION_WORLD_EDGE,
	// Item edge collections
	ITEM_ITEM_EDGE,
	ITEM_EVENT_EDGE,
	ITEM_WORLD_EDGE,
	// World edge collections
	WORLD_WORLD_EDGE,
	WORLD_EVENT_EDGE,
	// Event edge collections
	EVENT_EVENT_EDGE,
}
