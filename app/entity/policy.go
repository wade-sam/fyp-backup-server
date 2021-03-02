package entity

type Policy struct {
	PolicyID   string
	Policyname string
	Clients    []string
	Retention  int
	State      string
	Type       string
	Fullbackup []string
	IncBackup  []string
}

func NewPolicy(policyname, backupType string, retention int, fullbackup, incrementalbackup []string, clients []string) (*Policy, error) {
	p := &Policy{
		Policyname: policyname,
		Clients:    clients,
		Retention:  retention,
		Type:       backupType,
	}

	p.AddState()
	err := p.AddBackupPlan(fullbackup, incrementalbackup)
	if err != nil {
		return nil, ErrInvalidBackupPlan
	}
	err = p.ValidatePolicy()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return p, nil
}

func (p *Policy) AddBackupPlan(fullbackup, incrementalbackup []string) error {
	if p.Type == "full" && len(fullbackup) > 0 && len(incrementalbackup) == 0 {
		p.Fullbackup = append(fullbackup)
		p.IncBackup = append(incrementalbackup)
		return nil
	} else if p.Type == "both" && len(fullbackup) > 0 && len(incrementalbackup) > 0 {
		err := checkTwoForOverlappingDays(fullbackup, incrementalbackup)
		if err != nil {
			return ErrInvalidBackupPlan
		}
		p.Fullbackup = append(fullbackup)
		p.IncBackup = append(incrementalbackup)
		return nil
	} else {
		return ErrInvalidBackupPlan
	}
}

func checkTwoForOverlappingDays(fb []string, ib []string) error {
	for _, i := range fb {
		for _, j := range ib {
			if j == i {
				return ErrInvalidBackupPlan
			}
		}
	}
	return nil
}

func checkBackupPlan(backup []string) error {
	for _, i := range backup {
		for _, j := range backup {
			if j == i {
				return ErrInvalidBackupPlan
			}
		}
	}
	return nil
}

func (p *Policy) AddClient(client string) error {
	_, err := p.GetClient(client)
	if err == nil {
		return ErrClientAlreadyAdded
	}
	p.Clients = append(p.Clients, client)
	p.AddState()
	return nil
}

func (p *Policy) RemoveClient(client string) error {
	for i, j := range p.Clients {
		if j == client {
			p.Clients = append(p.Clients[:i], p.Clients[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

func (p *Policy) GetClient(client string) (string, error) {
	for _, v := range p.Clients {
		if v == client {
			return client, nil
		}
	}
	return client, ErrNotFound
}

func (p *Policy) GetState() (string, error) {
	if p.State == "" {
		return "", ErrNotFound

	}
	return p.State, nil

}

func (p *Policy) AddState() error {
	if len(p.Clients) > 0 {
		p.State = "active"
	} else {
		p.State = "inactive"
	}
	return nil
}

func (p *Policy) ValidatePolicy() error {
	if len(p.Clients) == 0 || p.Policyname == "" || p.Type == "" || p.Retention == 0 {

		return ErrInvalidEntity
	}
	return nil
}
