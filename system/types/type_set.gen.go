package types

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// system/types/types.yaml

type (

	// ApplicationSet slice of Application
	//
	// This type is auto-generated.
	ApplicationSet []*Application

	// AttachmentSet slice of Attachment
	//
	// This type is auto-generated.
	AttachmentSet []*Attachment

	// CredentialsSet slice of Credentials
	//
	// This type is auto-generated.
	CredentialsSet []*Credentials

	// ReminderSet slice of Reminder
	//
	// This type is auto-generated.
	ReminderSet []*Reminder

	// RoleSet slice of Role
	//
	// This type is auto-generated.
	RoleSet []*Role

	// RoleMemberSet slice of RoleMember
	//
	// This type is auto-generated.
	RoleMemberSet []*RoleMember

	// SettingValueSet slice of SettingValue
	//
	// This type is auto-generated.
	SettingValueSet []*SettingValue

	// TriggerSet slice of Trigger
	//
	// This type is auto-generated.
	TriggerSet []*Trigger

	// TriggerConstraintSet slice of TriggerConstraint
	//
	// This type is auto-generated.
	TriggerConstraintSet []*TriggerConstraint

	// UserSet slice of User
	//
	// This type is auto-generated.
	UserSet []*User

	// WorkflowSet slice of Workflow
	//
	// This type is auto-generated.
	WorkflowSet []*Workflow

	// WorkflowExpressionSet slice of WorkflowExpression
	//
	// This type is auto-generated.
	WorkflowExpressionSet []*WorkflowExpression

	// WorkflowFunctionSet slice of WorkflowFunction
	//
	// This type is auto-generated.
	WorkflowFunctionSet []*WorkflowFunction

	// WorkflowPathSet slice of WorkflowPath
	//
	// This type is auto-generated.
	WorkflowPathSet []*WorkflowPath

	// WorkflowSessionSet slice of WorkflowSession
	//
	// This type is auto-generated.
	WorkflowSessionSet []*WorkflowSession

	// WorkflowStateSet slice of WorkflowState
	//
	// This type is auto-generated.
	WorkflowStateSet []*WorkflowState

	// WorkflowStepSet slice of WorkflowStep
	//
	// This type is auto-generated.
	WorkflowStepSet []*WorkflowStep
)

