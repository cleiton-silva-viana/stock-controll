package position

import (
	"sync"

	"stock-controll/internal/domain/services/error/field"
)

// Positions devem ser acessadas e criadas em tempo de execução
// {
//     Admin.Key:         Admin,
//     Developer.Key:     Developer,
//     Client.Key:        Client,
//     Manager.Key:       Manager,
//     Seller.Key:        Seller,
//     Buyer.Key:         Buyer,
//     Conference.Key:    Conference,
//     HumanResource.Key: HumanResource,
// }

type Manager struct {
	positions map[string]Position
	mu        sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		positions: make(map[string]Position, 0),
	}
}

// TODO: refatorar o erro
func (pm *Manager) Position(uuid string) (*Position, error) {
	position, exists := pm.positions[uuid]
	if !exists {
		return nil, &field.FieldError{}
	}
	return &position, nil
}

func (pm *Manager) List() []Position {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	positions := make([]Position, 0, len(pm.positions))
	for _, position := range pm.positions {
		positions = append(positions, position)
	}
	return positions
}

// TODO: refatorar o erro
func (pm *Manager) Register(position Position) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	if _, exists := pm.positions[position.uuid.String()]; !exists {
		pm.positions[position.uuid.String()] = position
		return nil
	}
	return &field.FieldError{}
}

func (pm *Manager) Remove(uuid string) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	delete(pm.positions, uuid)
}

// TODO: refatorar o error
func (pm *Manager) Update(position Position) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	if _, exists := pm.positions[position.uuid.String()]; exists {
		pm.positions[position.uuid.String()] = position
		return nil
	}
	return &field.FieldError{}
}
