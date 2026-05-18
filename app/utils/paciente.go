package utils

import (
	"database/sql"
	"fmt"
)

type Paciente struct {
	ID          int    `json:"id"`
	Nome        string `json:"nome"`
	Email       string `json:"email"`
	Idade       int    `json:"idade"`
	Diagnostico string `json:"diagnostico"`
	Telefone    string `json:"telefone"`
	CPF         string `json:"cpf"`
}

// CreatePaciente insere um novo paciente no banco de dados
func CreatePaciente(p Paciente) error {
	query := `INSERT INTO pacientes (nome, email, idade, diagnostico, telefone, cpf) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := DB.Exec(query, p.Nome, p.Email, p.Idade, p.Diagnostico, p.Telefone, p.CPF)
	if err != nil {
		return fmt.Errorf("erro ao criar paciente: %v", err)
	}
	return nil
}

// GetPacientes retorna todos os pacientes cadastrados
func GetPacientes() ([]Paciente, error) {
	query := `SELECT id, nome, email, idade, diagnostico, telefone, cpf FROM pacientes ORDER BY id DESC`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar pacientes: %v", err)
	}
	defer rows.Close()

	var pacientes []Paciente
	for rows.Next() {
		var p Paciente
		// Usando sql.NullString para campos que podem ser nulos, caso não tenham sidos inseridos corretamente.
		// Porém, como modelamos com constraints na API podemos scanear direto, mas por segurança usamos variáveis temporárias se precisar.
		err := rows.Scan(&p.ID, &p.Nome, &p.Email, &p.Idade, &p.Diagnostico, &p.Telefone, &p.CPF)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler dados do paciente: %v", err)
		}
		pacientes = append(pacientes, p)
	}
	return pacientes, nil
}

// GetPacienteByID retorna um paciente específico
func GetPacienteByID(id int) (*Paciente, error) {
	query := `SELECT id, nome, email, idade, diagnostico, telefone, cpf FROM pacientes WHERE id = $1`
	var p Paciente
	err := DB.QueryRow(query, id).Scan(&p.ID, &p.Nome, &p.Email, &p.Idade, &p.Diagnostico, &p.Telefone, &p.CPF)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Não encontrado
		}
		return nil, fmt.Errorf("erro ao buscar paciente por id: %v", err)
	}
	return &p, nil
}

// UpdatePaciente atualiza os dados de um paciente
func UpdatePaciente(p Paciente) error {
	query := `UPDATE pacientes SET nome=$1, email=$2, idade=$3, diagnostico=$4, telefone=$5, cpf=$6 WHERE id=$7`
	_, err := DB.Exec(query, p.Nome, p.Email, p.Idade, p.Diagnostico, p.Telefone, p.CPF, p.ID)
	if err != nil {
		return fmt.Errorf("erro ao atualizar paciente: %v", err)
	}
	return nil
}

// DeletePaciente exclui um paciente
func DeletePaciente(id int) error {
	query := `DELETE FROM pacientes WHERE id=$1`
	_, err := DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erro ao excluir paciente: %v", err)
	}
	return nil
}
