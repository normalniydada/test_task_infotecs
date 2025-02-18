// Package storage содержит функции для работы с базой данных и управления
package storage

import (
	"github.com/normalniydada/test_task_infotecs/internal/models"
	"gorm.io/gorm"
)

// GetLastNTransactions получает последние N транзакций из базы данных
//
// Параметры:
//   - db (*gorm.DB): подключение к базе данных
//   - count (int): количество транзакций, которые необходимо вернуть
//
// Возвращает:
//   - []models.Transaction: массив последних N транзакций, отсортированных по убыванию времени создания
//   - error: ошибку при выполнении запроса или nil, если всё прошло успешно
//
// Логика работы:
//  1. Выполннение SQL-запроса с сортировкой `ORDER BY created_at DESC`
//  2. Ограничение количества результатов `LIMIT count`
//  3. Заполнение слайса `transactions` полученными данными
//  4. Возвращение полученных транзакций или ошибки при запросе
func GetLastNTransactions(db *gorm.DB, count int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := db.Order("created_at desc").Limit(count).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
