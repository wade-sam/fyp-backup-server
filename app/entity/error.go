package entity

import "errors"

var ErrNotFound = errors.New("Not found")

var ErrInvalidEntity = errors.New("Invalid entity")

var ErrClientAlreadyAdded = errors.New("The client has already been added to the client")
var ErrPolicyAlreadyAdded = errors.New("The policy has already been added to the client")

var ErrInvalidBackupPlan = errors.New("Invalid backup plan")

var ErrNoNewClient = errors.New("There are no new clients")

var ErrClientCannotBeDeleted = errors.New("This client could not be deleted, please check whether it is associated with any policies")
var ErrPolicyCantBeRemoved = errors.New("This policy could not be deleted from the client")

var ErrNoNewItem = errors.New("There are no new Items")

var ErrCouldNotAddItem = errors.New("Could not add the item requested")
var ErrCouldNotUpdateItem = errors.New("Could not update item")

var ErrNoMatchingTopic = errors.New("Topic Doesn't exist")
var ErrNoSubscribersForTopic = errors.New("No current subscribers for the topic")
var ErrChildAlreadyExists = errors.New("Child already exists")
var ErrFileNotFound = errors.New("Could not find file")
var ErrFailedDirectoryScan = errors.New("Failed the reading of the directory")
