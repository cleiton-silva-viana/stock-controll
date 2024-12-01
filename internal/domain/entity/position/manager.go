package position

import (
	"sync"
	
	"stock-controll/internal/domain/services/validate"
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

type PositionManager struct {
	positions map[string]Position
	mu        sync.RWMutex
}

func NewPositionManager() *PositionManager {
	return &PositionManager{
		positions: make(map[string]Position, 0),
	}
}

// TODO: refatorar o erro
func (pm *PositionManager) Position(uuid string) (*Position, error) {
	position, exists := pm.positions[uuid]
	if !exists {
		return nil, &validate.FieldError{}
	}
	return &position, nil
}

func (pm *PositionManager) ListPositions() []Position {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	positions := make([]Position, 0, len(pm.positions))
	for _, position := range pm.positions {
		positions = append(positions, position)
	}
	return positions
}

// TODO: refatorar o erro
func (pm *PositionManager) RegisterPosition(position Position) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	if _, exists := pm.positions[position.uuid]; !exists {
		pm.positions[position.uuid] = position
		return nil
	}
	return &validate.FieldError{}
}

func (pm *PositionManager) RemovePosition(uuid string) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	delete(pm.positions, uuid)
}

// TODO: refatorar o error
func (pm *PositionManager) UpdatePosition(position Position) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	if _, exists := pm.positions[position.uuid]; exists {
		pm.positions[position.uuid] = position
		return nil
	}
	return &validate.FieldError{}
}