// Walk iterates through every slice item and calls w(Application) err
//
// This function is auto-generated.
func (set ApplicationSet) Walk(w func(*Application) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(Application) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set ApplicationSet) Filter(f func(*Application) (bool, error)) (out ApplicationSet, err error) {
	var ok bool
	out = ApplicationSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set ApplicationSet) FindByID(ID uint64) *Application {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set ApplicationSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(Attachment) err
//
// This function is auto-generated.
func (set AttachmentSet) Walk(w func(*Attachment) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(Attachment) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set AttachmentSet) Filter(f func(*Attachment) (bool, error)) (out AttachmentSet, err error) {
	var ok bool
	out = AttachmentSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set AttachmentSet) FindByID(ID uint64) *Attachment {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set AttachmentSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(Credentials) err
//
// This function is auto-generated.
func (set CredentialsSet) Walk(w func(*Credentials) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(Credentials) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set CredentialsSet) Filter(f func(*Credentials) (bool, error)) (out CredentialsSet, err error) {
	var ok bool
	out = CredentialsSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set CredentialsSet) FindByID(ID uint64) *Credentials {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set CredentialsSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(Reminder) err
//
// This function is auto-generated.
func (set ReminderSet) Walk(w func(*Reminder) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(Reminder) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set ReminderSet) Filter(f func(*Reminder) (bool, error)) (out ReminderSet, err error) {
	var ok bool
	out = ReminderSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set ReminderSet) FindByID(ID uint64) *Reminder {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set ReminderSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(Role) err
//
// This function is auto-generated.
func (set RoleSet) Walk(w func(*Role) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(Role) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set RoleSet) Filter(f func(*Role) (bool, error)) (out RoleSet, err error) {
	var ok bool
	out = RoleSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set RoleSet) FindByID(ID uint64) *Role {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set RoleSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(RoleMember) err
//
// This function is auto-generated.
func (set RoleMemberSet) Walk(w func(*RoleMember) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(RoleMember) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set RoleMemberSet) Filter(f func(*RoleMember) (bool, error)) (out RoleMemberSet, err error) {
	var ok bool
	out = RoleMemberSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// Walk iterates through every slice item and calls w(SettingValue) err
//
// This function is auto-generated.
func (set SettingValueSet) Walk(w func(*SettingValue) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(SettingValue) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set SettingValueSet) Filter(f func(*SettingValue) (bool, error)) (out SettingValueSet, err error) {
	var ok bool
	out = SettingValueSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// Walk iterates through every slice item and calls w(Trigger) err
//
// This function is auto-generated.
func (set TriggerSet) Walk(w func(*Trigger) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(Trigger) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set TriggerSet) Filter(f func(*Trigger) (bool, error)) (out TriggerSet, err error) {
	var ok bool
	out = TriggerSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set TriggerSet) FindByID(ID uint64) *Trigger {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set TriggerSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(TriggerConstraint) err
//
// This function is auto-generated.
func (set TriggerConstraintSet) Walk(w func(*TriggerConstraint) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(TriggerConstraint) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set TriggerConstraintSet) Filter(f func(*TriggerConstraint) (bool, error)) (out TriggerConstraintSet, err error) {
	var ok bool
	out = TriggerConstraintSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// Walk iterates through every slice item and calls w(User) err
//
// This function is auto-generated.
func (set UserSet) Walk(w func(*User) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(User) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set UserSet) Filter(f func(*User) (bool, error)) (out UserSet, err error) {
	var ok bool
	out = UserSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set UserSet) FindByID(ID uint64) *User {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set UserSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(Workflow) err
//
// This function is auto-generated.
func (set WorkflowSet) Walk(w func(*Workflow) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(Workflow) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set WorkflowSet) Filter(f func(*Workflow) (bool, error)) (out WorkflowSet, err error) {
	var ok bool
	out = WorkflowSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set WorkflowSet) FindByID(ID uint64) *Workflow {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set WorkflowSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(WorkflowExpression) err
//
// This function is auto-generated.
func (set WorkflowExpressionSet) Walk(w func(*WorkflowExpression) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(WorkflowExpression) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set WorkflowExpressionSet) Filter(f func(*WorkflowExpression) (bool, error)) (out WorkflowExpressionSet, err error) {
	var ok bool
	out = WorkflowExpressionSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// Walk iterates through every slice item and calls w(WorkflowFunction) err
//
// This function is auto-generated.
func (set WorkflowFunctionSet) Walk(w func(*WorkflowFunction) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(WorkflowFunction) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set WorkflowFunctionSet) Filter(f func(*WorkflowFunction) (bool, error)) (out WorkflowFunctionSet, err error) {
	var ok bool
	out = WorkflowFunctionSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// Walk iterates through every slice item and calls w(WorkflowPath) err
//
// This function is auto-generated.
func (set WorkflowPathSet) Walk(w func(*WorkflowPath) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(WorkflowPath) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set WorkflowPathSet) Filter(f func(*WorkflowPath) (bool, error)) (out WorkflowPathSet, err error) {
	var ok bool
	out = WorkflowPathSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// Walk iterates through every slice item and calls w(WorkflowSession) err
//
// This function is auto-generated.
func (set WorkflowSessionSet) Walk(w func(*WorkflowSession) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(WorkflowSession) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set WorkflowSessionSet) Filter(f func(*WorkflowSession) (bool, error)) (out WorkflowSessionSet, err error) {
	var ok bool
	out = WorkflowSessionSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set WorkflowSessionSet) FindByID(ID uint64) *WorkflowSession {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set WorkflowSessionSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(WorkflowState) err
//
// This function is auto-generated.
func (set WorkflowStateSet) Walk(w func(*WorkflowState) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(WorkflowState) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set WorkflowStateSet) Filter(f func(*WorkflowState) (bool, error)) (out WorkflowStateSet, err error) {
	var ok bool
	out = WorkflowStateSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set WorkflowStateSet) FindByID(ID uint64) *WorkflowState {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set WorkflowStateSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}

// Walk iterates through every slice item and calls w(WorkflowStep) err
//
// This function is auto-generated.
func (set WorkflowStepSet) Walk(w func(*WorkflowStep) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(WorkflowStep) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set WorkflowStepSet) Filter(f func(*WorkflowStep) (bool, error)) (out WorkflowStepSet, err error) {
	var ok bool
	out = WorkflowStepSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set WorkflowStepSet) FindByID(ID uint64) *WorkflowStep {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set WorkflowStepSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}
