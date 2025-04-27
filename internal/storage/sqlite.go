package storage

import (
    "database/sql"
    "distributed_calculator/internal/models"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

type Storage struct {
    db *sql.DB
}

// NewStorage инициализирует подключение к базе данных и создает таблицы, если они еще не существуют.
func NewStorage(dbPath string) (*Storage, error) {
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return nil, err
    }

    s := &Storage{db: db}
    if err := s.init(); err != nil {
        return nil, err
    }

    return s, nil
}

// init создает таблицы для пользователей и вычислений в базе данных.
func (s *Storage) init() error {
    userTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        login TEXT UNIQUE,
        password TEXT
    );`
    calcTable := `
    CREATE TABLE IF NOT EXISTS calculations (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        expression TEXT,
        result REAL,
        FOREIGN KEY(user_id) REFERENCES users(id)
    );`
    _, err := s.db.Exec(userTable)
    if err != nil {
        return err
    }
    _, err = s.db.Exec(calcTable)
    return err
}

// CreateUser добавляет нового пользователя в базу данных.
func (s *Storage) CreateUser(login, password string) error {
    _, err := s.db.Exec("INSERT INTO users (login, password) VALUES (?, ?)", login, password)
    return err
}

// GetUserByLogin получает пользователя по его логину.
func (s *Storage) GetUserByLogin(login string) (models.User, error) {
    row := s.db.QueryRow("SELECT id, login, password FROM users WHERE login = ?", login)
    var user models.User
    err := row.Scan(&user.ID, &user.Login, &user.Password)
    return user, err
}

// SaveCalculation сохраняет результат вычисления в базе данных.
func (s *Storage) SaveCalculation(userID int, expression string, result float64) error {
    _, err := s.db.Exec("INSERT INTO calculations (user_id, expression, result) VALUES (?, ?, ?)", userID, expression, result)
    if err != nil {
        log.Printf("failed to save calculation: %v", err)
        return err
    }
    return nil
}