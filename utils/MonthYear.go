package utils

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type MonthYear time.Time

func (m *MonthYear) UnmarshalJSON(b []byte) error {
	// Убираем кавычки
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse("01-2006", s) // формат: 07-2025
	if err != nil {
		return fmt.Errorf("не удалось распарсить дату %s: %w", s, err)
	}
	*m = MonthYear(t)
	return nil
}

func (m MonthYear) MarshalJSON() ([]byte, error) {
	t := time.Time(m)
	formatted := fmt.Sprintf("\"%02d-%d\"", t.Month(), t.Year())
	return []byte(formatted), nil
}

func (m MonthYear) ToTime() time.Time {
	return time.Time(m)
}

// GORM: сохранение в БД
func (m MonthYear) Value() (driver.Value, error) {
	t := time.Time(m)
	return t.Format("2006-01-02"), nil // сохраняем как "2025-07-01"
}
